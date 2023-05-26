// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package recovery

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
)

// EmailGuardianMetaData contains all meta data concerning the EmailGuardian contract.
var EmailGuardianMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"email\",\"type\":\"bytes32\"}],\"name\":\"account\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// EmailGuardianABI is the input ABI used to generate the binding from.
// Deprecated: Use EmailGuardianMetaData.ABI instead.
var EmailGuardianABI = EmailGuardianMetaData.ABI

// EmailGuardian is an auto generated Go binding around an Ethereum contract.
type EmailGuardian struct {
	EmailGuardianCaller     // Read-only binding to the contract
	EmailGuardianTransactor // Write-only binding to the contract
	EmailGuardianFilterer   // Log filterer for contract events
}

// EmailGuardianCaller is an auto generated read-only Go binding around an Ethereum contract.
type EmailGuardianCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EmailGuardianTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EmailGuardianTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EmailGuardianFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EmailGuardianFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EmailGuardianSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EmailGuardianSession struct {
	Contract     *EmailGuardian    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EmailGuardianCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EmailGuardianCallerSession struct {
	Contract *EmailGuardianCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// EmailGuardianTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EmailGuardianTransactorSession struct {
	Contract     *EmailGuardianTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// EmailGuardianRaw is an auto generated low-level Go binding around an Ethereum contract.
type EmailGuardianRaw struct {
	Contract *EmailGuardian // Generic contract binding to access the raw methods on
}

// EmailGuardianCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EmailGuardianCallerRaw struct {
	Contract *EmailGuardianCaller // Generic read-only contract binding to access the raw methods on
}

// EmailGuardianTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EmailGuardianTransactorRaw struct {
	Contract *EmailGuardianTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEmailGuardian creates a new instance of EmailGuardian, bound to a specific deployed contract.
func NewEmailGuardian(address common.Address, backend bind.ContractBackend) (*EmailGuardian, error) {
	contract, err := bindEmailGuardian(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EmailGuardian{EmailGuardianCaller: EmailGuardianCaller{contract: contract}, EmailGuardianTransactor: EmailGuardianTransactor{contract: contract}, EmailGuardianFilterer: EmailGuardianFilterer{contract: contract}}, nil
}

// NewEmailGuardianCaller creates a new read-only instance of EmailGuardian, bound to a specific deployed contract.
func NewEmailGuardianCaller(address common.Address, caller bind.ContractCaller) (*EmailGuardianCaller, error) {
	contract, err := bindEmailGuardian(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EmailGuardianCaller{contract: contract}, nil
}

// NewEmailGuardianTransactor creates a new write-only instance of EmailGuardian, bound to a specific deployed contract.
func NewEmailGuardianTransactor(address common.Address, transactor bind.ContractTransactor) (*EmailGuardianTransactor, error) {
	contract, err := bindEmailGuardian(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EmailGuardianTransactor{contract: contract}, nil
}

// NewEmailGuardianFilterer creates a new log filterer instance of EmailGuardian, bound to a specific deployed contract.
func NewEmailGuardianFilterer(address common.Address, filterer bind.ContractFilterer) (*EmailGuardianFilterer, error) {
	contract, err := bindEmailGuardian(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EmailGuardianFilterer{contract: contract}, nil
}

// bindEmailGuardian binds a generic wrapper to an already deployed contract.
func bindEmailGuardian(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EmailGuardianABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EmailGuardian *EmailGuardianRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EmailGuardian.Contract.EmailGuardianCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EmailGuardian *EmailGuardianRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EmailGuardian.Contract.EmailGuardianTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EmailGuardian *EmailGuardianRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EmailGuardian.Contract.EmailGuardianTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EmailGuardian *EmailGuardianCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EmailGuardian.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EmailGuardian *EmailGuardianTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EmailGuardian.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EmailGuardian *EmailGuardianTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EmailGuardian.Contract.contract.Transact(opts, method, params...)
}

// Account is a free data retrieval call binding the contract method 0x882358ae.
//
// Solidity: function account(bytes32 email) view returns(address)
func (_EmailGuardian *EmailGuardianCaller) Account(opts *bind.CallOpts, email [32]byte) (common.Address, error) {
	var out []interface{}
	err := _EmailGuardian.contract.Call(opts, &out, "account", email)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Account is a free data retrieval call binding the contract method 0x882358ae.
//
// Solidity: function account(bytes32 email) view returns(address)
func (_EmailGuardian *EmailGuardianSession) Account(email [32]byte) (common.Address, error) {
	return _EmailGuardian.Contract.Account(&_EmailGuardian.CallOpts, email)
}

// Account is a free data retrieval call binding the contract method 0x882358ae.
//
// Solidity: function account(bytes32 email) view returns(address)
func (_EmailGuardian *EmailGuardianCallerSession) Account(email [32]byte) (common.Address, error) {
	return _EmailGuardian.Contract.Account(&_EmailGuardian.CallOpts, email)
}
