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

// BillingRegistryMetaData contains all meta data concerning the BillingRegistry contract.
var BillingRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"internalBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"externalBalance\",\"type\":\"uint256\"}],\"name\":\"BalanceInvariantViolated\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotSelfTransfer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptySendersList\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"have\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"want\",\"type\":\"uint32\"}],\"name\":\"GasLimitTooBig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectRequestID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCalldata\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"subscriptionId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"consumer\",\"type\":\"address\"}],\"name\":\"InvalidConsumer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"linkWei\",\"type\":\"int256\"}],\"name\":\"InvalidLinkWeiPrice\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSubscription\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proposedOwner\",\"type\":\"address\"}],\"name\":\"MustBeRequestedOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"MustBeSubOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAllowedToSetSenders\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotProposedOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableByOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableFromLink\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerMustBeSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PaymentTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingRequestExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Reentrant\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooManyConsumers\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnauthorizedSender\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"senders\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"changedBy\",\"type\":\"address\"}],\"name\":\"AuthorizedSendersChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"subscriptionId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"signerPayment\",\"type\":\"uint96\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"transmitterPayment\",\"type\":\"uint96\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"totalCost\",\"type\":\"uint96\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"name\":\"BillingEnd\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"subscriptionId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"client\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"gasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"don\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"donFee\",\"type\":\"uint96\"},{\"internalType\":\"uint96\",\"name\":\"registryFee\",\"type\":\"uint96\"},{\"internalType\":\"uint96\",\"name\":\"estimatedCost\",\"type\":\"uint96\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structFunctionsBillingRegistry.Commitment\",\"name\":\"commitment\",\"type\":\"tuple\"}],\"name\":\"BillingStart\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"maxGasLimit\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"stalenessSeconds\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasAfterPaymentCalculation\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"fallbackWeiPerUnitLink\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"gasOverhead\",\"type\":\"uint32\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FundsRecovered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"}],\"name\":\"RequestTimedOut\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"subscriptionId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SubscriptionCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"subscriptionId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"consumer\",\"type\":\"address\"}],\"name\":\"SubscriptionConsumerAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"subscriptionId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"consumer\",\"type\":\"address\"}],\"name\":\"SubscriptionConsumerRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"subscriptionId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"SubscriptionCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"subscriptionId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBalance\",\"type\":\"uint256\"}],\"name\":\"SubscriptionFunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"subscriptionId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"SubscriptionOwnerTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"subscriptionId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"SubscriptionOwnerTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_CONSUMERS\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"subscriptionId\",\"type\":\"uint64\"}],\"name\":\"acceptSubscriptionOwnerTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"subscriptionId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"consumer\",\"type\":\"address\"}],\"name\":\"addConsumer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"subscriptionId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"cancelSubscription\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"createSubscription\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"gasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint96\",\"name\":\"donFee\",\"type\":\"uint96\"},{\"internalType\":\"uint96\",\"name\":\"registryFee\",\"type\":\"uint96\"}],\"name\":\"estimateCost\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"response\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"err\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"transmitter\",\"type\":\"address\"},{\"internalType\":\"address[31]\",\"name\":\"signers\",\"type\":\"address[31]\"},{\"internalType\":\"uint8\",\"name\":\"signerCount\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"reportValidationGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initialGas\",\"type\":\"uint256\"}],\"name\":\"fulfillAndBill\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAuthorizedSenders\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConfig\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"maxGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"stalenessSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"gasAfterPaymentCalculation\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"fallbackWeiPerUnitLink\",\"type\":\"int256\"},{\"internalType\":\"uint32\",\"name\":\"gasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"linkAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"linkPriceFeed\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentsubscriptionId\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRequestConfig\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"subscriptionId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"client\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"gasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"}],\"internalType\":\"structFunctionsBillingRegistryInterface.RequestBilling\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"getRequiredFee\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"subscriptionId\",\"type\":\"uint64\"}],\"name\":\"getSubscription\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"balance\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"consumers\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"subscriptionId\",\"type\":\"uint64\"}],\"name\":\"getSubscriptionOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"link\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"linkEthFeed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"isAuthorizedSender\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onTokenTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"}],\"name\":\"oracleWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"subscriptionId\",\"type\":\"uint64\"}],\"name\":\"ownerCancelSubscription\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"subscriptionId\",\"type\":\"uint64\"}],\"name\":\"pendingRequestExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"recoverFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"subscriptionId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"consumer\",\"type\":\"address\"}],\"name\":\"removeConsumer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"subscriptionId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"requestSubscriptionOwnerTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"senders\",\"type\":\"address[]\"}],\"name\":\"setAuthorizedSenders\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"maxGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"stalenessSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"gasAfterPaymentCalculation\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"fallbackWeiPerUnitLink\",\"type\":\"int256\"},{\"internalType\":\"uint32\",\"name\":\"gasOverhead\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"requestTimeoutSeconds\",\"type\":\"uint32\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"subscriptionId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"client\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"gasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"}],\"internalType\":\"structFunctionsBillingRegistryInterface.RequestBilling\",\"name\":\"billing\",\"type\":\"tuple\"}],\"name\":\"startBilling\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"requestIdsToTimeout\",\"type\":\"bytes32[]\"}],\"name\":\"timeoutRequests\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BillingRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use BillingRegistryMetaData.ABI instead.
var BillingRegistryABI = BillingRegistryMetaData.ABI

// BillingRegistry is an auto generated Go binding around an Ethereum contract.
type BillingRegistry struct {
	BillingRegistryCaller     // Read-only binding to the contract
	BillingRegistryTransactor // Write-only binding to the contract
	BillingRegistryFilterer   // Log filterer for contract events
}

// BillingRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type BillingRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BillingRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BillingRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BillingRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BillingRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BillingRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BillingRegistrySession struct {
	Contract     *BillingRegistry  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BillingRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BillingRegistryCallerSession struct {
	Contract *BillingRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// BillingRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BillingRegistryTransactorSession struct {
	Contract     *BillingRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// BillingRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type BillingRegistryRaw struct {
	Contract *BillingRegistry // Generic contract binding to access the raw methods on
}

// BillingRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BillingRegistryCallerRaw struct {
	Contract *BillingRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// BillingRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BillingRegistryTransactorRaw struct {
	Contract *BillingRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBillingRegistry creates a new instance of BillingRegistry, bound to a specific deployed contract.
func NewBillingRegistry(address common.Address, backend bind.ContractBackend) (*BillingRegistry, error) {
	contract, err := bindBillingRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BillingRegistry{BillingRegistryCaller: BillingRegistryCaller{contract: contract}, BillingRegistryTransactor: BillingRegistryTransactor{contract: contract}, BillingRegistryFilterer: BillingRegistryFilterer{contract: contract}}, nil
}

// NewBillingRegistryCaller creates a new read-only instance of BillingRegistry, bound to a specific deployed contract.
func NewBillingRegistryCaller(address common.Address, caller bind.ContractCaller) (*BillingRegistryCaller, error) {
	contract, err := bindBillingRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BillingRegistryCaller{contract: contract}, nil
}

// NewBillingRegistryTransactor creates a new write-only instance of BillingRegistry, bound to a specific deployed contract.
func NewBillingRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*BillingRegistryTransactor, error) {
	contract, err := bindBillingRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BillingRegistryTransactor{contract: contract}, nil
}

// NewBillingRegistryFilterer creates a new log filterer instance of BillingRegistry, bound to a specific deployed contract.
func NewBillingRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*BillingRegistryFilterer, error) {
	contract, err := bindBillingRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BillingRegistryFilterer{contract: contract}, nil
}

// bindBillingRegistry binds a generic wrapper to an already deployed contract.
func bindBillingRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BillingRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BillingRegistry *BillingRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BillingRegistry.Contract.BillingRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BillingRegistry *BillingRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BillingRegistry.Contract.BillingRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BillingRegistry *BillingRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BillingRegistry.Contract.BillingRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BillingRegistry *BillingRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BillingRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BillingRegistry *BillingRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BillingRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BillingRegistry *BillingRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BillingRegistry.Contract.contract.Transact(opts, method, params...)
}

// MAXCONSUMERS is a free data retrieval call binding the contract method 0x64d51a2a.
//
// Solidity: function MAX_CONSUMERS() view returns(uint16)
func (_BillingRegistry *BillingRegistryCaller) MAXCONSUMERS(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _BillingRegistry.contract.Call(opts, &out, "MAX_CONSUMERS")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MAXCONSUMERS is a free data retrieval call binding the contract method 0x64d51a2a.
//
// Solidity: function MAX_CONSUMERS() view returns(uint16)
func (_BillingRegistry *BillingRegistrySession) MAXCONSUMERS() (uint16, error) {
	return _BillingRegistry.Contract.MAXCONSUMERS(&_BillingRegistry.CallOpts)
}

// MAXCONSUMERS is a free data retrieval call binding the contract method 0x64d51a2a.
//
// Solidity: function MAX_CONSUMERS() view returns(uint16)
func (_BillingRegistry *BillingRegistryCallerSession) MAXCONSUMERS() (uint16, error) {
	return _BillingRegistry.Contract.MAXCONSUMERS(&_BillingRegistry.CallOpts)
}

// EstimateCost is a free data retrieval call binding the contract method 0xa1a6d041.
//
// Solidity: function estimateCost(uint32 gasLimit, uint256 gasPrice, uint96 donFee, uint96 registryFee) view returns(uint96)
func (_BillingRegistry *BillingRegistryCaller) EstimateCost(opts *bind.CallOpts, gasLimit uint32, gasPrice *big.Int, donFee *big.Int, registryFee *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BillingRegistry.contract.Call(opts, &out, "estimateCost", gasLimit, gasPrice, donFee, registryFee)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimateCost is a free data retrieval call binding the contract method 0xa1a6d041.
//
// Solidity: function estimateCost(uint32 gasLimit, uint256 gasPrice, uint96 donFee, uint96 registryFee) view returns(uint96)
func (_BillingRegistry *BillingRegistrySession) EstimateCost(gasLimit uint32, gasPrice *big.Int, donFee *big.Int, registryFee *big.Int) (*big.Int, error) {
	return _BillingRegistry.Contract.EstimateCost(&_BillingRegistry.CallOpts, gasLimit, gasPrice, donFee, registryFee)
}

// EstimateCost is a free data retrieval call binding the contract method 0xa1a6d041.
//
// Solidity: function estimateCost(uint32 gasLimit, uint256 gasPrice, uint96 donFee, uint96 registryFee) view returns(uint96)
func (_BillingRegistry *BillingRegistryCallerSession) EstimateCost(gasLimit uint32, gasPrice *big.Int, donFee *big.Int, registryFee *big.Int) (*big.Int, error) {
	return _BillingRegistry.Contract.EstimateCost(&_BillingRegistry.CallOpts, gasLimit, gasPrice, donFee, registryFee)
}

// GetAuthorizedSenders is a free data retrieval call binding the contract method 0x2408afaa.
//
// Solidity: function getAuthorizedSenders() view returns(address[])
func (_BillingRegistry *BillingRegistryCaller) GetAuthorizedSenders(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _BillingRegistry.contract.Call(opts, &out, "getAuthorizedSenders")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAuthorizedSenders is a free data retrieval call binding the contract method 0x2408afaa.
//
// Solidity: function getAuthorizedSenders() view returns(address[])
func (_BillingRegistry *BillingRegistrySession) GetAuthorizedSenders() ([]common.Address, error) {
	return _BillingRegistry.Contract.GetAuthorizedSenders(&_BillingRegistry.CallOpts)
}

// GetAuthorizedSenders is a free data retrieval call binding the contract method 0x2408afaa.
//
// Solidity: function getAuthorizedSenders() view returns(address[])
func (_BillingRegistry *BillingRegistryCallerSession) GetAuthorizedSenders() ([]common.Address, error) {
	return _BillingRegistry.Contract.GetAuthorizedSenders(&_BillingRegistry.CallOpts)
}

// GetConfig is a free data retrieval call binding the contract method 0xc3f909d4.
//
// Solidity: function getConfig() view returns(uint32 maxGasLimit, uint32 stalenessSeconds, uint256 gasAfterPaymentCalculation, int256 fallbackWeiPerUnitLink, uint32 gasOverhead, address linkAddress, address linkPriceFeed)
func (_BillingRegistry *BillingRegistryCaller) GetConfig(opts *bind.CallOpts) (struct {
	MaxGasLimit                uint32
	StalenessSeconds           uint32
	GasAfterPaymentCalculation *big.Int
	FallbackWeiPerUnitLink     *big.Int
	GasOverhead                uint32
	LinkAddress                common.Address
	LinkPriceFeed              common.Address
}, error) {
	var out []interface{}
	err := _BillingRegistry.contract.Call(opts, &out, "getConfig")

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
func (_BillingRegistry *BillingRegistrySession) GetConfig() (struct {
	MaxGasLimit                uint32
	StalenessSeconds           uint32
	GasAfterPaymentCalculation *big.Int
	FallbackWeiPerUnitLink     *big.Int
	GasOverhead                uint32
	LinkAddress                common.Address
	LinkPriceFeed              common.Address
}, error) {
	return _BillingRegistry.Contract.GetConfig(&_BillingRegistry.CallOpts)
}

// GetConfig is a free data retrieval call binding the contract method 0xc3f909d4.
//
// Solidity: function getConfig() view returns(uint32 maxGasLimit, uint32 stalenessSeconds, uint256 gasAfterPaymentCalculation, int256 fallbackWeiPerUnitLink, uint32 gasOverhead, address linkAddress, address linkPriceFeed)
func (_BillingRegistry *BillingRegistryCallerSession) GetConfig() (struct {
	MaxGasLimit                uint32
	StalenessSeconds           uint32
	GasAfterPaymentCalculation *big.Int
	FallbackWeiPerUnitLink     *big.Int
	GasOverhead                uint32
	LinkAddress                common.Address
	LinkPriceFeed              common.Address
}, error) {
	return _BillingRegistry.Contract.GetConfig(&_BillingRegistry.CallOpts)
}

// GetCurrentsubscriptionId is a free data retrieval call binding the contract method 0x33652e3e.
//
// Solidity: function getCurrentsubscriptionId() view returns(uint64)
func (_BillingRegistry *BillingRegistryCaller) GetCurrentsubscriptionId(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _BillingRegistry.contract.Call(opts, &out, "getCurrentsubscriptionId")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetCurrentsubscriptionId is a free data retrieval call binding the contract method 0x33652e3e.
//
// Solidity: function getCurrentsubscriptionId() view returns(uint64)
func (_BillingRegistry *BillingRegistrySession) GetCurrentsubscriptionId() (uint64, error) {
	return _BillingRegistry.Contract.GetCurrentsubscriptionId(&_BillingRegistry.CallOpts)
}

// GetCurrentsubscriptionId is a free data retrieval call binding the contract method 0x33652e3e.
//
// Solidity: function getCurrentsubscriptionId() view returns(uint64)
func (_BillingRegistry *BillingRegistryCallerSession) GetCurrentsubscriptionId() (uint64, error) {
	return _BillingRegistry.Contract.GetCurrentsubscriptionId(&_BillingRegistry.CallOpts)
}

// GetRequestConfig is a free data retrieval call binding the contract method 0x00012291.
//
// Solidity: function getRequestConfig() view returns(uint32, address[])
func (_BillingRegistry *BillingRegistryCaller) GetRequestConfig(opts *bind.CallOpts) (uint32, []common.Address, error) {
	var out []interface{}
	err := _BillingRegistry.contract.Call(opts, &out, "getRequestConfig")

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
func (_BillingRegistry *BillingRegistrySession) GetRequestConfig() (uint32, []common.Address, error) {
	return _BillingRegistry.Contract.GetRequestConfig(&_BillingRegistry.CallOpts)
}

// GetRequestConfig is a free data retrieval call binding the contract method 0x00012291.
//
// Solidity: function getRequestConfig() view returns(uint32, address[])
func (_BillingRegistry *BillingRegistryCallerSession) GetRequestConfig() (uint32, []common.Address, error) {
	return _BillingRegistry.Contract.GetRequestConfig(&_BillingRegistry.CallOpts)
}

// GetRequiredFee is a free data retrieval call binding the contract method 0xf1e14a21.
//
// Solidity: function getRequiredFee(bytes , (uint64,address,uint32,uint256) ) pure returns(uint96)
func (_BillingRegistry *BillingRegistryCaller) GetRequiredFee(opts *bind.CallOpts, arg0 []byte, arg1 FunctionsBillingRegistryInterfaceRequestBilling) (*big.Int, error) {
	var out []interface{}
	err := _BillingRegistry.contract.Call(opts, &out, "getRequiredFee", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRequiredFee is a free data retrieval call binding the contract method 0xf1e14a21.
//
// Solidity: function getRequiredFee(bytes , (uint64,address,uint32,uint256) ) pure returns(uint96)
func (_BillingRegistry *BillingRegistrySession) GetRequiredFee(arg0 []byte, arg1 FunctionsBillingRegistryInterfaceRequestBilling) (*big.Int, error) {
	return _BillingRegistry.Contract.GetRequiredFee(&_BillingRegistry.CallOpts, arg0, arg1)
}

// GetRequiredFee is a free data retrieval call binding the contract method 0xf1e14a21.
//
// Solidity: function getRequiredFee(bytes , (uint64,address,uint32,uint256) ) pure returns(uint96)
func (_BillingRegistry *BillingRegistryCallerSession) GetRequiredFee(arg0 []byte, arg1 FunctionsBillingRegistryInterfaceRequestBilling) (*big.Int, error) {
	return _BillingRegistry.Contract.GetRequiredFee(&_BillingRegistry.CallOpts, arg0, arg1)
}

// GetSubscription is a free data retrieval call binding the contract method 0xa47c7696.
//
// Solidity: function getSubscription(uint64 subscriptionId) view returns(uint96 balance, address owner, address[] consumers)
func (_BillingRegistry *BillingRegistryCaller) GetSubscription(opts *bind.CallOpts, subscriptionId uint64) (struct {
	Balance   *big.Int
	Owner     common.Address
	Consumers []common.Address
}, error) {
	var out []interface{}
	err := _BillingRegistry.contract.Call(opts, &out, "getSubscription", subscriptionId)

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
func (_BillingRegistry *BillingRegistrySession) GetSubscription(subscriptionId uint64) (struct {
	Balance   *big.Int
	Owner     common.Address
	Consumers []common.Address
}, error) {
	return _BillingRegistry.Contract.GetSubscription(&_BillingRegistry.CallOpts, subscriptionId)
}

// GetSubscription is a free data retrieval call binding the contract method 0xa47c7696.
//
// Solidity: function getSubscription(uint64 subscriptionId) view returns(uint96 balance, address owner, address[] consumers)
func (_BillingRegistry *BillingRegistryCallerSession) GetSubscription(subscriptionId uint64) (struct {
	Balance   *big.Int
	Owner     common.Address
	Consumers []common.Address
}, error) {
	return _BillingRegistry.Contract.GetSubscription(&_BillingRegistry.CallOpts, subscriptionId)
}

// GetSubscriptionOwner is a free data retrieval call binding the contract method 0xb2a489ff.
//
// Solidity: function getSubscriptionOwner(uint64 subscriptionId) view returns(address owner)
func (_BillingRegistry *BillingRegistryCaller) GetSubscriptionOwner(opts *bind.CallOpts, subscriptionId uint64) (common.Address, error) {
	var out []interface{}
	err := _BillingRegistry.contract.Call(opts, &out, "getSubscriptionOwner", subscriptionId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSubscriptionOwner is a free data retrieval call binding the contract method 0xb2a489ff.
//
// Solidity: function getSubscriptionOwner(uint64 subscriptionId) view returns(address owner)
func (_BillingRegistry *BillingRegistrySession) GetSubscriptionOwner(subscriptionId uint64) (common.Address, error) {
	return _BillingRegistry.Contract.GetSubscriptionOwner(&_BillingRegistry.CallOpts, subscriptionId)
}

// GetSubscriptionOwner is a free data retrieval call binding the contract method 0xb2a489ff.
//
// Solidity: function getSubscriptionOwner(uint64 subscriptionId) view returns(address owner)
func (_BillingRegistry *BillingRegistryCallerSession) GetSubscriptionOwner(subscriptionId uint64) (common.Address, error) {
	return _BillingRegistry.Contract.GetSubscriptionOwner(&_BillingRegistry.CallOpts, subscriptionId)
}

// GetTotalBalance is a free data retrieval call binding the contract method 0x12b58349.
//
// Solidity: function getTotalBalance() view returns(uint256)
func (_BillingRegistry *BillingRegistryCaller) GetTotalBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BillingRegistry.contract.Call(opts, &out, "getTotalBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalBalance is a free data retrieval call binding the contract method 0x12b58349.
//
// Solidity: function getTotalBalance() view returns(uint256)
func (_BillingRegistry *BillingRegistrySession) GetTotalBalance() (*big.Int, error) {
	return _BillingRegistry.Contract.GetTotalBalance(&_BillingRegistry.CallOpts)
}

// GetTotalBalance is a free data retrieval call binding the contract method 0x12b58349.
//
// Solidity: function getTotalBalance() view returns(uint256)
func (_BillingRegistry *BillingRegistryCallerSession) GetTotalBalance() (*big.Int, error) {
	return _BillingRegistry.Contract.GetTotalBalance(&_BillingRegistry.CallOpts)
}

// IsAuthorizedSender is a free data retrieval call binding the contract method 0xfa00763a.
//
// Solidity: function isAuthorizedSender(address sender) view returns(bool)
func (_BillingRegistry *BillingRegistryCaller) IsAuthorizedSender(opts *bind.CallOpts, sender common.Address) (bool, error) {
	var out []interface{}
	err := _BillingRegistry.contract.Call(opts, &out, "isAuthorizedSender", sender)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAuthorizedSender is a free data retrieval call binding the contract method 0xfa00763a.
//
// Solidity: function isAuthorizedSender(address sender) view returns(bool)
func (_BillingRegistry *BillingRegistrySession) IsAuthorizedSender(sender common.Address) (bool, error) {
	return _BillingRegistry.Contract.IsAuthorizedSender(&_BillingRegistry.CallOpts, sender)
}

// IsAuthorizedSender is a free data retrieval call binding the contract method 0xfa00763a.
//
// Solidity: function isAuthorizedSender(address sender) view returns(bool)
func (_BillingRegistry *BillingRegistryCallerSession) IsAuthorizedSender(sender common.Address) (bool, error) {
	return _BillingRegistry.Contract.IsAuthorizedSender(&_BillingRegistry.CallOpts, sender)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BillingRegistry *BillingRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BillingRegistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BillingRegistry *BillingRegistrySession) Owner() (common.Address, error) {
	return _BillingRegistry.Contract.Owner(&_BillingRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BillingRegistry *BillingRegistryCallerSession) Owner() (common.Address, error) {
	return _BillingRegistry.Contract.Owner(&_BillingRegistry.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BillingRegistry *BillingRegistryCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BillingRegistry.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BillingRegistry *BillingRegistrySession) Paused() (bool, error) {
	return _BillingRegistry.Contract.Paused(&_BillingRegistry.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BillingRegistry *BillingRegistryCallerSession) Paused() (bool, error) {
	return _BillingRegistry.Contract.Paused(&_BillingRegistry.CallOpts)
}

// PendingRequestExists is a free data retrieval call binding the contract method 0xe82ad7d4.
//
// Solidity: function pendingRequestExists(uint64 subscriptionId) view returns(bool)
func (_BillingRegistry *BillingRegistryCaller) PendingRequestExists(opts *bind.CallOpts, subscriptionId uint64) (bool, error) {
	var out []interface{}
	err := _BillingRegistry.contract.Call(opts, &out, "pendingRequestExists", subscriptionId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PendingRequestExists is a free data retrieval call binding the contract method 0xe82ad7d4.
//
// Solidity: function pendingRequestExists(uint64 subscriptionId) view returns(bool)
func (_BillingRegistry *BillingRegistrySession) PendingRequestExists(subscriptionId uint64) (bool, error) {
	return _BillingRegistry.Contract.PendingRequestExists(&_BillingRegistry.CallOpts, subscriptionId)
}

// PendingRequestExists is a free data retrieval call binding the contract method 0xe82ad7d4.
//
// Solidity: function pendingRequestExists(uint64 subscriptionId) view returns(bool)
func (_BillingRegistry *BillingRegistryCallerSession) PendingRequestExists(subscriptionId uint64) (bool, error) {
	return _BillingRegistry.Contract.PendingRequestExists(&_BillingRegistry.CallOpts, subscriptionId)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_BillingRegistry *BillingRegistryTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BillingRegistry.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_BillingRegistry *BillingRegistrySession) AcceptOwnership() (*types.Transaction, error) {
	return _BillingRegistry.Contract.AcceptOwnership(&_BillingRegistry.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_BillingRegistry *BillingRegistryTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _BillingRegistry.Contract.AcceptOwnership(&_BillingRegistry.TransactOpts)
}

// AcceptSubscriptionOwnerTransfer is a paid mutator transaction binding the contract method 0x82359740.
//
// Solidity: function acceptSubscriptionOwnerTransfer(uint64 subscriptionId) returns()
func (_BillingRegistry *BillingRegistryTransactor) AcceptSubscriptionOwnerTransfer(opts *bind.TransactOpts, subscriptionId uint64) (*types.Transaction, error) {
	return _BillingRegistry.contract.Transact(opts, "acceptSubscriptionOwnerTransfer", subscriptionId)
}

// AcceptSubscriptionOwnerTransfer is a paid mutator transaction binding the contract method 0x82359740.
//
// Solidity: function acceptSubscriptionOwnerTransfer(uint64 subscriptionId) returns()
func (_BillingRegistry *BillingRegistrySession) AcceptSubscriptionOwnerTransfer(subscriptionId uint64) (*types.Transaction, error) {
	return _BillingRegistry.Contract.AcceptSubscriptionOwnerTransfer(&_BillingRegistry.TransactOpts, subscriptionId)
}

// AcceptSubscriptionOwnerTransfer is a paid mutator transaction binding the contract method 0x82359740.
//
// Solidity: function acceptSubscriptionOwnerTransfer(uint64 subscriptionId) returns()
func (_BillingRegistry *BillingRegistryTransactorSession) AcceptSubscriptionOwnerTransfer(subscriptionId uint64) (*types.Transaction, error) {
	return _BillingRegistry.Contract.AcceptSubscriptionOwnerTransfer(&_BillingRegistry.TransactOpts, subscriptionId)
}

// AddConsumer is a paid mutator transaction binding the contract method 0x7341c10c.
//
// Solidity: function addConsumer(uint64 subscriptionId, address consumer) returns()
func (_BillingRegistry *BillingRegistryTransactor) AddConsumer(opts *bind.TransactOpts, subscriptionId uint64, consumer common.Address) (*types.Transaction, error) {
	return _BillingRegistry.contract.Transact(opts, "addConsumer", subscriptionId, consumer)
}

// AddConsumer is a paid mutator transaction binding the contract method 0x7341c10c.
//
// Solidity: function addConsumer(uint64 subscriptionId, address consumer) returns()
func (_BillingRegistry *BillingRegistrySession) AddConsumer(subscriptionId uint64, consumer common.Address) (*types.Transaction, error) {
	return _BillingRegistry.Contract.AddConsumer(&_BillingRegistry.TransactOpts, subscriptionId, consumer)
}

// AddConsumer is a paid mutator transaction binding the contract method 0x7341c10c.
//
// Solidity: function addConsumer(uint64 subscriptionId, address consumer) returns()
func (_BillingRegistry *BillingRegistryTransactorSession) AddConsumer(subscriptionId uint64, consumer common.Address) (*types.Transaction, error) {
	return _BillingRegistry.Contract.AddConsumer(&_BillingRegistry.TransactOpts, subscriptionId, consumer)
}

// CancelSubscription is a paid mutator transaction binding the contract method 0xd7ae1d30.
//
// Solidity: function cancelSubscription(uint64 subscriptionId, address to) returns()
func (_BillingRegistry *BillingRegistryTransactor) CancelSubscription(opts *bind.TransactOpts, subscriptionId uint64, to common.Address) (*types.Transaction, error) {
	return _BillingRegistry.contract.Transact(opts, "cancelSubscription", subscriptionId, to)
}

// CancelSubscription is a paid mutator transaction binding the contract method 0xd7ae1d30.
//
// Solidity: function cancelSubscription(uint64 subscriptionId, address to) returns()
func (_BillingRegistry *BillingRegistrySession) CancelSubscription(subscriptionId uint64, to common.Address) (*types.Transaction, error) {
	return _BillingRegistry.Contract.CancelSubscription(&_BillingRegistry.TransactOpts, subscriptionId, to)
}

// CancelSubscription is a paid mutator transaction binding the contract method 0xd7ae1d30.
//
// Solidity: function cancelSubscription(uint64 subscriptionId, address to) returns()
func (_BillingRegistry *BillingRegistryTransactorSession) CancelSubscription(subscriptionId uint64, to common.Address) (*types.Transaction, error) {
	return _BillingRegistry.Contract.CancelSubscription(&_BillingRegistry.TransactOpts, subscriptionId, to)
}

// CreateSubscription is a paid mutator transaction binding the contract method 0xa21a23e4.
//
// Solidity: function createSubscription() returns(uint64)
func (_BillingRegistry *BillingRegistryTransactor) CreateSubscription(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BillingRegistry.contract.Transact(opts, "createSubscription")
}

// CreateSubscription is a paid mutator transaction binding the contract method 0xa21a23e4.
//
// Solidity: function createSubscription() returns(uint64)
func (_BillingRegistry *BillingRegistrySession) CreateSubscription() (*types.Transaction, error) {
	return _BillingRegistry.Contract.CreateSubscription(&_BillingRegistry.TransactOpts)
}

// CreateSubscription is a paid mutator transaction binding the contract method 0xa21a23e4.
//
// Solidity: function createSubscription() returns(uint64)
func (_BillingRegistry *BillingRegistryTransactorSession) CreateSubscription() (*types.Transaction, error) {
	return _BillingRegistry.Contract.CreateSubscription(&_BillingRegistry.TransactOpts)
}

// FulfillAndBill is a paid mutator transaction binding the contract method 0x0739e4f1.
//
// Solidity: function fulfillAndBill(bytes32 requestId, bytes response, bytes err, address transmitter, address[31] signers, uint8 signerCount, uint256 reportValidationGas, uint256 initialGas) returns(bool success)
func (_BillingRegistry *BillingRegistryTransactor) FulfillAndBill(opts *bind.TransactOpts, requestId [32]byte, response []byte, err []byte, transmitter common.Address, signers [31]common.Address, signerCount uint8, reportValidationGas *big.Int, initialGas *big.Int) (*types.Transaction, error) {
	return _BillingRegistry.contract.Transact(opts, "fulfillAndBill", requestId, response, err, transmitter, signers, signerCount, reportValidationGas, initialGas)
}

// FulfillAndBill is a paid mutator transaction binding the contract method 0x0739e4f1.
//
// Solidity: function fulfillAndBill(bytes32 requestId, bytes response, bytes err, address transmitter, address[31] signers, uint8 signerCount, uint256 reportValidationGas, uint256 initialGas) returns(bool success)
func (_BillingRegistry *BillingRegistrySession) FulfillAndBill(requestId [32]byte, response []byte, err []byte, transmitter common.Address, signers [31]common.Address, signerCount uint8, reportValidationGas *big.Int, initialGas *big.Int) (*types.Transaction, error) {
	return _BillingRegistry.Contract.FulfillAndBill(&_BillingRegistry.TransactOpts, requestId, response, err, transmitter, signers, signerCount, reportValidationGas, initialGas)
}

// FulfillAndBill is a paid mutator transaction binding the contract method 0x0739e4f1.
//
// Solidity: function fulfillAndBill(bytes32 requestId, bytes response, bytes err, address transmitter, address[31] signers, uint8 signerCount, uint256 reportValidationGas, uint256 initialGas) returns(bool success)
func (_BillingRegistry *BillingRegistryTransactorSession) FulfillAndBill(requestId [32]byte, response []byte, err []byte, transmitter common.Address, signers [31]common.Address, signerCount uint8, reportValidationGas *big.Int, initialGas *big.Int) (*types.Transaction, error) {
	return _BillingRegistry.Contract.FulfillAndBill(&_BillingRegistry.TransactOpts, requestId, response, err, transmitter, signers, signerCount, reportValidationGas, initialGas)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address link, address linkEthFeed, address oracle) returns()
func (_BillingRegistry *BillingRegistryTransactor) Initialize(opts *bind.TransactOpts, link common.Address, linkEthFeed common.Address, oracle common.Address) (*types.Transaction, error) {
	return _BillingRegistry.contract.Transact(opts, "initialize", link, linkEthFeed, oracle)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address link, address linkEthFeed, address oracle) returns()
func (_BillingRegistry *BillingRegistrySession) Initialize(link common.Address, linkEthFeed common.Address, oracle common.Address) (*types.Transaction, error) {
	return _BillingRegistry.Contract.Initialize(&_BillingRegistry.TransactOpts, link, linkEthFeed, oracle)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address link, address linkEthFeed, address oracle) returns()
func (_BillingRegistry *BillingRegistryTransactorSession) Initialize(link common.Address, linkEthFeed common.Address, oracle common.Address) (*types.Transaction, error) {
	return _BillingRegistry.Contract.Initialize(&_BillingRegistry.TransactOpts, link, linkEthFeed, oracle)
}

// OnTokenTransfer is a paid mutator transaction binding the contract method 0xa4c0ed36.
//
// Solidity: function onTokenTransfer(address , uint256 amount, bytes data) returns()
func (_BillingRegistry *BillingRegistryTransactor) OnTokenTransfer(opts *bind.TransactOpts, arg0 common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _BillingRegistry.contract.Transact(opts, "onTokenTransfer", arg0, amount, data)
}

// OnTokenTransfer is a paid mutator transaction binding the contract method 0xa4c0ed36.
//
// Solidity: function onTokenTransfer(address , uint256 amount, bytes data) returns()
func (_BillingRegistry *BillingRegistrySession) OnTokenTransfer(arg0 common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _BillingRegistry.Contract.OnTokenTransfer(&_BillingRegistry.TransactOpts, arg0, amount, data)
}

// OnTokenTransfer is a paid mutator transaction binding the contract method 0xa4c0ed36.
//
// Solidity: function onTokenTransfer(address , uint256 amount, bytes data) returns()
func (_BillingRegistry *BillingRegistryTransactorSession) OnTokenTransfer(arg0 common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _BillingRegistry.Contract.OnTokenTransfer(&_BillingRegistry.TransactOpts, arg0, amount, data)
}

// OracleWithdraw is a paid mutator transaction binding the contract method 0x66316d8d.
//
// Solidity: function oracleWithdraw(address recipient, uint96 amount) returns()
func (_BillingRegistry *BillingRegistryTransactor) OracleWithdraw(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BillingRegistry.contract.Transact(opts, "oracleWithdraw", recipient, amount)
}

// OracleWithdraw is a paid mutator transaction binding the contract method 0x66316d8d.
//
// Solidity: function oracleWithdraw(address recipient, uint96 amount) returns()
func (_BillingRegistry *BillingRegistrySession) OracleWithdraw(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BillingRegistry.Contract.OracleWithdraw(&_BillingRegistry.TransactOpts, recipient, amount)
}

// OracleWithdraw is a paid mutator transaction binding the contract method 0x66316d8d.
//
// Solidity: function oracleWithdraw(address recipient, uint96 amount) returns()
func (_BillingRegistry *BillingRegistryTransactorSession) OracleWithdraw(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BillingRegistry.Contract.OracleWithdraw(&_BillingRegistry.TransactOpts, recipient, amount)
}

// OwnerCancelSubscription is a paid mutator transaction binding the contract method 0x02bcc5b6.
//
// Solidity: function ownerCancelSubscription(uint64 subscriptionId) returns()
func (_BillingRegistry *BillingRegistryTransactor) OwnerCancelSubscription(opts *bind.TransactOpts, subscriptionId uint64) (*types.Transaction, error) {
	return _BillingRegistry.contract.Transact(opts, "ownerCancelSubscription", subscriptionId)
}

// OwnerCancelSubscription is a paid mutator transaction binding the contract method 0x02bcc5b6.
//
// Solidity: function ownerCancelSubscription(uint64 subscriptionId) returns()
func (_BillingRegistry *BillingRegistrySession) OwnerCancelSubscription(subscriptionId uint64) (*types.Transaction, error) {
	return _BillingRegistry.Contract.OwnerCancelSubscription(&_BillingRegistry.TransactOpts, subscriptionId)
}

// OwnerCancelSubscription is a paid mutator transaction binding the contract method 0x02bcc5b6.
//
// Solidity: function ownerCancelSubscription(uint64 subscriptionId) returns()
func (_BillingRegistry *BillingRegistryTransactorSession) OwnerCancelSubscription(subscriptionId uint64) (*types.Transaction, error) {
	return _BillingRegistry.Contract.OwnerCancelSubscription(&_BillingRegistry.TransactOpts, subscriptionId)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BillingRegistry *BillingRegistryTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BillingRegistry.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BillingRegistry *BillingRegistrySession) Pause() (*types.Transaction, error) {
	return _BillingRegistry.Contract.Pause(&_BillingRegistry.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BillingRegistry *BillingRegistryTransactorSession) Pause() (*types.Transaction, error) {
	return _BillingRegistry.Contract.Pause(&_BillingRegistry.TransactOpts)
}

// RecoverFunds is a paid mutator transaction binding the contract method 0xe72f6e30.
//
// Solidity: function recoverFunds(address to) returns()
func (_BillingRegistry *BillingRegistryTransactor) RecoverFunds(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _BillingRegistry.contract.Transact(opts, "recoverFunds", to)
}

// RecoverFunds is a paid mutator transaction binding the contract method 0xe72f6e30.
//
// Solidity: function recoverFunds(address to) returns()
func (_BillingRegistry *BillingRegistrySession) RecoverFunds(to common.Address) (*types.Transaction, error) {
	return _BillingRegistry.Contract.RecoverFunds(&_BillingRegistry.TransactOpts, to)
}

// RecoverFunds is a paid mutator transaction binding the contract method 0xe72f6e30.
//
// Solidity: function recoverFunds(address to) returns()
func (_BillingRegistry *BillingRegistryTransactorSession) RecoverFunds(to common.Address) (*types.Transaction, error) {
	return _BillingRegistry.Contract.RecoverFunds(&_BillingRegistry.TransactOpts, to)
}

// RemoveConsumer is a paid mutator transaction binding the contract method 0x9f87fad7.
//
// Solidity: function removeConsumer(uint64 subscriptionId, address consumer) returns()
func (_BillingRegistry *BillingRegistryTransactor) RemoveConsumer(opts *bind.TransactOpts, subscriptionId uint64, consumer common.Address) (*types.Transaction, error) {
	return _BillingRegistry.contract.Transact(opts, "removeConsumer", subscriptionId, consumer)
}

// RemoveConsumer is a paid mutator transaction binding the contract method 0x9f87fad7.
//
// Solidity: function removeConsumer(uint64 subscriptionId, address consumer) returns()
func (_BillingRegistry *BillingRegistrySession) RemoveConsumer(subscriptionId uint64, consumer common.Address) (*types.Transaction, error) {
	return _BillingRegistry.Contract.RemoveConsumer(&_BillingRegistry.TransactOpts, subscriptionId, consumer)
}

// RemoveConsumer is a paid mutator transaction binding the contract method 0x9f87fad7.
//
// Solidity: function removeConsumer(uint64 subscriptionId, address consumer) returns()
func (_BillingRegistry *BillingRegistryTransactorSession) RemoveConsumer(subscriptionId uint64, consumer common.Address) (*types.Transaction, error) {
	return _BillingRegistry.Contract.RemoveConsumer(&_BillingRegistry.TransactOpts, subscriptionId, consumer)
}

// RequestSubscriptionOwnerTransfer is a paid mutator transaction binding the contract method 0x04c357cb.
//
// Solidity: function requestSubscriptionOwnerTransfer(uint64 subscriptionId, address newOwner) returns()
func (_BillingRegistry *BillingRegistryTransactor) RequestSubscriptionOwnerTransfer(opts *bind.TransactOpts, subscriptionId uint64, newOwner common.Address) (*types.Transaction, error) {
	return _BillingRegistry.contract.Transact(opts, "requestSubscriptionOwnerTransfer", subscriptionId, newOwner)
}

// RequestSubscriptionOwnerTransfer is a paid mutator transaction binding the contract method 0x04c357cb.
//
// Solidity: function requestSubscriptionOwnerTransfer(uint64 subscriptionId, address newOwner) returns()
func (_BillingRegistry *BillingRegistrySession) RequestSubscriptionOwnerTransfer(subscriptionId uint64, newOwner common.Address) (*types.Transaction, error) {
	return _BillingRegistry.Contract.RequestSubscriptionOwnerTransfer(&_BillingRegistry.TransactOpts, subscriptionId, newOwner)
}

// RequestSubscriptionOwnerTransfer is a paid mutator transaction binding the contract method 0x04c357cb.
//
// Solidity: function requestSubscriptionOwnerTransfer(uint64 subscriptionId, address newOwner) returns()
func (_BillingRegistry *BillingRegistryTransactorSession) RequestSubscriptionOwnerTransfer(subscriptionId uint64, newOwner common.Address) (*types.Transaction, error) {
	return _BillingRegistry.Contract.RequestSubscriptionOwnerTransfer(&_BillingRegistry.TransactOpts, subscriptionId, newOwner)
}

// SetAuthorizedSenders is a paid mutator transaction binding the contract method 0xee56997b.
//
// Solidity: function setAuthorizedSenders(address[] senders) returns()
func (_BillingRegistry *BillingRegistryTransactor) SetAuthorizedSenders(opts *bind.TransactOpts, senders []common.Address) (*types.Transaction, error) {
	return _BillingRegistry.contract.Transact(opts, "setAuthorizedSenders", senders)
}

// SetAuthorizedSenders is a paid mutator transaction binding the contract method 0xee56997b.
//
// Solidity: function setAuthorizedSenders(address[] senders) returns()
func (_BillingRegistry *BillingRegistrySession) SetAuthorizedSenders(senders []common.Address) (*types.Transaction, error) {
	return _BillingRegistry.Contract.SetAuthorizedSenders(&_BillingRegistry.TransactOpts, senders)
}

// SetAuthorizedSenders is a paid mutator transaction binding the contract method 0xee56997b.
//
// Solidity: function setAuthorizedSenders(address[] senders) returns()
func (_BillingRegistry *BillingRegistryTransactorSession) SetAuthorizedSenders(senders []common.Address) (*types.Transaction, error) {
	return _BillingRegistry.Contract.SetAuthorizedSenders(&_BillingRegistry.TransactOpts, senders)
}

// SetConfig is a paid mutator transaction binding the contract method 0x27923e41.
//
// Solidity: function setConfig(uint32 maxGasLimit, uint32 stalenessSeconds, uint256 gasAfterPaymentCalculation, int256 fallbackWeiPerUnitLink, uint32 gasOverhead, uint32 requestTimeoutSeconds) returns()
func (_BillingRegistry *BillingRegistryTransactor) SetConfig(opts *bind.TransactOpts, maxGasLimit uint32, stalenessSeconds uint32, gasAfterPaymentCalculation *big.Int, fallbackWeiPerUnitLink *big.Int, gasOverhead uint32, requestTimeoutSeconds uint32) (*types.Transaction, error) {
	return _BillingRegistry.contract.Transact(opts, "setConfig", maxGasLimit, stalenessSeconds, gasAfterPaymentCalculation, fallbackWeiPerUnitLink, gasOverhead, requestTimeoutSeconds)
}

// SetConfig is a paid mutator transaction binding the contract method 0x27923e41.
//
// Solidity: function setConfig(uint32 maxGasLimit, uint32 stalenessSeconds, uint256 gasAfterPaymentCalculation, int256 fallbackWeiPerUnitLink, uint32 gasOverhead, uint32 requestTimeoutSeconds) returns()
func (_BillingRegistry *BillingRegistrySession) SetConfig(maxGasLimit uint32, stalenessSeconds uint32, gasAfterPaymentCalculation *big.Int, fallbackWeiPerUnitLink *big.Int, gasOverhead uint32, requestTimeoutSeconds uint32) (*types.Transaction, error) {
	return _BillingRegistry.Contract.SetConfig(&_BillingRegistry.TransactOpts, maxGasLimit, stalenessSeconds, gasAfterPaymentCalculation, fallbackWeiPerUnitLink, gasOverhead, requestTimeoutSeconds)
}

// SetConfig is a paid mutator transaction binding the contract method 0x27923e41.
//
// Solidity: function setConfig(uint32 maxGasLimit, uint32 stalenessSeconds, uint256 gasAfterPaymentCalculation, int256 fallbackWeiPerUnitLink, uint32 gasOverhead, uint32 requestTimeoutSeconds) returns()
func (_BillingRegistry *BillingRegistryTransactorSession) SetConfig(maxGasLimit uint32, stalenessSeconds uint32, gasAfterPaymentCalculation *big.Int, fallbackWeiPerUnitLink *big.Int, gasOverhead uint32, requestTimeoutSeconds uint32) (*types.Transaction, error) {
	return _BillingRegistry.Contract.SetConfig(&_BillingRegistry.TransactOpts, maxGasLimit, stalenessSeconds, gasAfterPaymentCalculation, fallbackWeiPerUnitLink, gasOverhead, requestTimeoutSeconds)
}

// StartBilling is a paid mutator transaction binding the contract method 0xa9d03c05.
//
// Solidity: function startBilling(bytes data, (uint64,address,uint32,uint256) billing) returns(bytes32)
func (_BillingRegistry *BillingRegistryTransactor) StartBilling(opts *bind.TransactOpts, data []byte, billing FunctionsBillingRegistryInterfaceRequestBilling) (*types.Transaction, error) {
	return _BillingRegistry.contract.Transact(opts, "startBilling", data, billing)
}

// StartBilling is a paid mutator transaction binding the contract method 0xa9d03c05.
//
// Solidity: function startBilling(bytes data, (uint64,address,uint32,uint256) billing) returns(bytes32)
func (_BillingRegistry *BillingRegistrySession) StartBilling(data []byte, billing FunctionsBillingRegistryInterfaceRequestBilling) (*types.Transaction, error) {
	return _BillingRegistry.Contract.StartBilling(&_BillingRegistry.TransactOpts, data, billing)
}

// StartBilling is a paid mutator transaction binding the contract method 0xa9d03c05.
//
// Solidity: function startBilling(bytes data, (uint64,address,uint32,uint256) billing) returns(bytes32)
func (_BillingRegistry *BillingRegistryTransactorSession) StartBilling(data []byte, billing FunctionsBillingRegistryInterfaceRequestBilling) (*types.Transaction, error) {
	return _BillingRegistry.Contract.StartBilling(&_BillingRegistry.TransactOpts, data, billing)
}

// TimeoutRequests is a paid mutator transaction binding the contract method 0x665871ec.
//
// Solidity: function timeoutRequests(bytes32[] requestIdsToTimeout) returns()
func (_BillingRegistry *BillingRegistryTransactor) TimeoutRequests(opts *bind.TransactOpts, requestIdsToTimeout [][32]byte) (*types.Transaction, error) {
	return _BillingRegistry.contract.Transact(opts, "timeoutRequests", requestIdsToTimeout)
}

// TimeoutRequests is a paid mutator transaction binding the contract method 0x665871ec.
//
// Solidity: function timeoutRequests(bytes32[] requestIdsToTimeout) returns()
func (_BillingRegistry *BillingRegistrySession) TimeoutRequests(requestIdsToTimeout [][32]byte) (*types.Transaction, error) {
	return _BillingRegistry.Contract.TimeoutRequests(&_BillingRegistry.TransactOpts, requestIdsToTimeout)
}

// TimeoutRequests is a paid mutator transaction binding the contract method 0x665871ec.
//
// Solidity: function timeoutRequests(bytes32[] requestIdsToTimeout) returns()
func (_BillingRegistry *BillingRegistryTransactorSession) TimeoutRequests(requestIdsToTimeout [][32]byte) (*types.Transaction, error) {
	return _BillingRegistry.Contract.TimeoutRequests(&_BillingRegistry.TransactOpts, requestIdsToTimeout)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address to) returns()
func (_BillingRegistry *BillingRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _BillingRegistry.contract.Transact(opts, "transferOwnership", to)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address to) returns()
func (_BillingRegistry *BillingRegistrySession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _BillingRegistry.Contract.TransferOwnership(&_BillingRegistry.TransactOpts, to)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address to) returns()
func (_BillingRegistry *BillingRegistryTransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _BillingRegistry.Contract.TransferOwnership(&_BillingRegistry.TransactOpts, to)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BillingRegistry *BillingRegistryTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BillingRegistry.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BillingRegistry *BillingRegistrySession) Unpause() (*types.Transaction, error) {
	return _BillingRegistry.Contract.Unpause(&_BillingRegistry.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BillingRegistry *BillingRegistryTransactorSession) Unpause() (*types.Transaction, error) {
	return _BillingRegistry.Contract.Unpause(&_BillingRegistry.TransactOpts)
}

// BillingRegistryAuthorizedSendersChangedIterator is returned from FilterAuthorizedSendersChanged and is used to iterate over the raw logs and unpacked data for AuthorizedSendersChanged events raised by the BillingRegistry contract.
type BillingRegistryAuthorizedSendersChangedIterator struct {
	Event *BillingRegistryAuthorizedSendersChanged // Event containing the contract specifics and raw log

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
func (it *BillingRegistryAuthorizedSendersChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BillingRegistryAuthorizedSendersChanged)
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
		it.Event = new(BillingRegistryAuthorizedSendersChanged)
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
func (it *BillingRegistryAuthorizedSendersChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BillingRegistryAuthorizedSendersChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BillingRegistryAuthorizedSendersChanged represents a AuthorizedSendersChanged event raised by the BillingRegistry contract.
type BillingRegistryAuthorizedSendersChanged struct {
	Senders   []common.Address
	ChangedBy common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAuthorizedSendersChanged is a free log retrieval operation binding the contract event 0xf263cfb3e4298332e776194610cf9fdc09ccb3ada8b9aa39764d882e11fbf0a0.
//
// Solidity: event AuthorizedSendersChanged(address[] senders, address changedBy)
func (_BillingRegistry *BillingRegistryFilterer) FilterAuthorizedSendersChanged(opts *bind.FilterOpts) (*BillingRegistryAuthorizedSendersChangedIterator, error) {

	logs, sub, err := _BillingRegistry.contract.FilterLogs(opts, "AuthorizedSendersChanged")
	if err != nil {
		return nil, err
	}
	return &BillingRegistryAuthorizedSendersChangedIterator{contract: _BillingRegistry.contract, event: "AuthorizedSendersChanged", logs: logs, sub: sub}, nil
}

// WatchAuthorizedSendersChanged is a free log subscription operation binding the contract event 0xf263cfb3e4298332e776194610cf9fdc09ccb3ada8b9aa39764d882e11fbf0a0.
//
// Solidity: event AuthorizedSendersChanged(address[] senders, address changedBy)
func (_BillingRegistry *BillingRegistryFilterer) WatchAuthorizedSendersChanged(opts *bind.WatchOpts, sink chan<- *BillingRegistryAuthorizedSendersChanged) (event.Subscription, error) {

	logs, sub, err := _BillingRegistry.contract.WatchLogs(opts, "AuthorizedSendersChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BillingRegistryAuthorizedSendersChanged)
				if err := _BillingRegistry.contract.UnpackLog(event, "AuthorizedSendersChanged", log); err != nil {
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
func (_BillingRegistry *BillingRegistryFilterer) ParseAuthorizedSendersChanged(log types.Log) (*BillingRegistryAuthorizedSendersChanged, error) {
	event := new(BillingRegistryAuthorizedSendersChanged)
	if err := _BillingRegistry.contract.UnpackLog(event, "AuthorizedSendersChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BillingRegistryBillingEndIterator is returned from FilterBillingEnd and is used to iterate over the raw logs and unpacked data for BillingEnd events raised by the BillingRegistry contract.
type BillingRegistryBillingEndIterator struct {
	Event *BillingRegistryBillingEnd // Event containing the contract specifics and raw log

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
func (it *BillingRegistryBillingEndIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BillingRegistryBillingEnd)
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
		it.Event = new(BillingRegistryBillingEnd)
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
func (it *BillingRegistryBillingEndIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BillingRegistryBillingEndIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BillingRegistryBillingEnd represents a BillingEnd event raised by the BillingRegistry contract.
type BillingRegistryBillingEnd struct {
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
func (_BillingRegistry *BillingRegistryFilterer) FilterBillingEnd(opts *bind.FilterOpts, requestId [][32]byte) (*BillingRegistryBillingEndIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _BillingRegistry.contract.FilterLogs(opts, "BillingEnd", requestIdRule)
	if err != nil {
		return nil, err
	}
	return &BillingRegistryBillingEndIterator{contract: _BillingRegistry.contract, event: "BillingEnd", logs: logs, sub: sub}, nil
}

// WatchBillingEnd is a free log subscription operation binding the contract event 0xc8dc973332de19a5f71b6026983110e9c2e04b0c98b87eb771ccb78607fd114f.
//
// Solidity: event BillingEnd(bytes32 indexed requestId, uint64 subscriptionId, uint96 signerPayment, uint96 transmitterPayment, uint96 totalCost, bool success)
func (_BillingRegistry *BillingRegistryFilterer) WatchBillingEnd(opts *bind.WatchOpts, sink chan<- *BillingRegistryBillingEnd, requestId [][32]byte) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _BillingRegistry.contract.WatchLogs(opts, "BillingEnd", requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BillingRegistryBillingEnd)
				if err := _BillingRegistry.contract.UnpackLog(event, "BillingEnd", log); err != nil {
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
func (_BillingRegistry *BillingRegistryFilterer) ParseBillingEnd(log types.Log) (*BillingRegistryBillingEnd, error) {
	event := new(BillingRegistryBillingEnd)
	if err := _BillingRegistry.contract.UnpackLog(event, "BillingEnd", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BillingRegistryBillingStartIterator is returned from FilterBillingStart and is used to iterate over the raw logs and unpacked data for BillingStart events raised by the BillingRegistry contract.
type BillingRegistryBillingStartIterator struct {
	Event *BillingRegistryBillingStart // Event containing the contract specifics and raw log

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
func (it *BillingRegistryBillingStartIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BillingRegistryBillingStart)
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
		it.Event = new(BillingRegistryBillingStart)
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
func (it *BillingRegistryBillingStartIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BillingRegistryBillingStartIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BillingRegistryBillingStart represents a BillingStart event raised by the BillingRegistry contract.
type BillingRegistryBillingStart struct {
	RequestId  [32]byte
	Commitment FunctionsBillingRegistryCommitment
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBillingStart is a free log retrieval operation binding the contract event 0x99f7f4e65b4b9fbabd4e357c47ed3099b36e57ecd3a43e84662f34c207d0ebe4.
//
// Solidity: event BillingStart(bytes32 indexed requestId, (uint64,address,uint32,uint256,address,uint96,uint96,uint96,uint256) commitment)
func (_BillingRegistry *BillingRegistryFilterer) FilterBillingStart(opts *bind.FilterOpts, requestId [][32]byte) (*BillingRegistryBillingStartIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _BillingRegistry.contract.FilterLogs(opts, "BillingStart", requestIdRule)
	if err != nil {
		return nil, err
	}
	return &BillingRegistryBillingStartIterator{contract: _BillingRegistry.contract, event: "BillingStart", logs: logs, sub: sub}, nil
}

// WatchBillingStart is a free log subscription operation binding the contract event 0x99f7f4e65b4b9fbabd4e357c47ed3099b36e57ecd3a43e84662f34c207d0ebe4.
//
// Solidity: event BillingStart(bytes32 indexed requestId, (uint64,address,uint32,uint256,address,uint96,uint96,uint96,uint256) commitment)
func (_BillingRegistry *BillingRegistryFilterer) WatchBillingStart(opts *bind.WatchOpts, sink chan<- *BillingRegistryBillingStart, requestId [][32]byte) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _BillingRegistry.contract.WatchLogs(opts, "BillingStart", requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BillingRegistryBillingStart)
				if err := _BillingRegistry.contract.UnpackLog(event, "BillingStart", log); err != nil {
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
func (_BillingRegistry *BillingRegistryFilterer) ParseBillingStart(log types.Log) (*BillingRegistryBillingStart, error) {
	event := new(BillingRegistryBillingStart)
	if err := _BillingRegistry.contract.UnpackLog(event, "BillingStart", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BillingRegistryConfigSetIterator is returned from FilterConfigSet and is used to iterate over the raw logs and unpacked data for ConfigSet events raised by the BillingRegistry contract.
type BillingRegistryConfigSetIterator struct {
	Event *BillingRegistryConfigSet // Event containing the contract specifics and raw log

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
func (it *BillingRegistryConfigSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BillingRegistryConfigSet)
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
		it.Event = new(BillingRegistryConfigSet)
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
func (it *BillingRegistryConfigSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BillingRegistryConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BillingRegistryConfigSet represents a ConfigSet event raised by the BillingRegistry contract.
type BillingRegistryConfigSet struct {
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
func (_BillingRegistry *BillingRegistryFilterer) FilterConfigSet(opts *bind.FilterOpts) (*BillingRegistryConfigSetIterator, error) {

	logs, sub, err := _BillingRegistry.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &BillingRegistryConfigSetIterator{contract: _BillingRegistry.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

// WatchConfigSet is a free log subscription operation binding the contract event 0x24d3d934adfef9b9029d6ffa463c07d0139ed47d26ee23506f85ece2879d2bd4.
//
// Solidity: event ConfigSet(uint32 maxGasLimit, uint32 stalenessSeconds, uint256 gasAfterPaymentCalculation, int256 fallbackWeiPerUnitLink, uint32 gasOverhead)
func (_BillingRegistry *BillingRegistryFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *BillingRegistryConfigSet) (event.Subscription, error) {

	logs, sub, err := _BillingRegistry.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BillingRegistryConfigSet)
				if err := _BillingRegistry.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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
func (_BillingRegistry *BillingRegistryFilterer) ParseConfigSet(log types.Log) (*BillingRegistryConfigSet, error) {
	event := new(BillingRegistryConfigSet)
	if err := _BillingRegistry.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BillingRegistryFundsRecoveredIterator is returned from FilterFundsRecovered and is used to iterate over the raw logs and unpacked data for FundsRecovered events raised by the BillingRegistry contract.
type BillingRegistryFundsRecoveredIterator struct {
	Event *BillingRegistryFundsRecovered // Event containing the contract specifics and raw log

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
func (it *BillingRegistryFundsRecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BillingRegistryFundsRecovered)
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
		it.Event = new(BillingRegistryFundsRecovered)
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
func (it *BillingRegistryFundsRecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BillingRegistryFundsRecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BillingRegistryFundsRecovered represents a FundsRecovered event raised by the BillingRegistry contract.
type BillingRegistryFundsRecovered struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFundsRecovered is a free log retrieval operation binding the contract event 0x59bfc682b673f8cbf945f1e454df9334834abf7dfe7f92237ca29ecb9b436600.
//
// Solidity: event FundsRecovered(address to, uint256 amount)
func (_BillingRegistry *BillingRegistryFilterer) FilterFundsRecovered(opts *bind.FilterOpts) (*BillingRegistryFundsRecoveredIterator, error) {

	logs, sub, err := _BillingRegistry.contract.FilterLogs(opts, "FundsRecovered")
	if err != nil {
		return nil, err
	}
	return &BillingRegistryFundsRecoveredIterator{contract: _BillingRegistry.contract, event: "FundsRecovered", logs: logs, sub: sub}, nil
}

// WatchFundsRecovered is a free log subscription operation binding the contract event 0x59bfc682b673f8cbf945f1e454df9334834abf7dfe7f92237ca29ecb9b436600.
//
// Solidity: event FundsRecovered(address to, uint256 amount)
func (_BillingRegistry *BillingRegistryFilterer) WatchFundsRecovered(opts *bind.WatchOpts, sink chan<- *BillingRegistryFundsRecovered) (event.Subscription, error) {

	logs, sub, err := _BillingRegistry.contract.WatchLogs(opts, "FundsRecovered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BillingRegistryFundsRecovered)
				if err := _BillingRegistry.contract.UnpackLog(event, "FundsRecovered", log); err != nil {
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
func (_BillingRegistry *BillingRegistryFilterer) ParseFundsRecovered(log types.Log) (*BillingRegistryFundsRecovered, error) {
	event := new(BillingRegistryFundsRecovered)
	if err := _BillingRegistry.contract.UnpackLog(event, "FundsRecovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BillingRegistryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the BillingRegistry contract.
type BillingRegistryInitializedIterator struct {
	Event *BillingRegistryInitialized // Event containing the contract specifics and raw log

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
func (it *BillingRegistryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BillingRegistryInitialized)
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
		it.Event = new(BillingRegistryInitialized)
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
func (it *BillingRegistryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BillingRegistryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BillingRegistryInitialized represents a Initialized event raised by the BillingRegistry contract.
type BillingRegistryInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_BillingRegistry *BillingRegistryFilterer) FilterInitialized(opts *bind.FilterOpts) (*BillingRegistryInitializedIterator, error) {

	logs, sub, err := _BillingRegistry.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BillingRegistryInitializedIterator{contract: _BillingRegistry.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_BillingRegistry *BillingRegistryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BillingRegistryInitialized) (event.Subscription, error) {

	logs, sub, err := _BillingRegistry.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BillingRegistryInitialized)
				if err := _BillingRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_BillingRegistry *BillingRegistryFilterer) ParseInitialized(log types.Log) (*BillingRegistryInitialized, error) {
	event := new(BillingRegistryInitialized)
	if err := _BillingRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BillingRegistryOwnershipTransferRequestedIterator is returned from FilterOwnershipTransferRequested and is used to iterate over the raw logs and unpacked data for OwnershipTransferRequested events raised by the BillingRegistry contract.
type BillingRegistryOwnershipTransferRequestedIterator struct {
	Event *BillingRegistryOwnershipTransferRequested // Event containing the contract specifics and raw log

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
func (it *BillingRegistryOwnershipTransferRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BillingRegistryOwnershipTransferRequested)
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
		it.Event = new(BillingRegistryOwnershipTransferRequested)
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
func (it *BillingRegistryOwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BillingRegistryOwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BillingRegistryOwnershipTransferRequested represents a OwnershipTransferRequested event raised by the BillingRegistry contract.
type BillingRegistryOwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferRequested is a free log retrieval operation binding the contract event 0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278.
//
// Solidity: event OwnershipTransferRequested(address indexed from, address indexed to)
func (_BillingRegistry *BillingRegistryFilterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BillingRegistryOwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BillingRegistry.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BillingRegistryOwnershipTransferRequestedIterator{contract: _BillingRegistry.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferRequested is a free log subscription operation binding the contract event 0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278.
//
// Solidity: event OwnershipTransferRequested(address indexed from, address indexed to)
func (_BillingRegistry *BillingRegistryFilterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *BillingRegistryOwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BillingRegistry.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BillingRegistryOwnershipTransferRequested)
				if err := _BillingRegistry.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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
func (_BillingRegistry *BillingRegistryFilterer) ParseOwnershipTransferRequested(log types.Log) (*BillingRegistryOwnershipTransferRequested, error) {
	event := new(BillingRegistryOwnershipTransferRequested)
	if err := _BillingRegistry.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BillingRegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BillingRegistry contract.
type BillingRegistryOwnershipTransferredIterator struct {
	Event *BillingRegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BillingRegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BillingRegistryOwnershipTransferred)
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
		it.Event = new(BillingRegistryOwnershipTransferred)
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
func (it *BillingRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BillingRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BillingRegistryOwnershipTransferred represents a OwnershipTransferred event raised by the BillingRegistry contract.
type BillingRegistryOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed from, address indexed to)
func (_BillingRegistry *BillingRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BillingRegistryOwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BillingRegistry.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BillingRegistryOwnershipTransferredIterator{contract: _BillingRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed from, address indexed to)
func (_BillingRegistry *BillingRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BillingRegistryOwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BillingRegistry.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BillingRegistryOwnershipTransferred)
				if err := _BillingRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_BillingRegistry *BillingRegistryFilterer) ParseOwnershipTransferred(log types.Log) (*BillingRegistryOwnershipTransferred, error) {
	event := new(BillingRegistryOwnershipTransferred)
	if err := _BillingRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BillingRegistryPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the BillingRegistry contract.
type BillingRegistryPausedIterator struct {
	Event *BillingRegistryPaused // Event containing the contract specifics and raw log

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
func (it *BillingRegistryPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BillingRegistryPaused)
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
		it.Event = new(BillingRegistryPaused)
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
func (it *BillingRegistryPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BillingRegistryPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BillingRegistryPaused represents a Paused event raised by the BillingRegistry contract.
type BillingRegistryPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_BillingRegistry *BillingRegistryFilterer) FilterPaused(opts *bind.FilterOpts) (*BillingRegistryPausedIterator, error) {

	logs, sub, err := _BillingRegistry.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &BillingRegistryPausedIterator{contract: _BillingRegistry.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_BillingRegistry *BillingRegistryFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *BillingRegistryPaused) (event.Subscription, error) {

	logs, sub, err := _BillingRegistry.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BillingRegistryPaused)
				if err := _BillingRegistry.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_BillingRegistry *BillingRegistryFilterer) ParsePaused(log types.Log) (*BillingRegistryPaused, error) {
	event := new(BillingRegistryPaused)
	if err := _BillingRegistry.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BillingRegistryRequestTimedOutIterator is returned from FilterRequestTimedOut and is used to iterate over the raw logs and unpacked data for RequestTimedOut events raised by the BillingRegistry contract.
type BillingRegistryRequestTimedOutIterator struct {
	Event *BillingRegistryRequestTimedOut // Event containing the contract specifics and raw log

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
func (it *BillingRegistryRequestTimedOutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BillingRegistryRequestTimedOut)
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
		it.Event = new(BillingRegistryRequestTimedOut)
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
func (it *BillingRegistryRequestTimedOutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BillingRegistryRequestTimedOutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BillingRegistryRequestTimedOut represents a RequestTimedOut event raised by the BillingRegistry contract.
type BillingRegistryRequestTimedOut struct {
	RequestId [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRequestTimedOut is a free log retrieval operation binding the contract event 0xf1ca1e9147be737b04a2b018a79405f687a97de8dd8a2559bbe62357343af414.
//
// Solidity: event RequestTimedOut(bytes32 indexed requestId)
func (_BillingRegistry *BillingRegistryFilterer) FilterRequestTimedOut(opts *bind.FilterOpts, requestId [][32]byte) (*BillingRegistryRequestTimedOutIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _BillingRegistry.contract.FilterLogs(opts, "RequestTimedOut", requestIdRule)
	if err != nil {
		return nil, err
	}
	return &BillingRegistryRequestTimedOutIterator{contract: _BillingRegistry.contract, event: "RequestTimedOut", logs: logs, sub: sub}, nil
}

// WatchRequestTimedOut is a free log subscription operation binding the contract event 0xf1ca1e9147be737b04a2b018a79405f687a97de8dd8a2559bbe62357343af414.
//
// Solidity: event RequestTimedOut(bytes32 indexed requestId)
func (_BillingRegistry *BillingRegistryFilterer) WatchRequestTimedOut(opts *bind.WatchOpts, sink chan<- *BillingRegistryRequestTimedOut, requestId [][32]byte) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _BillingRegistry.contract.WatchLogs(opts, "RequestTimedOut", requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BillingRegistryRequestTimedOut)
				if err := _BillingRegistry.contract.UnpackLog(event, "RequestTimedOut", log); err != nil {
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
func (_BillingRegistry *BillingRegistryFilterer) ParseRequestTimedOut(log types.Log) (*BillingRegistryRequestTimedOut, error) {
	event := new(BillingRegistryRequestTimedOut)
	if err := _BillingRegistry.contract.UnpackLog(event, "RequestTimedOut", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BillingRegistrySubscriptionCanceledIterator is returned from FilterSubscriptionCanceled and is used to iterate over the raw logs and unpacked data for SubscriptionCanceled events raised by the BillingRegistry contract.
type BillingRegistrySubscriptionCanceledIterator struct {
	Event *BillingRegistrySubscriptionCanceled // Event containing the contract specifics and raw log

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
func (it *BillingRegistrySubscriptionCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BillingRegistrySubscriptionCanceled)
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
		it.Event = new(BillingRegistrySubscriptionCanceled)
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
func (it *BillingRegistrySubscriptionCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BillingRegistrySubscriptionCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BillingRegistrySubscriptionCanceled represents a SubscriptionCanceled event raised by the BillingRegistry contract.
type BillingRegistrySubscriptionCanceled struct {
	SubscriptionId uint64
	To             common.Address
	Amount         *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSubscriptionCanceled is a free log retrieval operation binding the contract event 0xe8ed5b475a5b5987aa9165e8731bb78043f39eee32ec5a1169a89e27fcd49815.
//
// Solidity: event SubscriptionCanceled(uint64 indexed subscriptionId, address to, uint256 amount)
func (_BillingRegistry *BillingRegistryFilterer) FilterSubscriptionCanceled(opts *bind.FilterOpts, subscriptionId []uint64) (*BillingRegistrySubscriptionCanceledIterator, error) {

	var subscriptionIdRule []interface{}
	for _, subscriptionIdItem := range subscriptionId {
		subscriptionIdRule = append(subscriptionIdRule, subscriptionIdItem)
	}

	logs, sub, err := _BillingRegistry.contract.FilterLogs(opts, "SubscriptionCanceled", subscriptionIdRule)
	if err != nil {
		return nil, err
	}
	return &BillingRegistrySubscriptionCanceledIterator{contract: _BillingRegistry.contract, event: "SubscriptionCanceled", logs: logs, sub: sub}, nil
}

// WatchSubscriptionCanceled is a free log subscription operation binding the contract event 0xe8ed5b475a5b5987aa9165e8731bb78043f39eee32ec5a1169a89e27fcd49815.
//
// Solidity: event SubscriptionCanceled(uint64 indexed subscriptionId, address to, uint256 amount)
func (_BillingRegistry *BillingRegistryFilterer) WatchSubscriptionCanceled(opts *bind.WatchOpts, sink chan<- *BillingRegistrySubscriptionCanceled, subscriptionId []uint64) (event.Subscription, error) {

	var subscriptionIdRule []interface{}
	for _, subscriptionIdItem := range subscriptionId {
		subscriptionIdRule = append(subscriptionIdRule, subscriptionIdItem)
	}

	logs, sub, err := _BillingRegistry.contract.WatchLogs(opts, "SubscriptionCanceled", subscriptionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BillingRegistrySubscriptionCanceled)
				if err := _BillingRegistry.contract.UnpackLog(event, "SubscriptionCanceled", log); err != nil {
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
func (_BillingRegistry *BillingRegistryFilterer) ParseSubscriptionCanceled(log types.Log) (*BillingRegistrySubscriptionCanceled, error) {
	event := new(BillingRegistrySubscriptionCanceled)
	if err := _BillingRegistry.contract.UnpackLog(event, "SubscriptionCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BillingRegistrySubscriptionConsumerAddedIterator is returned from FilterSubscriptionConsumerAdded and is used to iterate over the raw logs and unpacked data for SubscriptionConsumerAdded events raised by the BillingRegistry contract.
type BillingRegistrySubscriptionConsumerAddedIterator struct {
	Event *BillingRegistrySubscriptionConsumerAdded // Event containing the contract specifics and raw log

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
func (it *BillingRegistrySubscriptionConsumerAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BillingRegistrySubscriptionConsumerAdded)
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
		it.Event = new(BillingRegistrySubscriptionConsumerAdded)
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
func (it *BillingRegistrySubscriptionConsumerAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BillingRegistrySubscriptionConsumerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BillingRegistrySubscriptionConsumerAdded represents a SubscriptionConsumerAdded event raised by the BillingRegistry contract.
type BillingRegistrySubscriptionConsumerAdded struct {
	SubscriptionId uint64
	Consumer       common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSubscriptionConsumerAdded is a free log retrieval operation binding the contract event 0x43dc749a04ac8fb825cbd514f7c0e13f13bc6f2ee66043b76629d51776cff8e0.
//
// Solidity: event SubscriptionConsumerAdded(uint64 indexed subscriptionId, address consumer)
func (_BillingRegistry *BillingRegistryFilterer) FilterSubscriptionConsumerAdded(opts *bind.FilterOpts, subscriptionId []uint64) (*BillingRegistrySubscriptionConsumerAddedIterator, error) {

	var subscriptionIdRule []interface{}
	for _, subscriptionIdItem := range subscriptionId {
		subscriptionIdRule = append(subscriptionIdRule, subscriptionIdItem)
	}

	logs, sub, err := _BillingRegistry.contract.FilterLogs(opts, "SubscriptionConsumerAdded", subscriptionIdRule)
	if err != nil {
		return nil, err
	}
	return &BillingRegistrySubscriptionConsumerAddedIterator{contract: _BillingRegistry.contract, event: "SubscriptionConsumerAdded", logs: logs, sub: sub}, nil
}

// WatchSubscriptionConsumerAdded is a free log subscription operation binding the contract event 0x43dc749a04ac8fb825cbd514f7c0e13f13bc6f2ee66043b76629d51776cff8e0.
//
// Solidity: event SubscriptionConsumerAdded(uint64 indexed subscriptionId, address consumer)
func (_BillingRegistry *BillingRegistryFilterer) WatchSubscriptionConsumerAdded(opts *bind.WatchOpts, sink chan<- *BillingRegistrySubscriptionConsumerAdded, subscriptionId []uint64) (event.Subscription, error) {

	var subscriptionIdRule []interface{}
	for _, subscriptionIdItem := range subscriptionId {
		subscriptionIdRule = append(subscriptionIdRule, subscriptionIdItem)
	}

	logs, sub, err := _BillingRegistry.contract.WatchLogs(opts, "SubscriptionConsumerAdded", subscriptionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BillingRegistrySubscriptionConsumerAdded)
				if err := _BillingRegistry.contract.UnpackLog(event, "SubscriptionConsumerAdded", log); err != nil {
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
func (_BillingRegistry *BillingRegistryFilterer) ParseSubscriptionConsumerAdded(log types.Log) (*BillingRegistrySubscriptionConsumerAdded, error) {
	event := new(BillingRegistrySubscriptionConsumerAdded)
	if err := _BillingRegistry.contract.UnpackLog(event, "SubscriptionConsumerAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BillingRegistrySubscriptionConsumerRemovedIterator is returned from FilterSubscriptionConsumerRemoved and is used to iterate over the raw logs and unpacked data for SubscriptionConsumerRemoved events raised by the BillingRegistry contract.
type BillingRegistrySubscriptionConsumerRemovedIterator struct {
	Event *BillingRegistrySubscriptionConsumerRemoved // Event containing the contract specifics and raw log

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
func (it *BillingRegistrySubscriptionConsumerRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BillingRegistrySubscriptionConsumerRemoved)
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
		it.Event = new(BillingRegistrySubscriptionConsumerRemoved)
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
func (it *BillingRegistrySubscriptionConsumerRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BillingRegistrySubscriptionConsumerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BillingRegistrySubscriptionConsumerRemoved represents a SubscriptionConsumerRemoved event raised by the BillingRegistry contract.
type BillingRegistrySubscriptionConsumerRemoved struct {
	SubscriptionId uint64
	Consumer       common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSubscriptionConsumerRemoved is a free log retrieval operation binding the contract event 0x182bff9831466789164ca77075fffd84916d35a8180ba73c27e45634549b445b.
//
// Solidity: event SubscriptionConsumerRemoved(uint64 indexed subscriptionId, address consumer)
func (_BillingRegistry *BillingRegistryFilterer) FilterSubscriptionConsumerRemoved(opts *bind.FilterOpts, subscriptionId []uint64) (*BillingRegistrySubscriptionConsumerRemovedIterator, error) {

	var subscriptionIdRule []interface{}
	for _, subscriptionIdItem := range subscriptionId {
		subscriptionIdRule = append(subscriptionIdRule, subscriptionIdItem)
	}

	logs, sub, err := _BillingRegistry.contract.FilterLogs(opts, "SubscriptionConsumerRemoved", subscriptionIdRule)
	if err != nil {
		return nil, err
	}
	return &BillingRegistrySubscriptionConsumerRemovedIterator{contract: _BillingRegistry.contract, event: "SubscriptionConsumerRemoved", logs: logs, sub: sub}, nil
}

// WatchSubscriptionConsumerRemoved is a free log subscription operation binding the contract event 0x182bff9831466789164ca77075fffd84916d35a8180ba73c27e45634549b445b.
//
// Solidity: event SubscriptionConsumerRemoved(uint64 indexed subscriptionId, address consumer)
func (_BillingRegistry *BillingRegistryFilterer) WatchSubscriptionConsumerRemoved(opts *bind.WatchOpts, sink chan<- *BillingRegistrySubscriptionConsumerRemoved, subscriptionId []uint64) (event.Subscription, error) {

	var subscriptionIdRule []interface{}
	for _, subscriptionIdItem := range subscriptionId {
		subscriptionIdRule = append(subscriptionIdRule, subscriptionIdItem)
	}

	logs, sub, err := _BillingRegistry.contract.WatchLogs(opts, "SubscriptionConsumerRemoved", subscriptionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BillingRegistrySubscriptionConsumerRemoved)
				if err := _BillingRegistry.contract.UnpackLog(event, "SubscriptionConsumerRemoved", log); err != nil {
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
func (_BillingRegistry *BillingRegistryFilterer) ParseSubscriptionConsumerRemoved(log types.Log) (*BillingRegistrySubscriptionConsumerRemoved, error) {
	event := new(BillingRegistrySubscriptionConsumerRemoved)
	if err := _BillingRegistry.contract.UnpackLog(event, "SubscriptionConsumerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BillingRegistrySubscriptionCreatedIterator is returned from FilterSubscriptionCreated and is used to iterate over the raw logs and unpacked data for SubscriptionCreated events raised by the BillingRegistry contract.
type BillingRegistrySubscriptionCreatedIterator struct {
	Event *BillingRegistrySubscriptionCreated // Event containing the contract specifics and raw log

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
func (it *BillingRegistrySubscriptionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BillingRegistrySubscriptionCreated)
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
		it.Event = new(BillingRegistrySubscriptionCreated)
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
func (it *BillingRegistrySubscriptionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BillingRegistrySubscriptionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BillingRegistrySubscriptionCreated represents a SubscriptionCreated event raised by the BillingRegistry contract.
type BillingRegistrySubscriptionCreated struct {
	SubscriptionId uint64
	Owner          common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSubscriptionCreated is a free log retrieval operation binding the contract event 0x464722b4166576d3dcbba877b999bc35cf911f4eaf434b7eba68fa113951d0bf.
//
// Solidity: event SubscriptionCreated(uint64 indexed subscriptionId, address owner)
func (_BillingRegistry *BillingRegistryFilterer) FilterSubscriptionCreated(opts *bind.FilterOpts, subscriptionId []uint64) (*BillingRegistrySubscriptionCreatedIterator, error) {

	var subscriptionIdRule []interface{}
	for _, subscriptionIdItem := range subscriptionId {
		subscriptionIdRule = append(subscriptionIdRule, subscriptionIdItem)
	}

	logs, sub, err := _BillingRegistry.contract.FilterLogs(opts, "SubscriptionCreated", subscriptionIdRule)
	if err != nil {
		return nil, err
	}
	return &BillingRegistrySubscriptionCreatedIterator{contract: _BillingRegistry.contract, event: "SubscriptionCreated", logs: logs, sub: sub}, nil
}

// WatchSubscriptionCreated is a free log subscription operation binding the contract event 0x464722b4166576d3dcbba877b999bc35cf911f4eaf434b7eba68fa113951d0bf.
//
// Solidity: event SubscriptionCreated(uint64 indexed subscriptionId, address owner)
func (_BillingRegistry *BillingRegistryFilterer) WatchSubscriptionCreated(opts *bind.WatchOpts, sink chan<- *BillingRegistrySubscriptionCreated, subscriptionId []uint64) (event.Subscription, error) {

	var subscriptionIdRule []interface{}
	for _, subscriptionIdItem := range subscriptionId {
		subscriptionIdRule = append(subscriptionIdRule, subscriptionIdItem)
	}

	logs, sub, err := _BillingRegistry.contract.WatchLogs(opts, "SubscriptionCreated", subscriptionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BillingRegistrySubscriptionCreated)
				if err := _BillingRegistry.contract.UnpackLog(event, "SubscriptionCreated", log); err != nil {
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
func (_BillingRegistry *BillingRegistryFilterer) ParseSubscriptionCreated(log types.Log) (*BillingRegistrySubscriptionCreated, error) {
	event := new(BillingRegistrySubscriptionCreated)
	if err := _BillingRegistry.contract.UnpackLog(event, "SubscriptionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BillingRegistrySubscriptionFundedIterator is returned from FilterSubscriptionFunded and is used to iterate over the raw logs and unpacked data for SubscriptionFunded events raised by the BillingRegistry contract.
type BillingRegistrySubscriptionFundedIterator struct {
	Event *BillingRegistrySubscriptionFunded // Event containing the contract specifics and raw log

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
func (it *BillingRegistrySubscriptionFundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BillingRegistrySubscriptionFunded)
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
		it.Event = new(BillingRegistrySubscriptionFunded)
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
func (it *BillingRegistrySubscriptionFundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BillingRegistrySubscriptionFundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BillingRegistrySubscriptionFunded represents a SubscriptionFunded event raised by the BillingRegistry contract.
type BillingRegistrySubscriptionFunded struct {
	SubscriptionId uint64
	OldBalance     *big.Int
	NewBalance     *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSubscriptionFunded is a free log retrieval operation binding the contract event 0xd39ec07f4e209f627a4c427971473820dc129761ba28de8906bd56f57101d4f8.
//
// Solidity: event SubscriptionFunded(uint64 indexed subscriptionId, uint256 oldBalance, uint256 newBalance)
func (_BillingRegistry *BillingRegistryFilterer) FilterSubscriptionFunded(opts *bind.FilterOpts, subscriptionId []uint64) (*BillingRegistrySubscriptionFundedIterator, error) {

	var subscriptionIdRule []interface{}
	for _, subscriptionIdItem := range subscriptionId {
		subscriptionIdRule = append(subscriptionIdRule, subscriptionIdItem)
	}

	logs, sub, err := _BillingRegistry.contract.FilterLogs(opts, "SubscriptionFunded", subscriptionIdRule)
	if err != nil {
		return nil, err
	}
	return &BillingRegistrySubscriptionFundedIterator{contract: _BillingRegistry.contract, event: "SubscriptionFunded", logs: logs, sub: sub}, nil
}

// WatchSubscriptionFunded is a free log subscription operation binding the contract event 0xd39ec07f4e209f627a4c427971473820dc129761ba28de8906bd56f57101d4f8.
//
// Solidity: event SubscriptionFunded(uint64 indexed subscriptionId, uint256 oldBalance, uint256 newBalance)
func (_BillingRegistry *BillingRegistryFilterer) WatchSubscriptionFunded(opts *bind.WatchOpts, sink chan<- *BillingRegistrySubscriptionFunded, subscriptionId []uint64) (event.Subscription, error) {

	var subscriptionIdRule []interface{}
	for _, subscriptionIdItem := range subscriptionId {
		subscriptionIdRule = append(subscriptionIdRule, subscriptionIdItem)
	}

	logs, sub, err := _BillingRegistry.contract.WatchLogs(opts, "SubscriptionFunded", subscriptionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BillingRegistrySubscriptionFunded)
				if err := _BillingRegistry.contract.UnpackLog(event, "SubscriptionFunded", log); err != nil {
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
func (_BillingRegistry *BillingRegistryFilterer) ParseSubscriptionFunded(log types.Log) (*BillingRegistrySubscriptionFunded, error) {
	event := new(BillingRegistrySubscriptionFunded)
	if err := _BillingRegistry.contract.UnpackLog(event, "SubscriptionFunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BillingRegistrySubscriptionOwnerTransferRequestedIterator is returned from FilterSubscriptionOwnerTransferRequested and is used to iterate over the raw logs and unpacked data for SubscriptionOwnerTransferRequested events raised by the BillingRegistry contract.
type BillingRegistrySubscriptionOwnerTransferRequestedIterator struct {
	Event *BillingRegistrySubscriptionOwnerTransferRequested // Event containing the contract specifics and raw log

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
func (it *BillingRegistrySubscriptionOwnerTransferRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BillingRegistrySubscriptionOwnerTransferRequested)
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
		it.Event = new(BillingRegistrySubscriptionOwnerTransferRequested)
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
func (it *BillingRegistrySubscriptionOwnerTransferRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BillingRegistrySubscriptionOwnerTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BillingRegistrySubscriptionOwnerTransferRequested represents a SubscriptionOwnerTransferRequested event raised by the BillingRegistry contract.
type BillingRegistrySubscriptionOwnerTransferRequested struct {
	SubscriptionId uint64
	From           common.Address
	To             common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSubscriptionOwnerTransferRequested is a free log retrieval operation binding the contract event 0x69436ea6df009049404f564eff6622cd00522b0bd6a89efd9e52a355c4a879be.
//
// Solidity: event SubscriptionOwnerTransferRequested(uint64 indexed subscriptionId, address from, address to)
func (_BillingRegistry *BillingRegistryFilterer) FilterSubscriptionOwnerTransferRequested(opts *bind.FilterOpts, subscriptionId []uint64) (*BillingRegistrySubscriptionOwnerTransferRequestedIterator, error) {

	var subscriptionIdRule []interface{}
	for _, subscriptionIdItem := range subscriptionId {
		subscriptionIdRule = append(subscriptionIdRule, subscriptionIdItem)
	}

	logs, sub, err := _BillingRegistry.contract.FilterLogs(opts, "SubscriptionOwnerTransferRequested", subscriptionIdRule)
	if err != nil {
		return nil, err
	}
	return &BillingRegistrySubscriptionOwnerTransferRequestedIterator{contract: _BillingRegistry.contract, event: "SubscriptionOwnerTransferRequested", logs: logs, sub: sub}, nil
}

// WatchSubscriptionOwnerTransferRequested is a free log subscription operation binding the contract event 0x69436ea6df009049404f564eff6622cd00522b0bd6a89efd9e52a355c4a879be.
//
// Solidity: event SubscriptionOwnerTransferRequested(uint64 indexed subscriptionId, address from, address to)
func (_BillingRegistry *BillingRegistryFilterer) WatchSubscriptionOwnerTransferRequested(opts *bind.WatchOpts, sink chan<- *BillingRegistrySubscriptionOwnerTransferRequested, subscriptionId []uint64) (event.Subscription, error) {

	var subscriptionIdRule []interface{}
	for _, subscriptionIdItem := range subscriptionId {
		subscriptionIdRule = append(subscriptionIdRule, subscriptionIdItem)
	}

	logs, sub, err := _BillingRegistry.contract.WatchLogs(opts, "SubscriptionOwnerTransferRequested", subscriptionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BillingRegistrySubscriptionOwnerTransferRequested)
				if err := _BillingRegistry.contract.UnpackLog(event, "SubscriptionOwnerTransferRequested", log); err != nil {
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
func (_BillingRegistry *BillingRegistryFilterer) ParseSubscriptionOwnerTransferRequested(log types.Log) (*BillingRegistrySubscriptionOwnerTransferRequested, error) {
	event := new(BillingRegistrySubscriptionOwnerTransferRequested)
	if err := _BillingRegistry.contract.UnpackLog(event, "SubscriptionOwnerTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BillingRegistrySubscriptionOwnerTransferredIterator is returned from FilterSubscriptionOwnerTransferred and is used to iterate over the raw logs and unpacked data for SubscriptionOwnerTransferred events raised by the BillingRegistry contract.
type BillingRegistrySubscriptionOwnerTransferredIterator struct {
	Event *BillingRegistrySubscriptionOwnerTransferred // Event containing the contract specifics and raw log

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
func (it *BillingRegistrySubscriptionOwnerTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BillingRegistrySubscriptionOwnerTransferred)
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
		it.Event = new(BillingRegistrySubscriptionOwnerTransferred)
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
func (it *BillingRegistrySubscriptionOwnerTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BillingRegistrySubscriptionOwnerTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BillingRegistrySubscriptionOwnerTransferred represents a SubscriptionOwnerTransferred event raised by the BillingRegistry contract.
type BillingRegistrySubscriptionOwnerTransferred struct {
	SubscriptionId uint64
	From           common.Address
	To             common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSubscriptionOwnerTransferred is a free log retrieval operation binding the contract event 0x6f1dc65165ffffedfd8e507b4a0f1fcfdada045ed11f6c26ba27cedfe87802f0.
//
// Solidity: event SubscriptionOwnerTransferred(uint64 indexed subscriptionId, address from, address to)
func (_BillingRegistry *BillingRegistryFilterer) FilterSubscriptionOwnerTransferred(opts *bind.FilterOpts, subscriptionId []uint64) (*BillingRegistrySubscriptionOwnerTransferredIterator, error) {

	var subscriptionIdRule []interface{}
	for _, subscriptionIdItem := range subscriptionId {
		subscriptionIdRule = append(subscriptionIdRule, subscriptionIdItem)
	}

	logs, sub, err := _BillingRegistry.contract.FilterLogs(opts, "SubscriptionOwnerTransferred", subscriptionIdRule)
	if err != nil {
		return nil, err
	}
	return &BillingRegistrySubscriptionOwnerTransferredIterator{contract: _BillingRegistry.contract, event: "SubscriptionOwnerTransferred", logs: logs, sub: sub}, nil
}

// WatchSubscriptionOwnerTransferred is a free log subscription operation binding the contract event 0x6f1dc65165ffffedfd8e507b4a0f1fcfdada045ed11f6c26ba27cedfe87802f0.
//
// Solidity: event SubscriptionOwnerTransferred(uint64 indexed subscriptionId, address from, address to)
func (_BillingRegistry *BillingRegistryFilterer) WatchSubscriptionOwnerTransferred(opts *bind.WatchOpts, sink chan<- *BillingRegistrySubscriptionOwnerTransferred, subscriptionId []uint64) (event.Subscription, error) {

	var subscriptionIdRule []interface{}
	for _, subscriptionIdItem := range subscriptionId {
		subscriptionIdRule = append(subscriptionIdRule, subscriptionIdItem)
	}

	logs, sub, err := _BillingRegistry.contract.WatchLogs(opts, "SubscriptionOwnerTransferred", subscriptionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BillingRegistrySubscriptionOwnerTransferred)
				if err := _BillingRegistry.contract.UnpackLog(event, "SubscriptionOwnerTransferred", log); err != nil {
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
func (_BillingRegistry *BillingRegistryFilterer) ParseSubscriptionOwnerTransferred(log types.Log) (*BillingRegistrySubscriptionOwnerTransferred, error) {
	event := new(BillingRegistrySubscriptionOwnerTransferred)
	if err := _BillingRegistry.contract.UnpackLog(event, "SubscriptionOwnerTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BillingRegistryUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the BillingRegistry contract.
type BillingRegistryUnpausedIterator struct {
	Event *BillingRegistryUnpaused // Event containing the contract specifics and raw log

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
func (it *BillingRegistryUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BillingRegistryUnpaused)
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
		it.Event = new(BillingRegistryUnpaused)
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
func (it *BillingRegistryUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BillingRegistryUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BillingRegistryUnpaused represents a Unpaused event raised by the BillingRegistry contract.
type BillingRegistryUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_BillingRegistry *BillingRegistryFilterer) FilterUnpaused(opts *bind.FilterOpts) (*BillingRegistryUnpausedIterator, error) {

	logs, sub, err := _BillingRegistry.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &BillingRegistryUnpausedIterator{contract: _BillingRegistry.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_BillingRegistry *BillingRegistryFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *BillingRegistryUnpaused) (event.Subscription, error) {

	logs, sub, err := _BillingRegistry.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BillingRegistryUnpaused)
				if err := _BillingRegistry.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_BillingRegistry *BillingRegistryFilterer) ParseUnpaused(log types.Log) (*BillingRegistryUnpaused, error) {
	event := new(BillingRegistryUnpaused)
	if err := _BillingRegistry.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
