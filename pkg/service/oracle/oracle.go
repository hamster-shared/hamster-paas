// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package oracle

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

// FunctionsBillingRegistryInterfaceRequestBilling is an auto generated low-level Go binding around an user-defined struct.
type FunctionsBillingRegistryInterfaceRequestBilling struct {
	SubscriptionId uint64
	Client         common.Address
	GasLimit       uint32
	GasPrice       *big.Int
}

// OracleMetaData contains all meta data concerning the Oracle contract.
var OracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AlreadySet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotSelfTransfer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyBillingRegistry\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyPublicKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyRequestData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptySendersList\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InconsistentReportData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRequestID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAllowedToSetSenders\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotProposedOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerMustBeSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReportInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnauthorizedPublicKeyChange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnauthorizedSender\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AuthorizedSendersActive\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"senders\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"changedBy\",\"type\":\"address\"}],\"name\":\"AuthorizedSendersChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AuthorizedSendersDeactive\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"requestingContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"requestInitiator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"subscriptionId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"subscriptionOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"OracleRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"}],\"name\":\"OracleResponse\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"name\":\"Transmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"UserCallbackError\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"lowLevelData\",\"type\":\"bytes\"}],\"name\":\"UserCallbackRawError\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activateAuthorizedReceiver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"senders\",\"type\":\"address[]\"}],\"name\":\"addAuthorizedSenders\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"authorizedReceiverActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deactivateAuthorizedReceiver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"}],\"name\":\"deleteNodePublicKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"subscriptionId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"gasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"}],\"name\":\"estimateCost\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllNodePublicKeys\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAuthorizedSenders\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDONPublicKey\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"subscriptionId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"client\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"gasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"}],\"internalType\":\"structFunctionsBillingRegistryInterface.RequestBilling\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"getRequiredFee\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"isAuthorizedSender\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDigestAndEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"scanLogs\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"senders\",\"type\":\"address[]\"}],\"name\":\"removeAuthorizedSenders\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"subscriptionId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"gasLimit\",\"type\":\"uint32\"}],\"name\":\"sendRequest\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"_f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"_onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"_offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_offchainConfig\",\"type\":\"bytes\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"donPublicKey\",\"type\":\"bytes\"}],\"name\":\"setDONPublicKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"}],\"name\":\"setNodePublicKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"registryAddress\",\"type\":\"address\"}],\"name\":\"setRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"rawVs\",\"type\":\"bytes32\"}],\"name\":\"transmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transmitters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// OracleABI is the input ABI used to generate the binding from.
// Deprecated: Use OracleMetaData.ABI instead.
var OracleABI = OracleMetaData.ABI

// Oracle is an auto generated Go binding around an Ethereum contract.
type Oracle struct {
	OracleCaller     // Read-only binding to the contract
	OracleTransactor // Write-only binding to the contract
	OracleFilterer   // Log filterer for contract events
}

// OracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type OracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OracleSession struct {
	Contract     *Oracle           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OracleCallerSession struct {
	Contract *OracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OracleTransactorSession struct {
	Contract     *OracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type OracleRaw struct {
	Contract *Oracle // Generic contract binding to access the raw methods on
}

// OracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OracleCallerRaw struct {
	Contract *OracleCaller // Generic read-only contract binding to access the raw methods on
}

// OracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OracleTransactorRaw struct {
	Contract *OracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOracle creates a new instance of Oracle, bound to a specific deployed contract.
func NewOracle(address common.Address, backend bind.ContractBackend) (*Oracle, error) {
	contract, err := bindOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Oracle{OracleCaller: OracleCaller{contract: contract}, OracleTransactor: OracleTransactor{contract: contract}, OracleFilterer: OracleFilterer{contract: contract}}, nil
}

// NewOracleCaller creates a new read-only instance of Oracle, bound to a specific deployed contract.
func NewOracleCaller(address common.Address, caller bind.ContractCaller) (*OracleCaller, error) {
	contract, err := bindOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OracleCaller{contract: contract}, nil
}

// NewOracleTransactor creates a new write-only instance of Oracle, bound to a specific deployed contract.
func NewOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*OracleTransactor, error) {
	contract, err := bindOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OracleTransactor{contract: contract}, nil
}

// NewOracleFilterer creates a new log filterer instance of Oracle, bound to a specific deployed contract.
func NewOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*OracleFilterer, error) {
	contract, err := bindOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OracleFilterer{contract: contract}, nil
}

// bindOracle binds a generic wrapper to an already deployed contract.
func bindOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OracleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oracle *OracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Oracle.Contract.OracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oracle *OracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oracle.Contract.OracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oracle *OracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Oracle.Contract.OracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oracle *OracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Oracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oracle *OracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oracle *OracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Oracle.Contract.contract.Transact(opts, method, params...)
}

// AuthorizedReceiverActive is a free data retrieval call binding the contract method 0x4b4fa0c1.
//
// Solidity: function authorizedReceiverActive() view returns(bool)
func (_Oracle *OracleCaller) AuthorizedReceiverActive(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "authorizedReceiverActive")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AuthorizedReceiverActive is a free data retrieval call binding the contract method 0x4b4fa0c1.
//
// Solidity: function authorizedReceiverActive() view returns(bool)
func (_Oracle *OracleSession) AuthorizedReceiverActive() (bool, error) {
	return _Oracle.Contract.AuthorizedReceiverActive(&_Oracle.CallOpts)
}

// AuthorizedReceiverActive is a free data retrieval call binding the contract method 0x4b4fa0c1.
//
// Solidity: function authorizedReceiverActive() view returns(bool)
func (_Oracle *OracleCallerSession) AuthorizedReceiverActive() (bool, error) {
	return _Oracle.Contract.AuthorizedReceiverActive(&_Oracle.CallOpts)
}

// EstimateCost is a free data retrieval call binding the contract method 0xd227d245.
//
// Solidity: function estimateCost(uint64 subscriptionId, bytes data, uint32 gasLimit, uint256 gasPrice) view returns(uint96)
func (_Oracle *OracleCaller) EstimateCost(opts *bind.CallOpts, subscriptionId uint64, data []byte, gasLimit uint32, gasPrice *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "estimateCost", subscriptionId, data, gasLimit, gasPrice)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimateCost is a free data retrieval call binding the contract method 0xd227d245.
//
// Solidity: function estimateCost(uint64 subscriptionId, bytes data, uint32 gasLimit, uint256 gasPrice) view returns(uint96)
func (_Oracle *OracleSession) EstimateCost(subscriptionId uint64, data []byte, gasLimit uint32, gasPrice *big.Int) (*big.Int, error) {
	return _Oracle.Contract.EstimateCost(&_Oracle.CallOpts, subscriptionId, data, gasLimit, gasPrice)
}

// EstimateCost is a free data retrieval call binding the contract method 0xd227d245.
//
// Solidity: function estimateCost(uint64 subscriptionId, bytes data, uint32 gasLimit, uint256 gasPrice) view returns(uint96)
func (_Oracle *OracleCallerSession) EstimateCost(subscriptionId uint64, data []byte, gasLimit uint32, gasPrice *big.Int) (*big.Int, error) {
	return _Oracle.Contract.EstimateCost(&_Oracle.CallOpts, subscriptionId, data, gasLimit, gasPrice)
}

// GetAllNodePublicKeys is a free data retrieval call binding the contract method 0x53398987.
//
// Solidity: function getAllNodePublicKeys() view returns(address[], bytes[])
func (_Oracle *OracleCaller) GetAllNodePublicKeys(opts *bind.CallOpts) ([]common.Address, [][]byte, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getAllNodePublicKeys")

	if err != nil {
		return *new([]common.Address), *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new([][]byte)).(*[][]byte)

	return out0, out1, err

}

// GetAllNodePublicKeys is a free data retrieval call binding the contract method 0x53398987.
//
// Solidity: function getAllNodePublicKeys() view returns(address[], bytes[])
func (_Oracle *OracleSession) GetAllNodePublicKeys() ([]common.Address, [][]byte, error) {
	return _Oracle.Contract.GetAllNodePublicKeys(&_Oracle.CallOpts)
}

// GetAllNodePublicKeys is a free data retrieval call binding the contract method 0x53398987.
//
// Solidity: function getAllNodePublicKeys() view returns(address[], bytes[])
func (_Oracle *OracleCallerSession) GetAllNodePublicKeys() ([]common.Address, [][]byte, error) {
	return _Oracle.Contract.GetAllNodePublicKeys(&_Oracle.CallOpts)
}

// GetAuthorizedSenders is a free data retrieval call binding the contract method 0x2408afaa.
//
// Solidity: function getAuthorizedSenders() view returns(address[])
func (_Oracle *OracleCaller) GetAuthorizedSenders(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getAuthorizedSenders")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAuthorizedSenders is a free data retrieval call binding the contract method 0x2408afaa.
//
// Solidity: function getAuthorizedSenders() view returns(address[])
func (_Oracle *OracleSession) GetAuthorizedSenders() ([]common.Address, error) {
	return _Oracle.Contract.GetAuthorizedSenders(&_Oracle.CallOpts)
}

// GetAuthorizedSenders is a free data retrieval call binding the contract method 0x2408afaa.
//
// Solidity: function getAuthorizedSenders() view returns(address[])
func (_Oracle *OracleCallerSession) GetAuthorizedSenders() ([]common.Address, error) {
	return _Oracle.Contract.GetAuthorizedSenders(&_Oracle.CallOpts)
}

// GetDONPublicKey is a free data retrieval call binding the contract method 0xd328a91e.
//
// Solidity: function getDONPublicKey() view returns(bytes)
func (_Oracle *OracleCaller) GetDONPublicKey(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getDONPublicKey")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetDONPublicKey is a free data retrieval call binding the contract method 0xd328a91e.
//
// Solidity: function getDONPublicKey() view returns(bytes)
func (_Oracle *OracleSession) GetDONPublicKey() ([]byte, error) {
	return _Oracle.Contract.GetDONPublicKey(&_Oracle.CallOpts)
}

// GetDONPublicKey is a free data retrieval call binding the contract method 0xd328a91e.
//
// Solidity: function getDONPublicKey() view returns(bytes)
func (_Oracle *OracleCallerSession) GetDONPublicKey() ([]byte, error) {
	return _Oracle.Contract.GetDONPublicKey(&_Oracle.CallOpts)
}

// GetRegistry is a free data retrieval call binding the contract method 0x5ab1bd53.
//
// Solidity: function getRegistry() view returns(address)
func (_Oracle *OracleCaller) GetRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRegistry is a free data retrieval call binding the contract method 0x5ab1bd53.
//
// Solidity: function getRegistry() view returns(address)
func (_Oracle *OracleSession) GetRegistry() (common.Address, error) {
	return _Oracle.Contract.GetRegistry(&_Oracle.CallOpts)
}

// GetRegistry is a free data retrieval call binding the contract method 0x5ab1bd53.
//
// Solidity: function getRegistry() view returns(address)
func (_Oracle *OracleCallerSession) GetRegistry() (common.Address, error) {
	return _Oracle.Contract.GetRegistry(&_Oracle.CallOpts)
}

// GetRequiredFee is a free data retrieval call binding the contract method 0xf1e14a21.
//
// Solidity: function getRequiredFee(bytes , (uint64,address,uint32,uint256) ) pure returns(uint96)
func (_Oracle *OracleCaller) GetRequiredFee(opts *bind.CallOpts, arg0 []byte, arg1 FunctionsBillingRegistryInterfaceRequestBilling) (*big.Int, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "getRequiredFee", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRequiredFee is a free data retrieval call binding the contract method 0xf1e14a21.
//
// Solidity: function getRequiredFee(bytes , (uint64,address,uint32,uint256) ) pure returns(uint96)
func (_Oracle *OracleSession) GetRequiredFee(arg0 []byte, arg1 FunctionsBillingRegistryInterfaceRequestBilling) (*big.Int, error) {
	return _Oracle.Contract.GetRequiredFee(&_Oracle.CallOpts, arg0, arg1)
}

// GetRequiredFee is a free data retrieval call binding the contract method 0xf1e14a21.
//
// Solidity: function getRequiredFee(bytes , (uint64,address,uint32,uint256) ) pure returns(uint96)
func (_Oracle *OracleCallerSession) GetRequiredFee(arg0 []byte, arg1 FunctionsBillingRegistryInterfaceRequestBilling) (*big.Int, error) {
	return _Oracle.Contract.GetRequiredFee(&_Oracle.CallOpts, arg0, arg1)
}

// IsAuthorizedSender is a free data retrieval call binding the contract method 0xfa00763a.
//
// Solidity: function isAuthorizedSender(address sender) view returns(bool)
func (_Oracle *OracleCaller) IsAuthorizedSender(opts *bind.CallOpts, sender common.Address) (bool, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "isAuthorizedSender", sender)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAuthorizedSender is a free data retrieval call binding the contract method 0xfa00763a.
//
// Solidity: function isAuthorizedSender(address sender) view returns(bool)
func (_Oracle *OracleSession) IsAuthorizedSender(sender common.Address) (bool, error) {
	return _Oracle.Contract.IsAuthorizedSender(&_Oracle.CallOpts, sender)
}

// IsAuthorizedSender is a free data retrieval call binding the contract method 0xfa00763a.
//
// Solidity: function isAuthorizedSender(address sender) view returns(bool)
func (_Oracle *OracleCallerSession) IsAuthorizedSender(sender common.Address) (bool, error) {
	return _Oracle.Contract.IsAuthorizedSender(&_Oracle.CallOpts, sender)
}

// LatestConfigDetails is a free data retrieval call binding the contract method 0x81ff7048.
//
// Solidity: function latestConfigDetails() view returns(uint32 configCount, uint32 blockNumber, bytes32 configDigest)
func (_Oracle *OracleCaller) LatestConfigDetails(opts *bind.CallOpts) (struct {
	ConfigCount  uint32
	BlockNumber  uint32
	ConfigDigest [32]byte
}, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "latestConfigDetails")

	outstruct := new(struct {
		ConfigCount  uint32
		BlockNumber  uint32
		ConfigDigest [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ConfigCount = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.BlockNumber = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.ConfigDigest = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// LatestConfigDetails is a free data retrieval call binding the contract method 0x81ff7048.
//
// Solidity: function latestConfigDetails() view returns(uint32 configCount, uint32 blockNumber, bytes32 configDigest)
func (_Oracle *OracleSession) LatestConfigDetails() (struct {
	ConfigCount  uint32
	BlockNumber  uint32
	ConfigDigest [32]byte
}, error) {
	return _Oracle.Contract.LatestConfigDetails(&_Oracle.CallOpts)
}

// LatestConfigDetails is a free data retrieval call binding the contract method 0x81ff7048.
//
// Solidity: function latestConfigDetails() view returns(uint32 configCount, uint32 blockNumber, bytes32 configDigest)
func (_Oracle *OracleCallerSession) LatestConfigDetails() (struct {
	ConfigCount  uint32
	BlockNumber  uint32
	ConfigDigest [32]byte
}, error) {
	return _Oracle.Contract.LatestConfigDetails(&_Oracle.CallOpts)
}

// LatestConfigDigestAndEpoch is a free data retrieval call binding the contract method 0xafcb95d7.
//
// Solidity: function latestConfigDigestAndEpoch() view returns(bool scanLogs, bytes32 configDigest, uint32 epoch)
func (_Oracle *OracleCaller) LatestConfigDigestAndEpoch(opts *bind.CallOpts) (struct {
	ScanLogs     bool
	ConfigDigest [32]byte
	Epoch        uint32
}, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "latestConfigDigestAndEpoch")

	outstruct := new(struct {
		ScanLogs     bool
		ConfigDigest [32]byte
		Epoch        uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ScanLogs = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.ConfigDigest = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.Epoch = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

// LatestConfigDigestAndEpoch is a free data retrieval call binding the contract method 0xafcb95d7.
//
// Solidity: function latestConfigDigestAndEpoch() view returns(bool scanLogs, bytes32 configDigest, uint32 epoch)
func (_Oracle *OracleSession) LatestConfigDigestAndEpoch() (struct {
	ScanLogs     bool
	ConfigDigest [32]byte
	Epoch        uint32
}, error) {
	return _Oracle.Contract.LatestConfigDigestAndEpoch(&_Oracle.CallOpts)
}

// LatestConfigDigestAndEpoch is a free data retrieval call binding the contract method 0xafcb95d7.
//
// Solidity: function latestConfigDigestAndEpoch() view returns(bool scanLogs, bytes32 configDigest, uint32 epoch)
func (_Oracle *OracleCallerSession) LatestConfigDigestAndEpoch() (struct {
	ScanLogs     bool
	ConfigDigest [32]byte
	Epoch        uint32
}, error) {
	return _Oracle.Contract.LatestConfigDigestAndEpoch(&_Oracle.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Oracle *OracleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Oracle *OracleSession) Owner() (common.Address, error) {
	return _Oracle.Contract.Owner(&_Oracle.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Oracle *OracleCallerSession) Owner() (common.Address, error) {
	return _Oracle.Contract.Owner(&_Oracle.CallOpts)
}

// Transmitters is a free data retrieval call binding the contract method 0x81411834.
//
// Solidity: function transmitters() view returns(address[])
func (_Oracle *OracleCaller) Transmitters(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "transmitters")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// Transmitters is a free data retrieval call binding the contract method 0x81411834.
//
// Solidity: function transmitters() view returns(address[])
func (_Oracle *OracleSession) Transmitters() ([]common.Address, error) {
	return _Oracle.Contract.Transmitters(&_Oracle.CallOpts)
}

// Transmitters is a free data retrieval call binding the contract method 0x81411834.
//
// Solidity: function transmitters() view returns(address[])
func (_Oracle *OracleCallerSession) Transmitters() ([]common.Address, error) {
	return _Oracle.Contract.Transmitters(&_Oracle.CallOpts)
}

// TypeAndVersion is a free data retrieval call binding the contract method 0x181f5a77.
//
// Solidity: function typeAndVersion() pure returns(string)
func (_Oracle *OracleCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Oracle.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TypeAndVersion is a free data retrieval call binding the contract method 0x181f5a77.
//
// Solidity: function typeAndVersion() pure returns(string)
func (_Oracle *OracleSession) TypeAndVersion() (string, error) {
	return _Oracle.Contract.TypeAndVersion(&_Oracle.CallOpts)
}

// TypeAndVersion is a free data retrieval call binding the contract method 0x181f5a77.
//
// Solidity: function typeAndVersion() pure returns(string)
func (_Oracle *OracleCallerSession) TypeAndVersion() (string, error) {
	return _Oracle.Contract.TypeAndVersion(&_Oracle.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Oracle *OracleTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Oracle *OracleSession) AcceptOwnership() (*types.Transaction, error) {
	return _Oracle.Contract.AcceptOwnership(&_Oracle.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Oracle *OracleTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Oracle.Contract.AcceptOwnership(&_Oracle.TransactOpts)
}

// ActivateAuthorizedReceiver is a paid mutator transaction binding the contract method 0x110254c8.
//
// Solidity: function activateAuthorizedReceiver() returns()
func (_Oracle *OracleTransactor) ActivateAuthorizedReceiver(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "activateAuthorizedReceiver")
}

// ActivateAuthorizedReceiver is a paid mutator transaction binding the contract method 0x110254c8.
//
// Solidity: function activateAuthorizedReceiver() returns()
func (_Oracle *OracleSession) ActivateAuthorizedReceiver() (*types.Transaction, error) {
	return _Oracle.Contract.ActivateAuthorizedReceiver(&_Oracle.TransactOpts)
}

// ActivateAuthorizedReceiver is a paid mutator transaction binding the contract method 0x110254c8.
//
// Solidity: function activateAuthorizedReceiver() returns()
func (_Oracle *OracleTransactorSession) ActivateAuthorizedReceiver() (*types.Transaction, error) {
	return _Oracle.Contract.ActivateAuthorizedReceiver(&_Oracle.TransactOpts)
}

// AddAuthorizedSenders is a paid mutator transaction binding the contract method 0x4dcef404.
//
// Solidity: function addAuthorizedSenders(address[] senders) returns()
func (_Oracle *OracleTransactor) AddAuthorizedSenders(opts *bind.TransactOpts, senders []common.Address) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "addAuthorizedSenders", senders)
}

// AddAuthorizedSenders is a paid mutator transaction binding the contract method 0x4dcef404.
//
// Solidity: function addAuthorizedSenders(address[] senders) returns()
func (_Oracle *OracleSession) AddAuthorizedSenders(senders []common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.AddAuthorizedSenders(&_Oracle.TransactOpts, senders)
}

// AddAuthorizedSenders is a paid mutator transaction binding the contract method 0x4dcef404.
//
// Solidity: function addAuthorizedSenders(address[] senders) returns()
func (_Oracle *OracleTransactorSession) AddAuthorizedSenders(senders []common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.AddAuthorizedSenders(&_Oracle.TransactOpts, senders)
}

// DeactivateAuthorizedReceiver is a paid mutator transaction binding the contract method 0x91bb64eb.
//
// Solidity: function deactivateAuthorizedReceiver() returns()
func (_Oracle *OracleTransactor) DeactivateAuthorizedReceiver(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "deactivateAuthorizedReceiver")
}

// DeactivateAuthorizedReceiver is a paid mutator transaction binding the contract method 0x91bb64eb.
//
// Solidity: function deactivateAuthorizedReceiver() returns()
func (_Oracle *OracleSession) DeactivateAuthorizedReceiver() (*types.Transaction, error) {
	return _Oracle.Contract.DeactivateAuthorizedReceiver(&_Oracle.TransactOpts)
}

// DeactivateAuthorizedReceiver is a paid mutator transaction binding the contract method 0x91bb64eb.
//
// Solidity: function deactivateAuthorizedReceiver() returns()
func (_Oracle *OracleTransactorSession) DeactivateAuthorizedReceiver() (*types.Transaction, error) {
	return _Oracle.Contract.DeactivateAuthorizedReceiver(&_Oracle.TransactOpts)
}

// DeleteNodePublicKey is a paid mutator transaction binding the contract method 0x26ceabac.
//
// Solidity: function deleteNodePublicKey(address node) returns()
func (_Oracle *OracleTransactor) DeleteNodePublicKey(opts *bind.TransactOpts, node common.Address) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "deleteNodePublicKey", node)
}

// DeleteNodePublicKey is a paid mutator transaction binding the contract method 0x26ceabac.
//
// Solidity: function deleteNodePublicKey(address node) returns()
func (_Oracle *OracleSession) DeleteNodePublicKey(node common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.DeleteNodePublicKey(&_Oracle.TransactOpts, node)
}

// DeleteNodePublicKey is a paid mutator transaction binding the contract method 0x26ceabac.
//
// Solidity: function deleteNodePublicKey(address node) returns()
func (_Oracle *OracleTransactorSession) DeleteNodePublicKey(node common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.DeleteNodePublicKey(&_Oracle.TransactOpts, node)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Oracle *OracleTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Oracle *OracleSession) Initialize() (*types.Transaction, error) {
	return _Oracle.Contract.Initialize(&_Oracle.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Oracle *OracleTransactorSession) Initialize() (*types.Transaction, error) {
	return _Oracle.Contract.Initialize(&_Oracle.TransactOpts)
}

// RemoveAuthorizedSenders is a paid mutator transaction binding the contract method 0x03e1bf23.
//
// Solidity: function removeAuthorizedSenders(address[] senders) returns()
func (_Oracle *OracleTransactor) RemoveAuthorizedSenders(opts *bind.TransactOpts, senders []common.Address) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "removeAuthorizedSenders", senders)
}

// RemoveAuthorizedSenders is a paid mutator transaction binding the contract method 0x03e1bf23.
//
// Solidity: function removeAuthorizedSenders(address[] senders) returns()
func (_Oracle *OracleSession) RemoveAuthorizedSenders(senders []common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.RemoveAuthorizedSenders(&_Oracle.TransactOpts, senders)
}

// RemoveAuthorizedSenders is a paid mutator transaction binding the contract method 0x03e1bf23.
//
// Solidity: function removeAuthorizedSenders(address[] senders) returns()
func (_Oracle *OracleTransactorSession) RemoveAuthorizedSenders(senders []common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.RemoveAuthorizedSenders(&_Oracle.TransactOpts, senders)
}

// SendRequest is a paid mutator transaction binding the contract method 0x28242b04.
//
// Solidity: function sendRequest(uint64 subscriptionId, bytes data, uint32 gasLimit) returns(bytes32)
func (_Oracle *OracleTransactor) SendRequest(opts *bind.TransactOpts, subscriptionId uint64, data []byte, gasLimit uint32) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "sendRequest", subscriptionId, data, gasLimit)
}

// SendRequest is a paid mutator transaction binding the contract method 0x28242b04.
//
// Solidity: function sendRequest(uint64 subscriptionId, bytes data, uint32 gasLimit) returns(bytes32)
func (_Oracle *OracleSession) SendRequest(subscriptionId uint64, data []byte, gasLimit uint32) (*types.Transaction, error) {
	return _Oracle.Contract.SendRequest(&_Oracle.TransactOpts, subscriptionId, data, gasLimit)
}

// SendRequest is a paid mutator transaction binding the contract method 0x28242b04.
//
// Solidity: function sendRequest(uint64 subscriptionId, bytes data, uint32 gasLimit) returns(bytes32)
func (_Oracle *OracleTransactorSession) SendRequest(subscriptionId uint64, data []byte, gasLimit uint32) (*types.Transaction, error) {
	return _Oracle.Contract.SendRequest(&_Oracle.TransactOpts, subscriptionId, data, gasLimit)
}

// SetConfig is a paid mutator transaction binding the contract method 0xe3d0e712.
//
// Solidity: function setConfig(address[] _signers, address[] _transmitters, uint8 _f, bytes _onchainConfig, uint64 _offchainConfigVersion, bytes _offchainConfig) returns()
func (_Oracle *OracleTransactor) SetConfig(opts *bind.TransactOpts, _signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "setConfig", _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
}

// SetConfig is a paid mutator transaction binding the contract method 0xe3d0e712.
//
// Solidity: function setConfig(address[] _signers, address[] _transmitters, uint8 _f, bytes _onchainConfig, uint64 _offchainConfigVersion, bytes _offchainConfig) returns()
func (_Oracle *OracleSession) SetConfig(_signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _Oracle.Contract.SetConfig(&_Oracle.TransactOpts, _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
}

// SetConfig is a paid mutator transaction binding the contract method 0xe3d0e712.
//
// Solidity: function setConfig(address[] _signers, address[] _transmitters, uint8 _f, bytes _onchainConfig, uint64 _offchainConfigVersion, bytes _offchainConfig) returns()
func (_Oracle *OracleTransactorSession) SetConfig(_signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _Oracle.Contract.SetConfig(&_Oracle.TransactOpts, _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
}

// SetDONPublicKey is a paid mutator transaction binding the contract method 0x7f15e166.
//
// Solidity: function setDONPublicKey(bytes donPublicKey) returns()
func (_Oracle *OracleTransactor) SetDONPublicKey(opts *bind.TransactOpts, donPublicKey []byte) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "setDONPublicKey", donPublicKey)
}

// SetDONPublicKey is a paid mutator transaction binding the contract method 0x7f15e166.
//
// Solidity: function setDONPublicKey(bytes donPublicKey) returns()
func (_Oracle *OracleSession) SetDONPublicKey(donPublicKey []byte) (*types.Transaction, error) {
	return _Oracle.Contract.SetDONPublicKey(&_Oracle.TransactOpts, donPublicKey)
}

// SetDONPublicKey is a paid mutator transaction binding the contract method 0x7f15e166.
//
// Solidity: function setDONPublicKey(bytes donPublicKey) returns()
func (_Oracle *OracleTransactorSession) SetDONPublicKey(donPublicKey []byte) (*types.Transaction, error) {
	return _Oracle.Contract.SetDONPublicKey(&_Oracle.TransactOpts, donPublicKey)
}

// SetNodePublicKey is a paid mutator transaction binding the contract method 0x80756031.
//
// Solidity: function setNodePublicKey(address node, bytes publicKey) returns()
func (_Oracle *OracleTransactor) SetNodePublicKey(opts *bind.TransactOpts, node common.Address, publicKey []byte) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "setNodePublicKey", node, publicKey)
}

// SetNodePublicKey is a paid mutator transaction binding the contract method 0x80756031.
//
// Solidity: function setNodePublicKey(address node, bytes publicKey) returns()
func (_Oracle *OracleSession) SetNodePublicKey(node common.Address, publicKey []byte) (*types.Transaction, error) {
	return _Oracle.Contract.SetNodePublicKey(&_Oracle.TransactOpts, node, publicKey)
}

// SetNodePublicKey is a paid mutator transaction binding the contract method 0x80756031.
//
// Solidity: function setNodePublicKey(address node, bytes publicKey) returns()
func (_Oracle *OracleTransactorSession) SetNodePublicKey(node common.Address, publicKey []byte) (*types.Transaction, error) {
	return _Oracle.Contract.SetNodePublicKey(&_Oracle.TransactOpts, node, publicKey)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address registryAddress) returns()
func (_Oracle *OracleTransactor) SetRegistry(opts *bind.TransactOpts, registryAddress common.Address) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "setRegistry", registryAddress)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address registryAddress) returns()
func (_Oracle *OracleSession) SetRegistry(registryAddress common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.SetRegistry(&_Oracle.TransactOpts, registryAddress)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address registryAddress) returns()
func (_Oracle *OracleTransactorSession) SetRegistry(registryAddress common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.SetRegistry(&_Oracle.TransactOpts, registryAddress)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address to) returns()
func (_Oracle *OracleTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "transferOwnership", to)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address to) returns()
func (_Oracle *OracleSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.TransferOwnership(&_Oracle.TransactOpts, to)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address to) returns()
func (_Oracle *OracleTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _Oracle.Contract.TransferOwnership(&_Oracle.TransactOpts, to)
}

// Transmit is a paid mutator transaction binding the contract method 0xb1dc65a4.
//
// Solidity: function transmit(bytes32[3] reportContext, bytes report, bytes32[] rs, bytes32[] ss, bytes32 rawVs) returns()
func (_Oracle *OracleTransactor) Transmit(opts *bind.TransactOpts, reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _Oracle.contract.Transact(opts, "transmit", reportContext, report, rs, ss, rawVs)
}

// Transmit is a paid mutator transaction binding the contract method 0xb1dc65a4.
//
// Solidity: function transmit(bytes32[3] reportContext, bytes report, bytes32[] rs, bytes32[] ss, bytes32 rawVs) returns()
func (_Oracle *OracleSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _Oracle.Contract.Transmit(&_Oracle.TransactOpts, reportContext, report, rs, ss, rawVs)
}

// Transmit is a paid mutator transaction binding the contract method 0xb1dc65a4.
//
// Solidity: function transmit(bytes32[3] reportContext, bytes report, bytes32[] rs, bytes32[] ss, bytes32 rawVs) returns()
func (_Oracle *OracleTransactorSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _Oracle.Contract.Transmit(&_Oracle.TransactOpts, reportContext, report, rs, ss, rawVs)
}

// OracleAuthorizedSendersActiveIterator is returned from FilterAuthorizedSendersActive and is used to iterate over the raw logs and unpacked data for AuthorizedSendersActive events raised by the Oracle contract.
type OracleAuthorizedSendersActiveIterator struct {
	Event *OracleAuthorizedSendersActive // Event containing the contract specifics and raw log

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
func (it *OracleAuthorizedSendersActiveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleAuthorizedSendersActive)
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
		it.Event = new(OracleAuthorizedSendersActive)
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
func (it *OracleAuthorizedSendersActiveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleAuthorizedSendersActiveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleAuthorizedSendersActive represents a AuthorizedSendersActive event raised by the Oracle contract.
type OracleAuthorizedSendersActive struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAuthorizedSendersActive is a free log retrieval operation binding the contract event 0xae51766a982895b0c444fc99fc1a560762b464d709e6c78376c85617f7eeb5ce.
//
// Solidity: event AuthorizedSendersActive(address account)
func (_Oracle *OracleFilterer) FilterAuthorizedSendersActive(opts *bind.FilterOpts) (*OracleAuthorizedSendersActiveIterator, error) {

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "AuthorizedSendersActive")
	if err != nil {
		return nil, err
	}
	return &OracleAuthorizedSendersActiveIterator{contract: _Oracle.contract, event: "AuthorizedSendersActive", logs: logs, sub: sub}, nil
}

// WatchAuthorizedSendersActive is a free log subscription operation binding the contract event 0xae51766a982895b0c444fc99fc1a560762b464d709e6c78376c85617f7eeb5ce.
//
// Solidity: event AuthorizedSendersActive(address account)
func (_Oracle *OracleFilterer) WatchAuthorizedSendersActive(opts *bind.WatchOpts, sink chan<- *OracleAuthorizedSendersActive) (event.Subscription, error) {

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "AuthorizedSendersActive")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleAuthorizedSendersActive)
				if err := _Oracle.contract.UnpackLog(event, "AuthorizedSendersActive", log); err != nil {
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

// ParseAuthorizedSendersActive is a log parse operation binding the contract event 0xae51766a982895b0c444fc99fc1a560762b464d709e6c78376c85617f7eeb5ce.
//
// Solidity: event AuthorizedSendersActive(address account)
func (_Oracle *OracleFilterer) ParseAuthorizedSendersActive(log types.Log) (*OracleAuthorizedSendersActive, error) {
	event := new(OracleAuthorizedSendersActive)
	if err := _Oracle.contract.UnpackLog(event, "AuthorizedSendersActive", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleAuthorizedSendersChangedIterator is returned from FilterAuthorizedSendersChanged and is used to iterate over the raw logs and unpacked data for AuthorizedSendersChanged events raised by the Oracle contract.
type OracleAuthorizedSendersChangedIterator struct {
	Event *OracleAuthorizedSendersChanged // Event containing the contract specifics and raw log

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
func (it *OracleAuthorizedSendersChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleAuthorizedSendersChanged)
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
		it.Event = new(OracleAuthorizedSendersChanged)
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
func (it *OracleAuthorizedSendersChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleAuthorizedSendersChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleAuthorizedSendersChanged represents a AuthorizedSendersChanged event raised by the Oracle contract.
type OracleAuthorizedSendersChanged struct {
	Senders   []common.Address
	ChangedBy common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAuthorizedSendersChanged is a free log retrieval operation binding the contract event 0xf263cfb3e4298332e776194610cf9fdc09ccb3ada8b9aa39764d882e11fbf0a0.
//
// Solidity: event AuthorizedSendersChanged(address[] senders, address changedBy)
func (_Oracle *OracleFilterer) FilterAuthorizedSendersChanged(opts *bind.FilterOpts) (*OracleAuthorizedSendersChangedIterator, error) {

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "AuthorizedSendersChanged")
	if err != nil {
		return nil, err
	}
	return &OracleAuthorizedSendersChangedIterator{contract: _Oracle.contract, event: "AuthorizedSendersChanged", logs: logs, sub: sub}, nil
}

// WatchAuthorizedSendersChanged is a free log subscription operation binding the contract event 0xf263cfb3e4298332e776194610cf9fdc09ccb3ada8b9aa39764d882e11fbf0a0.
//
// Solidity: event AuthorizedSendersChanged(address[] senders, address changedBy)
func (_Oracle *OracleFilterer) WatchAuthorizedSendersChanged(opts *bind.WatchOpts, sink chan<- *OracleAuthorizedSendersChanged) (event.Subscription, error) {

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "AuthorizedSendersChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleAuthorizedSendersChanged)
				if err := _Oracle.contract.UnpackLog(event, "AuthorizedSendersChanged", log); err != nil {
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

// ParseAuthorizedSendersChanged is a log parse operation binding the contract event 0xf263cfb3e4298332e776194610cf9fdc09ccb3ada8b9aa39764d882e11fbf0a0.
//
// Solidity: event AuthorizedSendersChanged(address[] senders, address changedBy)
func (_Oracle *OracleFilterer) ParseAuthorizedSendersChanged(log types.Log) (*OracleAuthorizedSendersChanged, error) {
	event := new(OracleAuthorizedSendersChanged)
	if err := _Oracle.contract.UnpackLog(event, "AuthorizedSendersChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleAuthorizedSendersDeactiveIterator is returned from FilterAuthorizedSendersDeactive and is used to iterate over the raw logs and unpacked data for AuthorizedSendersDeactive events raised by the Oracle contract.
type OracleAuthorizedSendersDeactiveIterator struct {
	Event *OracleAuthorizedSendersDeactive // Event containing the contract specifics and raw log

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
func (it *OracleAuthorizedSendersDeactiveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleAuthorizedSendersDeactive)
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
		it.Event = new(OracleAuthorizedSendersDeactive)
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
func (it *OracleAuthorizedSendersDeactiveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleAuthorizedSendersDeactiveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleAuthorizedSendersDeactive represents a AuthorizedSendersDeactive event raised by the Oracle contract.
type OracleAuthorizedSendersDeactive struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAuthorizedSendersDeactive is a free log retrieval operation binding the contract event 0xea3828816a323b8d7ff49d755efd105e7719166d6c76fad97a28eee5eccc3d9a.
//
// Solidity: event AuthorizedSendersDeactive(address account)
func (_Oracle *OracleFilterer) FilterAuthorizedSendersDeactive(opts *bind.FilterOpts) (*OracleAuthorizedSendersDeactiveIterator, error) {

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "AuthorizedSendersDeactive")
	if err != nil {
		return nil, err
	}
	return &OracleAuthorizedSendersDeactiveIterator{contract: _Oracle.contract, event: "AuthorizedSendersDeactive", logs: logs, sub: sub}, nil
}

// WatchAuthorizedSendersDeactive is a free log subscription operation binding the contract event 0xea3828816a323b8d7ff49d755efd105e7719166d6c76fad97a28eee5eccc3d9a.
//
// Solidity: event AuthorizedSendersDeactive(address account)
func (_Oracle *OracleFilterer) WatchAuthorizedSendersDeactive(opts *bind.WatchOpts, sink chan<- *OracleAuthorizedSendersDeactive) (event.Subscription, error) {

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "AuthorizedSendersDeactive")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleAuthorizedSendersDeactive)
				if err := _Oracle.contract.UnpackLog(event, "AuthorizedSendersDeactive", log); err != nil {
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

// ParseAuthorizedSendersDeactive is a log parse operation binding the contract event 0xea3828816a323b8d7ff49d755efd105e7719166d6c76fad97a28eee5eccc3d9a.
//
// Solidity: event AuthorizedSendersDeactive(address account)
func (_Oracle *OracleFilterer) ParseAuthorizedSendersDeactive(log types.Log) (*OracleAuthorizedSendersDeactive, error) {
	event := new(OracleAuthorizedSendersDeactive)
	if err := _Oracle.contract.UnpackLog(event, "AuthorizedSendersDeactive", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleConfigSetIterator is returned from FilterConfigSet and is used to iterate over the raw logs and unpacked data for ConfigSet events raised by the Oracle contract.
type OracleConfigSetIterator struct {
	Event *OracleConfigSet // Event containing the contract specifics and raw log

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
func (it *OracleConfigSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleConfigSet)
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
		it.Event = new(OracleConfigSet)
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
func (it *OracleConfigSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleConfigSet represents a ConfigSet event raised by the Oracle contract.
type OracleConfigSet struct {
	PreviousConfigBlockNumber uint32
	ConfigDigest              [32]byte
	ConfigCount               uint64
	Signers                   []common.Address
	Transmitters              []common.Address
	F                         uint8
	OnchainConfig             []byte
	OffchainConfigVersion     uint64
	OffchainConfig            []byte
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterConfigSet is a free log retrieval operation binding the contract event 0x1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e05.
//
// Solidity: event ConfigSet(uint32 previousConfigBlockNumber, bytes32 configDigest, uint64 configCount, address[] signers, address[] transmitters, uint8 f, bytes onchainConfig, uint64 offchainConfigVersion, bytes offchainConfig)
func (_Oracle *OracleFilterer) FilterConfigSet(opts *bind.FilterOpts) (*OracleConfigSetIterator, error) {

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &OracleConfigSetIterator{contract: _Oracle.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

// WatchConfigSet is a free log subscription operation binding the contract event 0x1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e05.
//
// Solidity: event ConfigSet(uint32 previousConfigBlockNumber, bytes32 configDigest, uint64 configCount, address[] signers, address[] transmitters, uint8 f, bytes onchainConfig, uint64 offchainConfigVersion, bytes offchainConfig)
func (_Oracle *OracleFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *OracleConfigSet) (event.Subscription, error) {

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleConfigSet)
				if err := _Oracle.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

// ParseConfigSet is a log parse operation binding the contract event 0x1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e05.
//
// Solidity: event ConfigSet(uint32 previousConfigBlockNumber, bytes32 configDigest, uint64 configCount, address[] signers, address[] transmitters, uint8 f, bytes onchainConfig, uint64 offchainConfigVersion, bytes offchainConfig)
func (_Oracle *OracleFilterer) ParseConfigSet(log types.Log) (*OracleConfigSet, error) {
	event := new(OracleConfigSet)
	if err := _Oracle.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Oracle contract.
type OracleInitializedIterator struct {
	Event *OracleInitialized // Event containing the contract specifics and raw log

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
func (it *OracleInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleInitialized)
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
		it.Event = new(OracleInitialized)
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
func (it *OracleInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleInitialized represents a Initialized event raised by the Oracle contract.
type OracleInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Oracle *OracleFilterer) FilterInitialized(opts *bind.FilterOpts) (*OracleInitializedIterator, error) {

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &OracleInitializedIterator{contract: _Oracle.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Oracle *OracleFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *OracleInitialized) (event.Subscription, error) {

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleInitialized)
				if err := _Oracle.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Oracle *OracleFilterer) ParseInitialized(log types.Log) (*OracleInitialized, error) {
	event := new(OracleInitialized)
	if err := _Oracle.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleOracleRequestIterator is returned from FilterOracleRequest and is used to iterate over the raw logs and unpacked data for OracleRequest events raised by the Oracle contract.
type OracleOracleRequestIterator struct {
	Event *OracleOracleRequest // Event containing the contract specifics and raw log

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
func (it *OracleOracleRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleOracleRequest)
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
		it.Event = new(OracleOracleRequest)
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
func (it *OracleOracleRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleOracleRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleOracleRequest represents a OracleRequest event raised by the Oracle contract.
type OracleOracleRequest struct {
	RequestId          [32]byte
	RequestingContract common.Address
	RequestInitiator   common.Address
	SubscriptionId     uint64
	SubscriptionOwner  common.Address
	Data               []byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterOracleRequest is a free log retrieval operation binding the contract event 0xa1ec73989d79578cd6f67d4f593ac3e0a4d1020e5c0164db52108d7ff785406c.
//
// Solidity: event OracleRequest(bytes32 indexed requestId, address requestingContract, address requestInitiator, uint64 subscriptionId, address subscriptionOwner, bytes data)
func (_Oracle *OracleFilterer) FilterOracleRequest(opts *bind.FilterOpts, requestId [][32]byte) (*OracleOracleRequestIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "OracleRequest", requestIdRule)
	if err != nil {
		return nil, err
	}
	return &OracleOracleRequestIterator{contract: _Oracle.contract, event: "OracleRequest", logs: logs, sub: sub}, nil
}

// WatchOracleRequest is a free log subscription operation binding the contract event 0xa1ec73989d79578cd6f67d4f593ac3e0a4d1020e5c0164db52108d7ff785406c.
//
// Solidity: event OracleRequest(bytes32 indexed requestId, address requestingContract, address requestInitiator, uint64 subscriptionId, address subscriptionOwner, bytes data)
func (_Oracle *OracleFilterer) WatchOracleRequest(opts *bind.WatchOpts, sink chan<- *OracleOracleRequest, requestId [][32]byte) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "OracleRequest", requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleOracleRequest)
				if err := _Oracle.contract.UnpackLog(event, "OracleRequest", log); err != nil {
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

// ParseOracleRequest is a log parse operation binding the contract event 0xa1ec73989d79578cd6f67d4f593ac3e0a4d1020e5c0164db52108d7ff785406c.
//
// Solidity: event OracleRequest(bytes32 indexed requestId, address requestingContract, address requestInitiator, uint64 subscriptionId, address subscriptionOwner, bytes data)
func (_Oracle *OracleFilterer) ParseOracleRequest(log types.Log) (*OracleOracleRequest, error) {
	event := new(OracleOracleRequest)
	if err := _Oracle.contract.UnpackLog(event, "OracleRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleOracleResponseIterator is returned from FilterOracleResponse and is used to iterate over the raw logs and unpacked data for OracleResponse events raised by the Oracle contract.
type OracleOracleResponseIterator struct {
	Event *OracleOracleResponse // Event containing the contract specifics and raw log

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
func (it *OracleOracleResponseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleOracleResponse)
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
		it.Event = new(OracleOracleResponse)
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
func (it *OracleOracleResponseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleOracleResponseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleOracleResponse represents a OracleResponse event raised by the Oracle contract.
type OracleOracleResponse struct {
	RequestId [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOracleResponse is a free log retrieval operation binding the contract event 0x9e9bc7616d42c2835d05ae617e508454e63b30b934be8aa932ebc125e0e58a64.
//
// Solidity: event OracleResponse(bytes32 indexed requestId)
func (_Oracle *OracleFilterer) FilterOracleResponse(opts *bind.FilterOpts, requestId [][32]byte) (*OracleOracleResponseIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "OracleResponse", requestIdRule)
	if err != nil {
		return nil, err
	}
	return &OracleOracleResponseIterator{contract: _Oracle.contract, event: "OracleResponse", logs: logs, sub: sub}, nil
}

// WatchOracleResponse is a free log subscription operation binding the contract event 0x9e9bc7616d42c2835d05ae617e508454e63b30b934be8aa932ebc125e0e58a64.
//
// Solidity: event OracleResponse(bytes32 indexed requestId)
func (_Oracle *OracleFilterer) WatchOracleResponse(opts *bind.WatchOpts, sink chan<- *OracleOracleResponse, requestId [][32]byte) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "OracleResponse", requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleOracleResponse)
				if err := _Oracle.contract.UnpackLog(event, "OracleResponse", log); err != nil {
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

// ParseOracleResponse is a log parse operation binding the contract event 0x9e9bc7616d42c2835d05ae617e508454e63b30b934be8aa932ebc125e0e58a64.
//
// Solidity: event OracleResponse(bytes32 indexed requestId)
func (_Oracle *OracleFilterer) ParseOracleResponse(log types.Log) (*OracleOracleResponse, error) {
	event := new(OracleOracleResponse)
	if err := _Oracle.contract.UnpackLog(event, "OracleResponse", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleOwnershipTransferRequestedIterator is returned from FilterOwnershipTransferRequested and is used to iterate over the raw logs and unpacked data for OwnershipTransferRequested events raised by the Oracle contract.
type OracleOwnershipTransferRequestedIterator struct {
	Event *OracleOwnershipTransferRequested // Event containing the contract specifics and raw log

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
func (it *OracleOwnershipTransferRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleOwnershipTransferRequested)
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
		it.Event = new(OracleOwnershipTransferRequested)
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
func (it *OracleOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleOwnershipTransferRequested represents a OwnershipTransferRequested event raised by the Oracle contract.
type OracleOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferRequested is a free log retrieval operation binding the contract event 0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278.
//
// Solidity: event OwnershipTransferRequested(address indexed from, address indexed to)
func (_Oracle *OracleFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OracleOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OracleOwnershipTransferRequestedIterator{contract: _Oracle.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferRequested is a free log subscription operation binding the contract event 0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278.
//
// Solidity: event OwnershipTransferRequested(address indexed from, address indexed to)
func (_Oracle *OracleFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *OracleOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleOwnershipTransferRequested)
				if err := _Oracle.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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
func (_Oracle *OracleFilterer) ParseOwnershipTransferRequested(log types.Log) (*OracleOwnershipTransferRequested, error) {
	event := new(OracleOwnershipTransferRequested)
	if err := _Oracle.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Oracle contract.
type OracleOwnershipTransferredIterator struct {
	Event *OracleOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OracleOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleOwnershipTransferred)
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
		it.Event = new(OracleOwnershipTransferred)
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
func (it *OracleOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleOwnershipTransferred represents a OwnershipTransferred event raised by the Oracle contract.
type OracleOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed from, address indexed to)
func (_Oracle *OracleFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OracleOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OracleOwnershipTransferredIterator{contract: _Oracle.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed from, address indexed to)
func (_Oracle *OracleFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OracleOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleOwnershipTransferred)
				if err := _Oracle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Oracle *OracleFilterer) ParseOwnershipTransferred(log types.Log) (*OracleOwnershipTransferred, error) {
	event := new(OracleOwnershipTransferred)
	if err := _Oracle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleTransmittedIterator is returned from FilterTransmitted and is used to iterate over the raw logs and unpacked data for Transmitted events raised by the Oracle contract.
type OracleTransmittedIterator struct {
	Event *OracleTransmitted // Event containing the contract specifics and raw log

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
func (it *OracleTransmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleTransmitted)
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
		it.Event = new(OracleTransmitted)
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
func (it *OracleTransmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleTransmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleTransmitted represents a Transmitted event raised by the Oracle contract.
type OracleTransmitted struct {
	ConfigDigest [32]byte
	Epoch        uint32
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTransmitted is a free log retrieval operation binding the contract event 0xb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62.
//
// Solidity: event Transmitted(bytes32 configDigest, uint32 epoch)
func (_Oracle *OracleFilterer) FilterTransmitted(opts *bind.FilterOpts) (*OracleTransmittedIterator, error) {

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "Transmitted")
	if err != nil {
		return nil, err
	}
	return &OracleTransmittedIterator{contract: _Oracle.contract, event: "Transmitted", logs: logs, sub: sub}, nil
}

// WatchTransmitted is a free log subscription operation binding the contract event 0xb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62.
//
// Solidity: event Transmitted(bytes32 configDigest, uint32 epoch)
func (_Oracle *OracleFilterer) WatchTransmitted(opts *bind.WatchOpts, sink chan<- *OracleTransmitted) (event.Subscription, error) {

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "Transmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleTransmitted)
				if err := _Oracle.contract.UnpackLog(event, "Transmitted", log); err != nil {
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

// ParseTransmitted is a log parse operation binding the contract event 0xb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62.
//
// Solidity: event Transmitted(bytes32 configDigest, uint32 epoch)
func (_Oracle *OracleFilterer) ParseTransmitted(log types.Log) (*OracleTransmitted, error) {
	event := new(OracleTransmitted)
	if err := _Oracle.contract.UnpackLog(event, "Transmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleUserCallbackErrorIterator is returned from FilterUserCallbackError and is used to iterate over the raw logs and unpacked data for UserCallbackError events raised by the Oracle contract.
type OracleUserCallbackErrorIterator struct {
	Event *OracleUserCallbackError // Event containing the contract specifics and raw log

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
func (it *OracleUserCallbackErrorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleUserCallbackError)
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
		it.Event = new(OracleUserCallbackError)
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
func (it *OracleUserCallbackErrorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleUserCallbackErrorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleUserCallbackError represents a UserCallbackError event raised by the Oracle contract.
type OracleUserCallbackError struct {
	RequestId [32]byte
	Reason    string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUserCallbackError is a free log retrieval operation binding the contract event 0xb2931868c372fe17a25643458add467d60ec5c51125a99b7309f41f5bcd2da6c.
//
// Solidity: event UserCallbackError(bytes32 indexed requestId, string reason)
func (_Oracle *OracleFilterer) FilterUserCallbackError(opts *bind.FilterOpts, requestId [][32]byte) (*OracleUserCallbackErrorIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "UserCallbackError", requestIdRule)
	if err != nil {
		return nil, err
	}
	return &OracleUserCallbackErrorIterator{contract: _Oracle.contract, event: "UserCallbackError", logs: logs, sub: sub}, nil
}

// WatchUserCallbackError is a free log subscription operation binding the contract event 0xb2931868c372fe17a25643458add467d60ec5c51125a99b7309f41f5bcd2da6c.
//
// Solidity: event UserCallbackError(bytes32 indexed requestId, string reason)
func (_Oracle *OracleFilterer) WatchUserCallbackError(opts *bind.WatchOpts, sink chan<- *OracleUserCallbackError, requestId [][32]byte) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "UserCallbackError", requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleUserCallbackError)
				if err := _Oracle.contract.UnpackLog(event, "UserCallbackError", log); err != nil {
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

// ParseUserCallbackError is a log parse operation binding the contract event 0xb2931868c372fe17a25643458add467d60ec5c51125a99b7309f41f5bcd2da6c.
//
// Solidity: event UserCallbackError(bytes32 indexed requestId, string reason)
func (_Oracle *OracleFilterer) ParseUserCallbackError(log types.Log) (*OracleUserCallbackError, error) {
	event := new(OracleUserCallbackError)
	if err := _Oracle.contract.UnpackLog(event, "UserCallbackError", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleUserCallbackRawErrorIterator is returned from FilterUserCallbackRawError and is used to iterate over the raw logs and unpacked data for UserCallbackRawError events raised by the Oracle contract.
type OracleUserCallbackRawErrorIterator struct {
	Event *OracleUserCallbackRawError // Event containing the contract specifics and raw log

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
func (it *OracleUserCallbackRawErrorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleUserCallbackRawError)
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
		it.Event = new(OracleUserCallbackRawError)
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
func (it *OracleUserCallbackRawErrorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleUserCallbackRawErrorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleUserCallbackRawError represents a UserCallbackRawError event raised by the Oracle contract.
type OracleUserCallbackRawError struct {
	RequestId    [32]byte
	LowLevelData []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUserCallbackRawError is a free log retrieval operation binding the contract event 0xe0b838ffe6ee22a0d3acf19a85db6a41b34a1ab739e2d6c759a2e42d95bdccb2.
//
// Solidity: event UserCallbackRawError(bytes32 indexed requestId, bytes lowLevelData)
func (_Oracle *OracleFilterer) FilterUserCallbackRawError(opts *bind.FilterOpts, requestId [][32]byte) (*OracleUserCallbackRawErrorIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _Oracle.contract.FilterLogs(opts, "UserCallbackRawError", requestIdRule)
	if err != nil {
		return nil, err
	}
	return &OracleUserCallbackRawErrorIterator{contract: _Oracle.contract, event: "UserCallbackRawError", logs: logs, sub: sub}, nil
}

// WatchUserCallbackRawError is a free log subscription operation binding the contract event 0xe0b838ffe6ee22a0d3acf19a85db6a41b34a1ab739e2d6c759a2e42d95bdccb2.
//
// Solidity: event UserCallbackRawError(bytes32 indexed requestId, bytes lowLevelData)
func (_Oracle *OracleFilterer) WatchUserCallbackRawError(opts *bind.WatchOpts, sink chan<- *OracleUserCallbackRawError, requestId [][32]byte) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _Oracle.contract.WatchLogs(opts, "UserCallbackRawError", requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleUserCallbackRawError)
				if err := _Oracle.contract.UnpackLog(event, "UserCallbackRawError", log); err != nil {
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

// ParseUserCallbackRawError is a log parse operation binding the contract event 0xe0b838ffe6ee22a0d3acf19a85db6a41b34a1ab739e2d6c759a2e42d95bdccb2.
//
// Solidity: event UserCallbackRawError(bytes32 indexed requestId, bytes lowLevelData)
func (_Oracle *OracleFilterer) ParseUserCallbackRawError(log types.Log) (*OracleUserCallbackRawError, error) {
	event := new(OracleUserCallbackRawError)
	if err := _Oracle.contract.UnpackLog(event, "UserCallbackRawError", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
