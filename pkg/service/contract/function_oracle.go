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

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AlreadySet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotSelfTransfer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyBillingRegistry\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyPublicKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyRequestData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptySendersList\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InconsistentReportData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRequestID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAllowedToSetSenders\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotProposedOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerMustBeSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReportInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnauthorizedPublicKeyChange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnauthorizedSender\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AuthorizedSendersActive\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"senders\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"changedBy\",\"type\":\"address\"}],\"name\":\"AuthorizedSendersChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AuthorizedSendersDeactive\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"previousConfigBlockNumber\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"configCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"transmitters\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"f\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"onchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"offchainConfigVersion\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"requestingContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"requestInitiator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"subscriptionId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"subscriptionOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"OracleRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"}],\"name\":\"OracleResponse\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"name\":\"Transmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"UserCallbackError\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"lowLevelData\",\"type\":\"bytes\"}],\"name\":\"UserCallbackRawError\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activateAuthorizedReceiver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"senders\",\"type\":\"address[]\"}],\"name\":\"addAuthorizedSenders\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"authorizedReceiverActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deactivateAuthorizedReceiver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"}],\"name\":\"deleteNodePublicKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"subscriptionId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"gasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"}],\"name\":\"estimateCost\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllNodePublicKeys\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAuthorizedSenders\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDONPublicKey\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"subscriptionId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"client\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"gasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"}],\"internalType\":\"structFunctionsBillingRegistryInterface.RequestBilling\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"getRequiredFee\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"isAuthorizedSender\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"configCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfigDigestAndEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"scanLogs\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"configDigest\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"epoch\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"senders\",\"type\":\"address[]\"}],\"name\":\"removeAuthorizedSenders\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"subscriptionId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"gasLimit\",\"type\":\"uint32\"}],\"name\":\"sendRequest\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_transmitters\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"_f\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"_onchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"_offchainConfigVersion\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_offchainConfig\",\"type\":\"bytes\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"donPublicKey\",\"type\":\"bytes\"}],\"name\":\"setDONPublicKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"}],\"name\":\"setNodePublicKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"registryAddress\",\"type\":\"address\"}],\"name\":\"setRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[3]\",\"name\":\"reportContext\",\"type\":\"bytes32[3]\"},{\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"rs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"ss\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"rawVs\",\"type\":\"bytes32\"}],\"name\":\"transmit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transmitters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// AuthorizedReceiverActive is a free data retrieval call binding the contract method 0x4b4fa0c1.
//
// Solidity: function authorizedReceiverActive() view returns(bool)
func (_Contract *ContractCaller) AuthorizedReceiverActive(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "authorizedReceiverActive")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AuthorizedReceiverActive is a free data retrieval call binding the contract method 0x4b4fa0c1.
//
// Solidity: function authorizedReceiverActive() view returns(bool)
func (_Contract *ContractSession) AuthorizedReceiverActive() (bool, error) {
	return _Contract.Contract.AuthorizedReceiverActive(&_Contract.CallOpts)
}

// AuthorizedReceiverActive is a free data retrieval call binding the contract method 0x4b4fa0c1.
//
// Solidity: function authorizedReceiverActive() view returns(bool)
func (_Contract *ContractCallerSession) AuthorizedReceiverActive() (bool, error) {
	return _Contract.Contract.AuthorizedReceiverActive(&_Contract.CallOpts)
}

// EstimateCost is a free data retrieval call binding the contract method 0xd227d245.
//
// Solidity: function estimateCost(uint64 subscriptionId, bytes data, uint32 gasLimit, uint256 gasPrice) view returns(uint96)
func (_Contract *ContractCaller) EstimateCost(opts *bind.CallOpts, subscriptionId uint64, data []byte, gasLimit uint32, gasPrice *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "estimateCost", subscriptionId, data, gasLimit, gasPrice)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimateCost is a free data retrieval call binding the contract method 0xd227d245.
//
// Solidity: function estimateCost(uint64 subscriptionId, bytes data, uint32 gasLimit, uint256 gasPrice) view returns(uint96)
func (_Contract *ContractSession) EstimateCost(subscriptionId uint64, data []byte, gasLimit uint32, gasPrice *big.Int) (*big.Int, error) {
	return _Contract.Contract.EstimateCost(&_Contract.CallOpts, subscriptionId, data, gasLimit, gasPrice)
}

// EstimateCost is a free data retrieval call binding the contract method 0xd227d245.
//
// Solidity: function estimateCost(uint64 subscriptionId, bytes data, uint32 gasLimit, uint256 gasPrice) view returns(uint96)
func (_Contract *ContractCallerSession) EstimateCost(subscriptionId uint64, data []byte, gasLimit uint32, gasPrice *big.Int) (*big.Int, error) {
	return _Contract.Contract.EstimateCost(&_Contract.CallOpts, subscriptionId, data, gasLimit, gasPrice)
}

// GetAllNodePublicKeys is a free data retrieval call binding the contract method 0x53398987.
//
// Solidity: function getAllNodePublicKeys() view returns(address[], bytes[])
func (_Contract *ContractCaller) GetAllNodePublicKeys(opts *bind.CallOpts) ([]common.Address, [][]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getAllNodePublicKeys")

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
func (_Contract *ContractSession) GetAllNodePublicKeys() ([]common.Address, [][]byte, error) {
	return _Contract.Contract.GetAllNodePublicKeys(&_Contract.CallOpts)
}

// GetAllNodePublicKeys is a free data retrieval call binding the contract method 0x53398987.
//
// Solidity: function getAllNodePublicKeys() view returns(address[], bytes[])
func (_Contract *ContractCallerSession) GetAllNodePublicKeys() ([]common.Address, [][]byte, error) {
	return _Contract.Contract.GetAllNodePublicKeys(&_Contract.CallOpts)
}

// GetAuthorizedSenders is a free data retrieval call binding the contract method 0x2408afaa.
//
// Solidity: function getAuthorizedSenders() view returns(address[])
func (_Contract *ContractCaller) GetAuthorizedSenders(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getAuthorizedSenders")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAuthorizedSenders is a free data retrieval call binding the contract method 0x2408afaa.
//
// Solidity: function getAuthorizedSenders() view returns(address[])
func (_Contract *ContractSession) GetAuthorizedSenders() ([]common.Address, error) {
	return _Contract.Contract.GetAuthorizedSenders(&_Contract.CallOpts)
}

// GetAuthorizedSenders is a free data retrieval call binding the contract method 0x2408afaa.
//
// Solidity: function getAuthorizedSenders() view returns(address[])
func (_Contract *ContractCallerSession) GetAuthorizedSenders() ([]common.Address, error) {
	return _Contract.Contract.GetAuthorizedSenders(&_Contract.CallOpts)
}

// GetDONPublicKey is a free data retrieval call binding the contract method 0xd328a91e.
//
// Solidity: function getDONPublicKey() view returns(bytes)
func (_Contract *ContractCaller) GetDONPublicKey(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getDONPublicKey")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetDONPublicKey is a free data retrieval call binding the contract method 0xd328a91e.
//
// Solidity: function getDONPublicKey() view returns(bytes)
func (_Contract *ContractSession) GetDONPublicKey() ([]byte, error) {
	return _Contract.Contract.GetDONPublicKey(&_Contract.CallOpts)
}

// GetDONPublicKey is a free data retrieval call binding the contract method 0xd328a91e.
//
// Solidity: function getDONPublicKey() view returns(bytes)
func (_Contract *ContractCallerSession) GetDONPublicKey() ([]byte, error) {
	return _Contract.Contract.GetDONPublicKey(&_Contract.CallOpts)
}

// GetRegistry is a free data retrieval call binding the contract method 0x5ab1bd53.
//
// Solidity: function getRegistry() view returns(address)
func (_Contract *ContractCaller) GetRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRegistry is a free data retrieval call binding the contract method 0x5ab1bd53.
//
// Solidity: function getRegistry() view returns(address)
func (_Contract *ContractSession) GetRegistry() (common.Address, error) {
	return _Contract.Contract.GetRegistry(&_Contract.CallOpts)
}

// GetRegistry is a free data retrieval call binding the contract method 0x5ab1bd53.
//
// Solidity: function getRegistry() view returns(address)
func (_Contract *ContractCallerSession) GetRegistry() (common.Address, error) {
	return _Contract.Contract.GetRegistry(&_Contract.CallOpts)
}

// GetRequiredFee is a free data retrieval call binding the contract method 0xf1e14a21.
//
// Solidity: function getRequiredFee(bytes , (uint64,address,uint32,uint256) ) pure returns(uint96)
func (_Contract *ContractCaller) GetRequiredFee(opts *bind.CallOpts, arg0 []byte, arg1 FunctionsBillingRegistryInterfaceRequestBilling) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getRequiredFee", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRequiredFee is a free data retrieval call binding the contract method 0xf1e14a21.
//
// Solidity: function getRequiredFee(bytes , (uint64,address,uint32,uint256) ) pure returns(uint96)
func (_Contract *ContractSession) GetRequiredFee(arg0 []byte, arg1 FunctionsBillingRegistryInterfaceRequestBilling) (*big.Int, error) {
	return _Contract.Contract.GetRequiredFee(&_Contract.CallOpts, arg0, arg1)
}

// GetRequiredFee is a free data retrieval call binding the contract method 0xf1e14a21.
//
// Solidity: function getRequiredFee(bytes , (uint64,address,uint32,uint256) ) pure returns(uint96)
func (_Contract *ContractCallerSession) GetRequiredFee(arg0 []byte, arg1 FunctionsBillingRegistryInterfaceRequestBilling) (*big.Int, error) {
	return _Contract.Contract.GetRequiredFee(&_Contract.CallOpts, arg0, arg1)
}

// IsAuthorizedSender is a free data retrieval call binding the contract method 0xfa00763a.
//
// Solidity: function isAuthorizedSender(address sender) view returns(bool)
func (_Contract *ContractCaller) IsAuthorizedSender(opts *bind.CallOpts, sender common.Address) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "isAuthorizedSender", sender)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAuthorizedSender is a free data retrieval call binding the contract method 0xfa00763a.
//
// Solidity: function isAuthorizedSender(address sender) view returns(bool)
func (_Contract *ContractSession) IsAuthorizedSender(sender common.Address) (bool, error) {
	return _Contract.Contract.IsAuthorizedSender(&_Contract.CallOpts, sender)
}

// IsAuthorizedSender is a free data retrieval call binding the contract method 0xfa00763a.
//
// Solidity: function isAuthorizedSender(address sender) view returns(bool)
func (_Contract *ContractCallerSession) IsAuthorizedSender(sender common.Address) (bool, error) {
	return _Contract.Contract.IsAuthorizedSender(&_Contract.CallOpts, sender)
}

// LatestConfigDetails is a free data retrieval call binding the contract method 0x81ff7048.
//
// Solidity: function latestConfigDetails() view returns(uint32 configCount, uint32 blockNumber, bytes32 configDigest)
func (_Contract *ContractCaller) LatestConfigDetails(opts *bind.CallOpts) (struct {
	ConfigCount  uint32
	BlockNumber  uint32
	ConfigDigest [32]byte
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "latestConfigDetails")

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
func (_Contract *ContractSession) LatestConfigDetails() (struct {
	ConfigCount  uint32
	BlockNumber  uint32
	ConfigDigest [32]byte
}, error) {
	return _Contract.Contract.LatestConfigDetails(&_Contract.CallOpts)
}

// LatestConfigDetails is a free data retrieval call binding the contract method 0x81ff7048.
//
// Solidity: function latestConfigDetails() view returns(uint32 configCount, uint32 blockNumber, bytes32 configDigest)
func (_Contract *ContractCallerSession) LatestConfigDetails() (struct {
	ConfigCount  uint32
	BlockNumber  uint32
	ConfigDigest [32]byte
}, error) {
	return _Contract.Contract.LatestConfigDetails(&_Contract.CallOpts)
}

// LatestConfigDigestAndEpoch is a free data retrieval call binding the contract method 0xafcb95d7.
//
// Solidity: function latestConfigDigestAndEpoch() view returns(bool scanLogs, bytes32 configDigest, uint32 epoch)
func (_Contract *ContractCaller) LatestConfigDigestAndEpoch(opts *bind.CallOpts) (struct {
	ScanLogs     bool
	ConfigDigest [32]byte
	Epoch        uint32
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "latestConfigDigestAndEpoch")

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
func (_Contract *ContractSession) LatestConfigDigestAndEpoch() (struct {
	ScanLogs     bool
	ConfigDigest [32]byte
	Epoch        uint32
}, error) {
	return _Contract.Contract.LatestConfigDigestAndEpoch(&_Contract.CallOpts)
}

// LatestConfigDigestAndEpoch is a free data retrieval call binding the contract method 0xafcb95d7.
//
// Solidity: function latestConfigDigestAndEpoch() view returns(bool scanLogs, bytes32 configDigest, uint32 epoch)
func (_Contract *ContractCallerSession) LatestConfigDigestAndEpoch() (struct {
	ScanLogs     bool
	ConfigDigest [32]byte
	Epoch        uint32
}, error) {
	return _Contract.Contract.LatestConfigDigestAndEpoch(&_Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCallerSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// Transmitters is a free data retrieval call binding the contract method 0x81411834.
//
// Solidity: function transmitters() view returns(address[])
func (_Contract *ContractCaller) Transmitters(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "transmitters")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// Transmitters is a free data retrieval call binding the contract method 0x81411834.
//
// Solidity: function transmitters() view returns(address[])
func (_Contract *ContractSession) Transmitters() ([]common.Address, error) {
	return _Contract.Contract.Transmitters(&_Contract.CallOpts)
}

// Transmitters is a free data retrieval call binding the contract method 0x81411834.
//
// Solidity: function transmitters() view returns(address[])
func (_Contract *ContractCallerSession) Transmitters() ([]common.Address, error) {
	return _Contract.Contract.Transmitters(&_Contract.CallOpts)
}

// TypeAndVersion is a free data retrieval call binding the contract method 0x181f5a77.
//
// Solidity: function typeAndVersion() pure returns(string)
func (_Contract *ContractCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TypeAndVersion is a free data retrieval call binding the contract method 0x181f5a77.
//
// Solidity: function typeAndVersion() pure returns(string)
func (_Contract *ContractSession) TypeAndVersion() (string, error) {
	return _Contract.Contract.TypeAndVersion(&_Contract.CallOpts)
}

// TypeAndVersion is a free data retrieval call binding the contract method 0x181f5a77.
//
// Solidity: function typeAndVersion() pure returns(string)
func (_Contract *ContractCallerSession) TypeAndVersion() (string, error) {
	return _Contract.Contract.TypeAndVersion(&_Contract.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Contract *ContractTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Contract *ContractSession) AcceptOwnership() (*types.Transaction, error) {
	return _Contract.Contract.AcceptOwnership(&_Contract.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Contract *ContractTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Contract.Contract.AcceptOwnership(&_Contract.TransactOpts)
}

// ActivateAuthorizedReceiver is a paid mutator transaction binding the contract method 0x110254c8.
//
// Solidity: function activateAuthorizedReceiver() returns()
func (_Contract *ContractTransactor) ActivateAuthorizedReceiver(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "activateAuthorizedReceiver")
}

// ActivateAuthorizedReceiver is a paid mutator transaction binding the contract method 0x110254c8.
//
// Solidity: function activateAuthorizedReceiver() returns()
func (_Contract *ContractSession) ActivateAuthorizedReceiver() (*types.Transaction, error) {
	return _Contract.Contract.ActivateAuthorizedReceiver(&_Contract.TransactOpts)
}

// ActivateAuthorizedReceiver is a paid mutator transaction binding the contract method 0x110254c8.
//
// Solidity: function activateAuthorizedReceiver() returns()
func (_Contract *ContractTransactorSession) ActivateAuthorizedReceiver() (*types.Transaction, error) {
	return _Contract.Contract.ActivateAuthorizedReceiver(&_Contract.TransactOpts)
}

// AddAuthorizedSenders is a paid mutator transaction binding the contract method 0x4dcef404.
//
// Solidity: function addAuthorizedSenders(address[] senders) returns()
func (_Contract *ContractTransactor) AddAuthorizedSenders(opts *bind.TransactOpts, senders []common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "addAuthorizedSenders", senders)
}

// AddAuthorizedSenders is a paid mutator transaction binding the contract method 0x4dcef404.
//
// Solidity: function addAuthorizedSenders(address[] senders) returns()
func (_Contract *ContractSession) AddAuthorizedSenders(senders []common.Address) (*types.Transaction, error) {
	return _Contract.Contract.AddAuthorizedSenders(&_Contract.TransactOpts, senders)
}

// AddAuthorizedSenders is a paid mutator transaction binding the contract method 0x4dcef404.
//
// Solidity: function addAuthorizedSenders(address[] senders) returns()
func (_Contract *ContractTransactorSession) AddAuthorizedSenders(senders []common.Address) (*types.Transaction, error) {
	return _Contract.Contract.AddAuthorizedSenders(&_Contract.TransactOpts, senders)
}

// DeactivateAuthorizedReceiver is a paid mutator transaction binding the contract method 0x91bb64eb.
//
// Solidity: function deactivateAuthorizedReceiver() returns()
func (_Contract *ContractTransactor) DeactivateAuthorizedReceiver(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "deactivateAuthorizedReceiver")
}

// DeactivateAuthorizedReceiver is a paid mutator transaction binding the contract method 0x91bb64eb.
//
// Solidity: function deactivateAuthorizedReceiver() returns()
func (_Contract *ContractSession) DeactivateAuthorizedReceiver() (*types.Transaction, error) {
	return _Contract.Contract.DeactivateAuthorizedReceiver(&_Contract.TransactOpts)
}

// DeactivateAuthorizedReceiver is a paid mutator transaction binding the contract method 0x91bb64eb.
//
// Solidity: function deactivateAuthorizedReceiver() returns()
func (_Contract *ContractTransactorSession) DeactivateAuthorizedReceiver() (*types.Transaction, error) {
	return _Contract.Contract.DeactivateAuthorizedReceiver(&_Contract.TransactOpts)
}

// DeleteNodePublicKey is a paid mutator transaction binding the contract method 0x26ceabac.
//
// Solidity: function deleteNodePublicKey(address node) returns()
func (_Contract *ContractTransactor) DeleteNodePublicKey(opts *bind.TransactOpts, node common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "deleteNodePublicKey", node)
}

// DeleteNodePublicKey is a paid mutator transaction binding the contract method 0x26ceabac.
//
// Solidity: function deleteNodePublicKey(address node) returns()
func (_Contract *ContractSession) DeleteNodePublicKey(node common.Address) (*types.Transaction, error) {
	return _Contract.Contract.DeleteNodePublicKey(&_Contract.TransactOpts, node)
}

// DeleteNodePublicKey is a paid mutator transaction binding the contract method 0x26ceabac.
//
// Solidity: function deleteNodePublicKey(address node) returns()
func (_Contract *ContractTransactorSession) DeleteNodePublicKey(node common.Address) (*types.Transaction, error) {
	return _Contract.Contract.DeleteNodePublicKey(&_Contract.TransactOpts, node)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Contract *ContractTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Contract *ContractSession) Initialize() (*types.Transaction, error) {
	return _Contract.Contract.Initialize(&_Contract.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Contract *ContractTransactorSession) Initialize() (*types.Transaction, error) {
	return _Contract.Contract.Initialize(&_Contract.TransactOpts)
}

// RemoveAuthorizedSenders is a paid mutator transaction binding the contract method 0x03e1bf23.
//
// Solidity: function removeAuthorizedSenders(address[] senders) returns()
func (_Contract *ContractTransactor) RemoveAuthorizedSenders(opts *bind.TransactOpts, senders []common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "removeAuthorizedSenders", senders)
}

// RemoveAuthorizedSenders is a paid mutator transaction binding the contract method 0x03e1bf23.
//
// Solidity: function removeAuthorizedSenders(address[] senders) returns()
func (_Contract *ContractSession) RemoveAuthorizedSenders(senders []common.Address) (*types.Transaction, error) {
	return _Contract.Contract.RemoveAuthorizedSenders(&_Contract.TransactOpts, senders)
}

// RemoveAuthorizedSenders is a paid mutator transaction binding the contract method 0x03e1bf23.
//
// Solidity: function removeAuthorizedSenders(address[] senders) returns()
func (_Contract *ContractTransactorSession) RemoveAuthorizedSenders(senders []common.Address) (*types.Transaction, error) {
	return _Contract.Contract.RemoveAuthorizedSenders(&_Contract.TransactOpts, senders)
}

// SendRequest is a paid mutator transaction binding the contract method 0x28242b04.
//
// Solidity: function sendRequest(uint64 subscriptionId, bytes data, uint32 gasLimit) returns(bytes32)
func (_Contract *ContractTransactor) SendRequest(opts *bind.TransactOpts, subscriptionId uint64, data []byte, gasLimit uint32) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "sendRequest", subscriptionId, data, gasLimit)
}

// SendRequest is a paid mutator transaction binding the contract method 0x28242b04.
//
// Solidity: function sendRequest(uint64 subscriptionId, bytes data, uint32 gasLimit) returns(bytes32)
func (_Contract *ContractSession) SendRequest(subscriptionId uint64, data []byte, gasLimit uint32) (*types.Transaction, error) {
	return _Contract.Contract.SendRequest(&_Contract.TransactOpts, subscriptionId, data, gasLimit)
}

// SendRequest is a paid mutator transaction binding the contract method 0x28242b04.
//
// Solidity: function sendRequest(uint64 subscriptionId, bytes data, uint32 gasLimit) returns(bytes32)
func (_Contract *ContractTransactorSession) SendRequest(subscriptionId uint64, data []byte, gasLimit uint32) (*types.Transaction, error) {
	return _Contract.Contract.SendRequest(&_Contract.TransactOpts, subscriptionId, data, gasLimit)
}

// SetConfig is a paid mutator transaction binding the contract method 0xe3d0e712.
//
// Solidity: function setConfig(address[] _signers, address[] _transmitters, uint8 _f, bytes _onchainConfig, uint64 _offchainConfigVersion, bytes _offchainConfig) returns()
func (_Contract *ContractTransactor) SetConfig(opts *bind.TransactOpts, _signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setConfig", _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
}

// SetConfig is a paid mutator transaction binding the contract method 0xe3d0e712.
//
// Solidity: function setConfig(address[] _signers, address[] _transmitters, uint8 _f, bytes _onchainConfig, uint64 _offchainConfigVersion, bytes _offchainConfig) returns()
func (_Contract *ContractSession) SetConfig(_signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _Contract.Contract.SetConfig(&_Contract.TransactOpts, _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
}

// SetConfig is a paid mutator transaction binding the contract method 0xe3d0e712.
//
// Solidity: function setConfig(address[] _signers, address[] _transmitters, uint8 _f, bytes _onchainConfig, uint64 _offchainConfigVersion, bytes _offchainConfig) returns()
func (_Contract *ContractTransactorSession) SetConfig(_signers []common.Address, _transmitters []common.Address, _f uint8, _onchainConfig []byte, _offchainConfigVersion uint64, _offchainConfig []byte) (*types.Transaction, error) {
	return _Contract.Contract.SetConfig(&_Contract.TransactOpts, _signers, _transmitters, _f, _onchainConfig, _offchainConfigVersion, _offchainConfig)
}

// SetDONPublicKey is a paid mutator transaction binding the contract method 0x7f15e166.
//
// Solidity: function setDONPublicKey(bytes donPublicKey) returns()
func (_Contract *ContractTransactor) SetDONPublicKey(opts *bind.TransactOpts, donPublicKey []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setDONPublicKey", donPublicKey)
}

// SetDONPublicKey is a paid mutator transaction binding the contract method 0x7f15e166.
//
// Solidity: function setDONPublicKey(bytes donPublicKey) returns()
func (_Contract *ContractSession) SetDONPublicKey(donPublicKey []byte) (*types.Transaction, error) {
	return _Contract.Contract.SetDONPublicKey(&_Contract.TransactOpts, donPublicKey)
}

// SetDONPublicKey is a paid mutator transaction binding the contract method 0x7f15e166.
//
// Solidity: function setDONPublicKey(bytes donPublicKey) returns()
func (_Contract *ContractTransactorSession) SetDONPublicKey(donPublicKey []byte) (*types.Transaction, error) {
	return _Contract.Contract.SetDONPublicKey(&_Contract.TransactOpts, donPublicKey)
}

// SetNodePublicKey is a paid mutator transaction binding the contract method 0x80756031.
//
// Solidity: function setNodePublicKey(address node, bytes publicKey) returns()
func (_Contract *ContractTransactor) SetNodePublicKey(opts *bind.TransactOpts, node common.Address, publicKey []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setNodePublicKey", node, publicKey)
}

// SetNodePublicKey is a paid mutator transaction binding the contract method 0x80756031.
//
// Solidity: function setNodePublicKey(address node, bytes publicKey) returns()
func (_Contract *ContractSession) SetNodePublicKey(node common.Address, publicKey []byte) (*types.Transaction, error) {
	return _Contract.Contract.SetNodePublicKey(&_Contract.TransactOpts, node, publicKey)
}

// SetNodePublicKey is a paid mutator transaction binding the contract method 0x80756031.
//
// Solidity: function setNodePublicKey(address node, bytes publicKey) returns()
func (_Contract *ContractTransactorSession) SetNodePublicKey(node common.Address, publicKey []byte) (*types.Transaction, error) {
	return _Contract.Contract.SetNodePublicKey(&_Contract.TransactOpts, node, publicKey)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address registryAddress) returns()
func (_Contract *ContractTransactor) SetRegistry(opts *bind.TransactOpts, registryAddress common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setRegistry", registryAddress)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address registryAddress) returns()
func (_Contract *ContractSession) SetRegistry(registryAddress common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetRegistry(&_Contract.TransactOpts, registryAddress)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address registryAddress) returns()
func (_Contract *ContractTransactorSession) SetRegistry(registryAddress common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetRegistry(&_Contract.TransactOpts, registryAddress)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address to) returns()
func (_Contract *ContractTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transferOwnership", to)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address to) returns()
func (_Contract *ContractSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, to)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address to) returns()
func (_Contract *ContractTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, to)
}

// Transmit is a paid mutator transaction binding the contract method 0xb1dc65a4.
//
// Solidity: function transmit(bytes32[3] reportContext, bytes report, bytes32[] rs, bytes32[] ss, bytes32 rawVs) returns()
func (_Contract *ContractTransactor) Transmit(opts *bind.TransactOpts, reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transmit", reportContext, report, rs, ss, rawVs)
}

// Transmit is a paid mutator transaction binding the contract method 0xb1dc65a4.
//
// Solidity: function transmit(bytes32[3] reportContext, bytes report, bytes32[] rs, bytes32[] ss, bytes32 rawVs) returns()
func (_Contract *ContractSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.Transmit(&_Contract.TransactOpts, reportContext, report, rs, ss, rawVs)
}

// Transmit is a paid mutator transaction binding the contract method 0xb1dc65a4.
//
// Solidity: function transmit(bytes32[3] reportContext, bytes report, bytes32[] rs, bytes32[] ss, bytes32 rawVs) returns()
func (_Contract *ContractTransactorSession) Transmit(reportContext [3][32]byte, report []byte, rs [][32]byte, ss [][32]byte, rawVs [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.Transmit(&_Contract.TransactOpts, reportContext, report, rs, ss, rawVs)
}

// ContractAuthorizedSendersActiveIterator is returned from FilterAuthorizedSendersActive and is used to iterate over the raw logs and unpacked data for AuthorizedSendersActive events raised by the Contract contract.
type ContractAuthorizedSendersActiveIterator struct {
	Event *ContractAuthorizedSendersActive // Event containing the contract specifics and raw log

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
func (it *ContractAuthorizedSendersActiveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAuthorizedSendersActive)
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
		it.Event = new(ContractAuthorizedSendersActive)
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
func (it *ContractAuthorizedSendersActiveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAuthorizedSendersActiveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAuthorizedSendersActive represents a AuthorizedSendersActive event raised by the Contract contract.
type ContractAuthorizedSendersActive struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAuthorizedSendersActive is a free log retrieval operation binding the contract event 0xae51766a982895b0c444fc99fc1a560762b464d709e6c78376c85617f7eeb5ce.
//
// Solidity: event AuthorizedSendersActive(address account)
func (_Contract *ContractFilterer) FilterAuthorizedSendersActive(opts *bind.FilterOpts) (*ContractAuthorizedSendersActiveIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "AuthorizedSendersActive")
	if err != nil {
		return nil, err
	}
	return &ContractAuthorizedSendersActiveIterator{contract: _Contract.contract, event: "AuthorizedSendersActive", logs: logs, sub: sub}, nil
}

// WatchAuthorizedSendersActive is a free log subscription operation binding the contract event 0xae51766a982895b0c444fc99fc1a560762b464d709e6c78376c85617f7eeb5ce.
//
// Solidity: event AuthorizedSendersActive(address account)
func (_Contract *ContractFilterer) WatchAuthorizedSendersActive(opts *bind.WatchOpts, sink chan<- *ContractAuthorizedSendersActive) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "AuthorizedSendersActive")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAuthorizedSendersActive)
				if err := _Contract.contract.UnpackLog(event, "AuthorizedSendersActive", log); err != nil {
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
func (_Contract *ContractFilterer) ParseAuthorizedSendersActive(log types.Log) (*ContractAuthorizedSendersActive, error) {
	event := new(ContractAuthorizedSendersActive)
	if err := _Contract.contract.UnpackLog(event, "AuthorizedSendersActive", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAuthorizedSendersChangedIterator is returned from FilterAuthorizedSendersChanged and is used to iterate over the raw logs and unpacked data for AuthorizedSendersChanged events raised by the Contract contract.
type ContractAuthorizedSendersChangedIterator struct {
	Event *ContractAuthorizedSendersChanged // Event containing the contract specifics and raw log

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
func (it *ContractAuthorizedSendersChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAuthorizedSendersChanged)
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
		it.Event = new(ContractAuthorizedSendersChanged)
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
func (it *ContractAuthorizedSendersChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAuthorizedSendersChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAuthorizedSendersChanged represents a AuthorizedSendersChanged event raised by the Contract contract.
type ContractAuthorizedSendersChanged struct {
	Senders   []common.Address
	ChangedBy common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAuthorizedSendersChanged is a free log retrieval operation binding the contract event 0xf263cfb3e4298332e776194610cf9fdc09ccb3ada8b9aa39764d882e11fbf0a0.
//
// Solidity: event AuthorizedSendersChanged(address[] senders, address changedBy)
func (_Contract *ContractFilterer) FilterAuthorizedSendersChanged(opts *bind.FilterOpts) (*ContractAuthorizedSendersChangedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "AuthorizedSendersChanged")
	if err != nil {
		return nil, err
	}
	return &ContractAuthorizedSendersChangedIterator{contract: _Contract.contract, event: "AuthorizedSendersChanged", logs: logs, sub: sub}, nil
}

// WatchAuthorizedSendersChanged is a free log subscription operation binding the contract event 0xf263cfb3e4298332e776194610cf9fdc09ccb3ada8b9aa39764d882e11fbf0a0.
//
// Solidity: event AuthorizedSendersChanged(address[] senders, address changedBy)
func (_Contract *ContractFilterer) WatchAuthorizedSendersChanged(opts *bind.WatchOpts, sink chan<- *ContractAuthorizedSendersChanged) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "AuthorizedSendersChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAuthorizedSendersChanged)
				if err := _Contract.contract.UnpackLog(event, "AuthorizedSendersChanged", log); err != nil {
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
func (_Contract *ContractFilterer) ParseAuthorizedSendersChanged(log types.Log) (*ContractAuthorizedSendersChanged, error) {
	event := new(ContractAuthorizedSendersChanged)
	if err := _Contract.contract.UnpackLog(event, "AuthorizedSendersChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractAuthorizedSendersDeactiveIterator is returned from FilterAuthorizedSendersDeactive and is used to iterate over the raw logs and unpacked data for AuthorizedSendersDeactive events raised by the Contract contract.
type ContractAuthorizedSendersDeactiveIterator struct {
	Event *ContractAuthorizedSendersDeactive // Event containing the contract specifics and raw log

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
func (it *ContractAuthorizedSendersDeactiveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAuthorizedSendersDeactive)
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
		it.Event = new(ContractAuthorizedSendersDeactive)
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
func (it *ContractAuthorizedSendersDeactiveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAuthorizedSendersDeactiveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAuthorizedSendersDeactive represents a AuthorizedSendersDeactive event raised by the Contract contract.
type ContractAuthorizedSendersDeactive struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAuthorizedSendersDeactive is a free log retrieval operation binding the contract event 0xea3828816a323b8d7ff49d755efd105e7719166d6c76fad97a28eee5eccc3d9a.
//
// Solidity: event AuthorizedSendersDeactive(address account)
func (_Contract *ContractFilterer) FilterAuthorizedSendersDeactive(opts *bind.FilterOpts) (*ContractAuthorizedSendersDeactiveIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "AuthorizedSendersDeactive")
	if err != nil {
		return nil, err
	}
	return &ContractAuthorizedSendersDeactiveIterator{contract: _Contract.contract, event: "AuthorizedSendersDeactive", logs: logs, sub: sub}, nil
}

// WatchAuthorizedSendersDeactive is a free log subscription operation binding the contract event 0xea3828816a323b8d7ff49d755efd105e7719166d6c76fad97a28eee5eccc3d9a.
//
// Solidity: event AuthorizedSendersDeactive(address account)
func (_Contract *ContractFilterer) WatchAuthorizedSendersDeactive(opts *bind.WatchOpts, sink chan<- *ContractAuthorizedSendersDeactive) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "AuthorizedSendersDeactive")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAuthorizedSendersDeactive)
				if err := _Contract.contract.UnpackLog(event, "AuthorizedSendersDeactive", log); err != nil {
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
func (_Contract *ContractFilterer) ParseAuthorizedSendersDeactive(log types.Log) (*ContractAuthorizedSendersDeactive, error) {
	event := new(ContractAuthorizedSendersDeactive)
	if err := _Contract.contract.UnpackLog(event, "AuthorizedSendersDeactive", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractConfigSetIterator is returned from FilterConfigSet and is used to iterate over the raw logs and unpacked data for ConfigSet events raised by the Contract contract.
type ContractConfigSetIterator struct {
	Event *ContractConfigSet // Event containing the contract specifics and raw log

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
func (it *ContractConfigSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractConfigSet)
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
		it.Event = new(ContractConfigSet)
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
func (it *ContractConfigSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractConfigSet represents a ConfigSet event raised by the Contract contract.
type ContractConfigSet struct {
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
func (_Contract *ContractFilterer) FilterConfigSet(opts *bind.FilterOpts) (*ContractConfigSetIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &ContractConfigSetIterator{contract: _Contract.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

// WatchConfigSet is a free log subscription operation binding the contract event 0x1591690b8638f5fb2dbec82ac741805ac5da8b45dc5263f4875b0496fdce4e05.
//
// Solidity: event ConfigSet(uint32 previousConfigBlockNumber, bytes32 configDigest, uint64 configCount, address[] signers, address[] transmitters, uint8 f, bytes onchainConfig, uint64 offchainConfigVersion, bytes offchainConfig)
func (_Contract *ContractFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *ContractConfigSet) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractConfigSet)
				if err := _Contract.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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
func (_Contract *ContractFilterer) ParseConfigSet(log types.Log) (*ContractConfigSet, error) {
	event := new(ContractConfigSet)
	if err := _Contract.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Contract contract.
type ContractInitializedIterator struct {
	Event *ContractInitialized // Event containing the contract specifics and raw log

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
func (it *ContractInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractInitialized)
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
		it.Event = new(ContractInitialized)
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
func (it *ContractInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractInitialized represents a Initialized event raised by the Contract contract.
type ContractInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Contract *ContractFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractInitializedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractInitializedIterator{contract: _Contract.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Contract *ContractFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractInitialized) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractInitialized)
				if err := _Contract.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Contract *ContractFilterer) ParseInitialized(log types.Log) (*ContractInitialized, error) {
	event := new(ContractInitialized)
	if err := _Contract.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOracleRequestIterator is returned from FilterOracleRequest and is used to iterate over the raw logs and unpacked data for OracleRequest events raised by the Contract contract.
type ContractOracleRequestIterator struct {
	Event *ContractOracleRequest // Event containing the contract specifics and raw log

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
func (it *ContractOracleRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOracleRequest)
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
		it.Event = new(ContractOracleRequest)
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
func (it *ContractOracleRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOracleRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOracleRequest represents a OracleRequest event raised by the Contract contract.
type ContractOracleRequest struct {
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
func (_Contract *ContractFilterer) FilterOracleRequest(opts *bind.FilterOpts, requestId [][32]byte) (*ContractOracleRequestIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "OracleRequest", requestIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractOracleRequestIterator{contract: _Contract.contract, event: "OracleRequest", logs: logs, sub: sub}, nil
}

// WatchOracleRequest is a free log subscription operation binding the contract event 0xa1ec73989d79578cd6f67d4f593ac3e0a4d1020e5c0164db52108d7ff785406c.
//
// Solidity: event OracleRequest(bytes32 indexed requestId, address requestingContract, address requestInitiator, uint64 subscriptionId, address subscriptionOwner, bytes data)
func (_Contract *ContractFilterer) WatchOracleRequest(opts *bind.WatchOpts, sink chan<- *ContractOracleRequest, requestId [][32]byte) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "OracleRequest", requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOracleRequest)
				if err := _Contract.contract.UnpackLog(event, "OracleRequest", log); err != nil {
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
func (_Contract *ContractFilterer) ParseOracleRequest(log types.Log) (*ContractOracleRequest, error) {
	event := new(ContractOracleRequest)
	if err := _Contract.contract.UnpackLog(event, "OracleRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOracleResponseIterator is returned from FilterOracleResponse and is used to iterate over the raw logs and unpacked data for OracleResponse events raised by the Contract contract.
type ContractOracleResponseIterator struct {
	Event *ContractOracleResponse // Event containing the contract specifics and raw log

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
func (it *ContractOracleResponseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOracleResponse)
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
		it.Event = new(ContractOracleResponse)
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
func (it *ContractOracleResponseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOracleResponseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOracleResponse represents a OracleResponse event raised by the Contract contract.
type ContractOracleResponse struct {
	RequestId [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOracleResponse is a free log retrieval operation binding the contract event 0x9e9bc7616d42c2835d05ae617e508454e63b30b934be8aa932ebc125e0e58a64.
//
// Solidity: event OracleResponse(bytes32 indexed requestId)
func (_Contract *ContractFilterer) FilterOracleResponse(opts *bind.FilterOpts, requestId [][32]byte) (*ContractOracleResponseIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "OracleResponse", requestIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractOracleResponseIterator{contract: _Contract.contract, event: "OracleResponse", logs: logs, sub: sub}, nil
}

// WatchOracleResponse is a free log subscription operation binding the contract event 0x9e9bc7616d42c2835d05ae617e508454e63b30b934be8aa932ebc125e0e58a64.
//
// Solidity: event OracleResponse(bytes32 indexed requestId)
func (_Contract *ContractFilterer) WatchOracleResponse(opts *bind.WatchOpts, sink chan<- *ContractOracleResponse, requestId [][32]byte) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "OracleResponse", requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOracleResponse)
				if err := _Contract.contract.UnpackLog(event, "OracleResponse", log); err != nil {
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
func (_Contract *ContractFilterer) ParseOracleResponse(log types.Log) (*ContractOracleResponse, error) {
	event := new(ContractOracleResponse)
	if err := _Contract.contract.UnpackLog(event, "OracleResponse", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOwnershipTransferRequestedIterator is returned from FilterOwnershipTransferRequested and is used to iterate over the raw logs and unpacked data for OwnershipTransferRequested events raised by the Contract contract.
type ContractOwnershipTransferRequestedIterator struct {
	Event *ContractOwnershipTransferRequested // Event containing the contract specifics and raw log

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
func (it *ContractOwnershipTransferRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOwnershipTransferRequested)
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
		it.Event = new(ContractOwnershipTransferRequested)
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
func (it *ContractOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOwnershipTransferRequested represents a OwnershipTransferRequested event raised by the Contract contract.
type ContractOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferRequested is a free log retrieval operation binding the contract event 0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278.
//
// Solidity: event OwnershipTransferRequested(address indexed from, address indexed to)
func (_Contract *ContractFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ContractOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ContractOwnershipTransferRequestedIterator{contract: _Contract.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferRequested is a free log subscription operation binding the contract event 0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278.
//
// Solidity: event OwnershipTransferRequested(address indexed from, address indexed to)
func (_Contract *ContractFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *ContractOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOwnershipTransferRequested)
				if err := _Contract.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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
func (_Contract *ContractFilterer) ParseOwnershipTransferRequested(log types.Log) (*ContractOwnershipTransferRequested, error) {
	event := new(ContractOwnershipTransferRequested)
	if err := _Contract.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Contract contract.
type ContractOwnershipTransferredIterator struct {
	Event *ContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOwnershipTransferred)
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
		it.Event = new(ContractOwnershipTransferred)
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
func (it *ContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOwnershipTransferred represents a OwnershipTransferred event raised by the Contract contract.
type ContractOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed from, address indexed to)
func (_Contract *ContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ContractOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ContractOwnershipTransferredIterator{contract: _Contract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed from, address indexed to)
func (_Contract *ContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOwnershipTransferred)
				if err := _Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Contract *ContractFilterer) ParseOwnershipTransferred(log types.Log) (*ContractOwnershipTransferred, error) {
	event := new(ContractOwnershipTransferred)
	if err := _Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTransmittedIterator is returned from FilterTransmitted and is used to iterate over the raw logs and unpacked data for Transmitted events raised by the Contract contract.
type ContractTransmittedIterator struct {
	Event *ContractTransmitted // Event containing the contract specifics and raw log

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
func (it *ContractTransmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTransmitted)
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
		it.Event = new(ContractTransmitted)
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
func (it *ContractTransmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTransmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTransmitted represents a Transmitted event raised by the Contract contract.
type ContractTransmitted struct {
	ConfigDigest [32]byte
	Epoch        uint32
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTransmitted is a free log retrieval operation binding the contract event 0xb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62.
//
// Solidity: event Transmitted(bytes32 configDigest, uint32 epoch)
func (_Contract *ContractFilterer) FilterTransmitted(opts *bind.FilterOpts) (*ContractTransmittedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Transmitted")
	if err != nil {
		return nil, err
	}
	return &ContractTransmittedIterator{contract: _Contract.contract, event: "Transmitted", logs: logs, sub: sub}, nil
}

// WatchTransmitted is a free log subscription operation binding the contract event 0xb04e63db38c49950639fa09d29872f21f5d49d614f3a969d8adf3d4b52e41a62.
//
// Solidity: event Transmitted(bytes32 configDigest, uint32 epoch)
func (_Contract *ContractFilterer) WatchTransmitted(opts *bind.WatchOpts, sink chan<- *ContractTransmitted) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Transmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTransmitted)
				if err := _Contract.contract.UnpackLog(event, "Transmitted", log); err != nil {
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
func (_Contract *ContractFilterer) ParseTransmitted(log types.Log) (*ContractTransmitted, error) {
	event := new(ContractTransmitted)
	if err := _Contract.contract.UnpackLog(event, "Transmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractUserCallbackErrorIterator is returned from FilterUserCallbackError and is used to iterate over the raw logs and unpacked data for UserCallbackError events raised by the Contract contract.
type ContractUserCallbackErrorIterator struct {
	Event *ContractUserCallbackError // Event containing the contract specifics and raw log

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
func (it *ContractUserCallbackErrorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractUserCallbackError)
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
		it.Event = new(ContractUserCallbackError)
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
func (it *ContractUserCallbackErrorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractUserCallbackErrorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractUserCallbackError represents a UserCallbackError event raised by the Contract contract.
type ContractUserCallbackError struct {
	RequestId [32]byte
	Reason    string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUserCallbackError is a free log retrieval operation binding the contract event 0xb2931868c372fe17a25643458add467d60ec5c51125a99b7309f41f5bcd2da6c.
//
// Solidity: event UserCallbackError(bytes32 indexed requestId, string reason)
func (_Contract *ContractFilterer) FilterUserCallbackError(opts *bind.FilterOpts, requestId [][32]byte) (*ContractUserCallbackErrorIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "UserCallbackError", requestIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractUserCallbackErrorIterator{contract: _Contract.contract, event: "UserCallbackError", logs: logs, sub: sub}, nil
}

// WatchUserCallbackError is a free log subscription operation binding the contract event 0xb2931868c372fe17a25643458add467d60ec5c51125a99b7309f41f5bcd2da6c.
//
// Solidity: event UserCallbackError(bytes32 indexed requestId, string reason)
func (_Contract *ContractFilterer) WatchUserCallbackError(opts *bind.WatchOpts, sink chan<- *ContractUserCallbackError, requestId [][32]byte) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "UserCallbackError", requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractUserCallbackError)
				if err := _Contract.contract.UnpackLog(event, "UserCallbackError", log); err != nil {
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
func (_Contract *ContractFilterer) ParseUserCallbackError(log types.Log) (*ContractUserCallbackError, error) {
	event := new(ContractUserCallbackError)
	if err := _Contract.contract.UnpackLog(event, "UserCallbackError", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractUserCallbackRawErrorIterator is returned from FilterUserCallbackRawError and is used to iterate over the raw logs and unpacked data for UserCallbackRawError events raised by the Contract contract.
type ContractUserCallbackRawErrorIterator struct {
	Event *ContractUserCallbackRawError // Event containing the contract specifics and raw log

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
func (it *ContractUserCallbackRawErrorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractUserCallbackRawError)
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
		it.Event = new(ContractUserCallbackRawError)
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
func (it *ContractUserCallbackRawErrorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractUserCallbackRawErrorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractUserCallbackRawError represents a UserCallbackRawError event raised by the Contract contract.
type ContractUserCallbackRawError struct {
	RequestId    [32]byte
	LowLevelData []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUserCallbackRawError is a free log retrieval operation binding the contract event 0xe0b838ffe6ee22a0d3acf19a85db6a41b34a1ab739e2d6c759a2e42d95bdccb2.
//
// Solidity: event UserCallbackRawError(bytes32 indexed requestId, bytes lowLevelData)
func (_Contract *ContractFilterer) FilterUserCallbackRawError(opts *bind.FilterOpts, requestId [][32]byte) (*ContractUserCallbackRawErrorIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "UserCallbackRawError", requestIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractUserCallbackRawErrorIterator{contract: _Contract.contract, event: "UserCallbackRawError", logs: logs, sub: sub}, nil
}

// WatchUserCallbackRawError is a free log subscription operation binding the contract event 0xe0b838ffe6ee22a0d3acf19a85db6a41b34a1ab739e2d6c759a2e42d95bdccb2.
//
// Solidity: event UserCallbackRawError(bytes32 indexed requestId, bytes lowLevelData)
func (_Contract *ContractFilterer) WatchUserCallbackRawError(opts *bind.WatchOpts, sink chan<- *ContractUserCallbackRawError, requestId [][32]byte) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "UserCallbackRawError", requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractUserCallbackRawError)
				if err := _Contract.contract.UnpackLog(event, "UserCallbackRawError", log); err != nil {
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
func (_Contract *ContractFilterer) ParseUserCallbackRawError(log types.Log) (*ContractUserCallbackRawError, error) {
	event := new(ContractUserCallbackRawError)
	if err := _Contract.contract.UnpackLog(event, "UserCallbackRawError", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
