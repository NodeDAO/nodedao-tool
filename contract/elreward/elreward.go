// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package elreward

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

// ElrewardMetaData contains all meta data concerning the Elreward contract.
var ElrewardMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"changePool\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claimDaoRewards\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"dao\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"daoTreasury\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"addresspayable\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDaoRewards\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPoolRewards\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"implementation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_ownerAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_pool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_dao\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_daoTreasury\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_poolConfig\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pool\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"addresspayable\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"poolConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"reinvestment\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDao\",\"inputs\":[{\"name\":\"_dao\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDaoTreasury\",\"inputs\":[{\"name\":\"_daoTreasury\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPoolConfig\",\"inputs\":[{\"name\":\"_poolConfig\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"upgradeTo\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"pure\"},{\"type\":\"event\",\"name\":\"AdminChanged\",\"inputs\":[{\"name\":\"previousAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BeaconUpgraded\",\"inputs\":[{\"name\":\"beacon\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DaoChanged\",\"inputs\":[{\"name\":\"_oldDao\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"_dao\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DaoRewardsClaimed\",\"inputs\":[{\"name\":\"_to\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"_daoRewards\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DaoTreasuryChanged\",\"inputs\":[{\"name\":\"_oldDaoTreasury\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"_daoTreasury\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PoolChanged\",\"inputs\":[{\"name\":\"_oldPool\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"_pool\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PoolConfigChanged\",\"inputs\":[{\"name\":\"_oldOperatorRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"_operatorRegistry\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PoolSet\",\"inputs\":[{\"name\":\"_pool\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Received\",\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Settle\",\"inputs\":[{\"name\":\"_daoRewards\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_poolRewards\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"_to\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidAddr\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidVersion\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidtypeId\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PermissionDenied\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TransferFailed\",\"inputs\":[]}]",
}

// ElrewardABI is the input ABI used to generate the binding from.
// Deprecated: Use ElrewardMetaData.ABI instead.
var ElrewardABI = ElrewardMetaData.ABI

// Elreward is an auto generated Go binding around an Ethereum contract.
type Elreward struct {
	ElrewardCaller     // Read-only binding to the contract
	ElrewardTransactor // Write-only binding to the contract
	ElrewardFilterer   // Log filterer for contract events
}

// ElrewardCaller is an auto generated read-only Go binding around an Ethereum contract.
type ElrewardCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ElrewardTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ElrewardTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ElrewardFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ElrewardFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ElrewardSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ElrewardSession struct {
	Contract     *Elreward         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ElrewardCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ElrewardCallerSession struct {
	Contract *ElrewardCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ElrewardTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ElrewardTransactorSession struct {
	Contract     *ElrewardTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ElrewardRaw is an auto generated low-level Go binding around an Ethereum contract.
type ElrewardRaw struct {
	Contract *Elreward // Generic contract binding to access the raw methods on
}

// ElrewardCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ElrewardCallerRaw struct {
	Contract *ElrewardCaller // Generic read-only contract binding to access the raw methods on
}

// ElrewardTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ElrewardTransactorRaw struct {
	Contract *ElrewardTransactor // Generic write-only contract binding to access the raw methods on
}

// NewElreward creates a new instance of Elreward, bound to a specific deployed contract.
func NewElreward(address common.Address, backend bind.ContractBackend) (*Elreward, error) {
	contract, err := bindElreward(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Elreward{ElrewardCaller: ElrewardCaller{contract: contract}, ElrewardTransactor: ElrewardTransactor{contract: contract}, ElrewardFilterer: ElrewardFilterer{contract: contract}}, nil
}

// NewElrewardCaller creates a new read-only instance of Elreward, bound to a specific deployed contract.
func NewElrewardCaller(address common.Address, caller bind.ContractCaller) (*ElrewardCaller, error) {
	contract, err := bindElreward(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ElrewardCaller{contract: contract}, nil
}

// NewElrewardTransactor creates a new write-only instance of Elreward, bound to a specific deployed contract.
func NewElrewardTransactor(address common.Address, transactor bind.ContractTransactor) (*ElrewardTransactor, error) {
	contract, err := bindElreward(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ElrewardTransactor{contract: contract}, nil
}

// NewElrewardFilterer creates a new log filterer instance of Elreward, bound to a specific deployed contract.
func NewElrewardFilterer(address common.Address, filterer bind.ContractFilterer) (*ElrewardFilterer, error) {
	contract, err := bindElreward(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ElrewardFilterer{contract: contract}, nil
}

// bindElreward binds a generic wrapper to an already deployed contract.
func bindElreward(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ElrewardABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Elreward *ElrewardRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Elreward.Contract.ElrewardCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Elreward *ElrewardRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Elreward.Contract.ElrewardTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Elreward *ElrewardRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Elreward.Contract.ElrewardTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Elreward *ElrewardCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Elreward.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Elreward *ElrewardTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Elreward.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Elreward *ElrewardTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Elreward.Contract.contract.Transact(opts, method, params...)
}

// Dao is a free data retrieval call binding the contract method 0x4162169f.
//
// Solidity: function dao() view returns(address)
func (_Elreward *ElrewardCaller) Dao(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Elreward.contract.Call(opts, &out, "dao")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Dao is a free data retrieval call binding the contract method 0x4162169f.
//
// Solidity: function dao() view returns(address)
func (_Elreward *ElrewardSession) Dao() (common.Address, error) {
	return _Elreward.Contract.Dao(&_Elreward.CallOpts)
}

// Dao is a free data retrieval call binding the contract method 0x4162169f.
//
// Solidity: function dao() view returns(address)
func (_Elreward *ElrewardCallerSession) Dao() (common.Address, error) {
	return _Elreward.Contract.Dao(&_Elreward.CallOpts)
}

// DaoTreasury is a free data retrieval call binding the contract method 0x79022a9f.
//
// Solidity: function daoTreasury() view returns(address)
func (_Elreward *ElrewardCaller) DaoTreasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Elreward.contract.Call(opts, &out, "daoTreasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DaoTreasury is a free data retrieval call binding the contract method 0x79022a9f.
//
// Solidity: function daoTreasury() view returns(address)
func (_Elreward *ElrewardSession) DaoTreasury() (common.Address, error) {
	return _Elreward.Contract.DaoTreasury(&_Elreward.CallOpts)
}

// DaoTreasury is a free data retrieval call binding the contract method 0x79022a9f.
//
// Solidity: function daoTreasury() view returns(address)
func (_Elreward *ElrewardCallerSession) DaoTreasury() (common.Address, error) {
	return _Elreward.Contract.DaoTreasury(&_Elreward.CallOpts)
}

// GetDaoRewards is a free data retrieval call binding the contract method 0x5002b601.
//
// Solidity: function getDaoRewards() view returns(uint256)
func (_Elreward *ElrewardCaller) GetDaoRewards(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Elreward.contract.Call(opts, &out, "getDaoRewards")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDaoRewards is a free data retrieval call binding the contract method 0x5002b601.
//
// Solidity: function getDaoRewards() view returns(uint256)
func (_Elreward *ElrewardSession) GetDaoRewards() (*big.Int, error) {
	return _Elreward.Contract.GetDaoRewards(&_Elreward.CallOpts)
}

// GetDaoRewards is a free data retrieval call binding the contract method 0x5002b601.
//
// Solidity: function getDaoRewards() view returns(uint256)
func (_Elreward *ElrewardCallerSession) GetDaoRewards() (*big.Int, error) {
	return _Elreward.Contract.GetDaoRewards(&_Elreward.CallOpts)
}

// GetPoolRewards is a free data retrieval call binding the contract method 0x8f3d63da.
//
// Solidity: function getPoolRewards() view returns(uint256)
func (_Elreward *ElrewardCaller) GetPoolRewards(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Elreward.contract.Call(opts, &out, "getPoolRewards")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPoolRewards is a free data retrieval call binding the contract method 0x8f3d63da.
//
// Solidity: function getPoolRewards() view returns(uint256)
func (_Elreward *ElrewardSession) GetPoolRewards() (*big.Int, error) {
	return _Elreward.Contract.GetPoolRewards(&_Elreward.CallOpts)
}

// GetPoolRewards is a free data retrieval call binding the contract method 0x8f3d63da.
//
// Solidity: function getPoolRewards() view returns(uint256)
func (_Elreward *ElrewardCallerSession) GetPoolRewards() (*big.Int, error) {
	return _Elreward.Contract.GetPoolRewards(&_Elreward.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_Elreward *ElrewardCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Elreward.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_Elreward *ElrewardSession) Implementation() (common.Address, error) {
	return _Elreward.Contract.Implementation(&_Elreward.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_Elreward *ElrewardCallerSession) Implementation() (common.Address, error) {
	return _Elreward.Contract.Implementation(&_Elreward.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Elreward *ElrewardCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Elreward.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Elreward *ElrewardSession) Owner() (common.Address, error) {
	return _Elreward.Contract.Owner(&_Elreward.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Elreward *ElrewardCallerSession) Owner() (common.Address, error) {
	return _Elreward.Contract.Owner(&_Elreward.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Elreward *ElrewardCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Elreward.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Elreward *ElrewardSession) Paused() (bool, error) {
	return _Elreward.Contract.Paused(&_Elreward.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Elreward *ElrewardCallerSession) Paused() (bool, error) {
	return _Elreward.Contract.Paused(&_Elreward.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_Elreward *ElrewardCaller) Pool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Elreward.contract.Call(opts, &out, "pool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_Elreward *ElrewardSession) Pool() (common.Address, error) {
	return _Elreward.Contract.Pool(&_Elreward.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_Elreward *ElrewardCallerSession) Pool() (common.Address, error) {
	return _Elreward.Contract.Pool(&_Elreward.CallOpts)
}

// PoolConfig is a free data retrieval call binding the contract method 0x9695e195.
//
// Solidity: function poolConfig() view returns(address)
func (_Elreward *ElrewardCaller) PoolConfig(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Elreward.contract.Call(opts, &out, "poolConfig")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolConfig is a free data retrieval call binding the contract method 0x9695e195.
//
// Solidity: function poolConfig() view returns(address)
func (_Elreward *ElrewardSession) PoolConfig() (common.Address, error) {
	return _Elreward.Contract.PoolConfig(&_Elreward.CallOpts)
}

// PoolConfig is a free data retrieval call binding the contract method 0x9695e195.
//
// Solidity: function poolConfig() view returns(address)
func (_Elreward *ElrewardCallerSession) PoolConfig() (common.Address, error) {
	return _Elreward.Contract.PoolConfig(&_Elreward.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Elreward *ElrewardCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Elreward.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Elreward *ElrewardSession) ProxiableUUID() ([32]byte, error) {
	return _Elreward.Contract.ProxiableUUID(&_Elreward.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Elreward *ElrewardCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Elreward.Contract.ProxiableUUID(&_Elreward.CallOpts)
}

// TypeId is a free data retrieval call binding the contract method 0x8f940f63.
//
// Solidity: function typeId() pure returns(bytes32)
func (_Elreward *ElrewardCaller) TypeId(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Elreward.contract.Call(opts, &out, "typeId")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TypeId is a free data retrieval call binding the contract method 0x8f940f63.
//
// Solidity: function typeId() pure returns(bytes32)
func (_Elreward *ElrewardSession) TypeId() ([32]byte, error) {
	return _Elreward.Contract.TypeId(&_Elreward.CallOpts)
}

// TypeId is a free data retrieval call binding the contract method 0x8f940f63.
//
// Solidity: function typeId() pure returns(bytes32)
func (_Elreward *ElrewardCallerSession) TypeId() ([32]byte, error) {
	return _Elreward.Contract.TypeId(&_Elreward.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint8)
func (_Elreward *ElrewardCaller) Version(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Elreward.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint8)
func (_Elreward *ElrewardSession) Version() (uint8, error) {
	return _Elreward.Contract.Version(&_Elreward.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint8)
func (_Elreward *ElrewardCallerSession) Version() (uint8, error) {
	return _Elreward.Contract.Version(&_Elreward.CallOpts)
}

// ChangePool is a paid mutator transaction binding the contract method 0x4339bc30.
//
// Solidity: function changePool(address _pool) returns()
func (_Elreward *ElrewardTransactor) ChangePool(opts *bind.TransactOpts, _pool common.Address) (*types.Transaction, error) {
	return _Elreward.contract.Transact(opts, "changePool", _pool)
}

// ChangePool is a paid mutator transaction binding the contract method 0x4339bc30.
//
// Solidity: function changePool(address _pool) returns()
func (_Elreward *ElrewardSession) ChangePool(_pool common.Address) (*types.Transaction, error) {
	return _Elreward.Contract.ChangePool(&_Elreward.TransactOpts, _pool)
}

// ChangePool is a paid mutator transaction binding the contract method 0x4339bc30.
//
// Solidity: function changePool(address _pool) returns()
func (_Elreward *ElrewardTransactorSession) ChangePool(_pool common.Address) (*types.Transaction, error) {
	return _Elreward.Contract.ChangePool(&_Elreward.TransactOpts, _pool)
}

// ClaimDaoRewards is a paid mutator transaction binding the contract method 0xdc3d20a6.
//
// Solidity: function claimDaoRewards() returns()
func (_Elreward *ElrewardTransactor) ClaimDaoRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Elreward.contract.Transact(opts, "claimDaoRewards")
}

// ClaimDaoRewards is a paid mutator transaction binding the contract method 0xdc3d20a6.
//
// Solidity: function claimDaoRewards() returns()
func (_Elreward *ElrewardSession) ClaimDaoRewards() (*types.Transaction, error) {
	return _Elreward.Contract.ClaimDaoRewards(&_Elreward.TransactOpts)
}

// ClaimDaoRewards is a paid mutator transaction binding the contract method 0xdc3d20a6.
//
// Solidity: function claimDaoRewards() returns()
func (_Elreward *ElrewardTransactorSession) ClaimDaoRewards() (*types.Transaction, error) {
	return _Elreward.Contract.ClaimDaoRewards(&_Elreward.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _ownerAddr, address _pool, address _dao, address _daoTreasury, address _poolConfig) returns()
func (_Elreward *ElrewardTransactor) Initialize(opts *bind.TransactOpts, _ownerAddr common.Address, _pool common.Address, _dao common.Address, _daoTreasury common.Address, _poolConfig common.Address) (*types.Transaction, error) {
	return _Elreward.contract.Transact(opts, "initialize", _ownerAddr, _pool, _dao, _daoTreasury, _poolConfig)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _ownerAddr, address _pool, address _dao, address _daoTreasury, address _poolConfig) returns()
func (_Elreward *ElrewardSession) Initialize(_ownerAddr common.Address, _pool common.Address, _dao common.Address, _daoTreasury common.Address, _poolConfig common.Address) (*types.Transaction, error) {
	return _Elreward.Contract.Initialize(&_Elreward.TransactOpts, _ownerAddr, _pool, _dao, _daoTreasury, _poolConfig)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _ownerAddr, address _pool, address _dao, address _daoTreasury, address _poolConfig) returns()
func (_Elreward *ElrewardTransactorSession) Initialize(_ownerAddr common.Address, _pool common.Address, _dao common.Address, _daoTreasury common.Address, _poolConfig common.Address) (*types.Transaction, error) {
	return _Elreward.Contract.Initialize(&_Elreward.TransactOpts, _ownerAddr, _pool, _dao, _daoTreasury, _poolConfig)
}

// Reinvestment is a paid mutator transaction binding the contract method 0xb5e86bea.
//
// Solidity: function reinvestment() returns()
func (_Elreward *ElrewardTransactor) Reinvestment(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Elreward.contract.Transact(opts, "reinvestment")
}

// Reinvestment is a paid mutator transaction binding the contract method 0xb5e86bea.
//
// Solidity: function reinvestment() returns()
func (_Elreward *ElrewardSession) Reinvestment() (*types.Transaction, error) {
	return _Elreward.Contract.Reinvestment(&_Elreward.TransactOpts)
}

// Reinvestment is a paid mutator transaction binding the contract method 0xb5e86bea.
//
// Solidity: function reinvestment() returns()
func (_Elreward *ElrewardTransactorSession) Reinvestment() (*types.Transaction, error) {
	return _Elreward.Contract.Reinvestment(&_Elreward.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Elreward *ElrewardTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Elreward.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Elreward *ElrewardSession) RenounceOwnership() (*types.Transaction, error) {
	return _Elreward.Contract.RenounceOwnership(&_Elreward.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Elreward *ElrewardTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Elreward.Contract.RenounceOwnership(&_Elreward.TransactOpts)
}

// SetDao is a paid mutator transaction binding the contract method 0x6637b882.
//
// Solidity: function setDao(address _dao) returns()
func (_Elreward *ElrewardTransactor) SetDao(opts *bind.TransactOpts, _dao common.Address) (*types.Transaction, error) {
	return _Elreward.contract.Transact(opts, "setDao", _dao)
}

// SetDao is a paid mutator transaction binding the contract method 0x6637b882.
//
// Solidity: function setDao(address _dao) returns()
func (_Elreward *ElrewardSession) SetDao(_dao common.Address) (*types.Transaction, error) {
	return _Elreward.Contract.SetDao(&_Elreward.TransactOpts, _dao)
}

// SetDao is a paid mutator transaction binding the contract method 0x6637b882.
//
// Solidity: function setDao(address _dao) returns()
func (_Elreward *ElrewardTransactorSession) SetDao(_dao common.Address) (*types.Transaction, error) {
	return _Elreward.Contract.SetDao(&_Elreward.TransactOpts, _dao)
}

// SetDaoTreasury is a paid mutator transaction binding the contract method 0xa124df5d.
//
// Solidity: function setDaoTreasury(address _daoTreasury) returns()
func (_Elreward *ElrewardTransactor) SetDaoTreasury(opts *bind.TransactOpts, _daoTreasury common.Address) (*types.Transaction, error) {
	return _Elreward.contract.Transact(opts, "setDaoTreasury", _daoTreasury)
}

// SetDaoTreasury is a paid mutator transaction binding the contract method 0xa124df5d.
//
// Solidity: function setDaoTreasury(address _daoTreasury) returns()
func (_Elreward *ElrewardSession) SetDaoTreasury(_daoTreasury common.Address) (*types.Transaction, error) {
	return _Elreward.Contract.SetDaoTreasury(&_Elreward.TransactOpts, _daoTreasury)
}

// SetDaoTreasury is a paid mutator transaction binding the contract method 0xa124df5d.
//
// Solidity: function setDaoTreasury(address _daoTreasury) returns()
func (_Elreward *ElrewardTransactorSession) SetDaoTreasury(_daoTreasury common.Address) (*types.Transaction, error) {
	return _Elreward.Contract.SetDaoTreasury(&_Elreward.TransactOpts, _daoTreasury)
}

// SetPoolConfig is a paid mutator transaction binding the contract method 0x6bfad652.
//
// Solidity: function setPoolConfig(address _poolConfig) returns()
func (_Elreward *ElrewardTransactor) SetPoolConfig(opts *bind.TransactOpts, _poolConfig common.Address) (*types.Transaction, error) {
	return _Elreward.contract.Transact(opts, "setPoolConfig", _poolConfig)
}

// SetPoolConfig is a paid mutator transaction binding the contract method 0x6bfad652.
//
// Solidity: function setPoolConfig(address _poolConfig) returns()
func (_Elreward *ElrewardSession) SetPoolConfig(_poolConfig common.Address) (*types.Transaction, error) {
	return _Elreward.Contract.SetPoolConfig(&_Elreward.TransactOpts, _poolConfig)
}

// SetPoolConfig is a paid mutator transaction binding the contract method 0x6bfad652.
//
// Solidity: function setPoolConfig(address _poolConfig) returns()
func (_Elreward *ElrewardTransactorSession) SetPoolConfig(_poolConfig common.Address) (*types.Transaction, error) {
	return _Elreward.Contract.SetPoolConfig(&_Elreward.TransactOpts, _poolConfig)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Elreward *ElrewardTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Elreward.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Elreward *ElrewardSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Elreward.Contract.TransferOwnership(&_Elreward.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Elreward *ElrewardTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Elreward.Contract.TransferOwnership(&_Elreward.TransactOpts, newOwner)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Elreward *ElrewardTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Elreward.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Elreward *ElrewardSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Elreward.Contract.UpgradeTo(&_Elreward.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Elreward *ElrewardTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Elreward.Contract.UpgradeTo(&_Elreward.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Elreward *ElrewardTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Elreward.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Elreward *ElrewardSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Elreward.Contract.UpgradeToAndCall(&_Elreward.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Elreward *ElrewardTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Elreward.Contract.UpgradeToAndCall(&_Elreward.TransactOpts, newImplementation, data)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Elreward *ElrewardTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Elreward.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Elreward *ElrewardSession) Receive() (*types.Transaction, error) {
	return _Elreward.Contract.Receive(&_Elreward.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Elreward *ElrewardTransactorSession) Receive() (*types.Transaction, error) {
	return _Elreward.Contract.Receive(&_Elreward.TransactOpts)
}

// ElrewardAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Elreward contract.
type ElrewardAdminChangedIterator struct {
	Event *ElrewardAdminChanged // Event containing the contract specifics and raw log

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
func (it *ElrewardAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ElrewardAdminChanged)
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
		it.Event = new(ElrewardAdminChanged)
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
func (it *ElrewardAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ElrewardAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ElrewardAdminChanged represents a AdminChanged event raised by the Elreward contract.
type ElrewardAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Elreward *ElrewardFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*ElrewardAdminChangedIterator, error) {

	logs, sub, err := _Elreward.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &ElrewardAdminChangedIterator{contract: _Elreward.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Elreward *ElrewardFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *ElrewardAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Elreward.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ElrewardAdminChanged)
				if err := _Elreward.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_Elreward *ElrewardFilterer) ParseAdminChanged(log types.Log) (*ElrewardAdminChanged, error) {
	event := new(ElrewardAdminChanged)
	if err := _Elreward.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ElrewardBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the Elreward contract.
type ElrewardBeaconUpgradedIterator struct {
	Event *ElrewardBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *ElrewardBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ElrewardBeaconUpgraded)
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
		it.Event = new(ElrewardBeaconUpgraded)
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
func (it *ElrewardBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ElrewardBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ElrewardBeaconUpgraded represents a BeaconUpgraded event raised by the Elreward contract.
type ElrewardBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Elreward *ElrewardFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*ElrewardBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Elreward.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &ElrewardBeaconUpgradedIterator{contract: _Elreward.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Elreward *ElrewardFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *ElrewardBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Elreward.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ElrewardBeaconUpgraded)
				if err := _Elreward.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_Elreward *ElrewardFilterer) ParseBeaconUpgraded(log types.Log) (*ElrewardBeaconUpgraded, error) {
	event := new(ElrewardBeaconUpgraded)
	if err := _Elreward.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ElrewardDaoChangedIterator is returned from FilterDaoChanged and is used to iterate over the raw logs and unpacked data for DaoChanged events raised by the Elreward contract.
type ElrewardDaoChangedIterator struct {
	Event *ElrewardDaoChanged // Event containing the contract specifics and raw log

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
func (it *ElrewardDaoChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ElrewardDaoChanged)
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
		it.Event = new(ElrewardDaoChanged)
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
func (it *ElrewardDaoChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ElrewardDaoChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ElrewardDaoChanged represents a DaoChanged event raised by the Elreward contract.
type ElrewardDaoChanged struct {
	OldDao common.Address
	Dao    common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDaoChanged is a free log retrieval operation binding the contract event 0xfcde6c827a52b0870bc44ed9b10212272e18c9ea1725b772e9b493750afd8da4.
//
// Solidity: event DaoChanged(address _oldDao, address _dao)
func (_Elreward *ElrewardFilterer) FilterDaoChanged(opts *bind.FilterOpts) (*ElrewardDaoChangedIterator, error) {

	logs, sub, err := _Elreward.contract.FilterLogs(opts, "DaoChanged")
	if err != nil {
		return nil, err
	}
	return &ElrewardDaoChangedIterator{contract: _Elreward.contract, event: "DaoChanged", logs: logs, sub: sub}, nil
}

// WatchDaoChanged is a free log subscription operation binding the contract event 0xfcde6c827a52b0870bc44ed9b10212272e18c9ea1725b772e9b493750afd8da4.
//
// Solidity: event DaoChanged(address _oldDao, address _dao)
func (_Elreward *ElrewardFilterer) WatchDaoChanged(opts *bind.WatchOpts, sink chan<- *ElrewardDaoChanged) (event.Subscription, error) {

	logs, sub, err := _Elreward.contract.WatchLogs(opts, "DaoChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ElrewardDaoChanged)
				if err := _Elreward.contract.UnpackLog(event, "DaoChanged", log); err != nil {
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

// ParseDaoChanged is a log parse operation binding the contract event 0xfcde6c827a52b0870bc44ed9b10212272e18c9ea1725b772e9b493750afd8da4.
//
// Solidity: event DaoChanged(address _oldDao, address _dao)
func (_Elreward *ElrewardFilterer) ParseDaoChanged(log types.Log) (*ElrewardDaoChanged, error) {
	event := new(ElrewardDaoChanged)
	if err := _Elreward.contract.UnpackLog(event, "DaoChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ElrewardDaoRewardsClaimedIterator is returned from FilterDaoRewardsClaimed and is used to iterate over the raw logs and unpacked data for DaoRewardsClaimed events raised by the Elreward contract.
type ElrewardDaoRewardsClaimedIterator struct {
	Event *ElrewardDaoRewardsClaimed // Event containing the contract specifics and raw log

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
func (it *ElrewardDaoRewardsClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ElrewardDaoRewardsClaimed)
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
		it.Event = new(ElrewardDaoRewardsClaimed)
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
func (it *ElrewardDaoRewardsClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ElrewardDaoRewardsClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ElrewardDaoRewardsClaimed represents a DaoRewardsClaimed event raised by the Elreward contract.
type ElrewardDaoRewardsClaimed struct {
	To         common.Address
	DaoRewards *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDaoRewardsClaimed is a free log retrieval operation binding the contract event 0x9983c1b42e93483099d1468c7d059b296a20f7e9abe05f6a95742627829e3a8e.
//
// Solidity: event DaoRewardsClaimed(address _to, uint256 _daoRewards)
func (_Elreward *ElrewardFilterer) FilterDaoRewardsClaimed(opts *bind.FilterOpts) (*ElrewardDaoRewardsClaimedIterator, error) {

	logs, sub, err := _Elreward.contract.FilterLogs(opts, "DaoRewardsClaimed")
	if err != nil {
		return nil, err
	}
	return &ElrewardDaoRewardsClaimedIterator{contract: _Elreward.contract, event: "DaoRewardsClaimed", logs: logs, sub: sub}, nil
}

// WatchDaoRewardsClaimed is a free log subscription operation binding the contract event 0x9983c1b42e93483099d1468c7d059b296a20f7e9abe05f6a95742627829e3a8e.
//
// Solidity: event DaoRewardsClaimed(address _to, uint256 _daoRewards)
func (_Elreward *ElrewardFilterer) WatchDaoRewardsClaimed(opts *bind.WatchOpts, sink chan<- *ElrewardDaoRewardsClaimed) (event.Subscription, error) {

	logs, sub, err := _Elreward.contract.WatchLogs(opts, "DaoRewardsClaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ElrewardDaoRewardsClaimed)
				if err := _Elreward.contract.UnpackLog(event, "DaoRewardsClaimed", log); err != nil {
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

// ParseDaoRewardsClaimed is a log parse operation binding the contract event 0x9983c1b42e93483099d1468c7d059b296a20f7e9abe05f6a95742627829e3a8e.
//
// Solidity: event DaoRewardsClaimed(address _to, uint256 _daoRewards)
func (_Elreward *ElrewardFilterer) ParseDaoRewardsClaimed(log types.Log) (*ElrewardDaoRewardsClaimed, error) {
	event := new(ElrewardDaoRewardsClaimed)
	if err := _Elreward.contract.UnpackLog(event, "DaoRewardsClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ElrewardDaoTreasuryChangedIterator is returned from FilterDaoTreasuryChanged and is used to iterate over the raw logs and unpacked data for DaoTreasuryChanged events raised by the Elreward contract.
type ElrewardDaoTreasuryChangedIterator struct {
	Event *ElrewardDaoTreasuryChanged // Event containing the contract specifics and raw log

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
func (it *ElrewardDaoTreasuryChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ElrewardDaoTreasuryChanged)
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
		it.Event = new(ElrewardDaoTreasuryChanged)
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
func (it *ElrewardDaoTreasuryChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ElrewardDaoTreasuryChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ElrewardDaoTreasuryChanged represents a DaoTreasuryChanged event raised by the Elreward contract.
type ElrewardDaoTreasuryChanged struct {
	OldDaoTreasury common.Address
	DaoTreasury    common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDaoTreasuryChanged is a free log retrieval operation binding the contract event 0x2e8287a8375ab89afc23865fa94ab0b1ba259061c9b31397e193ca114161676d.
//
// Solidity: event DaoTreasuryChanged(address _oldDaoTreasury, address _daoTreasury)
func (_Elreward *ElrewardFilterer) FilterDaoTreasuryChanged(opts *bind.FilterOpts) (*ElrewardDaoTreasuryChangedIterator, error) {

	logs, sub, err := _Elreward.contract.FilterLogs(opts, "DaoTreasuryChanged")
	if err != nil {
		return nil, err
	}
	return &ElrewardDaoTreasuryChangedIterator{contract: _Elreward.contract, event: "DaoTreasuryChanged", logs: logs, sub: sub}, nil
}

// WatchDaoTreasuryChanged is a free log subscription operation binding the contract event 0x2e8287a8375ab89afc23865fa94ab0b1ba259061c9b31397e193ca114161676d.
//
// Solidity: event DaoTreasuryChanged(address _oldDaoTreasury, address _daoTreasury)
func (_Elreward *ElrewardFilterer) WatchDaoTreasuryChanged(opts *bind.WatchOpts, sink chan<- *ElrewardDaoTreasuryChanged) (event.Subscription, error) {

	logs, sub, err := _Elreward.contract.WatchLogs(opts, "DaoTreasuryChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ElrewardDaoTreasuryChanged)
				if err := _Elreward.contract.UnpackLog(event, "DaoTreasuryChanged", log); err != nil {
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

// ParseDaoTreasuryChanged is a log parse operation binding the contract event 0x2e8287a8375ab89afc23865fa94ab0b1ba259061c9b31397e193ca114161676d.
//
// Solidity: event DaoTreasuryChanged(address _oldDaoTreasury, address _daoTreasury)
func (_Elreward *ElrewardFilterer) ParseDaoTreasuryChanged(log types.Log) (*ElrewardDaoTreasuryChanged, error) {
	event := new(ElrewardDaoTreasuryChanged)
	if err := _Elreward.contract.UnpackLog(event, "DaoTreasuryChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ElrewardInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Elreward contract.
type ElrewardInitializedIterator struct {
	Event *ElrewardInitialized // Event containing the contract specifics and raw log

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
func (it *ElrewardInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ElrewardInitialized)
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
		it.Event = new(ElrewardInitialized)
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
func (it *ElrewardInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ElrewardInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ElrewardInitialized represents a Initialized event raised by the Elreward contract.
type ElrewardInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Elreward *ElrewardFilterer) FilterInitialized(opts *bind.FilterOpts) (*ElrewardInitializedIterator, error) {

	logs, sub, err := _Elreward.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ElrewardInitializedIterator{contract: _Elreward.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Elreward *ElrewardFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ElrewardInitialized) (event.Subscription, error) {

	logs, sub, err := _Elreward.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ElrewardInitialized)
				if err := _Elreward.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Elreward *ElrewardFilterer) ParseInitialized(log types.Log) (*ElrewardInitialized, error) {
	event := new(ElrewardInitialized)
	if err := _Elreward.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ElrewardOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Elreward contract.
type ElrewardOwnershipTransferredIterator struct {
	Event *ElrewardOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ElrewardOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ElrewardOwnershipTransferred)
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
		it.Event = new(ElrewardOwnershipTransferred)
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
func (it *ElrewardOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ElrewardOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ElrewardOwnershipTransferred represents a OwnershipTransferred event raised by the Elreward contract.
type ElrewardOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Elreward *ElrewardFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ElrewardOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Elreward.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ElrewardOwnershipTransferredIterator{contract: _Elreward.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Elreward *ElrewardFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ElrewardOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Elreward.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ElrewardOwnershipTransferred)
				if err := _Elreward.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Elreward *ElrewardFilterer) ParseOwnershipTransferred(log types.Log) (*ElrewardOwnershipTransferred, error) {
	event := new(ElrewardOwnershipTransferred)
	if err := _Elreward.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ElrewardPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Elreward contract.
type ElrewardPausedIterator struct {
	Event *ElrewardPaused // Event containing the contract specifics and raw log

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
func (it *ElrewardPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ElrewardPaused)
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
		it.Event = new(ElrewardPaused)
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
func (it *ElrewardPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ElrewardPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ElrewardPaused represents a Paused event raised by the Elreward contract.
type ElrewardPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Elreward *ElrewardFilterer) FilterPaused(opts *bind.FilterOpts) (*ElrewardPausedIterator, error) {

	logs, sub, err := _Elreward.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &ElrewardPausedIterator{contract: _Elreward.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Elreward *ElrewardFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ElrewardPaused) (event.Subscription, error) {

	logs, sub, err := _Elreward.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ElrewardPaused)
				if err := _Elreward.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Elreward *ElrewardFilterer) ParsePaused(log types.Log) (*ElrewardPaused, error) {
	event := new(ElrewardPaused)
	if err := _Elreward.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ElrewardPoolChangedIterator is returned from FilterPoolChanged and is used to iterate over the raw logs and unpacked data for PoolChanged events raised by the Elreward contract.
type ElrewardPoolChangedIterator struct {
	Event *ElrewardPoolChanged // Event containing the contract specifics and raw log

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
func (it *ElrewardPoolChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ElrewardPoolChanged)
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
		it.Event = new(ElrewardPoolChanged)
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
func (it *ElrewardPoolChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ElrewardPoolChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ElrewardPoolChanged represents a PoolChanged event raised by the Elreward contract.
type ElrewardPoolChanged struct {
	OldPool common.Address
	Pool    common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPoolChanged is a free log retrieval operation binding the contract event 0x89f24f59f1f74b74999221ad0a9ab9b4d1d2b27bddbf6f91b0c773ca0f940643.
//
// Solidity: event PoolChanged(address _oldPool, address _pool)
func (_Elreward *ElrewardFilterer) FilterPoolChanged(opts *bind.FilterOpts) (*ElrewardPoolChangedIterator, error) {

	logs, sub, err := _Elreward.contract.FilterLogs(opts, "PoolChanged")
	if err != nil {
		return nil, err
	}
	return &ElrewardPoolChangedIterator{contract: _Elreward.contract, event: "PoolChanged", logs: logs, sub: sub}, nil
}

// WatchPoolChanged is a free log subscription operation binding the contract event 0x89f24f59f1f74b74999221ad0a9ab9b4d1d2b27bddbf6f91b0c773ca0f940643.
//
// Solidity: event PoolChanged(address _oldPool, address _pool)
func (_Elreward *ElrewardFilterer) WatchPoolChanged(opts *bind.WatchOpts, sink chan<- *ElrewardPoolChanged) (event.Subscription, error) {

	logs, sub, err := _Elreward.contract.WatchLogs(opts, "PoolChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ElrewardPoolChanged)
				if err := _Elreward.contract.UnpackLog(event, "PoolChanged", log); err != nil {
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

// ParsePoolChanged is a log parse operation binding the contract event 0x89f24f59f1f74b74999221ad0a9ab9b4d1d2b27bddbf6f91b0c773ca0f940643.
//
// Solidity: event PoolChanged(address _oldPool, address _pool)
func (_Elreward *ElrewardFilterer) ParsePoolChanged(log types.Log) (*ElrewardPoolChanged, error) {
	event := new(ElrewardPoolChanged)
	if err := _Elreward.contract.UnpackLog(event, "PoolChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ElrewardPoolConfigChangedIterator is returned from FilterPoolConfigChanged and is used to iterate over the raw logs and unpacked data for PoolConfigChanged events raised by the Elreward contract.
type ElrewardPoolConfigChangedIterator struct {
	Event *ElrewardPoolConfigChanged // Event containing the contract specifics and raw log

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
func (it *ElrewardPoolConfigChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ElrewardPoolConfigChanged)
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
		it.Event = new(ElrewardPoolConfigChanged)
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
func (it *ElrewardPoolConfigChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ElrewardPoolConfigChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ElrewardPoolConfigChanged represents a PoolConfigChanged event raised by the Elreward contract.
type ElrewardPoolConfigChanged struct {
	OldOperatorRegistry common.Address
	OperatorRegistry    common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterPoolConfigChanged is a free log retrieval operation binding the contract event 0x3b4ee357b2324b710644b79f0128178b3d3d4151d9886318feedad4d78aa8e5e.
//
// Solidity: event PoolConfigChanged(address _oldOperatorRegistry, address _operatorRegistry)
func (_Elreward *ElrewardFilterer) FilterPoolConfigChanged(opts *bind.FilterOpts) (*ElrewardPoolConfigChangedIterator, error) {

	logs, sub, err := _Elreward.contract.FilterLogs(opts, "PoolConfigChanged")
	if err != nil {
		return nil, err
	}
	return &ElrewardPoolConfigChangedIterator{contract: _Elreward.contract, event: "PoolConfigChanged", logs: logs, sub: sub}, nil
}

// WatchPoolConfigChanged is a free log subscription operation binding the contract event 0x3b4ee357b2324b710644b79f0128178b3d3d4151d9886318feedad4d78aa8e5e.
//
// Solidity: event PoolConfigChanged(address _oldOperatorRegistry, address _operatorRegistry)
func (_Elreward *ElrewardFilterer) WatchPoolConfigChanged(opts *bind.WatchOpts, sink chan<- *ElrewardPoolConfigChanged) (event.Subscription, error) {

	logs, sub, err := _Elreward.contract.WatchLogs(opts, "PoolConfigChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ElrewardPoolConfigChanged)
				if err := _Elreward.contract.UnpackLog(event, "PoolConfigChanged", log); err != nil {
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

// ParsePoolConfigChanged is a log parse operation binding the contract event 0x3b4ee357b2324b710644b79f0128178b3d3d4151d9886318feedad4d78aa8e5e.
//
// Solidity: event PoolConfigChanged(address _oldOperatorRegistry, address _operatorRegistry)
func (_Elreward *ElrewardFilterer) ParsePoolConfigChanged(log types.Log) (*ElrewardPoolConfigChanged, error) {
	event := new(ElrewardPoolConfigChanged)
	if err := _Elreward.contract.UnpackLog(event, "PoolConfigChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ElrewardPoolSetIterator is returned from FilterPoolSet and is used to iterate over the raw logs and unpacked data for PoolSet events raised by the Elreward contract.
type ElrewardPoolSetIterator struct {
	Event *ElrewardPoolSet // Event containing the contract specifics and raw log

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
func (it *ElrewardPoolSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ElrewardPoolSet)
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
		it.Event = new(ElrewardPoolSet)
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
func (it *ElrewardPoolSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ElrewardPoolSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ElrewardPoolSet represents a PoolSet event raised by the Elreward contract.
type ElrewardPoolSet struct {
	Pool common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPoolSet is a free log retrieval operation binding the contract event 0x025f89b99c8ce32af8da7624f4575b920a86ebf07870d85a9fb545fee349ddce.
//
// Solidity: event PoolSet(address _pool)
func (_Elreward *ElrewardFilterer) FilterPoolSet(opts *bind.FilterOpts) (*ElrewardPoolSetIterator, error) {

	logs, sub, err := _Elreward.contract.FilterLogs(opts, "PoolSet")
	if err != nil {
		return nil, err
	}
	return &ElrewardPoolSetIterator{contract: _Elreward.contract, event: "PoolSet", logs: logs, sub: sub}, nil
}

// WatchPoolSet is a free log subscription operation binding the contract event 0x025f89b99c8ce32af8da7624f4575b920a86ebf07870d85a9fb545fee349ddce.
//
// Solidity: event PoolSet(address _pool)
func (_Elreward *ElrewardFilterer) WatchPoolSet(opts *bind.WatchOpts, sink chan<- *ElrewardPoolSet) (event.Subscription, error) {

	logs, sub, err := _Elreward.contract.WatchLogs(opts, "PoolSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ElrewardPoolSet)
				if err := _Elreward.contract.UnpackLog(event, "PoolSet", log); err != nil {
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

// ParsePoolSet is a log parse operation binding the contract event 0x025f89b99c8ce32af8da7624f4575b920a86ebf07870d85a9fb545fee349ddce.
//
// Solidity: event PoolSet(address _pool)
func (_Elreward *ElrewardFilterer) ParsePoolSet(log types.Log) (*ElrewardPoolSet, error) {
	event := new(ElrewardPoolSet)
	if err := _Elreward.contract.UnpackLog(event, "PoolSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ElrewardReceivedIterator is returned from FilterReceived and is used to iterate over the raw logs and unpacked data for Received events raised by the Elreward contract.
type ElrewardReceivedIterator struct {
	Event *ElrewardReceived // Event containing the contract specifics and raw log

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
func (it *ElrewardReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ElrewardReceived)
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
		it.Event = new(ElrewardReceived)
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
func (it *ElrewardReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ElrewardReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ElrewardReceived represents a Received event raised by the Elreward contract.
type ElrewardReceived struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReceived is a free log retrieval operation binding the contract event 0xa8142743f8f70a4c26f3691cf4ed59718381fb2f18070ec52be1f1022d855557.
//
// Solidity: event Received(uint256 _amount)
func (_Elreward *ElrewardFilterer) FilterReceived(opts *bind.FilterOpts) (*ElrewardReceivedIterator, error) {

	logs, sub, err := _Elreward.contract.FilterLogs(opts, "Received")
	if err != nil {
		return nil, err
	}
	return &ElrewardReceivedIterator{contract: _Elreward.contract, event: "Received", logs: logs, sub: sub}, nil
}

// WatchReceived is a free log subscription operation binding the contract event 0xa8142743f8f70a4c26f3691cf4ed59718381fb2f18070ec52be1f1022d855557.
//
// Solidity: event Received(uint256 _amount)
func (_Elreward *ElrewardFilterer) WatchReceived(opts *bind.WatchOpts, sink chan<- *ElrewardReceived) (event.Subscription, error) {

	logs, sub, err := _Elreward.contract.WatchLogs(opts, "Received")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ElrewardReceived)
				if err := _Elreward.contract.UnpackLog(event, "Received", log); err != nil {
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

// ParseReceived is a log parse operation binding the contract event 0xa8142743f8f70a4c26f3691cf4ed59718381fb2f18070ec52be1f1022d855557.
//
// Solidity: event Received(uint256 _amount)
func (_Elreward *ElrewardFilterer) ParseReceived(log types.Log) (*ElrewardReceived, error) {
	event := new(ElrewardReceived)
	if err := _Elreward.contract.UnpackLog(event, "Received", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ElrewardSettleIterator is returned from FilterSettle and is used to iterate over the raw logs and unpacked data for Settle events raised by the Elreward contract.
type ElrewardSettleIterator struct {
	Event *ElrewardSettle // Event containing the contract specifics and raw log

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
func (it *ElrewardSettleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ElrewardSettle)
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
		it.Event = new(ElrewardSettle)
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
func (it *ElrewardSettleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ElrewardSettleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ElrewardSettle represents a Settle event raised by the Elreward contract.
type ElrewardSettle struct {
	DaoRewards  *big.Int
	PoolRewards *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSettle is a free log retrieval operation binding the contract event 0x88a84ea6dd274b386afd27dbbe11b6192b25017f5e60bb8c4053dfddb45c294d.
//
// Solidity: event Settle(uint256 _daoRewards, uint256 _poolRewards)
func (_Elreward *ElrewardFilterer) FilterSettle(opts *bind.FilterOpts) (*ElrewardSettleIterator, error) {

	logs, sub, err := _Elreward.contract.FilterLogs(opts, "Settle")
	if err != nil {
		return nil, err
	}
	return &ElrewardSettleIterator{contract: _Elreward.contract, event: "Settle", logs: logs, sub: sub}, nil
}

// WatchSettle is a free log subscription operation binding the contract event 0x88a84ea6dd274b386afd27dbbe11b6192b25017f5e60bb8c4053dfddb45c294d.
//
// Solidity: event Settle(uint256 _daoRewards, uint256 _poolRewards)
func (_Elreward *ElrewardFilterer) WatchSettle(opts *bind.WatchOpts, sink chan<- *ElrewardSettle) (event.Subscription, error) {

	logs, sub, err := _Elreward.contract.WatchLogs(opts, "Settle")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ElrewardSettle)
				if err := _Elreward.contract.UnpackLog(event, "Settle", log); err != nil {
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

// ParseSettle is a log parse operation binding the contract event 0x88a84ea6dd274b386afd27dbbe11b6192b25017f5e60bb8c4053dfddb45c294d.
//
// Solidity: event Settle(uint256 _daoRewards, uint256 _poolRewards)
func (_Elreward *ElrewardFilterer) ParseSettle(log types.Log) (*ElrewardSettle, error) {
	event := new(ElrewardSettle)
	if err := _Elreward.contract.UnpackLog(event, "Settle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ElrewardTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Elreward contract.
type ElrewardTransferIterator struct {
	Event *ElrewardTransfer // Event containing the contract specifics and raw log

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
func (it *ElrewardTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ElrewardTransfer)
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
		it.Event = new(ElrewardTransfer)
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
func (it *ElrewardTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ElrewardTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ElrewardTransfer represents a Transfer event raised by the Elreward contract.
type ElrewardTransfer struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0x69ca02dd4edd7bf0a4abb9ed3b7af3f14778db5d61921c7dc7cd545266326de2.
//
// Solidity: event Transfer(address _to, uint256 _amount)
func (_Elreward *ElrewardFilterer) FilterTransfer(opts *bind.FilterOpts) (*ElrewardTransferIterator, error) {

	logs, sub, err := _Elreward.contract.FilterLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return &ElrewardTransferIterator{contract: _Elreward.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0x69ca02dd4edd7bf0a4abb9ed3b7af3f14778db5d61921c7dc7cd545266326de2.
//
// Solidity: event Transfer(address _to, uint256 _amount)
func (_Elreward *ElrewardFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ElrewardTransfer) (event.Subscription, error) {

	logs, sub, err := _Elreward.contract.WatchLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ElrewardTransfer)
				if err := _Elreward.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0x69ca02dd4edd7bf0a4abb9ed3b7af3f14778db5d61921c7dc7cd545266326de2.
//
// Solidity: event Transfer(address _to, uint256 _amount)
func (_Elreward *ElrewardFilterer) ParseTransfer(log types.Log) (*ElrewardTransfer, error) {
	event := new(ElrewardTransfer)
	if err := _Elreward.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ElrewardUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Elreward contract.
type ElrewardUnpausedIterator struct {
	Event *ElrewardUnpaused // Event containing the contract specifics and raw log

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
func (it *ElrewardUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ElrewardUnpaused)
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
		it.Event = new(ElrewardUnpaused)
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
func (it *ElrewardUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ElrewardUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ElrewardUnpaused represents a Unpaused event raised by the Elreward contract.
type ElrewardUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Elreward *ElrewardFilterer) FilterUnpaused(opts *bind.FilterOpts) (*ElrewardUnpausedIterator, error) {

	logs, sub, err := _Elreward.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &ElrewardUnpausedIterator{contract: _Elreward.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Elreward *ElrewardFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ElrewardUnpaused) (event.Subscription, error) {

	logs, sub, err := _Elreward.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ElrewardUnpaused)
				if err := _Elreward.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_Elreward *ElrewardFilterer) ParseUnpaused(log types.Log) (*ElrewardUnpaused, error) {
	event := new(ElrewardUnpaused)
	if err := _Elreward.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ElrewardUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Elreward contract.
type ElrewardUpgradedIterator struct {
	Event *ElrewardUpgraded // Event containing the contract specifics and raw log

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
func (it *ElrewardUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ElrewardUpgraded)
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
		it.Event = new(ElrewardUpgraded)
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
func (it *ElrewardUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ElrewardUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ElrewardUpgraded represents a Upgraded event raised by the Elreward contract.
type ElrewardUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Elreward *ElrewardFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*ElrewardUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Elreward.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &ElrewardUpgradedIterator{contract: _Elreward.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Elreward *ElrewardFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *ElrewardUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Elreward.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ElrewardUpgraded)
				if err := _Elreward.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_Elreward *ElrewardFilterer) ParseUpgraded(log types.Log) (*ElrewardUpgraded, error) {
	event := new(ElrewardUpgraded)
	if err := _Elreward.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
