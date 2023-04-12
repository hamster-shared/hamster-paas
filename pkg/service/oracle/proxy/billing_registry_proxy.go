// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package oracle_proxy

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

// FunctionsBillingRegistryCommitment is an auto generated low-level Go binding around an user-defined struct.
type FunctionsBillingRegistryCommitment struct {
	SubscriptionId uint64
	Client         common.Address
	GasLimit       uint32
	GasPrice       *big.Int
	Don            common.Address
	DonFee         *big.Int
	RegistryFee    *big.Int
	EstimatedCost  *big.Int
	Timestamp      *big.Int
}

// FunctionsBillingRegistryInterfaceRequestBilling is an auto generated low-level Go binding around an user-defined struct.
type FunctionsBillingRegistryInterfaceRequestBilling struct {
	SubscriptionId uint64
	Client         common.Address
	GasLimit       uint32
	GasPrice       *big.Int
}

// OracleProxyMetaData contains all meta data concerning the OracleProxy contract.
var OracleProxyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"internalBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"externalBalance\",\"type\":\"uint256\"}],\"name\":\"BalanceInvariantViolated\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotSelfTransfer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptySendersList\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"have\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"want\",\"type\":\"uint32\"}],\"name\":\"GasLimitTooBig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectRequestID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCalldata\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"subscriptionId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"consumer\",\"type\":\"address\"}],\"name\":\"InvalidConsumer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"linkWei\",\"type\":\"int256\"}],\"name\":\"InvalidLinkWeiPrice\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSubscription\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proposedOwner\",\"type\":\"address\"}],\"name\":\"MustBeRequestedOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"MustBeSubOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAllowedToSetSenders\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotProposedOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableFromLink\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerMustBeSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PaymentTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingRequestExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Reentrant\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooManyConsumers\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnauthorizedSender\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"senders\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"changedBy\",\"type\":\"address\"}],\"name\":\"AuthorizedSendersChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"subscriptionId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"signerPayment\",\"type\":\"uint96\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"transmitterPayment\",\"type\":\"uint96\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"totalCost\",\"type\":\"uint96\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"name\":\"BillingEnd\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"subscriptionId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"client\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"gasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"don\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"donFee\",\"type\":\"uint96\"},{\"internalType\":\"uint96\",\"name\":\"registryFee\",\"type\":\"uint96\"},{\"internalType\":\"uint96\",\"name\":\"estimatedCost\",\"type\":\"uint96\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structFunctionsBillingRegistry.Commitment\",\"name\":\"commitment\",\"type\":\"tuple\"}],\"name\":\"BillingStart\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"maxGasLimit\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"stalenessSeconds\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasAfterPaymentCalculation\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"fallbackWeiPerUnitLink\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"gasOverhead\",\"type\":\"uint32\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FundsRecovered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"}],\"name\":\"RequestTimedOut\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"subscriptionId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SubscriptionCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"subscriptionId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"consumer\",\"type\":\"address\"}],\"name\":\"SubscriptionConsumerAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"subscriptionId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"consumer\",\"type\":\"address\"}],\"name\":\"SubscriptionConsumerRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"subscriptionId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"SubscriptionCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"subscriptionId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBalance\",\"type\":\"uint256\"}],\"name\":\"SubscriptionFunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"subscriptionId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"SubscriptionOwnerTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"subscriptionId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"SubscriptionOwnerTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_CONSUMERS\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"subscriptionId\",\"type\":\"uint64\"}],\"name\":\"acceptSubscriptionOwnerTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"subscriptionId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"consumer\",\"type\":\"address\"}],\"name\":\"addConsumer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"subscriptionId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"cancelSubscription\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"createSubscription\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"gasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint96\",\"name\":\"donFee\",\"type\":\"uint96\"},{\"internalType\":\"uint96\",\"name\":\"registryFee\",\"type\":\"uint96\"}],\"name\":\"estimateCost\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"response\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"err\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"internalType\":\"address[31]\",\"name\":\"signers\",\"type\":\"address[31]\"},{\"internalType\":\"uint8\",\"name\":\"signerCount\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"reportValidationGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initialGas\",\"type\":\"uint256\"}],\"name\":\"fulfillAndBill\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAuthorizedSenders\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConfig\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"maxGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"stalenessSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"gasAfterPaymentCalculation\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"fallbackWeiPerUnitLink\",\"type\":\"int256\"},{\"internalType\":\"uint32\",\"name\":\"gasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"linkAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"linkPriceFeed\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentsubscriptionId\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRequestConfig\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"subscriptionId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"client\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"gasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"}],\"internalType\":\"structFunctionsBillingRegistryInterface.RequestBilling\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"getRequiredFee\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"subscriptionId\",\"type\":\"uint64\"}],\"name\":\"getSubscription\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"balance\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"consumers\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"subscriptionId\",\"type\":\"uint64\"}],\"name\":\"getSubscriptionOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"link\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"linkEthFeed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"isAuthorizedSender\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onTokenTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"}],\"name\":\"oracleWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"subscriptionId\",\"type\":\"uint64\"}],\"name\":\"ownerCancelSubscription\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"subscriptionId\",\"type\":\"uint64\"}],\"name\":\"pendingRequestExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"recoverFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"subscriptionId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"consumer\",\"type\":\"address\"}],\"name\":\"removeConsumer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"subscriptionId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"requestSubscriptionOwnerTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"senders\",\"type\":\"address[]\"}],\"name\":\"setAuthorizedSenders\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"maxGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"stalenessSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"gasAfterPaymentCalculation\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"fallbackWeiPerUnitLink\",\"type\":\"int256\"},{\"internalType\":\"uint32\",\"name\":\"gasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"requestTimeoutSeconds\",\"type\":\"uint32\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"subscriptionId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"client\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"gasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"}],\"internalType\":\"structFunctionsBillingRegistryInterface.RequestBilling\",\"name\":\"billing\",\"type\":\"tuple\"}],\"name\":\"startBilling\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"requestIdsToTimeout\",\"type\":\"bytes32[]\"}],\"name\":\"timeoutRequests\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// OracleProxyABI is the input ABI used to generate the binding from.
// Deprecated: Use OracleProxyMetaData.ABI instead.
var OracleProxyABI = OracleProxyMetaData.ABI

// OracleProxy is an auto generated Go binding around an Ethereum contract.
type OracleProxy struct {
	OracleProxyCaller     // Read-only binding to the contract
	OracleProxyTransactor // Write-only binding to the contract
	OracleProxyFilterer   // Log filterer for contract events
}

// OracleProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type OracleProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OracleProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OracleProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OracleProxySession struct {
	Contract     *OracleProxy      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OracleProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OracleProxyCallerSession struct {
	Contract *OracleProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// OracleProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OracleProxyTransactorSession struct {
	Contract     *OracleProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// OracleProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type OracleProxyRaw struct {
	Contract *OracleProxy // Generic contract binding to access the raw methods on
}

// OracleProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OracleProxyCallerRaw struct {
	Contract *OracleProxyCaller // Generic read-only contract binding to access the raw methods on
}

// OracleProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OracleProxyTransactorRaw struct {
	Contract *OracleProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOracleProxy creates a new instance of OracleProxy, bound to a specific deployed contract.
func NewOracleProxy(address common.Address, backend bind.ContractBackend) (*OracleProxy, error) {
	contract, err := bindOracleProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OracleProxy{OracleProxyCaller: OracleProxyCaller{contract: contract}, OracleProxyTransactor: OracleProxyTransactor{contract: contract}, OracleProxyFilterer: OracleProxyFilterer{contract: contract}}, nil
}

// NewOracleProxyCaller creates a new read-only instance of OracleProxy, bound to a specific deployed contract.
func NewOracleProxyCaller(address common.Address, caller bind.ContractCaller) (*OracleProxyCaller, error) {
	contract, err := bindOracleProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OracleProxyCaller{contract: contract}, nil
}

// NewOracleProxyTransactor creates a new write-only instance of OracleProxy, bound to a specific deployed contract.
func NewOracleProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*OracleProxyTransactor, error) {
	contract, err := bindOracleProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OracleProxyTransactor{contract: contract}, nil
}

// NewOracleProxyFilterer creates a new log filterer instance of OracleProxy, bound to a specific deployed contract.
func NewOracleProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*OracleProxyFilterer, error) {
	contract, err := bindOracleProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OracleProxyFilterer{contract: contract}, nil
}

// bindOracleProxy binds a generic wrapper to an already deployed contract.
func bindOracleProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OracleProxyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OracleProxy *OracleProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OracleProxy.Contract.OracleProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OracleProxy *OracleProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleProxy.Contract.OracleProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OracleProxy *OracleProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OracleProxy.Contract.OracleProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OracleProxy *OracleProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OracleProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OracleProxy *OracleProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OracleProxy *OracleProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OracleProxy.Contract.contract.Transact(opts, method, params...)
}

// MAXCONSUMERS is a free data retrieval call binding the contract method 0x64d51a2a.
//
// Solidity: function MAX_CONSUMERS() view returns(uint16)
func (_OracleProxy *OracleProxyCaller) MAXCONSUMERS(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _OracleProxy.contract.Call(opts, &out, "MAX_CONSUMERS")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MAXCONSUMERS is a free data retrieval call binding the contract method 0x64d51a2a.
//
// Solidity: function MAX_CONSUMERS() view returns(uint16)
func (_OracleProxy *OracleProxySession) MAXCONSUMERS() (uint16, error) {
	return _OracleProxy.Contract.MAXCONSUMERS(&_OracleProxy.CallOpts)
}

// MAXCONSUMERS is a free data retrieval call binding the contract method 0x64d51a2a.
//
// Solidity: function MAX_CONSUMERS() view returns(uint16)
func (_OracleProxy *OracleProxyCallerSession) MAXCONSUMERS() (uint16, error) {
	return _OracleProxy.Contract.MAXCONSUMERS(&_OracleProxy.CallOpts)
}

// EstimateCost is a free data retrieval call binding the contract method 0xa1a6d041.
//
// Solidity: function estimateCost(uint32 gasLimit, uint256 gasPrice, uint96 donFee, uint96 registryFee) view returns(uint96)
func (_OracleProxy *OracleProxyCaller) EstimateCost(opts *bind.CallOpts, gasLimit uint32, gasPrice *big.Int, donFee *big.Int, registryFee *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _OracleProxy.contract.Call(opts, &out, "estimateCost", gasLimit, gasPrice, donFee, registryFee)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimateCost is a free data retrieval call binding the contract method 0xa1a6d041.
//
// Solidity: function estimateCost(uint32 gasLimit, uint256 gasPrice, uint96 donFee, uint96 registryFee) view returns(uint96)
func (_OracleProxy *OracleProxySession) EstimateCost(gasLimit uint32, gasPrice *big.Int, donFee *big.Int, registryFee *big.Int) (*big.Int, error) {
	return _OracleProxy.Contract.EstimateCost(&_OracleProxy.CallOpts, gasLimit, gasPrice, donFee, registryFee)
}

// EstimateCost is a free data retrieval call binding the contract method 0xa1a6d041.
//
// Solidity: function estimateCost(uint32 gasLimit, uint256 gasPrice, uint96 donFee, uint96 registryFee) view returns(uint96)
func (_OracleProxy *OracleProxyCallerSession) EstimateCost(gasLimit uint32, gasPrice *big.Int, donFee *big.Int, registryFee *big.Int) (*big.Int, error) {
	return _OracleProxy.Contract.EstimateCost(&_OracleProxy.CallOpts, gasLimit, gasPrice, donFee, registryFee)
}

// GetAuthorizedSenders is a free data retrieval call binding the contract method 0x2408afaa.
//
// Solidity: function getAuthorizedSenders() view returns(address[])
func (_OracleProxy *OracleProxyCaller) GetAuthorizedSenders(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _OracleProxy.contract.Call(opts, &out, "getAuthorizedSenders")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAuthorizedSenders is a free data retrieval call binding the contract method 0x2408afaa.
//
// Solidity: function getAuthorizedSenders() view returns(address[])
func (_OracleProxy *OracleProxySession) GetAuthorizedSenders() ([]common.Address, error) {
	return _OracleProxy.Contract.GetAuthorizedSenders(&_OracleProxy.CallOpts)
}

// GetAuthorizedSenders is a free data retrieval call binding the contract method 0x2408afaa.
//
// Solidity: function getAuthorizedSenders() view returns(address[])
func (_OracleProxy *OracleProxyCallerSession) GetAuthorizedSenders() ([]common.Address, error) {
	return _OracleProxy.Contract.GetAuthorizedSenders(&_OracleProxy.CallOpts)
}

// GetConfig is a free data retrieval call binding the contract method 0xc3f909d4.
//
// Solidity: function getConfig() view returns(uint32 maxGasLimit, uint32 stalenessSeconds, uint256 gasAfterPaymentCalculation, int256 fallbackWeiPerUnitLink, uint32 gasOverhead, address linkAddress, address linkPriceFeed)
func (_OracleProxy *OracleProxyCaller) GetConfig(opts *bind.CallOpts) (struct {
	MaxGasLimit                uint32
	StalenessSeconds           uint32
	GasAfterPaymentCalculation *big.Int
	FallbackWeiPerUnitLink     *big.Int
	GasOverhead                uint32
	LinkAddress                common.Address
	LinkPriceFeed              common.Address
}, error) {
	var out []interface{}
	err := _OracleProxy.contract.Call(opts, &out, "getConfig")

	outstruct := new(struct {
		MaxGasLimit                uint32
		StalenessSeconds           uint32
		GasAfterPaymentCalculation *big.Int
		FallbackWeiPerUnitLink     *big.Int
		GasOverhead                uint32
		LinkAddress                common.Address
		LinkPriceFeed              common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MaxGasLimit = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.StalenessSeconds = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.GasAfterPaymentCalculation = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.FallbackWeiPerUnitLink = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.GasOverhead = *abi.ConvertType(out[4], new(uint32)).(*uint32)
	outstruct.LinkAddress = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	outstruct.LinkPriceFeed = *abi.ConvertType(out[6], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// GetConfig is a free data retrieval call binding the contract method 0xc3f909d4.
//
// Solidity: function getConfig() view returns(uint32 maxGasLimit, uint32 stalenessSeconds, uint256 gasAfterPaymentCalculation, int256 fallbackWeiPerUnitLink, uint32 gasOverhead, address linkAddress, address linkPriceFeed)
func (_OracleProxy *OracleProxySession) GetConfig() (struct {
	MaxGasLimit                uint32
	StalenessSeconds           uint32
	GasAfterPaymentCalculation *big.Int
	FallbackWeiPerUnitLink     *big.Int
	GasOverhead                uint32
	LinkAddress                common.Address
	LinkPriceFeed              common.Address
}, error) {
	return _OracleProxy.Contract.GetConfig(&_OracleProxy.CallOpts)
}

// GetConfig is a free data retrieval call binding the contract method 0xc3f909d4.
//
// Solidity: function getConfig() view returns(uint32 maxGasLimit, uint32 stalenessSeconds, uint256 gasAfterPaymentCalculation, int256 fallbackWeiPerUnitLink, uint32 gasOverhead, address linkAddress, address linkPriceFeed)
func (_OracleProxy *OracleProxyCallerSession) GetConfig() (struct {
	MaxGasLimit                uint32
	StalenessSeconds           uint32
	GasAfterPaymentCalculation *big.Int
	FallbackWeiPerUnitLink     *big.Int
	GasOverhead                uint32
	LinkAddress                common.Address
	LinkPriceFeed              common.Address
}, error) {
	return _OracleProxy.Contract.GetConfig(&_OracleProxy.CallOpts)
}

// GetCurrentsubscriptionId is a free data retrieval call binding the contract method 0x33652e3e.
//
// Solidity: function getCurrentsubscriptionId() view returns(uint64)
func (_OracleProxy *OracleProxyCaller) GetCurrentsubscriptionId(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _OracleProxy.contract.Call(opts, &out, "getCurrentsubscriptionId")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetCurrentsubscriptionId is a free data retrieval call binding the contract method 0x33652e3e.
//
// Solidity: function getCurrentsubscriptionId() view returns(uint64)
func (_OracleProxy *OracleProxySession) GetCurrentsubscriptionId() (uint64, error) {
	return _OracleProxy.Contract.GetCurrentsubscriptionId(&_OracleProxy.CallOpts)
}

// GetCurrentsubscriptionId is a free data retrieval call binding the contract method 0x33652e3e.
//
// Solidity: function getCurrentsubscriptionId() view returns(uint64)
func (_OracleProxy *OracleProxyCallerSession) GetCurrentsubscriptionId() (uint64, error) {
	return _OracleProxy.Contract.GetCurrentsubscriptionId(&_OracleProxy.CallOpts)
}

// GetRequestConfig is a free data retrieval call binding the contract method 0x00012291.
//
// Solidity: function getRequestConfig() view returns(uint32, address[])
func (_OracleProxy *OracleProxyCaller) GetRequestConfig(opts *bind.CallOpts) (uint32, []common.Address, error) {
	var out []interface{}
	err := _OracleProxy.contract.Call(opts, &out, "getRequestConfig")

	if err != nil {
		return *new(uint32), *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)
	out1 := *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)

	return out0, out1, err

}

// GetRequestConfig is a free data retrieval call binding the contract method 0x00012291.
//
// Solidity: function getRequestConfig() view returns(uint32, address[])
func (_OracleProxy *OracleProxySession) GetRequestConfig() (uint32, []common.Address, error) {
	return _OracleProxy.Contract.GetRequestConfig(&_OracleProxy.CallOpts)
}

// GetRequestConfig is a free data retrieval call binding the contract method 0x00012291.
//
// Solidity: function getRequestConfig() view returns(uint32, address[])
func (_OracleProxy *OracleProxyCallerSession) GetRequestConfig() (uint32, []common.Address, error) {
	return _OracleProxy.Contract.GetRequestConfig(&_OracleProxy.CallOpts)
}

// GetRequiredFee is a free data retrieval call binding the contract method 0xf1e14a21.
//
// Solidity: function getRequiredFee(bytes , (uint64,address,uint32,uint256) ) pure returns(uint96)
func (_OracleProxy *OracleProxyCaller) GetRequiredFee(opts *bind.CallOpts, arg0 []byte, arg1 FunctionsBillingRegistryInterfaceRequestBilling) (*big.Int, error) {
	var out []interface{}
	err := _OracleProxy.contract.Call(opts, &out, "getRequiredFee", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRequiredFee is a free data retrieval call binding the contract method 0xf1e14a21.
//
// Solidity: function getRequiredFee(bytes , (uint64,address,uint32,uint256) ) pure returns(uint96)
func (_OracleProxy *OracleProxySession) GetRequiredFee(arg0 []byte, arg1 FunctionsBillingRegistryInterfaceRequestBilling) (*big.Int, error) {
	return _OracleProxy.Contract.GetRequiredFee(&_OracleProxy.CallOpts, arg0, arg1)
}

// GetRequiredFee is a free data retrieval call binding the contract method 0xf1e14a21.
//
// Solidity: function getRequiredFee(bytes , (uint64,address,uint32,uint256) ) pure returns(uint96)
func (_OracleProxy *OracleProxyCallerSession) GetRequiredFee(arg0 []byte, arg1 FunctionsBillingRegistryInterfaceRequestBilling) (*big.Int, error) {
	return _OracleProxy.Contract.GetRequiredFee(&_OracleProxy.CallOpts, arg0, arg1)
}

// GetSubscription is a free data retrieval call binding the contract method 0xa47c7696.
//
// Solidity: function getSubscription(uint64 subscriptionId) view returns(uint96 balance, address owner, address[] consumers)
func (_OracleProxy *OracleProxyCaller) GetSubscription(opts *bind.CallOpts, subscriptionId uint64) (struct {
	Balance   *big.Int
	Owner     common.Address
	Consumers []common.Address
}, error) {
	var out []interface{}
	err := _OracleProxy.contract.Call(opts, &out, "getSubscription", subscriptionId)

	outstruct := new(struct {
		Balance   *big.Int
		Owner     common.Address
		Consumers []common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Balance = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Owner = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Consumers = *abi.ConvertType(out[2], new([]common.Address)).(*[]common.Address)

	return *outstruct, err

}

// GetSubscription is a free data retrieval call binding the contract method 0xa47c7696.
//
// Solidity: function getSubscription(uint64 subscriptionId) view returns(uint96 balance, address owner, address[] consumers)
func (_OracleProxy *OracleProxySession) GetSubscription(subscriptionId uint64) (struct {
	Balance   *big.Int
	Owner     common.Address
	Consumers []common.Address
}, error) {
	return _OracleProxy.Contract.GetSubscription(&_OracleProxy.CallOpts, subscriptionId)
}

// GetSubscription is a free data retrieval call binding the contract method 0xa47c7696.
//
// Solidity: function getSubscription(uint64 subscriptionId) view returns(uint96 balance, address owner, address[] consumers)
func (_OracleProxy *OracleProxyCallerSession) GetSubscription(subscriptionId uint64) (struct {
	Balance   *big.Int
	Owner     common.Address
	Consumers []common.Address
}, error) {
	return _OracleProxy.Contract.GetSubscription(&_OracleProxy.CallOpts, subscriptionId)
}

// GetSubscriptionOwner is a free data retrieval call binding the contract method 0xb2a489ff.
//
// Solidity: function getSubscriptionOwner(uint64 subscriptionId) view returns(address owner)
func (_OracleProxy *OracleProxyCaller) GetSubscriptionOwner(opts *bind.CallOpts, subscriptionId uint64) (common.Address, error) {
	var out []interface{}
	err := _OracleProxy.contract.Call(opts, &out, "getSubscriptionOwner", subscriptionId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSubscriptionOwner is a free data retrieval call binding the contract method 0xb2a489ff.
//
// Solidity: function getSubscriptionOwner(uint64 subscriptionId) view returns(address owner)
func (_OracleProxy *OracleProxySession) GetSubscriptionOwner(subscriptionId uint64) (common.Address, error) {
	return _OracleProxy.Contract.GetSubscriptionOwner(&_OracleProxy.CallOpts, subscriptionId)
}

// GetSubscriptionOwner is a free data retrieval call binding the contract method 0xb2a489ff.
//
// Solidity: function getSubscriptionOwner(uint64 subscriptionId) view returns(address owner)
func (_OracleProxy *OracleProxyCallerSession) GetSubscriptionOwner(subscriptionId uint64) (common.Address, error) {
	return _OracleProxy.Contract.GetSubscriptionOwner(&_OracleProxy.CallOpts, subscriptionId)
}

// GetTotalBalance is a free data retrieval call binding the contract method 0x12b58349.
//
// Solidity: function getTotalBalance() view returns(uint256)
func (_OracleProxy *OracleProxyCaller) GetTotalBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OracleProxy.contract.Call(opts, &out, "getTotalBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalBalance is a free data retrieval call binding the contract method 0x12b58349.
//
// Solidity: function getTotalBalance() view returns(uint256)
func (_OracleProxy *OracleProxySession) GetTotalBalance() (*big.Int, error) {
	return _OracleProxy.Contract.GetTotalBalance(&_OracleProxy.CallOpts)
}

// GetTotalBalance is a free data retrieval call binding the contract method 0x12b58349.
//
// Solidity: function getTotalBalance() view returns(uint256)
func (_OracleProxy *OracleProxyCallerSession) GetTotalBalance() (*big.Int, error) {
	return _OracleProxy.Contract.GetTotalBalance(&_OracleProxy.CallOpts)
}

// IsAuthorizedSender is a free data retrieval call binding the contract method 0xfa00763a.
//
// Solidity: function isAuthorizedSender(address sender) view returns(bool)
func (_OracleProxy *OracleProxyCaller) IsAuthorizedSender(opts *bind.CallOpts, sender common.Address) (bool, error) {
	var out []interface{}
	err := _OracleProxy.contract.Call(opts, &out, "isAuthorizedSender", sender)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAuthorizedSender is a free data retrieval call binding the contract method 0xfa00763a.
//
// Solidity: function isAuthorizedSender(address sender) view returns(bool)
func (_OracleProxy *OracleProxySession) IsAuthorizedSender(sender common.Address) (bool, error) {
	return _OracleProxy.Contract.IsAuthorizedSender(&_OracleProxy.CallOpts, sender)
}

// IsAuthorizedSender is a free data retrieval call binding the contract method 0xfa00763a.
//
// Solidity: function isAuthorizedSender(address sender) view returns(bool)
func (_OracleProxy *OracleProxyCallerSession) IsAuthorizedSender(sender common.Address) (bool, error) {
	return _OracleProxy.Contract.IsAuthorizedSender(&_OracleProxy.CallOpts, sender)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OracleProxy *OracleProxyCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OracleProxy.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OracleProxy *OracleProxySession) Owner() (common.Address, error) {
	return _OracleProxy.Contract.Owner(&_OracleProxy.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OracleProxy *OracleProxyCallerSession) Owner() (common.Address, error) {
	return _OracleProxy.Contract.Owner(&_OracleProxy.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_OracleProxy *OracleProxyCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _OracleProxy.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_OracleProxy *OracleProxySession) Paused() (bool, error) {
	return _OracleProxy.Contract.Paused(&_OracleProxy.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_OracleProxy *OracleProxyCallerSession) Paused() (bool, error) {
	return _OracleProxy.Contract.Paused(&_OracleProxy.CallOpts)
}

// PendingRequestExists is a free data retrieval call binding the contract method 0xe82ad7d4.
//
// Solidity: function pendingRequestExists(uint64 subscriptionId) view returns(bool)
func (_OracleProxy *OracleProxyCaller) PendingRequestExists(opts *bind.CallOpts, subscriptionId uint64) (bool, error) {
	var out []interface{}
	err := _OracleProxy.contract.Call(opts, &out, "pendingRequestExists", subscriptionId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PendingRequestExists is a free data retrieval call binding the contract method 0xe82ad7d4.
//
// Solidity: function pendingRequestExists(uint64 subscriptionId) view returns(bool)
func (_OracleProxy *OracleProxySession) PendingRequestExists(subscriptionId uint64) (bool, error) {
	return _OracleProxy.Contract.PendingRequestExists(&_OracleProxy.CallOpts, subscriptionId)
}

// PendingRequestExists is a free data retrieval call binding the contract method 0xe82ad7d4.
//
// Solidity: function pendingRequestExists(uint64 subscriptionId) view returns(bool)
func (_OracleProxy *OracleProxyCallerSession) PendingRequestExists(subscriptionId uint64) (bool, error) {
	return _OracleProxy.Contract.PendingRequestExists(&_OracleProxy.CallOpts, subscriptionId)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_OracleProxy *OracleProxyTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleProxy.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_OracleProxy *OracleProxySession) AcceptOwnership() (*types.Transaction, error) {
	return _OracleProxy.Contract.AcceptOwnership(&_OracleProxy.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_OracleProxy *OracleProxyTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _OracleProxy.Contract.AcceptOwnership(&_OracleProxy.TransactOpts)
}

// AcceptSubscriptionOwnerTransfer is a paid mutator transaction binding the contract method 0x82359740.
//
// Solidity: function acceptSubscriptionOwnerTransfer(uint64 subscriptionId) returns()
func (_OracleProxy *OracleProxyTransactor) AcceptSubscriptionOwnerTransfer(opts *bind.TransactOpts, subscriptionId uint64) (*types.Transaction, error) {
	return _OracleProxy.contract.Transact(opts, "acceptSubscriptionOwnerTransfer", subscriptionId)
}

// AcceptSubscriptionOwnerTransfer is a paid mutator transaction binding the contract method 0x82359740.
//
// Solidity: function acceptSubscriptionOwnerTransfer(uint64 subscriptionId) returns()
func (_OracleProxy *OracleProxySession) AcceptSubscriptionOwnerTransfer(subscriptionId uint64) (*types.Transaction, error) {
	return _OracleProxy.Contract.AcceptSubscriptionOwnerTransfer(&_OracleProxy.TransactOpts, subscriptionId)
}

// AcceptSubscriptionOwnerTransfer is a paid mutator transaction binding the contract method 0x82359740.
//
// Solidity: function acceptSubscriptionOwnerTransfer(uint64 subscriptionId) returns()
func (_OracleProxy *OracleProxyTransactorSession) AcceptSubscriptionOwnerTransfer(subscriptionId uint64) (*types.Transaction, error) {
	return _OracleProxy.Contract.AcceptSubscriptionOwnerTransfer(&_OracleProxy.TransactOpts, subscriptionId)
}

// AddConsumer is a paid mutator transaction binding the contract method 0x7341c10c.
//
// Solidity: function addConsumer(uint64 subscriptionId, address consumer) returns()
func (_OracleProxy *OracleProxyTransactor) AddConsumer(opts *bind.TransactOpts, subscriptionId uint64, consumer common.Address) (*types.Transaction, error) {
	return _OracleProxy.contract.Transact(opts, "addConsumer", subscriptionId, consumer)
}

// AddConsumer is a paid mutator transaction binding the contract method 0x7341c10c.
//
// Solidity: function addConsumer(uint64 subscriptionId, address consumer) returns()
func (_OracleProxy *OracleProxySession) AddConsumer(subscriptionId uint64, consumer common.Address) (*types.Transaction, error) {
	return _OracleProxy.Contract.AddConsumer(&_OracleProxy.TransactOpts, subscriptionId, consumer)
}

// AddConsumer is a paid mutator transaction binding the contract method 0x7341c10c.
//
// Solidity: function addConsumer(uint64 subscriptionId, address consumer) returns()
func (_OracleProxy *OracleProxyTransactorSession) AddConsumer(subscriptionId uint64, consumer common.Address) (*types.Transaction, error) {
	return _OracleProxy.Contract.AddConsumer(&_OracleProxy.TransactOpts, subscriptionId, consumer)
}

// CancelSubscription is a paid mutator transaction binding the contract method 0xd7ae1d30.
//
// Solidity: function cancelSubscription(uint64 subscriptionId, address to) returns()
func (_OracleProxy *OracleProxyTransactor) CancelSubscription(opts *bind.TransactOpts, subscriptionId uint64, to common.Address) (*types.Transaction, error) {
	return _OracleProxy.contract.Transact(opts, "cancelSubscription", subscriptionId, to)
}

// CancelSubscription is a paid mutator transaction binding the contract method 0xd7ae1d30.
//
// Solidity: function cancelSubscription(uint64 subscriptionId, address to) returns()
func (_OracleProxy *OracleProxySession) CancelSubscription(subscriptionId uint64, to common.Address) (*types.Transaction, error) {
	return _OracleProxy.Contract.CancelSubscription(&_OracleProxy.TransactOpts, subscriptionId, to)
}

// CancelSubscription is a paid mutator transaction binding the contract method 0xd7ae1d30.
//
// Solidity: function cancelSubscription(uint64 subscriptionId, address to) returns()
func (_OracleProxy *OracleProxyTransactorSession) CancelSubscription(subscriptionId uint64, to common.Address) (*types.Transaction, error) {
	return _OracleProxy.Contract.CancelSubscription(&_OracleProxy.TransactOpts, subscriptionId, to)
}

// CreateSubscription is a paid mutator transaction binding the contract method 0xa21a23e4.
//
// Solidity: function createSubscription() returns(uint64)
func (_OracleProxy *OracleProxyTransactor) CreateSubscription(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleProxy.contract.Transact(opts, "createSubscription")
}

// CreateSubscription is a paid mutator transaction binding the contract method 0xa21a23e4.
//
// Solidity: function createSubscription() returns(uint64)
func (_OracleProxy *OracleProxySession) CreateSubscription() (*types.Transaction, error) {
	return _OracleProxy.Contract.CreateSubscription(&_OracleProxy.TransactOpts)
}

// CreateSubscription is a paid mutator transaction binding the contract method 0xa21a23e4.
//
// Solidity: function createSubscription() returns(uint64)
func (_OracleProxy *OracleProxyTransactorSession) CreateSubscription() (*types.Transaction, error) {
	return _OracleProxy.Contract.CreateSubscription(&_OracleProxy.TransactOpts)
}

// FulfillAndBill is a paid mutator transaction binding the contract method 0x0739e4f1.
//
// Solidity: function fulfillAndBill(bytes32 requestId, bytes response, bytes err, address transmitter, address[31] signers, uint8 signerCount, uint256 reportValidationGas, uint256 initialGas) returns(bool success)
func (_OracleProxy *OracleProxyTransactor) FulfillAndBill(opts *bind.TransactOpts, requestId [32]byte, response []byte, err []byte, transmitter common.Address, signers [31]common.Address, signerCount uint8, reportValidationGas *big.Int, initialGas *big.Int) (*types.Transaction, error) {
	return _OracleProxy.contract.Transact(opts, "fulfillAndBill", requestId, response, err, transmitter, signers, signerCount, reportValidationGas, initialGas)
}

// FulfillAndBill is a paid mutator transaction binding the contract method 0x0739e4f1.
//
// Solidity: function fulfillAndBill(bytes32 requestId, bytes response, bytes err, address transmitter, address[31] signers, uint8 signerCount, uint256 reportValidationGas, uint256 initialGas) returns(bool success)
func (_OracleProxy *OracleProxySession) FulfillAndBill(requestId [32]byte, response []byte, err []byte, transmitter common.Address, signers [31]common.Address, signerCount uint8, reportValidationGas *big.Int, initialGas *big.Int) (*types.Transaction, error) {
	return _OracleProxy.Contract.FulfillAndBill(&_OracleProxy.TransactOpts, requestId, response, err, transmitter, signers, signerCount, reportValidationGas, initialGas)
}

// FulfillAndBill is a paid mutator transaction binding the contract method 0x0739e4f1.
//
// Solidity: function fulfillAndBill(bytes32 requestId, bytes response, bytes err, address transmitter, address[31] signers, uint8 signerCount, uint256 reportValidationGas, uint256 initialGas) returns(bool success)
func (_OracleProxy *OracleProxyTransactorSession) FulfillAndBill(requestId [32]byte, response []byte, err []byte, transmitter common.Address, signers [31]common.Address, signerCount uint8, reportValidationGas *big.Int, initialGas *big.Int) (*types.Transaction, error) {
	return _OracleProxy.Contract.FulfillAndBill(&_OracleProxy.TransactOpts, requestId, response, err, transmitter, signers, signerCount, reportValidationGas, initialGas)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address link, address linkEthFeed, address oracle) returns()
func (_OracleProxy *OracleProxyTransactor) Initialize(opts *bind.TransactOpts, link common.Address, linkEthFeed common.Address, oracle common.Address) (*types.Transaction, error) {
	return _OracleProxy.contract.Transact(opts, "initialize", link, linkEthFeed, oracle)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address link, address linkEthFeed, address oracle) returns()
func (_OracleProxy *OracleProxySession) Initialize(link common.Address, linkEthFeed common.Address, oracle common.Address) (*types.Transaction, error) {
	return _OracleProxy.Contract.Initialize(&_OracleProxy.TransactOpts, link, linkEthFeed, oracle)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address link, address linkEthFeed, address oracle) returns()
func (_OracleProxy *OracleProxyTransactorSession) Initialize(link common.Address, linkEthFeed common.Address, oracle common.Address) (*types.Transaction, error) {
	return _OracleProxy.Contract.Initialize(&_OracleProxy.TransactOpts, link, linkEthFeed, oracle)
}

// OnTokenTransfer is a paid mutator transaction binding the contract method 0xa4c0ed36.
//
// Solidity: function onTokenTransfer(address , uint256 amount, bytes data) returns()
func (_OracleProxy *OracleProxyTransactor) OnTokenTransfer(opts *bind.TransactOpts, arg0 common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _OracleProxy.contract.Transact(opts, "onTokenTransfer", arg0, amount, data)
}

// OnTokenTransfer is a paid mutator transaction binding the contract method 0xa4c0ed36.
//
// Solidity: function onTokenTransfer(address , uint256 amount, bytes data) returns()
func (_OracleProxy *OracleProxySession) OnTokenTransfer(arg0 common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _OracleProxy.Contract.OnTokenTransfer(&_OracleProxy.TransactOpts, arg0, amount, data)
}

// OnTokenTransfer is a paid mutator transaction binding the contract method 0xa4c0ed36.
//
// Solidity: function onTokenTransfer(address , uint256 amount, bytes data) returns()
func (_OracleProxy *OracleProxyTransactorSession) OnTokenTransfer(arg0 common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _OracleProxy.Contract.OnTokenTransfer(&_OracleProxy.TransactOpts, arg0, amount, data)
}

// OracleWithdraw is a paid mutator transaction binding the contract method 0x66316d8d.
//
// Solidity: function oracleWithdraw(address recipient, uint96 amount) returns()
func (_OracleProxy *OracleProxyTransactor) OracleWithdraw(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OracleProxy.contract.Transact(opts, "oracleWithdraw", recipient, amount)
}

// OracleWithdraw is a paid mutator transaction binding the contract method 0x66316d8d.
//
// Solidity: function oracleWithdraw(address recipient, uint96 amount) returns()
func (_OracleProxy *OracleProxySession) OracleWithdraw(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OracleProxy.Contract.OracleWithdraw(&_OracleProxy.TransactOpts, recipient, amount)
}

// OracleWithdraw is a paid mutator transaction binding the contract method 0x66316d8d.
//
// Solidity: function oracleWithdraw(address recipient, uint96 amount) returns()
func (_OracleProxy *OracleProxyTransactorSession) OracleWithdraw(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OracleProxy.Contract.OracleWithdraw(&_OracleProxy.TransactOpts, recipient, amount)
}

// OwnerCancelSubscription is a paid mutator transaction binding the contract method 0x02bcc5b6.
//
// Solidity: function ownerCancelSubscription(uint64 subscriptionId) returns()
func (_OracleProxy *OracleProxyTransactor) OwnerCancelSubscription(opts *bind.TransactOpts, subscriptionId uint64) (*types.Transaction, error) {
	return _OracleProxy.contract.Transact(opts, "ownerCancelSubscription", subscriptionId)
}

// OwnerCancelSubscription is a paid mutator transaction binding the contract method 0x02bcc5b6.
//
// Solidity: function ownerCancelSubscription(uint64 subscriptionId) returns()
func (_OracleProxy *OracleProxySession) OwnerCancelSubscription(subscriptionId uint64) (*types.Transaction, error) {
	return _OracleProxy.Contract.OwnerCancelSubscription(&_OracleProxy.TransactOpts, subscriptionId)
}

// OwnerCancelSubscription is a paid mutator transaction binding the contract method 0x02bcc5b6.
//
// Solidity: function ownerCancelSubscription(uint64 subscriptionId) returns()
func (_OracleProxy *OracleProxyTransactorSession) OwnerCancelSubscription(subscriptionId uint64) (*types.Transaction, error) {
	return _OracleProxy.Contract.OwnerCancelSubscription(&_OracleProxy.TransactOpts, subscriptionId)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_OracleProxy *OracleProxyTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleProxy.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_OracleProxy *OracleProxySession) Pause() (*types.Transaction, error) {
	return _OracleProxy.Contract.Pause(&_OracleProxy.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_OracleProxy *OracleProxyTransactorSession) Pause() (*types.Transaction, error) {
	return _OracleProxy.Contract.Pause(&_OracleProxy.TransactOpts)
}

// RecoverFunds is a paid mutator transaction binding the contract method 0xe72f6e30.
//
// Solidity: function recoverFunds(address to) returns()
func (_OracleProxy *OracleProxyTransactor) RecoverFunds(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _OracleProxy.contract.Transact(opts, "recoverFunds", to)
}

// RecoverFunds is a paid mutator transaction binding the contract method 0xe72f6e30.
//
// Solidity: function recoverFunds(address to) returns()
func (_OracleProxy *OracleProxySession) RecoverFunds(to common.Address) (*types.Transaction, error) {
	return _OracleProxy.Contract.RecoverFunds(&_OracleProxy.TransactOpts, to)
}

// RecoverFunds is a paid mutator transaction binding the contract method 0xe72f6e30.
//
// Solidity: function recoverFunds(address to) returns()
func (_OracleProxy *OracleProxyTransactorSession) RecoverFunds(to common.Address) (*types.Transaction, error) {
	return _OracleProxy.Contract.RecoverFunds(&_OracleProxy.TransactOpts, to)
}

// RemoveConsumer is a paid mutator transaction binding the contract method 0x9f87fad7.
//
// Solidity: function removeConsumer(uint64 subscriptionId, address consumer) returns()
func (_OracleProxy *OracleProxyTransactor) RemoveConsumer(opts *bind.TransactOpts, subscriptionId uint64, consumer common.Address) (*types.Transaction, error) {
	return _OracleProxy.contract.Transact(opts, "removeConsumer", subscriptionId, consumer)
}

// RemoveConsumer is a paid mutator transaction binding the contract method 0x9f87fad7.
//
// Solidity: function removeConsumer(uint64 subscriptionId, address consumer) returns()
func (_OracleProxy *OracleProxySession) RemoveConsumer(subscriptionId uint64, consumer common.Address) (*types.Transaction, error) {
	return _OracleProxy.Contract.RemoveConsumer(&_OracleProxy.TransactOpts, subscriptionId, consumer)
}

// RemoveConsumer is a paid mutator transaction binding the contract method 0x9f87fad7.
//
// Solidity: function removeConsumer(uint64 subscriptionId, address consumer) returns()
func (_OracleProxy *OracleProxyTransactorSession) RemoveConsumer(subscriptionId uint64, consumer common.Address) (*types.Transaction, error) {
	return _OracleProxy.Contract.RemoveConsumer(&_OracleProxy.TransactOpts, subscriptionId, consumer)
}

// RequestSubscriptionOwnerTransfer is a paid mutator transaction binding the contract method 0x04c357cb.
//
// Solidity: function requestSubscriptionOwnerTransfer(uint64 subscriptionId, address newOwner) returns()
func (_OracleProxy *OracleProxyTransactor) RequestSubscriptionOwnerTransfer(opts *bind.TransactOpts, subscriptionId uint64, newOwner common.Address) (*types.Transaction, error) {
	return _OracleProxy.contract.Transact(opts, "requestSubscriptionOwnerTransfer", subscriptionId, newOwner)
}

// RequestSubscriptionOwnerTransfer is a paid mutator transaction binding the contract method 0x04c357cb.
//
// Solidity: function requestSubscriptionOwnerTransfer(uint64 subscriptionId, address newOwner) returns()
func (_OracleProxy *OracleProxySession) RequestSubscriptionOwnerTransfer(subscriptionId uint64, newOwner common.Address) (*types.Transaction, error) {
	return _OracleProxy.Contract.RequestSubscriptionOwnerTransfer(&_OracleProxy.TransactOpts, subscriptionId, newOwner)
}

// RequestSubscriptionOwnerTransfer is a paid mutator transaction binding the contract method 0x04c357cb.
//
// Solidity: function requestSubscriptionOwnerTransfer(uint64 subscriptionId, address newOwner) returns()
func (_OracleProxy *OracleProxyTransactorSession) RequestSubscriptionOwnerTransfer(subscriptionId uint64, newOwner common.Address) (*types.Transaction, error) {
	return _OracleProxy.Contract.RequestSubscriptionOwnerTransfer(&_OracleProxy.TransactOpts, subscriptionId, newOwner)
}

// SetAuthorizedSenders is a paid mutator transaction binding the contract method 0xee56997b.
//
// Solidity: function setAuthorizedSenders(address[] senders) returns()
func (_OracleProxy *OracleProxyTransactor) SetAuthorizedSenders(opts *bind.TransactOpts, senders []common.Address) (*types.Transaction, error) {
	return _OracleProxy.contract.Transact(opts, "setAuthorizedSenders", senders)
}

// SetAuthorizedSenders is a paid mutator transaction binding the contract method 0xee56997b.
//
// Solidity: function setAuthorizedSenders(address[] senders) returns()
func (_OracleProxy *OracleProxySession) SetAuthorizedSenders(senders []common.Address) (*types.Transaction, error) {
	return _OracleProxy.Contract.SetAuthorizedSenders(&_OracleProxy.TransactOpts, senders)
}

// SetAuthorizedSenders is a paid mutator transaction binding the contract method 0xee56997b.
//
// Solidity: function setAuthorizedSenders(address[] senders) returns()
func (_OracleProxy *OracleProxyTransactorSession) SetAuthorizedSenders(senders []common.Address) (*types.Transaction, error) {
	return _OracleProxy.Contract.SetAuthorizedSenders(&_OracleProxy.TransactOpts, senders)
}

// SetConfig is a paid mutator transaction binding the contract method 0x27923e41.
//
// Solidity: function setConfig(uint32 maxGasLimit, uint32 stalenessSeconds, uint256 gasAfterPaymentCalculation, int256 fallbackWeiPerUnitLink, uint32 gasOverhead, uint32 requestTimeoutSeconds) returns()
func (_OracleProxy *OracleProxyTransactor) SetConfig(opts *bind.TransactOpts, maxGasLimit uint32, stalenessSeconds uint32, gasAfterPaymentCalculation *big.Int, fallbackWeiPerUnitLink *big.Int, gasOverhead uint32, requestTimeoutSeconds uint32) (*types.Transaction, error) {
	return _OracleProxy.contract.Transact(opts, "setConfig", maxGasLimit, stalenessSeconds, gasAfterPaymentCalculation, fallbackWeiPerUnitLink, gasOverhead, requestTimeoutSeconds)
}

// SetConfig is a paid mutator transaction binding the contract method 0x27923e41.
//
// Solidity: function setConfig(uint32 maxGasLimit, uint32 stalenessSeconds, uint256 gasAfterPaymentCalculation, int256 fallbackWeiPerUnitLink, uint32 gasOverhead, uint32 requestTimeoutSeconds) returns()
func (_OracleProxy *OracleProxySession) SetConfig(maxGasLimit uint32, stalenessSeconds uint32, gasAfterPaymentCalculation *big.Int, fallbackWeiPerUnitLink *big.Int, gasOverhead uint32, requestTimeoutSeconds uint32) (*types.Transaction, error) {
	return _OracleProxy.Contract.SetConfig(&_OracleProxy.TransactOpts, maxGasLimit, stalenessSeconds, gasAfterPaymentCalculation, fallbackWeiPerUnitLink, gasOverhead, requestTimeoutSeconds)
}

// SetConfig is a paid mutator transaction binding the contract method 0x27923e41.
//
// Solidity: function setConfig(uint32 maxGasLimit, uint32 stalenessSeconds, uint256 gasAfterPaymentCalculation, int256 fallbackWeiPerUnitLink, uint32 gasOverhead, uint32 requestTimeoutSeconds) returns()
func (_OracleProxy *OracleProxyTransactorSession) SetConfig(maxGasLimit uint32, stalenessSeconds uint32, gasAfterPaymentCalculation *big.Int, fallbackWeiPerUnitLink *big.Int, gasOverhead uint32, requestTimeoutSeconds uint32) (*types.Transaction, error) {
	return _OracleProxy.Contract.SetConfig(&_OracleProxy.TransactOpts, maxGasLimit, stalenessSeconds, gasAfterPaymentCalculation, fallbackWeiPerUnitLink, gasOverhead, requestTimeoutSeconds)
}

// StartBilling is a paid mutator transaction binding the contract method 0xa9d03c05.
//
// Solidity: function startBilling(bytes data, (uint64,address,uint32,uint256) billing) returns(bytes32)
func (_OracleProxy *OracleProxyTransactor) StartBilling(opts *bind.TransactOpts, data []byte, billing FunctionsBillingRegistryInterfaceRequestBilling) (*types.Transaction, error) {
	return _OracleProxy.contract.Transact(opts, "startBilling", data, billing)
}

// StartBilling is a paid mutator transaction binding the contract method 0xa9d03c05.
//
// Solidity: function startBilling(bytes data, (uint64,address,uint32,uint256) billing) returns(bytes32)
func (_OracleProxy *OracleProxySession) StartBilling(data []byte, billing FunctionsBillingRegistryInterfaceRequestBilling) (*types.Transaction, error) {
	return _OracleProxy.Contract.StartBilling(&_OracleProxy.TransactOpts, data, billing)
}

// StartBilling is a paid mutator transaction binding the contract method 0xa9d03c05.
//
// Solidity: function startBilling(bytes data, (uint64,address,uint32,uint256) billing) returns(bytes32)
func (_OracleProxy *OracleProxyTransactorSession) StartBilling(data []byte, billing FunctionsBillingRegistryInterfaceRequestBilling) (*types.Transaction, error) {
	return _OracleProxy.Contract.StartBilling(&_OracleProxy.TransactOpts, data, billing)
}

// TimeoutRequests is a paid mutator transaction binding the contract method 0x665871ec.
//
// Solidity: function timeoutRequests(bytes32[] requestIdsToTimeout) returns()
func (_OracleProxy *OracleProxyTransactor) TimeoutRequests(opts *bind.TransactOpts, requestIdsToTimeout [][32]byte) (*types.Transaction, error) {
	return _OracleProxy.contract.Transact(opts, "timeoutRequests", requestIdsToTimeout)
}

// TimeoutRequests is a paid mutator transaction binding the contract method 0x665871ec.
//
// Solidity: function timeoutRequests(bytes32[] requestIdsToTimeout) returns()
func (_OracleProxy *OracleProxySession) TimeoutRequests(requestIdsToTimeout [][32]byte) (*types.Transaction, error) {
	return _OracleProxy.Contract.TimeoutRequests(&_OracleProxy.TransactOpts, requestIdsToTimeout)
}

// TimeoutRequests is a paid mutator transaction binding the contract method 0x665871ec.
//
// Solidity: function timeoutRequests(bytes32[] requestIdsToTimeout) returns()
func (_OracleProxy *OracleProxyTransactorSession) TimeoutRequests(requestIdsToTimeout [][32]byte) (*types.Transaction, error) {
	return _OracleProxy.Contract.TimeoutRequests(&_OracleProxy.TransactOpts, requestIdsToTimeout)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address to) returns()
func (_OracleProxy *OracleProxyTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _OracleProxy.contract.Transact(opts, "transferOwnership", to)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address to) returns()
func (_OracleProxy *OracleProxySession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _OracleProxy.Contract.TransferOwnership(&_OracleProxy.TransactOpts, to)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address to) returns()
func (_OracleProxy *OracleProxyTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _OracleProxy.Contract.TransferOwnership(&_OracleProxy.TransactOpts, to)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_OracleProxy *OracleProxyTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleProxy.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_OracleProxy *OracleProxySession) Unpause() (*types.Transaction, error) {
	return _OracleProxy.Contract.Unpause(&_OracleProxy.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_OracleProxy *OracleProxyTransactorSession) Unpause() (*types.Transaction, error) {
	return _OracleProxy.Contract.Unpause(&_OracleProxy.TransactOpts)
}

// OracleProxyAuthorizedSendersChangedIterator is returned from FilterAuthorizedSendersChanged and is used to iterate over the raw logs and unpacked data for AuthorizedSendersChanged events raised by the OracleProxy contract.
type OracleProxyAuthorizedSendersChangedIterator struct {
	Event *OracleProxyAuthorizedSendersChanged // Event containing the contract specifics and raw log

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
func (it *OracleProxyAuthorizedSendersChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleProxyAuthorizedSendersChanged)
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
		it.Event = new(OracleProxyAuthorizedSendersChanged)
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
func (it *OracleProxyAuthorizedSendersChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleProxyAuthorizedSendersChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleProxyAuthorizedSendersChanged represents a AuthorizedSendersChanged event raised by the OracleProxy contract.
type OracleProxyAuthorizedSendersChanged struct {
	Senders   []common.Address
	ChangedBy common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAuthorizedSendersChanged is a free log retrieval operation binding the contract event 0xf263cfb3e4298332e776194610cf9fdc09ccb3ada8b9aa39764d882e11fbf0a0.
//
// Solidity: event AuthorizedSendersChanged(address[] senders, address changedBy)
func (_OracleProxy *OracleProxyFilterer) FilterAuthorizedSendersChanged(opts *bind.FilterOpts) (*OracleProxyAuthorizedSendersChangedIterator, error) {

	logs, sub, err := _OracleProxy.contract.FilterLogs(opts, "AuthorizedSendersChanged")
	if err != nil {
		return nil, err
	}
	return &OracleProxyAuthorizedSendersChangedIterator{contract: _OracleProxy.contract, event: "AuthorizedSendersChanged", logs: logs, sub: sub}, nil
}

// WatchAuthorizedSendersChanged is a free log subscription operation binding the contract event 0xf263cfb3e4298332e776194610cf9fdc09ccb3ada8b9aa39764d882e11fbf0a0.
//
// Solidity: event AuthorizedSendersChanged(address[] senders, address changedBy)
func (_OracleProxy *OracleProxyFilterer) WatchAuthorizedSendersChanged(opts *bind.WatchOpts, sink chan<- *OracleProxyAuthorizedSendersChanged) (event.Subscription, error) {

	logs, sub, err := _OracleProxy.contract.WatchLogs(opts, "AuthorizedSendersChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleProxyAuthorizedSendersChanged)
				if err := _OracleProxy.contract.UnpackLog(event, "AuthorizedSendersChanged", log); err != nil {
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
func (_OracleProxy *OracleProxyFilterer) ParseAuthorizedSendersChanged(log types.Log) (*OracleProxyAuthorizedSendersChanged, error) {
	event := new(OracleProxyAuthorizedSendersChanged)
	if err := _OracleProxy.contract.UnpackLog(event, "AuthorizedSendersChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleProxyBillingEndIterator is returned from FilterBillingEnd and is used to iterate over the raw logs and unpacked data for BillingEnd events raised by the OracleProxy contract.
type OracleProxyBillingEndIterator struct {
	Event *OracleProxyBillingEnd // Event containing the contract specifics and raw log

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
func (it *OracleProxyBillingEndIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleProxyBillingEnd)
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
		it.Event = new(OracleProxyBillingEnd)
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
func (it *OracleProxyBillingEndIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleProxyBillingEndIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleProxyBillingEnd represents a BillingEnd event raised by the OracleProxy contract.
type OracleProxyBillingEnd struct {
	RequestId          [32]byte
	SubscriptionId     uint64
	SignerPayment      *big.Int
	TransmitterPayment *big.Int
	TotalCost          *big.Int
	Success            bool
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterBillingEnd is a free log retrieval operation binding the contract event 0xc8dc973332de19a5f71b6026983110e9c2e04b0c98b87eb771ccb78607fd114f.
//
// Solidity: event BillingEnd(bytes32 indexed requestId, uint64 subscriptionId, uint96 signerPayment, uint96 transmitterPayment, uint96 totalCost, bool success)
func (_OracleProxy *OracleProxyFilterer) FilterBillingEnd(opts *bind.FilterOpts, requestId [][32]byte) (*OracleProxyBillingEndIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _OracleProxy.contract.FilterLogs(opts, "BillingEnd", requestIdRule)
	if err != nil {
		return nil, err
	}
	return &OracleProxyBillingEndIterator{contract: _OracleProxy.contract, event: "BillingEnd", logs: logs, sub: sub}, nil
}

// WatchBillingEnd is a free log subscription operation binding the contract event 0xc8dc973332de19a5f71b6026983110e9c2e04b0c98b87eb771ccb78607fd114f.
//
// Solidity: event BillingEnd(bytes32 indexed requestId, uint64 subscriptionId, uint96 signerPayment, uint96 transmitterPayment, uint96 totalCost, bool success)
func (_OracleProxy *OracleProxyFilterer) WatchBillingEnd(opts *bind.WatchOpts, sink chan<- *OracleProxyBillingEnd, requestId [][32]byte) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _OracleProxy.contract.WatchLogs(opts, "BillingEnd", requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleProxyBillingEnd)
				if err := _OracleProxy.contract.UnpackLog(event, "BillingEnd", log); err != nil {
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

// ParseBillingEnd is a log parse operation binding the contract event 0xc8dc973332de19a5f71b6026983110e9c2e04b0c98b87eb771ccb78607fd114f.
//
// Solidity: event BillingEnd(bytes32 indexed requestId, uint64 subscriptionId, uint96 signerPayment, uint96 transmitterPayment, uint96 totalCost, bool success)
func (_OracleProxy *OracleProxyFilterer) ParseBillingEnd(log types.Log) (*OracleProxyBillingEnd, error) {
	event := new(OracleProxyBillingEnd)
	if err := _OracleProxy.contract.UnpackLog(event, "BillingEnd", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleProxyBillingStartIterator is returned from FilterBillingStart and is used to iterate over the raw logs and unpacked data for BillingStart events raised by the OracleProxy contract.
type OracleProxyBillingStartIterator struct {
	Event *OracleProxyBillingStart // Event containing the contract specifics and raw log

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
func (it *OracleProxyBillingStartIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleProxyBillingStart)
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
		it.Event = new(OracleProxyBillingStart)
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
func (it *OracleProxyBillingStartIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleProxyBillingStartIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleProxyBillingStart represents a BillingStart event raised by the OracleProxy contract.
type OracleProxyBillingStart struct {
	RequestId  [32]byte
	Commitment FunctionsBillingRegistryCommitment
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBillingStart is a free log retrieval operation binding the contract event 0x99f7f4e65b4b9fbabd4e357c47ed3099b36e57ecd3a43e84662f34c207d0ebe4.
//
// Solidity: event BillingStart(bytes32 indexed requestId, (uint64,address,uint32,uint256,address,uint96,uint96,uint96,uint256) commitment)
func (_OracleProxy *OracleProxyFilterer) FilterBillingStart(opts *bind.FilterOpts, requestId [][32]byte) (*OracleProxyBillingStartIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _OracleProxy.contract.FilterLogs(opts, "BillingStart", requestIdRule)
	if err != nil {
		return nil, err
	}
	return &OracleProxyBillingStartIterator{contract: _OracleProxy.contract, event: "BillingStart", logs: logs, sub: sub}, nil
}

// WatchBillingStart is a free log subscription operation binding the contract event 0x99f7f4e65b4b9fbabd4e357c47ed3099b36e57ecd3a43e84662f34c207d0ebe4.
//
// Solidity: event BillingStart(bytes32 indexed requestId, (uint64,address,uint32,uint256,address,uint96,uint96,uint96,uint256) commitment)
func (_OracleProxy *OracleProxyFilterer) WatchBillingStart(opts *bind.WatchOpts, sink chan<- *OracleProxyBillingStart, requestId [][32]byte) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _OracleProxy.contract.WatchLogs(opts, "BillingStart", requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleProxyBillingStart)
				if err := _OracleProxy.contract.UnpackLog(event, "BillingStart", log); err != nil {
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

// ParseBillingStart is a log parse operation binding the contract event 0x99f7f4e65b4b9fbabd4e357c47ed3099b36e57ecd3a43e84662f34c207d0ebe4.
//
// Solidity: event BillingStart(bytes32 indexed requestId, (uint64,address,uint32,uint256,address,uint96,uint96,uint96,uint256) commitment)
func (_OracleProxy *OracleProxyFilterer) ParseBillingStart(log types.Log) (*OracleProxyBillingStart, error) {
	event := new(OracleProxyBillingStart)
	if err := _OracleProxy.contract.UnpackLog(event, "BillingStart", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleProxyConfigSetIterator is returned from FilterConfigSet and is used to iterate over the raw logs and unpacked data for ConfigSet events raised by the OracleProxy contract.
type OracleProxyConfigSetIterator struct {
	Event *OracleProxyConfigSet // Event containing the contract specifics and raw log

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
func (it *OracleProxyConfigSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleProxyConfigSet)
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
		it.Event = new(OracleProxyConfigSet)
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
func (it *OracleProxyConfigSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleProxyConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleProxyConfigSet represents a ConfigSet event raised by the OracleProxy contract.
type OracleProxyConfigSet struct {
	MaxGasLimit                uint32
	StalenessSeconds           uint32
	GasAfterPaymentCalculation *big.Int
	FallbackWeiPerUnitLink     *big.Int
	GasOverhead                uint32
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterConfigSet is a free log retrieval operation binding the contract event 0x24d3d934adfef9b9029d6ffa463c07d0139ed47d26ee23506f85ece2879d2bd4.
//
// Solidity: event ConfigSet(uint32 maxGasLimit, uint32 stalenessSeconds, uint256 gasAfterPaymentCalculation, int256 fallbackWeiPerUnitLink, uint32 gasOverhead)
func (_OracleProxy *OracleProxyFilterer) FilterConfigSet(opts *bind.FilterOpts) (*OracleProxyConfigSetIterator, error) {

	logs, sub, err := _OracleProxy.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &OracleProxyConfigSetIterator{contract: _OracleProxy.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

// WatchConfigSet is a free log subscription operation binding the contract event 0x24d3d934adfef9b9029d6ffa463c07d0139ed47d26ee23506f85ece2879d2bd4.
//
// Solidity: event ConfigSet(uint32 maxGasLimit, uint32 stalenessSeconds, uint256 gasAfterPaymentCalculation, int256 fallbackWeiPerUnitLink, uint32 gasOverhead)
func (_OracleProxy *OracleProxyFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *OracleProxyConfigSet) (event.Subscription, error) {

	logs, sub, err := _OracleProxy.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleProxyConfigSet)
				if err := _OracleProxy.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

// ParseConfigSet is a log parse operation binding the contract event 0x24d3d934adfef9b9029d6ffa463c07d0139ed47d26ee23506f85ece2879d2bd4.
//
// Solidity: event ConfigSet(uint32 maxGasLimit, uint32 stalenessSeconds, uint256 gasAfterPaymentCalculation, int256 fallbackWeiPerUnitLink, uint32 gasOverhead)
func (_OracleProxy *OracleProxyFilterer) ParseConfigSet(log types.Log) (*OracleProxyConfigSet, error) {
	event := new(OracleProxyConfigSet)
	if err := _OracleProxy.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleProxyFundsRecoveredIterator is returned from FilterFundsRecovered and is used to iterate over the raw logs and unpacked data for FundsRecovered events raised by the OracleProxy contract.
type OracleProxyFundsRecoveredIterator struct {
	Event *OracleProxyFundsRecovered // Event containing the contract specifics and raw log

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
func (it *OracleProxyFundsRecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleProxyFundsRecovered)
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
		it.Event = new(OracleProxyFundsRecovered)
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
func (it *OracleProxyFundsRecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleProxyFundsRecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleProxyFundsRecovered represents a FundsRecovered event raised by the OracleProxy contract.
type OracleProxyFundsRecovered struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFundsRecovered is a free log retrieval operation binding the contract event 0x59bfc682b673f8cbf945f1e454df9334834abf7dfe7f92237ca29ecb9b436600.
//
// Solidity: event FundsRecovered(address to, uint256 amount)
func (_OracleProxy *OracleProxyFilterer) FilterFundsRecovered(opts *bind.FilterOpts) (*OracleProxyFundsRecoveredIterator, error) {

	logs, sub, err := _OracleProxy.contract.FilterLogs(opts, "FundsRecovered")
	if err != nil {
		return nil, err
	}
	return &OracleProxyFundsRecoveredIterator{contract: _OracleProxy.contract, event: "FundsRecovered", logs: logs, sub: sub}, nil
}

// WatchFundsRecovered is a free log subscription operation binding the contract event 0x59bfc682b673f8cbf945f1e454df9334834abf7dfe7f92237ca29ecb9b436600.
//
// Solidity: event FundsRecovered(address to, uint256 amount)
func (_OracleProxy *OracleProxyFilterer) WatchFundsRecovered(opts *bind.WatchOpts, sink chan<- *OracleProxyFundsRecovered) (event.Subscription, error) {

	logs, sub, err := _OracleProxy.contract.WatchLogs(opts, "FundsRecovered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleProxyFundsRecovered)
				if err := _OracleProxy.contract.UnpackLog(event, "FundsRecovered", log); err != nil {
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

// ParseFundsRecovered is a log parse operation binding the contract event 0x59bfc682b673f8cbf945f1e454df9334834abf7dfe7f92237ca29ecb9b436600.
//
// Solidity: event FundsRecovered(address to, uint256 amount)
func (_OracleProxy *OracleProxyFilterer) ParseFundsRecovered(log types.Log) (*OracleProxyFundsRecovered, error) {
	event := new(OracleProxyFundsRecovered)
	if err := _OracleProxy.contract.UnpackLog(event, "FundsRecovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleProxyInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the OracleProxy contract.
type OracleProxyInitializedIterator struct {
	Event *OracleProxyInitialized // Event containing the contract specifics and raw log

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
func (it *OracleProxyInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleProxyInitialized)
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
		it.Event = new(OracleProxyInitialized)
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
func (it *OracleProxyInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleProxyInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleProxyInitialized represents a Initialized event raised by the OracleProxy contract.
type OracleProxyInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_OracleProxy *OracleProxyFilterer) FilterInitialized(opts *bind.FilterOpts) (*OracleProxyInitializedIterator, error) {

	logs, sub, err := _OracleProxy.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &OracleProxyInitializedIterator{contract: _OracleProxy.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_OracleProxy *OracleProxyFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *OracleProxyInitialized) (event.Subscription, error) {

	logs, sub, err := _OracleProxy.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleProxyInitialized)
				if err := _OracleProxy.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_OracleProxy *OracleProxyFilterer) ParseInitialized(log types.Log) (*OracleProxyInitialized, error) {
	event := new(OracleProxyInitialized)
	if err := _OracleProxy.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleProxyOwnershipTransferRequestedIterator is returned from FilterOwnershipTransferRequested and is used to iterate over the raw logs and unpacked data for OwnershipTransferRequested events raised by the OracleProxy contract.
type OracleProxyOwnershipTransferRequestedIterator struct {
	Event *OracleProxyOwnershipTransferRequested // Event containing the contract specifics and raw log

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
func (it *OracleProxyOwnershipTransferRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleProxyOwnershipTransferRequested)
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
		it.Event = new(OracleProxyOwnershipTransferRequested)
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
func (it *OracleProxyOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleProxyOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleProxyOwnershipTransferRequested represents a OwnershipTransferRequested event raised by the OracleProxy contract.
type OracleProxyOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferRequested is a free log retrieval operation binding the contract event 0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278.
//
// Solidity: event OwnershipTransferRequested(address indexed from, address indexed to)
func (_OracleProxy *OracleProxyFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OracleProxyOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OracleProxy.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OracleProxyOwnershipTransferRequestedIterator{contract: _OracleProxy.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferRequested is a free log subscription operation binding the contract event 0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278.
//
// Solidity: event OwnershipTransferRequested(address indexed from, address indexed to)
func (_OracleProxy *OracleProxyFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *OracleProxyOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OracleProxy.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleProxyOwnershipTransferRequested)
				if err := _OracleProxy.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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
func (_OracleProxy *OracleProxyFilterer) ParseOwnershipTransferRequested(log types.Log) (*OracleProxyOwnershipTransferRequested, error) {
	event := new(OracleProxyOwnershipTransferRequested)
	if err := _OracleProxy.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleProxyOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the OracleProxy contract.
type OracleProxyOwnershipTransferredIterator struct {
	Event *OracleProxyOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OracleProxyOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleProxyOwnershipTransferred)
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
		it.Event = new(OracleProxyOwnershipTransferred)
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
func (it *OracleProxyOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleProxyOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleProxyOwnershipTransferred represents a OwnershipTransferred event raised by the OracleProxy contract.
type OracleProxyOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed from, address indexed to)
func (_OracleProxy *OracleProxyFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OracleProxyOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OracleProxy.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OracleProxyOwnershipTransferredIterator{contract: _OracleProxy.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed from, address indexed to)
func (_OracleProxy *OracleProxyFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OracleProxyOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OracleProxy.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleProxyOwnershipTransferred)
				if err := _OracleProxy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_OracleProxy *OracleProxyFilterer) ParseOwnershipTransferred(log types.Log) (*OracleProxyOwnershipTransferred, error) {
	event := new(OracleProxyOwnershipTransferred)
	if err := _OracleProxy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleProxyPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the OracleProxy contract.
type OracleProxyPausedIterator struct {
	Event *OracleProxyPaused // Event containing the contract specifics and raw log

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
func (it *OracleProxyPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleProxyPaused)
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
		it.Event = new(OracleProxyPaused)
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
func (it *OracleProxyPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleProxyPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleProxyPaused represents a Paused event raised by the OracleProxy contract.
type OracleProxyPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_OracleProxy *OracleProxyFilterer) FilterPaused(opts *bind.FilterOpts) (*OracleProxyPausedIterator, error) {

	logs, sub, err := _OracleProxy.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &OracleProxyPausedIterator{contract: _OracleProxy.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_OracleProxy *OracleProxyFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *OracleProxyPaused) (event.Subscription, error) {

	logs, sub, err := _OracleProxy.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleProxyPaused)
				if err := _OracleProxy.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_OracleProxy *OracleProxyFilterer) ParsePaused(log types.Log) (*OracleProxyPaused, error) {
	event := new(OracleProxyPaused)
	if err := _OracleProxy.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleProxyRequestTimedOutIterator is returned from FilterRequestTimedOut and is used to iterate over the raw logs and unpacked data for RequestTimedOut events raised by the OracleProxy contract.
type OracleProxyRequestTimedOutIterator struct {
	Event *OracleProxyRequestTimedOut // Event containing the contract specifics and raw log

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
func (it *OracleProxyRequestTimedOutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleProxyRequestTimedOut)
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
		it.Event = new(OracleProxyRequestTimedOut)
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
func (it *OracleProxyRequestTimedOutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleProxyRequestTimedOutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleProxyRequestTimedOut represents a RequestTimedOut event raised by the OracleProxy contract.
type OracleProxyRequestTimedOut struct {
	RequestId [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRequestTimedOut is a free log retrieval operation binding the contract event 0xf1ca1e9147be737b04a2b018a79405f687a97de8dd8a2559bbe62357343af414.
//
// Solidity: event RequestTimedOut(bytes32 indexed requestId)
func (_OracleProxy *OracleProxyFilterer) FilterRequestTimedOut(opts *bind.FilterOpts, requestId [][32]byte) (*OracleProxyRequestTimedOutIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _OracleProxy.contract.FilterLogs(opts, "RequestTimedOut", requestIdRule)
	if err != nil {
		return nil, err
	}
	return &OracleProxyRequestTimedOutIterator{contract: _OracleProxy.contract, event: "RequestTimedOut", logs: logs, sub: sub}, nil
}

// WatchRequestTimedOut is a free log subscription operation binding the contract event 0xf1ca1e9147be737b04a2b018a79405f687a97de8dd8a2559bbe62357343af414.
//
// Solidity: event RequestTimedOut(bytes32 indexed requestId)
func (_OracleProxy *OracleProxyFilterer) WatchRequestTimedOut(opts *bind.WatchOpts, sink chan<- *OracleProxyRequestTimedOut, requestId [][32]byte) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _OracleProxy.contract.WatchLogs(opts, "RequestTimedOut", requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleProxyRequestTimedOut)
				if err := _OracleProxy.contract.UnpackLog(event, "RequestTimedOut", log); err != nil {
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

// ParseRequestTimedOut is a log parse operation binding the contract event 0xf1ca1e9147be737b04a2b018a79405f687a97de8dd8a2559bbe62357343af414.
//
// Solidity: event RequestTimedOut(bytes32 indexed requestId)
func (_OracleProxy *OracleProxyFilterer) ParseRequestTimedOut(log types.Log) (*OracleProxyRequestTimedOut, error) {
	event := new(OracleProxyRequestTimedOut)
	if err := _OracleProxy.contract.UnpackLog(event, "RequestTimedOut", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleProxySubscriptionCanceledIterator is returned from FilterSubscriptionCanceled and is used to iterate over the raw logs and unpacked data for SubscriptionCanceled events raised by the OracleProxy contract.
type OracleProxySubscriptionCanceledIterator struct {
	Event *OracleProxySubscriptionCanceled // Event containing the contract specifics and raw log

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
func (it *OracleProxySubscriptionCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleProxySubscriptionCanceled)
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
		it.Event = new(OracleProxySubscriptionCanceled)
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
func (it *OracleProxySubscriptionCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleProxySubscriptionCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleProxySubscriptionCanceled represents a SubscriptionCanceled event raised by the OracleProxy contract.
type OracleProxySubscriptionCanceled struct {
	SubscriptionId uint64
	To             common.Address
	Amount         *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSubscriptionCanceled is a free log retrieval operation binding the contract event 0xe8ed5b475a5b5987aa9165e8731bb78043f39eee32ec5a1169a89e27fcd49815.
//
// Solidity: event SubscriptionCanceled(uint64 indexed subscriptionId, address to, uint256 amount)
func (_OracleProxy *OracleProxyFilterer) FilterSubscriptionCanceled(opts *bind.FilterOpts, subscriptionId []uint64) (*OracleProxySubscriptionCanceledIterator, error) {

	var subscriptionIdRule []interface{}
	for _, subscriptionIdItem := range subscriptionId {
		subscriptionIdRule = append(subscriptionIdRule, subscriptionIdItem)
	}

	logs, sub, err := _OracleProxy.contract.FilterLogs(opts, "SubscriptionCanceled", subscriptionIdRule)
	if err != nil {
		return nil, err
	}
	return &OracleProxySubscriptionCanceledIterator{contract: _OracleProxy.contract, event: "SubscriptionCanceled", logs: logs, sub: sub}, nil
}

// WatchSubscriptionCanceled is a free log subscription operation binding the contract event 0xe8ed5b475a5b5987aa9165e8731bb78043f39eee32ec5a1169a89e27fcd49815.
//
// Solidity: event SubscriptionCanceled(uint64 indexed subscriptionId, address to, uint256 amount)
func (_OracleProxy *OracleProxyFilterer) WatchSubscriptionCanceled(opts *bind.WatchOpts, sink chan<- *OracleProxySubscriptionCanceled, subscriptionId []uint64) (event.Subscription, error) {

	var subscriptionIdRule []interface{}
	for _, subscriptionIdItem := range subscriptionId {
		subscriptionIdRule = append(subscriptionIdRule, subscriptionIdItem)
	}

	logs, sub, err := _OracleProxy.contract.WatchLogs(opts, "SubscriptionCanceled", subscriptionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleProxySubscriptionCanceled)
				if err := _OracleProxy.contract.UnpackLog(event, "SubscriptionCanceled", log); err != nil {
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

// ParseSubscriptionCanceled is a log parse operation binding the contract event 0xe8ed5b475a5b5987aa9165e8731bb78043f39eee32ec5a1169a89e27fcd49815.
//
// Solidity: event SubscriptionCanceled(uint64 indexed subscriptionId, address to, uint256 amount)
func (_OracleProxy *OracleProxyFilterer) ParseSubscriptionCanceled(log types.Log) (*OracleProxySubscriptionCanceled, error) {
	event := new(OracleProxySubscriptionCanceled)
	if err := _OracleProxy.contract.UnpackLog(event, "SubscriptionCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleProxySubscriptionConsumerAddedIterator is returned from FilterSubscriptionConsumerAdded and is used to iterate over the raw logs and unpacked data for SubscriptionConsumerAdded events raised by the OracleProxy contract.
type OracleProxySubscriptionConsumerAddedIterator struct {
	Event *OracleProxySubscriptionConsumerAdded // Event containing the contract specifics and raw log

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
func (it *OracleProxySubscriptionConsumerAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleProxySubscriptionConsumerAdded)
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
		it.Event = new(OracleProxySubscriptionConsumerAdded)
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
func (it *OracleProxySubscriptionConsumerAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleProxySubscriptionConsumerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleProxySubscriptionConsumerAdded represents a SubscriptionConsumerAdded event raised by the OracleProxy contract.
type OracleProxySubscriptionConsumerAdded struct {
	SubscriptionId uint64
	Consumer       common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSubscriptionConsumerAdded is a free log retrieval operation binding the contract event 0x43dc749a04ac8fb825cbd514f7c0e13f13bc6f2ee66043b76629d51776cff8e0.
//
// Solidity: event SubscriptionConsumerAdded(uint64 indexed subscriptionId, address consumer)
func (_OracleProxy *OracleProxyFilterer) FilterSubscriptionConsumerAdded(opts *bind.FilterOpts, subscriptionId []uint64) (*OracleProxySubscriptionConsumerAddedIterator, error) {

	var subscriptionIdRule []interface{}
	for _, subscriptionIdItem := range subscriptionId {
		subscriptionIdRule = append(subscriptionIdRule, subscriptionIdItem)
	}

	logs, sub, err := _OracleProxy.contract.FilterLogs(opts, "SubscriptionConsumerAdded", subscriptionIdRule)
	if err != nil {
		return nil, err
	}
	return &OracleProxySubscriptionConsumerAddedIterator{contract: _OracleProxy.contract, event: "SubscriptionConsumerAdded", logs: logs, sub: sub}, nil
}

// WatchSubscriptionConsumerAdded is a free log subscription operation binding the contract event 0x43dc749a04ac8fb825cbd514f7c0e13f13bc6f2ee66043b76629d51776cff8e0.
//
// Solidity: event SubscriptionConsumerAdded(uint64 indexed subscriptionId, address consumer)
func (_OracleProxy *OracleProxyFilterer) WatchSubscriptionConsumerAdded(opts *bind.WatchOpts, sink chan<- *OracleProxySubscriptionConsumerAdded, subscriptionId []uint64) (event.Subscription, error) {

	var subscriptionIdRule []interface{}
	for _, subscriptionIdItem := range subscriptionId {
		subscriptionIdRule = append(subscriptionIdRule, subscriptionIdItem)
	}

	logs, sub, err := _OracleProxy.contract.WatchLogs(opts, "SubscriptionConsumerAdded", subscriptionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleProxySubscriptionConsumerAdded)
				if err := _OracleProxy.contract.UnpackLog(event, "SubscriptionConsumerAdded", log); err != nil {
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

// ParseSubscriptionConsumerAdded is a log parse operation binding the contract event 0x43dc749a04ac8fb825cbd514f7c0e13f13bc6f2ee66043b76629d51776cff8e0.
//
// Solidity: event SubscriptionConsumerAdded(uint64 indexed subscriptionId, address consumer)
func (_OracleProxy *OracleProxyFilterer) ParseSubscriptionConsumerAdded(log types.Log) (*OracleProxySubscriptionConsumerAdded, error) {
	event := new(OracleProxySubscriptionConsumerAdded)
	if err := _OracleProxy.contract.UnpackLog(event, "SubscriptionConsumerAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleProxySubscriptionConsumerRemovedIterator is returned from FilterSubscriptionConsumerRemoved and is used to iterate over the raw logs and unpacked data for SubscriptionConsumerRemoved events raised by the OracleProxy contract.
type OracleProxySubscriptionConsumerRemovedIterator struct {
	Event *OracleProxySubscriptionConsumerRemoved // Event containing the contract specifics and raw log

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
func (it *OracleProxySubscriptionConsumerRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleProxySubscriptionConsumerRemoved)
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
		it.Event = new(OracleProxySubscriptionConsumerRemoved)
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
func (it *OracleProxySubscriptionConsumerRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleProxySubscriptionConsumerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleProxySubscriptionConsumerRemoved represents a SubscriptionConsumerRemoved event raised by the OracleProxy contract.
type OracleProxySubscriptionConsumerRemoved struct {
	SubscriptionId uint64
	Consumer       common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSubscriptionConsumerRemoved is a free log retrieval operation binding the contract event 0x182bff9831466789164ca77075fffd84916d35a8180ba73c27e45634549b445b.
//
// Solidity: event SubscriptionConsumerRemoved(uint64 indexed subscriptionId, address consumer)
func (_OracleProxy *OracleProxyFilterer) FilterSubscriptionConsumerRemoved(opts *bind.FilterOpts, subscriptionId []uint64) (*OracleProxySubscriptionConsumerRemovedIterator, error) {

	var subscriptionIdRule []interface{}
	for _, subscriptionIdItem := range subscriptionId {
		subscriptionIdRule = append(subscriptionIdRule, subscriptionIdItem)
	}

	logs, sub, err := _OracleProxy.contract.FilterLogs(opts, "SubscriptionConsumerRemoved", subscriptionIdRule)
	if err != nil {
		return nil, err
	}
	return &OracleProxySubscriptionConsumerRemovedIterator{contract: _OracleProxy.contract, event: "SubscriptionConsumerRemoved", logs: logs, sub: sub}, nil
}

// WatchSubscriptionConsumerRemoved is a free log subscription operation binding the contract event 0x182bff9831466789164ca77075fffd84916d35a8180ba73c27e45634549b445b.
//
// Solidity: event SubscriptionConsumerRemoved(uint64 indexed subscriptionId, address consumer)
func (_OracleProxy *OracleProxyFilterer) WatchSubscriptionConsumerRemoved(opts *bind.WatchOpts, sink chan<- *OracleProxySubscriptionConsumerRemoved, subscriptionId []uint64) (event.Subscription, error) {

	var subscriptionIdRule []interface{}
	for _, subscriptionIdItem := range subscriptionId {
		subscriptionIdRule = append(subscriptionIdRule, subscriptionIdItem)
	}

	logs, sub, err := _OracleProxy.contract.WatchLogs(opts, "SubscriptionConsumerRemoved", subscriptionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleProxySubscriptionConsumerRemoved)
				if err := _OracleProxy.contract.UnpackLog(event, "SubscriptionConsumerRemoved", log); err != nil {
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

// ParseSubscriptionConsumerRemoved is a log parse operation binding the contract event 0x182bff9831466789164ca77075fffd84916d35a8180ba73c27e45634549b445b.
//
// Solidity: event SubscriptionConsumerRemoved(uint64 indexed subscriptionId, address consumer)
func (_OracleProxy *OracleProxyFilterer) ParseSubscriptionConsumerRemoved(log types.Log) (*OracleProxySubscriptionConsumerRemoved, error) {
	event := new(OracleProxySubscriptionConsumerRemoved)
	if err := _OracleProxy.contract.UnpackLog(event, "SubscriptionConsumerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleProxySubscriptionCreatedIterator is returned from FilterSubscriptionCreated and is used to iterate over the raw logs and unpacked data for SubscriptionCreated events raised by the OracleProxy contract.
type OracleProxySubscriptionCreatedIterator struct {
	Event *OracleProxySubscriptionCreated // Event containing the contract specifics and raw log

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
func (it *OracleProxySubscriptionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleProxySubscriptionCreated)
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
		it.Event = new(OracleProxySubscriptionCreated)
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
func (it *OracleProxySubscriptionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleProxySubscriptionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleProxySubscriptionCreated represents a SubscriptionCreated event raised by the OracleProxy contract.
type OracleProxySubscriptionCreated struct {
	SubscriptionId uint64
	Owner          common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSubscriptionCreated is a free log retrieval operation binding the contract event 0x464722b4166576d3dcbba877b999bc35cf911f4eaf434b7eba68fa113951d0bf.
//
// Solidity: event SubscriptionCreated(uint64 indexed subscriptionId, address owner)
func (_OracleProxy *OracleProxyFilterer) FilterSubscriptionCreated(opts *bind.FilterOpts, subscriptionId []uint64) (*OracleProxySubscriptionCreatedIterator, error) {

	var subscriptionIdRule []interface{}
	for _, subscriptionIdItem := range subscriptionId {
		subscriptionIdRule = append(subscriptionIdRule, subscriptionIdItem)
	}

	logs, sub, err := _OracleProxy.contract.FilterLogs(opts, "SubscriptionCreated", subscriptionIdRule)
	if err != nil {
		return nil, err
	}
	return &OracleProxySubscriptionCreatedIterator{contract: _OracleProxy.contract, event: "SubscriptionCreated", logs: logs, sub: sub}, nil
}

// WatchSubscriptionCreated is a free log subscription operation binding the contract event 0x464722b4166576d3dcbba877b999bc35cf911f4eaf434b7eba68fa113951d0bf.
//
// Solidity: event SubscriptionCreated(uint64 indexed subscriptionId, address owner)
func (_OracleProxy *OracleProxyFilterer) WatchSubscriptionCreated(opts *bind.WatchOpts, sink chan<- *OracleProxySubscriptionCreated, subscriptionId []uint64) (event.Subscription, error) {

	var subscriptionIdRule []interface{}
	for _, subscriptionIdItem := range subscriptionId {
		subscriptionIdRule = append(subscriptionIdRule, subscriptionIdItem)
	}

	logs, sub, err := _OracleProxy.contract.WatchLogs(opts, "SubscriptionCreated", subscriptionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleProxySubscriptionCreated)
				if err := _OracleProxy.contract.UnpackLog(event, "SubscriptionCreated", log); err != nil {
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

// ParseSubscriptionCreated is a log parse operation binding the contract event 0x464722b4166576d3dcbba877b999bc35cf911f4eaf434b7eba68fa113951d0bf.
//
// Solidity: event SubscriptionCreated(uint64 indexed subscriptionId, address owner)
func (_OracleProxy *OracleProxyFilterer) ParseSubscriptionCreated(log types.Log) (*OracleProxySubscriptionCreated, error) {
	event := new(OracleProxySubscriptionCreated)
	if err := _OracleProxy.contract.UnpackLog(event, "SubscriptionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleProxySubscriptionFundedIterator is returned from FilterSubscriptionFunded and is used to iterate over the raw logs and unpacked data for SubscriptionFunded events raised by the OracleProxy contract.
type OracleProxySubscriptionFundedIterator struct {
	Event *OracleProxySubscriptionFunded // Event containing the contract specifics and raw log

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
func (it *OracleProxySubscriptionFundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleProxySubscriptionFunded)
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
		it.Event = new(OracleProxySubscriptionFunded)
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
func (it *OracleProxySubscriptionFundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleProxySubscriptionFundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleProxySubscriptionFunded represents a SubscriptionFunded event raised by the OracleProxy contract.
type OracleProxySubscriptionFunded struct {
	SubscriptionId uint64
	OldBalance     *big.Int
	NewBalance     *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSubscriptionFunded is a free log retrieval operation binding the contract event 0xd39ec07f4e209f627a4c427971473820dc129761ba28de8906bd56f57101d4f8.
//
// Solidity: event SubscriptionFunded(uint64 indexed subscriptionId, uint256 oldBalance, uint256 newBalance)
func (_OracleProxy *OracleProxyFilterer) FilterSubscriptionFunded(opts *bind.FilterOpts, subscriptionId []uint64) (*OracleProxySubscriptionFundedIterator, error) {

	var subscriptionIdRule []interface{}
	for _, subscriptionIdItem := range subscriptionId {
		subscriptionIdRule = append(subscriptionIdRule, subscriptionIdItem)
	}

	logs, sub, err := _OracleProxy.contract.FilterLogs(opts, "SubscriptionFunded", subscriptionIdRule)
	if err != nil {
		return nil, err
	}
	return &OracleProxySubscriptionFundedIterator{contract: _OracleProxy.contract, event: "SubscriptionFunded", logs: logs, sub: sub}, nil
}

// WatchSubscriptionFunded is a free log subscription operation binding the contract event 0xd39ec07f4e209f627a4c427971473820dc129761ba28de8906bd56f57101d4f8.
//
// Solidity: event SubscriptionFunded(uint64 indexed subscriptionId, uint256 oldBalance, uint256 newBalance)
func (_OracleProxy *OracleProxyFilterer) WatchSubscriptionFunded(opts *bind.WatchOpts, sink chan<- *OracleProxySubscriptionFunded, subscriptionId []uint64) (event.Subscription, error) {

	var subscriptionIdRule []interface{}
	for _, subscriptionIdItem := range subscriptionId {
		subscriptionIdRule = append(subscriptionIdRule, subscriptionIdItem)
	}

	logs, sub, err := _OracleProxy.contract.WatchLogs(opts, "SubscriptionFunded", subscriptionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleProxySubscriptionFunded)
				if err := _OracleProxy.contract.UnpackLog(event, "SubscriptionFunded", log); err != nil {
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

// ParseSubscriptionFunded is a log parse operation binding the contract event 0xd39ec07f4e209f627a4c427971473820dc129761ba28de8906bd56f57101d4f8.
//
// Solidity: event SubscriptionFunded(uint64 indexed subscriptionId, uint256 oldBalance, uint256 newBalance)
func (_OracleProxy *OracleProxyFilterer) ParseSubscriptionFunded(log types.Log) (*OracleProxySubscriptionFunded, error) {
	event := new(OracleProxySubscriptionFunded)
	if err := _OracleProxy.contract.UnpackLog(event, "SubscriptionFunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleProxySubscriptionOwnerTransferRequestedIterator is returned from FilterSubscriptionOwnerTransferRequested and is used to iterate over the raw logs and unpacked data for SubscriptionOwnerTransferRequested events raised by the OracleProxy contract.
type OracleProxySubscriptionOwnerTransferRequestedIterator struct {
	Event *OracleProxySubscriptionOwnerTransferRequested // Event containing the contract specifics and raw log

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
func (it *OracleProxySubscriptionOwnerTransferRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleProxySubscriptionOwnerTransferRequested)
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
		it.Event = new(OracleProxySubscriptionOwnerTransferRequested)
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
func (it *OracleProxySubscriptionOwnerTransferRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleProxySubscriptionOwnerTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleProxySubscriptionOwnerTransferRequested represents a SubscriptionOwnerTransferRequested event raised by the OracleProxy contract.
type OracleProxySubscriptionOwnerTransferRequested struct {
	SubscriptionId uint64
	From           common.Address
	To             common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSubscriptionOwnerTransferRequested is a free log retrieval operation binding the contract event 0x69436ea6df009049404f564eff6622cd00522b0bd6a89efd9e52a355c4a879be.
//
// Solidity: event SubscriptionOwnerTransferRequested(uint64 indexed subscriptionId, address from, address to)
func (_OracleProxy *OracleProxyFilterer) FilterSubscriptionOwnerTransferRequested(opts *bind.FilterOpts, subscriptionId []uint64) (*OracleProxySubscriptionOwnerTransferRequestedIterator, error) {

	var subscriptionIdRule []interface{}
	for _, subscriptionIdItem := range subscriptionId {
		subscriptionIdRule = append(subscriptionIdRule, subscriptionIdItem)
	}

	logs, sub, err := _OracleProxy.contract.FilterLogs(opts, "SubscriptionOwnerTransferRequested", subscriptionIdRule)
	if err != nil {
		return nil, err
	}
	return &OracleProxySubscriptionOwnerTransferRequestedIterator{contract: _OracleProxy.contract, event: "SubscriptionOwnerTransferRequested", logs: logs, sub: sub}, nil
}

// WatchSubscriptionOwnerTransferRequested is a free log subscription operation binding the contract event 0x69436ea6df009049404f564eff6622cd00522b0bd6a89efd9e52a355c4a879be.
//
// Solidity: event SubscriptionOwnerTransferRequested(uint64 indexed subscriptionId, address from, address to)
func (_OracleProxy *OracleProxyFilterer) WatchSubscriptionOwnerTransferRequested(opts *bind.WatchOpts, sink chan<- *OracleProxySubscriptionOwnerTransferRequested, subscriptionId []uint64) (event.Subscription, error) {

	var subscriptionIdRule []interface{}
	for _, subscriptionIdItem := range subscriptionId {
		subscriptionIdRule = append(subscriptionIdRule, subscriptionIdItem)
	}

	logs, sub, err := _OracleProxy.contract.WatchLogs(opts, "SubscriptionOwnerTransferRequested", subscriptionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleProxySubscriptionOwnerTransferRequested)
				if err := _OracleProxy.contract.UnpackLog(event, "SubscriptionOwnerTransferRequested", log); err != nil {
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

// ParseSubscriptionOwnerTransferRequested is a log parse operation binding the contract event 0x69436ea6df009049404f564eff6622cd00522b0bd6a89efd9e52a355c4a879be.
//
// Solidity: event SubscriptionOwnerTransferRequested(uint64 indexed subscriptionId, address from, address to)
func (_OracleProxy *OracleProxyFilterer) ParseSubscriptionOwnerTransferRequested(log types.Log) (*OracleProxySubscriptionOwnerTransferRequested, error) {
	event := new(OracleProxySubscriptionOwnerTransferRequested)
	if err := _OracleProxy.contract.UnpackLog(event, "SubscriptionOwnerTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleProxySubscriptionOwnerTransferredIterator is returned from FilterSubscriptionOwnerTransferred and is used to iterate over the raw logs and unpacked data for SubscriptionOwnerTransferred events raised by the OracleProxy contract.
type OracleProxySubscriptionOwnerTransferredIterator struct {
	Event *OracleProxySubscriptionOwnerTransferred // Event containing the contract specifics and raw log

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
func (it *OracleProxySubscriptionOwnerTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleProxySubscriptionOwnerTransferred)
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
		it.Event = new(OracleProxySubscriptionOwnerTransferred)
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
func (it *OracleProxySubscriptionOwnerTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleProxySubscriptionOwnerTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleProxySubscriptionOwnerTransferred represents a SubscriptionOwnerTransferred event raised by the OracleProxy contract.
type OracleProxySubscriptionOwnerTransferred struct {
	SubscriptionId uint64
	From           common.Address
	To             common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSubscriptionOwnerTransferred is a free log retrieval operation binding the contract event 0x6f1dc65165ffffedfd8e507b4a0f1fcfdada045ed11f6c26ba27cedfe87802f0.
//
// Solidity: event SubscriptionOwnerTransferred(uint64 indexed subscriptionId, address from, address to)
func (_OracleProxy *OracleProxyFilterer) FilterSubscriptionOwnerTransferred(opts *bind.FilterOpts, subscriptionId []uint64) (*OracleProxySubscriptionOwnerTransferredIterator, error) {

	var subscriptionIdRule []interface{}
	for _, subscriptionIdItem := range subscriptionId {
		subscriptionIdRule = append(subscriptionIdRule, subscriptionIdItem)
	}

	logs, sub, err := _OracleProxy.contract.FilterLogs(opts, "SubscriptionOwnerTransferred", subscriptionIdRule)
	if err != nil {
		return nil, err
	}
	return &OracleProxySubscriptionOwnerTransferredIterator{contract: _OracleProxy.contract, event: "SubscriptionOwnerTransferred", logs: logs, sub: sub}, nil
}

// WatchSubscriptionOwnerTransferred is a free log subscription operation binding the contract event 0x6f1dc65165ffffedfd8e507b4a0f1fcfdada045ed11f6c26ba27cedfe87802f0.
//
// Solidity: event SubscriptionOwnerTransferred(uint64 indexed subscriptionId, address from, address to)
func (_OracleProxy *OracleProxyFilterer) WatchSubscriptionOwnerTransferred(opts *bind.WatchOpts, sink chan<- *OracleProxySubscriptionOwnerTransferred, subscriptionId []uint64) (event.Subscription, error) {

	var subscriptionIdRule []interface{}
	for _, subscriptionIdItem := range subscriptionId {
		subscriptionIdRule = append(subscriptionIdRule, subscriptionIdItem)
	}

	logs, sub, err := _OracleProxy.contract.WatchLogs(opts, "SubscriptionOwnerTransferred", subscriptionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleProxySubscriptionOwnerTransferred)
				if err := _OracleProxy.contract.UnpackLog(event, "SubscriptionOwnerTransferred", log); err != nil {
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

// ParseSubscriptionOwnerTransferred is a log parse operation binding the contract event 0x6f1dc65165ffffedfd8e507b4a0f1fcfdada045ed11f6c26ba27cedfe87802f0.
//
// Solidity: event SubscriptionOwnerTransferred(uint64 indexed subscriptionId, address from, address to)
func (_OracleProxy *OracleProxyFilterer) ParseSubscriptionOwnerTransferred(log types.Log) (*OracleProxySubscriptionOwnerTransferred, error) {
	event := new(OracleProxySubscriptionOwnerTransferred)
	if err := _OracleProxy.contract.UnpackLog(event, "SubscriptionOwnerTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OracleProxyUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the OracleProxy contract.
type OracleProxyUnpausedIterator struct {
	Event *OracleProxyUnpaused // Event containing the contract specifics and raw log

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
func (it *OracleProxyUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OracleProxyUnpaused)
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
		it.Event = new(OracleProxyUnpaused)
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
func (it *OracleProxyUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OracleProxyUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OracleProxyUnpaused represents a Unpaused event raised by the OracleProxy contract.
type OracleProxyUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_OracleProxy *OracleProxyFilterer) FilterUnpaused(opts *bind.FilterOpts) (*OracleProxyUnpausedIterator, error) {

	logs, sub, err := _OracleProxy.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &OracleProxyUnpausedIterator{contract: _OracleProxy.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_OracleProxy *OracleProxyFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *OracleProxyUnpaused) (event.Subscription, error) {

	logs, sub, err := _OracleProxy.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OracleProxyUnpaused)
				if err := _OracleProxy.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_OracleProxy *OracleProxyFilterer) ParseUnpaused(log types.Log) (*OracleProxyUnpaused, error) {
	event := new(OracleProxyUnpaused)
	if err := _OracleProxy.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
