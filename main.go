package main

import (
	"log"

	nodeType "temp11/node"

	bn128_blsPKG "temp11/bn128_bls"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/herumi/bls-eth-go-binary/bls"
	logger "github.com/inconshreveable/log15"
	viperCFG "github.com/spf13/viper"
	// "github.com/syndtr/goleveldb/leveldb/opt"
)

func main() {
	bls.Init(bls.BLS12_381)
	bls.SetETHmode(bls.EthModeDraft07)
	viper := viperCFG.New()
	viper.SetConfigFile("config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Panic("Could Not Read Config File: ", err)
	}
	nodeSecretKeyStr := viper.GetString("secretKey")
	nodeKey := new(bls.SecretKey)
	if len(nodeSecretKeyStr) != 64 {
		logger.Warn("Unable To Read Node Secret Key, Generating And Storing Random Key In Config File")
		nodeKey.SetByCSPRNG()
		viper.Set("secretKey", nodeKey.SerializeToHexStr())
		if err := viper.WriteConfig(); err != nil {
			log.Panic("Failed To Write Node Secret Key In Config File: ", err)
		}
	} else {
		err = nodeKey.DeserializeHexStr(nodeSecretKeyStr)
		if err != nil {
			log.Panic("Secret Key Seems Invalid: ", err)
		}
	}

	nodeECDSAPrivateKeyStr := viper.GetString("ecdsaPrivateKey")
	if len(nodeECDSAPrivateKeyStr) != 64 {
		logger.Error("Unable To Read Node ECDSA Private Key")
		log.Panic("Unable To Read Node ECDSA Private Key")
	}
	nodeECDSAPrivateKey, err := crypto.HexToECDSA(nodeECDSAPrivateKeyStr)
	if err != nil {
		logger.Error("Invalid ECDSA Private Key", "Error", err)
		log.Panic("Invalid ECDSA Private Key:", err)
	}

	bn128_bls := bn128_blsPKG.NewBls()
	nodeRegisteredIdPrivateKeyStr := viper.GetString("registeredIdPrivateKey")
	if len(nodeRegisteredIdPrivateKeyStr) != 64 {
		logger.Error("Unable To Read Node Registered Id Private Key")
		log.Panic("Unable To Read Node Registered Id Private Key")
	}
	nodeRegisteredIdKeyPair, err := bn128_bls.NewKeyPair(nodeRegisteredIdPrivateKeyStr)
	if err != nil {
		logger.Error("Invalid Node Registered Id Private Key", "Error", err)
		log.Panic("Invalid Node Registered Id Private Key:", err)
	}

	nodeId := nodeKey.GetPublicKey().SerializeToHexStr()
	listenAddr := viper.GetString("listenAddr")
	requestAddr := viper.GetString("requestAddr")
	isBootstrapping := viper.GetBool("bootstrap")
	logger.Info("Basic Config", "NodeId", nodeId, "Bootstrapping", isBootstrapping)

	node := new(nodeType.Node)
	node.Id = nodeKey.GetPublicKey().SerializeToHexStr()
	node.ECDSAAddress = crypto.PubkeyToAddress(nodeECDSAPrivateKey.PublicKey)
	node.SetCredentials(bn128_bls, nodeKey, nodeECDSAPrivateKey, nodeRegisteredIdKeyPair)
	node.StartNode(viper, nodeType.PeerReqAndListenAddr{
		ListenAddr:  listenAddr,
		RequestAddr: requestAddr,
	}, nodeType.PeersRegisteredId{
		G1: nodeRegisteredIdKeyPair.PubKeyG1,
		G2: nodeRegisteredIdKeyPair.PubKey,
	}, isBootstrapping)
	// To Keep Main Thread Alive
	var tempChan chan bool
	<-tempChan
}
