package temp11

import (
	"bytes"
	"context"
	ecdsa "crypto/ecdsa"
	"encoding/gob"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"sort"
	"strings"
	"sync"
	bn128_blsPKG "temp11/bn128_bls"
	constants "temp11/constants"
	"temp11/requestFulfiller"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/mux"
	gorillaRPC "github.com/gorilla/rpc"
	gorillaJSON "github.com/gorilla/rpc/json"
	"github.com/herumi/bls-eth-go-binary/bls"
	logger "github.com/inconshreveable/log15"
	"github.com/mitchellh/mapstructure"
	viperCFG "github.com/spf13/viper"
	"golang.org/x/exp/maps"
)

type RequestsToFulFill struct {
	TxHash   string
	Receiver common.Address
}

type PeerReqAndListenAddr struct {
	ListenAddr  string
	RequestAddr string
}

type PeersRegisteredId struct {
	G1 [3]*big.Int
	G2 [3][2]*big.Int
}

type UnFinalizedRandDataSupportInfo struct {
	RandData  *big.Int
	Signature string
	Nonce     int
}

type FinalizedRandDataSupportInfo struct {
	SignatureForETH [3]*big.Int
}

type DataRequestInfo struct {
	FinalizedRandData             *big.Int
	LeaderNodeId                  string
	LeaderNodeSignature           string
	SignatureNonce                int
	FulFillTXHash                 string
	Receiver                      common.Address
	SupportForUnFinalizedRandData map[string]UnFinalizedRandDataSupportInfo
	SupportForFinalizedRandData   map[string]FinalizedRandDataSupportInfo
}

type Node struct {
	bn128_bls *bn128_blsPKG.BLS
	ethClient *ethclient.Client

	nodeKey                 *bls.SecretKey
	ecdsaPrivateKey         *ecdsa.PrivateKey
	nodeRegisteredIdKeyPair *bn128_blsPKG.KeyPair

	Id                   string
	ECDSAAddress         common.Address
	NodeRegisteredId     [3][2]*big.Int
	PeerList             []string
	PeersAddr            map[string]PeerReqAndListenAddr
	PeersRegisteredId    map[string]PeersRegisteredId
	PeersUsedNonces      map[string][]int
	LeaderPeerIndex      int
	NextLeaderChangeTime int64

	DataForRequest map[string]*DataRequestInfo
}

func (node *Node) StartNode(viper *viperCFG.Viper, peerAddrs PeerReqAndListenAddr, peerRegisteredId PeersRegisteredId, isBootstrapping bool) {
	client, err := ethclient.Dial(constants.EthereumNodeRPCURL)
	if err != nil {
		log.Panic("Failed To Connect To Ethereum Client:", err)
	}
	node.ethClient = client
	if isBootstrapping {
		// db, err = leveldb.OpenFile("nodeData", &opt.Options{
		// 	ErrorIfExist: true,
		// })
		// if err != nil {
		// 	log.Panic("Could Not Create Database, May Be Database Already Exists: ", err)
		// }
		bootstrappingPeers, peersRegisteredId := ParseBootstrappingPeerConfigData(node.bn128_bls, viper.GetStringMap("bootstrappingPeers"))
		bootstrappingPeers[node.Id] = peerAddrs
		peersList := maps.Keys(bootstrappingPeers)
		sort.Strings(peersList)
		peersRegisteredId[node.Id] = peerRegisteredId

		node.PeerList = peersList
		node.PeersAddr = bootstrappingPeers
		node.PeersRegisteredId = peersRegisteredId
		node.PeersUsedNonces = make(map[string][]int)
		node.LeaderPeerIndex = 0
		node.NextLeaderChangeTime = viper.GetInt64("nextLeaderChangeTime")
		node.DataForRequest = make(map[string]*DataRequestInfo)

		go node.runNode()
		node.connectToBootstrappingNodes()
		// utils.StoreStringMap(db, PeersKey, peers)
		logger.Info("Bootstrapping Done, Network Is Ready")
	} else {
		// db, err = leveldb.OpenFile("nodeData", nil)
		// if err != nil {
		// 	log.Panic("Could Not Open Or Create Database: ", err)
		// }
		// // Not Optimized Way

		existenceShowerServer := node.startExistenceShowerServer(peerAddrs.ListenAddr)
		leaderPeerListenAdder := viper.GetString("leaderPeerListenAdder")
		peersData := node.getPeersDataOrRegisterPeer(&peerAddrs, &peerRegisteredId, leaderPeerListenAdder)
		if peersData != nil {
			if _, ok := peersData.PeersAddr[node.Id]; ok {
				logger.Info("Node Registration Done")
				existenceShowerServer.Shutdown(context.Background())
				logger.Info("Existence Shower Server Is Stopped And Starting Node")

				node.PeerList = peersData.PeersList
				node.PeersAddr = peersData.PeersAddr
				node.PeersRegisteredId = peersData.PeersRegisteredId
				node.PeersUsedNonces = peersData.PeersUsedNonces
				node.LeaderPeerIndex = peersData.LeaderPeerIndex
				node.NextLeaderChangeTime = peersData.NextLeaderChangeTime
				node.DataForRequest = peersData.DataForRequest
				go node.runNode()
			} else {
				logger.Info("Node Is Not Registered")
			}
		} else {
			log.Panic("Failed To Get Peer List From Helping Nodes")
		}
	}
	// defer db.Close()
}

func (node *Node) runNode() {
	router := mux.NewRouter()
	router.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("Node Is Running\n"))
	})

	// Registering JSON-RPC Services
	rpcServer := gorillaRPC.NewServer()
	rpcServer.RegisterCodec(gorillaJSON.NewCodec(), "application/json")
	go node.manageLeaderSelection(node.NextLeaderChangeTime)
	rpcServer.RegisterService(node, "")
	router.Handle("/node", rpcServer)

	listenAddr := node.PeersAddr[node.Id].ListenAddr
	logger.Info("Node Is Listening", "ListenAddr", listenAddr)
	err := http.ListenAndServe(strings.TrimPrefix(listenAddr, "http://"), router)
	if err != nil {
		log.Panic("Something Went Wrong While Starting Server: ", err)
	}
}

func (node *Node) connectToBootstrappingNodes() {
	remainingNodes := CloneStringMap(node.PeersAddr)
	logger.Info("Trying To Connect To Bootstrapping Peers", "PeersId", maps.Keys(remainingNodes))
	sleepingMessage := fmt.Sprintf("Sleeping For %d", constants.WaitingTimeForPeersInEachRound)
	for {
		for id, peer := range remainingNodes {
			logger.Info("Pinging Peer", "PeerId", id)
			resp, err := http.Get(peer.ListenAddr)
			if err != nil {
				logger.Error("Couldn't Connecting To Peer", "PeerId", id, "ListenAddr", peer.ListenAddr, "Error", err)
			} else {
				if resp.StatusCode == http.StatusOK {
					logger.Info("Connected To Peer", "PeerId", id, "ListenAddr", peer.ListenAddr)
					resp.Body.Close()
					delete(remainingNodes, id)
				} else {
					logger.Error("Bad Request", "Response", resp)
				}
			}
			time.Sleep(time.Second * 1)
		}
		logger.Info(sleepingMessage, "RemainingNodes", maps.Keys(remainingNodes))
		if len(remainingNodes) < 1 {
			return
		}
		time.Sleep(time.Second * constants.WaitingTimeForPeersInEachRound)
	}
}

func (node *Node) startExistenceShowerServer(listenAddr string) *http.Server {
	router := mux.NewRouter()
	router.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("Node Is Running\n"))
	})
	server := &http.Server{
		Addr:    strings.TrimPrefix(listenAddr, "http://"),
		Handler: router,
	}
	logger.Info("Existence Shower Server Is Listening", "ListenAddr", listenAddr)
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				log.Panic("Something Went Wrong While Starting Server: ", err)
			}
		}
	}()
	return server
}

func (node *Node) callRequestFulFillerContract(contract *requestFulfiller.RequestFulfiller, finalizeRandData *big.Int, txHash string, receiverAddr common.Address) (string, error) {
	nonce, err := node.ethClient.PendingNonceAt(context.Background(), node.ECDSAAddress)
	if err != nil {
		logger.Error("Failed To Get Nonce Of Account", "Error", err)
		return "", err
	}
	gasPrice, err := node.ethClient.SuggestGasPrice(context.Background())
	if err != nil {
		logger.Error("Failed To Get Gas Price", "Error", err)
		return "", err
	}

	auth, _ := bind.NewKeyedTransactorWithChainID(node.ecdsaPrivateKey, big.NewInt(constants.ChainID))
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(constants.MaxGasForTransaction)
	auth.GasPrice = gasPrice

	signersPubKeyG1 := [][3]*big.Int{}
	signersPubKeyG2 := [][3][2]*big.Int{}
	signersParsedPubKey := [][4]*big.Int{}
	signatures := [][3]*big.Int{}
	for peerId, data := range node.DataForRequest[txHash].SupportForFinalizedRandData {
		signersPubKeyG1 = append(signersPubKeyG1, node.PeersRegisteredId[peerId].G1)
		signersPubKeyG2 = append(signersPubKeyG2, node.PeersRegisteredId[peerId].G2)
		signersParsedPubKey = append(signersParsedPubKey, node.bn128_bls.ParsePubKey(node.PeersRegisteredId[peerId].G2))
		signatures = append(signatures, data.SignatureForETH)
	}
	aggregatedPubKeyG1, aggregatedPubKeyG2, err := node.bn128_bls.AggregatePubKeys(signersPubKeyG1, signersPubKeyG2)
	if err != nil {
		logger.Error("Failed To Aggregate PubKeys", "Error", err)
		return "", err
	}
	aggregatedSignatures, err := node.bn128_bls.AggregateSignatures(signatures)
	if err != nil {
		logger.Error("Failed To Aggregate Signatures", "Error", err)
		return "", err
	}
	tx, err := contract.FullFillRequestUsingUnRegisteredAggPubKey(auth, signersParsedPubKey, node.bn128_bls.ParseSignature(aggregatedSignatures), node.bn128_bls.ParsePubKeyG1(aggregatedPubKeyG1), node.bn128_bls.ParsePubKey(aggregatedPubKeyG2), receiverAddr, finalizeRandData)
	if err != nil {
		logger.Error("Failed To Send Transaction", "Error", err)
		return "", err
	}
	return tx.Hash().Hex(), nil
}

func (node *Node) handleRandDataRequest(txHash string, receiverAddr common.Address) {
	logger.Info("Handling Rand Data Request", "TxHash", txHash, "Receiver", receiverAddr)
	// Requesting For RandData For RandData Request
	signatureMessageForUnFinalizedRequest, nonceForUnFinalizedRequest, err := node.generatePeerSignatureMessage(RandDataRequestSigMessage{
		TxHash:   txHash,
		Receiver: receiverAddr,
	})
	if err != nil {
		logger.Error("Failed To Generate Signature Message, Avoiding This Data Request", "TxHash", txHash)
		return
	}
	signatureForUnFinalizedRequest := node.nodeKey.SignByte(signatureMessageForUnFinalizedRequest).SerializeToHexStr()
	logger.Info("Signature Generated By Leader Node For UnFinalized Request", "Signature", signatureForUnFinalizedRequest)
	resDataForUnFinalizedRequest := make(map[string]interface{})
	node.broadCastMessage("Node.GetRandDataForRequest", RandDataForRequestArgs{
		LeaderSignature: signatureForUnFinalizedRequest,
		Nonce:           nonceForUnFinalizedRequest,
		TxHash:          txHash,
		Receiver:        receiverAddr,
	}, resDataForUnFinalizedRequest)
	node.DataForRequest[txHash] = &DataRequestInfo{
		FinalizedRandData:             big.NewInt(0),
		LeaderNodeId:                  node.PeerList[node.LeaderPeerIndex],
		SupportForUnFinalizedRandData: make(map[string]UnFinalizedRandDataSupportInfo),
		SupportForFinalizedRandData:   make(map[string]FinalizedRandDataSupportInfo),
		Receiver:                      receiverAddr,
	}
	finalizeRandData := big.NewInt(0)
	for peerId, data := range resDataForUnFinalizedRequest {
		dataInInterfacePTRType, _ := data.(*interface{})
		dataInInterfaceType := *dataInInterfacePTRType
		unparsedData, _ := dataInInterfaceType.(map[string]interface{})
		parsedData := new(RandDataForRequestReply)
		if err := mapstructure.Decode(unparsedData, parsedData); err == nil {
			logger.Info("Peer Res Data Decoded", "PeerId", peerId, "Data", parsedData)
			unFinalizedRandData, ok := new(big.Int).SetString(parsedData.UnFinalizedRandData, 10)
			if ok {
				if err := node.verifySignature(peerId, parsedData.Signature, parsedData.Nonce,
					PeerSignatureMessage{
						PlatformName: constants.PlatformName,
						Nonce:        parsedData.Nonce,
						Data: UnFinalizedRandData{
							RandData: parsedData.UnFinalizedRandData,
							TxHash:   txHash,
						},
					}); err == nil {
					finalizeRandData.Add(finalizeRandData, unFinalizedRandData)
					fmt.Println(finalizeRandData, peerId)
					node.DataForRequest[txHash].SupportForUnFinalizedRandData[peerId] = UnFinalizedRandDataSupportInfo{
						RandData:  unFinalizedRandData,
						Signature: parsedData.Signature,
						Nonce:     parsedData.Nonce,
					}
				} else {
					logger.Error("Failed To Verify Signature Of Request Supporter Peer", "PeerId", peerId, "Signature", parsedData.Signature)
				}
			} else {
				logger.Error("Failed To Convert String To BigInt")
			}
		} else {
			logger.Error("Failed To Parse request Supporter Peer Response")
		}
	}
	node.DataForRequest[txHash].FinalizedRandData = finalizeRandData
	logger.Info("Expected Finalized RandData Generated", "RandData", finalizeRandData)

	// Broadcasting Finalized RandData And UnFinalized RandData Support Data To Other Peers
	signatureMessageForFinalizedRequest, nonceForFinalizedRequest, err := node.generatePeerSignatureMessage(
		FinalizedRandData{
			RandData: finalizeRandData.String(),
			TxHash:   txHash,
		})
	if err != nil {
		logger.Error("Failed To Generate Signature Message, Avoiding This Data Request", "TxHash", txHash)
		return
	}
	signatureForFinalizedRequest := node.nodeKey.SignByte(signatureMessageForFinalizedRequest).SerializeToHexStr()
	logger.Info("Signature Generated By Leader Node For Finalized Request", "Signature", signatureForFinalizedRequest)

	node.DataForRequest[txHash].FinalizedRandData = finalizeRandData
	node.DataForRequest[txHash].LeaderNodeSignature = signatureForFinalizedRequest
	node.DataForRequest[txHash].SignatureNonce = nonceForFinalizedRequest

	resDataForFinalizedRequest := make(map[string]interface{})
	node.broadCastMessage("Node.ReceiveFinalizedRandData", ReceiveFinalizedRandDataArgs{
		FinalizedRandData:             finalizeRandData.String(),
		LeaderSignature:               signatureForFinalizedRequest,
		Nonce:                         nonceForFinalizedRequest,
		TxHash:                        txHash,
		SupportForUnFinalizedRandData: node.DataForRequest[txHash].SupportForUnFinalizedRandData,
	}, resDataForFinalizedRequest)

	logger.Info("Received Data For Finalized Request", "ResData", resDataForFinalizedRequest)
	contract, err := requestFulfiller.NewRequestFulfiller(common.HexToAddress(constants.RequestFulfillerContract), node.ethClient)
	if err != nil {
		logger.Error("Failed To Create Request Fulfiller Contract Instance", "Error", err)
		return
	}
	messageX, messageY, err := contract.GetMessageData(nil, receiverAddr, finalizeRandData)
	if err != nil {
		logger.Error("Failed To Get Signature Message", "Error", err)
		return
	}
	for peerId, data := range resDataForFinalizedRequest {
		dataInInterfacePTRType, _ := data.(*interface{})
		dataInInterfaceType := *dataInInterfacePTRType
		unparsedData, _ := dataInInterfaceType.(map[string]interface{})
		parsedData := new(ReceiveFinalizedRandDataReply)
		if err := mapstructure.Decode(unparsedData, parsedData); err == nil {
			logger.Info("Peer Res Data Decoded", "PeerId", peerId, "Data", parsedData)
			signatureBytes, err := hex.DecodeString(parsedData.SignatureForETH)
			if err != nil {
				logger.Error("Failed To Decoded Hex Signature Signature String For Ethereum", "Error", err)
				return
			}
			signatureBuffer := bytes.NewBuffer(signatureBytes)
			decoder := gob.NewDecoder(signatureBuffer)
			decodedSignature := [3]*big.Int{}
			err = decoder.Decode(&decodedSignature)
			if err != nil {
				logger.Error("Failed To encode Signature For Ethereum", "Error", err)
				return
			}
			if ok, _ := node.bn128_bls.VerifySignature(decodedSignature, node.PeersRegisteredId[peerId].G2, hex.EncodeToString(messageX), hex.EncodeToString(messageY)); ok {
				node.DataForRequest[txHash].SupportForFinalizedRandData[peerId] = FinalizedRandDataSupportInfo{
					SignatureForETH: decodedSignature,
				}
			} else {
				logger.Error("Failed To Verify Signature Of Request Supporter Peer", "PeerId", peerId, "Signature", parsedData.SignatureForETH)
			}
		} else {
			logger.Error("Failed To Parse request Supporter Peer Response")
		}
	}
	logger.Info("Finalized RandData Declared", "RandData", finalizeRandData)

	// Broadcasting Finalized RandData Support To Other Peers
	signatureMessageForFinalizedRequestDataStore, nonceForFinalizedRequestDataStore, err := node.generatePeerSignatureMessage(txHash)
	if err != nil {
		logger.Error("Failed To Generate Signature Message, Avoiding This Data Request", "TxHash", txHash)
		return
	}
	signatureForFinalizedRequestDataStore := node.nodeKey.SignByte(signatureMessageForFinalizedRequestDataStore).SerializeToHexStr()
	logger.Info("Signature Generated By Leader Node For Finalized Request Data Store", "Signature", signatureMessageForFinalizedRequestDataStore)
	resDataForFinalizedRequestDataStore := make(map[string]interface{})
	node.broadCastMessage("Node.ReceiveFinalizedRandDataSupportInfo", ReceiveFinalizedRandDataSupportInfoArgs{
		LeaderSignature:             signatureForFinalizedRequestDataStore,
		Nonce:                       nonceForFinalizedRequestDataStore,
		TxHash:                      txHash,
		SupportForFinalizedRandData: node.DataForRequest[txHash].SupportForFinalizedRandData,
	}, resDataForFinalizedRequestDataStore)
	logger.Info("Finalize Request Data Info Is Broadcasted To Peers", "ResData", resDataForFinalizedRequestDataStore)

	// Calling Request FulFiller Contract Which Will Send Data To Receiver And Manage Payments
	fulFillRequestHash, err := node.callRequestFulFillerContract(contract, finalizeRandData, txHash, receiverAddr)
	if err != nil {
		return
	}
	logger.Info("Transaction Is Sent Successfully", "TxHash", fulFillRequestHash)
	logger.Info("RandData Request Is FullFilled", "RequestTXHash", txHash, "FulFillTxHash", fulFillRequestHash)
	node.DataForRequest[txHash].FulFillTXHash = fulFillRequestHash

	// Broadcasting FulFill Request TxHash To Other Peers
	signatureMessageForReceiveFulFillRequestTxHash, nonceForReceiveFulFillRequestTxHash, err := node.generatePeerSignatureMessage(RequestAndFulFullTxHash{
		RequestTxHash:      txHash,
		FulFillRequestHash: fulFillRequestHash,
	})
	if err != nil {
		logger.Error("Failed To Generate Signature Message, Avoiding This Data Request", "TxHash", txHash)
		return
	}
	signatureForReceiveFulFillRequestTxHash := node.nodeKey.SignByte(signatureMessageForReceiveFulFillRequestTxHash).SerializeToHexStr()
	logger.Info("Signature Generated By Leader Node For Finalized Request Data Store", "Signature", signatureForReceiveFulFillRequestTxHash)
	resDataForReceiveFulFillRequestTxHash := make(map[string]interface{})
	node.broadCastMessage("Node.ReceiveFulFillRequestTxHash", ReceiveFulFillRequestTxHashArgs{
		LeaderSignature:    signatureForReceiveFulFillRequestTxHash,
		Nonce:              nonceForReceiveFulFillRequestTxHash,
		RequestTxHash:      txHash,
		FulFillRequestHash: fulFillRequestHash,
	}, resDataForReceiveFulFillRequestTxHash)
	logger.Info("FulFill Request TxHash Is Broadcasted To Peers", "ResData", resDataForReceiveFulFillRequestTxHash)
}

func (node *Node) listenRandDataRequest(stopTransactionListening chan bool) {
	client, err := ethclient.Dial(constants.EthereumNodeSocketURL)
	if err != nil {
		log.Panic("Failed To Create A Ethereum Client:", err)
	}
	contractAddress := common.HexToAddress(constants.RequestReceiverContract)
	eventTopicHash := common.HexToHash(constants.RequestReceivedEventTopicHash)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
		Topics:    [][]common.Hash{{eventTopicHash}},
	}
	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Panic("Received Error While Subscribing Event:", err)
	}
	logger.Info("Listener For Verified Data Request Transaction Is Started")
	for {
		select {
		case err := <-sub.Err():
			log.Panic("Received Error In Receiving Logs:", err)

		case logsData := <-logs:
			fmt.Println("Transaction Hash:", logsData.TxHash)
			fmt.Println("TopicHash:", logsData.Topics[0])
			fmt.Println("ReceiverAddress:", common.HexToAddress(logsData.Topics[1].Hex()))

			client, err := ethclient.Dial(constants.EthereumNodeRPCURL)
			if err != nil {
				logger.Error("Failed To Create A Ethereum Client", "Error", err)
				break
			}
			balance, err := client.BalanceAt(context.Background(), node.ECDSAAddress, nil)
			if err != nil {
				logger.Error("Failed To Get Node's ECDSA Address Balance", "Error", err)
			}
			if balance.Int64() < constants.MinETHBalanceForRequestFulfillment {
				logger.Error("Node Have Less Balance Than Minimum Balance For Request FulFillment")
				break
			}
			go node.handleRandDataRequest(logsData.TxHash.String(), common.HexToAddress(logsData.Topics[1].Hex()))

		case <-stopTransactionListening:
			logger.Info("Stopping Transaction Listening Because Leaders Effective Leading Time Has Been Changed")
			return
		}
	}
}

func (node *Node) listenRandRequestForFuture(stopTransactionListeningForFuture chan bool) {
	client, err := ethclient.Dial(constants.EthereumNodeSocketURL)
	if err != nil {
		log.Panic("Failed To Create A Ethereum Client:", err)
	}
	contractAddress := common.HexToAddress(constants.RequestReceiverContract)
	eventTopicHash := common.HexToHash(constants.RequestReceivedEventTopicHash)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
		Topics:    [][]common.Hash{{eventTopicHash}},
	}
	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Panic("Received Error While Subscribing Event:", err)
	}
	requestsToFulFull := []RequestsToFulFill{}
	logger.Info("Listener For Futures Verified Data Request Transaction Is Started")
	for {
		select {
		case err := <-sub.Err():
			log.Panic("Received Error In Receiving Logs:", err)

		case logsData := <-logs:
			fmt.Println("Transaction Hash:", logsData.TxHash)
			fmt.Println("TopicHash:", logsData.Topics[0])
			fmt.Println("ReceiverAddress:", common.HexToAddress(logsData.Topics[1].Hex()))

			client, err := ethclient.Dial(constants.EthereumNodeRPCURL)
			if err != nil {
				logger.Error("Failed To Create A Ethereum Client", "Error", err)
				break
			}
			balance, err := client.BalanceAt(context.Background(), node.ECDSAAddress, nil)
			if err != nil {
				logger.Error("Failed To Get Node's ECDSA Address Balance", "Error", err)
			}
			if balance.Int64() < constants.MinETHBalanceForRequestFulfillment {
				logger.Error("Node Have Less Balance Than Minimum Balance For Request FulFillment")
				break
			}
			requestsToFulFull = append(requestsToFulFull, RequestsToFulFill{
				TxHash:   logsData.TxHash.String(),
				Receiver: common.HexToAddress(logsData.Topics[1].Hex()),
			})

		case <-stopTransactionListeningForFuture:
			logger.Info("Stopping Futures Transaction Listening Because Waiting Time For Old Leader Task Completion Is Over And Leader Is Changed")
			logger.Info("FulFilling Pending Requests")
			for _, data := range requestsToFulFull {
				go node.handleRandDataRequest(data.TxHash, data.Receiver)
			}
			return
		}
	}
}

func (node *Node) manageLeaderSelection(nextLeaderChangeTime int64) {
	nextLeaderChangeTimeInSec := nextLeaderChangeTime - time.Now().Unix()
	if nextLeaderChangeTimeInSec < constants.LeadingPowerReleaseDuration {
		log.Panic("Invalid NextLeaderChangeTime")
	}

	stopTransactionListening := make(chan bool)
	if node.IsLeader() {
		go node.listenRandDataRequest(stopTransactionListening)
	}
	time.Sleep(time.Second * time.Duration(nextLeaderChangeTimeInSec-constants.LeadingPowerReleaseDuration))
	if node.IsLeader() {
		stopTransactionListening <- true
	}

	nextLeaderPeerIndex := (node.LeaderPeerIndex + 1) % len(node.PeerList)
	logger.Info("Waiting For Current Leader To Complete Their Remaining Tasks", "WaitingTime", constants.LeadingPowerReleaseDuration, "CurrentLeader", node.PeerList[node.LeaderPeerIndex], "NextLeader", node.PeerList[nextLeaderPeerIndex])
	stopTransactionListeningForFuture := make(chan bool)
	if node.Id == node.PeerList[nextLeaderPeerIndex] {
		go node.listenRandRequestForFuture(stopTransactionListeningForFuture)
	}
	time.Sleep(time.Second * time.Duration(constants.LeadingPowerReleaseDuration))
	node.LeaderPeerIndex = nextLeaderPeerIndex
	node.NextLeaderChangeTime = time.Now().Unix() + constants.TotalLeadingDuration
	logger.Info("Leader Is Changed", "New Leader", node.PeerList[node.LeaderPeerIndex])
	logger.Info("Network Is In Break Time")
	time.Sleep(time.Second * time.Duration(constants.BreakTime))
	if node.Id == node.PeerList[nextLeaderPeerIndex] {
		stopTransactionListeningForFuture <- true
	}
	logger.Info("Break Time Is Over")
	for {
		stopTransactionListening := make(chan bool)
		if node.IsLeader() {
			go node.listenRandDataRequest(stopTransactionListening)
		}
		time.Sleep(time.Second * time.Duration(constants.EffectiveLeadingDuration))
		if node.IsLeader() {
			stopTransactionListening <- true
		}

		nextLeaderPeerIndex := (node.LeaderPeerIndex + 1) % len(node.PeerList)
		logger.Info("Waiting For Current Leader To Complete Their Remaining Tasks", "WaitingTime", constants.LeadingPowerReleaseDuration, "CurrentLeader", node.PeerList[node.LeaderPeerIndex], "NextLeader", node.PeerList[nextLeaderPeerIndex])
		stopTransactionListeningForFuture := make(chan bool)
		if node.Id == node.PeerList[nextLeaderPeerIndex] {
			go node.listenRandRequestForFuture(stopTransactionListeningForFuture)
		}
		time.Sleep(time.Second * time.Duration(constants.LeadingPowerReleaseDuration))
		node.LeaderPeerIndex = (node.LeaderPeerIndex + 1) % len(node.PeerList)
		node.NextLeaderChangeTime = time.Now().Unix() + constants.TotalLeadingDuration
		logger.Info("Leader Is Changed", "New Leader", node.PeerList[node.LeaderPeerIndex])
		logger.Info("Network Is In Break Time")
		time.Sleep(time.Second * time.Duration(constants.BreakTime))
		if node.Id == node.PeerList[nextLeaderPeerIndex] {
			stopTransactionListeningForFuture <- true
		}
		logger.Info("Break Time Is Over")
	}
}

func (node *Node) requestForPeerRegistration(peerAddrs *PeerReqAndListenAddr, peerRegisteredId *PeersRegisteredId, leaderPeerListenAdder string, networkInfo *PeerDataReply) *PeerDataReply {
	nodeId := node.Id
	logger.Info("Requesting For Peer Registration To Helping Peers")
	nonce := 0
	if _, ok := networkInfo.PeersUsedNonces[nodeId]; ok {
		nonce = networkInfo.PeersUsedNonces[nodeId][(len(networkInfo.PeersUsedNonces[nodeId])-1)] + 1
	}
	sleepingMessage := fmt.Sprintf("Sleeping For %d", constants.WaitingTimeForPeersInEachRound)
	signatureMessage, _ := json.Marshal(RegistrationRequestSignatureMessage{
		PlatformName:     constants.PlatformName,
		RequestMessage:   constants.RegistrationRequestMessage,
		PeerAddrs:        *peerAddrs,
		PeerRegisteredId: *peerRegisteredId,
		Nonce:            nonce,
	})
	signatureHash := crypto.Keccak256(signatureMessage)
	signature := node.nodeKey.SignByte(signatureHash)
	logger.Info("Signature For Registration Request Is Generated", "SigMessage", signature.SerializeToHexStr(), "SigMessage", signatureHash)

	req := RegisterNewPeerArgs{
		NodeId:           nodeId,
		Signature:        signature.SerializeToHexStr(),
		Nonce:            nonce,
		RequesterAddrs:   *peerAddrs,
		PeerRegisteredId: *peerRegisteredId,
		RequestMessage:   constants.RegistrationRequestMessage,
	}
	reqData, _ := gorillaJSON.EncodeClientRequest("Node.RegisterNewPeer", req)
	body := bytes.NewBuffer(reqData)
	for i := 0; i < constants.MaxRetryForPeerListRequest; i++ {
		logger.Info("Pinging Leader Peer", "RequestAddr", leaderPeerListenAdder)
		resp, err := http.Post(leaderPeerListenAdder+"/node", "application/json", body)
		if err != nil {
			logger.Error("Something Went Wrong While Sending Request", "Error", err)
		} else {
			resData := PeerDataReply{}
			err = gorillaJSON.DecodeClientResponse(resp.Body, &resData)
			if err != nil {
				log.Panic("Could Not Decode Response: ", err)
			}
			return &resData
		}
		logger.Info(sleepingMessage)
		time.Sleep(time.Second * constants.WaitingTimeForPeersInEachRound)
	}
	return nil
}

func (node *Node) getPeersDataOrRegisterPeer(peerAddrs *PeerReqAndListenAddr, peerRegisteredId *PeersRegisteredId, leaderPeerListenAdder string) *PeerDataReply {
	logger.Info("Requesting Leader Peers For Network Info")
	sleepingMessage := fmt.Sprintf("Sleeping For %d", constants.WaitingTimeForPeersInEachRound)
	req := EmptyArgs{}
	reqData, _ := gorillaJSON.EncodeClientRequest("Node.GetPeersData", req)
	body := bytes.NewBuffer(reqData)
	for i := 0; i < constants.MaxRetryForPeerListRequest; i++ {
		logger.Info("Pinging Leader Peer", "RequestAddr", leaderPeerListenAdder)
		resp, err := http.Post(leaderPeerListenAdder+"/node", "application/json", body)
		if err != nil {
			logger.Error("Something Went Wrong While Sending Request", "Error", err)
		} else {
			resData := &PeerDataReply{}
			err = gorillaJSON.DecodeClientResponse(resp.Body, resData)
			if err != nil {
				log.Panic("Could Not Decode Response: ", err)
			}
			if _, ok := resData.PeersAddr[node.Id]; ok {
				return resData
			} else {
				resData = node.requestForPeerRegistration(peerAddrs, peerRegisteredId, leaderPeerListenAdder, resData)
			}
			return resData
		}
		logger.Info(sleepingMessage)
		time.Sleep(time.Second * constants.WaitingTimeForPeersInEachRound)
	}
	return nil
}

func (node *Node) SetCredentials(bn128_bls *bn128_blsPKG.BLS, nodeKey *bls.SecretKey, ecdsaPrivateKey *ecdsa.PrivateKey, nodeRegisteredIdPrivateKey *bn128_blsPKG.KeyPair) {
	node.bn128_bls = bn128_bls
	node.nodeKey = nodeKey
	node.ecdsaPrivateKey = ecdsaPrivateKey
	node.nodeRegisteredIdKeyPair = nodeRegisteredIdPrivateKey
}

func (node *Node) IsLeader() bool {
	return node.PeerList[node.LeaderPeerIndex] == node.Id
}

func (node *Node) validatePeerNonce(peerId string, nonce int) error {
	if _, ok := node.PeersUsedNonces[peerId]; !ok {
		if nonce != 0 {
			logger.Error("Invalid Nonce Is Used For Signature, First Nonce Must Be Zero")
			return fmt.Errorf("invalid nonce, first nonce must be 0")
		}
		node.PeersUsedNonces[peerId] = []int{0}
	} else {
		peerUsedNonces := node.PeersUsedNonces[peerId]
		totalNoncesUsed := len(peerUsedNonces)
		if peerUsedNonces[(totalNoncesUsed-1)] > nonce {
			sort.Ints(peerUsedNonces)
		}
		if sort.SearchInts(peerUsedNonces, nonce) != totalNoncesUsed {
			logger.Error("Used Nonce Is Used In Signature", "Nonce", nonce)
			return fmt.Errorf("used nonce is used in signature")
		}
		node.PeersUsedNonces[peerId] = append(peerUsedNonces, nonce)
	}
	return nil
}

func (node *Node) verifySignature(pubKeyStr string, signatureStr string, nonce int, message interface{}) error {
	signatureMessage, err := json.Marshal(message)
	if err != nil {
		logger.Error("Failed To Marshal Signature Message: %v", err)
		return fmt.Errorf("failed to marshal signature message: %v", err)
	}
	messageHash := crypto.Keccak256(signatureMessage)
	pubKey := new(bls.PublicKey)
	err = pubKey.DeserializeHexStr(pubKeyStr)
	if err != nil {
		logger.Error("PubKey Seems Invalid", "Error", err)
		return fmt.Errorf("pubKey seems invalid: %v", err)
	}
	signature := new(bls.Sign)
	err = signature.DeserializeHexStr(signatureStr)
	if err != nil {
		logger.Error("Signature Seems Invalid", "Error", err)
		return fmt.Errorf("signature seems invalid: %v", err)
	}
	if !signature.VerifyByte(pubKey, messageHash) {
		logger.Error("Signature Verification Failed", "Signature", signatureStr, "MessageHash", messageHash)
		return fmt.Errorf("signature verification failed: %v", err)
	}
	if err := node.validatePeerNonce(pubKeyStr, nonce); err != nil {
		return err
	}
	return nil
}

func (node *Node) generatePeerSignatureMessage(data interface{}) ([]byte, int, error) {
	var nonce int
	if _, ok := node.PeersUsedNonces[node.Id]; !ok {
		nonce = 0
		node.PeersUsedNonces[node.Id] = []int{0}
	} else {
		nonce = node.PeersUsedNonces[node.Id][(len(node.PeersUsedNonces[node.Id])-1)] + 1
		node.PeersUsedNonces[node.Id] = append(node.PeersUsedNonces[node.Id], nonce)
	}
	logger.Info("Message Nonce Generated", "Nonce", nonce)
	messageDataInBytes, err := json.Marshal(PeerSignatureMessage{
		PlatformName: constants.PlatformName,
		Nonce:        nonce,
		Data:         data,
	})
	if err != nil {
		logger.Error("Failed T0 Marshal Signature Message")
		return []byte{}, 0, fmt.Errorf("failed to marshal signature message: %v", err)
	}
	messageHash := crypto.Keccak256(messageDataInBytes)
	logger.Info("Message Hash Generated", "Hash", messageHash)
	return messageHash, nonce, nil
}

func (node *Node) broadCastMessage(method string, message interface{}, resData map[string]interface{}) {
	logger.Info("Broadcasting Message")
	reqData, err := gorillaJSON.EncodeClientRequest(method, message)
	wg := new(sync.WaitGroup)
	wg.Add(len(node.PeerList) - 1)
	mutex := new(sync.Mutex)
	for _, peer := range node.PeerList {
		if peer != node.Id {
			go func(wg *sync.WaitGroup, mutex *sync.Mutex, peerId string, reqBodyData []byte) {
				defer wg.Done()
				if err != nil {
					logger.Error("Failed To Encode Client Request", "Error", err)
				}
				resp, err := http.Post(node.PeersAddr[peerId].ListenAddr+"/node", "application/json", bytes.NewBuffer(reqBodyData))
				if err != nil {
					logger.Error("Failed To Send Request", "Error", err)
				} else {
					res := new(interface{})
					err = gorillaJSON.DecodeClientResponse(resp.Body, res)
					mutex.Lock()
					resData[peerId] = res
					mutex.Unlock()
					if err != nil {
						logger.Error("Request Rejected By Peer", "PeerId", peerId, "Response", err)
					} else {
						logger.Info("Data Sent To Node", "PeerId", peerId)
					}
				}
			}(wg, mutex, peer, reqData)
		}
	}
	wg.Wait()
}
