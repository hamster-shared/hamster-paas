// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// FunctionsRequest is an auto generated low-level Go binding around an user-defined struct.
type FunctionsRequest struct {
	CodeLocation    uint8
	SecretsLocation uint8
	Language        uint8
	Source          string
	Secrets         []byte
	Args            []string
}

// FunctionConsumerMetaData contains all meta data concerning the FunctionConsumer contract.
var FunctionConsumerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"EmptyArgs\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptySecrets\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptySource\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RequestIsAlreadyPending\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RequestIsNotPending\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SenderIsNotRegistry\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"result\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"err\",\"type\":\"bytes\"}],\"name\":\"OCRResponse\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"RequestFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"RequestSent\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oracleAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"}],\"name\":\"addSimulatedRequestId\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumFunctions.Location\",\"name\":\"codeLocation\",\"type\":\"uint8\"},{\"internalType\":\"enumFunctions.Location\",\"name\":\"secretsLocation\",\"type\":\"uint8\"},{\"internalType\":\"enumFunctions.CodeLanguage\",\"name\":\"language\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"source\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"secrets\",\"type\":\"bytes\"},{\"internalType\":\"string[]\",\"name\":\"args\",\"type\":\"string[]\"}],\"internalType\":\"structFunctions.Request\",\"name\":\"req\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"subscriptionId\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"gasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"}],\"name\":\"estimateCost\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"source\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"secrets\",\"type\":\"bytes\"},{\"internalType\":\"enumFunctions.Location\",\"name\":\"secretsLocation\",\"type\":\"uint8\"},{\"internalType\":\"string[]\",\"name\":\"args\",\"type\":\"string[]\"},{\"internalType\":\"uint64\",\"name\":\"subscriptionId\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"gasLimit\",\"type\":\"uint32\"}],\"name\":\"executeRequest\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDONPublicKey\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"response\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"err\",\"type\":\"bytes\"}],\"name\":\"handleOracleFulfillment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestError\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestRequestId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestResponse\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"}],\"name\":\"updateOracleAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// FunctionConsumerABI is the input ABI used to generate the binding from.
// Deprecated: Use FunctionConsumerMetaData.ABI instead.
var FunctionConsumerABI = FunctionConsumerMetaData.ABI

// FunctionConsumer is an auto generated Go binding around an Ethereum contract.
type FunctionConsumer struct {
	FunctionConsumerCaller     // Read-only binding to the contract
	FunctionConsumerTransactor // Write-only binding to the contract
	FunctionConsumerFilterer   // Log filterer for contract events
}

// FunctionConsumerCaller is an auto generated read-only Go binding around an Ethereum contract.
type FunctionConsumerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FunctionConsumerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FunctionConsumerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FunctionConsumerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FunctionConsumerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FunctionConsumerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FunctionConsumerSession struct {
	Contract     *FunctionConsumer // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FunctionConsumerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FunctionConsumerCallerSession struct {
	Contract *FunctionConsumerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// FunctionConsumerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FunctionConsumerTransactorSession struct {
	Contract     *FunctionConsumerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// FunctionConsumerRaw is an auto generated low-level Go binding around an Ethereum contract.
type FunctionConsumerRaw struct {
	Contract *FunctionConsumer // Generic contract binding to access the raw methods on
}

// FunctionConsumerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FunctionConsumerCallerRaw struct {
	Contract *FunctionConsumerCaller // Generic read-only contract binding to access the raw methods on
}

// FunctionConsumerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FunctionConsumerTransactorRaw struct {
	Contract *FunctionConsumerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFunctionConsumer creates a new instance of FunctionConsumer, bound to a specific deployed contract.
func NewFunctionConsumer(address common.Address, backend bind.ContractBackend) (*FunctionConsumer, error) {
	contract, err := bindFunctionConsumer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FunctionConsumer{FunctionConsumerCaller: FunctionConsumerCaller{contract: contract}, FunctionConsumerTransactor: FunctionConsumerTransactor{contract: contract}, FunctionConsumerFilterer: FunctionConsumerFilterer{contract: contract}}, nil
}

// NewFunctionConsumerCaller creates a new read-only instance of FunctionConsumer, bound to a specific deployed contract.
func NewFunctionConsumerCaller(address common.Address, caller bind.ContractCaller) (*FunctionConsumerCaller, error) {
	contract, err := bindFunctionConsumer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FunctionConsumerCaller{contract: contract}, nil
}

// NewFunctionConsumerTransactor creates a new write-only instance of FunctionConsumer, bound to a specific deployed contract.
func NewFunctionConsumerTransactor(address common.Address, transactor bind.ContractTransactor) (*FunctionConsumerTransactor, error) {
	contract, err := bindFunctionConsumer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FunctionConsumerTransactor{contract: contract}, nil
}

// NewFunctionConsumerFilterer creates a new log filterer instance of FunctionConsumer, bound to a specific deployed contract.
func NewFunctionConsumerFilterer(address common.Address, filterer bind.ContractFilterer) (*FunctionConsumerFilterer, error) {
	contract, err := bindFunctionConsumer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FunctionConsumerFilterer{contract: contract}, nil
}

// bindFunctionConsumer binds a generic wrapper to an already deployed contract.
func bindFunctionConsumer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FunctionConsumerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FunctionConsumer *FunctionConsumerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FunctionConsumer.Contract.FunctionConsumerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FunctionConsumer *FunctionConsumerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FunctionConsumer.Contract.FunctionConsumerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FunctionConsumer *FunctionConsumerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FunctionConsumer.Contract.FunctionConsumerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FunctionConsumer *FunctionConsumerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FunctionConsumer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FunctionConsumer *FunctionConsumerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FunctionConsumer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FunctionConsumer *FunctionConsumerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FunctionConsumer.Contract.contract.Transact(opts, method, params...)
}

// EstimateCost is a free data retrieval call binding the contract method 0xd4b39175.
//
// Solidity: function estimateCost((uint8,uint8,uint8,string,bytes,string[]) req, uint64 subscriptionId, uint32 gasLimit, uint256 gasPrice) view returns(uint96)
func (_FunctionConsumer *FunctionConsumerCaller) EstimateCost(opts *bind.CallOpts, req FunctionsRequest, subscriptionId uint64, gasLimit uint32, gasPrice *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FunctionConsumer.contract.Call(opts, &out, "estimateCost", req, subscriptionId, gasLimit, gasPrice)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimateCost is a free data retrieval call binding the contract method 0xd4b39175.
//
// Solidity: function estimateCost((uint8,uint8,uint8,string,bytes,string[]) req, uint64 subscriptionId, uint32 gasLimit, uint256 gasPrice) view returns(uint96)
func (_FunctionConsumer *FunctionConsumerSession) EstimateCost(req FunctionsRequest, subscriptionId uint64, gasLimit uint32, gasPrice *big.Int) (*big.Int, error) {
	return _FunctionConsumer.Contract.EstimateCost(&_FunctionConsumer.CallOpts, req, subscriptionId, gasLimit, gasPrice)
}

// EstimateCost is a free data retrieval call binding the contract method 0xd4b39175.
//
// Solidity: function estimateCost((uint8,uint8,uint8,string,bytes,string[]) req, uint64 subscriptionId, uint32 gasLimit, uint256 gasPrice) view returns(uint96)
func (_FunctionConsumer *FunctionConsumerCallerSession) EstimateCost(req FunctionsRequest, subscriptionId uint64, gasLimit uint32, gasPrice *big.Int) (*big.Int, error) {
	return _FunctionConsumer.Contract.EstimateCost(&_FunctionConsumer.CallOpts, req, subscriptionId, gasLimit, gasPrice)
}

// GetDONPublicKey is a free data retrieval call binding the contract method 0xd328a91e.
//
// Solidity: function getDONPublicKey() view returns(bytes)
func (_FunctionConsumer *FunctionConsumerCaller) GetDONPublicKey(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _FunctionConsumer.contract.Call(opts, &out, "getDONPublicKey")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetDONPublicKey is a free data retrieval call binding the contract method 0xd328a91e.
//
// Solidity: function getDONPublicKey() view returns(bytes)
func (_FunctionConsumer *FunctionConsumerSession) GetDONPublicKey() ([]byte, error) {
	return _FunctionConsumer.Contract.GetDONPublicKey(&_FunctionConsumer.CallOpts)
}

// GetDONPublicKey is a free data retrieval call binding the contract method 0xd328a91e.
//
// Solidity: function getDONPublicKey() view returns(bytes)
func (_FunctionConsumer *FunctionConsumerCallerSession) GetDONPublicKey() ([]byte, error) {
	return _FunctionConsumer.Contract.GetDONPublicKey(&_FunctionConsumer.CallOpts)
}

// LatestError is a free data retrieval call binding the contract method 0xfffeb84e.
//
// Solidity: function latestError() view returns(bytes)
func (_FunctionConsumer *FunctionConsumerCaller) LatestError(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _FunctionConsumer.contract.Call(opts, &out, "latestError")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// LatestError is a free data retrieval call binding the contract method 0xfffeb84e.
//
// Solidity: function latestError() view returns(bytes)
func (_FunctionConsumer *FunctionConsumerSession) LatestError() ([]byte, error) {
	return _FunctionConsumer.Contract.LatestError(&_FunctionConsumer.CallOpts)
}

// LatestError is a free data retrieval call binding the contract method 0xfffeb84e.
//
// Solidity: function latestError() view returns(bytes)
func (_FunctionConsumer *FunctionConsumerCallerSession) LatestError() ([]byte, error) {
	return _FunctionConsumer.Contract.LatestError(&_FunctionConsumer.CallOpts)
}

// LatestRequestId is a free data retrieval call binding the contract method 0x1aa46f59.
//
// Solidity: function latestRequestId() view returns(bytes32)
func (_FunctionConsumer *FunctionConsumerCaller) LatestRequestId(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _FunctionConsumer.contract.Call(opts, &out, "latestRequestId")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LatestRequestId is a free data retrieval call binding the contract method 0x1aa46f59.
//
// Solidity: function latestRequestId() view returns(bytes32)
func (_FunctionConsumer *FunctionConsumerSession) LatestRequestId() ([32]byte, error) {
	return _FunctionConsumer.Contract.LatestRequestId(&_FunctionConsumer.CallOpts)
}

// LatestRequestId is a free data retrieval call binding the contract method 0x1aa46f59.
//
// Solidity: function latestRequestId() view returns(bytes32)
func (_FunctionConsumer *FunctionConsumerCallerSession) LatestRequestId() ([32]byte, error) {
	return _FunctionConsumer.Contract.LatestRequestId(&_FunctionConsumer.CallOpts)
}

// LatestResponse is a free data retrieval call binding the contract method 0xbef3a2f0.
//
// Solidity: function latestResponse() view returns(bytes)
func (_FunctionConsumer *FunctionConsumerCaller) LatestResponse(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _FunctionConsumer.contract.Call(opts, &out, "latestResponse")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// LatestResponse is a free data retrieval call binding the contract method 0xbef3a2f0.
//
// Solidity: function latestResponse() view returns(bytes)
func (_FunctionConsumer *FunctionConsumerSession) LatestResponse() ([]byte, error) {
	return _FunctionConsumer.Contract.LatestResponse(&_FunctionConsumer.CallOpts)
}

// LatestResponse is a free data retrieval call binding the contract method 0xbef3a2f0.
//
// Solidity: function latestResponse() view returns(bytes)
func (_FunctionConsumer *FunctionConsumerCallerSession) LatestResponse() ([]byte, error) {
	return _FunctionConsumer.Contract.LatestResponse(&_FunctionConsumer.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FunctionConsumer *FunctionConsumerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FunctionConsumer.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FunctionConsumer *FunctionConsumerSession) Owner() (common.Address, error) {
	return _FunctionConsumer.Contract.Owner(&_FunctionConsumer.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FunctionConsumer *FunctionConsumerCallerSession) Owner() (common.Address, error) {
	return _FunctionConsumer.Contract.Owner(&_FunctionConsumer.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_FunctionConsumer *FunctionConsumerTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FunctionConsumer.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_FunctionConsumer *FunctionConsumerSession) AcceptOwnership() (*types.Transaction, error) {
	return _FunctionConsumer.Contract.AcceptOwnership(&_FunctionConsumer.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_FunctionConsumer *FunctionConsumerTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _FunctionConsumer.Contract.AcceptOwnership(&_FunctionConsumer.TransactOpts)
}

// AddSimulatedRequestId is a paid mutator transaction binding the contract method 0x30bda99d.
//
// Solidity: function addSimulatedRequestId(address oracleAddress, bytes32 requestId) returns()
func (_FunctionConsumer *FunctionConsumerTransactor) AddSimulatedRequestId(opts *bind.TransactOpts, oracleAddress common.Address, requestId [32]byte) (*types.Transaction, error) {
	return _FunctionConsumer.contract.Transact(opts, "addSimulatedRequestId", oracleAddress, requestId)
}

// AddSimulatedRequestId is a paid mutator transaction binding the contract method 0x30bda99d.
//
// Solidity: function addSimulatedRequestId(address oracleAddress, bytes32 requestId) returns()
func (_FunctionConsumer *FunctionConsumerSession) AddSimulatedRequestId(oracleAddress common.Address, requestId [32]byte) (*types.Transaction, error) {
	return _FunctionConsumer.Contract.AddSimulatedRequestId(&_FunctionConsumer.TransactOpts, oracleAddress, requestId)
}

// AddSimulatedRequestId is a paid mutator transaction binding the contract method 0x30bda99d.
//
// Solidity: function addSimulatedRequestId(address oracleAddress, bytes32 requestId) returns()
func (_FunctionConsumer *FunctionConsumerTransactorSession) AddSimulatedRequestId(oracleAddress common.Address, requestId [32]byte) (*types.Transaction, error) {
	return _FunctionConsumer.Contract.AddSimulatedRequestId(&_FunctionConsumer.TransactOpts, oracleAddress, requestId)
}

// ExecuteRequest is a paid mutator transaction binding the contract method 0xd22b224a.
//
// Solidity: function executeRequest(string source, bytes secrets, uint8 secretsLocation, string[] args, uint64 subscriptionId, uint32 gasLimit) returns(bytes32)
func (_FunctionConsumer *FunctionConsumerTransactor) ExecuteRequest(opts *bind.TransactOpts, source string, secrets []byte, secretsLocation uint8, args []string, subscriptionId uint64, gasLimit uint32) (*types.Transaction, error) {
	return _FunctionConsumer.contract.Transact(opts, "executeRequest", source, secrets, secretsLocation, args, subscriptionId, gasLimit)
}

// ExecuteRequest is a paid mutator transaction binding the contract method 0xd22b224a.
//
// Solidity: function executeRequest(string source, bytes secrets, uint8 secretsLocation, string[] args, uint64 subscriptionId, uint32 gasLimit) returns(bytes32)
func (_FunctionConsumer *FunctionConsumerSession) ExecuteRequest(source string, secrets []byte, secretsLocation uint8, args []string, subscriptionId uint64, gasLimit uint32) (*types.Transaction, error) {
	return _FunctionConsumer.Contract.ExecuteRequest(&_FunctionConsumer.TransactOpts, source, secrets, secretsLocation, args, subscriptionId, gasLimit)
}

// ExecuteRequest is a paid mutator transaction binding the contract method 0xd22b224a.
//
// Solidity: function executeRequest(string source, bytes secrets, uint8 secretsLocation, string[] args, uint64 subscriptionId, uint32 gasLimit) returns(bytes32)
func (_FunctionConsumer *FunctionConsumerTransactorSession) ExecuteRequest(source string, secrets []byte, secretsLocation uint8, args []string, subscriptionId uint64, gasLimit uint32) (*types.Transaction, error) {
	return _FunctionConsumer.Contract.ExecuteRequest(&_FunctionConsumer.TransactOpts, source, secrets, secretsLocation, args, subscriptionId, gasLimit)
}

// HandleOracleFulfillment is a paid mutator transaction binding the contract method 0x0ca76175.
//
// Solidity: function handleOracleFulfillment(bytes32 requestId, bytes response, bytes err) returns()
func (_FunctionConsumer *FunctionConsumerTransactor) HandleOracleFulfillment(opts *bind.TransactOpts, requestId [32]byte, response []byte, err []byte) (*types.Transaction, error) {
	return _FunctionConsumer.contract.Transact(opts, "handleOracleFulfillment", requestId, response, err)
}

// HandleOracleFulfillment is a paid mutator transaction binding the contract method 0x0ca76175.
//
// Solidity: function handleOracleFulfillment(bytes32 requestId, bytes response, bytes err) returns()
func (_FunctionConsumer *FunctionConsumerSession) HandleOracleFulfillment(requestId [32]byte, response []byte, err []byte) (*types.Transaction, error) {
	return _FunctionConsumer.Contract.HandleOracleFulfillment(&_FunctionConsumer.TransactOpts, requestId, response, err)
}

// HandleOracleFulfillment is a paid mutator transaction binding the contract method 0x0ca76175.
//
// Solidity: function handleOracleFulfillment(bytes32 requestId, bytes response, bytes err) returns()
func (_FunctionConsumer *FunctionConsumerTransactorSession) HandleOracleFulfillment(requestId [32]byte, response []byte, err []byte) (*types.Transaction, error) {
	return _FunctionConsumer.Contract.HandleOracleFulfillment(&_FunctionConsumer.TransactOpts, requestId, response, err)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address to) returns()
func (_FunctionConsumer *FunctionConsumerTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _FunctionConsumer.contract.Transact(opts, "transferOwnership", to)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address to) returns()
func (_FunctionConsumer *FunctionConsumerSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _FunctionConsumer.Contract.TransferOwnership(&_FunctionConsumer.TransactOpts, to)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address to) returns()
func (_FunctionConsumer *FunctionConsumerTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _FunctionConsumer.Contract.TransferOwnership(&_FunctionConsumer.TransactOpts, to)
}

// UpdateOracleAddress is a paid mutator transaction binding the contract method 0xf7023bb6.
//
// Solidity: function updateOracleAddress(address oracle) returns()
func (_FunctionConsumer *FunctionConsumerTransactor) UpdateOracleAddress(opts *bind.TransactOpts, oracle common.Address) (*types.Transaction, error) {
	return _FunctionConsumer.contract.Transact(opts, "updateOracleAddress", oracle)
}

// UpdateOracleAddress is a paid mutator transaction binding the contract method 0xf7023bb6.
//
// Solidity: function updateOracleAddress(address oracle) returns()
func (_FunctionConsumer *FunctionConsumerSession) UpdateOracleAddress(oracle common.Address) (*types.Transaction, error) {
	return _FunctionConsumer.Contract.UpdateOracleAddress(&_FunctionConsumer.TransactOpts, oracle)
}

// UpdateOracleAddress is a paid mutator transaction binding the contract method 0xf7023bb6.
//
// Solidity: function updateOracleAddress(address oracle) returns()
func (_FunctionConsumer *FunctionConsumerTransactorSession) UpdateOracleAddress(oracle common.Address) (*types.Transaction, error) {
	return _FunctionConsumer.Contract.UpdateOracleAddress(&_FunctionConsumer.TransactOpts, oracle)
}

// FunctionConsumerOCRResponseIterator is returned from FilterOCRResponse and is used to iterate over the raw logs and unpacked data for OCRResponse events raised by the FunctionConsumer contract.
type FunctionConsumerOCRResponseIterator struct {
	Event *FunctionConsumerOCRResponse // Event containing the contract specifics and raw log

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
func (it *FunctionConsumerOCRResponseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FunctionConsumerOCRResponse)
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
		it.Event = new(FunctionConsumerOCRResponse)
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
func (it *FunctionConsumerOCRResponseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FunctionConsumerOCRResponseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FunctionConsumerOCRResponse represents a OCRResponse event raised by the FunctionConsumer contract.
type FunctionConsumerOCRResponse struct {
	RequestId [32]byte
	Result    []byte
	Err       []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOCRResponse is a free log retrieval operation binding the contract event 0x7bab0ec163b5c132c72b8146ac4d6e067e82ed58f8b131150aa71c9258911562.
//
// Solidity: event OCRResponse(bytes32 indexed requestId, bytes result, bytes err)
func (_FunctionConsumer *FunctionConsumerFilterer) FilterOCRResponse(opts *bind.FilterOpts, requestId [][32]byte) (*FunctionConsumerOCRResponseIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _FunctionConsumer.contract.FilterLogs(opts, "OCRResponse", requestIdRule)
	if err != nil {
		return nil, err
	}
	return &FunctionConsumerOCRResponseIterator{contract: _FunctionConsumer.contract, event: "OCRResponse", logs: logs, sub: sub}, nil
}

// WatchOCRResponse is a free log subscription operation binding the contract event 0x7bab0ec163b5c132c72b8146ac4d6e067e82ed58f8b131150aa71c9258911562.
//
// Solidity: event OCRResponse(bytes32 indexed requestId, bytes result, bytes err)
func (_FunctionConsumer *FunctionConsumerFilterer) WatchOCRResponse(opts *bind.WatchOpts, sink chan<- *FunctionConsumerOCRResponse, requestId [][32]byte) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _FunctionConsumer.contract.WatchLogs(opts, "OCRResponse", requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FunctionConsumerOCRResponse)
				if err := _FunctionConsumer.contract.UnpackLog(event, "OCRResponse", log); err != nil {
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

// ParseOCRResponse is a log parse operation binding the contract event 0x7bab0ec163b5c132c72b8146ac4d6e067e82ed58f8b131150aa71c9258911562.
//
// Solidity: event OCRResponse(bytes32 indexed requestId, bytes result, bytes err)
func (_FunctionConsumer *FunctionConsumerFilterer) ParseOCRResponse(log types.Log) (*FunctionConsumerOCRResponse, error) {
	event := new(FunctionConsumerOCRResponse)
	if err := _FunctionConsumer.contract.UnpackLog(event, "OCRResponse", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FunctionConsumerOwnershipTransferRequestedIterator is returned from FilterOwnershipTransferRequested and is used to iterate over the raw logs and unpacked data for OwnershipTransferRequested events raised by the FunctionConsumer contract.
type FunctionConsumerOwnershipTransferRequestedIterator struct {
	Event *FunctionConsumerOwnershipTransferRequested // Event containing the contract specifics and raw log

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
func (it *FunctionConsumerOwnershipTransferRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FunctionConsumerOwnershipTransferRequested)
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
		it.Event = new(FunctionConsumerOwnershipTransferRequested)
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
func (it *FunctionConsumerOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FunctionConsumerOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FunctionConsumerOwnershipTransferRequested represents a OwnershipTransferRequested event raised by the FunctionConsumer contract.
type FunctionConsumerOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferRequested is a free log retrieval operation binding the contract event 0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278.
//
// Solidity: event OwnershipTransferRequested(address indexed from, address indexed to)
func (_FunctionConsumer *FunctionConsumerFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*FunctionConsumerOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FunctionConsumer.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &FunctionConsumerOwnershipTransferRequestedIterator{contract: _FunctionConsumer.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferRequested is a free log subscription operation binding the contract event 0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278.
//
// Solidity: event OwnershipTransferRequested(address indexed from, address indexed to)
func (_FunctionConsumer *FunctionConsumerFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *FunctionConsumerOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FunctionConsumer.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FunctionConsumerOwnershipTransferRequested)
				if err := _FunctionConsumer.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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

// ParseOwnershipTransferRequested is a log parse operation binding the contract event 0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278.
//
// Solidity: event OwnershipTransferRequested(address indexed from, address indexed to)
func (_FunctionConsumer *FunctionConsumerFilterer) ParseOwnershipTransferRequested(log types.Log) (*FunctionConsumerOwnershipTransferRequested, error) {
	event := new(FunctionConsumerOwnershipTransferRequested)
	if err := _FunctionConsumer.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FunctionConsumerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the FunctionConsumer contract.
type FunctionConsumerOwnershipTransferredIterator struct {
	Event *FunctionConsumerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FunctionConsumerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FunctionConsumerOwnershipTransferred)
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
		it.Event = new(FunctionConsumerOwnershipTransferred)
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
func (it *FunctionConsumerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FunctionConsumerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FunctionConsumerOwnershipTransferred represents a OwnershipTransferred event raised by the FunctionConsumer contract.
type FunctionConsumerOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed from, address indexed to)
func (_FunctionConsumer *FunctionConsumerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*FunctionConsumerOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FunctionConsumer.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &FunctionConsumerOwnershipTransferredIterator{contract: _FunctionConsumer.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed from, address indexed to)
func (_FunctionConsumer *FunctionConsumerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FunctionConsumerOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _FunctionConsumer.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FunctionConsumerOwnershipTransferred)
				if err := _FunctionConsumer.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
// Solidity: event OwnershipTransferred(address indexed from, address indexed to)
func (_FunctionConsumer *FunctionConsumerFilterer) ParseOwnershipTransferred(log types.Log) (*FunctionConsumerOwnershipTransferred, error) {
	event := new(FunctionConsumerOwnershipTransferred)
	if err := _FunctionConsumer.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FunctionConsumerRequestFulfilledIterator is returned from FilterRequestFulfilled and is used to iterate over the raw logs and unpacked data for RequestFulfilled events raised by the FunctionConsumer contract.
type FunctionConsumerRequestFulfilledIterator struct {
	Event *FunctionConsumerRequestFulfilled // Event containing the contract specifics and raw log

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
func (it *FunctionConsumerRequestFulfilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FunctionConsumerRequestFulfilled)
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
		it.Event = new(FunctionConsumerRequestFulfilled)
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
func (it *FunctionConsumerRequestFulfilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FunctionConsumerRequestFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FunctionConsumerRequestFulfilled represents a RequestFulfilled event raised by the FunctionConsumer contract.
type FunctionConsumerRequestFulfilled struct {
	Id  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRequestFulfilled is a free log retrieval operation binding the contract event 0x85e1543bf2f84fe80c6badbce3648c8539ad1df4d2b3d822938ca0538be727e6.
//
// Solidity: event RequestFulfilled(bytes32 indexed id)
func (_FunctionConsumer *FunctionConsumerFilterer) FilterRequestFulfilled(opts *bind.FilterOpts, id [][32]byte) (*FunctionConsumerRequestFulfilledIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _FunctionConsumer.contract.FilterLogs(opts, "RequestFulfilled", idRule)
	if err != nil {
		return nil, err
	}
	return &FunctionConsumerRequestFulfilledIterator{contract: _FunctionConsumer.contract, event: "RequestFulfilled", logs: logs, sub: sub}, nil
}

// WatchRequestFulfilled is a free log subscription operation binding the contract event 0x85e1543bf2f84fe80c6badbce3648c8539ad1df4d2b3d822938ca0538be727e6.
//
// Solidity: event RequestFulfilled(bytes32 indexed id)
func (_FunctionConsumer *FunctionConsumerFilterer) WatchRequestFulfilled(opts *bind.WatchOpts, sink chan<- *FunctionConsumerRequestFulfilled, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _FunctionConsumer.contract.WatchLogs(opts, "RequestFulfilled", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FunctionConsumerRequestFulfilled)
				if err := _FunctionConsumer.contract.UnpackLog(event, "RequestFulfilled", log); err != nil {
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

// ParseRequestFulfilled is a log parse operation binding the contract event 0x85e1543bf2f84fe80c6badbce3648c8539ad1df4d2b3d822938ca0538be727e6.
//
// Solidity: event RequestFulfilled(bytes32 indexed id)
func (_FunctionConsumer *FunctionConsumerFilterer) ParseRequestFulfilled(log types.Log) (*FunctionConsumerRequestFulfilled, error) {
	event := new(FunctionConsumerRequestFulfilled)
	if err := _FunctionConsumer.contract.UnpackLog(event, "RequestFulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FunctionConsumerRequestSentIterator is returned from FilterRequestSent and is used to iterate over the raw logs and unpacked data for RequestSent events raised by the FunctionConsumer contract.
type FunctionConsumerRequestSentIterator struct {
	Event *FunctionConsumerRequestSent // Event containing the contract specifics and raw log

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
func (it *FunctionConsumerRequestSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FunctionConsumerRequestSent)
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
		it.Event = new(FunctionConsumerRequestSent)
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
func (it *FunctionConsumerRequestSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FunctionConsumerRequestSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FunctionConsumerRequestSent represents a RequestSent event raised by the FunctionConsumer contract.
type FunctionConsumerRequestSent struct {
	Id  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRequestSent is a free log retrieval operation binding the contract event 0x1131472297a800fee664d1d89cfa8f7676ff07189ecc53f80bbb5f4969099db8.
//
// Solidity: event RequestSent(bytes32 indexed id)
func (_FunctionConsumer *FunctionConsumerFilterer) FilterRequestSent(opts *bind.FilterOpts, id [][32]byte) (*FunctionConsumerRequestSentIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _FunctionConsumer.contract.FilterLogs(opts, "RequestSent", idRule)
	if err != nil {
		return nil, err
	}
	return &FunctionConsumerRequestSentIterator{contract: _FunctionConsumer.contract, event: "RequestSent", logs: logs, sub: sub}, nil
}

// WatchRequestSent is a free log subscription operation binding the contract event 0x1131472297a800fee664d1d89cfa8f7676ff07189ecc53f80bbb5f4969099db8.
//
// Solidity: event RequestSent(bytes32 indexed id)
func (_FunctionConsumer *FunctionConsumerFilterer) WatchRequestSent(opts *bind.WatchOpts, sink chan<- *FunctionConsumerRequestSent, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _FunctionConsumer.contract.WatchLogs(opts, "RequestSent", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FunctionConsumerRequestSent)
				if err := _FunctionConsumer.contract.UnpackLog(event, "RequestSent", log); err != nil {
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

// ParseRequestSent is a log parse operation binding the contract event 0x1131472297a800fee664d1d89cfa8f7676ff07189ecc53f80bbb5f4969099db8.
//
// Solidity: event RequestSent(bytes32 indexed id)
func (_FunctionConsumer *FunctionConsumerFilterer) ParseRequestSent(log types.Log) (*FunctionConsumerRequestSent, error) {
	event := new(FunctionConsumerRequestSent)
	if err := _FunctionConsumer.contract.UnpackLog(event, "RequestSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
