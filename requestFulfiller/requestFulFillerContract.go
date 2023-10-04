// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package requestFulfiller

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// RequestFulfillerMetaData contains all meta data concerning the RequestFulfiller contract.
var RequestFulfillerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256[4][]\",\"name\":\"members\",\"type\":\"uint256[4][]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"aggPubKeyHash\",\"type\":\"bytes32\"}],\"name\":\"AggregatedPubKeyRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"packetId\",\"type\":\"uint256\"}],\"name\":\"PacketVerified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"RequestFulFilled\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"aggregatedSignature\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[4]\",\"name\":\"registeredAggPubKeyG2\",\"type\":\"uint256[4]\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"fullFillRequestUsingRegisteredAggPubKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[4][]\",\"name\":\"signersPubKey\",\"type\":\"uint256[4][]\"},{\"internalType\":\"uint256[2]\",\"name\":\"aggregatedSignature\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"aggregatedPubKeyG1\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[4]\",\"name\":\"aggregatedPubKeyG2\",\"type\":\"uint256[4]\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"fullFillRequestUsingUnRegisteredAggPubKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"getMessageData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"hashToPoint\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"isMembers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"isRegisteredAggPubKey\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"isUsedData\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"publicKey\",\"type\":\"uint256[4]\"}],\"name\":\"isValidPublicKey\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"signature\",\"type\":\"uint256[2]\"}],\"name\":\"isValidSignature\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[4][]\",\"name\":\"signersPubKey\",\"type\":\"uint256[4][]\"},{\"internalType\":\"uint256[2]\",\"name\":\"aggregatedPubKeyG1\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[4]\",\"name\":\"aggregatedPubKeyG2\",\"type\":\"uint256[4]\"}],\"name\":\"registerAggPubKeyFromBasePubKeys\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[4][]\",\"name\":\"registeredAggPubKeys\",\"type\":\"uint256[4][]\"},{\"internalType\":\"uint256[2]\",\"name\":\"aggregatedPubKeyG1\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[4]\",\"name\":\"aggregatedPubKeyG2\",\"type\":\"uint256[4]\"}],\"name\":\"registerAggPubKeyFromRegisteredAggPubKeys\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[4][]\",\"name\":\"signersPubKey\",\"type\":\"uint256[4][]\"},{\"internalType\":\"uint256[4][]\",\"name\":\"registeredAggPubKeys\",\"type\":\"uint256[4][]\"},{\"internalType\":\"uint256[2]\",\"name\":\"aggregatedPubKeyG1\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[4]\",\"name\":\"aggregatedPubKeyG2\",\"type\":\"uint256[4]\"}],\"name\":\"registerAggPubKeyFromRegisteredAggPubKeysAndBasePubKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[4][]\",\"name\":\"members\",\"type\":\"uint256[4][]\"}],\"name\":\"removeMembers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[4][]\",\"name\":\"members\",\"type\":\"uint256[4][]\"}],\"name\":\"setNewMembers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"pubkeyG1\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[4]\",\"name\":\"pubkeyG2\",\"type\":\"uint256[4]\"}],\"name\":\"verifyHelpedAggregationPublicKeys\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"aggPubkeyG1\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[4][]\",\"name\":\"pubkeysG2\",\"type\":\"uint256[4][]\"}],\"name\":\"verifyHelpedAggregationPublicKeysMultiple\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// RequestFulfillerABI is the input ABI used to generate the binding from.
// Deprecated: Use RequestFulfillerMetaData.ABI instead.
var RequestFulfillerABI = RequestFulfillerMetaData.ABI

// RequestFulfiller is an auto generated Go binding around an Ethereum contract.
type RequestFulfiller struct {
	RequestFulfillerCaller     // Read-only binding to the contract
	RequestFulfillerTransactor // Write-only binding to the contract
	RequestFulfillerFilterer   // Log filterer for contract events
}

// RequestFulfillerCaller is an auto generated read-only Go binding around an Ethereum contract.
type RequestFulfillerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RequestFulfillerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RequestFulfillerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RequestFulfillerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RequestFulfillerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RequestFulfillerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RequestFulfillerSession struct {
	Contract     *RequestFulfiller // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RequestFulfillerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RequestFulfillerCallerSession struct {
	Contract *RequestFulfillerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// RequestFulfillerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RequestFulfillerTransactorSession struct {
	Contract     *RequestFulfillerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// RequestFulfillerRaw is an auto generated low-level Go binding around an Ethereum contract.
type RequestFulfillerRaw struct {
	Contract *RequestFulfiller // Generic contract binding to access the raw methods on
}

// RequestFulfillerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RequestFulfillerCallerRaw struct {
	Contract *RequestFulfillerCaller // Generic read-only contract binding to access the raw methods on
}

// RequestFulfillerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RequestFulfillerTransactorRaw struct {
	Contract *RequestFulfillerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRequestFulfiller creates a new instance of RequestFulfiller, bound to a specific deployed contract.
func NewRequestFulfiller(address common.Address, backend bind.ContractBackend) (*RequestFulfiller, error) {
	contract, err := bindRequestFulfiller(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RequestFulfiller{RequestFulfillerCaller: RequestFulfillerCaller{contract: contract}, RequestFulfillerTransactor: RequestFulfillerTransactor{contract: contract}, RequestFulfillerFilterer: RequestFulfillerFilterer{contract: contract}}, nil
}

// NewRequestFulfillerCaller creates a new read-only instance of RequestFulfiller, bound to a specific deployed contract.
func NewRequestFulfillerCaller(address common.Address, caller bind.ContractCaller) (*RequestFulfillerCaller, error) {
	contract, err := bindRequestFulfiller(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RequestFulfillerCaller{contract: contract}, nil
}

// NewRequestFulfillerTransactor creates a new write-only instance of RequestFulfiller, bound to a specific deployed contract.
func NewRequestFulfillerTransactor(address common.Address, transactor bind.ContractTransactor) (*RequestFulfillerTransactor, error) {
	contract, err := bindRequestFulfiller(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RequestFulfillerTransactor{contract: contract}, nil
}

// NewRequestFulfillerFilterer creates a new log filterer instance of RequestFulfiller, bound to a specific deployed contract.
func NewRequestFulfillerFilterer(address common.Address, filterer bind.ContractFilterer) (*RequestFulfillerFilterer, error) {
	contract, err := bindRequestFulfiller(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RequestFulfillerFilterer{contract: contract}, nil
}

// bindRequestFulfiller binds a generic wrapper to an already deployed contract.
func bindRequestFulfiller(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RequestFulfillerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RequestFulfiller *RequestFulfillerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RequestFulfiller.Contract.RequestFulfillerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RequestFulfiller *RequestFulfillerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RequestFulfiller.Contract.RequestFulfillerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RequestFulfiller *RequestFulfillerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RequestFulfiller.Contract.RequestFulfillerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RequestFulfiller *RequestFulfillerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RequestFulfiller.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RequestFulfiller *RequestFulfillerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RequestFulfiller.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RequestFulfiller *RequestFulfillerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RequestFulfiller.Contract.contract.Transact(opts, method, params...)
}

// GetMessageData is a free data retrieval call binding the contract method 0x27041819.
//
// Solidity: function getMessageData(address recipient, uint256 data) view returns(bytes, bytes)
func (_RequestFulfiller *RequestFulfillerCaller) GetMessageData(opts *bind.CallOpts, recipient common.Address, data *big.Int) ([]byte, []byte, error) {
	var out []interface{}
	err := _RequestFulfiller.contract.Call(opts, &out, "getMessageData", recipient, data)

	if err != nil {
		return *new([]byte), *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	out1 := *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return out0, out1, err

}

// GetMessageData is a free data retrieval call binding the contract method 0x27041819.
//
// Solidity: function getMessageData(address recipient, uint256 data) view returns(bytes, bytes)
func (_RequestFulfiller *RequestFulfillerSession) GetMessageData(recipient common.Address, data *big.Int) ([]byte, []byte, error) {
	return _RequestFulfiller.Contract.GetMessageData(&_RequestFulfiller.CallOpts, recipient, data)
}

// GetMessageData is a free data retrieval call binding the contract method 0x27041819.
//
// Solidity: function getMessageData(address recipient, uint256 data) view returns(bytes, bytes)
func (_RequestFulfiller *RequestFulfillerCallerSession) GetMessageData(recipient common.Address, data *big.Int) ([]byte, []byte, error) {
	return _RequestFulfiller.Contract.GetMessageData(&_RequestFulfiller.CallOpts, recipient, data)
}

// HashToPoint is a free data retrieval call binding the contract method 0x3033cc51.
//
// Solidity: function hashToPoint(bytes data) view returns(bytes, bytes)
func (_RequestFulfiller *RequestFulfillerCaller) HashToPoint(opts *bind.CallOpts, data []byte) ([]byte, []byte, error) {
	var out []interface{}
	err := _RequestFulfiller.contract.Call(opts, &out, "hashToPoint", data)

	if err != nil {
		return *new([]byte), *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	out1 := *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return out0, out1, err

}

// HashToPoint is a free data retrieval call binding the contract method 0x3033cc51.
//
// Solidity: function hashToPoint(bytes data) view returns(bytes, bytes)
func (_RequestFulfiller *RequestFulfillerSession) HashToPoint(data []byte) ([]byte, []byte, error) {
	return _RequestFulfiller.Contract.HashToPoint(&_RequestFulfiller.CallOpts, data)
}

// HashToPoint is a free data retrieval call binding the contract method 0x3033cc51.
//
// Solidity: function hashToPoint(bytes data) view returns(bytes, bytes)
func (_RequestFulfiller *RequestFulfillerCallerSession) HashToPoint(data []byte) ([]byte, []byte, error) {
	return _RequestFulfiller.Contract.HashToPoint(&_RequestFulfiller.CallOpts, data)
}

// IsMembers is a free data retrieval call binding the contract method 0x824eae37.
//
// Solidity: function isMembers(bytes32 ) view returns(bool)
func (_RequestFulfiller *RequestFulfillerCaller) IsMembers(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _RequestFulfiller.contract.Call(opts, &out, "isMembers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMembers is a free data retrieval call binding the contract method 0x824eae37.
//
// Solidity: function isMembers(bytes32 ) view returns(bool)
func (_RequestFulfiller *RequestFulfillerSession) IsMembers(arg0 [32]byte) (bool, error) {
	return _RequestFulfiller.Contract.IsMembers(&_RequestFulfiller.CallOpts, arg0)
}

// IsMembers is a free data retrieval call binding the contract method 0x824eae37.
//
// Solidity: function isMembers(bytes32 ) view returns(bool)
func (_RequestFulfiller *RequestFulfillerCallerSession) IsMembers(arg0 [32]byte) (bool, error) {
	return _RequestFulfiller.Contract.IsMembers(&_RequestFulfiller.CallOpts, arg0)
}

// IsRegisteredAggPubKey is a free data retrieval call binding the contract method 0x7ad7a5dc.
//
// Solidity: function isRegisteredAggPubKey(bytes32 ) view returns(bool)
func (_RequestFulfiller *RequestFulfillerCaller) IsRegisteredAggPubKey(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _RequestFulfiller.contract.Call(opts, &out, "isRegisteredAggPubKey", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRegisteredAggPubKey is a free data retrieval call binding the contract method 0x7ad7a5dc.
//
// Solidity: function isRegisteredAggPubKey(bytes32 ) view returns(bool)
func (_RequestFulfiller *RequestFulfillerSession) IsRegisteredAggPubKey(arg0 [32]byte) (bool, error) {
	return _RequestFulfiller.Contract.IsRegisteredAggPubKey(&_RequestFulfiller.CallOpts, arg0)
}

// IsRegisteredAggPubKey is a free data retrieval call binding the contract method 0x7ad7a5dc.
//
// Solidity: function isRegisteredAggPubKey(bytes32 ) view returns(bool)
func (_RequestFulfiller *RequestFulfillerCallerSession) IsRegisteredAggPubKey(arg0 [32]byte) (bool, error) {
	return _RequestFulfiller.Contract.IsRegisteredAggPubKey(&_RequestFulfiller.CallOpts, arg0)
}

// IsUsedData is a free data retrieval call binding the contract method 0xb308ab37.
//
// Solidity: function isUsedData(uint256 ) view returns(bool)
func (_RequestFulfiller *RequestFulfillerCaller) IsUsedData(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _RequestFulfiller.contract.Call(opts, &out, "isUsedData", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsUsedData is a free data retrieval call binding the contract method 0xb308ab37.
//
// Solidity: function isUsedData(uint256 ) view returns(bool)
func (_RequestFulfiller *RequestFulfillerSession) IsUsedData(arg0 *big.Int) (bool, error) {
	return _RequestFulfiller.Contract.IsUsedData(&_RequestFulfiller.CallOpts, arg0)
}

// IsUsedData is a free data retrieval call binding the contract method 0xb308ab37.
//
// Solidity: function isUsedData(uint256 ) view returns(bool)
func (_RequestFulfiller *RequestFulfillerCallerSession) IsUsedData(arg0 *big.Int) (bool, error) {
	return _RequestFulfiller.Contract.IsUsedData(&_RequestFulfiller.CallOpts, arg0)
}

// IsValidPublicKey is a free data retrieval call binding the contract method 0x6fda2c79.
//
// Solidity: function isValidPublicKey(uint256[4] publicKey) pure returns(bool)
func (_RequestFulfiller *RequestFulfillerCaller) IsValidPublicKey(opts *bind.CallOpts, publicKey [4]*big.Int) (bool, error) {
	var out []interface{}
	err := _RequestFulfiller.contract.Call(opts, &out, "isValidPublicKey", publicKey)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidPublicKey is a free data retrieval call binding the contract method 0x6fda2c79.
//
// Solidity: function isValidPublicKey(uint256[4] publicKey) pure returns(bool)
func (_RequestFulfiller *RequestFulfillerSession) IsValidPublicKey(publicKey [4]*big.Int) (bool, error) {
	return _RequestFulfiller.Contract.IsValidPublicKey(&_RequestFulfiller.CallOpts, publicKey)
}

// IsValidPublicKey is a free data retrieval call binding the contract method 0x6fda2c79.
//
// Solidity: function isValidPublicKey(uint256[4] publicKey) pure returns(bool)
func (_RequestFulfiller *RequestFulfillerCallerSession) IsValidPublicKey(publicKey [4]*big.Int) (bool, error) {
	return _RequestFulfiller.Contract.IsValidPublicKey(&_RequestFulfiller.CallOpts, publicKey)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x247dd9fb.
//
// Solidity: function isValidSignature(uint256[2] signature) pure returns(bool)
func (_RequestFulfiller *RequestFulfillerCaller) IsValidSignature(opts *bind.CallOpts, signature [2]*big.Int) (bool, error) {
	var out []interface{}
	err := _RequestFulfiller.contract.Call(opts, &out, "isValidSignature", signature)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidSignature is a free data retrieval call binding the contract method 0x247dd9fb.
//
// Solidity: function isValidSignature(uint256[2] signature) pure returns(bool)
func (_RequestFulfiller *RequestFulfillerSession) IsValidSignature(signature [2]*big.Int) (bool, error) {
	return _RequestFulfiller.Contract.IsValidSignature(&_RequestFulfiller.CallOpts, signature)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x247dd9fb.
//
// Solidity: function isValidSignature(uint256[2] signature) pure returns(bool)
func (_RequestFulfiller *RequestFulfillerCallerSession) IsValidSignature(signature [2]*big.Int) (bool, error) {
	return _RequestFulfiller.Contract.IsValidSignature(&_RequestFulfiller.CallOpts, signature)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RequestFulfiller *RequestFulfillerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RequestFulfiller.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RequestFulfiller *RequestFulfillerSession) Owner() (common.Address, error) {
	return _RequestFulfiller.Contract.Owner(&_RequestFulfiller.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RequestFulfiller *RequestFulfillerCallerSession) Owner() (common.Address, error) {
	return _RequestFulfiller.Contract.Owner(&_RequestFulfiller.CallOpts)
}

// VerifyHelpedAggregationPublicKeys is a free data retrieval call binding the contract method 0x700dc39a.
//
// Solidity: function verifyHelpedAggregationPublicKeys(uint256[2] pubkeyG1, uint256[4] pubkeyG2) view returns(bool)
func (_RequestFulfiller *RequestFulfillerCaller) VerifyHelpedAggregationPublicKeys(opts *bind.CallOpts, pubkeyG1 [2]*big.Int, pubkeyG2 [4]*big.Int) (bool, error) {
	var out []interface{}
	err := _RequestFulfiller.contract.Call(opts, &out, "verifyHelpedAggregationPublicKeys", pubkeyG1, pubkeyG2)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyHelpedAggregationPublicKeys is a free data retrieval call binding the contract method 0x700dc39a.
//
// Solidity: function verifyHelpedAggregationPublicKeys(uint256[2] pubkeyG1, uint256[4] pubkeyG2) view returns(bool)
func (_RequestFulfiller *RequestFulfillerSession) VerifyHelpedAggregationPublicKeys(pubkeyG1 [2]*big.Int, pubkeyG2 [4]*big.Int) (bool, error) {
	return _RequestFulfiller.Contract.VerifyHelpedAggregationPublicKeys(&_RequestFulfiller.CallOpts, pubkeyG1, pubkeyG2)
}

// VerifyHelpedAggregationPublicKeys is a free data retrieval call binding the contract method 0x700dc39a.
//
// Solidity: function verifyHelpedAggregationPublicKeys(uint256[2] pubkeyG1, uint256[4] pubkeyG2) view returns(bool)
func (_RequestFulfiller *RequestFulfillerCallerSession) VerifyHelpedAggregationPublicKeys(pubkeyG1 [2]*big.Int, pubkeyG2 [4]*big.Int) (bool, error) {
	return _RequestFulfiller.Contract.VerifyHelpedAggregationPublicKeys(&_RequestFulfiller.CallOpts, pubkeyG1, pubkeyG2)
}

// VerifyHelpedAggregationPublicKeysMultiple is a free data retrieval call binding the contract method 0x4b527783.
//
// Solidity: function verifyHelpedAggregationPublicKeysMultiple(uint256[2] aggPubkeyG1, uint256[4][] pubkeysG2) view returns(bool)
func (_RequestFulfiller *RequestFulfillerCaller) VerifyHelpedAggregationPublicKeysMultiple(opts *bind.CallOpts, aggPubkeyG1 [2]*big.Int, pubkeysG2 [][4]*big.Int) (bool, error) {
	var out []interface{}
	err := _RequestFulfiller.contract.Call(opts, &out, "verifyHelpedAggregationPublicKeysMultiple", aggPubkeyG1, pubkeysG2)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyHelpedAggregationPublicKeysMultiple is a free data retrieval call binding the contract method 0x4b527783.
//
// Solidity: function verifyHelpedAggregationPublicKeysMultiple(uint256[2] aggPubkeyG1, uint256[4][] pubkeysG2) view returns(bool)
func (_RequestFulfiller *RequestFulfillerSession) VerifyHelpedAggregationPublicKeysMultiple(aggPubkeyG1 [2]*big.Int, pubkeysG2 [][4]*big.Int) (bool, error) {
	return _RequestFulfiller.Contract.VerifyHelpedAggregationPublicKeysMultiple(&_RequestFulfiller.CallOpts, aggPubkeyG1, pubkeysG2)
}

// VerifyHelpedAggregationPublicKeysMultiple is a free data retrieval call binding the contract method 0x4b527783.
//
// Solidity: function verifyHelpedAggregationPublicKeysMultiple(uint256[2] aggPubkeyG1, uint256[4][] pubkeysG2) view returns(bool)
func (_RequestFulfiller *RequestFulfillerCallerSession) VerifyHelpedAggregationPublicKeysMultiple(aggPubkeyG1 [2]*big.Int, pubkeysG2 [][4]*big.Int) (bool, error) {
	return _RequestFulfiller.Contract.VerifyHelpedAggregationPublicKeysMultiple(&_RequestFulfiller.CallOpts, aggPubkeyG1, pubkeysG2)
}

// FullFillRequestUsingRegisteredAggPubKey is a paid mutator transaction binding the contract method 0x29ceecb1.
//
// Solidity: function fullFillRequestUsingRegisteredAggPubKey(uint256[2] aggregatedSignature, uint256[4] registeredAggPubKeyG2, address receiver, uint256 data) returns()
func (_RequestFulfiller *RequestFulfillerTransactor) FullFillRequestUsingRegisteredAggPubKey(opts *bind.TransactOpts, aggregatedSignature [2]*big.Int, registeredAggPubKeyG2 [4]*big.Int, receiver common.Address, data *big.Int) (*types.Transaction, error) {
	return _RequestFulfiller.contract.Transact(opts, "fullFillRequestUsingRegisteredAggPubKey", aggregatedSignature, registeredAggPubKeyG2, receiver, data)
}

// FullFillRequestUsingRegisteredAggPubKey is a paid mutator transaction binding the contract method 0x29ceecb1.
//
// Solidity: function fullFillRequestUsingRegisteredAggPubKey(uint256[2] aggregatedSignature, uint256[4] registeredAggPubKeyG2, address receiver, uint256 data) returns()
func (_RequestFulfiller *RequestFulfillerSession) FullFillRequestUsingRegisteredAggPubKey(aggregatedSignature [2]*big.Int, registeredAggPubKeyG2 [4]*big.Int, receiver common.Address, data *big.Int) (*types.Transaction, error) {
	return _RequestFulfiller.Contract.FullFillRequestUsingRegisteredAggPubKey(&_RequestFulfiller.TransactOpts, aggregatedSignature, registeredAggPubKeyG2, receiver, data)
}

// FullFillRequestUsingRegisteredAggPubKey is a paid mutator transaction binding the contract method 0x29ceecb1.
//
// Solidity: function fullFillRequestUsingRegisteredAggPubKey(uint256[2] aggregatedSignature, uint256[4] registeredAggPubKeyG2, address receiver, uint256 data) returns()
func (_RequestFulfiller *RequestFulfillerTransactorSession) FullFillRequestUsingRegisteredAggPubKey(aggregatedSignature [2]*big.Int, registeredAggPubKeyG2 [4]*big.Int, receiver common.Address, data *big.Int) (*types.Transaction, error) {
	return _RequestFulfiller.Contract.FullFillRequestUsingRegisteredAggPubKey(&_RequestFulfiller.TransactOpts, aggregatedSignature, registeredAggPubKeyG2, receiver, data)
}

// FullFillRequestUsingUnRegisteredAggPubKey is a paid mutator transaction binding the contract method 0x2f27bc89.
//
// Solidity: function fullFillRequestUsingUnRegisteredAggPubKey(uint256[4][] signersPubKey, uint256[2] aggregatedSignature, uint256[2] aggregatedPubKeyG1, uint256[4] aggregatedPubKeyG2, address receiver, uint256 data) returns()
func (_RequestFulfiller *RequestFulfillerTransactor) FullFillRequestUsingUnRegisteredAggPubKey(opts *bind.TransactOpts, signersPubKey [][4]*big.Int, aggregatedSignature [2]*big.Int, aggregatedPubKeyG1 [2]*big.Int, aggregatedPubKeyG2 [4]*big.Int, receiver common.Address, data *big.Int) (*types.Transaction, error) {
	return _RequestFulfiller.contract.Transact(opts, "fullFillRequestUsingUnRegisteredAggPubKey", signersPubKey, aggregatedSignature, aggregatedPubKeyG1, aggregatedPubKeyG2, receiver, data)
}

// FullFillRequestUsingUnRegisteredAggPubKey is a paid mutator transaction binding the contract method 0x2f27bc89.
//
// Solidity: function fullFillRequestUsingUnRegisteredAggPubKey(uint256[4][] signersPubKey, uint256[2] aggregatedSignature, uint256[2] aggregatedPubKeyG1, uint256[4] aggregatedPubKeyG2, address receiver, uint256 data) returns()
func (_RequestFulfiller *RequestFulfillerSession) FullFillRequestUsingUnRegisteredAggPubKey(signersPubKey [][4]*big.Int, aggregatedSignature [2]*big.Int, aggregatedPubKeyG1 [2]*big.Int, aggregatedPubKeyG2 [4]*big.Int, receiver common.Address, data *big.Int) (*types.Transaction, error) {
	return _RequestFulfiller.Contract.FullFillRequestUsingUnRegisteredAggPubKey(&_RequestFulfiller.TransactOpts, signersPubKey, aggregatedSignature, aggregatedPubKeyG1, aggregatedPubKeyG2, receiver, data)
}

// FullFillRequestUsingUnRegisteredAggPubKey is a paid mutator transaction binding the contract method 0x2f27bc89.
//
// Solidity: function fullFillRequestUsingUnRegisteredAggPubKey(uint256[4][] signersPubKey, uint256[2] aggregatedSignature, uint256[2] aggregatedPubKeyG1, uint256[4] aggregatedPubKeyG2, address receiver, uint256 data) returns()
func (_RequestFulfiller *RequestFulfillerTransactorSession) FullFillRequestUsingUnRegisteredAggPubKey(signersPubKey [][4]*big.Int, aggregatedSignature [2]*big.Int, aggregatedPubKeyG1 [2]*big.Int, aggregatedPubKeyG2 [4]*big.Int, receiver common.Address, data *big.Int) (*types.Transaction, error) {
	return _RequestFulfiller.Contract.FullFillRequestUsingUnRegisteredAggPubKey(&_RequestFulfiller.TransactOpts, signersPubKey, aggregatedSignature, aggregatedPubKeyG1, aggregatedPubKeyG2, receiver, data)
}

// RegisterAggPubKeyFromBasePubKeys is a paid mutator transaction binding the contract method 0x2b2f9a4f.
//
// Solidity: function registerAggPubKeyFromBasePubKeys(uint256[4][] signersPubKey, uint256[2] aggregatedPubKeyG1, uint256[4] aggregatedPubKeyG2) returns()
func (_RequestFulfiller *RequestFulfillerTransactor) RegisterAggPubKeyFromBasePubKeys(opts *bind.TransactOpts, signersPubKey [][4]*big.Int, aggregatedPubKeyG1 [2]*big.Int, aggregatedPubKeyG2 [4]*big.Int) (*types.Transaction, error) {
	return _RequestFulfiller.contract.Transact(opts, "registerAggPubKeyFromBasePubKeys", signersPubKey, aggregatedPubKeyG1, aggregatedPubKeyG2)
}

// RegisterAggPubKeyFromBasePubKeys is a paid mutator transaction binding the contract method 0x2b2f9a4f.
//
// Solidity: function registerAggPubKeyFromBasePubKeys(uint256[4][] signersPubKey, uint256[2] aggregatedPubKeyG1, uint256[4] aggregatedPubKeyG2) returns()
func (_RequestFulfiller *RequestFulfillerSession) RegisterAggPubKeyFromBasePubKeys(signersPubKey [][4]*big.Int, aggregatedPubKeyG1 [2]*big.Int, aggregatedPubKeyG2 [4]*big.Int) (*types.Transaction, error) {
	return _RequestFulfiller.Contract.RegisterAggPubKeyFromBasePubKeys(&_RequestFulfiller.TransactOpts, signersPubKey, aggregatedPubKeyG1, aggregatedPubKeyG2)
}

// RegisterAggPubKeyFromBasePubKeys is a paid mutator transaction binding the contract method 0x2b2f9a4f.
//
// Solidity: function registerAggPubKeyFromBasePubKeys(uint256[4][] signersPubKey, uint256[2] aggregatedPubKeyG1, uint256[4] aggregatedPubKeyG2) returns()
func (_RequestFulfiller *RequestFulfillerTransactorSession) RegisterAggPubKeyFromBasePubKeys(signersPubKey [][4]*big.Int, aggregatedPubKeyG1 [2]*big.Int, aggregatedPubKeyG2 [4]*big.Int) (*types.Transaction, error) {
	return _RequestFulfiller.Contract.RegisterAggPubKeyFromBasePubKeys(&_RequestFulfiller.TransactOpts, signersPubKey, aggregatedPubKeyG1, aggregatedPubKeyG2)
}

// RegisterAggPubKeyFromRegisteredAggPubKeys is a paid mutator transaction binding the contract method 0xd6ec91bf.
//
// Solidity: function registerAggPubKeyFromRegisteredAggPubKeys(uint256[4][] registeredAggPubKeys, uint256[2] aggregatedPubKeyG1, uint256[4] aggregatedPubKeyG2) returns()
func (_RequestFulfiller *RequestFulfillerTransactor) RegisterAggPubKeyFromRegisteredAggPubKeys(opts *bind.TransactOpts, registeredAggPubKeys [][4]*big.Int, aggregatedPubKeyG1 [2]*big.Int, aggregatedPubKeyG2 [4]*big.Int) (*types.Transaction, error) {
	return _RequestFulfiller.contract.Transact(opts, "registerAggPubKeyFromRegisteredAggPubKeys", registeredAggPubKeys, aggregatedPubKeyG1, aggregatedPubKeyG2)
}

// RegisterAggPubKeyFromRegisteredAggPubKeys is a paid mutator transaction binding the contract method 0xd6ec91bf.
//
// Solidity: function registerAggPubKeyFromRegisteredAggPubKeys(uint256[4][] registeredAggPubKeys, uint256[2] aggregatedPubKeyG1, uint256[4] aggregatedPubKeyG2) returns()
func (_RequestFulfiller *RequestFulfillerSession) RegisterAggPubKeyFromRegisteredAggPubKeys(registeredAggPubKeys [][4]*big.Int, aggregatedPubKeyG1 [2]*big.Int, aggregatedPubKeyG2 [4]*big.Int) (*types.Transaction, error) {
	return _RequestFulfiller.Contract.RegisterAggPubKeyFromRegisteredAggPubKeys(&_RequestFulfiller.TransactOpts, registeredAggPubKeys, aggregatedPubKeyG1, aggregatedPubKeyG2)
}

// RegisterAggPubKeyFromRegisteredAggPubKeys is a paid mutator transaction binding the contract method 0xd6ec91bf.
//
// Solidity: function registerAggPubKeyFromRegisteredAggPubKeys(uint256[4][] registeredAggPubKeys, uint256[2] aggregatedPubKeyG1, uint256[4] aggregatedPubKeyG2) returns()
func (_RequestFulfiller *RequestFulfillerTransactorSession) RegisterAggPubKeyFromRegisteredAggPubKeys(registeredAggPubKeys [][4]*big.Int, aggregatedPubKeyG1 [2]*big.Int, aggregatedPubKeyG2 [4]*big.Int) (*types.Transaction, error) {
	return _RequestFulfiller.Contract.RegisterAggPubKeyFromRegisteredAggPubKeys(&_RequestFulfiller.TransactOpts, registeredAggPubKeys, aggregatedPubKeyG1, aggregatedPubKeyG2)
}

// RegisterAggPubKeyFromRegisteredAggPubKeysAndBasePubKey is a paid mutator transaction binding the contract method 0x37773514.
//
// Solidity: function registerAggPubKeyFromRegisteredAggPubKeysAndBasePubKey(uint256[4][] signersPubKey, uint256[4][] registeredAggPubKeys, uint256[2] aggregatedPubKeyG1, uint256[4] aggregatedPubKeyG2) returns()
func (_RequestFulfiller *RequestFulfillerTransactor) RegisterAggPubKeyFromRegisteredAggPubKeysAndBasePubKey(opts *bind.TransactOpts, signersPubKey [][4]*big.Int, registeredAggPubKeys [][4]*big.Int, aggregatedPubKeyG1 [2]*big.Int, aggregatedPubKeyG2 [4]*big.Int) (*types.Transaction, error) {
	return _RequestFulfiller.contract.Transact(opts, "registerAggPubKeyFromRegisteredAggPubKeysAndBasePubKey", signersPubKey, registeredAggPubKeys, aggregatedPubKeyG1, aggregatedPubKeyG2)
}

// RegisterAggPubKeyFromRegisteredAggPubKeysAndBasePubKey is a paid mutator transaction binding the contract method 0x37773514.
//
// Solidity: function registerAggPubKeyFromRegisteredAggPubKeysAndBasePubKey(uint256[4][] signersPubKey, uint256[4][] registeredAggPubKeys, uint256[2] aggregatedPubKeyG1, uint256[4] aggregatedPubKeyG2) returns()
func (_RequestFulfiller *RequestFulfillerSession) RegisterAggPubKeyFromRegisteredAggPubKeysAndBasePubKey(signersPubKey [][4]*big.Int, registeredAggPubKeys [][4]*big.Int, aggregatedPubKeyG1 [2]*big.Int, aggregatedPubKeyG2 [4]*big.Int) (*types.Transaction, error) {
	return _RequestFulfiller.Contract.RegisterAggPubKeyFromRegisteredAggPubKeysAndBasePubKey(&_RequestFulfiller.TransactOpts, signersPubKey, registeredAggPubKeys, aggregatedPubKeyG1, aggregatedPubKeyG2)
}

// RegisterAggPubKeyFromRegisteredAggPubKeysAndBasePubKey is a paid mutator transaction binding the contract method 0x37773514.
//
// Solidity: function registerAggPubKeyFromRegisteredAggPubKeysAndBasePubKey(uint256[4][] signersPubKey, uint256[4][] registeredAggPubKeys, uint256[2] aggregatedPubKeyG1, uint256[4] aggregatedPubKeyG2) returns()
func (_RequestFulfiller *RequestFulfillerTransactorSession) RegisterAggPubKeyFromRegisteredAggPubKeysAndBasePubKey(signersPubKey [][4]*big.Int, registeredAggPubKeys [][4]*big.Int, aggregatedPubKeyG1 [2]*big.Int, aggregatedPubKeyG2 [4]*big.Int) (*types.Transaction, error) {
	return _RequestFulfiller.Contract.RegisterAggPubKeyFromRegisteredAggPubKeysAndBasePubKey(&_RequestFulfiller.TransactOpts, signersPubKey, registeredAggPubKeys, aggregatedPubKeyG1, aggregatedPubKeyG2)
}

// RemoveMembers is a paid mutator transaction binding the contract method 0x35454cf6.
//
// Solidity: function removeMembers(uint256[4][] members) returns()
func (_RequestFulfiller *RequestFulfillerTransactor) RemoveMembers(opts *bind.TransactOpts, members [][4]*big.Int) (*types.Transaction, error) {
	return _RequestFulfiller.contract.Transact(opts, "removeMembers", members)
}

// RemoveMembers is a paid mutator transaction binding the contract method 0x35454cf6.
//
// Solidity: function removeMembers(uint256[4][] members) returns()
func (_RequestFulfiller *RequestFulfillerSession) RemoveMembers(members [][4]*big.Int) (*types.Transaction, error) {
	return _RequestFulfiller.Contract.RemoveMembers(&_RequestFulfiller.TransactOpts, members)
}

// RemoveMembers is a paid mutator transaction binding the contract method 0x35454cf6.
//
// Solidity: function removeMembers(uint256[4][] members) returns()
func (_RequestFulfiller *RequestFulfillerTransactorSession) RemoveMembers(members [][4]*big.Int) (*types.Transaction, error) {
	return _RequestFulfiller.Contract.RemoveMembers(&_RequestFulfiller.TransactOpts, members)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RequestFulfiller *RequestFulfillerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RequestFulfiller.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RequestFulfiller *RequestFulfillerSession) RenounceOwnership() (*types.Transaction, error) {
	return _RequestFulfiller.Contract.RenounceOwnership(&_RequestFulfiller.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RequestFulfiller *RequestFulfillerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _RequestFulfiller.Contract.RenounceOwnership(&_RequestFulfiller.TransactOpts)
}

// SetNewMembers is a paid mutator transaction binding the contract method 0x6cc6a88e.
//
// Solidity: function setNewMembers(uint256[4][] members) returns()
func (_RequestFulfiller *RequestFulfillerTransactor) SetNewMembers(opts *bind.TransactOpts, members [][4]*big.Int) (*types.Transaction, error) {
	return _RequestFulfiller.contract.Transact(opts, "setNewMembers", members)
}

// SetNewMembers is a paid mutator transaction binding the contract method 0x6cc6a88e.
//
// Solidity: function setNewMembers(uint256[4][] members) returns()
func (_RequestFulfiller *RequestFulfillerSession) SetNewMembers(members [][4]*big.Int) (*types.Transaction, error) {
	return _RequestFulfiller.Contract.SetNewMembers(&_RequestFulfiller.TransactOpts, members)
}

// SetNewMembers is a paid mutator transaction binding the contract method 0x6cc6a88e.
//
// Solidity: function setNewMembers(uint256[4][] members) returns()
func (_RequestFulfiller *RequestFulfillerTransactorSession) SetNewMembers(members [][4]*big.Int) (*types.Transaction, error) {
	return _RequestFulfiller.Contract.SetNewMembers(&_RequestFulfiller.TransactOpts, members)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RequestFulfiller *RequestFulfillerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _RequestFulfiller.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RequestFulfiller *RequestFulfillerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RequestFulfiller.Contract.TransferOwnership(&_RequestFulfiller.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RequestFulfiller *RequestFulfillerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RequestFulfiller.Contract.TransferOwnership(&_RequestFulfiller.TransactOpts, newOwner)
}

// RequestFulfillerAggregatedPubKeyRegisteredIterator is returned from FilterAggregatedPubKeyRegistered and is used to iterate over the raw logs and unpacked data for AggregatedPubKeyRegistered events raised by the RequestFulfiller contract.
type RequestFulfillerAggregatedPubKeyRegisteredIterator struct {
	Event *RequestFulfillerAggregatedPubKeyRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RequestFulfillerAggregatedPubKeyRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RequestFulfillerAggregatedPubKeyRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RequestFulfillerAggregatedPubKeyRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RequestFulfillerAggregatedPubKeyRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RequestFulfillerAggregatedPubKeyRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RequestFulfillerAggregatedPubKeyRegistered represents a AggregatedPubKeyRegistered event raised by the RequestFulfiller contract.
type RequestFulfillerAggregatedPubKeyRegistered struct {
	AggPubKeyHash [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAggregatedPubKeyRegistered is a free log retrieval operation binding the contract event 0x5d61eeaf19cb6e3ecae846008772c9367e5b01deb90f015956a48189be2d0edc.
//
// Solidity: event AggregatedPubKeyRegistered(bytes32 indexed aggPubKeyHash)
func (_RequestFulfiller *RequestFulfillerFilterer) FilterAggregatedPubKeyRegistered(opts *bind.FilterOpts, aggPubKeyHash [][32]byte) (*RequestFulfillerAggregatedPubKeyRegisteredIterator, error) {

	var aggPubKeyHashRule []interface{}
	for _, aggPubKeyHashItem := range aggPubKeyHash {
		aggPubKeyHashRule = append(aggPubKeyHashRule, aggPubKeyHashItem)
	}

	logs, sub, err := _RequestFulfiller.contract.FilterLogs(opts, "AggregatedPubKeyRegistered", aggPubKeyHashRule)
	if err != nil {
		return nil, err
	}
	return &RequestFulfillerAggregatedPubKeyRegisteredIterator{contract: _RequestFulfiller.contract, event: "AggregatedPubKeyRegistered", logs: logs, sub: sub}, nil
}

// WatchAggregatedPubKeyRegistered is a free log subscription operation binding the contract event 0x5d61eeaf19cb6e3ecae846008772c9367e5b01deb90f015956a48189be2d0edc.
//
// Solidity: event AggregatedPubKeyRegistered(bytes32 indexed aggPubKeyHash)
func (_RequestFulfiller *RequestFulfillerFilterer) WatchAggregatedPubKeyRegistered(opts *bind.WatchOpts, sink chan<- *RequestFulfillerAggregatedPubKeyRegistered, aggPubKeyHash [][32]byte) (event.Subscription, error) {

	var aggPubKeyHashRule []interface{}
	for _, aggPubKeyHashItem := range aggPubKeyHash {
		aggPubKeyHashRule = append(aggPubKeyHashRule, aggPubKeyHashItem)
	}

	logs, sub, err := _RequestFulfiller.contract.WatchLogs(opts, "AggregatedPubKeyRegistered", aggPubKeyHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RequestFulfillerAggregatedPubKeyRegistered)
				if err := _RequestFulfiller.contract.UnpackLog(event, "AggregatedPubKeyRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAggregatedPubKeyRegistered is a log parse operation binding the contract event 0x5d61eeaf19cb6e3ecae846008772c9367e5b01deb90f015956a48189be2d0edc.
//
// Solidity: event AggregatedPubKeyRegistered(bytes32 indexed aggPubKeyHash)
func (_RequestFulfiller *RequestFulfillerFilterer) ParseAggregatedPubKeyRegistered(log types.Log) (*RequestFulfillerAggregatedPubKeyRegistered, error) {
	event := new(RequestFulfillerAggregatedPubKeyRegistered)
	if err := _RequestFulfiller.contract.UnpackLog(event, "AggregatedPubKeyRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RequestFulfillerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the RequestFulfiller contract.
type RequestFulfillerOwnershipTransferredIterator struct {
	Event *RequestFulfillerOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RequestFulfillerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RequestFulfillerOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RequestFulfillerOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RequestFulfillerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RequestFulfillerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RequestFulfillerOwnershipTransferred represents a OwnershipTransferred event raised by the RequestFulfiller contract.
type RequestFulfillerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RequestFulfiller *RequestFulfillerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RequestFulfillerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RequestFulfiller.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RequestFulfillerOwnershipTransferredIterator{contract: _RequestFulfiller.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RequestFulfiller *RequestFulfillerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RequestFulfillerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RequestFulfiller.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RequestFulfillerOwnershipTransferred)
				if err := _RequestFulfiller.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RequestFulfiller *RequestFulfillerFilterer) ParseOwnershipTransferred(log types.Log) (*RequestFulfillerOwnershipTransferred, error) {
	event := new(RequestFulfillerOwnershipTransferred)
	if err := _RequestFulfiller.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RequestFulfillerPacketVerifiedIterator is returned from FilterPacketVerified and is used to iterate over the raw logs and unpacked data for PacketVerified events raised by the RequestFulfiller contract.
type RequestFulfillerPacketVerifiedIterator struct {
	Event *RequestFulfillerPacketVerified // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RequestFulfillerPacketVerifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RequestFulfillerPacketVerified)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RequestFulfillerPacketVerified)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RequestFulfillerPacketVerifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RequestFulfillerPacketVerifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RequestFulfillerPacketVerified represents a PacketVerified event raised by the RequestFulfiller contract.
type RequestFulfillerPacketVerified struct {
	Recipient common.Address
	PacketId  *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPacketVerified is a free log retrieval operation binding the contract event 0xf8fc62df97ba80cfffd86e8635ed59f825c620e81180e1b7c898a9121e473b6e.
//
// Solidity: event PacketVerified(address indexed recipient, uint256 packetId)
func (_RequestFulfiller *RequestFulfillerFilterer) FilterPacketVerified(opts *bind.FilterOpts, recipient []common.Address) (*RequestFulfillerPacketVerifiedIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _RequestFulfiller.contract.FilterLogs(opts, "PacketVerified", recipientRule)
	if err != nil {
		return nil, err
	}
	return &RequestFulfillerPacketVerifiedIterator{contract: _RequestFulfiller.contract, event: "PacketVerified", logs: logs, sub: sub}, nil
}

// WatchPacketVerified is a free log subscription operation binding the contract event 0xf8fc62df97ba80cfffd86e8635ed59f825c620e81180e1b7c898a9121e473b6e.
//
// Solidity: event PacketVerified(address indexed recipient, uint256 packetId)
func (_RequestFulfiller *RequestFulfillerFilterer) WatchPacketVerified(opts *bind.WatchOpts, sink chan<- *RequestFulfillerPacketVerified, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _RequestFulfiller.contract.WatchLogs(opts, "PacketVerified", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RequestFulfillerPacketVerified)
				if err := _RequestFulfiller.contract.UnpackLog(event, "PacketVerified", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePacketVerified is a log parse operation binding the contract event 0xf8fc62df97ba80cfffd86e8635ed59f825c620e81180e1b7c898a9121e473b6e.
//
// Solidity: event PacketVerified(address indexed recipient, uint256 packetId)
func (_RequestFulfiller *RequestFulfillerFilterer) ParsePacketVerified(log types.Log) (*RequestFulfillerPacketVerified, error) {
	event := new(RequestFulfillerPacketVerified)
	if err := _RequestFulfiller.contract.UnpackLog(event, "PacketVerified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RequestFulfillerRequestFulFilledIterator is returned from FilterRequestFulFilled and is used to iterate over the raw logs and unpacked data for RequestFulFilled events raised by the RequestFulfiller contract.
type RequestFulfillerRequestFulFilledIterator struct {
	Event *RequestFulfillerRequestFulFilled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RequestFulfillerRequestFulFilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RequestFulfillerRequestFulFilled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RequestFulfillerRequestFulFilled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RequestFulfillerRequestFulFilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RequestFulfillerRequestFulFilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RequestFulfillerRequestFulFilled represents a RequestFulFilled event raised by the RequestFulfiller contract.
type RequestFulfillerRequestFulFilled struct {
	Receiver common.Address
	Data     *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRequestFulFilled is a free log retrieval operation binding the contract event 0xdd0b10aecdaced9ec6193f302aa4d1ca496c897cb696bcd65398613dca929f60.
//
// Solidity: event RequestFulFilled(address indexed receiver, uint256 indexed data)
func (_RequestFulfiller *RequestFulfillerFilterer) FilterRequestFulFilled(opts *bind.FilterOpts, receiver []common.Address, data []*big.Int) (*RequestFulfillerRequestFulFilledIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var dataRule []interface{}
	for _, dataItem := range data {
		dataRule = append(dataRule, dataItem)
	}

	logs, sub, err := _RequestFulfiller.contract.FilterLogs(opts, "RequestFulFilled", receiverRule, dataRule)
	if err != nil {
		return nil, err
	}
	return &RequestFulfillerRequestFulFilledIterator{contract: _RequestFulfiller.contract, event: "RequestFulFilled", logs: logs, sub: sub}, nil
}

// WatchRequestFulFilled is a free log subscription operation binding the contract event 0xdd0b10aecdaced9ec6193f302aa4d1ca496c897cb696bcd65398613dca929f60.
//
// Solidity: event RequestFulFilled(address indexed receiver, uint256 indexed data)
func (_RequestFulfiller *RequestFulfillerFilterer) WatchRequestFulFilled(opts *bind.WatchOpts, sink chan<- *RequestFulfillerRequestFulFilled, receiver []common.Address, data []*big.Int) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var dataRule []interface{}
	for _, dataItem := range data {
		dataRule = append(dataRule, dataItem)
	}

	logs, sub, err := _RequestFulfiller.contract.WatchLogs(opts, "RequestFulFilled", receiverRule, dataRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RequestFulfillerRequestFulFilled)
				if err := _RequestFulfiller.contract.UnpackLog(event, "RequestFulFilled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRequestFulFilled is a log parse operation binding the contract event 0xdd0b10aecdaced9ec6193f302aa4d1ca496c897cb696bcd65398613dca929f60.
//
// Solidity: event RequestFulFilled(address indexed receiver, uint256 indexed data)
func (_RequestFulfiller *RequestFulfillerFilterer) ParseRequestFulFilled(log types.Log) (*RequestFulfillerRequestFulFilled, error) {
	event := new(RequestFulfillerRequestFulFilled)
	if err := _RequestFulfiller.contract.UnpackLog(event, "RequestFulFilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
