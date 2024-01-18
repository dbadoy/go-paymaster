// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

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

	"github.com/dbadoy/go-paymaster/pkg/userop"
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

// VerifiyingPaymasterMetaData contains all meta data concerning the VerifiyingPaymaster contract.
var VerifiyingPaymasterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"contractIEntryPoint\",\"name\":\"_entryPoint\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_verifyingSigner\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"CallerIsNotAnEntryPoint\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CanNotWithdrawToZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DepositCanNotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EntryPointCannotBeZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountRequired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentBalance\",\"type\":\"uint256\"}],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sigLength\",\"type\":\"uint256\"}],\"name\":\"InvalidPaymasterSignatureLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PaymasterIdCannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VerifyingSignerCannotBeZero\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_oldValue\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_newValue\",\"type\":\"uint256\"}],\"name\":\"EPGasOverheadChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_paymasterId\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_charge\",\"type\":\"uint256\"}],\"name\":\"GasBalanceDeducted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_paymasterId\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"GasDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_paymasterId\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"GasWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_oldSigner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_newSigner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_actor\",\"type\":\"address\"}],\"name\":\"VerifyingSignerChanged\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"unstakeDelaySec\",\"type\":\"uint32\"}],\"name\":\"addStake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"paymasterId\",\"type\":\"address\"}],\"name\":\"depositFor\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"entryPoint\",\"outputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"paymasterId\",\"type\":\"address\"}],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"callGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"verificationGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"paymasterId\",\"type\":\"address\"}],\"name\":\"getHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"paymasterIdBalances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumIPaymaster.PostOpMode\",\"name\":\"mode\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"context\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"actualGasCost\",\"type\":\"uint256\"}],\"name\":\"postOp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newVerifyingSigner\",\"type\":\"address\"}],\"name\":\"setSigner\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setUnaccountedEPGasOverhead\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unlockStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"callGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"verificationGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"maxCost\",\"type\":\"uint256\"}],\"name\":\"validatePaymasterUserOp\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"context\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"validationData\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verifyingSigner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"withdrawAddress\",\"type\":\"address\"}],\"name\":\"withdrawStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"withdrawAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// VerifiyingPaymasterABI is the input ABI used to generate the binding from.
// Deprecated: Use VerifiyingPaymasterMetaData.ABI instead.
var VerifiyingPaymasterABI = VerifiyingPaymasterMetaData.ABI

// VerifiyingPaymaster is an auto generated Go binding around an Ethereum contract.
type VerifiyingPaymaster struct {
	VerifiyingPaymasterCaller     // Read-only binding to the contract
	VerifiyingPaymasterTransactor // Write-only binding to the contract
	VerifiyingPaymasterFilterer   // Log filterer for contract events
}

// VerifiyingPaymasterCaller is an auto generated read-only Go binding around an Ethereum contract.
type VerifiyingPaymasterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifiyingPaymasterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VerifiyingPaymasterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifiyingPaymasterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VerifiyingPaymasterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifiyingPaymasterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VerifiyingPaymasterSession struct {
	Contract     *VerifiyingPaymaster // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// VerifiyingPaymasterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VerifiyingPaymasterCallerSession struct {
	Contract *VerifiyingPaymasterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// VerifiyingPaymasterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VerifiyingPaymasterTransactorSession struct {
	Contract     *VerifiyingPaymasterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// VerifiyingPaymasterRaw is an auto generated low-level Go binding around an Ethereum contract.
type VerifiyingPaymasterRaw struct {
	Contract *VerifiyingPaymaster // Generic contract binding to access the raw methods on
}

// VerifiyingPaymasterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VerifiyingPaymasterCallerRaw struct {
	Contract *VerifiyingPaymasterCaller // Generic read-only contract binding to access the raw methods on
}

// VerifiyingPaymasterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VerifiyingPaymasterTransactorRaw struct {
	Contract *VerifiyingPaymasterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVerifiyingPaymaster creates a new instance of VerifiyingPaymaster, bound to a specific deployed contract.
func NewVerifiyingPaymaster(address common.Address, backend bind.ContractBackend) (*VerifiyingPaymaster, error) {
	contract, err := bindVerifiyingPaymaster(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VerifiyingPaymaster{VerifiyingPaymasterCaller: VerifiyingPaymasterCaller{contract: contract}, VerifiyingPaymasterTransactor: VerifiyingPaymasterTransactor{contract: contract}, VerifiyingPaymasterFilterer: VerifiyingPaymasterFilterer{contract: contract}}, nil
}

// NewVerifiyingPaymasterCaller creates a new read-only instance of VerifiyingPaymaster, bound to a specific deployed contract.
func NewVerifiyingPaymasterCaller(address common.Address, caller bind.ContractCaller) (*VerifiyingPaymasterCaller, error) {
	contract, err := bindVerifiyingPaymaster(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VerifiyingPaymasterCaller{contract: contract}, nil
}

// NewVerifiyingPaymasterTransactor creates a new write-only instance of VerifiyingPaymaster, bound to a specific deployed contract.
func NewVerifiyingPaymasterTransactor(address common.Address, transactor bind.ContractTransactor) (*VerifiyingPaymasterTransactor, error) {
	contract, err := bindVerifiyingPaymaster(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VerifiyingPaymasterTransactor{contract: contract}, nil
}

// NewVerifiyingPaymasterFilterer creates a new log filterer instance of VerifiyingPaymaster, bound to a specific deployed contract.
func NewVerifiyingPaymasterFilterer(address common.Address, filterer bind.ContractFilterer) (*VerifiyingPaymasterFilterer, error) {
	contract, err := bindVerifiyingPaymaster(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VerifiyingPaymasterFilterer{contract: contract}, nil
}

// bindVerifiyingPaymaster binds a generic wrapper to an already deployed contract.
func bindVerifiyingPaymaster(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VerifiyingPaymasterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VerifiyingPaymaster *VerifiyingPaymasterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VerifiyingPaymaster.Contract.VerifiyingPaymasterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VerifiyingPaymaster *VerifiyingPaymasterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifiyingPaymaster.Contract.VerifiyingPaymasterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VerifiyingPaymaster *VerifiyingPaymasterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VerifiyingPaymaster.Contract.VerifiyingPaymasterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VerifiyingPaymaster *VerifiyingPaymasterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VerifiyingPaymaster.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VerifiyingPaymaster *VerifiyingPaymasterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifiyingPaymaster.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VerifiyingPaymaster *VerifiyingPaymasterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VerifiyingPaymaster.Contract.contract.Transact(opts, method, params...)
}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_VerifiyingPaymaster *VerifiyingPaymasterCaller) EntryPoint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VerifiyingPaymaster.contract.Call(opts, &out, "entryPoint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_VerifiyingPaymaster *VerifiyingPaymasterSession) EntryPoint() (common.Address, error) {
	return _VerifiyingPaymaster.Contract.EntryPoint(&_VerifiyingPaymaster.CallOpts)
}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_VerifiyingPaymaster *VerifiyingPaymasterCallerSession) EntryPoint() (common.Address, error) {
	return _VerifiyingPaymaster.Contract.EntryPoint(&_VerifiyingPaymaster.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address paymasterId) view returns(uint256 balance)
func (_VerifiyingPaymaster *VerifiyingPaymasterCaller) GetBalance(opts *bind.CallOpts, paymasterId common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VerifiyingPaymaster.contract.Call(opts, &out, "getBalance", paymasterId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address paymasterId) view returns(uint256 balance)
func (_VerifiyingPaymaster *VerifiyingPaymasterSession) GetBalance(paymasterId common.Address) (*big.Int, error) {
	return _VerifiyingPaymaster.Contract.GetBalance(&_VerifiyingPaymaster.CallOpts, paymasterId)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address paymasterId) view returns(uint256 balance)
func (_VerifiyingPaymaster *VerifiyingPaymasterCallerSession) GetBalance(paymasterId common.Address) (*big.Int, error) {
	return _VerifiyingPaymaster.Contract.GetBalance(&_VerifiyingPaymaster.CallOpts, paymasterId)
}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_VerifiyingPaymaster *VerifiyingPaymasterCaller) GetDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VerifiyingPaymaster.contract.Call(opts, &out, "getDeposit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_VerifiyingPaymaster *VerifiyingPaymasterSession) GetDeposit() (*big.Int, error) {
	return _VerifiyingPaymaster.Contract.GetDeposit(&_VerifiyingPaymaster.CallOpts)
}

// GetDeposit is a free data retrieval call binding the contract method 0xc399ec88.
//
// Solidity: function getDeposit() view returns(uint256)
func (_VerifiyingPaymaster *VerifiyingPaymasterCallerSession) GetDeposit() (*big.Int, error) {
	return _VerifiyingPaymaster.Contract.GetDeposit(&_VerifiyingPaymaster.CallOpts)
}

// GetHash is a free data retrieval call binding the contract method 0x1f67de80.
//
// Solidity: function getHash((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp, address paymasterId) view returns(bytes32)
func (_VerifiyingPaymaster *VerifiyingPaymasterCaller) GetHash(opts *bind.CallOpts, userOp *userop.UserOperation, paymasterId common.Address) ([32]byte, error) {
	var out []interface{}
	err := _VerifiyingPaymaster.contract.Call(opts, &out, "getHash", userOp, paymasterId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetHash is a free data retrieval call binding the contract method 0x1f67de80.
//
// Solidity: function getHash((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp, address paymasterId) view returns(bytes32)
func (_VerifiyingPaymaster *VerifiyingPaymasterSession) GetHash(userOp *userop.UserOperation, paymasterId common.Address) ([32]byte, error) {
	return _VerifiyingPaymaster.Contract.GetHash(&_VerifiyingPaymaster.CallOpts, userOp, paymasterId)
}

// GetHash is a free data retrieval call binding the contract method 0x1f67de80.
//
// Solidity: function getHash((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp, address paymasterId) view returns(bytes32)
func (_VerifiyingPaymaster *VerifiyingPaymasterCallerSession) GetHash(userOp *userop.UserOperation, paymasterId common.Address) ([32]byte, error) {
	return _VerifiyingPaymaster.Contract.GetHash(&_VerifiyingPaymaster.CallOpts, userOp, paymasterId)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VerifiyingPaymaster *VerifiyingPaymasterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VerifiyingPaymaster.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VerifiyingPaymaster *VerifiyingPaymasterSession) Owner() (common.Address, error) {
	return _VerifiyingPaymaster.Contract.Owner(&_VerifiyingPaymaster.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VerifiyingPaymaster *VerifiyingPaymasterCallerSession) Owner() (common.Address, error) {
	return _VerifiyingPaymaster.Contract.Owner(&_VerifiyingPaymaster.CallOpts)
}

// PaymasterIdBalances is a free data retrieval call binding the contract method 0xa40a7ddc.
//
// Solidity: function paymasterIdBalances(address ) view returns(uint256)
func (_VerifiyingPaymaster *VerifiyingPaymasterCaller) PaymasterIdBalances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VerifiyingPaymaster.contract.Call(opts, &out, "paymasterIdBalances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PaymasterIdBalances is a free data retrieval call binding the contract method 0xa40a7ddc.
//
// Solidity: function paymasterIdBalances(address ) view returns(uint256)
func (_VerifiyingPaymaster *VerifiyingPaymasterSession) PaymasterIdBalances(arg0 common.Address) (*big.Int, error) {
	return _VerifiyingPaymaster.Contract.PaymasterIdBalances(&_VerifiyingPaymaster.CallOpts, arg0)
}

// PaymasterIdBalances is a free data retrieval call binding the contract method 0xa40a7ddc.
//
// Solidity: function paymasterIdBalances(address ) view returns(uint256)
func (_VerifiyingPaymaster *VerifiyingPaymasterCallerSession) PaymasterIdBalances(arg0 common.Address) (*big.Int, error) {
	return _VerifiyingPaymaster.Contract.PaymasterIdBalances(&_VerifiyingPaymaster.CallOpts, arg0)
}

// VerifyingSigner is a free data retrieval call binding the contract method 0x23d9ac9b.
//
// Solidity: function verifyingSigner() view returns(address)
func (_VerifiyingPaymaster *VerifiyingPaymasterCaller) VerifyingSigner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VerifiyingPaymaster.contract.Call(opts, &out, "verifyingSigner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VerifyingSigner is a free data retrieval call binding the contract method 0x23d9ac9b.
//
// Solidity: function verifyingSigner() view returns(address)
func (_VerifiyingPaymaster *VerifiyingPaymasterSession) VerifyingSigner() (common.Address, error) {
	return _VerifiyingPaymaster.Contract.VerifyingSigner(&_VerifiyingPaymaster.CallOpts)
}

// VerifyingSigner is a free data retrieval call binding the contract method 0x23d9ac9b.
//
// Solidity: function verifyingSigner() view returns(address)
func (_VerifiyingPaymaster *VerifiyingPaymasterCallerSession) VerifyingSigner() (common.Address, error) {
	return _VerifiyingPaymaster.Contract.VerifyingSigner(&_VerifiyingPaymaster.CallOpts)
}

// AddStake is a paid mutator transaction binding the contract method 0x0396cb60.
//
// Solidity: function addStake(uint32 unstakeDelaySec) payable returns()
func (_VerifiyingPaymaster *VerifiyingPaymasterTransactor) AddStake(opts *bind.TransactOpts, unstakeDelaySec uint32) (*types.Transaction, error) {
	return _VerifiyingPaymaster.contract.Transact(opts, "addStake", unstakeDelaySec)
}

// AddStake is a paid mutator transaction binding the contract method 0x0396cb60.
//
// Solidity: function addStake(uint32 unstakeDelaySec) payable returns()
func (_VerifiyingPaymaster *VerifiyingPaymasterSession) AddStake(unstakeDelaySec uint32) (*types.Transaction, error) {
	return _VerifiyingPaymaster.Contract.AddStake(&_VerifiyingPaymaster.TransactOpts, unstakeDelaySec)
}

// AddStake is a paid mutator transaction binding the contract method 0x0396cb60.
//
// Solidity: function addStake(uint32 unstakeDelaySec) payable returns()
func (_VerifiyingPaymaster *VerifiyingPaymasterTransactorSession) AddStake(unstakeDelaySec uint32) (*types.Transaction, error) {
	return _VerifiyingPaymaster.Contract.AddStake(&_VerifiyingPaymaster.TransactOpts, unstakeDelaySec)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_VerifiyingPaymaster *VerifiyingPaymasterTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifiyingPaymaster.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_VerifiyingPaymaster *VerifiyingPaymasterSession) Deposit() (*types.Transaction, error) {
	return _VerifiyingPaymaster.Contract.Deposit(&_VerifiyingPaymaster.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_VerifiyingPaymaster *VerifiyingPaymasterTransactorSession) Deposit() (*types.Transaction, error) {
	return _VerifiyingPaymaster.Contract.Deposit(&_VerifiyingPaymaster.TransactOpts)
}

// DepositFor is a paid mutator transaction binding the contract method 0xaa67c919.
//
// Solidity: function depositFor(address paymasterId) payable returns()
func (_VerifiyingPaymaster *VerifiyingPaymasterTransactor) DepositFor(opts *bind.TransactOpts, paymasterId common.Address) (*types.Transaction, error) {
	return _VerifiyingPaymaster.contract.Transact(opts, "depositFor", paymasterId)
}

// DepositFor is a paid mutator transaction binding the contract method 0xaa67c919.
//
// Solidity: function depositFor(address paymasterId) payable returns()
func (_VerifiyingPaymaster *VerifiyingPaymasterSession) DepositFor(paymasterId common.Address) (*types.Transaction, error) {
	return _VerifiyingPaymaster.Contract.DepositFor(&_VerifiyingPaymaster.TransactOpts, paymasterId)
}

// DepositFor is a paid mutator transaction binding the contract method 0xaa67c919.
//
// Solidity: function depositFor(address paymasterId) payable returns()
func (_VerifiyingPaymaster *VerifiyingPaymasterTransactorSession) DepositFor(paymasterId common.Address) (*types.Transaction, error) {
	return _VerifiyingPaymaster.Contract.DepositFor(&_VerifiyingPaymaster.TransactOpts, paymasterId)
}

// PostOp is a paid mutator transaction binding the contract method 0xa9a23409.
//
// Solidity: function postOp(uint8 mode, bytes context, uint256 actualGasCost) returns()
func (_VerifiyingPaymaster *VerifiyingPaymasterTransactor) PostOp(opts *bind.TransactOpts, mode uint8, context []byte, actualGasCost *big.Int) (*types.Transaction, error) {
	return _VerifiyingPaymaster.contract.Transact(opts, "postOp", mode, context, actualGasCost)
}

// PostOp is a paid mutator transaction binding the contract method 0xa9a23409.
//
// Solidity: function postOp(uint8 mode, bytes context, uint256 actualGasCost) returns()
func (_VerifiyingPaymaster *VerifiyingPaymasterSession) PostOp(mode uint8, context []byte, actualGasCost *big.Int) (*types.Transaction, error) {
	return _VerifiyingPaymaster.Contract.PostOp(&_VerifiyingPaymaster.TransactOpts, mode, context, actualGasCost)
}

// PostOp is a paid mutator transaction binding the contract method 0xa9a23409.
//
// Solidity: function postOp(uint8 mode, bytes context, uint256 actualGasCost) returns()
func (_VerifiyingPaymaster *VerifiyingPaymasterTransactorSession) PostOp(mode uint8, context []byte, actualGasCost *big.Int) (*types.Transaction, error) {
	return _VerifiyingPaymaster.Contract.PostOp(&_VerifiyingPaymaster.TransactOpts, mode, context, actualGasCost)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VerifiyingPaymaster *VerifiyingPaymasterTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifiyingPaymaster.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VerifiyingPaymaster *VerifiyingPaymasterSession) RenounceOwnership() (*types.Transaction, error) {
	return _VerifiyingPaymaster.Contract.RenounceOwnership(&_VerifiyingPaymaster.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VerifiyingPaymaster *VerifiyingPaymasterTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _VerifiyingPaymaster.Contract.RenounceOwnership(&_VerifiyingPaymaster.TransactOpts)
}

// SetSigner is a paid mutator transaction binding the contract method 0x6c19e783.
//
// Solidity: function setSigner(address _newVerifyingSigner) payable returns()
func (_VerifiyingPaymaster *VerifiyingPaymasterTransactor) SetSigner(opts *bind.TransactOpts, _newVerifyingSigner common.Address) (*types.Transaction, error) {
	return _VerifiyingPaymaster.contract.Transact(opts, "setSigner", _newVerifyingSigner)
}

// SetSigner is a paid mutator transaction binding the contract method 0x6c19e783.
//
// Solidity: function setSigner(address _newVerifyingSigner) payable returns()
func (_VerifiyingPaymaster *VerifiyingPaymasterSession) SetSigner(_newVerifyingSigner common.Address) (*types.Transaction, error) {
	return _VerifiyingPaymaster.Contract.SetSigner(&_VerifiyingPaymaster.TransactOpts, _newVerifyingSigner)
}

// SetSigner is a paid mutator transaction binding the contract method 0x6c19e783.
//
// Solidity: function setSigner(address _newVerifyingSigner) payable returns()
func (_VerifiyingPaymaster *VerifiyingPaymasterTransactorSession) SetSigner(_newVerifyingSigner common.Address) (*types.Transaction, error) {
	return _VerifiyingPaymaster.Contract.SetSigner(&_VerifiyingPaymaster.TransactOpts, _newVerifyingSigner)
}

// SetUnaccountedEPGasOverhead is a paid mutator transaction binding the contract method 0xdeeb3874.
//
// Solidity: function setUnaccountedEPGasOverhead(uint256 value) returns()
func (_VerifiyingPaymaster *VerifiyingPaymasterTransactor) SetUnaccountedEPGasOverhead(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _VerifiyingPaymaster.contract.Transact(opts, "setUnaccountedEPGasOverhead", value)
}

// SetUnaccountedEPGasOverhead is a paid mutator transaction binding the contract method 0xdeeb3874.
//
// Solidity: function setUnaccountedEPGasOverhead(uint256 value) returns()
func (_VerifiyingPaymaster *VerifiyingPaymasterSession) SetUnaccountedEPGasOverhead(value *big.Int) (*types.Transaction, error) {
	return _VerifiyingPaymaster.Contract.SetUnaccountedEPGasOverhead(&_VerifiyingPaymaster.TransactOpts, value)
}

// SetUnaccountedEPGasOverhead is a paid mutator transaction binding the contract method 0xdeeb3874.
//
// Solidity: function setUnaccountedEPGasOverhead(uint256 value) returns()
func (_VerifiyingPaymaster *VerifiyingPaymasterTransactorSession) SetUnaccountedEPGasOverhead(value *big.Int) (*types.Transaction, error) {
	return _VerifiyingPaymaster.Contract.SetUnaccountedEPGasOverhead(&_VerifiyingPaymaster.TransactOpts, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VerifiyingPaymaster *VerifiyingPaymasterTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _VerifiyingPaymaster.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VerifiyingPaymaster *VerifiyingPaymasterSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VerifiyingPaymaster.Contract.TransferOwnership(&_VerifiyingPaymaster.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VerifiyingPaymaster *VerifiyingPaymasterTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VerifiyingPaymaster.Contract.TransferOwnership(&_VerifiyingPaymaster.TransactOpts, newOwner)
}

// UnlockStake is a paid mutator transaction binding the contract method 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns()
func (_VerifiyingPaymaster *VerifiyingPaymasterTransactor) UnlockStake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifiyingPaymaster.contract.Transact(opts, "unlockStake")
}

// UnlockStake is a paid mutator transaction binding the contract method 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns()
func (_VerifiyingPaymaster *VerifiyingPaymasterSession) UnlockStake() (*types.Transaction, error) {
	return _VerifiyingPaymaster.Contract.UnlockStake(&_VerifiyingPaymaster.TransactOpts)
}

// UnlockStake is a paid mutator transaction binding the contract method 0xbb9fe6bf.
//
// Solidity: function unlockStake() returns()
func (_VerifiyingPaymaster *VerifiyingPaymasterTransactorSession) UnlockStake() (*types.Transaction, error) {
	return _VerifiyingPaymaster.Contract.UnlockStake(&_VerifiyingPaymaster.TransactOpts)
}

// ValidatePaymasterUserOp is a paid mutator transaction binding the contract method 0xf465c77e.
//
// Solidity: function validatePaymasterUserOp((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp, bytes32 userOpHash, uint256 maxCost) returns(bytes context, uint256 validationData)
func (_VerifiyingPaymaster *VerifiyingPaymasterTransactor) ValidatePaymasterUserOp(opts *bind.TransactOpts, userOp *userop.UserOperation, userOpHash [32]byte, maxCost *big.Int) (*types.Transaction, error) {
	return _VerifiyingPaymaster.contract.Transact(opts, "validatePaymasterUserOp", userOp, userOpHash, maxCost)
}

// ValidatePaymasterUserOp is a paid mutator transaction binding the contract method 0xf465c77e.
//
// Solidity: function validatePaymasterUserOp((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp, bytes32 userOpHash, uint256 maxCost) returns(bytes context, uint256 validationData)
func (_VerifiyingPaymaster *VerifiyingPaymasterSession) ValidatePaymasterUserOp(userOp *userop.UserOperation, userOpHash [32]byte, maxCost *big.Int) (*types.Transaction, error) {
	return _VerifiyingPaymaster.Contract.ValidatePaymasterUserOp(&_VerifiyingPaymaster.TransactOpts, userOp, userOpHash, maxCost)
}

// ValidatePaymasterUserOp is a paid mutator transaction binding the contract method 0xf465c77e.
//
// Solidity: function validatePaymasterUserOp((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) userOp, bytes32 userOpHash, uint256 maxCost) returns(bytes context, uint256 validationData)
func (_VerifiyingPaymaster *VerifiyingPaymasterTransactorSession) ValidatePaymasterUserOp(userOp *userop.UserOperation, userOpHash [32]byte, maxCost *big.Int) (*types.Transaction, error) {
	return _VerifiyingPaymaster.Contract.ValidatePaymasterUserOp(&_VerifiyingPaymaster.TransactOpts, userOp, userOpHash, maxCost)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address withdrawAddress) returns()
func (_VerifiyingPaymaster *VerifiyingPaymasterTransactor) WithdrawStake(opts *bind.TransactOpts, withdrawAddress common.Address) (*types.Transaction, error) {
	return _VerifiyingPaymaster.contract.Transact(opts, "withdrawStake", withdrawAddress)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address withdrawAddress) returns()
func (_VerifiyingPaymaster *VerifiyingPaymasterSession) WithdrawStake(withdrawAddress common.Address) (*types.Transaction, error) {
	return _VerifiyingPaymaster.Contract.WithdrawStake(&_VerifiyingPaymaster.TransactOpts, withdrawAddress)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address withdrawAddress) returns()
func (_VerifiyingPaymaster *VerifiyingPaymasterTransactorSession) WithdrawStake(withdrawAddress common.Address) (*types.Transaction, error) {
	return _VerifiyingPaymaster.Contract.WithdrawStake(&_VerifiyingPaymaster.TransactOpts, withdrawAddress)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address withdrawAddress, uint256 amount) returns()
func (_VerifiyingPaymaster *VerifiyingPaymasterTransactor) WithdrawTo(opts *bind.TransactOpts, withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VerifiyingPaymaster.contract.Transact(opts, "withdrawTo", withdrawAddress, amount)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address withdrawAddress, uint256 amount) returns()
func (_VerifiyingPaymaster *VerifiyingPaymasterSession) WithdrawTo(withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VerifiyingPaymaster.Contract.WithdrawTo(&_VerifiyingPaymaster.TransactOpts, withdrawAddress, amount)
}

// WithdrawTo is a paid mutator transaction binding the contract method 0x205c2878.
//
// Solidity: function withdrawTo(address withdrawAddress, uint256 amount) returns()
func (_VerifiyingPaymaster *VerifiyingPaymasterTransactorSession) WithdrawTo(withdrawAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VerifiyingPaymaster.Contract.WithdrawTo(&_VerifiyingPaymaster.TransactOpts, withdrawAddress, amount)
}

// VerifiyingPaymasterEPGasOverheadChangedIterator is returned from FilterEPGasOverheadChanged and is used to iterate over the raw logs and unpacked data for EPGasOverheadChanged events raised by the VerifiyingPaymaster contract.
type VerifiyingPaymasterEPGasOverheadChangedIterator struct {
	Event *VerifiyingPaymasterEPGasOverheadChanged // Event containing the contract specifics and raw log

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
func (it *VerifiyingPaymasterEPGasOverheadChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifiyingPaymasterEPGasOverheadChanged)
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
		it.Event = new(VerifiyingPaymasterEPGasOverheadChanged)
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
func (it *VerifiyingPaymasterEPGasOverheadChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VerifiyingPaymasterEPGasOverheadChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VerifiyingPaymasterEPGasOverheadChanged represents a EPGasOverheadChanged event raised by the VerifiyingPaymaster contract.
type VerifiyingPaymasterEPGasOverheadChanged struct {
	OldValue *big.Int
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterEPGasOverheadChanged is a free log retrieval operation binding the contract event 0x0b2f957fc0a9306310238ef9a212935360e98fe3f8bc525f4cb69d38b1efa859.
//
// Solidity: event EPGasOverheadChanged(uint256 indexed _oldValue, uint256 indexed _newValue)
func (_VerifiyingPaymaster *VerifiyingPaymasterFilterer) FilterEPGasOverheadChanged(opts *bind.FilterOpts, _oldValue []*big.Int, _newValue []*big.Int) (*VerifiyingPaymasterEPGasOverheadChangedIterator, error) {

	var _oldValueRule []interface{}
	for _, _oldValueItem := range _oldValue {
		_oldValueRule = append(_oldValueRule, _oldValueItem)
	}
	var _newValueRule []interface{}
	for _, _newValueItem := range _newValue {
		_newValueRule = append(_newValueRule, _newValueItem)
	}

	logs, sub, err := _VerifiyingPaymaster.contract.FilterLogs(opts, "EPGasOverheadChanged", _oldValueRule, _newValueRule)
	if err != nil {
		return nil, err
	}
	return &VerifiyingPaymasterEPGasOverheadChangedIterator{contract: _VerifiyingPaymaster.contract, event: "EPGasOverheadChanged", logs: logs, sub: sub}, nil
}

// WatchEPGasOverheadChanged is a free log subscription operation binding the contract event 0x0b2f957fc0a9306310238ef9a212935360e98fe3f8bc525f4cb69d38b1efa859.
//
// Solidity: event EPGasOverheadChanged(uint256 indexed _oldValue, uint256 indexed _newValue)
func (_VerifiyingPaymaster *VerifiyingPaymasterFilterer) WatchEPGasOverheadChanged(opts *bind.WatchOpts, sink chan<- *VerifiyingPaymasterEPGasOverheadChanged, _oldValue []*big.Int, _newValue []*big.Int) (event.Subscription, error) {

	var _oldValueRule []interface{}
	for _, _oldValueItem := range _oldValue {
		_oldValueRule = append(_oldValueRule, _oldValueItem)
	}
	var _newValueRule []interface{}
	for _, _newValueItem := range _newValue {
		_newValueRule = append(_newValueRule, _newValueItem)
	}

	logs, sub, err := _VerifiyingPaymaster.contract.WatchLogs(opts, "EPGasOverheadChanged", _oldValueRule, _newValueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VerifiyingPaymasterEPGasOverheadChanged)
				if err := _VerifiyingPaymaster.contract.UnpackLog(event, "EPGasOverheadChanged", log); err != nil {
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

// ParseEPGasOverheadChanged is a log parse operation binding the contract event 0x0b2f957fc0a9306310238ef9a212935360e98fe3f8bc525f4cb69d38b1efa859.
//
// Solidity: event EPGasOverheadChanged(uint256 indexed _oldValue, uint256 indexed _newValue)
func (_VerifiyingPaymaster *VerifiyingPaymasterFilterer) ParseEPGasOverheadChanged(log types.Log) (*VerifiyingPaymasterEPGasOverheadChanged, error) {
	event := new(VerifiyingPaymasterEPGasOverheadChanged)
	if err := _VerifiyingPaymaster.contract.UnpackLog(event, "EPGasOverheadChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VerifiyingPaymasterGasBalanceDeductedIterator is returned from FilterGasBalanceDeducted and is used to iterate over the raw logs and unpacked data for GasBalanceDeducted events raised by the VerifiyingPaymaster contract.
type VerifiyingPaymasterGasBalanceDeductedIterator struct {
	Event *VerifiyingPaymasterGasBalanceDeducted // Event containing the contract specifics and raw log

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
func (it *VerifiyingPaymasterGasBalanceDeductedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifiyingPaymasterGasBalanceDeducted)
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
		it.Event = new(VerifiyingPaymasterGasBalanceDeducted)
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
func (it *VerifiyingPaymasterGasBalanceDeductedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VerifiyingPaymasterGasBalanceDeductedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VerifiyingPaymasterGasBalanceDeducted represents a GasBalanceDeducted event raised by the VerifiyingPaymaster contract.
type VerifiyingPaymasterGasBalanceDeducted struct {
	PaymasterId common.Address
	Charge      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterGasBalanceDeducted is a free log retrieval operation binding the contract event 0x5dc1c754041954fe976773fa441397a7928c7127a1c83904214a7d2563399007.
//
// Solidity: event GasBalanceDeducted(address indexed _paymasterId, uint256 indexed _charge)
func (_VerifiyingPaymaster *VerifiyingPaymasterFilterer) FilterGasBalanceDeducted(opts *bind.FilterOpts, _paymasterId []common.Address, _charge []*big.Int) (*VerifiyingPaymasterGasBalanceDeductedIterator, error) {

	var _paymasterIdRule []interface{}
	for _, _paymasterIdItem := range _paymasterId {
		_paymasterIdRule = append(_paymasterIdRule, _paymasterIdItem)
	}
	var _chargeRule []interface{}
	for _, _chargeItem := range _charge {
		_chargeRule = append(_chargeRule, _chargeItem)
	}

	logs, sub, err := _VerifiyingPaymaster.contract.FilterLogs(opts, "GasBalanceDeducted", _paymasterIdRule, _chargeRule)
	if err != nil {
		return nil, err
	}
	return &VerifiyingPaymasterGasBalanceDeductedIterator{contract: _VerifiyingPaymaster.contract, event: "GasBalanceDeducted", logs: logs, sub: sub}, nil
}

// WatchGasBalanceDeducted is a free log subscription operation binding the contract event 0x5dc1c754041954fe976773fa441397a7928c7127a1c83904214a7d2563399007.
//
// Solidity: event GasBalanceDeducted(address indexed _paymasterId, uint256 indexed _charge)
func (_VerifiyingPaymaster *VerifiyingPaymasterFilterer) WatchGasBalanceDeducted(opts *bind.WatchOpts, sink chan<- *VerifiyingPaymasterGasBalanceDeducted, _paymasterId []common.Address, _charge []*big.Int) (event.Subscription, error) {

	var _paymasterIdRule []interface{}
	for _, _paymasterIdItem := range _paymasterId {
		_paymasterIdRule = append(_paymasterIdRule, _paymasterIdItem)
	}
	var _chargeRule []interface{}
	for _, _chargeItem := range _charge {
		_chargeRule = append(_chargeRule, _chargeItem)
	}

	logs, sub, err := _VerifiyingPaymaster.contract.WatchLogs(opts, "GasBalanceDeducted", _paymasterIdRule, _chargeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VerifiyingPaymasterGasBalanceDeducted)
				if err := _VerifiyingPaymaster.contract.UnpackLog(event, "GasBalanceDeducted", log); err != nil {
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

// ParseGasBalanceDeducted is a log parse operation binding the contract event 0x5dc1c754041954fe976773fa441397a7928c7127a1c83904214a7d2563399007.
//
// Solidity: event GasBalanceDeducted(address indexed _paymasterId, uint256 indexed _charge)
func (_VerifiyingPaymaster *VerifiyingPaymasterFilterer) ParseGasBalanceDeducted(log types.Log) (*VerifiyingPaymasterGasBalanceDeducted, error) {
	event := new(VerifiyingPaymasterGasBalanceDeducted)
	if err := _VerifiyingPaymaster.contract.UnpackLog(event, "GasBalanceDeducted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VerifiyingPaymasterGasDepositedIterator is returned from FilterGasDeposited and is used to iterate over the raw logs and unpacked data for GasDeposited events raised by the VerifiyingPaymaster contract.
type VerifiyingPaymasterGasDepositedIterator struct {
	Event *VerifiyingPaymasterGasDeposited // Event containing the contract specifics and raw log

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
func (it *VerifiyingPaymasterGasDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifiyingPaymasterGasDeposited)
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
		it.Event = new(VerifiyingPaymasterGasDeposited)
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
func (it *VerifiyingPaymasterGasDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VerifiyingPaymasterGasDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VerifiyingPaymasterGasDeposited represents a GasDeposited event raised by the VerifiyingPaymaster contract.
type VerifiyingPaymasterGasDeposited struct {
	PaymasterId common.Address
	Value       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterGasDeposited is a free log retrieval operation binding the contract event 0x1dbbf474736d6415d6a265fabee708fe6e988f6fd0c9d870ded36cab380898dd.
//
// Solidity: event GasDeposited(address indexed _paymasterId, uint256 indexed _value)
func (_VerifiyingPaymaster *VerifiyingPaymasterFilterer) FilterGasDeposited(opts *bind.FilterOpts, _paymasterId []common.Address, _value []*big.Int) (*VerifiyingPaymasterGasDepositedIterator, error) {

	var _paymasterIdRule []interface{}
	for _, _paymasterIdItem := range _paymasterId {
		_paymasterIdRule = append(_paymasterIdRule, _paymasterIdItem)
	}
	var _valueRule []interface{}
	for _, _valueItem := range _value {
		_valueRule = append(_valueRule, _valueItem)
	}

	logs, sub, err := _VerifiyingPaymaster.contract.FilterLogs(opts, "GasDeposited", _paymasterIdRule, _valueRule)
	if err != nil {
		return nil, err
	}
	return &VerifiyingPaymasterGasDepositedIterator{contract: _VerifiyingPaymaster.contract, event: "GasDeposited", logs: logs, sub: sub}, nil
}

// WatchGasDeposited is a free log subscription operation binding the contract event 0x1dbbf474736d6415d6a265fabee708fe6e988f6fd0c9d870ded36cab380898dd.
//
// Solidity: event GasDeposited(address indexed _paymasterId, uint256 indexed _value)
func (_VerifiyingPaymaster *VerifiyingPaymasterFilterer) WatchGasDeposited(opts *bind.WatchOpts, sink chan<- *VerifiyingPaymasterGasDeposited, _paymasterId []common.Address, _value []*big.Int) (event.Subscription, error) {

	var _paymasterIdRule []interface{}
	for _, _paymasterIdItem := range _paymasterId {
		_paymasterIdRule = append(_paymasterIdRule, _paymasterIdItem)
	}
	var _valueRule []interface{}
	for _, _valueItem := range _value {
		_valueRule = append(_valueRule, _valueItem)
	}

	logs, sub, err := _VerifiyingPaymaster.contract.WatchLogs(opts, "GasDeposited", _paymasterIdRule, _valueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VerifiyingPaymasterGasDeposited)
				if err := _VerifiyingPaymaster.contract.UnpackLog(event, "GasDeposited", log); err != nil {
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

// ParseGasDeposited is a log parse operation binding the contract event 0x1dbbf474736d6415d6a265fabee708fe6e988f6fd0c9d870ded36cab380898dd.
//
// Solidity: event GasDeposited(address indexed _paymasterId, uint256 indexed _value)
func (_VerifiyingPaymaster *VerifiyingPaymasterFilterer) ParseGasDeposited(log types.Log) (*VerifiyingPaymasterGasDeposited, error) {
	event := new(VerifiyingPaymasterGasDeposited)
	if err := _VerifiyingPaymaster.contract.UnpackLog(event, "GasDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VerifiyingPaymasterGasWithdrawnIterator is returned from FilterGasWithdrawn and is used to iterate over the raw logs and unpacked data for GasWithdrawn events raised by the VerifiyingPaymaster contract.
type VerifiyingPaymasterGasWithdrawnIterator struct {
	Event *VerifiyingPaymasterGasWithdrawn // Event containing the contract specifics and raw log

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
func (it *VerifiyingPaymasterGasWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifiyingPaymasterGasWithdrawn)
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
		it.Event = new(VerifiyingPaymasterGasWithdrawn)
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
func (it *VerifiyingPaymasterGasWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VerifiyingPaymasterGasWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VerifiyingPaymasterGasWithdrawn represents a GasWithdrawn event raised by the VerifiyingPaymaster contract.
type VerifiyingPaymasterGasWithdrawn struct {
	PaymasterId common.Address
	To          common.Address
	Value       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterGasWithdrawn is a free log retrieval operation binding the contract event 0x926a144b6fffc1d73f115b81af7ec66a7c12aed0ff73197c39a683753fc1d925.
//
// Solidity: event GasWithdrawn(address indexed _paymasterId, address indexed _to, uint256 indexed _value)
func (_VerifiyingPaymaster *VerifiyingPaymasterFilterer) FilterGasWithdrawn(opts *bind.FilterOpts, _paymasterId []common.Address, _to []common.Address, _value []*big.Int) (*VerifiyingPaymasterGasWithdrawnIterator, error) {

	var _paymasterIdRule []interface{}
	for _, _paymasterIdItem := range _paymasterId {
		_paymasterIdRule = append(_paymasterIdRule, _paymasterIdItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _valueRule []interface{}
	for _, _valueItem := range _value {
		_valueRule = append(_valueRule, _valueItem)
	}

	logs, sub, err := _VerifiyingPaymaster.contract.FilterLogs(opts, "GasWithdrawn", _paymasterIdRule, _toRule, _valueRule)
	if err != nil {
		return nil, err
	}
	return &VerifiyingPaymasterGasWithdrawnIterator{contract: _VerifiyingPaymaster.contract, event: "GasWithdrawn", logs: logs, sub: sub}, nil
}

// WatchGasWithdrawn is a free log subscription operation binding the contract event 0x926a144b6fffc1d73f115b81af7ec66a7c12aed0ff73197c39a683753fc1d925.
//
// Solidity: event GasWithdrawn(address indexed _paymasterId, address indexed _to, uint256 indexed _value)
func (_VerifiyingPaymaster *VerifiyingPaymasterFilterer) WatchGasWithdrawn(opts *bind.WatchOpts, sink chan<- *VerifiyingPaymasterGasWithdrawn, _paymasterId []common.Address, _to []common.Address, _value []*big.Int) (event.Subscription, error) {

	var _paymasterIdRule []interface{}
	for _, _paymasterIdItem := range _paymasterId {
		_paymasterIdRule = append(_paymasterIdRule, _paymasterIdItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _valueRule []interface{}
	for _, _valueItem := range _value {
		_valueRule = append(_valueRule, _valueItem)
	}

	logs, sub, err := _VerifiyingPaymaster.contract.WatchLogs(opts, "GasWithdrawn", _paymasterIdRule, _toRule, _valueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VerifiyingPaymasterGasWithdrawn)
				if err := _VerifiyingPaymaster.contract.UnpackLog(event, "GasWithdrawn", log); err != nil {
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

// ParseGasWithdrawn is a log parse operation binding the contract event 0x926a144b6fffc1d73f115b81af7ec66a7c12aed0ff73197c39a683753fc1d925.
//
// Solidity: event GasWithdrawn(address indexed _paymasterId, address indexed _to, uint256 indexed _value)
func (_VerifiyingPaymaster *VerifiyingPaymasterFilterer) ParseGasWithdrawn(log types.Log) (*VerifiyingPaymasterGasWithdrawn, error) {
	event := new(VerifiyingPaymasterGasWithdrawn)
	if err := _VerifiyingPaymaster.contract.UnpackLog(event, "GasWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VerifiyingPaymasterOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the VerifiyingPaymaster contract.
type VerifiyingPaymasterOwnershipTransferredIterator struct {
	Event *VerifiyingPaymasterOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *VerifiyingPaymasterOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifiyingPaymasterOwnershipTransferred)
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
		it.Event = new(VerifiyingPaymasterOwnershipTransferred)
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
func (it *VerifiyingPaymasterOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VerifiyingPaymasterOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VerifiyingPaymasterOwnershipTransferred represents a OwnershipTransferred event raised by the VerifiyingPaymaster contract.
type VerifiyingPaymasterOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VerifiyingPaymaster *VerifiyingPaymasterFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*VerifiyingPaymasterOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VerifiyingPaymaster.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &VerifiyingPaymasterOwnershipTransferredIterator{contract: _VerifiyingPaymaster.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VerifiyingPaymaster *VerifiyingPaymasterFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VerifiyingPaymasterOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VerifiyingPaymaster.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VerifiyingPaymasterOwnershipTransferred)
				if err := _VerifiyingPaymaster.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_VerifiyingPaymaster *VerifiyingPaymasterFilterer) ParseOwnershipTransferred(log types.Log) (*VerifiyingPaymasterOwnershipTransferred, error) {
	event := new(VerifiyingPaymasterOwnershipTransferred)
	if err := _VerifiyingPaymaster.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VerifiyingPaymasterVerifyingSignerChangedIterator is returned from FilterVerifyingSignerChanged and is used to iterate over the raw logs and unpacked data for VerifyingSignerChanged events raised by the VerifiyingPaymaster contract.
type VerifiyingPaymasterVerifyingSignerChangedIterator struct {
	Event *VerifiyingPaymasterVerifyingSignerChanged // Event containing the contract specifics and raw log

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
func (it *VerifiyingPaymasterVerifyingSignerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VerifiyingPaymasterVerifyingSignerChanged)
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
		it.Event = new(VerifiyingPaymasterVerifyingSignerChanged)
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
func (it *VerifiyingPaymasterVerifyingSignerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VerifiyingPaymasterVerifyingSignerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VerifiyingPaymasterVerifyingSignerChanged represents a VerifyingSignerChanged event raised by the VerifiyingPaymaster contract.
type VerifiyingPaymasterVerifyingSignerChanged struct {
	OldSigner common.Address
	NewSigner common.Address
	Actor     common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterVerifyingSignerChanged is a free log retrieval operation binding the contract event 0xe1f62c0e6d7bb6d470828565415bf2e87dbfea50e52d2d753788b529bd0c6d62.
//
// Solidity: event VerifyingSignerChanged(address indexed _oldSigner, address indexed _newSigner, address indexed _actor)
func (_VerifiyingPaymaster *VerifiyingPaymasterFilterer) FilterVerifyingSignerChanged(opts *bind.FilterOpts, _oldSigner []common.Address, _newSigner []common.Address, _actor []common.Address) (*VerifiyingPaymasterVerifyingSignerChangedIterator, error) {

	var _oldSignerRule []interface{}
	for _, _oldSignerItem := range _oldSigner {
		_oldSignerRule = append(_oldSignerRule, _oldSignerItem)
	}
	var _newSignerRule []interface{}
	for _, _newSignerItem := range _newSigner {
		_newSignerRule = append(_newSignerRule, _newSignerItem)
	}
	var _actorRule []interface{}
	for _, _actorItem := range _actor {
		_actorRule = append(_actorRule, _actorItem)
	}

	logs, sub, err := _VerifiyingPaymaster.contract.FilterLogs(opts, "VerifyingSignerChanged", _oldSignerRule, _newSignerRule, _actorRule)
	if err != nil {
		return nil, err
	}
	return &VerifiyingPaymasterVerifyingSignerChangedIterator{contract: _VerifiyingPaymaster.contract, event: "VerifyingSignerChanged", logs: logs, sub: sub}, nil
}

// WatchVerifyingSignerChanged is a free log subscription operation binding the contract event 0xe1f62c0e6d7bb6d470828565415bf2e87dbfea50e52d2d753788b529bd0c6d62.
//
// Solidity: event VerifyingSignerChanged(address indexed _oldSigner, address indexed _newSigner, address indexed _actor)
func (_VerifiyingPaymaster *VerifiyingPaymasterFilterer) WatchVerifyingSignerChanged(opts *bind.WatchOpts, sink chan<- *VerifiyingPaymasterVerifyingSignerChanged, _oldSigner []common.Address, _newSigner []common.Address, _actor []common.Address) (event.Subscription, error) {

	var _oldSignerRule []interface{}
	for _, _oldSignerItem := range _oldSigner {
		_oldSignerRule = append(_oldSignerRule, _oldSignerItem)
	}
	var _newSignerRule []interface{}
	for _, _newSignerItem := range _newSigner {
		_newSignerRule = append(_newSignerRule, _newSignerItem)
	}
	var _actorRule []interface{}
	for _, _actorItem := range _actor {
		_actorRule = append(_actorRule, _actorItem)
	}

	logs, sub, err := _VerifiyingPaymaster.contract.WatchLogs(opts, "VerifyingSignerChanged", _oldSignerRule, _newSignerRule, _actorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VerifiyingPaymasterVerifyingSignerChanged)
				if err := _VerifiyingPaymaster.contract.UnpackLog(event, "VerifyingSignerChanged", log); err != nil {
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

// ParseVerifyingSignerChanged is a log parse operation binding the contract event 0xe1f62c0e6d7bb6d470828565415bf2e87dbfea50e52d2d753788b529bd0c6d62.
//
// Solidity: event VerifyingSignerChanged(address indexed _oldSigner, address indexed _newSigner, address indexed _actor)
func (_VerifiyingPaymaster *VerifiyingPaymasterFilterer) ParseVerifyingSignerChanged(log types.Log) (*VerifiyingPaymasterVerifyingSignerChanged, error) {
	event := new(VerifiyingPaymasterVerifyingSignerChanged)
	if err := _VerifiyingPaymaster.contract.UnpackLog(event, "VerifyingSignerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
