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

// UserOperation is an auto generated low-level Go binding around an user-defined struct.
type UserOperation struct {
	Sender               common.Address
	Nonce                *big.Int
	InitCode             []byte
	CallData             []byte
	CallGasLimit         *big.Int
	VerificationGasLimit *big.Int
	PreVerificationGas   *big.Int
	MaxFeePerGas         *big.Int
	MaxPriorityFeePerGas *big.Int
	PaymasterAndData     []byte
	Signature            []byte
}

// AccountMetaData contains all meta data concerning the Account contract.
var AccountMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"anEntryPoint\",\"type\":\"address\"},{\"internalType\":\"contractISecp256r1\",\"name\":\"aSecp256r1\",\"type\":\"address\"},{\"internalType\":\"contractEmailGuardian\",\"name\":\"anEmailGuardian\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"}],\"name\":\"AccountPendingRecovey\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"}],\"name\":\"AccountRecovered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"AccountRecoveryStopped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"}],\"name\":\"AccountResetted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"email\",\"type\":\"bytes32\"}],\"name\":\"EmailGuardianAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EmailGuardianRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIEntryPoint\",\"name\":\"entryPoint\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractISecp256r1\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractEmailGuardian\",\"name\":\"guardian\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"}],\"name\":\"P256AccountInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"addDeposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_email\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"addEmailGuardian\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"entryPoint\",\"outputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dest\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"func\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"dest\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"func\",\"type\":\"bytes[]\"}],\"name\":\"executeBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_publicKey\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingPublicKey\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"server\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"}],\"name\":\"pendingRecovery\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"publicKey\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"recovery\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removeEmailGuardian\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"}],\"name\":\"resetPublicKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stopRecovery\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"tokensReceived\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"callGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"verificationGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"missingAccountFunds\",\"type\":\"uint256\"}],\"name\":\"validateUserOp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"validationData\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"withdrawAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawDepositTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// AccountABI is the input ABI used to generate the binding from.
// Deprecated: Use AccountMetaData.ABI instead.
var AccountABI = AccountMetaData.ABI

// Account is an auto generated Go binding around an Ethereum contract.
type Account struct {
	AccountCaller     // Read-only binding to the contract
	AccountTransactor // Write-only binding to the contract
	AccountFilterer   // Log filterer for contract events
}

// AccountCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccountCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccountTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AccountFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AccountSession struct {
	Contract     *Account          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AccountCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AccountCallerSession struct {
	Contract *AccountCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// AccountTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AccountTransactorSession struct {
	Contract     *AccountTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AccountRaw is an auto generated low-level Go binding around an Ethereum contract.
type AccountRaw struct {
	Contract *Account // Generic contract binding to access the raw methods on
}

// AccountCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AccountCallerRaw struct {
	Contract *AccountCaller // Generic read-only contract binding to access the raw methods on
}

// AccountTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AccountTransactorRaw struct {
	Contract *AccountTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAccount creates a new instance of Account, bound to a specific deployed contract.
func NewAccount(address common.Address, backend bind.ContractBackend) (*Account, error) {
	contract, err := bindAccount(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Account{AccountCaller: AccountCaller{contract: contract}, AccountTransactor: AccountTransactor{contract: contract}, AccountFilterer: AccountFilterer{contract: contract}}, nil
}

// NewAccountCaller creates a new read-only instance of Account, bound to a specific deployed contract.
func NewAccountCaller(address common.Address, caller bind.ContractCaller) (*AccountCaller, error) {
	contract, err := bindAccount(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccountCaller{contract: contract}, nil
}

// NewAccountTransactor creates a new write-only instance of Account, bound to a specific deployed contract.
func NewAccountTransactor(address common.Address, transactor bind.ContractTransactor) (*AccountTransactor, error) {
	contract, err := bindAccount(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccountTransactor{contract: contract}, nil
}

// NewAccountFilterer creates a new log filterer instance of Account, bound to a specific deployed contract.
func NewAccountFilterer(address common.Address, filterer bind.ContractFilterer) (*AccountFilterer, error) {
	contract, err := bindAccount(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccountFilterer{contract: contract}, nil
}

// bindAccount binds a generic wrapper to an already deployed contract.
func bindAccount(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AccountABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Account *AccountRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Account.Contract.AccountCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Account *AccountRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Account.Contract.AccountTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Account *AccountRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Account.Contract.AccountTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Account *AccountCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Account.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Account *AccountTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Account.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Account *AccountTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Account.Contract.contract.Transact(opts, method, params...)
}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_Account *AccountCaller) EntryPoint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Account.contract.Call(opts, &out, "entryPoint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_Account *AccountSession) EntryPoint() (common.Address, error) {
	return _Account.Contract.EntryPoint(&_Account.CallOpts)
}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_Account *AccountCallerSession) EntryPoint() (common.Address, error) {
	return _Account.Contract.EntryPoint(&_Account.CallOpts)
}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_Account *AccountCaller) GetDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Account.contract.Call(opts, &out, "getDeposit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_Account *AccountSession) GetDeposit() (*big.Int, error) {
	return _Account.Contract.GetDeposit(&_Account.CallOpts)
}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_Account *AccountCallerSession) GetDeposit() (*big.Int, error) {
	return _Account.Contract.GetDeposit(&_Account.CallOpts)
}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(uint256)
func (_Account *AccountCaller) GetNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Account.contract.Call(opts, &out, "getNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(uint256)
func (_Account *AccountSession) GetNonce() (*big.Int, error) {
	return _Account.Contract.GetNonce(&_Account.CallOpts)
}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(uint256)
func (_Account *AccountCallerSession) GetNonce() (*big.Int, error) {
	return _Account.Contract.GetNonce(&_Account.CallOpts)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_Account *AccountCaller) OnERC1155BatchReceived(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	var out []interface{}
	err := _Account.contract.Call(opts, &out, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_Account *AccountSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	return _Account.Contract.OnERC1155BatchReceived(&_Account.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_Account *AccountCallerSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	return _Account.Contract.OnERC1155BatchReceived(&_Account.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_Account *AccountCaller) OnERC1155Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	var out []interface{}
	err := _Account.contract.Call(opts, &out, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_Account *AccountSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	return _Account.Contract.OnERC1155Received(&_Account.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_Account *AccountCallerSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	return _Account.Contract.OnERC1155Received(&_Account.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_Account *AccountCaller) OnERC721Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	var out []interface{}
	err := _Account.contract.Call(opts, &out, "onERC721Received", arg0, arg1, arg2, arg3)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_Account *AccountSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _Account.Contract.OnERC721Received(&_Account.CallOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_Account *AccountCallerSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _Account.Contract.OnERC721Received(&_Account.CallOpts, arg0, arg1, arg2, arg3)
}

// PendingPublicKey is a free data retrieval call binding the contract method 0xb4de27dd.
//
// Solidity: function pendingPublicKey() view returns(uint256 timestamp, bytes publicKey)
func (_Account *AccountCaller) PendingPublicKey(opts *bind.CallOpts) (struct {
	Timestamp *big.Int
	PublicKey []byte
}, error) {
	var out []interface{}
	err := _Account.contract.Call(opts, &out, "pendingPublicKey")

	outstruct := new(struct {
		Timestamp *big.Int
		PublicKey []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Timestamp = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.PublicKey = *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return *outstruct, err

}

// PendingPublicKey is a free data retrieval call binding the contract method 0xb4de27dd.
//
// Solidity: function pendingPublicKey() view returns(uint256 timestamp, bytes publicKey)
func (_Account *AccountSession) PendingPublicKey() (struct {
	Timestamp *big.Int
	PublicKey []byte
}, error) {
	return _Account.Contract.PendingPublicKey(&_Account.CallOpts)
}

// PendingPublicKey is a free data retrieval call binding the contract method 0xb4de27dd.
//
// Solidity: function pendingPublicKey() view returns(uint256 timestamp, bytes publicKey)
func (_Account *AccountCallerSession) PendingPublicKey() (struct {
	Timestamp *big.Int
	PublicKey []byte
}, error) {
	return _Account.Contract.PendingPublicKey(&_Account.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Account *AccountCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Account.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Account *AccountSession) ProxiableUUID() ([32]byte, error) {
	return _Account.Contract.ProxiableUUID(&_Account.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Account *AccountCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Account.Contract.ProxiableUUID(&_Account.CallOpts)
}

// PublicKey is a free data retrieval call binding the contract method 0x63ffab31.
//
// Solidity: function publicKey() view returns(bytes)
func (_Account *AccountCaller) PublicKey(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Account.contract.Call(opts, &out, "publicKey")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// PublicKey is a free data retrieval call binding the contract method 0x63ffab31.
//
// Solidity: function publicKey() view returns(bytes)
func (_Account *AccountSession) PublicKey() ([]byte, error) {
	return _Account.Contract.PublicKey(&_Account.CallOpts)
}

// PublicKey is a free data retrieval call binding the contract method 0x63ffab31.
//
// Solidity: function publicKey() view returns(bytes)
func (_Account *AccountCallerSession) PublicKey() ([]byte, error) {
	return _Account.Contract.PublicKey(&_Account.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Account *AccountCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Account.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Account *AccountSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Account.Contract.SupportsInterface(&_Account.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Account *AccountCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Account.Contract.SupportsInterface(&_Account.CallOpts, interfaceId)
}

// TokensReceived is a free data retrieval call binding the contract method 0x0023de29.
//
// Solidity: function tokensReceived(address , address , address , uint256 , bytes , bytes ) pure returns()
func (_Account *AccountCaller) TokensReceived(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 common.Address, arg3 *big.Int, arg4 []byte, arg5 []byte) error {
	var out []interface{}
	err := _Account.contract.Call(opts, &out, "tokensReceived", arg0, arg1, arg2, arg3, arg4, arg5)

	if err != nil {
		return err
	}

	return err

}

// TokensReceived is a free data retrieval call binding the contract method 0x0023de29.
//
// Solidity: function tokensReceived(address , address , address , uint256 , bytes , bytes ) pure returns()
func (_Account *AccountSession) TokensReceived(arg0 common.Address, arg1 common.Address, arg2 common.Address, arg3 *big.Int, arg4 []byte, arg5 []byte) error {
	return _Account.Contract.TokensReceived(&_Account.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// TokensReceived is a free data retrieval call binding the contract method 0x0023de29.
//
// Solidity: function tokensReceived(address , address , address , uint256 , bytes , bytes ) pure returns()
func (_Account *AccountCallerSession) TokensReceived(arg0 common.Address, arg1 common.Address, arg2 common.Address, arg3 *big.Int, arg4 []byte, arg5 []byte) error {
	return _Account.Contract.TokensReceived(&_Account.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// AddDeposit is a paid mutator transaction binding the contract method 0x4a58db19.
//
// Solidity: function addDeposit() payable returns()
func (_Account *AccountTransactor) AddDeposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Account.contract.Transact(opts, "addDeposit")
}

// AddDeposit is a paid mutator transaction binding the contract method 0x4a58db19.
//
// Solidity: function addDeposit() payable returns()
func (_Account *AccountSession) AddDeposit() (*types.Transaction, error) {
	return _Account.Contract.AddDeposit(&_Account.TransactOpts)
}

// AddDeposit is a paid mutator transaction binding the contract method 0x4a58db19.
//
// Solidity: function addDeposit() payable returns()
func (_Account *AccountTransactorSession) AddDeposit() (*types.Transaction, error) {
	return _Account.Contract.AddDeposit(&_Account.TransactOpts)
}

// AddEmailGuardian is a paid mutator transaction binding the contract method 0x7c5b84e7.
//
// Solidity: function addEmailGuardian(bytes32 _email, bytes _signature) returns()
func (_Account *AccountTransactor) AddEmailGuardian(opts *bind.TransactOpts, _email [32]byte, _signature []byte) (*types.Transaction, error) {
	return _Account.contract.Transact(opts, "addEmailGuardian", _email, _signature)
}

// AddEmailGuardian is a paid mutator transaction binding the contract method 0x7c5b84e7.
//
// Solidity: function addEmailGuardian(bytes32 _email, bytes _signature) returns()
func (_Account *AccountSession) AddEmailGuardian(_email [32]byte, _signature []byte) (*types.Transaction, error) {
	return _Account.Contract.AddEmailGuardian(&_Account.TransactOpts, _email, _signature)
}

// AddEmailGuardian is a paid mutator transaction binding the contract method 0x7c5b84e7.
//
// Solidity: function addEmailGuardian(bytes32 _email, bytes _signature) returns()
func (_Account *AccountTransactorSession) AddEmailGuardian(_email [32]byte, _signature []byte) (*types.Transaction, error) {
	return _Account.Contract.AddEmailGuardian(&_Account.TransactOpts, _email, _signature)
}

// Execute is a paid mutator transaction binding the contract method 0xb61d27f6.
//
// Solidity: function execute(address dest, uint256 value, bytes func) returns()
func (_Account *AccountTransactor) Execute(opts *bind.TransactOpts, dest common.Address, value *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _Account.contract.Transact(opts, "execute", dest, value, arg2)
}

// Execute is a paid mutator transaction binding the contract method 0xb61d27f6.
//
// Solidity: function execute(address dest, uint256 value, bytes func) returns()
func (_Account *AccountSession) Execute(dest common.Address, value *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _Account.Contract.Execute(&_Account.TransactOpts, dest, value, arg2)
}

// Execute is a paid mutator transaction binding the contract method 0xb61d27f6.
//
// Solidity: function execute(address dest, uint256 value, bytes func) returns()
func (_Account *AccountTransactorSession) Execute(dest common.Address, value *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _Account.Contract.Execute(&_Account.TransactOpts, dest, value, arg2)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0x47e1da2a.
//
// Solidity: function executeBatch(address[] dest, uint256[] values, bytes[] func) returns()
func (_Account *AccountTransactor) ExecuteBatch(opts *bind.TransactOpts, dest []common.Address, values []*big.Int, arg2 [][]byte) (*types.Transaction, error) {
	return _Account.contract.Transact(opts, "executeBatch", dest, values, arg2)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0x47e1da2a.
//
// Solidity: function executeBatch(address[] dest, uint256[] values, bytes[] func) returns()
func (_Account *AccountSession) ExecuteBatch(dest []common.Address, values []*big.Int, arg2 [][]byte) (*types.Transaction, error) {
	return _Account.Contract.ExecuteBatch(&_Account.TransactOpts, dest, values, arg2)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0x47e1da2a.
//
// Solidity: function executeBatch(address[] dest, uint256[] values, bytes[] func) returns()
func (_Account *AccountTransactorSession) ExecuteBatch(dest []common.Address, values []*big.Int, arg2 [][]byte) (*types.Transaction, error) {
	return _Account.Contract.ExecuteBatch(&_Account.TransactOpts, dest, values, arg2)
}

// Initialize is a paid mutator transaction binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes _publicKey) returns()
func (_Account *AccountTransactor) Initialize(opts *bind.TransactOpts, _publicKey []byte) (*types.Transaction, error) {
	return _Account.contract.Transact(opts, "initialize", _publicKey)
}

// Initialize is a paid mutator transaction binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes _publicKey) returns()
func (_Account *AccountSession) Initialize(_publicKey []byte) (*types.Transaction, error) {
	return _Account.Contract.Initialize(&_Account.TransactOpts, _publicKey)
}

// Initialize is a paid mutator transaction binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes _publicKey) returns()
func (_Account *AccountTransactorSession) Initialize(_publicKey []byte) (*types.Transaction, error) {
	return _Account.Contract.Initialize(&_Account.TransactOpts, _publicKey)
}

// PendingRecovery is a paid mutator transaction binding the contract method 0x5a91d239.
//
// Solidity: function pendingRecovery(bytes32 server, bytes data, bytes signature, bytes pubkey) returns()
func (_Account *AccountTransactor) PendingRecovery(opts *bind.TransactOpts, server [32]byte, data []byte, signature []byte, pubkey []byte) (*types.Transaction, error) {
	return _Account.contract.Transact(opts, "pendingRecovery", server, data, signature, pubkey)
}

// PendingRecovery is a paid mutator transaction binding the contract method 0x5a91d239.
//
// Solidity: function pendingRecovery(bytes32 server, bytes data, bytes signature, bytes pubkey) returns()
func (_Account *AccountSession) PendingRecovery(server [32]byte, data []byte, signature []byte, pubkey []byte) (*types.Transaction, error) {
	return _Account.Contract.PendingRecovery(&_Account.TransactOpts, server, data, signature, pubkey)
}

// PendingRecovery is a paid mutator transaction binding the contract method 0x5a91d239.
//
// Solidity: function pendingRecovery(bytes32 server, bytes data, bytes signature, bytes pubkey) returns()
func (_Account *AccountTransactorSession) PendingRecovery(server [32]byte, data []byte, signature []byte, pubkey []byte) (*types.Transaction, error) {
	return _Account.Contract.PendingRecovery(&_Account.TransactOpts, server, data, signature, pubkey)
}

// Recovery is a paid mutator transaction binding the contract method 0xddceafa9.
//
// Solidity: function recovery() returns()
func (_Account *AccountTransactor) Recovery(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Account.contract.Transact(opts, "recovery")
}

// Recovery is a paid mutator transaction binding the contract method 0xddceafa9.
//
// Solidity: function recovery() returns()
func (_Account *AccountSession) Recovery() (*types.Transaction, error) {
	return _Account.Contract.Recovery(&_Account.TransactOpts)
}

// Recovery is a paid mutator transaction binding the contract method 0xddceafa9.
//
// Solidity: function recovery() returns()
func (_Account *AccountTransactorSession) Recovery() (*types.Transaction, error) {
	return _Account.Contract.Recovery(&_Account.TransactOpts)
}

// RemoveEmailGuardian is a paid mutator transaction binding the contract method 0x1e5f36b3.
//
// Solidity: function removeEmailGuardian() returns()
func (_Account *AccountTransactor) RemoveEmailGuardian(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Account.contract.Transact(opts, "removeEmailGuardian")
}

// RemoveEmailGuardian is a paid mutator transaction binding the contract method 0x1e5f36b3.
//
// Solidity: function removeEmailGuardian() returns()
func (_Account *AccountSession) RemoveEmailGuardian() (*types.Transaction, error) {
	return _Account.Contract.RemoveEmailGuardian(&_Account.TransactOpts)
}

// RemoveEmailGuardian is a paid mutator transaction binding the contract method 0x1e5f36b3.
//
// Solidity: function removeEmailGuardian() returns()
func (_Account *AccountTransactorSession) RemoveEmailGuardian() (*types.Transaction, error) {
	return _Account.Contract.RemoveEmailGuardian(&_Account.TransactOpts)
}

// ResetPublicKey is a paid mutator transaction binding the contract method 0xa08e18f4.
//
// Solidity: function resetPublicKey(bytes pubkey) returns()
func (_Account *AccountTransactor) ResetPublicKey(opts *bind.TransactOpts, pubkey []byte) (*types.Transaction, error) {
	return _Account.contract.Transact(opts, "resetPublicKey", pubkey)
}

// ResetPublicKey is a paid mutator transaction binding the contract method 0xa08e18f4.
//
// Solidity: function resetPublicKey(bytes pubkey) returns()
func (_Account *AccountSession) ResetPublicKey(pubkey []byte) (*types.Transaction, error) {
	return _Account.Contract.ResetPublicKey(&_Account.TransactOpts, pubkey)
}

// ResetPublicKey is a paid mutator transaction binding the contract method 0xa08e18f4.
//
// Solidity: function resetPublicKey(bytes pubkey) returns()
func (_Account *AccountTransactorSession) ResetPublicKey(pubkey []byte) (*types.Transaction, error) {
	return _Account.Contract.ResetPublicKey(&_Account.TransactOpts, pubkey)
}

// StopRecovery is a paid mutator transaction binding the contract method 0xbbd855f8.
//
// Solidity: function stopRecovery() returns()
func (_Account *AccountTransactor) StopRecovery(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Account.contract.Transact(opts, "stopRecovery")
}

// StopRecovery is a paid mutator transaction binding the contract method 0xbbd855f8.
//
// Solidity: function stopRecovery() returns()
func (_Account *AccountSession) StopRecovery() (*types.Transaction, error) {
	return _Account.Contract.StopRecovery(&_Account.TransactOpts)
}

// StopRecovery is a paid mutator transaction binding the contract method 0xbbd855f8.
//
// Solidity: function stopRecovery() returns()
func (_Account *AccountTransactorSession) StopRecovery() (*types.Transaction, error) {
	return _Account.Contract.StopRecovery(&_Account.TransactOpts)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Account *AccountTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Account.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Account *AccountSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Account.Contract.UpgradeTo(&_Account.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Account *AccountTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Account.Contract.UpgradeTo(&_Account.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Account *AccountTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Account.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Account *AccountSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Account.Contract.UpgradeToAndCall(&_Account.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Account *AccountTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Account.Contract.UpgradeToAndCall(&_Account.TransactOpts, newImplementation, data)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x3a871cdd.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp, bytes32 userOpHash, uint256 missingAccountFunds) returns(uint256 validationData)
func (_Account *AccountTransactor) ValidateUserOp(opts *bind.TransactOpts, userOp UserOperation, userOpHash [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _Account.contract.Transact(opts, "validateUserOp", userOp, userOpHash, missingAccountFunds)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x3a871cdd.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp, bytes32 userOpHash, uint256 missingAccountFunds) returns(uint256 validationData)
func (_Account *AccountSession) ValidateUserOp(userOp UserOperation, userOpHash [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _Account.Contract.ValidateUserOp(&_Account.TransactOpts, userOp, userOpHash, missingAccountFunds)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x3a871cdd.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp, bytes32 userOpHash, uint256 missingAccountFunds) returns(uint256 validationData)
func (_Account *AccountTransactorSession) ValidateUserOp(userOp UserOperation, userOpHash [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _Account.Contract.ValidateUserOp(&_Account.TransactOpts, userOp, userOpHash, missingAccountFunds)
}

// WithdrawDepositTo is a paid mutator transaction binding the contract method 0x4d44560d.
//
// Solidity: function withdrawDepositTo(address withdrawAddress, uint256 amount) returns()
func (_Account *AccountTransactor) WithdrawDepositTo(opts *bind.TransactOpts, withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Account.contract.Transact(opts, "withdrawDepositTo", withdrawAddress, amount)
}

// WithdrawDepositTo is a paid mutator transaction binding the contract method 0x4d44560d.
//
// Solidity: function withdrawDepositTo(address withdrawAddress, uint256 amount) returns()
func (_Account *AccountSession) WithdrawDepositTo(withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Account.Contract.WithdrawDepositTo(&_Account.TransactOpts, withdrawAddress, amount)
}

// WithdrawDepositTo is a paid mutator transaction binding the contract method 0x4d44560d.
//
// Solidity: function withdrawDepositTo(address withdrawAddress, uint256 amount) returns()
func (_Account *AccountTransactorSession) WithdrawDepositTo(withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Account.Contract.WithdrawDepositTo(&_Account.TransactOpts, withdrawAddress, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Account *AccountTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Account.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Account *AccountSession) Receive() (*types.Transaction, error) {
	return _Account.Contract.Receive(&_Account.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Account *AccountTransactorSession) Receive() (*types.Transaction, error) {
	return _Account.Contract.Receive(&_Account.TransactOpts)
}

// AccountAccountPendingRecoveyIterator is returned from FilterAccountPendingRecovey and is used to iterate over the raw logs and unpacked data for AccountPendingRecovey events raised by the Account contract.
type AccountAccountPendingRecoveyIterator struct {
	Event *AccountAccountPendingRecovey // Event containing the contract specifics and raw log

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
func (it *AccountAccountPendingRecoveyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountAccountPendingRecovey)
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
		it.Event = new(AccountAccountPendingRecovey)
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
func (it *AccountAccountPendingRecoveyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountAccountPendingRecoveyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountAccountPendingRecovey represents a AccountPendingRecovey event raised by the Account contract.
type AccountAccountPendingRecovey struct {
	Timestamp *big.Int
	PublicKey []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAccountPendingRecovey is a free log retrieval operation binding the contract event 0x4ea6e027fa89115fc2e5188f9665f4755a8edd56443147876507e1fd1db35eac.
//
// Solidity: event AccountPendingRecovey(uint256 timestamp, bytes publicKey)
func (_Account *AccountFilterer) FilterAccountPendingRecovey(opts *bind.FilterOpts) (*AccountAccountPendingRecoveyIterator, error) {

	logs, sub, err := _Account.contract.FilterLogs(opts, "AccountPendingRecovey")
	if err != nil {
		return nil, err
	}
	return &AccountAccountPendingRecoveyIterator{contract: _Account.contract, event: "AccountPendingRecovey", logs: logs, sub: sub}, nil
}

// WatchAccountPendingRecovey is a free log subscription operation binding the contract event 0x4ea6e027fa89115fc2e5188f9665f4755a8edd56443147876507e1fd1db35eac.
//
// Solidity: event AccountPendingRecovey(uint256 timestamp, bytes publicKey)
func (_Account *AccountFilterer) WatchAccountPendingRecovey(opts *bind.WatchOpts, sink chan<- *AccountAccountPendingRecovey) (event.Subscription, error) {

	logs, sub, err := _Account.contract.WatchLogs(opts, "AccountPendingRecovey")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountAccountPendingRecovey)
				if err := _Account.contract.UnpackLog(event, "AccountPendingRecovey", log); err != nil {
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

// ParseAccountPendingRecovey is a log parse operation binding the contract event 0x4ea6e027fa89115fc2e5188f9665f4755a8edd56443147876507e1fd1db35eac.
//
// Solidity: event AccountPendingRecovey(uint256 timestamp, bytes publicKey)
func (_Account *AccountFilterer) ParseAccountPendingRecovey(log types.Log) (*AccountAccountPendingRecovey, error) {
	event := new(AccountAccountPendingRecovey)
	if err := _Account.contract.UnpackLog(event, "AccountPendingRecovey", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountAccountRecoveredIterator is returned from FilterAccountRecovered and is used to iterate over the raw logs and unpacked data for AccountRecovered events raised by the Account contract.
type AccountAccountRecoveredIterator struct {
	Event *AccountAccountRecovered // Event containing the contract specifics and raw log

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
func (it *AccountAccountRecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountAccountRecovered)
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
		it.Event = new(AccountAccountRecovered)
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
func (it *AccountAccountRecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountAccountRecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountAccountRecovered represents a AccountRecovered event raised by the Account contract.
type AccountAccountRecovered struct {
	PublicKey []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAccountRecovered is a free log retrieval operation binding the contract event 0x7dc0147d22488d2f12c2c4b776f1bc504e5819e683d3f35d83bfc84eed73dc10.
//
// Solidity: event AccountRecovered(bytes publicKey)
func (_Account *AccountFilterer) FilterAccountRecovered(opts *bind.FilterOpts) (*AccountAccountRecoveredIterator, error) {

	logs, sub, err := _Account.contract.FilterLogs(opts, "AccountRecovered")
	if err != nil {
		return nil, err
	}
	return &AccountAccountRecoveredIterator{contract: _Account.contract, event: "AccountRecovered", logs: logs, sub: sub}, nil
}

// WatchAccountRecovered is a free log subscription operation binding the contract event 0x7dc0147d22488d2f12c2c4b776f1bc504e5819e683d3f35d83bfc84eed73dc10.
//
// Solidity: event AccountRecovered(bytes publicKey)
func (_Account *AccountFilterer) WatchAccountRecovered(opts *bind.WatchOpts, sink chan<- *AccountAccountRecovered) (event.Subscription, error) {

	logs, sub, err := _Account.contract.WatchLogs(opts, "AccountRecovered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountAccountRecovered)
				if err := _Account.contract.UnpackLog(event, "AccountRecovered", log); err != nil {
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

// ParseAccountRecovered is a log parse operation binding the contract event 0x7dc0147d22488d2f12c2c4b776f1bc504e5819e683d3f35d83bfc84eed73dc10.
//
// Solidity: event AccountRecovered(bytes publicKey)
func (_Account *AccountFilterer) ParseAccountRecovered(log types.Log) (*AccountAccountRecovered, error) {
	event := new(AccountAccountRecovered)
	if err := _Account.contract.UnpackLog(event, "AccountRecovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountAccountRecoveryStoppedIterator is returned from FilterAccountRecoveryStopped and is used to iterate over the raw logs and unpacked data for AccountRecoveryStopped events raised by the Account contract.
type AccountAccountRecoveryStoppedIterator struct {
	Event *AccountAccountRecoveryStopped // Event containing the contract specifics and raw log

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
func (it *AccountAccountRecoveryStoppedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountAccountRecoveryStopped)
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
		it.Event = new(AccountAccountRecoveryStopped)
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
func (it *AccountAccountRecoveryStoppedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountAccountRecoveryStoppedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountAccountRecoveryStopped represents a AccountRecoveryStopped event raised by the Account contract.
type AccountAccountRecoveryStopped struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAccountRecoveryStopped is a free log retrieval operation binding the contract event 0x4c53d9c9c5e0b9bcd418f17f35450c36a222a35a337f0b3b5b2b83b757d59154.
//
// Solidity: event AccountRecoveryStopped()
func (_Account *AccountFilterer) FilterAccountRecoveryStopped(opts *bind.FilterOpts) (*AccountAccountRecoveryStoppedIterator, error) {

	logs, sub, err := _Account.contract.FilterLogs(opts, "AccountRecoveryStopped")
	if err != nil {
		return nil, err
	}
	return &AccountAccountRecoveryStoppedIterator{contract: _Account.contract, event: "AccountRecoveryStopped", logs: logs, sub: sub}, nil
}

// WatchAccountRecoveryStopped is a free log subscription operation binding the contract event 0x4c53d9c9c5e0b9bcd418f17f35450c36a222a35a337f0b3b5b2b83b757d59154.
//
// Solidity: event AccountRecoveryStopped()
func (_Account *AccountFilterer) WatchAccountRecoveryStopped(opts *bind.WatchOpts, sink chan<- *AccountAccountRecoveryStopped) (event.Subscription, error) {

	logs, sub, err := _Account.contract.WatchLogs(opts, "AccountRecoveryStopped")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountAccountRecoveryStopped)
				if err := _Account.contract.UnpackLog(event, "AccountRecoveryStopped", log); err != nil {
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

// ParseAccountRecoveryStopped is a log parse operation binding the contract event 0x4c53d9c9c5e0b9bcd418f17f35450c36a222a35a337f0b3b5b2b83b757d59154.
//
// Solidity: event AccountRecoveryStopped()
func (_Account *AccountFilterer) ParseAccountRecoveryStopped(log types.Log) (*AccountAccountRecoveryStopped, error) {
	event := new(AccountAccountRecoveryStopped)
	if err := _Account.contract.UnpackLog(event, "AccountRecoveryStopped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountAccountResettedIterator is returned from FilterAccountResetted and is used to iterate over the raw logs and unpacked data for AccountResetted events raised by the Account contract.
type AccountAccountResettedIterator struct {
	Event *AccountAccountResetted // Event containing the contract specifics and raw log

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
func (it *AccountAccountResettedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountAccountResetted)
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
		it.Event = new(AccountAccountResetted)
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
func (it *AccountAccountResettedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountAccountResettedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountAccountResetted represents a AccountResetted event raised by the Account contract.
type AccountAccountResetted struct {
	PublicKey []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAccountResetted is a free log retrieval operation binding the contract event 0x70afa444db04f5ca57be5c4925d02957c3f573f55430e6e899b47d1228060262.
//
// Solidity: event AccountResetted(bytes publicKey)
func (_Account *AccountFilterer) FilterAccountResetted(opts *bind.FilterOpts) (*AccountAccountResettedIterator, error) {

	logs, sub, err := _Account.contract.FilterLogs(opts, "AccountResetted")
	if err != nil {
		return nil, err
	}
	return &AccountAccountResettedIterator{contract: _Account.contract, event: "AccountResetted", logs: logs, sub: sub}, nil
}

// WatchAccountResetted is a free log subscription operation binding the contract event 0x70afa444db04f5ca57be5c4925d02957c3f573f55430e6e899b47d1228060262.
//
// Solidity: event AccountResetted(bytes publicKey)
func (_Account *AccountFilterer) WatchAccountResetted(opts *bind.WatchOpts, sink chan<- *AccountAccountResetted) (event.Subscription, error) {

	logs, sub, err := _Account.contract.WatchLogs(opts, "AccountResetted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountAccountResetted)
				if err := _Account.contract.UnpackLog(event, "AccountResetted", log); err != nil {
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

// ParseAccountResetted is a log parse operation binding the contract event 0x70afa444db04f5ca57be5c4925d02957c3f573f55430e6e899b47d1228060262.
//
// Solidity: event AccountResetted(bytes publicKey)
func (_Account *AccountFilterer) ParseAccountResetted(log types.Log) (*AccountAccountResetted, error) {
	event := new(AccountAccountResetted)
	if err := _Account.contract.UnpackLog(event, "AccountResetted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Account contract.
type AccountAdminChangedIterator struct {
	Event *AccountAdminChanged // Event containing the contract specifics and raw log

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
func (it *AccountAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountAdminChanged)
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
		it.Event = new(AccountAdminChanged)
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
func (it *AccountAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountAdminChanged represents a AdminChanged event raised by the Account contract.
type AccountAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Account *AccountFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*AccountAdminChangedIterator, error) {

	logs, sub, err := _Account.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &AccountAdminChangedIterator{contract: _Account.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Account *AccountFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *AccountAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Account.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountAdminChanged)
				if err := _Account.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Account *AccountFilterer) ParseAdminChanged(log types.Log) (*AccountAdminChanged, error) {
	event := new(AccountAdminChanged)
	if err := _Account.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the Account contract.
type AccountBeaconUpgradedIterator struct {
	Event *AccountBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *AccountBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountBeaconUpgraded)
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
		it.Event = new(AccountBeaconUpgraded)
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
func (it *AccountBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountBeaconUpgraded represents a BeaconUpgraded event raised by the Account contract.
type AccountBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Account *AccountFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*AccountBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Account.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &AccountBeaconUpgradedIterator{contract: _Account.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Account *AccountFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *AccountBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Account.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountBeaconUpgraded)
				if err := _Account.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Account *AccountFilterer) ParseBeaconUpgraded(log types.Log) (*AccountBeaconUpgraded, error) {
	event := new(AccountBeaconUpgraded)
	if err := _Account.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountEmailGuardianAddedIterator is returned from FilterEmailGuardianAdded and is used to iterate over the raw logs and unpacked data for EmailGuardianAdded events raised by the Account contract.
type AccountEmailGuardianAddedIterator struct {
	Event *AccountEmailGuardianAdded // Event containing the contract specifics and raw log

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
func (it *AccountEmailGuardianAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountEmailGuardianAdded)
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
		it.Event = new(AccountEmailGuardianAdded)
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
func (it *AccountEmailGuardianAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountEmailGuardianAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountEmailGuardianAdded represents a EmailGuardianAdded event raised by the Account contract.
type AccountEmailGuardianAdded struct {
	Email [32]byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterEmailGuardianAdded is a free log retrieval operation binding the contract event 0xd8adaafedab553012b2226f396e06cbdc69b9391df6050ac3ccac863bd4e9467.
//
// Solidity: event EmailGuardianAdded(bytes32 email)
func (_Account *AccountFilterer) FilterEmailGuardianAdded(opts *bind.FilterOpts) (*AccountEmailGuardianAddedIterator, error) {

	logs, sub, err := _Account.contract.FilterLogs(opts, "EmailGuardianAdded")
	if err != nil {
		return nil, err
	}
	return &AccountEmailGuardianAddedIterator{contract: _Account.contract, event: "EmailGuardianAdded", logs: logs, sub: sub}, nil
}

// WatchEmailGuardianAdded is a free log subscription operation binding the contract event 0xd8adaafedab553012b2226f396e06cbdc69b9391df6050ac3ccac863bd4e9467.
//
// Solidity: event EmailGuardianAdded(bytes32 email)
func (_Account *AccountFilterer) WatchEmailGuardianAdded(opts *bind.WatchOpts, sink chan<- *AccountEmailGuardianAdded) (event.Subscription, error) {

	logs, sub, err := _Account.contract.WatchLogs(opts, "EmailGuardianAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountEmailGuardianAdded)
				if err := _Account.contract.UnpackLog(event, "EmailGuardianAdded", log); err != nil {
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

// ParseEmailGuardianAdded is a log parse operation binding the contract event 0xd8adaafedab553012b2226f396e06cbdc69b9391df6050ac3ccac863bd4e9467.
//
// Solidity: event EmailGuardianAdded(bytes32 email)
func (_Account *AccountFilterer) ParseEmailGuardianAdded(log types.Log) (*AccountEmailGuardianAdded, error) {
	event := new(AccountEmailGuardianAdded)
	if err := _Account.contract.UnpackLog(event, "EmailGuardianAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountEmailGuardianRemovedIterator is returned from FilterEmailGuardianRemoved and is used to iterate over the raw logs and unpacked data for EmailGuardianRemoved events raised by the Account contract.
type AccountEmailGuardianRemovedIterator struct {
	Event *AccountEmailGuardianRemoved // Event containing the contract specifics and raw log

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
func (it *AccountEmailGuardianRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountEmailGuardianRemoved)
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
		it.Event = new(AccountEmailGuardianRemoved)
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
func (it *AccountEmailGuardianRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountEmailGuardianRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountEmailGuardianRemoved represents a EmailGuardianRemoved event raised by the Account contract.
type AccountEmailGuardianRemoved struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEmailGuardianRemoved is a free log retrieval operation binding the contract event 0x131e7b86b68982bc41e68953e2069d8ce9de299e35ec2bc90b27949aabced73d.
//
// Solidity: event EmailGuardianRemoved()
func (_Account *AccountFilterer) FilterEmailGuardianRemoved(opts *bind.FilterOpts) (*AccountEmailGuardianRemovedIterator, error) {

	logs, sub, err := _Account.contract.FilterLogs(opts, "EmailGuardianRemoved")
	if err != nil {
		return nil, err
	}
	return &AccountEmailGuardianRemovedIterator{contract: _Account.contract, event: "EmailGuardianRemoved", logs: logs, sub: sub}, nil
}

// WatchEmailGuardianRemoved is a free log subscription operation binding the contract event 0x131e7b86b68982bc41e68953e2069d8ce9de299e35ec2bc90b27949aabced73d.
//
// Solidity: event EmailGuardianRemoved()
func (_Account *AccountFilterer) WatchEmailGuardianRemoved(opts *bind.WatchOpts, sink chan<- *AccountEmailGuardianRemoved) (event.Subscription, error) {

	logs, sub, err := _Account.contract.WatchLogs(opts, "EmailGuardianRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountEmailGuardianRemoved)
				if err := _Account.contract.UnpackLog(event, "EmailGuardianRemoved", log); err != nil {
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

// ParseEmailGuardianRemoved is a log parse operation binding the contract event 0x131e7b86b68982bc41e68953e2069d8ce9de299e35ec2bc90b27949aabced73d.
//
// Solidity: event EmailGuardianRemoved()
func (_Account *AccountFilterer) ParseEmailGuardianRemoved(log types.Log) (*AccountEmailGuardianRemoved, error) {
	event := new(AccountEmailGuardianRemoved)
	if err := _Account.contract.UnpackLog(event, "EmailGuardianRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Account contract.
type AccountInitializedIterator struct {
	Event *AccountInitialized // Event containing the contract specifics and raw log

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
func (it *AccountInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountInitialized)
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
		it.Event = new(AccountInitialized)
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
func (it *AccountInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountInitialized represents a Initialized event raised by the Account contract.
type AccountInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Account *AccountFilterer) FilterInitialized(opts *bind.FilterOpts) (*AccountInitializedIterator, error) {

	logs, sub, err := _Account.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AccountInitializedIterator{contract: _Account.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Account *AccountFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AccountInitialized) (event.Subscription, error) {

	logs, sub, err := _Account.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountInitialized)
				if err := _Account.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Account *AccountFilterer) ParseInitialized(log types.Log) (*AccountInitialized, error) {
	event := new(AccountInitialized)
	if err := _Account.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountP256AccountInitializedIterator is returned from FilterP256AccountInitialized and is used to iterate over the raw logs and unpacked data for P256AccountInitialized events raised by the Account contract.
type AccountP256AccountInitializedIterator struct {
	Event *AccountP256AccountInitialized // Event containing the contract specifics and raw log

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
func (it *AccountP256AccountInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountP256AccountInitialized)
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
		it.Event = new(AccountP256AccountInitialized)
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
func (it *AccountP256AccountInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountP256AccountInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountP256AccountInitialized represents a P256AccountInitialized event raised by the Account contract.
type AccountP256AccountInitialized struct {
	EntryPoint common.Address
	Validator  common.Address
	Guardian   common.Address
	PublicKey  []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterP256AccountInitialized is a free log retrieval operation binding the contract event 0x3d6f32c3e0c0a7f846ba1a822f1e8f4c5f1ec789e59f454c54deff94ccd31037.
//
// Solidity: event P256AccountInitialized(address indexed entryPoint, address validator, address guardian, bytes publicKey)
func (_Account *AccountFilterer) FilterP256AccountInitialized(opts *bind.FilterOpts, entryPoint []common.Address) (*AccountP256AccountInitializedIterator, error) {

	var entryPointRule []interface{}
	for _, entryPointItem := range entryPoint {
		entryPointRule = append(entryPointRule, entryPointItem)
	}

	logs, sub, err := _Account.contract.FilterLogs(opts, "P256AccountInitialized", entryPointRule)
	if err != nil {
		return nil, err
	}
	return &AccountP256AccountInitializedIterator{contract: _Account.contract, event: "P256AccountInitialized", logs: logs, sub: sub}, nil
}

// WatchP256AccountInitialized is a free log subscription operation binding the contract event 0x3d6f32c3e0c0a7f846ba1a822f1e8f4c5f1ec789e59f454c54deff94ccd31037.
//
// Solidity: event P256AccountInitialized(address indexed entryPoint, address validator, address guardian, bytes publicKey)
func (_Account *AccountFilterer) WatchP256AccountInitialized(opts *bind.WatchOpts, sink chan<- *AccountP256AccountInitialized, entryPoint []common.Address) (event.Subscription, error) {

	var entryPointRule []interface{}
	for _, entryPointItem := range entryPoint {
		entryPointRule = append(entryPointRule, entryPointItem)
	}

	logs, sub, err := _Account.contract.WatchLogs(opts, "P256AccountInitialized", entryPointRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountP256AccountInitialized)
				if err := _Account.contract.UnpackLog(event, "P256AccountInitialized", log); err != nil {
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

// ParseP256AccountInitialized is a log parse operation binding the contract event 0x3d6f32c3e0c0a7f846ba1a822f1e8f4c5f1ec789e59f454c54deff94ccd31037.
//
// Solidity: event P256AccountInitialized(address indexed entryPoint, address validator, address guardian, bytes publicKey)
func (_Account *AccountFilterer) ParseP256AccountInitialized(log types.Log) (*AccountP256AccountInitialized, error) {
	event := new(AccountP256AccountInitialized)
	if err := _Account.contract.UnpackLog(event, "P256AccountInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccountUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Account contract.
type AccountUpgradedIterator struct {
	Event *AccountUpgraded // Event containing the contract specifics and raw log

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
func (it *AccountUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountUpgraded)
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
		it.Event = new(AccountUpgraded)
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
func (it *AccountUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountUpgraded represents a Upgraded event raised by the Account contract.
type AccountUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Account *AccountFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*AccountUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Account.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &AccountUpgradedIterator{contract: _Account.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Account *AccountFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *AccountUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Account.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountUpgraded)
				if err := _Account.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Account *AccountFilterer) ParseUpgraded(log types.Log) (*AccountUpgraded, error) {
	event := new(AccountUpgraded)
	if err := _Account.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
