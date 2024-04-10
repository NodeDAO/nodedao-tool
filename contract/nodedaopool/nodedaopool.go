// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package nodedaopool

import (
	"errors"
	"fmt"
	"math/big"
	"os"
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

// WithdrawalRequestWithdrawalInfo is an auto generated low-level Go binding around an user-defined struct.
type WithdrawalRequestWithdrawalInfo struct {
	WithdrawalHeight   *big.Int
	WithdrawalExchange *big.Int
	IsClaim            uint64
	WithdrawalAmount   *big.Int
	ClaimAmount        *big.Int
}

// PoolMetaData contains all meta data concerning the Pool contract.
var PoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CanUnstakeETH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DelayTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DepositRootMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EigenPodMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientFunds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidApr\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidParameter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRequestId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidVersion\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidtypeId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PermissionDenied\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RestakingPodNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnstakeNotEnabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UpdateTimelocked\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WithrawalsRequestCannotClaimed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_oldApr\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_apr\",\"type\":\"uint256\"}],\"name\":\"AprUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalUnderlyingAsset\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_estimatedRewards\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"AssetsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_oldDao\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_dao\",\"type\":\"address\"}],\"name\":\"DaoChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_stakeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_mintAmount\",\"type\":\"uint256\"}],\"name\":\"EthStake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_unstakeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_ethAmount\",\"type\":\"uint256\"}],\"name\":\"EthUnstake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_oldAprManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_aprManager\",\"type\":\"address\"}],\"name\":\"RateManagerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Received\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_restakingPod\",\"type\":\"address\"}],\"name\":\"RestakingPodAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_oldUnstakeAllowed\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_unstakeAllowed\",\"type\":\"bool\"}],\"name\":\"UnstakeAllowedUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_oldValidatorManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_validatorManager\",\"type\":\"address\"}],\"name\":\"ValidatorManagerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"_pubkeys\",\"type\":\"bytes[]\"}],\"name\":\"ValidatorRegistration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_oldValidatorManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_validatorManager\",\"type\":\"address\"}],\"name\":\"ValidatorRegistryChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_oldWithdrawalDelayBlocks\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_withdrawalDelayBlocks\",\"type\":\"uint256\"}],\"name\":\"WithdrawalDelayChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_claimAmount\",\"type\":\"uint256\"}],\"name\":\"WithdrawalsClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_withdrawalAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"WithdrawalsRequest\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"APR_PERCENT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BLOCK_NUMBER_PER_YEAR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_APR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_WITHDRAWAL_DELAY_BLOCKS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPDATE_BLOCK_LIMIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_withdrawalDelayBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_apr\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalUnderlyingAsset\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_dao\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_poolToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rateManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_validatorManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_validatorRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_depositContract\",\"type\":\"address\"}],\"name\":\"__BasePool_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_restakingPod\",\"type\":\"address\"}],\"name\":\"addRestakingPod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"canClaimWithdrawal\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimDelayedWithdrawals\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_requestIds\",\"type\":\"uint256[]\"}],\"name\":\"claimWithdrawals\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_unstakeAmount\",\"type\":\"uint256\"}],\"name\":\"convertToAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stakeAmount\",\"type\":\"uint256\"}],\"name\":\"convertToShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentApr\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dao\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositContract\",\"outputs\":[{\"internalType\":\"contractIDepositContract\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eigenLayerEigenPodManager\",\"outputs\":[{\"internalType\":\"contractIEigenPodManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"elRewardsAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exchangeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCLVaultAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRestakingPods\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"getUserWithdrawals\",\"outputs\":[{\"components\":[{\"internalType\":\"uint96\",\"name\":\"withdrawalHeight\",\"type\":\"uint96\"},{\"internalType\":\"uint96\",\"name\":\"withdrawalExchange\",\"type\":\"uint96\"},{\"internalType\":\"uint64\",\"name\":\"isClaim\",\"type\":\"uint64\"},{\"internalType\":\"uint128\",\"name\":\"withdrawalAmount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"claimAmount\",\"type\":\"uint128\"}],\"internalType\":\"structWithdrawalRequest.WithdrawalInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_apr\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_dao\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_elRewardsAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_poolToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rateManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_validatorManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_validatorRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_depositContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_eigenLayerEigenPodManager\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_restakingPods\",\"type\":\"address[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolToken\",\"outputs\":[{\"internalType\":\"contractILsdETH\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rateManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"receiveRewards\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_depositContractRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_restakingPod\",\"type\":\"address\"},{\"internalType\":\"bytes[]\",\"name\":\"_pubkeys\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_signatures\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_depositDataRoots\",\"type\":\"bytes32[]\"}],\"name\":\"registerValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_unstakeAmount\",\"type\":\"uint256\"}],\"name\":\"requestWithdrawals\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardsUpdateBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rateManager\",\"type\":\"address\"}],\"name\":\"setRateManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_unstakeAllowed\",\"type\":\"bool\"}],\"name\":\"setUnstakeAllowed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validatorManager\",\"type\":\"address\"}],\"name\":\"setValidatorManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validatorRegistry\",\"type\":\"address\"}],\"name\":\"setValidatorRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_withdrawalDelayBlocks\",\"type\":\"uint256\"}],\"name\":\"setWithdrawalDelayBlocks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalUnderlyingAsset\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalWithdrawalAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unstakeAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_unstakeAmount\",\"type\":\"uint256\"}],\"name\":\"unstakeETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_apr\",\"type\":\"uint256\"}],\"name\":\"updateApr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorRegistry\",\"outputs\":[{\"internalType\":\"contractIValidatorRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawCredentials\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawalDelayBlocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// PoolABI is the input ABI used to generate the binding from.
// Deprecated: Use PoolMetaData.ABI instead.
var PoolABI = PoolMetaData.ABI
var NodeDaoPoolABI abi.ABI

func init() {
	var err error
	NodeDaoPoolABI, err = abi.JSON(strings.NewReader(PoolABI))
	if err != nil {
		fmt.Println("err: ", err.Error())
		os.Exit(1)
	}
}

// Pool is an auto generated Go binding around an Ethereum contract.
type Pool struct {
	PoolCaller     // Read-only binding to the contract
	PoolTransactor // Write-only binding to the contract
	PoolFilterer   // Log filterer for contract events
}

// PoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type PoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PoolSession struct {
	Contract     *Pool             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PoolCallerSession struct {
	Contract *PoolCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PoolTransactorSession struct {
	Contract     *PoolTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type PoolRaw struct {
	Contract *Pool // Generic contract binding to access the raw methods on
}

// PoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PoolCallerRaw struct {
	Contract *PoolCaller // Generic read-only contract binding to access the raw methods on
}

// PoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PoolTransactorRaw struct {
	Contract *PoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPool creates a new instance of Pool, bound to a specific deployed contract.
func NewPool(address common.Address, backend bind.ContractBackend) (*Pool, error) {
	contract, err := bindPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pool{PoolCaller: PoolCaller{contract: contract}, PoolTransactor: PoolTransactor{contract: contract}, PoolFilterer: PoolFilterer{contract: contract}}, nil
}

// NewPoolCaller creates a new read-only instance of Pool, bound to a specific deployed contract.
func NewPoolCaller(address common.Address, caller bind.ContractCaller) (*PoolCaller, error) {
	contract, err := bindPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PoolCaller{contract: contract}, nil
}

// NewPoolTransactor creates a new write-only instance of Pool, bound to a specific deployed contract.
func NewPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*PoolTransactor, error) {
	contract, err := bindPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PoolTransactor{contract: contract}, nil
}

// NewPoolFilterer creates a new log filterer instance of Pool, bound to a specific deployed contract.
func NewPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*PoolFilterer, error) {
	contract, err := bindPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PoolFilterer{contract: contract}, nil
}

// bindPool binds a generic wrapper to an already deployed contract.
func bindPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pool *PoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pool.Contract.PoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pool *PoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool.Contract.PoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pool *PoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pool.Contract.PoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pool *PoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pool *PoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pool *PoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pool.Contract.contract.Transact(opts, method, params...)
}

// APRPERCENT is a free data retrieval call binding the contract method 0x99ca63bb.
//
// Solidity: function APR_PERCENT() view returns(uint256)
func (_Pool *PoolCaller) APRPERCENT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "APR_PERCENT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// APRPERCENT is a free data retrieval call binding the contract method 0x99ca63bb.
//
// Solidity: function APR_PERCENT() view returns(uint256)
func (_Pool *PoolSession) APRPERCENT() (*big.Int, error) {
	return _Pool.Contract.APRPERCENT(&_Pool.CallOpts)
}

// APRPERCENT is a free data retrieval call binding the contract method 0x99ca63bb.
//
// Solidity: function APR_PERCENT() view returns(uint256)
func (_Pool *PoolCallerSession) APRPERCENT() (*big.Int, error) {
	return _Pool.Contract.APRPERCENT(&_Pool.CallOpts)
}

// BLOCKNUMBERPERYEAR is a free data retrieval call binding the contract method 0xef213ddc.
//
// Solidity: function BLOCK_NUMBER_PER_YEAR() view returns(uint256)
func (_Pool *PoolCaller) BLOCKNUMBERPERYEAR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "BLOCK_NUMBER_PER_YEAR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BLOCKNUMBERPERYEAR is a free data retrieval call binding the contract method 0xef213ddc.
//
// Solidity: function BLOCK_NUMBER_PER_YEAR() view returns(uint256)
func (_Pool *PoolSession) BLOCKNUMBERPERYEAR() (*big.Int, error) {
	return _Pool.Contract.BLOCKNUMBERPERYEAR(&_Pool.CallOpts)
}

// BLOCKNUMBERPERYEAR is a free data retrieval call binding the contract method 0xef213ddc.
//
// Solidity: function BLOCK_NUMBER_PER_YEAR() view returns(uint256)
func (_Pool *PoolCallerSession) BLOCKNUMBERPERYEAR() (*big.Int, error) {
	return _Pool.Contract.BLOCKNUMBERPERYEAR(&_Pool.CallOpts)
}

// MAXAPR is a free data retrieval call binding the contract method 0x86d8f78d.
//
// Solidity: function MAX_APR() view returns(uint256)
func (_Pool *PoolCaller) MAXAPR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "MAX_APR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXAPR is a free data retrieval call binding the contract method 0x86d8f78d.
//
// Solidity: function MAX_APR() view returns(uint256)
func (_Pool *PoolSession) MAXAPR() (*big.Int, error) {
	return _Pool.Contract.MAXAPR(&_Pool.CallOpts)
}

// MAXAPR is a free data retrieval call binding the contract method 0x86d8f78d.
//
// Solidity: function MAX_APR() view returns(uint256)
func (_Pool *PoolCallerSession) MAXAPR() (*big.Int, error) {
	return _Pool.Contract.MAXAPR(&_Pool.CallOpts)
}

// MAXWITHDRAWALDELAYBLOCKS is a free data retrieval call binding the contract method 0xca661c04.
//
// Solidity: function MAX_WITHDRAWAL_DELAY_BLOCKS() view returns(uint256)
func (_Pool *PoolCaller) MAXWITHDRAWALDELAYBLOCKS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "MAX_WITHDRAWAL_DELAY_BLOCKS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXWITHDRAWALDELAYBLOCKS is a free data retrieval call binding the contract method 0xca661c04.
//
// Solidity: function MAX_WITHDRAWAL_DELAY_BLOCKS() view returns(uint256)
func (_Pool *PoolSession) MAXWITHDRAWALDELAYBLOCKS() (*big.Int, error) {
	return _Pool.Contract.MAXWITHDRAWALDELAYBLOCKS(&_Pool.CallOpts)
}

// MAXWITHDRAWALDELAYBLOCKS is a free data retrieval call binding the contract method 0xca661c04.
//
// Solidity: function MAX_WITHDRAWAL_DELAY_BLOCKS() view returns(uint256)
func (_Pool *PoolCallerSession) MAXWITHDRAWALDELAYBLOCKS() (*big.Int, error) {
	return _Pool.Contract.MAXWITHDRAWALDELAYBLOCKS(&_Pool.CallOpts)
}

// UPDATEBLOCKLIMIT is a free data retrieval call binding the contract method 0xb7e34b0c.
//
// Solidity: function UPDATE_BLOCK_LIMIT() view returns(uint256)
func (_Pool *PoolCaller) UPDATEBLOCKLIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "UPDATE_BLOCK_LIMIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UPDATEBLOCKLIMIT is a free data retrieval call binding the contract method 0xb7e34b0c.
//
// Solidity: function UPDATE_BLOCK_LIMIT() view returns(uint256)
func (_Pool *PoolSession) UPDATEBLOCKLIMIT() (*big.Int, error) {
	return _Pool.Contract.UPDATEBLOCKLIMIT(&_Pool.CallOpts)
}

// UPDATEBLOCKLIMIT is a free data retrieval call binding the contract method 0xb7e34b0c.
//
// Solidity: function UPDATE_BLOCK_LIMIT() view returns(uint256)
func (_Pool *PoolCallerSession) UPDATEBLOCKLIMIT() (*big.Int, error) {
	return _Pool.Contract.UPDATEBLOCKLIMIT(&_Pool.CallOpts)
}

// CanClaimWithdrawal is a free data retrieval call binding the contract method 0x78865853.
//
// Solidity: function canClaimWithdrawal(address _receiver, uint256 _requestId) view returns(bool)
func (_Pool *PoolCaller) CanClaimWithdrawal(opts *bind.CallOpts, _receiver common.Address, _requestId *big.Int) (bool, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "canClaimWithdrawal", _receiver, _requestId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanClaimWithdrawal is a free data retrieval call binding the contract method 0x78865853.
//
// Solidity: function canClaimWithdrawal(address _receiver, uint256 _requestId) view returns(bool)
func (_Pool *PoolSession) CanClaimWithdrawal(_receiver common.Address, _requestId *big.Int) (bool, error) {
	return _Pool.Contract.CanClaimWithdrawal(&_Pool.CallOpts, _receiver, _requestId)
}

// CanClaimWithdrawal is a free data retrieval call binding the contract method 0x78865853.
//
// Solidity: function canClaimWithdrawal(address _receiver, uint256 _requestId) view returns(bool)
func (_Pool *PoolCallerSession) CanClaimWithdrawal(_receiver common.Address, _requestId *big.Int) (bool, error) {
	return _Pool.Contract.CanClaimWithdrawal(&_Pool.CallOpts, _receiver, _requestId)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 _unstakeAmount) view returns(uint256)
func (_Pool *PoolCaller) ConvertToAssets(opts *bind.CallOpts, _unstakeAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "convertToAssets", _unstakeAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 _unstakeAmount) view returns(uint256)
func (_Pool *PoolSession) ConvertToAssets(_unstakeAmount *big.Int) (*big.Int, error) {
	return _Pool.Contract.ConvertToAssets(&_Pool.CallOpts, _unstakeAmount)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 _unstakeAmount) view returns(uint256)
func (_Pool *PoolCallerSession) ConvertToAssets(_unstakeAmount *big.Int) (*big.Int, error) {
	return _Pool.Contract.ConvertToAssets(&_Pool.CallOpts, _unstakeAmount)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 _stakeAmount) view returns(uint256)
func (_Pool *PoolCaller) ConvertToShares(opts *bind.CallOpts, _stakeAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "convertToShares", _stakeAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 _stakeAmount) view returns(uint256)
func (_Pool *PoolSession) ConvertToShares(_stakeAmount *big.Int) (*big.Int, error) {
	return _Pool.Contract.ConvertToShares(&_Pool.CallOpts, _stakeAmount)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 _stakeAmount) view returns(uint256)
func (_Pool *PoolCallerSession) ConvertToShares(_stakeAmount *big.Int) (*big.Int, error) {
	return _Pool.Contract.ConvertToShares(&_Pool.CallOpts, _stakeAmount)
}

// CurrentApr is a free data retrieval call binding the contract method 0x4f322ae8.
//
// Solidity: function currentApr() view returns(uint256)
func (_Pool *PoolCaller) CurrentApr(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "currentApr")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentApr is a free data retrieval call binding the contract method 0x4f322ae8.
//
// Solidity: function currentApr() view returns(uint256)
func (_Pool *PoolSession) CurrentApr() (*big.Int, error) {
	return _Pool.Contract.CurrentApr(&_Pool.CallOpts)
}

// CurrentApr is a free data retrieval call binding the contract method 0x4f322ae8.
//
// Solidity: function currentApr() view returns(uint256)
func (_Pool *PoolCallerSession) CurrentApr() (*big.Int, error) {
	return _Pool.Contract.CurrentApr(&_Pool.CallOpts)
}

// Dao is a free data retrieval call binding the contract method 0x4162169f.
//
// Solidity: function dao() view returns(address)
func (_Pool *PoolCaller) Dao(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "dao")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Dao is a free data retrieval call binding the contract method 0x4162169f.
//
// Solidity: function dao() view returns(address)
func (_Pool *PoolSession) Dao() (common.Address, error) {
	return _Pool.Contract.Dao(&_Pool.CallOpts)
}

// Dao is a free data retrieval call binding the contract method 0x4162169f.
//
// Solidity: function dao() view returns(address)
func (_Pool *PoolCallerSession) Dao() (common.Address, error) {
	return _Pool.Contract.Dao(&_Pool.CallOpts)
}

// DepositContract is a free data retrieval call binding the contract method 0xe94ad65b.
//
// Solidity: function depositContract() view returns(address)
func (_Pool *PoolCaller) DepositContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "depositContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DepositContract is a free data retrieval call binding the contract method 0xe94ad65b.
//
// Solidity: function depositContract() view returns(address)
func (_Pool *PoolSession) DepositContract() (common.Address, error) {
	return _Pool.Contract.DepositContract(&_Pool.CallOpts)
}

// DepositContract is a free data retrieval call binding the contract method 0xe94ad65b.
//
// Solidity: function depositContract() view returns(address)
func (_Pool *PoolCallerSession) DepositContract() (common.Address, error) {
	return _Pool.Contract.DepositContract(&_Pool.CallOpts)
}

// EigenLayerEigenPodManager is a free data retrieval call binding the contract method 0x2ffb004f.
//
// Solidity: function eigenLayerEigenPodManager() view returns(address)
func (_Pool *PoolCaller) EigenLayerEigenPodManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "eigenLayerEigenPodManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EigenLayerEigenPodManager is a free data retrieval call binding the contract method 0x2ffb004f.
//
// Solidity: function eigenLayerEigenPodManager() view returns(address)
func (_Pool *PoolSession) EigenLayerEigenPodManager() (common.Address, error) {
	return _Pool.Contract.EigenLayerEigenPodManager(&_Pool.CallOpts)
}

// EigenLayerEigenPodManager is a free data retrieval call binding the contract method 0x2ffb004f.
//
// Solidity: function eigenLayerEigenPodManager() view returns(address)
func (_Pool *PoolCallerSession) EigenLayerEigenPodManager() (common.Address, error) {
	return _Pool.Contract.EigenLayerEigenPodManager(&_Pool.CallOpts)
}

// ElRewardsAddress is a free data retrieval call binding the contract method 0x215bd1b1.
//
// Solidity: function elRewardsAddress() view returns(address)
func (_Pool *PoolCaller) ElRewardsAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "elRewardsAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ElRewardsAddress is a free data retrieval call binding the contract method 0x215bd1b1.
//
// Solidity: function elRewardsAddress() view returns(address)
func (_Pool *PoolSession) ElRewardsAddress() (common.Address, error) {
	return _Pool.Contract.ElRewardsAddress(&_Pool.CallOpts)
}

// ElRewardsAddress is a free data retrieval call binding the contract method 0x215bd1b1.
//
// Solidity: function elRewardsAddress() view returns(address)
func (_Pool *PoolCallerSession) ElRewardsAddress() (common.Address, error) {
	return _Pool.Contract.ElRewardsAddress(&_Pool.CallOpts)
}

// ExchangeRate is a free data retrieval call binding the contract method 0x3ba0b9a9.
//
// Solidity: function exchangeRate() view returns(uint256)
func (_Pool *PoolCaller) ExchangeRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "exchangeRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExchangeRate is a free data retrieval call binding the contract method 0x3ba0b9a9.
//
// Solidity: function exchangeRate() view returns(uint256)
func (_Pool *PoolSession) ExchangeRate() (*big.Int, error) {
	return _Pool.Contract.ExchangeRate(&_Pool.CallOpts)
}

// ExchangeRate is a free data retrieval call binding the contract method 0x3ba0b9a9.
//
// Solidity: function exchangeRate() view returns(uint256)
func (_Pool *PoolCallerSession) ExchangeRate() (*big.Int, error) {
	return _Pool.Contract.ExchangeRate(&_Pool.CallOpts)
}

// GetCLVaultAmount is a free data retrieval call binding the contract method 0xf60b6fe6.
//
// Solidity: function getCLVaultAmount() view returns(uint256)
func (_Pool *PoolCaller) GetCLVaultAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getCLVaultAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCLVaultAmount is a free data retrieval call binding the contract method 0xf60b6fe6.
//
// Solidity: function getCLVaultAmount() view returns(uint256)
func (_Pool *PoolSession) GetCLVaultAmount() (*big.Int, error) {
	return _Pool.Contract.GetCLVaultAmount(&_Pool.CallOpts)
}

// GetCLVaultAmount is a free data retrieval call binding the contract method 0xf60b6fe6.
//
// Solidity: function getCLVaultAmount() view returns(uint256)
func (_Pool *PoolCallerSession) GetCLVaultAmount() (*big.Int, error) {
	return _Pool.Contract.GetCLVaultAmount(&_Pool.CallOpts)
}

// GetPoolAmount is a free data retrieval call binding the contract method 0x4ab4ba42.
//
// Solidity: function getPoolAmount() view returns(uint256)
func (_Pool *PoolCaller) GetPoolAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getPoolAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPoolAmount is a free data retrieval call binding the contract method 0x4ab4ba42.
//
// Solidity: function getPoolAmount() view returns(uint256)
func (_Pool *PoolSession) GetPoolAmount() (*big.Int, error) {
	return _Pool.Contract.GetPoolAmount(&_Pool.CallOpts)
}

// GetPoolAmount is a free data retrieval call binding the contract method 0x4ab4ba42.
//
// Solidity: function getPoolAmount() view returns(uint256)
func (_Pool *PoolCallerSession) GetPoolAmount() (*big.Int, error) {
	return _Pool.Contract.GetPoolAmount(&_Pool.CallOpts)
}

// GetRestakingPods is a free data retrieval call binding the contract method 0x61e9ecf3.
//
// Solidity: function getRestakingPods() view returns(address[])
func (_Pool *PoolCaller) GetRestakingPods(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getRestakingPods")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRestakingPods is a free data retrieval call binding the contract method 0x61e9ecf3.
//
// Solidity: function getRestakingPods() view returns(address[])
func (_Pool *PoolSession) GetRestakingPods() ([]common.Address, error) {
	return _Pool.Contract.GetRestakingPods(&_Pool.CallOpts)
}

// GetRestakingPods is a free data retrieval call binding the contract method 0x61e9ecf3.
//
// Solidity: function getRestakingPods() view returns(address[])
func (_Pool *PoolCallerSession) GetRestakingPods() ([]common.Address, error) {
	return _Pool.Contract.GetRestakingPods(&_Pool.CallOpts)
}

// GetUserWithdrawals is a free data retrieval call binding the contract method 0xe502eb68.
//
// Solidity: function getUserWithdrawals(address _receiver) view returns((uint96,uint96,uint64,uint128,uint128)[])
func (_Pool *PoolCaller) GetUserWithdrawals(opts *bind.CallOpts, _receiver common.Address) ([]WithdrawalRequestWithdrawalInfo, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getUserWithdrawals", _receiver)

	if err != nil {
		return *new([]WithdrawalRequestWithdrawalInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]WithdrawalRequestWithdrawalInfo)).(*[]WithdrawalRequestWithdrawalInfo)

	return out0, err

}

// GetUserWithdrawals is a free data retrieval call binding the contract method 0xe502eb68.
//
// Solidity: function getUserWithdrawals(address _receiver) view returns((uint96,uint96,uint64,uint128,uint128)[])
func (_Pool *PoolSession) GetUserWithdrawals(_receiver common.Address) ([]WithdrawalRequestWithdrawalInfo, error) {
	return _Pool.Contract.GetUserWithdrawals(&_Pool.CallOpts, _receiver)
}

// GetUserWithdrawals is a free data retrieval call binding the contract method 0xe502eb68.
//
// Solidity: function getUserWithdrawals(address _receiver) view returns((uint96,uint96,uint64,uint128,uint128)[])
func (_Pool *PoolCallerSession) GetUserWithdrawals(_receiver common.Address) ([]WithdrawalRequestWithdrawalInfo, error) {
	return _Pool.Contract.GetUserWithdrawals(&_Pool.CallOpts, _receiver)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_Pool *PoolCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_Pool *PoolSession) Implementation() (common.Address, error) {
	return _Pool.Contract.Implementation(&_Pool.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_Pool *PoolCallerSession) Implementation() (common.Address, error) {
	return _Pool.Contract.Implementation(&_Pool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pool *PoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pool *PoolSession) Owner() (common.Address, error) {
	return _Pool.Contract.Owner(&_Pool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pool *PoolCallerSession) Owner() (common.Address, error) {
	return _Pool.Contract.Owner(&_Pool.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Pool *PoolCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Pool *PoolSession) Paused() (bool, error) {
	return _Pool.Contract.Paused(&_Pool.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Pool *PoolCallerSession) Paused() (bool, error) {
	return _Pool.Contract.Paused(&_Pool.CallOpts)
}

// PoolToken is a free data retrieval call binding the contract method 0xcbdf382c.
//
// Solidity: function poolToken() view returns(address)
func (_Pool *PoolCaller) PoolToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "poolToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolToken is a free data retrieval call binding the contract method 0xcbdf382c.
//
// Solidity: function poolToken() view returns(address)
func (_Pool *PoolSession) PoolToken() (common.Address, error) {
	return _Pool.Contract.PoolToken(&_Pool.CallOpts)
}

// PoolToken is a free data retrieval call binding the contract method 0xcbdf382c.
//
// Solidity: function poolToken() view returns(address)
func (_Pool *PoolCallerSession) PoolToken() (common.Address, error) {
	return _Pool.Contract.PoolToken(&_Pool.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Pool *PoolCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Pool *PoolSession) ProxiableUUID() ([32]byte, error) {
	return _Pool.Contract.ProxiableUUID(&_Pool.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Pool *PoolCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Pool.Contract.ProxiableUUID(&_Pool.CallOpts)
}

// RateManager is a free data retrieval call binding the contract method 0x71c3cd88.
//
// Solidity: function rateManager() view returns(address)
func (_Pool *PoolCaller) RateManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "rateManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RateManager is a free data retrieval call binding the contract method 0x71c3cd88.
//
// Solidity: function rateManager() view returns(address)
func (_Pool *PoolSession) RateManager() (common.Address, error) {
	return _Pool.Contract.RateManager(&_Pool.CallOpts)
}

// RateManager is a free data retrieval call binding the contract method 0x71c3cd88.
//
// Solidity: function rateManager() view returns(address)
func (_Pool *PoolCallerSession) RateManager() (common.Address, error) {
	return _Pool.Contract.RateManager(&_Pool.CallOpts)
}

// RewardsUpdateBlock is a free data retrieval call binding the contract method 0x88f36cdd.
//
// Solidity: function rewardsUpdateBlock() view returns(uint256)
func (_Pool *PoolCaller) RewardsUpdateBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "rewardsUpdateBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardsUpdateBlock is a free data retrieval call binding the contract method 0x88f36cdd.
//
// Solidity: function rewardsUpdateBlock() view returns(uint256)
func (_Pool *PoolSession) RewardsUpdateBlock() (*big.Int, error) {
	return _Pool.Contract.RewardsUpdateBlock(&_Pool.CallOpts)
}

// RewardsUpdateBlock is a free data retrieval call binding the contract method 0x88f36cdd.
//
// Solidity: function rewardsUpdateBlock() view returns(uint256)
func (_Pool *PoolCallerSession) RewardsUpdateBlock() (*big.Int, error) {
	return _Pool.Contract.RewardsUpdateBlock(&_Pool.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_Pool *PoolCaller) TotalAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "totalAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_Pool *PoolSession) TotalAssets() (*big.Int, error) {
	return _Pool.Contract.TotalAssets(&_Pool.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_Pool *PoolCallerSession) TotalAssets() (*big.Int, error) {
	return _Pool.Contract.TotalAssets(&_Pool.CallOpts)
}

// TotalUnderlyingAsset is a free data retrieval call binding the contract method 0xddfa63ae.
//
// Solidity: function totalUnderlyingAsset() view returns(uint256)
func (_Pool *PoolCaller) TotalUnderlyingAsset(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "totalUnderlyingAsset")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalUnderlyingAsset is a free data retrieval call binding the contract method 0xddfa63ae.
//
// Solidity: function totalUnderlyingAsset() view returns(uint256)
func (_Pool *PoolSession) TotalUnderlyingAsset() (*big.Int, error) {
	return _Pool.Contract.TotalUnderlyingAsset(&_Pool.CallOpts)
}

// TotalUnderlyingAsset is a free data retrieval call binding the contract method 0xddfa63ae.
//
// Solidity: function totalUnderlyingAsset() view returns(uint256)
func (_Pool *PoolCallerSession) TotalUnderlyingAsset() (*big.Int, error) {
	return _Pool.Contract.TotalUnderlyingAsset(&_Pool.CallOpts)
}

// TotalWithdrawalAmount is a free data retrieval call binding the contract method 0xe23e0e08.
//
// Solidity: function totalWithdrawalAmount() view returns(uint256)
func (_Pool *PoolCaller) TotalWithdrawalAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "totalWithdrawalAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalWithdrawalAmount is a free data retrieval call binding the contract method 0xe23e0e08.
//
// Solidity: function totalWithdrawalAmount() view returns(uint256)
func (_Pool *PoolSession) TotalWithdrawalAmount() (*big.Int, error) {
	return _Pool.Contract.TotalWithdrawalAmount(&_Pool.CallOpts)
}

// TotalWithdrawalAmount is a free data retrieval call binding the contract method 0xe23e0e08.
//
// Solidity: function totalWithdrawalAmount() view returns(uint256)
func (_Pool *PoolCallerSession) TotalWithdrawalAmount() (*big.Int, error) {
	return _Pool.Contract.TotalWithdrawalAmount(&_Pool.CallOpts)
}

// TypeId is a free data retrieval call binding the contract method 0x8f940f63.
//
// Solidity: function typeId() pure returns(bytes32)
func (_Pool *PoolCaller) TypeId(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "typeId")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TypeId is a free data retrieval call binding the contract method 0x8f940f63.
//
// Solidity: function typeId() pure returns(bytes32)
func (_Pool *PoolSession) TypeId() ([32]byte, error) {
	return _Pool.Contract.TypeId(&_Pool.CallOpts)
}

// TypeId is a free data retrieval call binding the contract method 0x8f940f63.
//
// Solidity: function typeId() pure returns(bytes32)
func (_Pool *PoolCallerSession) TypeId() ([32]byte, error) {
	return _Pool.Contract.TypeId(&_Pool.CallOpts)
}

// UnstakeAllowed is a free data retrieval call binding the contract method 0xee018b6d.
//
// Solidity: function unstakeAllowed() view returns(bool)
func (_Pool *PoolCaller) UnstakeAllowed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "unstakeAllowed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UnstakeAllowed is a free data retrieval call binding the contract method 0xee018b6d.
//
// Solidity: function unstakeAllowed() view returns(bool)
func (_Pool *PoolSession) UnstakeAllowed() (bool, error) {
	return _Pool.Contract.UnstakeAllowed(&_Pool.CallOpts)
}

// UnstakeAllowed is a free data retrieval call binding the contract method 0xee018b6d.
//
// Solidity: function unstakeAllowed() view returns(bool)
func (_Pool *PoolCallerSession) UnstakeAllowed() (bool, error) {
	return _Pool.Contract.UnstakeAllowed(&_Pool.CallOpts)
}

// ValidatorManager is a free data retrieval call binding the contract method 0xfe55bde9.
//
// Solidity: function validatorManager() view returns(address)
func (_Pool *PoolCaller) ValidatorManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "validatorManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ValidatorManager is a free data retrieval call binding the contract method 0xfe55bde9.
//
// Solidity: function validatorManager() view returns(address)
func (_Pool *PoolSession) ValidatorManager() (common.Address, error) {
	return _Pool.Contract.ValidatorManager(&_Pool.CallOpts)
}

// ValidatorManager is a free data retrieval call binding the contract method 0xfe55bde9.
//
// Solidity: function validatorManager() view returns(address)
func (_Pool *PoolCallerSession) ValidatorManager() (common.Address, error) {
	return _Pool.Contract.ValidatorManager(&_Pool.CallOpts)
}

// ValidatorRegistry is a free data retrieval call binding the contract method 0xf376ebbb.
//
// Solidity: function validatorRegistry() view returns(address)
func (_Pool *PoolCaller) ValidatorRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "validatorRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ValidatorRegistry is a free data retrieval call binding the contract method 0xf376ebbb.
//
// Solidity: function validatorRegistry() view returns(address)
func (_Pool *PoolSession) ValidatorRegistry() (common.Address, error) {
	return _Pool.Contract.ValidatorRegistry(&_Pool.CallOpts)
}

// ValidatorRegistry is a free data retrieval call binding the contract method 0xf376ebbb.
//
// Solidity: function validatorRegistry() view returns(address)
func (_Pool *PoolCallerSession) ValidatorRegistry() (common.Address, error) {
	return _Pool.Contract.ValidatorRegistry(&_Pool.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint8)
func (_Pool *PoolCaller) Version(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint8)
func (_Pool *PoolSession) Version() (uint8, error) {
	return _Pool.Contract.Version(&_Pool.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint8)
func (_Pool *PoolCallerSession) Version() (uint8, error) {
	return _Pool.Contract.Version(&_Pool.CallOpts)
}

// WithdrawCredentials is a free data retrieval call binding the contract method 0xb62be945.
//
// Solidity: function withdrawCredentials() view returns(bytes[])
func (_Pool *PoolCaller) WithdrawCredentials(opts *bind.CallOpts) ([][]byte, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "withdrawCredentials")

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

// WithdrawCredentials is a free data retrieval call binding the contract method 0xb62be945.
//
// Solidity: function withdrawCredentials() view returns(bytes[])
func (_Pool *PoolSession) WithdrawCredentials() ([][]byte, error) {
	return _Pool.Contract.WithdrawCredentials(&_Pool.CallOpts)
}

// WithdrawCredentials is a free data retrieval call binding the contract method 0xb62be945.
//
// Solidity: function withdrawCredentials() view returns(bytes[])
func (_Pool *PoolCallerSession) WithdrawCredentials() ([][]byte, error) {
	return _Pool.Contract.WithdrawCredentials(&_Pool.CallOpts)
}

// WithdrawalDelayBlocks is a free data retrieval call binding the contract method 0x50f73e7c.
//
// Solidity: function withdrawalDelayBlocks() view returns(uint256)
func (_Pool *PoolCaller) WithdrawalDelayBlocks(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "withdrawalDelayBlocks")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawalDelayBlocks is a free data retrieval call binding the contract method 0x50f73e7c.
//
// Solidity: function withdrawalDelayBlocks() view returns(uint256)
func (_Pool *PoolSession) WithdrawalDelayBlocks() (*big.Int, error) {
	return _Pool.Contract.WithdrawalDelayBlocks(&_Pool.CallOpts)
}

// WithdrawalDelayBlocks is a free data retrieval call binding the contract method 0x50f73e7c.
//
// Solidity: function withdrawalDelayBlocks() view returns(uint256)
func (_Pool *PoolCallerSession) WithdrawalDelayBlocks() (*big.Int, error) {
	return _Pool.Contract.WithdrawalDelayBlocks(&_Pool.CallOpts)
}

// BasePoolInit is a paid mutator transaction binding the contract method 0xdb764e6b.
//
// Solidity: function __BasePool_init(uint256 _withdrawalDelayBlocks, uint256 _apr, uint256 _totalUnderlyingAsset, address _dao, address _poolToken, address _rateManager, address _validatorManager, address _validatorRegistry, address _depositContract) returns()
func (_Pool *PoolTransactor) BasePoolInit(opts *bind.TransactOpts, _withdrawalDelayBlocks *big.Int, _apr *big.Int, _totalUnderlyingAsset *big.Int, _dao common.Address, _poolToken common.Address, _rateManager common.Address, _validatorManager common.Address, _validatorRegistry common.Address, _depositContract common.Address) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "__BasePool_init", _withdrawalDelayBlocks, _apr, _totalUnderlyingAsset, _dao, _poolToken, _rateManager, _validatorManager, _validatorRegistry, _depositContract)
}

// BasePoolInit is a paid mutator transaction binding the contract method 0xdb764e6b.
//
// Solidity: function __BasePool_init(uint256 _withdrawalDelayBlocks, uint256 _apr, uint256 _totalUnderlyingAsset, address _dao, address _poolToken, address _rateManager, address _validatorManager, address _validatorRegistry, address _depositContract) returns()
func (_Pool *PoolSession) BasePoolInit(_withdrawalDelayBlocks *big.Int, _apr *big.Int, _totalUnderlyingAsset *big.Int, _dao common.Address, _poolToken common.Address, _rateManager common.Address, _validatorManager common.Address, _validatorRegistry common.Address, _depositContract common.Address) (*types.Transaction, error) {
	return _Pool.Contract.BasePoolInit(&_Pool.TransactOpts, _withdrawalDelayBlocks, _apr, _totalUnderlyingAsset, _dao, _poolToken, _rateManager, _validatorManager, _validatorRegistry, _depositContract)
}

// BasePoolInit is a paid mutator transaction binding the contract method 0xdb764e6b.
//
// Solidity: function __BasePool_init(uint256 _withdrawalDelayBlocks, uint256 _apr, uint256 _totalUnderlyingAsset, address _dao, address _poolToken, address _rateManager, address _validatorManager, address _validatorRegistry, address _depositContract) returns()
func (_Pool *PoolTransactorSession) BasePoolInit(_withdrawalDelayBlocks *big.Int, _apr *big.Int, _totalUnderlyingAsset *big.Int, _dao common.Address, _poolToken common.Address, _rateManager common.Address, _validatorManager common.Address, _validatorRegistry common.Address, _depositContract common.Address) (*types.Transaction, error) {
	return _Pool.Contract.BasePoolInit(&_Pool.TransactOpts, _withdrawalDelayBlocks, _apr, _totalUnderlyingAsset, _dao, _poolToken, _rateManager, _validatorManager, _validatorRegistry, _depositContract)
}

// AddRestakingPod is a paid mutator transaction binding the contract method 0xebeac515.
//
// Solidity: function addRestakingPod(address _restakingPod) returns()
func (_Pool *PoolTransactor) AddRestakingPod(opts *bind.TransactOpts, _restakingPod common.Address) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "addRestakingPod", _restakingPod)
}

// AddRestakingPod is a paid mutator transaction binding the contract method 0xebeac515.
//
// Solidity: function addRestakingPod(address _restakingPod) returns()
func (_Pool *PoolSession) AddRestakingPod(_restakingPod common.Address) (*types.Transaction, error) {
	return _Pool.Contract.AddRestakingPod(&_Pool.TransactOpts, _restakingPod)
}

// AddRestakingPod is a paid mutator transaction binding the contract method 0xebeac515.
//
// Solidity: function addRestakingPod(address _restakingPod) returns()
func (_Pool *PoolTransactorSession) AddRestakingPod(_restakingPod common.Address) (*types.Transaction, error) {
	return _Pool.Contract.AddRestakingPod(&_Pool.TransactOpts, _restakingPod)
}

// ClaimDelayedWithdrawals is a paid mutator transaction binding the contract method 0x28837c07.
//
// Solidity: function claimDelayedWithdrawals() returns()
func (_Pool *PoolTransactor) ClaimDelayedWithdrawals(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "claimDelayedWithdrawals")
}

// ClaimDelayedWithdrawals is a paid mutator transaction binding the contract method 0x28837c07.
//
// Solidity: function claimDelayedWithdrawals() returns()
func (_Pool *PoolSession) ClaimDelayedWithdrawals() (*types.Transaction, error) {
	return _Pool.Contract.ClaimDelayedWithdrawals(&_Pool.TransactOpts)
}

// ClaimDelayedWithdrawals is a paid mutator transaction binding the contract method 0x28837c07.
//
// Solidity: function claimDelayedWithdrawals() returns()
func (_Pool *PoolTransactorSession) ClaimDelayedWithdrawals() (*types.Transaction, error) {
	return _Pool.Contract.ClaimDelayedWithdrawals(&_Pool.TransactOpts)
}

// ClaimWithdrawals is a paid mutator transaction binding the contract method 0x7139053e.
//
// Solidity: function claimWithdrawals(address _receiver, uint256[] _requestIds) returns()
func (_Pool *PoolTransactor) ClaimWithdrawals(opts *bind.TransactOpts, _receiver common.Address, _requestIds []*big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "claimWithdrawals", _receiver, _requestIds)
}

// ClaimWithdrawals is a paid mutator transaction binding the contract method 0x7139053e.
//
// Solidity: function claimWithdrawals(address _receiver, uint256[] _requestIds) returns()
func (_Pool *PoolSession) ClaimWithdrawals(_receiver common.Address, _requestIds []*big.Int) (*types.Transaction, error) {
	return _Pool.Contract.ClaimWithdrawals(&_Pool.TransactOpts, _receiver, _requestIds)
}

// ClaimWithdrawals is a paid mutator transaction binding the contract method 0x7139053e.
//
// Solidity: function claimWithdrawals(address _receiver, uint256[] _requestIds) returns()
func (_Pool *PoolTransactorSession) ClaimWithdrawals(_receiver common.Address, _requestIds []*big.Int) (*types.Transaction, error) {
	return _Pool.Contract.ClaimWithdrawals(&_Pool.TransactOpts, _receiver, _requestIds)
}

// Initialize is a paid mutator transaction binding the contract method 0x23df75e6.
//
// Solidity: function initialize(uint256 _apr, address _dao, address _elRewardsAddress, address _poolToken, address _rateManager, address _validatorManager, address _validatorRegistry, address _depositContract, address _eigenLayerEigenPodManager, address[] _restakingPods) returns()
func (_Pool *PoolTransactor) Initialize(opts *bind.TransactOpts, _apr *big.Int, _dao common.Address, _elRewardsAddress common.Address, _poolToken common.Address, _rateManager common.Address, _validatorManager common.Address, _validatorRegistry common.Address, _depositContract common.Address, _eigenLayerEigenPodManager common.Address, _restakingPods []common.Address) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "initialize", _apr, _dao, _elRewardsAddress, _poolToken, _rateManager, _validatorManager, _validatorRegistry, _depositContract, _eigenLayerEigenPodManager, _restakingPods)
}

// Initialize is a paid mutator transaction binding the contract method 0x23df75e6.
//
// Solidity: function initialize(uint256 _apr, address _dao, address _elRewardsAddress, address _poolToken, address _rateManager, address _validatorManager, address _validatorRegistry, address _depositContract, address _eigenLayerEigenPodManager, address[] _restakingPods) returns()
func (_Pool *PoolSession) Initialize(_apr *big.Int, _dao common.Address, _elRewardsAddress common.Address, _poolToken common.Address, _rateManager common.Address, _validatorManager common.Address, _validatorRegistry common.Address, _depositContract common.Address, _eigenLayerEigenPodManager common.Address, _restakingPods []common.Address) (*types.Transaction, error) {
	return _Pool.Contract.Initialize(&_Pool.TransactOpts, _apr, _dao, _elRewardsAddress, _poolToken, _rateManager, _validatorManager, _validatorRegistry, _depositContract, _eigenLayerEigenPodManager, _restakingPods)
}

// Initialize is a paid mutator transaction binding the contract method 0x23df75e6.
//
// Solidity: function initialize(uint256 _apr, address _dao, address _elRewardsAddress, address _poolToken, address _rateManager, address _validatorManager, address _validatorRegistry, address _depositContract, address _eigenLayerEigenPodManager, address[] _restakingPods) returns()
func (_Pool *PoolTransactorSession) Initialize(_apr *big.Int, _dao common.Address, _elRewardsAddress common.Address, _poolToken common.Address, _rateManager common.Address, _validatorManager common.Address, _validatorRegistry common.Address, _depositContract common.Address, _eigenLayerEigenPodManager common.Address, _restakingPods []common.Address) (*types.Transaction, error) {
	return _Pool.Contract.Initialize(&_Pool.TransactOpts, _apr, _dao, _elRewardsAddress, _poolToken, _rateManager, _validatorManager, _validatorRegistry, _depositContract, _eigenLayerEigenPodManager, _restakingPods)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Pool *PoolTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Pool *PoolSession) Pause() (*types.Transaction, error) {
	return _Pool.Contract.Pause(&_Pool.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Pool *PoolTransactorSession) Pause() (*types.Transaction, error) {
	return _Pool.Contract.Pause(&_Pool.TransactOpts)
}

// ReceiveRewards is a paid mutator transaction binding the contract method 0xd77868da.
//
// Solidity: function receiveRewards(uint256 ) payable returns()
func (_Pool *PoolTransactor) ReceiveRewards(opts *bind.TransactOpts, arg0 *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "receiveRewards", arg0)
}

// ReceiveRewards is a paid mutator transaction binding the contract method 0xd77868da.
//
// Solidity: function receiveRewards(uint256 ) payable returns()
func (_Pool *PoolSession) ReceiveRewards(arg0 *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.ReceiveRewards(&_Pool.TransactOpts, arg0)
}

// ReceiveRewards is a paid mutator transaction binding the contract method 0xd77868da.
//
// Solidity: function receiveRewards(uint256 ) payable returns()
func (_Pool *PoolTransactorSession) ReceiveRewards(arg0 *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.ReceiveRewards(&_Pool.TransactOpts, arg0)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0xe0b05b61.
//
// Solidity: function registerValidator(bytes32 _depositContractRoot, address _restakingPod, bytes[] _pubkeys, bytes[] _signatures, bytes32[] _depositDataRoots) returns()
func (_Pool *PoolTransactor) RegisterValidator(opts *bind.TransactOpts, _depositContractRoot [32]byte, _restakingPod common.Address, _pubkeys [][]byte, _signatures [][]byte, _depositDataRoots [][32]byte) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "registerValidator", _depositContractRoot, _restakingPod, _pubkeys, _signatures, _depositDataRoots)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0xe0b05b61.
//
// Solidity: function registerValidator(bytes32 _depositContractRoot, address _restakingPod, bytes[] _pubkeys, bytes[] _signatures, bytes32[] _depositDataRoots) returns()
func (_Pool *PoolSession) RegisterValidator(_depositContractRoot [32]byte, _restakingPod common.Address, _pubkeys [][]byte, _signatures [][]byte, _depositDataRoots [][32]byte) (*types.Transaction, error) {
	return _Pool.Contract.RegisterValidator(&_Pool.TransactOpts, _depositContractRoot, _restakingPod, _pubkeys, _signatures, _depositDataRoots)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0xe0b05b61.
//
// Solidity: function registerValidator(bytes32 _depositContractRoot, address _restakingPod, bytes[] _pubkeys, bytes[] _signatures, bytes32[] _depositDataRoots) returns()
func (_Pool *PoolTransactorSession) RegisterValidator(_depositContractRoot [32]byte, _restakingPod common.Address, _pubkeys [][]byte, _signatures [][]byte, _depositDataRoots [][32]byte) (*types.Transaction, error) {
	return _Pool.Contract.RegisterValidator(&_Pool.TransactOpts, _depositContractRoot, _restakingPod, _pubkeys, _signatures, _depositDataRoots)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Pool *PoolTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Pool *PoolSession) RenounceOwnership() (*types.Transaction, error) {
	return _Pool.Contract.RenounceOwnership(&_Pool.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Pool *PoolTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Pool.Contract.RenounceOwnership(&_Pool.TransactOpts)
}

// RequestWithdrawals is a paid mutator transaction binding the contract method 0x2963a287.
//
// Solidity: function requestWithdrawals(uint256 _unstakeAmount) returns()
func (_Pool *PoolTransactor) RequestWithdrawals(opts *bind.TransactOpts, _unstakeAmount *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "requestWithdrawals", _unstakeAmount)
}

// RequestWithdrawals is a paid mutator transaction binding the contract method 0x2963a287.
//
// Solidity: function requestWithdrawals(uint256 _unstakeAmount) returns()
func (_Pool *PoolSession) RequestWithdrawals(_unstakeAmount *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.RequestWithdrawals(&_Pool.TransactOpts, _unstakeAmount)
}

// RequestWithdrawals is a paid mutator transaction binding the contract method 0x2963a287.
//
// Solidity: function requestWithdrawals(uint256 _unstakeAmount) returns()
func (_Pool *PoolTransactorSession) RequestWithdrawals(_unstakeAmount *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.RequestWithdrawals(&_Pool.TransactOpts, _unstakeAmount)
}

// SetRateManager is a paid mutator transaction binding the contract method 0xc49db0cd.
//
// Solidity: function setRateManager(address _rateManager) returns()
func (_Pool *PoolTransactor) SetRateManager(opts *bind.TransactOpts, _rateManager common.Address) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "setRateManager", _rateManager)
}

// SetRateManager is a paid mutator transaction binding the contract method 0xc49db0cd.
//
// Solidity: function setRateManager(address _rateManager) returns()
func (_Pool *PoolSession) SetRateManager(_rateManager common.Address) (*types.Transaction, error) {
	return _Pool.Contract.SetRateManager(&_Pool.TransactOpts, _rateManager)
}

// SetRateManager is a paid mutator transaction binding the contract method 0xc49db0cd.
//
// Solidity: function setRateManager(address _rateManager) returns()
func (_Pool *PoolTransactorSession) SetRateManager(_rateManager common.Address) (*types.Transaction, error) {
	return _Pool.Contract.SetRateManager(&_Pool.TransactOpts, _rateManager)
}

// SetUnstakeAllowed is a paid mutator transaction binding the contract method 0xbba59090.
//
// Solidity: function setUnstakeAllowed(bool _unstakeAllowed) returns()
func (_Pool *PoolTransactor) SetUnstakeAllowed(opts *bind.TransactOpts, _unstakeAllowed bool) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "setUnstakeAllowed", _unstakeAllowed)
}

// SetUnstakeAllowed is a paid mutator transaction binding the contract method 0xbba59090.
//
// Solidity: function setUnstakeAllowed(bool _unstakeAllowed) returns()
func (_Pool *PoolSession) SetUnstakeAllowed(_unstakeAllowed bool) (*types.Transaction, error) {
	return _Pool.Contract.SetUnstakeAllowed(&_Pool.TransactOpts, _unstakeAllowed)
}

// SetUnstakeAllowed is a paid mutator transaction binding the contract method 0xbba59090.
//
// Solidity: function setUnstakeAllowed(bool _unstakeAllowed) returns()
func (_Pool *PoolTransactorSession) SetUnstakeAllowed(_unstakeAllowed bool) (*types.Transaction, error) {
	return _Pool.Contract.SetUnstakeAllowed(&_Pool.TransactOpts, _unstakeAllowed)
}

// SetValidatorManager is a paid mutator transaction binding the contract method 0x45f34e92.
//
// Solidity: function setValidatorManager(address _validatorManager) returns()
func (_Pool *PoolTransactor) SetValidatorManager(opts *bind.TransactOpts, _validatorManager common.Address) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "setValidatorManager", _validatorManager)
}

// SetValidatorManager is a paid mutator transaction binding the contract method 0x45f34e92.
//
// Solidity: function setValidatorManager(address _validatorManager) returns()
func (_Pool *PoolSession) SetValidatorManager(_validatorManager common.Address) (*types.Transaction, error) {
	return _Pool.Contract.SetValidatorManager(&_Pool.TransactOpts, _validatorManager)
}

// SetValidatorManager is a paid mutator transaction binding the contract method 0x45f34e92.
//
// Solidity: function setValidatorManager(address _validatorManager) returns()
func (_Pool *PoolTransactorSession) SetValidatorManager(_validatorManager common.Address) (*types.Transaction, error) {
	return _Pool.Contract.SetValidatorManager(&_Pool.TransactOpts, _validatorManager)
}

// SetValidatorRegistry is a paid mutator transaction binding the contract method 0x49773050.
//
// Solidity: function setValidatorRegistry(address _validatorRegistry) returns()
func (_Pool *PoolTransactor) SetValidatorRegistry(opts *bind.TransactOpts, _validatorRegistry common.Address) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "setValidatorRegistry", _validatorRegistry)
}

// SetValidatorRegistry is a paid mutator transaction binding the contract method 0x49773050.
//
// Solidity: function setValidatorRegistry(address _validatorRegistry) returns()
func (_Pool *PoolSession) SetValidatorRegistry(_validatorRegistry common.Address) (*types.Transaction, error) {
	return _Pool.Contract.SetValidatorRegistry(&_Pool.TransactOpts, _validatorRegistry)
}

// SetValidatorRegistry is a paid mutator transaction binding the contract method 0x49773050.
//
// Solidity: function setValidatorRegistry(address _validatorRegistry) returns()
func (_Pool *PoolTransactorSession) SetValidatorRegistry(_validatorRegistry common.Address) (*types.Transaction, error) {
	return _Pool.Contract.SetValidatorRegistry(&_Pool.TransactOpts, _validatorRegistry)
}

// SetWithdrawalDelayBlocks is a paid mutator transaction binding the contract method 0x4d50f9a4.
//
// Solidity: function setWithdrawalDelayBlocks(uint256 _withdrawalDelayBlocks) returns()
func (_Pool *PoolTransactor) SetWithdrawalDelayBlocks(opts *bind.TransactOpts, _withdrawalDelayBlocks *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "setWithdrawalDelayBlocks", _withdrawalDelayBlocks)
}

// SetWithdrawalDelayBlocks is a paid mutator transaction binding the contract method 0x4d50f9a4.
//
// Solidity: function setWithdrawalDelayBlocks(uint256 _withdrawalDelayBlocks) returns()
func (_Pool *PoolSession) SetWithdrawalDelayBlocks(_withdrawalDelayBlocks *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SetWithdrawalDelayBlocks(&_Pool.TransactOpts, _withdrawalDelayBlocks)
}

// SetWithdrawalDelayBlocks is a paid mutator transaction binding the contract method 0x4d50f9a4.
//
// Solidity: function setWithdrawalDelayBlocks(uint256 _withdrawalDelayBlocks) returns()
func (_Pool *PoolTransactorSession) SetWithdrawalDelayBlocks(_withdrawalDelayBlocks *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SetWithdrawalDelayBlocks(&_Pool.TransactOpts, _withdrawalDelayBlocks)
}

// StakeETH is a paid mutator transaction binding the contract method 0xdceb986d.
//
// Solidity: function stakeETH() payable returns()
func (_Pool *PoolTransactor) StakeETH(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "stakeETH")
}

// StakeETH is a paid mutator transaction binding the contract method 0xdceb986d.
//
// Solidity: function stakeETH() payable returns()
func (_Pool *PoolSession) StakeETH() (*types.Transaction, error) {
	return _Pool.Contract.StakeETH(&_Pool.TransactOpts)
}

// StakeETH is a paid mutator transaction binding the contract method 0xdceb986d.
//
// Solidity: function stakeETH() payable returns()
func (_Pool *PoolTransactorSession) StakeETH() (*types.Transaction, error) {
	return _Pool.Contract.StakeETH(&_Pool.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Pool *PoolTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Pool *PoolSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Pool.Contract.TransferOwnership(&_Pool.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Pool *PoolTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Pool.Contract.TransferOwnership(&_Pool.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Pool *PoolTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Pool *PoolSession) Unpause() (*types.Transaction, error) {
	return _Pool.Contract.Unpause(&_Pool.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Pool *PoolTransactorSession) Unpause() (*types.Transaction, error) {
	return _Pool.Contract.Unpause(&_Pool.TransactOpts)
}

// UnstakeETH is a paid mutator transaction binding the contract method 0x62d53403.
//
// Solidity: function unstakeETH(uint256 _unstakeAmount) returns()
func (_Pool *PoolTransactor) UnstakeETH(opts *bind.TransactOpts, _unstakeAmount *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "unstakeETH", _unstakeAmount)
}

// UnstakeETH is a paid mutator transaction binding the contract method 0x62d53403.
//
// Solidity: function unstakeETH(uint256 _unstakeAmount) returns()
func (_Pool *PoolSession) UnstakeETH(_unstakeAmount *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.UnstakeETH(&_Pool.TransactOpts, _unstakeAmount)
}

// UnstakeETH is a paid mutator transaction binding the contract method 0x62d53403.
//
// Solidity: function unstakeETH(uint256 _unstakeAmount) returns()
func (_Pool *PoolTransactorSession) UnstakeETH(_unstakeAmount *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.UnstakeETH(&_Pool.TransactOpts, _unstakeAmount)
}

// UpdateApr is a paid mutator transaction binding the contract method 0x8552bf90.
//
// Solidity: function updateApr(uint256 _apr) returns()
func (_Pool *PoolTransactor) UpdateApr(opts *bind.TransactOpts, _apr *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "updateApr", _apr)
}

// UpdateApr is a paid mutator transaction binding the contract method 0x8552bf90.
//
// Solidity: function updateApr(uint256 _apr) returns()
func (_Pool *PoolSession) UpdateApr(_apr *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.UpdateApr(&_Pool.TransactOpts, _apr)
}

// UpdateApr is a paid mutator transaction binding the contract method 0x8552bf90.
//
// Solidity: function updateApr(uint256 _apr) returns()
func (_Pool *PoolTransactorSession) UpdateApr(_apr *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.UpdateApr(&_Pool.TransactOpts, _apr)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Pool *PoolTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Pool *PoolSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Pool.Contract.UpgradeTo(&_Pool.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Pool *PoolTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Pool.Contract.UpgradeTo(&_Pool.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Pool *PoolTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Pool *PoolSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Pool.Contract.UpgradeToAndCall(&_Pool.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Pool *PoolTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Pool.Contract.UpgradeToAndCall(&_Pool.TransactOpts, newImplementation, data)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Pool *PoolTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Pool *PoolSession) Receive() (*types.Transaction, error) {
	return _Pool.Contract.Receive(&_Pool.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Pool *PoolTransactorSession) Receive() (*types.Transaction, error) {
	return _Pool.Contract.Receive(&_Pool.TransactOpts)
}

// PoolAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Pool contract.
type PoolAdminChangedIterator struct {
	Event *PoolAdminChanged // Event containing the contract specifics and raw log

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
func (it *PoolAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolAdminChanged)
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
		it.Event = new(PoolAdminChanged)
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
func (it *PoolAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolAdminChanged represents a AdminChanged event raised by the Pool contract.
type PoolAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Pool *PoolFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*PoolAdminChangedIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &PoolAdminChangedIterator{contract: _Pool.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Pool *PoolFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *PoolAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolAdminChanged)
				if err := _Pool.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_Pool *PoolFilterer) ParseAdminChanged(log types.Log) (*PoolAdminChanged, error) {
	event := new(PoolAdminChanged)
	if err := _Pool.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolAprUpdatedIterator is returned from FilterAprUpdated and is used to iterate over the raw logs and unpacked data for AprUpdated events raised by the Pool contract.
type PoolAprUpdatedIterator struct {
	Event *PoolAprUpdated // Event containing the contract specifics and raw log

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
func (it *PoolAprUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolAprUpdated)
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
		it.Event = new(PoolAprUpdated)
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
func (it *PoolAprUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolAprUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolAprUpdated represents a AprUpdated event raised by the Pool contract.
type PoolAprUpdated struct {
	OldApr *big.Int
	Apr    *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAprUpdated is a free log retrieval operation binding the contract event 0x782f84f1274a11befd10700001e41dfdbf825313fb93311d474b857dfd8f1c2b.
//
// Solidity: event AprUpdated(uint256 _oldApr, uint256 _apr)
func (_Pool *PoolFilterer) FilterAprUpdated(opts *bind.FilterOpts) (*PoolAprUpdatedIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "AprUpdated")
	if err != nil {
		return nil, err
	}
	return &PoolAprUpdatedIterator{contract: _Pool.contract, event: "AprUpdated", logs: logs, sub: sub}, nil
}

// WatchAprUpdated is a free log subscription operation binding the contract event 0x782f84f1274a11befd10700001e41dfdbf825313fb93311d474b857dfd8f1c2b.
//
// Solidity: event AprUpdated(uint256 _oldApr, uint256 _apr)
func (_Pool *PoolFilterer) WatchAprUpdated(opts *bind.WatchOpts, sink chan<- *PoolAprUpdated) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "AprUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolAprUpdated)
				if err := _Pool.contract.UnpackLog(event, "AprUpdated", log); err != nil {
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

// ParseAprUpdated is a log parse operation binding the contract event 0x782f84f1274a11befd10700001e41dfdbf825313fb93311d474b857dfd8f1c2b.
//
// Solidity: event AprUpdated(uint256 _oldApr, uint256 _apr)
func (_Pool *PoolFilterer) ParseAprUpdated(log types.Log) (*PoolAprUpdated, error) {
	event := new(PoolAprUpdated)
	if err := _Pool.contract.UnpackLog(event, "AprUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolAssetsUpdatedIterator is returned from FilterAssetsUpdated and is used to iterate over the raw logs and unpacked data for AssetsUpdated events raised by the Pool contract.
type PoolAssetsUpdatedIterator struct {
	Event *PoolAssetsUpdated // Event containing the contract specifics and raw log

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
func (it *PoolAssetsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolAssetsUpdated)
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
		it.Event = new(PoolAssetsUpdated)
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
func (it *PoolAssetsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolAssetsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolAssetsUpdated represents a AssetsUpdated event raised by the Pool contract.
type PoolAssetsUpdated struct {
	TotalUnderlyingAsset *big.Int
	EstimatedRewards     *big.Int
	BlockNumber          *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterAssetsUpdated is a free log retrieval operation binding the contract event 0x35a901c4413e585f9121eb5cf07e67760bd4ac498dd031249e5cd2cd225f74e4.
//
// Solidity: event AssetsUpdated(uint256 _totalUnderlyingAsset, uint256 _estimatedRewards, uint256 _blockNumber)
func (_Pool *PoolFilterer) FilterAssetsUpdated(opts *bind.FilterOpts) (*PoolAssetsUpdatedIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "AssetsUpdated")
	if err != nil {
		return nil, err
	}
	return &PoolAssetsUpdatedIterator{contract: _Pool.contract, event: "AssetsUpdated", logs: logs, sub: sub}, nil
}

// WatchAssetsUpdated is a free log subscription operation binding the contract event 0x35a901c4413e585f9121eb5cf07e67760bd4ac498dd031249e5cd2cd225f74e4.
//
// Solidity: event AssetsUpdated(uint256 _totalUnderlyingAsset, uint256 _estimatedRewards, uint256 _blockNumber)
func (_Pool *PoolFilterer) WatchAssetsUpdated(opts *bind.WatchOpts, sink chan<- *PoolAssetsUpdated) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "AssetsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolAssetsUpdated)
				if err := _Pool.contract.UnpackLog(event, "AssetsUpdated", log); err != nil {
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

// ParseAssetsUpdated is a log parse operation binding the contract event 0x35a901c4413e585f9121eb5cf07e67760bd4ac498dd031249e5cd2cd225f74e4.
//
// Solidity: event AssetsUpdated(uint256 _totalUnderlyingAsset, uint256 _estimatedRewards, uint256 _blockNumber)
func (_Pool *PoolFilterer) ParseAssetsUpdated(log types.Log) (*PoolAssetsUpdated, error) {
	event := new(PoolAssetsUpdated)
	if err := _Pool.contract.UnpackLog(event, "AssetsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the Pool contract.
type PoolBeaconUpgradedIterator struct {
	Event *PoolBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *PoolBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolBeaconUpgraded)
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
		it.Event = new(PoolBeaconUpgraded)
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
func (it *PoolBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolBeaconUpgraded represents a BeaconUpgraded event raised by the Pool contract.
type PoolBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Pool *PoolFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*PoolBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &PoolBeaconUpgradedIterator{contract: _Pool.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Pool *PoolFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *PoolBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolBeaconUpgraded)
				if err := _Pool.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_Pool *PoolFilterer) ParseBeaconUpgraded(log types.Log) (*PoolBeaconUpgraded, error) {
	event := new(PoolBeaconUpgraded)
	if err := _Pool.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolDaoChangedIterator is returned from FilterDaoChanged and is used to iterate over the raw logs and unpacked data for DaoChanged events raised by the Pool contract.
type PoolDaoChangedIterator struct {
	Event *PoolDaoChanged // Event containing the contract specifics and raw log

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
func (it *PoolDaoChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolDaoChanged)
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
		it.Event = new(PoolDaoChanged)
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
func (it *PoolDaoChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolDaoChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolDaoChanged represents a DaoChanged event raised by the Pool contract.
type PoolDaoChanged struct {
	OldDao common.Address
	Dao    common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDaoChanged is a free log retrieval operation binding the contract event 0xfcde6c827a52b0870bc44ed9b10212272e18c9ea1725b772e9b493750afd8da4.
//
// Solidity: event DaoChanged(address _oldDao, address _dao)
func (_Pool *PoolFilterer) FilterDaoChanged(opts *bind.FilterOpts) (*PoolDaoChangedIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "DaoChanged")
	if err != nil {
		return nil, err
	}
	return &PoolDaoChangedIterator{contract: _Pool.contract, event: "DaoChanged", logs: logs, sub: sub}, nil
}

// WatchDaoChanged is a free log subscription operation binding the contract event 0xfcde6c827a52b0870bc44ed9b10212272e18c9ea1725b772e9b493750afd8da4.
//
// Solidity: event DaoChanged(address _oldDao, address _dao)
func (_Pool *PoolFilterer) WatchDaoChanged(opts *bind.WatchOpts, sink chan<- *PoolDaoChanged) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "DaoChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolDaoChanged)
				if err := _Pool.contract.UnpackLog(event, "DaoChanged", log); err != nil {
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
func (_Pool *PoolFilterer) ParseDaoChanged(log types.Log) (*PoolDaoChanged, error) {
	event := new(PoolDaoChanged)
	if err := _Pool.contract.UnpackLog(event, "DaoChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolEthStakeIterator is returned from FilterEthStake and is used to iterate over the raw logs and unpacked data for EthStake events raised by the Pool contract.
type PoolEthStakeIterator struct {
	Event *PoolEthStake // Event containing the contract specifics and raw log

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
func (it *PoolEthStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolEthStake)
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
		it.Event = new(PoolEthStake)
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
func (it *PoolEthStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolEthStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolEthStake represents a EthStake event raised by the Pool contract.
type PoolEthStake struct {
	Staker      common.Address
	StakeAmount *big.Int
	MintAmount  *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterEthStake is a free log retrieval operation binding the contract event 0x838d17987e57e587c458220b9b38723c41fbc3f397550b506712960a73ef19f9.
//
// Solidity: event EthStake(address _staker, uint256 _stakeAmount, uint256 _mintAmount)
func (_Pool *PoolFilterer) FilterEthStake(opts *bind.FilterOpts) (*PoolEthStakeIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "EthStake")
	if err != nil {
		return nil, err
	}
	return &PoolEthStakeIterator{contract: _Pool.contract, event: "EthStake", logs: logs, sub: sub}, nil
}

// WatchEthStake is a free log subscription operation binding the contract event 0x838d17987e57e587c458220b9b38723c41fbc3f397550b506712960a73ef19f9.
//
// Solidity: event EthStake(address _staker, uint256 _stakeAmount, uint256 _mintAmount)
func (_Pool *PoolFilterer) WatchEthStake(opts *bind.WatchOpts, sink chan<- *PoolEthStake) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "EthStake")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolEthStake)
				if err := _Pool.contract.UnpackLog(event, "EthStake", log); err != nil {
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

// ParseEthStake is a log parse operation binding the contract event 0x838d17987e57e587c458220b9b38723c41fbc3f397550b506712960a73ef19f9.
//
// Solidity: event EthStake(address _staker, uint256 _stakeAmount, uint256 _mintAmount)
func (_Pool *PoolFilterer) ParseEthStake(log types.Log) (*PoolEthStake, error) {
	event := new(PoolEthStake)
	if err := _Pool.contract.UnpackLog(event, "EthStake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolEthUnstakeIterator is returned from FilterEthUnstake and is used to iterate over the raw logs and unpacked data for EthUnstake events raised by the Pool contract.
type PoolEthUnstakeIterator struct {
	Event *PoolEthUnstake // Event containing the contract specifics and raw log

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
func (it *PoolEthUnstakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolEthUnstake)
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
		it.Event = new(PoolEthUnstake)
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
func (it *PoolEthUnstakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolEthUnstakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolEthUnstake represents a EthUnstake event raised by the Pool contract.
type PoolEthUnstake struct {
	Sender        common.Address
	UnstakeAmount *big.Int
	EthAmount     *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterEthUnstake is a free log retrieval operation binding the contract event 0xdef4a00a06705bb4b80fd0c912337f4f4e01fb81d39a572514be68d328599abf.
//
// Solidity: event EthUnstake(address _sender, uint256 _unstakeAmount, uint256 _ethAmount)
func (_Pool *PoolFilterer) FilterEthUnstake(opts *bind.FilterOpts) (*PoolEthUnstakeIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "EthUnstake")
	if err != nil {
		return nil, err
	}
	return &PoolEthUnstakeIterator{contract: _Pool.contract, event: "EthUnstake", logs: logs, sub: sub}, nil
}

// WatchEthUnstake is a free log subscription operation binding the contract event 0xdef4a00a06705bb4b80fd0c912337f4f4e01fb81d39a572514be68d328599abf.
//
// Solidity: event EthUnstake(address _sender, uint256 _unstakeAmount, uint256 _ethAmount)
func (_Pool *PoolFilterer) WatchEthUnstake(opts *bind.WatchOpts, sink chan<- *PoolEthUnstake) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "EthUnstake")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolEthUnstake)
				if err := _Pool.contract.UnpackLog(event, "EthUnstake", log); err != nil {
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

// ParseEthUnstake is a log parse operation binding the contract event 0xdef4a00a06705bb4b80fd0c912337f4f4e01fb81d39a572514be68d328599abf.
//
// Solidity: event EthUnstake(address _sender, uint256 _unstakeAmount, uint256 _ethAmount)
func (_Pool *PoolFilterer) ParseEthUnstake(log types.Log) (*PoolEthUnstake, error) {
	event := new(PoolEthUnstake)
	if err := _Pool.contract.UnpackLog(event, "EthUnstake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Pool contract.
type PoolInitializedIterator struct {
	Event *PoolInitialized // Event containing the contract specifics and raw log

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
func (it *PoolInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolInitialized)
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
		it.Event = new(PoolInitialized)
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
func (it *PoolInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolInitialized represents a Initialized event raised by the Pool contract.
type PoolInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Pool *PoolFilterer) FilterInitialized(opts *bind.FilterOpts) (*PoolInitializedIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PoolInitializedIterator{contract: _Pool.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Pool *PoolFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PoolInitialized) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolInitialized)
				if err := _Pool.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Pool *PoolFilterer) ParseInitialized(log types.Log) (*PoolInitialized, error) {
	event := new(PoolInitialized)
	if err := _Pool.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Pool contract.
type PoolOwnershipTransferredIterator struct {
	Event *PoolOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PoolOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolOwnershipTransferred)
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
		it.Event = new(PoolOwnershipTransferred)
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
func (it *PoolOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolOwnershipTransferred represents a OwnershipTransferred event raised by the Pool contract.
type PoolOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Pool *PoolFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PoolOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PoolOwnershipTransferredIterator{contract: _Pool.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Pool *PoolFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PoolOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolOwnershipTransferred)
				if err := _Pool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Pool *PoolFilterer) ParseOwnershipTransferred(log types.Log) (*PoolOwnershipTransferred, error) {
	event := new(PoolOwnershipTransferred)
	if err := _Pool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Pool contract.
type PoolPausedIterator struct {
	Event *PoolPaused // Event containing the contract specifics and raw log

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
func (it *PoolPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolPaused)
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
		it.Event = new(PoolPaused)
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
func (it *PoolPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolPaused represents a Paused event raised by the Pool contract.
type PoolPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Pool *PoolFilterer) FilterPaused(opts *bind.FilterOpts) (*PoolPausedIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &PoolPausedIterator{contract: _Pool.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Pool *PoolFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *PoolPaused) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolPaused)
				if err := _Pool.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Pool *PoolFilterer) ParsePaused(log types.Log) (*PoolPaused, error) {
	event := new(PoolPaused)
	if err := _Pool.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolRateManagerChangedIterator is returned from FilterRateManagerChanged and is used to iterate over the raw logs and unpacked data for RateManagerChanged events raised by the Pool contract.
type PoolRateManagerChangedIterator struct {
	Event *PoolRateManagerChanged // Event containing the contract specifics and raw log

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
func (it *PoolRateManagerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolRateManagerChanged)
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
		it.Event = new(PoolRateManagerChanged)
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
func (it *PoolRateManagerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolRateManagerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolRateManagerChanged represents a RateManagerChanged event raised by the Pool contract.
type PoolRateManagerChanged struct {
	OldAprManager common.Address
	AprManager    common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRateManagerChanged is a free log retrieval operation binding the contract event 0x6b8a1f2fd09d355f8419738a3593f646cbd0e7be553ffcce23694ce968cc6425.
//
// Solidity: event RateManagerChanged(address _oldAprManager, address _aprManager)
func (_Pool *PoolFilterer) FilterRateManagerChanged(opts *bind.FilterOpts) (*PoolRateManagerChangedIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "RateManagerChanged")
	if err != nil {
		return nil, err
	}
	return &PoolRateManagerChangedIterator{contract: _Pool.contract, event: "RateManagerChanged", logs: logs, sub: sub}, nil
}

// WatchRateManagerChanged is a free log subscription operation binding the contract event 0x6b8a1f2fd09d355f8419738a3593f646cbd0e7be553ffcce23694ce968cc6425.
//
// Solidity: event RateManagerChanged(address _oldAprManager, address _aprManager)
func (_Pool *PoolFilterer) WatchRateManagerChanged(opts *bind.WatchOpts, sink chan<- *PoolRateManagerChanged) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "RateManagerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolRateManagerChanged)
				if err := _Pool.contract.UnpackLog(event, "RateManagerChanged", log); err != nil {
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

// ParseRateManagerChanged is a log parse operation binding the contract event 0x6b8a1f2fd09d355f8419738a3593f646cbd0e7be553ffcce23694ce968cc6425.
//
// Solidity: event RateManagerChanged(address _oldAprManager, address _aprManager)
func (_Pool *PoolFilterer) ParseRateManagerChanged(log types.Log) (*PoolRateManagerChanged, error) {
	event := new(PoolRateManagerChanged)
	if err := _Pool.contract.UnpackLog(event, "RateManagerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolReceivedIterator is returned from FilterReceived and is used to iterate over the raw logs and unpacked data for Received events raised by the Pool contract.
type PoolReceivedIterator struct {
	Event *PoolReceived // Event containing the contract specifics and raw log

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
func (it *PoolReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolReceived)
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
		it.Event = new(PoolReceived)
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
func (it *PoolReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolReceived represents a Received event raised by the Pool contract.
type PoolReceived struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReceived is a free log retrieval operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address _sender, uint256 _amount)
func (_Pool *PoolFilterer) FilterReceived(opts *bind.FilterOpts) (*PoolReceivedIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "Received")
	if err != nil {
		return nil, err
	}
	return &PoolReceivedIterator{contract: _Pool.contract, event: "Received", logs: logs, sub: sub}, nil
}

// WatchReceived is a free log subscription operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address _sender, uint256 _amount)
func (_Pool *PoolFilterer) WatchReceived(opts *bind.WatchOpts, sink chan<- *PoolReceived) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "Received")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolReceived)
				if err := _Pool.contract.UnpackLog(event, "Received", log); err != nil {
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

// ParseReceived is a log parse operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address _sender, uint256 _amount)
func (_Pool *PoolFilterer) ParseReceived(log types.Log) (*PoolReceived, error) {
	event := new(PoolReceived)
	if err := _Pool.contract.UnpackLog(event, "Received", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolRestakingPodAddedIterator is returned from FilterRestakingPodAdded and is used to iterate over the raw logs and unpacked data for RestakingPodAdded events raised by the Pool contract.
type PoolRestakingPodAddedIterator struct {
	Event *PoolRestakingPodAdded // Event containing the contract specifics and raw log

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
func (it *PoolRestakingPodAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolRestakingPodAdded)
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
		it.Event = new(PoolRestakingPodAdded)
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
func (it *PoolRestakingPodAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolRestakingPodAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolRestakingPodAdded represents a RestakingPodAdded event raised by the Pool contract.
type PoolRestakingPodAdded struct {
	RestakingPod common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRestakingPodAdded is a free log retrieval operation binding the contract event 0x84cf405f0a1827114df681d11bd5006d56d6414d63c57e4cf6569fabd9a8e7ff.
//
// Solidity: event RestakingPodAdded(address _restakingPod)
func (_Pool *PoolFilterer) FilterRestakingPodAdded(opts *bind.FilterOpts) (*PoolRestakingPodAddedIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "RestakingPodAdded")
	if err != nil {
		return nil, err
	}
	return &PoolRestakingPodAddedIterator{contract: _Pool.contract, event: "RestakingPodAdded", logs: logs, sub: sub}, nil
}

// WatchRestakingPodAdded is a free log subscription operation binding the contract event 0x84cf405f0a1827114df681d11bd5006d56d6414d63c57e4cf6569fabd9a8e7ff.
//
// Solidity: event RestakingPodAdded(address _restakingPod)
func (_Pool *PoolFilterer) WatchRestakingPodAdded(opts *bind.WatchOpts, sink chan<- *PoolRestakingPodAdded) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "RestakingPodAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolRestakingPodAdded)
				if err := _Pool.contract.UnpackLog(event, "RestakingPodAdded", log); err != nil {
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

// ParseRestakingPodAdded is a log parse operation binding the contract event 0x84cf405f0a1827114df681d11bd5006d56d6414d63c57e4cf6569fabd9a8e7ff.
//
// Solidity: event RestakingPodAdded(address _restakingPod)
func (_Pool *PoolFilterer) ParseRestakingPodAdded(log types.Log) (*PoolRestakingPodAdded, error) {
	event := new(PoolRestakingPodAdded)
	if err := _Pool.contract.UnpackLog(event, "RestakingPodAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Pool contract.
type PoolUnpausedIterator struct {
	Event *PoolUnpaused // Event containing the contract specifics and raw log

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
func (it *PoolUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolUnpaused)
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
		it.Event = new(PoolUnpaused)
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
func (it *PoolUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolUnpaused represents a Unpaused event raised by the Pool contract.
type PoolUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Pool *PoolFilterer) FilterUnpaused(opts *bind.FilterOpts) (*PoolUnpausedIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &PoolUnpausedIterator{contract: _Pool.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Pool *PoolFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *PoolUnpaused) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolUnpaused)
				if err := _Pool.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_Pool *PoolFilterer) ParseUnpaused(log types.Log) (*PoolUnpaused, error) {
	event := new(PoolUnpaused)
	if err := _Pool.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolUnstakeAllowedUpdatedIterator is returned from FilterUnstakeAllowedUpdated and is used to iterate over the raw logs and unpacked data for UnstakeAllowedUpdated events raised by the Pool contract.
type PoolUnstakeAllowedUpdatedIterator struct {
	Event *PoolUnstakeAllowedUpdated // Event containing the contract specifics and raw log

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
func (it *PoolUnstakeAllowedUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolUnstakeAllowedUpdated)
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
		it.Event = new(PoolUnstakeAllowedUpdated)
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
func (it *PoolUnstakeAllowedUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolUnstakeAllowedUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolUnstakeAllowedUpdated represents a UnstakeAllowedUpdated event raised by the Pool contract.
type PoolUnstakeAllowedUpdated struct {
	OldUnstakeAllowed bool
	UnstakeAllowed    bool
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterUnstakeAllowedUpdated is a free log retrieval operation binding the contract event 0xf0700b7cce37d0ddc289445e07408f3709033f9785391e5e05b908cb9705748b.
//
// Solidity: event UnstakeAllowedUpdated(bool _oldUnstakeAllowed, bool _unstakeAllowed)
func (_Pool *PoolFilterer) FilterUnstakeAllowedUpdated(opts *bind.FilterOpts) (*PoolUnstakeAllowedUpdatedIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "UnstakeAllowedUpdated")
	if err != nil {
		return nil, err
	}
	return &PoolUnstakeAllowedUpdatedIterator{contract: _Pool.contract, event: "UnstakeAllowedUpdated", logs: logs, sub: sub}, nil
}

// WatchUnstakeAllowedUpdated is a free log subscription operation binding the contract event 0xf0700b7cce37d0ddc289445e07408f3709033f9785391e5e05b908cb9705748b.
//
// Solidity: event UnstakeAllowedUpdated(bool _oldUnstakeAllowed, bool _unstakeAllowed)
func (_Pool *PoolFilterer) WatchUnstakeAllowedUpdated(opts *bind.WatchOpts, sink chan<- *PoolUnstakeAllowedUpdated) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "UnstakeAllowedUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolUnstakeAllowedUpdated)
				if err := _Pool.contract.UnpackLog(event, "UnstakeAllowedUpdated", log); err != nil {
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

// ParseUnstakeAllowedUpdated is a log parse operation binding the contract event 0xf0700b7cce37d0ddc289445e07408f3709033f9785391e5e05b908cb9705748b.
//
// Solidity: event UnstakeAllowedUpdated(bool _oldUnstakeAllowed, bool _unstakeAllowed)
func (_Pool *PoolFilterer) ParseUnstakeAllowedUpdated(log types.Log) (*PoolUnstakeAllowedUpdated, error) {
	event := new(PoolUnstakeAllowedUpdated)
	if err := _Pool.contract.UnpackLog(event, "UnstakeAllowedUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Pool contract.
type PoolUpgradedIterator struct {
	Event *PoolUpgraded // Event containing the contract specifics and raw log

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
func (it *PoolUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolUpgraded)
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
		it.Event = new(PoolUpgraded)
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
func (it *PoolUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolUpgraded represents a Upgraded event raised by the Pool contract.
type PoolUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Pool *PoolFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*PoolUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &PoolUpgradedIterator{contract: _Pool.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Pool *PoolFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *PoolUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolUpgraded)
				if err := _Pool.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_Pool *PoolFilterer) ParseUpgraded(log types.Log) (*PoolUpgraded, error) {
	event := new(PoolUpgraded)
	if err := _Pool.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolValidatorManagerChangedIterator is returned from FilterValidatorManagerChanged and is used to iterate over the raw logs and unpacked data for ValidatorManagerChanged events raised by the Pool contract.
type PoolValidatorManagerChangedIterator struct {
	Event *PoolValidatorManagerChanged // Event containing the contract specifics and raw log

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
func (it *PoolValidatorManagerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolValidatorManagerChanged)
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
		it.Event = new(PoolValidatorManagerChanged)
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
func (it *PoolValidatorManagerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolValidatorManagerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolValidatorManagerChanged represents a ValidatorManagerChanged event raised by the Pool contract.
type PoolValidatorManagerChanged struct {
	OldValidatorManager common.Address
	ValidatorManager    common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterValidatorManagerChanged is a free log retrieval operation binding the contract event 0xfcf16285f1e14b0b8a544860d3664317dd81073a42e811e33afe75a22886443e.
//
// Solidity: event ValidatorManagerChanged(address _oldValidatorManager, address _validatorManager)
func (_Pool *PoolFilterer) FilterValidatorManagerChanged(opts *bind.FilterOpts) (*PoolValidatorManagerChangedIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "ValidatorManagerChanged")
	if err != nil {
		return nil, err
	}
	return &PoolValidatorManagerChangedIterator{contract: _Pool.contract, event: "ValidatorManagerChanged", logs: logs, sub: sub}, nil
}

// WatchValidatorManagerChanged is a free log subscription operation binding the contract event 0xfcf16285f1e14b0b8a544860d3664317dd81073a42e811e33afe75a22886443e.
//
// Solidity: event ValidatorManagerChanged(address _oldValidatorManager, address _validatorManager)
func (_Pool *PoolFilterer) WatchValidatorManagerChanged(opts *bind.WatchOpts, sink chan<- *PoolValidatorManagerChanged) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "ValidatorManagerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolValidatorManagerChanged)
				if err := _Pool.contract.UnpackLog(event, "ValidatorManagerChanged", log); err != nil {
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

// ParseValidatorManagerChanged is a log parse operation binding the contract event 0xfcf16285f1e14b0b8a544860d3664317dd81073a42e811e33afe75a22886443e.
//
// Solidity: event ValidatorManagerChanged(address _oldValidatorManager, address _validatorManager)
func (_Pool *PoolFilterer) ParseValidatorManagerChanged(log types.Log) (*PoolValidatorManagerChanged, error) {
	event := new(PoolValidatorManagerChanged)
	if err := _Pool.contract.UnpackLog(event, "ValidatorManagerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolValidatorRegistrationIterator is returned from FilterValidatorRegistration and is used to iterate over the raw logs and unpacked data for ValidatorRegistration events raised by the Pool contract.
type PoolValidatorRegistrationIterator struct {
	Event *PoolValidatorRegistration // Event containing the contract specifics and raw log

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
func (it *PoolValidatorRegistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolValidatorRegistration)
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
		it.Event = new(PoolValidatorRegistration)
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
func (it *PoolValidatorRegistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolValidatorRegistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolValidatorRegistration represents a ValidatorRegistration event raised by the Pool contract.
type PoolValidatorRegistration struct {
	Pubkeys [][]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterValidatorRegistration is a free log retrieval operation binding the contract event 0xe585eadb0042252d35431ecfed1027caf5672811fdf1db08cb49e32fee170505.
//
// Solidity: event ValidatorRegistration(bytes[] _pubkeys)
func (_Pool *PoolFilterer) FilterValidatorRegistration(opts *bind.FilterOpts) (*PoolValidatorRegistrationIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "ValidatorRegistration")
	if err != nil {
		return nil, err
	}
	return &PoolValidatorRegistrationIterator{contract: _Pool.contract, event: "ValidatorRegistration", logs: logs, sub: sub}, nil
}

// WatchValidatorRegistration is a free log subscription operation binding the contract event 0xe585eadb0042252d35431ecfed1027caf5672811fdf1db08cb49e32fee170505.
//
// Solidity: event ValidatorRegistration(bytes[] _pubkeys)
func (_Pool *PoolFilterer) WatchValidatorRegistration(opts *bind.WatchOpts, sink chan<- *PoolValidatorRegistration) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "ValidatorRegistration")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolValidatorRegistration)
				if err := _Pool.contract.UnpackLog(event, "ValidatorRegistration", log); err != nil {
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

// ParseValidatorRegistration is a log parse operation binding the contract event 0xe585eadb0042252d35431ecfed1027caf5672811fdf1db08cb49e32fee170505.
//
// Solidity: event ValidatorRegistration(bytes[] _pubkeys)
func (_Pool *PoolFilterer) ParseValidatorRegistration(log types.Log) (*PoolValidatorRegistration, error) {
	event := new(PoolValidatorRegistration)
	if err := _Pool.contract.UnpackLog(event, "ValidatorRegistration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolValidatorRegistryChangedIterator is returned from FilterValidatorRegistryChanged and is used to iterate over the raw logs and unpacked data for ValidatorRegistryChanged events raised by the Pool contract.
type PoolValidatorRegistryChangedIterator struct {
	Event *PoolValidatorRegistryChanged // Event containing the contract specifics and raw log

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
func (it *PoolValidatorRegistryChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolValidatorRegistryChanged)
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
		it.Event = new(PoolValidatorRegistryChanged)
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
func (it *PoolValidatorRegistryChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolValidatorRegistryChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolValidatorRegistryChanged represents a ValidatorRegistryChanged event raised by the Pool contract.
type PoolValidatorRegistryChanged struct {
	OldValidatorManager common.Address
	ValidatorManager    common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterValidatorRegistryChanged is a free log retrieval operation binding the contract event 0x5f98eb9c39a016a522ab1ca3601f349c89b557a9b6471ce04afa8770d5589062.
//
// Solidity: event ValidatorRegistryChanged(address _oldValidatorManager, address _validatorManager)
func (_Pool *PoolFilterer) FilterValidatorRegistryChanged(opts *bind.FilterOpts) (*PoolValidatorRegistryChangedIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "ValidatorRegistryChanged")
	if err != nil {
		return nil, err
	}
	return &PoolValidatorRegistryChangedIterator{contract: _Pool.contract, event: "ValidatorRegistryChanged", logs: logs, sub: sub}, nil
}

// WatchValidatorRegistryChanged is a free log subscription operation binding the contract event 0x5f98eb9c39a016a522ab1ca3601f349c89b557a9b6471ce04afa8770d5589062.
//
// Solidity: event ValidatorRegistryChanged(address _oldValidatorManager, address _validatorManager)
func (_Pool *PoolFilterer) WatchValidatorRegistryChanged(opts *bind.WatchOpts, sink chan<- *PoolValidatorRegistryChanged) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "ValidatorRegistryChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolValidatorRegistryChanged)
				if err := _Pool.contract.UnpackLog(event, "ValidatorRegistryChanged", log); err != nil {
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

// ParseValidatorRegistryChanged is a log parse operation binding the contract event 0x5f98eb9c39a016a522ab1ca3601f349c89b557a9b6471ce04afa8770d5589062.
//
// Solidity: event ValidatorRegistryChanged(address _oldValidatorManager, address _validatorManager)
func (_Pool *PoolFilterer) ParseValidatorRegistryChanged(log types.Log) (*PoolValidatorRegistryChanged, error) {
	event := new(PoolValidatorRegistryChanged)
	if err := _Pool.contract.UnpackLog(event, "ValidatorRegistryChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolWithdrawalDelayChangedIterator is returned from FilterWithdrawalDelayChanged and is used to iterate over the raw logs and unpacked data for WithdrawalDelayChanged events raised by the Pool contract.
type PoolWithdrawalDelayChangedIterator struct {
	Event *PoolWithdrawalDelayChanged // Event containing the contract specifics and raw log

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
func (it *PoolWithdrawalDelayChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolWithdrawalDelayChanged)
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
		it.Event = new(PoolWithdrawalDelayChanged)
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
func (it *PoolWithdrawalDelayChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolWithdrawalDelayChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolWithdrawalDelayChanged represents a WithdrawalDelayChanged event raised by the Pool contract.
type PoolWithdrawalDelayChanged struct {
	OldWithdrawalDelayBlocks *big.Int
	WithdrawalDelayBlocks    *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalDelayChanged is a free log retrieval operation binding the contract event 0xab3f1d5eaee409b7067167f77f1fa3f8a863366d6fb2b88559cd4f9b8e03e182.
//
// Solidity: event WithdrawalDelayChanged(uint256 _oldWithdrawalDelayBlocks, uint256 _withdrawalDelayBlocks)
func (_Pool *PoolFilterer) FilterWithdrawalDelayChanged(opts *bind.FilterOpts) (*PoolWithdrawalDelayChangedIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "WithdrawalDelayChanged")
	if err != nil {
		return nil, err
	}
	return &PoolWithdrawalDelayChangedIterator{contract: _Pool.contract, event: "WithdrawalDelayChanged", logs: logs, sub: sub}, nil
}

// WatchWithdrawalDelayChanged is a free log subscription operation binding the contract event 0xab3f1d5eaee409b7067167f77f1fa3f8a863366d6fb2b88559cd4f9b8e03e182.
//
// Solidity: event WithdrawalDelayChanged(uint256 _oldWithdrawalDelayBlocks, uint256 _withdrawalDelayBlocks)
func (_Pool *PoolFilterer) WatchWithdrawalDelayChanged(opts *bind.WatchOpts, sink chan<- *PoolWithdrawalDelayChanged) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "WithdrawalDelayChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolWithdrawalDelayChanged)
				if err := _Pool.contract.UnpackLog(event, "WithdrawalDelayChanged", log); err != nil {
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

// ParseWithdrawalDelayChanged is a log parse operation binding the contract event 0xab3f1d5eaee409b7067167f77f1fa3f8a863366d6fb2b88559cd4f9b8e03e182.
//
// Solidity: event WithdrawalDelayChanged(uint256 _oldWithdrawalDelayBlocks, uint256 _withdrawalDelayBlocks)
func (_Pool *PoolFilterer) ParseWithdrawalDelayChanged(log types.Log) (*PoolWithdrawalDelayChanged, error) {
	event := new(PoolWithdrawalDelayChanged)
	if err := _Pool.contract.UnpackLog(event, "WithdrawalDelayChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolWithdrawalsClaimedIterator is returned from FilterWithdrawalsClaimed and is used to iterate over the raw logs and unpacked data for WithdrawalsClaimed events raised by the Pool contract.
type PoolWithdrawalsClaimedIterator struct {
	Event *PoolWithdrawalsClaimed // Event containing the contract specifics and raw log

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
func (it *PoolWithdrawalsClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolWithdrawalsClaimed)
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
		it.Event = new(PoolWithdrawalsClaimed)
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
func (it *PoolWithdrawalsClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolWithdrawalsClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolWithdrawalsClaimed represents a WithdrawalsClaimed event raised by the Pool contract.
type PoolWithdrawalsClaimed struct {
	Receiver    common.Address
	RequestId   *big.Int
	ClaimAmount *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalsClaimed is a free log retrieval operation binding the contract event 0xf39bbe3e7fb2d887fe7e6e23dea5a53c1e720410dfb13e87567b873845136cf4.
//
// Solidity: event WithdrawalsClaimed(address _receiver, uint256 _requestId, uint256 _claimAmount)
func (_Pool *PoolFilterer) FilterWithdrawalsClaimed(opts *bind.FilterOpts) (*PoolWithdrawalsClaimedIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "WithdrawalsClaimed")
	if err != nil {
		return nil, err
	}
	return &PoolWithdrawalsClaimedIterator{contract: _Pool.contract, event: "WithdrawalsClaimed", logs: logs, sub: sub}, nil
}

// WatchWithdrawalsClaimed is a free log subscription operation binding the contract event 0xf39bbe3e7fb2d887fe7e6e23dea5a53c1e720410dfb13e87567b873845136cf4.
//
// Solidity: event WithdrawalsClaimed(address _receiver, uint256 _requestId, uint256 _claimAmount)
func (_Pool *PoolFilterer) WatchWithdrawalsClaimed(opts *bind.WatchOpts, sink chan<- *PoolWithdrawalsClaimed) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "WithdrawalsClaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolWithdrawalsClaimed)
				if err := _Pool.contract.UnpackLog(event, "WithdrawalsClaimed", log); err != nil {
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

// ParseWithdrawalsClaimed is a log parse operation binding the contract event 0xf39bbe3e7fb2d887fe7e6e23dea5a53c1e720410dfb13e87567b873845136cf4.
//
// Solidity: event WithdrawalsClaimed(address _receiver, uint256 _requestId, uint256 _claimAmount)
func (_Pool *PoolFilterer) ParseWithdrawalsClaimed(log types.Log) (*PoolWithdrawalsClaimed, error) {
	event := new(PoolWithdrawalsClaimed)
	if err := _Pool.contract.UnpackLog(event, "WithdrawalsClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolWithdrawalsRequestIterator is returned from FilterWithdrawalsRequest and is used to iterate over the raw logs and unpacked data for WithdrawalsRequest events raised by the Pool contract.
type PoolWithdrawalsRequestIterator struct {
	Event *PoolWithdrawalsRequest // Event containing the contract specifics and raw log

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
func (it *PoolWithdrawalsRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolWithdrawalsRequest)
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
		it.Event = new(PoolWithdrawalsRequest)
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
func (it *PoolWithdrawalsRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolWithdrawalsRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolWithdrawalsRequest represents a WithdrawalsRequest event raised by the Pool contract.
type PoolWithdrawalsRequest struct {
	Receiver         common.Address
	WithdrawalAmount *big.Int
	BlockNumber      *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalsRequest is a free log retrieval operation binding the contract event 0x74ffedfd7821cf30dd556fac01944f4d077a2099ae55773bb89444cce29755f4.
//
// Solidity: event WithdrawalsRequest(address _receiver, uint256 _withdrawalAmount, uint256 _blockNumber)
func (_Pool *PoolFilterer) FilterWithdrawalsRequest(opts *bind.FilterOpts) (*PoolWithdrawalsRequestIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "WithdrawalsRequest")
	if err != nil {
		return nil, err
	}
	return &PoolWithdrawalsRequestIterator{contract: _Pool.contract, event: "WithdrawalsRequest", logs: logs, sub: sub}, nil
}

// WatchWithdrawalsRequest is a free log subscription operation binding the contract event 0x74ffedfd7821cf30dd556fac01944f4d077a2099ae55773bb89444cce29755f4.
//
// Solidity: event WithdrawalsRequest(address _receiver, uint256 _withdrawalAmount, uint256 _blockNumber)
func (_Pool *PoolFilterer) WatchWithdrawalsRequest(opts *bind.WatchOpts, sink chan<- *PoolWithdrawalsRequest) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "WithdrawalsRequest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolWithdrawalsRequest)
				if err := _Pool.contract.UnpackLog(event, "WithdrawalsRequest", log); err != nil {
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

// ParseWithdrawalsRequest is a log parse operation binding the contract event 0x74ffedfd7821cf30dd556fac01944f4d077a2099ae55773bb89444cce29755f4.
//
// Solidity: event WithdrawalsRequest(address _receiver, uint256 _withdrawalAmount, uint256 _blockNumber)
func (_Pool *PoolFilterer) ParseWithdrawalsRequest(log types.Log) (*PoolWithdrawalsRequest, error) {
	event := new(PoolWithdrawalsRequest)
	if err := _Pool.contract.UnpackLog(event, "WithdrawalsRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
