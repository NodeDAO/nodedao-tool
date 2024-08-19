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
	_ = abi.ConvertType
)

// WithdrawalRequestWithdrawalInfo is an auto generated low-level Go binding around an user-defined struct.
type WithdrawalRequestWithdrawalInfo struct {
	WithdrawalHeight   *big.Int
	WithdrawalExchange *big.Int
	IsClaim            uint64
	WithdrawalAmount   *big.Int
	ClaimAmount        *big.Int
}

// NodedaopoolMetaData contains all meta data concerning the Nodedaopool contract.
var NodedaopoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CanUnstakeETH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DelayTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DepositRootMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EigenPodMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientFunds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddr\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidApr\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidParameter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRequestId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidVersion\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidtypeId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PermissionDenied\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RestakingPodNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnstakeNotEnabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UpdateTimelocked\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WithrawalsRequestCannotClaimed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_oldApr\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_apr\",\"type\":\"uint256\"}],\"name\":\"AprUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalUnderlyingAsset\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_estimatedRewards\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"AssetsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_oldDao\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_dao\",\"type\":\"address\"}],\"name\":\"DaoChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_stakeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_mintAmount\",\"type\":\"uint256\"}],\"name\":\"EthStake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_unstakeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_ethAmount\",\"type\":\"uint256\"}],\"name\":\"EthUnstake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_oldAprManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_aprManager\",\"type\":\"address\"}],\"name\":\"RateManagerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Received\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_restakingPod\",\"type\":\"address\"}],\"name\":\"RestakingPodAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_strategyAmount\",\"type\":\"uint256\"}],\"name\":\"StrategyDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"StrategyReturn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_oldStrategyVault\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_strategyVault\",\"type\":\"address\"}],\"name\":\"StrategyVaultChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_oldUnstakeAllowed\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_unstakeAllowed\",\"type\":\"bool\"}],\"name\":\"UnstakeAllowedUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_oldValidatorManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_validatorManager\",\"type\":\"address\"}],\"name\":\"ValidatorManagerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"_pubkeys\",\"type\":\"bytes[]\"}],\"name\":\"ValidatorRegistration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_oldValidatorManager\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_validatorManager\",\"type\":\"address\"}],\"name\":\"ValidatorRegistryChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_oldWithdrawalDelayBlocks\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_withdrawalDelayBlocks\",\"type\":\"uint256\"}],\"name\":\"WithdrawalDelayChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_claimAmount\",\"type\":\"uint256\"}],\"name\":\"WithdrawalsClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_withdrawalAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"WithdrawalsRequest\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"APR_PERCENT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BLOCK_NUMBER_PER_YEAR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_APR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_WITHDRAWAL_DELAY_BLOCKS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPDATE_BLOCK_LIMIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ownerAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_withdrawalDelayBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_apr\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_totalUnderlyingAsset\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_dao\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_poolToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rateManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_validatorManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_depositContract\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_unstakeAllowed\",\"type\":\"bool\"}],\"name\":\"__BasePool_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_restakingPod\",\"type\":\"address\"}],\"name\":\"addRestakingPod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"canClaimWithdrawal\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimDelayedWithdrawals\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_requestIds\",\"type\":\"uint256[]\"}],\"name\":\"claimWithdrawals\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_unstakeAmount\",\"type\":\"uint256\"}],\"name\":\"convertToAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stakeAmount\",\"type\":\"uint256\"}],\"name\":\"convertToShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentApr\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dao\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositContract\",\"outputs\":[{\"internalType\":\"contractIDepositContract\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eigenLayerEigenPodManager\",\"outputs\":[{\"internalType\":\"contractIEigenPodManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"elRewardsAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exchangeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCLVaultAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRestakingPods\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"getUserWithdrawals\",\"outputs\":[{\"components\":[{\"internalType\":\"uint96\",\"name\":\"withdrawalHeight\",\"type\":\"uint96\"},{\"internalType\":\"uint96\",\"name\":\"withdrawalExchange\",\"type\":\"uint96\"},{\"internalType\":\"uint64\",\"name\":\"isClaim\",\"type\":\"uint64\"},{\"internalType\":\"uint128\",\"name\":\"withdrawalAmount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"claimAmount\",\"type\":\"uint128\"}],\"internalType\":\"structWithdrawalRequest.WithdrawalInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ownerAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_apr\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_dao\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_elRewardsAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_poolToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rateManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_validatorManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_depositContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_eigenLayerEigenPodManager\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_restakingPods\",\"type\":\"address[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolToken\",\"outputs\":[{\"internalType\":\"contractILsdETH\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rateManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"receiveRewards\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_depositContractRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_restakingPod\",\"type\":\"address\"},{\"internalType\":\"bytes[]\",\"name\":\"_pubkeys\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_signatures\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_depositDataRoots\",\"type\":\"bytes32[]\"}],\"name\":\"registerValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_unstakeAmount\",\"type\":\"uint256\"}],\"name\":\"requestWithdrawals\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardsUpdateBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_dao\",\"type\":\"address\"}],\"name\":\"setDao\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rateManager\",\"type\":\"address\"}],\"name\":\"setRateManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_strategyVault\",\"type\":\"address\"}],\"name\":\"setStrategyVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_unstakeAllowed\",\"type\":\"bool\"}],\"name\":\"setUnstakeAllowed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validatorManager\",\"type\":\"address\"}],\"name\":\"setValidatorManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validatorRegistry\",\"type\":\"address\"}],\"name\":\"setValidatorRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_withdrawalDelayBlocks\",\"type\":\"uint256\"}],\"name\":\"setWithdrawalDelayBlocks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"strategyAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"strategyDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"strategyReturn\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"strategyVault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalUnderlyingAsset\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalWithdrawalAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unstakeAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_unstakeAmount\",\"type\":\"uint256\"}],\"name\":\"unstakeETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_apr\",\"type\":\"uint256\"}],\"name\":\"updateApr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorRegistry\",\"outputs\":[{\"internalType\":\"contractIValidatorRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawCredentials\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawalDelayBlocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// NodedaopoolABI is the input ABI used to generate the binding from.
// Deprecated: Use NodedaopoolMetaData.ABI instead.
var NodedaopoolABI = NodedaopoolMetaData.ABI

var NodeDaoPoolABI abi.ABI

func init() {
	var err error
	NodeDaoPoolABI, err = abi.JSON(strings.NewReader(NodedaopoolABI))
	if err != nil {
		fmt.Println("err: ", err.Error())
		os.Exit(1)
	}
}

// Nodedaopool is an auto generated Go binding around an Ethereum contract.
type Nodedaopool struct {
	NodedaopoolCaller     // Read-only binding to the contract
	NodedaopoolTransactor // Write-only binding to the contract
	NodedaopoolFilterer   // Log filterer for contract events
}

// NodedaopoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type NodedaopoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodedaopoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NodedaopoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodedaopoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NodedaopoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodedaopoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NodedaopoolSession struct {
	Contract     *Nodedaopool      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NodedaopoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NodedaopoolCallerSession struct {
	Contract *NodedaopoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// NodedaopoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NodedaopoolTransactorSession struct {
	Contract     *NodedaopoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// NodedaopoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type NodedaopoolRaw struct {
	Contract *Nodedaopool // Generic contract binding to access the raw methods on
}

// NodedaopoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NodedaopoolCallerRaw struct {
	Contract *NodedaopoolCaller // Generic read-only contract binding to access the raw methods on
}

// NodedaopoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NodedaopoolTransactorRaw struct {
	Contract *NodedaopoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNodedaopool creates a new instance of Nodedaopool, bound to a specific deployed contract.
func NewNodedaopool(address common.Address, backend bind.ContractBackend) (*Nodedaopool, error) {
	contract, err := bindNodedaopool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Nodedaopool{NodedaopoolCaller: NodedaopoolCaller{contract: contract}, NodedaopoolTransactor: NodedaopoolTransactor{contract: contract}, NodedaopoolFilterer: NodedaopoolFilterer{contract: contract}}, nil
}

// NewNodedaopoolCaller creates a new read-only instance of Nodedaopool, bound to a specific deployed contract.
func NewNodedaopoolCaller(address common.Address, caller bind.ContractCaller) (*NodedaopoolCaller, error) {
	contract, err := bindNodedaopool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NodedaopoolCaller{contract: contract}, nil
}

// NewNodedaopoolTransactor creates a new write-only instance of Nodedaopool, bound to a specific deployed contract.
func NewNodedaopoolTransactor(address common.Address, transactor bind.ContractTransactor) (*NodedaopoolTransactor, error) {
	contract, err := bindNodedaopool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NodedaopoolTransactor{contract: contract}, nil
}

// NewNodedaopoolFilterer creates a new log filterer instance of Nodedaopool, bound to a specific deployed contract.
func NewNodedaopoolFilterer(address common.Address, filterer bind.ContractFilterer) (*NodedaopoolFilterer, error) {
	contract, err := bindNodedaopool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NodedaopoolFilterer{contract: contract}, nil
}

// bindNodedaopool binds a generic wrapper to an already deployed contract.
func bindNodedaopool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NodedaopoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Nodedaopool *NodedaopoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Nodedaopool.Contract.NodedaopoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Nodedaopool *NodedaopoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nodedaopool.Contract.NodedaopoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Nodedaopool *NodedaopoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Nodedaopool.Contract.NodedaopoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Nodedaopool *NodedaopoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Nodedaopool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Nodedaopool *NodedaopoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nodedaopool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Nodedaopool *NodedaopoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Nodedaopool.Contract.contract.Transact(opts, method, params...)
}

// APRPERCENT is a free data retrieval call binding the contract method 0x99ca63bb.
//
// Solidity: function APR_PERCENT() view returns(uint256)
func (_Nodedaopool *NodedaopoolCaller) APRPERCENT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nodedaopool.contract.Call(opts, &out, "APR_PERCENT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// APRPERCENT is a free data retrieval call binding the contract method 0x99ca63bb.
//
// Solidity: function APR_PERCENT() view returns(uint256)
func (_Nodedaopool *NodedaopoolSession) APRPERCENT() (*big.Int, error) {
	return _Nodedaopool.Contract.APRPERCENT(&_Nodedaopool.CallOpts)
}

// APRPERCENT is a free data retrieval call binding the contract method 0x99ca63bb.
//
// Solidity: function APR_PERCENT() view returns(uint256)
func (_Nodedaopool *NodedaopoolCallerSession) APRPERCENT() (*big.Int, error) {
	return _Nodedaopool.Contract.APRPERCENT(&_Nodedaopool.CallOpts)
}

// BLOCKNUMBERPERYEAR is a free data retrieval call binding the contract method 0xef213ddc.
//
// Solidity: function BLOCK_NUMBER_PER_YEAR() view returns(uint256)
func (_Nodedaopool *NodedaopoolCaller) BLOCKNUMBERPERYEAR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nodedaopool.contract.Call(opts, &out, "BLOCK_NUMBER_PER_YEAR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BLOCKNUMBERPERYEAR is a free data retrieval call binding the contract method 0xef213ddc.
//
// Solidity: function BLOCK_NUMBER_PER_YEAR() view returns(uint256)
func (_Nodedaopool *NodedaopoolSession) BLOCKNUMBERPERYEAR() (*big.Int, error) {
	return _Nodedaopool.Contract.BLOCKNUMBERPERYEAR(&_Nodedaopool.CallOpts)
}

// BLOCKNUMBERPERYEAR is a free data retrieval call binding the contract method 0xef213ddc.
//
// Solidity: function BLOCK_NUMBER_PER_YEAR() view returns(uint256)
func (_Nodedaopool *NodedaopoolCallerSession) BLOCKNUMBERPERYEAR() (*big.Int, error) {
	return _Nodedaopool.Contract.BLOCKNUMBERPERYEAR(&_Nodedaopool.CallOpts)
}

// MAXAPR is a free data retrieval call binding the contract method 0x86d8f78d.
//
// Solidity: function MAX_APR() view returns(uint256)
func (_Nodedaopool *NodedaopoolCaller) MAXAPR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nodedaopool.contract.Call(opts, &out, "MAX_APR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXAPR is a free data retrieval call binding the contract method 0x86d8f78d.
//
// Solidity: function MAX_APR() view returns(uint256)
func (_Nodedaopool *NodedaopoolSession) MAXAPR() (*big.Int, error) {
	return _Nodedaopool.Contract.MAXAPR(&_Nodedaopool.CallOpts)
}

// MAXAPR is a free data retrieval call binding the contract method 0x86d8f78d.
//
// Solidity: function MAX_APR() view returns(uint256)
func (_Nodedaopool *NodedaopoolCallerSession) MAXAPR() (*big.Int, error) {
	return _Nodedaopool.Contract.MAXAPR(&_Nodedaopool.CallOpts)
}

// MAXWITHDRAWALDELAYBLOCKS is a free data retrieval call binding the contract method 0xca661c04.
//
// Solidity: function MAX_WITHDRAWAL_DELAY_BLOCKS() view returns(uint256)
func (_Nodedaopool *NodedaopoolCaller) MAXWITHDRAWALDELAYBLOCKS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nodedaopool.contract.Call(opts, &out, "MAX_WITHDRAWAL_DELAY_BLOCKS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXWITHDRAWALDELAYBLOCKS is a free data retrieval call binding the contract method 0xca661c04.
//
// Solidity: function MAX_WITHDRAWAL_DELAY_BLOCKS() view returns(uint256)
func (_Nodedaopool *NodedaopoolSession) MAXWITHDRAWALDELAYBLOCKS() (*big.Int, error) {
	return _Nodedaopool.Contract.MAXWITHDRAWALDELAYBLOCKS(&_Nodedaopool.CallOpts)
}

// MAXWITHDRAWALDELAYBLOCKS is a free data retrieval call binding the contract method 0xca661c04.
//
// Solidity: function MAX_WITHDRAWAL_DELAY_BLOCKS() view returns(uint256)
func (_Nodedaopool *NodedaopoolCallerSession) MAXWITHDRAWALDELAYBLOCKS() (*big.Int, error) {
	return _Nodedaopool.Contract.MAXWITHDRAWALDELAYBLOCKS(&_Nodedaopool.CallOpts)
}

// UPDATEBLOCKLIMIT is a free data retrieval call binding the contract method 0xb7e34b0c.
//
// Solidity: function UPDATE_BLOCK_LIMIT() view returns(uint256)
func (_Nodedaopool *NodedaopoolCaller) UPDATEBLOCKLIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nodedaopool.contract.Call(opts, &out, "UPDATE_BLOCK_LIMIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UPDATEBLOCKLIMIT is a free data retrieval call binding the contract method 0xb7e34b0c.
//
// Solidity: function UPDATE_BLOCK_LIMIT() view returns(uint256)
func (_Nodedaopool *NodedaopoolSession) UPDATEBLOCKLIMIT() (*big.Int, error) {
	return _Nodedaopool.Contract.UPDATEBLOCKLIMIT(&_Nodedaopool.CallOpts)
}

// UPDATEBLOCKLIMIT is a free data retrieval call binding the contract method 0xb7e34b0c.
//
// Solidity: function UPDATE_BLOCK_LIMIT() view returns(uint256)
func (_Nodedaopool *NodedaopoolCallerSession) UPDATEBLOCKLIMIT() (*big.Int, error) {
	return _Nodedaopool.Contract.UPDATEBLOCKLIMIT(&_Nodedaopool.CallOpts)
}

// CanClaimWithdrawal is a free data retrieval call binding the contract method 0x78865853.
//
// Solidity: function canClaimWithdrawal(address _receiver, uint256 _requestId) view returns(bool)
func (_Nodedaopool *NodedaopoolCaller) CanClaimWithdrawal(opts *bind.CallOpts, _receiver common.Address, _requestId *big.Int) (bool, error) {
	var out []interface{}
	err := _Nodedaopool.contract.Call(opts, &out, "canClaimWithdrawal", _receiver, _requestId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanClaimWithdrawal is a free data retrieval call binding the contract method 0x78865853.
//
// Solidity: function canClaimWithdrawal(address _receiver, uint256 _requestId) view returns(bool)
func (_Nodedaopool *NodedaopoolSession) CanClaimWithdrawal(_receiver common.Address, _requestId *big.Int) (bool, error) {
	return _Nodedaopool.Contract.CanClaimWithdrawal(&_Nodedaopool.CallOpts, _receiver, _requestId)
}

// CanClaimWithdrawal is a free data retrieval call binding the contract method 0x78865853.
//
// Solidity: function canClaimWithdrawal(address _receiver, uint256 _requestId) view returns(bool)
func (_Nodedaopool *NodedaopoolCallerSession) CanClaimWithdrawal(_receiver common.Address, _requestId *big.Int) (bool, error) {
	return _Nodedaopool.Contract.CanClaimWithdrawal(&_Nodedaopool.CallOpts, _receiver, _requestId)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 _unstakeAmount) view returns(uint256)
func (_Nodedaopool *NodedaopoolCaller) ConvertToAssets(opts *bind.CallOpts, _unstakeAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Nodedaopool.contract.Call(opts, &out, "convertToAssets", _unstakeAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 _unstakeAmount) view returns(uint256)
func (_Nodedaopool *NodedaopoolSession) ConvertToAssets(_unstakeAmount *big.Int) (*big.Int, error) {
	return _Nodedaopool.Contract.ConvertToAssets(&_Nodedaopool.CallOpts, _unstakeAmount)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 _unstakeAmount) view returns(uint256)
func (_Nodedaopool *NodedaopoolCallerSession) ConvertToAssets(_unstakeAmount *big.Int) (*big.Int, error) {
	return _Nodedaopool.Contract.ConvertToAssets(&_Nodedaopool.CallOpts, _unstakeAmount)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 _stakeAmount) view returns(uint256)
func (_Nodedaopool *NodedaopoolCaller) ConvertToShares(opts *bind.CallOpts, _stakeAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Nodedaopool.contract.Call(opts, &out, "convertToShares", _stakeAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 _stakeAmount) view returns(uint256)
func (_Nodedaopool *NodedaopoolSession) ConvertToShares(_stakeAmount *big.Int) (*big.Int, error) {
	return _Nodedaopool.Contract.ConvertToShares(&_Nodedaopool.CallOpts, _stakeAmount)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 _stakeAmount) view returns(uint256)
func (_Nodedaopool *NodedaopoolCallerSession) ConvertToShares(_stakeAmount *big.Int) (*big.Int, error) {
	return _Nodedaopool.Contract.ConvertToShares(&_Nodedaopool.CallOpts, _stakeAmount)
}

// CurrentApr is a free data retrieval call binding the contract method 0x4f322ae8.
//
// Solidity: function currentApr() view returns(uint256)
func (_Nodedaopool *NodedaopoolCaller) CurrentApr(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nodedaopool.contract.Call(opts, &out, "currentApr")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentApr is a free data retrieval call binding the contract method 0x4f322ae8.
//
// Solidity: function currentApr() view returns(uint256)
func (_Nodedaopool *NodedaopoolSession) CurrentApr() (*big.Int, error) {
	return _Nodedaopool.Contract.CurrentApr(&_Nodedaopool.CallOpts)
}

// CurrentApr is a free data retrieval call binding the contract method 0x4f322ae8.
//
// Solidity: function currentApr() view returns(uint256)
func (_Nodedaopool *NodedaopoolCallerSession) CurrentApr() (*big.Int, error) {
	return _Nodedaopool.Contract.CurrentApr(&_Nodedaopool.CallOpts)
}

// Dao is a free data retrieval call binding the contract method 0x4162169f.
//
// Solidity: function dao() view returns(address)
func (_Nodedaopool *NodedaopoolCaller) Dao(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Nodedaopool.contract.Call(opts, &out, "dao")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Dao is a free data retrieval call binding the contract method 0x4162169f.
//
// Solidity: function dao() view returns(address)
func (_Nodedaopool *NodedaopoolSession) Dao() (common.Address, error) {
	return _Nodedaopool.Contract.Dao(&_Nodedaopool.CallOpts)
}

// Dao is a free data retrieval call binding the contract method 0x4162169f.
//
// Solidity: function dao() view returns(address)
func (_Nodedaopool *NodedaopoolCallerSession) Dao() (common.Address, error) {
	return _Nodedaopool.Contract.Dao(&_Nodedaopool.CallOpts)
}

// DepositContract is a free data retrieval call binding the contract method 0xe94ad65b.
//
// Solidity: function depositContract() view returns(address)
func (_Nodedaopool *NodedaopoolCaller) DepositContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Nodedaopool.contract.Call(opts, &out, "depositContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DepositContract is a free data retrieval call binding the contract method 0xe94ad65b.
//
// Solidity: function depositContract() view returns(address)
func (_Nodedaopool *NodedaopoolSession) DepositContract() (common.Address, error) {
	return _Nodedaopool.Contract.DepositContract(&_Nodedaopool.CallOpts)
}

// DepositContract is a free data retrieval call binding the contract method 0xe94ad65b.
//
// Solidity: function depositContract() view returns(address)
func (_Nodedaopool *NodedaopoolCallerSession) DepositContract() (common.Address, error) {
	return _Nodedaopool.Contract.DepositContract(&_Nodedaopool.CallOpts)
}

// EigenLayerEigenPodManager is a free data retrieval call binding the contract method 0x2ffb004f.
//
// Solidity: function eigenLayerEigenPodManager() view returns(address)
func (_Nodedaopool *NodedaopoolCaller) EigenLayerEigenPodManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Nodedaopool.contract.Call(opts, &out, "eigenLayerEigenPodManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EigenLayerEigenPodManager is a free data retrieval call binding the contract method 0x2ffb004f.
//
// Solidity: function eigenLayerEigenPodManager() view returns(address)
func (_Nodedaopool *NodedaopoolSession) EigenLayerEigenPodManager() (common.Address, error) {
	return _Nodedaopool.Contract.EigenLayerEigenPodManager(&_Nodedaopool.CallOpts)
}

// EigenLayerEigenPodManager is a free data retrieval call binding the contract method 0x2ffb004f.
//
// Solidity: function eigenLayerEigenPodManager() view returns(address)
func (_Nodedaopool *NodedaopoolCallerSession) EigenLayerEigenPodManager() (common.Address, error) {
	return _Nodedaopool.Contract.EigenLayerEigenPodManager(&_Nodedaopool.CallOpts)
}

// ElRewardsAddress is a free data retrieval call binding the contract method 0x215bd1b1.
//
// Solidity: function elRewardsAddress() view returns(address)
func (_Nodedaopool *NodedaopoolCaller) ElRewardsAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Nodedaopool.contract.Call(opts, &out, "elRewardsAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ElRewardsAddress is a free data retrieval call binding the contract method 0x215bd1b1.
//
// Solidity: function elRewardsAddress() view returns(address)
func (_Nodedaopool *NodedaopoolSession) ElRewardsAddress() (common.Address, error) {
	return _Nodedaopool.Contract.ElRewardsAddress(&_Nodedaopool.CallOpts)
}

// ElRewardsAddress is a free data retrieval call binding the contract method 0x215bd1b1.
//
// Solidity: function elRewardsAddress() view returns(address)
func (_Nodedaopool *NodedaopoolCallerSession) ElRewardsAddress() (common.Address, error) {
	return _Nodedaopool.Contract.ElRewardsAddress(&_Nodedaopool.CallOpts)
}

// ExchangeRate is a free data retrieval call binding the contract method 0x3ba0b9a9.
//
// Solidity: function exchangeRate() view returns(uint256)
func (_Nodedaopool *NodedaopoolCaller) ExchangeRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nodedaopool.contract.Call(opts, &out, "exchangeRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExchangeRate is a free data retrieval call binding the contract method 0x3ba0b9a9.
//
// Solidity: function exchangeRate() view returns(uint256)
func (_Nodedaopool *NodedaopoolSession) ExchangeRate() (*big.Int, error) {
	return _Nodedaopool.Contract.ExchangeRate(&_Nodedaopool.CallOpts)
}

// ExchangeRate is a free data retrieval call binding the contract method 0x3ba0b9a9.
//
// Solidity: function exchangeRate() view returns(uint256)
func (_Nodedaopool *NodedaopoolCallerSession) ExchangeRate() (*big.Int, error) {
	return _Nodedaopool.Contract.ExchangeRate(&_Nodedaopool.CallOpts)
}

// GetCLVaultAmount is a free data retrieval call binding the contract method 0xf60b6fe6.
//
// Solidity: function getCLVaultAmount() view returns(uint256)
func (_Nodedaopool *NodedaopoolCaller) GetCLVaultAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nodedaopool.contract.Call(opts, &out, "getCLVaultAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCLVaultAmount is a free data retrieval call binding the contract method 0xf60b6fe6.
//
// Solidity: function getCLVaultAmount() view returns(uint256)
func (_Nodedaopool *NodedaopoolSession) GetCLVaultAmount() (*big.Int, error) {
	return _Nodedaopool.Contract.GetCLVaultAmount(&_Nodedaopool.CallOpts)
}

// GetCLVaultAmount is a free data retrieval call binding the contract method 0xf60b6fe6.
//
// Solidity: function getCLVaultAmount() view returns(uint256)
func (_Nodedaopool *NodedaopoolCallerSession) GetCLVaultAmount() (*big.Int, error) {
	return _Nodedaopool.Contract.GetCLVaultAmount(&_Nodedaopool.CallOpts)
}

// GetPoolAmount is a free data retrieval call binding the contract method 0x4ab4ba42.
//
// Solidity: function getPoolAmount() view returns(uint256)
func (_Nodedaopool *NodedaopoolCaller) GetPoolAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nodedaopool.contract.Call(opts, &out, "getPoolAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPoolAmount is a free data retrieval call binding the contract method 0x4ab4ba42.
//
// Solidity: function getPoolAmount() view returns(uint256)
func (_Nodedaopool *NodedaopoolSession) GetPoolAmount() (*big.Int, error) {
	return _Nodedaopool.Contract.GetPoolAmount(&_Nodedaopool.CallOpts)
}

// GetPoolAmount is a free data retrieval call binding the contract method 0x4ab4ba42.
//
// Solidity: function getPoolAmount() view returns(uint256)
func (_Nodedaopool *NodedaopoolCallerSession) GetPoolAmount() (*big.Int, error) {
	return _Nodedaopool.Contract.GetPoolAmount(&_Nodedaopool.CallOpts)
}

// GetRestakingPods is a free data retrieval call binding the contract method 0x61e9ecf3.
//
// Solidity: function getRestakingPods() view returns(address[])
func (_Nodedaopool *NodedaopoolCaller) GetRestakingPods(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Nodedaopool.contract.Call(opts, &out, "getRestakingPods")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRestakingPods is a free data retrieval call binding the contract method 0x61e9ecf3.
//
// Solidity: function getRestakingPods() view returns(address[])
func (_Nodedaopool *NodedaopoolSession) GetRestakingPods() ([]common.Address, error) {
	return _Nodedaopool.Contract.GetRestakingPods(&_Nodedaopool.CallOpts)
}

// GetRestakingPods is a free data retrieval call binding the contract method 0x61e9ecf3.
//
// Solidity: function getRestakingPods() view returns(address[])
func (_Nodedaopool *NodedaopoolCallerSession) GetRestakingPods() ([]common.Address, error) {
	return _Nodedaopool.Contract.GetRestakingPods(&_Nodedaopool.CallOpts)
}

// GetUserWithdrawals is a free data retrieval call binding the contract method 0xe502eb68.
//
// Solidity: function getUserWithdrawals(address _receiver) view returns((uint96,uint96,uint64,uint128,uint128)[])
func (_Nodedaopool *NodedaopoolCaller) GetUserWithdrawals(opts *bind.CallOpts, _receiver common.Address) ([]WithdrawalRequestWithdrawalInfo, error) {
	var out []interface{}
	err := _Nodedaopool.contract.Call(opts, &out, "getUserWithdrawals", _receiver)

	if err != nil {
		return *new([]WithdrawalRequestWithdrawalInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]WithdrawalRequestWithdrawalInfo)).(*[]WithdrawalRequestWithdrawalInfo)

	return out0, err

}

// GetUserWithdrawals is a free data retrieval call binding the contract method 0xe502eb68.
//
// Solidity: function getUserWithdrawals(address _receiver) view returns((uint96,uint96,uint64,uint128,uint128)[])
func (_Nodedaopool *NodedaopoolSession) GetUserWithdrawals(_receiver common.Address) ([]WithdrawalRequestWithdrawalInfo, error) {
	return _Nodedaopool.Contract.GetUserWithdrawals(&_Nodedaopool.CallOpts, _receiver)
}

// GetUserWithdrawals is a free data retrieval call binding the contract method 0xe502eb68.
//
// Solidity: function getUserWithdrawals(address _receiver) view returns((uint96,uint96,uint64,uint128,uint128)[])
func (_Nodedaopool *NodedaopoolCallerSession) GetUserWithdrawals(_receiver common.Address) ([]WithdrawalRequestWithdrawalInfo, error) {
	return _Nodedaopool.Contract.GetUserWithdrawals(&_Nodedaopool.CallOpts, _receiver)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_Nodedaopool *NodedaopoolCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Nodedaopool.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_Nodedaopool *NodedaopoolSession) Implementation() (common.Address, error) {
	return _Nodedaopool.Contract.Implementation(&_Nodedaopool.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_Nodedaopool *NodedaopoolCallerSession) Implementation() (common.Address, error) {
	return _Nodedaopool.Contract.Implementation(&_Nodedaopool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Nodedaopool *NodedaopoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Nodedaopool.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Nodedaopool *NodedaopoolSession) Owner() (common.Address, error) {
	return _Nodedaopool.Contract.Owner(&_Nodedaopool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Nodedaopool *NodedaopoolCallerSession) Owner() (common.Address, error) {
	return _Nodedaopool.Contract.Owner(&_Nodedaopool.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Nodedaopool *NodedaopoolCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Nodedaopool.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Nodedaopool *NodedaopoolSession) Paused() (bool, error) {
	return _Nodedaopool.Contract.Paused(&_Nodedaopool.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Nodedaopool *NodedaopoolCallerSession) Paused() (bool, error) {
	return _Nodedaopool.Contract.Paused(&_Nodedaopool.CallOpts)
}

// PoolToken is a free data retrieval call binding the contract method 0xcbdf382c.
//
// Solidity: function poolToken() view returns(address)
func (_Nodedaopool *NodedaopoolCaller) PoolToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Nodedaopool.contract.Call(opts, &out, "poolToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolToken is a free data retrieval call binding the contract method 0xcbdf382c.
//
// Solidity: function poolToken() view returns(address)
func (_Nodedaopool *NodedaopoolSession) PoolToken() (common.Address, error) {
	return _Nodedaopool.Contract.PoolToken(&_Nodedaopool.CallOpts)
}

// PoolToken is a free data retrieval call binding the contract method 0xcbdf382c.
//
// Solidity: function poolToken() view returns(address)
func (_Nodedaopool *NodedaopoolCallerSession) PoolToken() (common.Address, error) {
	return _Nodedaopool.Contract.PoolToken(&_Nodedaopool.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Nodedaopool *NodedaopoolCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Nodedaopool.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Nodedaopool *NodedaopoolSession) ProxiableUUID() ([32]byte, error) {
	return _Nodedaopool.Contract.ProxiableUUID(&_Nodedaopool.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Nodedaopool *NodedaopoolCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Nodedaopool.Contract.ProxiableUUID(&_Nodedaopool.CallOpts)
}

// RateManager is a free data retrieval call binding the contract method 0x71c3cd88.
//
// Solidity: function rateManager() view returns(address)
func (_Nodedaopool *NodedaopoolCaller) RateManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Nodedaopool.contract.Call(opts, &out, "rateManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RateManager is a free data retrieval call binding the contract method 0x71c3cd88.
//
// Solidity: function rateManager() view returns(address)
func (_Nodedaopool *NodedaopoolSession) RateManager() (common.Address, error) {
	return _Nodedaopool.Contract.RateManager(&_Nodedaopool.CallOpts)
}

// RateManager is a free data retrieval call binding the contract method 0x71c3cd88.
//
// Solidity: function rateManager() view returns(address)
func (_Nodedaopool *NodedaopoolCallerSession) RateManager() (common.Address, error) {
	return _Nodedaopool.Contract.RateManager(&_Nodedaopool.CallOpts)
}

// RewardsUpdateBlock is a free data retrieval call binding the contract method 0x88f36cdd.
//
// Solidity: function rewardsUpdateBlock() view returns(uint256)
func (_Nodedaopool *NodedaopoolCaller) RewardsUpdateBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nodedaopool.contract.Call(opts, &out, "rewardsUpdateBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardsUpdateBlock is a free data retrieval call binding the contract method 0x88f36cdd.
//
// Solidity: function rewardsUpdateBlock() view returns(uint256)
func (_Nodedaopool *NodedaopoolSession) RewardsUpdateBlock() (*big.Int, error) {
	return _Nodedaopool.Contract.RewardsUpdateBlock(&_Nodedaopool.CallOpts)
}

// RewardsUpdateBlock is a free data retrieval call binding the contract method 0x88f36cdd.
//
// Solidity: function rewardsUpdateBlock() view returns(uint256)
func (_Nodedaopool *NodedaopoolCallerSession) RewardsUpdateBlock() (*big.Int, error) {
	return _Nodedaopool.Contract.RewardsUpdateBlock(&_Nodedaopool.CallOpts)
}

// StrategyAmount is a free data retrieval call binding the contract method 0x8dd7f293.
//
// Solidity: function strategyAmount() view returns(uint256)
func (_Nodedaopool *NodedaopoolCaller) StrategyAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nodedaopool.contract.Call(opts, &out, "strategyAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StrategyAmount is a free data retrieval call binding the contract method 0x8dd7f293.
//
// Solidity: function strategyAmount() view returns(uint256)
func (_Nodedaopool *NodedaopoolSession) StrategyAmount() (*big.Int, error) {
	return _Nodedaopool.Contract.StrategyAmount(&_Nodedaopool.CallOpts)
}

// StrategyAmount is a free data retrieval call binding the contract method 0x8dd7f293.
//
// Solidity: function strategyAmount() view returns(uint256)
func (_Nodedaopool *NodedaopoolCallerSession) StrategyAmount() (*big.Int, error) {
	return _Nodedaopool.Contract.StrategyAmount(&_Nodedaopool.CallOpts)
}

// StrategyVault is a free data retrieval call binding the contract method 0x40099098.
//
// Solidity: function strategyVault() view returns(address)
func (_Nodedaopool *NodedaopoolCaller) StrategyVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Nodedaopool.contract.Call(opts, &out, "strategyVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StrategyVault is a free data retrieval call binding the contract method 0x40099098.
//
// Solidity: function strategyVault() view returns(address)
func (_Nodedaopool *NodedaopoolSession) StrategyVault() (common.Address, error) {
	return _Nodedaopool.Contract.StrategyVault(&_Nodedaopool.CallOpts)
}

// StrategyVault is a free data retrieval call binding the contract method 0x40099098.
//
// Solidity: function strategyVault() view returns(address)
func (_Nodedaopool *NodedaopoolCallerSession) StrategyVault() (common.Address, error) {
	return _Nodedaopool.Contract.StrategyVault(&_Nodedaopool.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_Nodedaopool *NodedaopoolCaller) TotalAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nodedaopool.contract.Call(opts, &out, "totalAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_Nodedaopool *NodedaopoolSession) TotalAssets() (*big.Int, error) {
	return _Nodedaopool.Contract.TotalAssets(&_Nodedaopool.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_Nodedaopool *NodedaopoolCallerSession) TotalAssets() (*big.Int, error) {
	return _Nodedaopool.Contract.TotalAssets(&_Nodedaopool.CallOpts)
}

// TotalUnderlyingAsset is a free data retrieval call binding the contract method 0xddfa63ae.
//
// Solidity: function totalUnderlyingAsset() view returns(uint256)
func (_Nodedaopool *NodedaopoolCaller) TotalUnderlyingAsset(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nodedaopool.contract.Call(opts, &out, "totalUnderlyingAsset")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalUnderlyingAsset is a free data retrieval call binding the contract method 0xddfa63ae.
//
// Solidity: function totalUnderlyingAsset() view returns(uint256)
func (_Nodedaopool *NodedaopoolSession) TotalUnderlyingAsset() (*big.Int, error) {
	return _Nodedaopool.Contract.TotalUnderlyingAsset(&_Nodedaopool.CallOpts)
}

// TotalUnderlyingAsset is a free data retrieval call binding the contract method 0xddfa63ae.
//
// Solidity: function totalUnderlyingAsset() view returns(uint256)
func (_Nodedaopool *NodedaopoolCallerSession) TotalUnderlyingAsset() (*big.Int, error) {
	return _Nodedaopool.Contract.TotalUnderlyingAsset(&_Nodedaopool.CallOpts)
}

// TotalWithdrawalAmount is a free data retrieval call binding the contract method 0xe23e0e08.
//
// Solidity: function totalWithdrawalAmount() view returns(uint256)
func (_Nodedaopool *NodedaopoolCaller) TotalWithdrawalAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nodedaopool.contract.Call(opts, &out, "totalWithdrawalAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalWithdrawalAmount is a free data retrieval call binding the contract method 0xe23e0e08.
//
// Solidity: function totalWithdrawalAmount() view returns(uint256)
func (_Nodedaopool *NodedaopoolSession) TotalWithdrawalAmount() (*big.Int, error) {
	return _Nodedaopool.Contract.TotalWithdrawalAmount(&_Nodedaopool.CallOpts)
}

// TotalWithdrawalAmount is a free data retrieval call binding the contract method 0xe23e0e08.
//
// Solidity: function totalWithdrawalAmount() view returns(uint256)
func (_Nodedaopool *NodedaopoolCallerSession) TotalWithdrawalAmount() (*big.Int, error) {
	return _Nodedaopool.Contract.TotalWithdrawalAmount(&_Nodedaopool.CallOpts)
}

// TypeId is a free data retrieval call binding the contract method 0x8f940f63.
//
// Solidity: function typeId() pure returns(bytes32)
func (_Nodedaopool *NodedaopoolCaller) TypeId(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Nodedaopool.contract.Call(opts, &out, "typeId")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TypeId is a free data retrieval call binding the contract method 0x8f940f63.
//
// Solidity: function typeId() pure returns(bytes32)
func (_Nodedaopool *NodedaopoolSession) TypeId() ([32]byte, error) {
	return _Nodedaopool.Contract.TypeId(&_Nodedaopool.CallOpts)
}

// TypeId is a free data retrieval call binding the contract method 0x8f940f63.
//
// Solidity: function typeId() pure returns(bytes32)
func (_Nodedaopool *NodedaopoolCallerSession) TypeId() ([32]byte, error) {
	return _Nodedaopool.Contract.TypeId(&_Nodedaopool.CallOpts)
}

// UnstakeAllowed is a free data retrieval call binding the contract method 0xee018b6d.
//
// Solidity: function unstakeAllowed() view returns(bool)
func (_Nodedaopool *NodedaopoolCaller) UnstakeAllowed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Nodedaopool.contract.Call(opts, &out, "unstakeAllowed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UnstakeAllowed is a free data retrieval call binding the contract method 0xee018b6d.
//
// Solidity: function unstakeAllowed() view returns(bool)
func (_Nodedaopool *NodedaopoolSession) UnstakeAllowed() (bool, error) {
	return _Nodedaopool.Contract.UnstakeAllowed(&_Nodedaopool.CallOpts)
}

// UnstakeAllowed is a free data retrieval call binding the contract method 0xee018b6d.
//
// Solidity: function unstakeAllowed() view returns(bool)
func (_Nodedaopool *NodedaopoolCallerSession) UnstakeAllowed() (bool, error) {
	return _Nodedaopool.Contract.UnstakeAllowed(&_Nodedaopool.CallOpts)
}

// ValidatorManager is a free data retrieval call binding the contract method 0xfe55bde9.
//
// Solidity: function validatorManager() view returns(address)
func (_Nodedaopool *NodedaopoolCaller) ValidatorManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Nodedaopool.contract.Call(opts, &out, "validatorManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ValidatorManager is a free data retrieval call binding the contract method 0xfe55bde9.
//
// Solidity: function validatorManager() view returns(address)
func (_Nodedaopool *NodedaopoolSession) ValidatorManager() (common.Address, error) {
	return _Nodedaopool.Contract.ValidatorManager(&_Nodedaopool.CallOpts)
}

// ValidatorManager is a free data retrieval call binding the contract method 0xfe55bde9.
//
// Solidity: function validatorManager() view returns(address)
func (_Nodedaopool *NodedaopoolCallerSession) ValidatorManager() (common.Address, error) {
	return _Nodedaopool.Contract.ValidatorManager(&_Nodedaopool.CallOpts)
}

// ValidatorRegistry is a free data retrieval call binding the contract method 0xf376ebbb.
//
// Solidity: function validatorRegistry() view returns(address)
func (_Nodedaopool *NodedaopoolCaller) ValidatorRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Nodedaopool.contract.Call(opts, &out, "validatorRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ValidatorRegistry is a free data retrieval call binding the contract method 0xf376ebbb.
//
// Solidity: function validatorRegistry() view returns(address)
func (_Nodedaopool *NodedaopoolSession) ValidatorRegistry() (common.Address, error) {
	return _Nodedaopool.Contract.ValidatorRegistry(&_Nodedaopool.CallOpts)
}

// ValidatorRegistry is a free data retrieval call binding the contract method 0xf376ebbb.
//
// Solidity: function validatorRegistry() view returns(address)
func (_Nodedaopool *NodedaopoolCallerSession) ValidatorRegistry() (common.Address, error) {
	return _Nodedaopool.Contract.ValidatorRegistry(&_Nodedaopool.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint8)
func (_Nodedaopool *NodedaopoolCaller) Version(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Nodedaopool.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint8)
func (_Nodedaopool *NodedaopoolSession) Version() (uint8, error) {
	return _Nodedaopool.Contract.Version(&_Nodedaopool.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint8)
func (_Nodedaopool *NodedaopoolCallerSession) Version() (uint8, error) {
	return _Nodedaopool.Contract.Version(&_Nodedaopool.CallOpts)
}

// WithdrawCredentials is a free data retrieval call binding the contract method 0xb62be945.
//
// Solidity: function withdrawCredentials() view returns(bytes[])
func (_Nodedaopool *NodedaopoolCaller) WithdrawCredentials(opts *bind.CallOpts) ([][]byte, error) {
	var out []interface{}
	err := _Nodedaopool.contract.Call(opts, &out, "withdrawCredentials")

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

// WithdrawCredentials is a free data retrieval call binding the contract method 0xb62be945.
//
// Solidity: function withdrawCredentials() view returns(bytes[])
func (_Nodedaopool *NodedaopoolSession) WithdrawCredentials() ([][]byte, error) {
	return _Nodedaopool.Contract.WithdrawCredentials(&_Nodedaopool.CallOpts)
}

// WithdrawCredentials is a free data retrieval call binding the contract method 0xb62be945.
//
// Solidity: function withdrawCredentials() view returns(bytes[])
func (_Nodedaopool *NodedaopoolCallerSession) WithdrawCredentials() ([][]byte, error) {
	return _Nodedaopool.Contract.WithdrawCredentials(&_Nodedaopool.CallOpts)
}

// WithdrawalDelayBlocks is a free data retrieval call binding the contract method 0x50f73e7c.
//
// Solidity: function withdrawalDelayBlocks() view returns(uint256)
func (_Nodedaopool *NodedaopoolCaller) WithdrawalDelayBlocks(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nodedaopool.contract.Call(opts, &out, "withdrawalDelayBlocks")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawalDelayBlocks is a free data retrieval call binding the contract method 0x50f73e7c.
//
// Solidity: function withdrawalDelayBlocks() view returns(uint256)
func (_Nodedaopool *NodedaopoolSession) WithdrawalDelayBlocks() (*big.Int, error) {
	return _Nodedaopool.Contract.WithdrawalDelayBlocks(&_Nodedaopool.CallOpts)
}

// WithdrawalDelayBlocks is a free data retrieval call binding the contract method 0x50f73e7c.
//
// Solidity: function withdrawalDelayBlocks() view returns(uint256)
func (_Nodedaopool *NodedaopoolCallerSession) WithdrawalDelayBlocks() (*big.Int, error) {
	return _Nodedaopool.Contract.WithdrawalDelayBlocks(&_Nodedaopool.CallOpts)
}

// BasePoolInit is a paid mutator transaction binding the contract method 0x73c62d5d.
//
// Solidity: function __BasePool_init(address _ownerAddr, uint256 _withdrawalDelayBlocks, uint256 _apr, uint256 _totalUnderlyingAsset, address _dao, address _poolToken, address _rateManager, address _validatorManager, address _depositContract, bool _unstakeAllowed) returns()
func (_Nodedaopool *NodedaopoolTransactor) BasePoolInit(opts *bind.TransactOpts, _ownerAddr common.Address, _withdrawalDelayBlocks *big.Int, _apr *big.Int, _totalUnderlyingAsset *big.Int, _dao common.Address, _poolToken common.Address, _rateManager common.Address, _validatorManager common.Address, _depositContract common.Address, _unstakeAllowed bool) (*types.Transaction, error) {
	return _Nodedaopool.contract.Transact(opts, "__BasePool_init", _ownerAddr, _withdrawalDelayBlocks, _apr, _totalUnderlyingAsset, _dao, _poolToken, _rateManager, _validatorManager, _depositContract, _unstakeAllowed)
}

// BasePoolInit is a paid mutator transaction binding the contract method 0x73c62d5d.
//
// Solidity: function __BasePool_init(address _ownerAddr, uint256 _withdrawalDelayBlocks, uint256 _apr, uint256 _totalUnderlyingAsset, address _dao, address _poolToken, address _rateManager, address _validatorManager, address _depositContract, bool _unstakeAllowed) returns()
func (_Nodedaopool *NodedaopoolSession) BasePoolInit(_ownerAddr common.Address, _withdrawalDelayBlocks *big.Int, _apr *big.Int, _totalUnderlyingAsset *big.Int, _dao common.Address, _poolToken common.Address, _rateManager common.Address, _validatorManager common.Address, _depositContract common.Address, _unstakeAllowed bool) (*types.Transaction, error) {
	return _Nodedaopool.Contract.BasePoolInit(&_Nodedaopool.TransactOpts, _ownerAddr, _withdrawalDelayBlocks, _apr, _totalUnderlyingAsset, _dao, _poolToken, _rateManager, _validatorManager, _depositContract, _unstakeAllowed)
}

// BasePoolInit is a paid mutator transaction binding the contract method 0x73c62d5d.
//
// Solidity: function __BasePool_init(address _ownerAddr, uint256 _withdrawalDelayBlocks, uint256 _apr, uint256 _totalUnderlyingAsset, address _dao, address _poolToken, address _rateManager, address _validatorManager, address _depositContract, bool _unstakeAllowed) returns()
func (_Nodedaopool *NodedaopoolTransactorSession) BasePoolInit(_ownerAddr common.Address, _withdrawalDelayBlocks *big.Int, _apr *big.Int, _totalUnderlyingAsset *big.Int, _dao common.Address, _poolToken common.Address, _rateManager common.Address, _validatorManager common.Address, _depositContract common.Address, _unstakeAllowed bool) (*types.Transaction, error) {
	return _Nodedaopool.Contract.BasePoolInit(&_Nodedaopool.TransactOpts, _ownerAddr, _withdrawalDelayBlocks, _apr, _totalUnderlyingAsset, _dao, _poolToken, _rateManager, _validatorManager, _depositContract, _unstakeAllowed)
}

// AddRestakingPod is a paid mutator transaction binding the contract method 0xebeac515.
//
// Solidity: function addRestakingPod(address _restakingPod) returns()
func (_Nodedaopool *NodedaopoolTransactor) AddRestakingPod(opts *bind.TransactOpts, _restakingPod common.Address) (*types.Transaction, error) {
	return _Nodedaopool.contract.Transact(opts, "addRestakingPod", _restakingPod)
}

// AddRestakingPod is a paid mutator transaction binding the contract method 0xebeac515.
//
// Solidity: function addRestakingPod(address _restakingPod) returns()
func (_Nodedaopool *NodedaopoolSession) AddRestakingPod(_restakingPod common.Address) (*types.Transaction, error) {
	return _Nodedaopool.Contract.AddRestakingPod(&_Nodedaopool.TransactOpts, _restakingPod)
}

// AddRestakingPod is a paid mutator transaction binding the contract method 0xebeac515.
//
// Solidity: function addRestakingPod(address _restakingPod) returns()
func (_Nodedaopool *NodedaopoolTransactorSession) AddRestakingPod(_restakingPod common.Address) (*types.Transaction, error) {
	return _Nodedaopool.Contract.AddRestakingPod(&_Nodedaopool.TransactOpts, _restakingPod)
}

// ClaimDelayedWithdrawals is a paid mutator transaction binding the contract method 0x28837c07.
//
// Solidity: function claimDelayedWithdrawals() returns()
func (_Nodedaopool *NodedaopoolTransactor) ClaimDelayedWithdrawals(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nodedaopool.contract.Transact(opts, "claimDelayedWithdrawals")
}

// ClaimDelayedWithdrawals is a paid mutator transaction binding the contract method 0x28837c07.
//
// Solidity: function claimDelayedWithdrawals() returns()
func (_Nodedaopool *NodedaopoolSession) ClaimDelayedWithdrawals() (*types.Transaction, error) {
	return _Nodedaopool.Contract.ClaimDelayedWithdrawals(&_Nodedaopool.TransactOpts)
}

// ClaimDelayedWithdrawals is a paid mutator transaction binding the contract method 0x28837c07.
//
// Solidity: function claimDelayedWithdrawals() returns()
func (_Nodedaopool *NodedaopoolTransactorSession) ClaimDelayedWithdrawals() (*types.Transaction, error) {
	return _Nodedaopool.Contract.ClaimDelayedWithdrawals(&_Nodedaopool.TransactOpts)
}

// ClaimWithdrawals is a paid mutator transaction binding the contract method 0x7139053e.
//
// Solidity: function claimWithdrawals(address _receiver, uint256[] _requestIds) returns()
func (_Nodedaopool *NodedaopoolTransactor) ClaimWithdrawals(opts *bind.TransactOpts, _receiver common.Address, _requestIds []*big.Int) (*types.Transaction, error) {
	return _Nodedaopool.contract.Transact(opts, "claimWithdrawals", _receiver, _requestIds)
}

// ClaimWithdrawals is a paid mutator transaction binding the contract method 0x7139053e.
//
// Solidity: function claimWithdrawals(address _receiver, uint256[] _requestIds) returns()
func (_Nodedaopool *NodedaopoolSession) ClaimWithdrawals(_receiver common.Address, _requestIds []*big.Int) (*types.Transaction, error) {
	return _Nodedaopool.Contract.ClaimWithdrawals(&_Nodedaopool.TransactOpts, _receiver, _requestIds)
}

// ClaimWithdrawals is a paid mutator transaction binding the contract method 0x7139053e.
//
// Solidity: function claimWithdrawals(address _receiver, uint256[] _requestIds) returns()
func (_Nodedaopool *NodedaopoolTransactorSession) ClaimWithdrawals(_receiver common.Address, _requestIds []*big.Int) (*types.Transaction, error) {
	return _Nodedaopool.Contract.ClaimWithdrawals(&_Nodedaopool.TransactOpts, _receiver, _requestIds)
}

// Initialize is a paid mutator transaction binding the contract method 0x983e87bc.
//
// Solidity: function initialize(address _ownerAddr, uint256 _apr, address _dao, address _elRewardsAddress, address _poolToken, address _rateManager, address _validatorManager, address _depositContract, address _eigenLayerEigenPodManager, address[] _restakingPods) returns()
func (_Nodedaopool *NodedaopoolTransactor) Initialize(opts *bind.TransactOpts, _ownerAddr common.Address, _apr *big.Int, _dao common.Address, _elRewardsAddress common.Address, _poolToken common.Address, _rateManager common.Address, _validatorManager common.Address, _depositContract common.Address, _eigenLayerEigenPodManager common.Address, _restakingPods []common.Address) (*types.Transaction, error) {
	return _Nodedaopool.contract.Transact(opts, "initialize", _ownerAddr, _apr, _dao, _elRewardsAddress, _poolToken, _rateManager, _validatorManager, _depositContract, _eigenLayerEigenPodManager, _restakingPods)
}

// Initialize is a paid mutator transaction binding the contract method 0x983e87bc.
//
// Solidity: function initialize(address _ownerAddr, uint256 _apr, address _dao, address _elRewardsAddress, address _poolToken, address _rateManager, address _validatorManager, address _depositContract, address _eigenLayerEigenPodManager, address[] _restakingPods) returns()
func (_Nodedaopool *NodedaopoolSession) Initialize(_ownerAddr common.Address, _apr *big.Int, _dao common.Address, _elRewardsAddress common.Address, _poolToken common.Address, _rateManager common.Address, _validatorManager common.Address, _depositContract common.Address, _eigenLayerEigenPodManager common.Address, _restakingPods []common.Address) (*types.Transaction, error) {
	return _Nodedaopool.Contract.Initialize(&_Nodedaopool.TransactOpts, _ownerAddr, _apr, _dao, _elRewardsAddress, _poolToken, _rateManager, _validatorManager, _depositContract, _eigenLayerEigenPodManager, _restakingPods)
}

// Initialize is a paid mutator transaction binding the contract method 0x983e87bc.
//
// Solidity: function initialize(address _ownerAddr, uint256 _apr, address _dao, address _elRewardsAddress, address _poolToken, address _rateManager, address _validatorManager, address _depositContract, address _eigenLayerEigenPodManager, address[] _restakingPods) returns()
func (_Nodedaopool *NodedaopoolTransactorSession) Initialize(_ownerAddr common.Address, _apr *big.Int, _dao common.Address, _elRewardsAddress common.Address, _poolToken common.Address, _rateManager common.Address, _validatorManager common.Address, _depositContract common.Address, _eigenLayerEigenPodManager common.Address, _restakingPods []common.Address) (*types.Transaction, error) {
	return _Nodedaopool.Contract.Initialize(&_Nodedaopool.TransactOpts, _ownerAddr, _apr, _dao, _elRewardsAddress, _poolToken, _rateManager, _validatorManager, _depositContract, _eigenLayerEigenPodManager, _restakingPods)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Nodedaopool *NodedaopoolTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nodedaopool.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Nodedaopool *NodedaopoolSession) Pause() (*types.Transaction, error) {
	return _Nodedaopool.Contract.Pause(&_Nodedaopool.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Nodedaopool *NodedaopoolTransactorSession) Pause() (*types.Transaction, error) {
	return _Nodedaopool.Contract.Pause(&_Nodedaopool.TransactOpts)
}

// ReceiveRewards is a paid mutator transaction binding the contract method 0xd77868da.
//
// Solidity: function receiveRewards(uint256 ) payable returns()
func (_Nodedaopool *NodedaopoolTransactor) ReceiveRewards(opts *bind.TransactOpts, arg0 *big.Int) (*types.Transaction, error) {
	return _Nodedaopool.contract.Transact(opts, "receiveRewards", arg0)
}

// ReceiveRewards is a paid mutator transaction binding the contract method 0xd77868da.
//
// Solidity: function receiveRewards(uint256 ) payable returns()
func (_Nodedaopool *NodedaopoolSession) ReceiveRewards(arg0 *big.Int) (*types.Transaction, error) {
	return _Nodedaopool.Contract.ReceiveRewards(&_Nodedaopool.TransactOpts, arg0)
}

// ReceiveRewards is a paid mutator transaction binding the contract method 0xd77868da.
//
// Solidity: function receiveRewards(uint256 ) payable returns()
func (_Nodedaopool *NodedaopoolTransactorSession) ReceiveRewards(arg0 *big.Int) (*types.Transaction, error) {
	return _Nodedaopool.Contract.ReceiveRewards(&_Nodedaopool.TransactOpts, arg0)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0xe0b05b61.
//
// Solidity: function registerValidator(bytes32 _depositContractRoot, address _restakingPod, bytes[] _pubkeys, bytes[] _signatures, bytes32[] _depositDataRoots) returns()
func (_Nodedaopool *NodedaopoolTransactor) RegisterValidator(opts *bind.TransactOpts, _depositContractRoot [32]byte, _restakingPod common.Address, _pubkeys [][]byte, _signatures [][]byte, _depositDataRoots [][32]byte) (*types.Transaction, error) {
	return _Nodedaopool.contract.Transact(opts, "registerValidator", _depositContractRoot, _restakingPod, _pubkeys, _signatures, _depositDataRoots)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0xe0b05b61.
//
// Solidity: function registerValidator(bytes32 _depositContractRoot, address _restakingPod, bytes[] _pubkeys, bytes[] _signatures, bytes32[] _depositDataRoots) returns()
func (_Nodedaopool *NodedaopoolSession) RegisterValidator(_depositContractRoot [32]byte, _restakingPod common.Address, _pubkeys [][]byte, _signatures [][]byte, _depositDataRoots [][32]byte) (*types.Transaction, error) {
	return _Nodedaopool.Contract.RegisterValidator(&_Nodedaopool.TransactOpts, _depositContractRoot, _restakingPod, _pubkeys, _signatures, _depositDataRoots)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0xe0b05b61.
//
// Solidity: function registerValidator(bytes32 _depositContractRoot, address _restakingPod, bytes[] _pubkeys, bytes[] _signatures, bytes32[] _depositDataRoots) returns()
func (_Nodedaopool *NodedaopoolTransactorSession) RegisterValidator(_depositContractRoot [32]byte, _restakingPod common.Address, _pubkeys [][]byte, _signatures [][]byte, _depositDataRoots [][32]byte) (*types.Transaction, error) {
	return _Nodedaopool.Contract.RegisterValidator(&_Nodedaopool.TransactOpts, _depositContractRoot, _restakingPod, _pubkeys, _signatures, _depositDataRoots)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Nodedaopool *NodedaopoolTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nodedaopool.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Nodedaopool *NodedaopoolSession) RenounceOwnership() (*types.Transaction, error) {
	return _Nodedaopool.Contract.RenounceOwnership(&_Nodedaopool.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Nodedaopool *NodedaopoolTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Nodedaopool.Contract.RenounceOwnership(&_Nodedaopool.TransactOpts)
}

// RequestWithdrawals is a paid mutator transaction binding the contract method 0x2963a287.
//
// Solidity: function requestWithdrawals(uint256 _unstakeAmount) returns()
func (_Nodedaopool *NodedaopoolTransactor) RequestWithdrawals(opts *bind.TransactOpts, _unstakeAmount *big.Int) (*types.Transaction, error) {
	return _Nodedaopool.contract.Transact(opts, "requestWithdrawals", _unstakeAmount)
}

// RequestWithdrawals is a paid mutator transaction binding the contract method 0x2963a287.
//
// Solidity: function requestWithdrawals(uint256 _unstakeAmount) returns()
func (_Nodedaopool *NodedaopoolSession) RequestWithdrawals(_unstakeAmount *big.Int) (*types.Transaction, error) {
	return _Nodedaopool.Contract.RequestWithdrawals(&_Nodedaopool.TransactOpts, _unstakeAmount)
}

// RequestWithdrawals is a paid mutator transaction binding the contract method 0x2963a287.
//
// Solidity: function requestWithdrawals(uint256 _unstakeAmount) returns()
func (_Nodedaopool *NodedaopoolTransactorSession) RequestWithdrawals(_unstakeAmount *big.Int) (*types.Transaction, error) {
	return _Nodedaopool.Contract.RequestWithdrawals(&_Nodedaopool.TransactOpts, _unstakeAmount)
}

// SetDao is a paid mutator transaction binding the contract method 0x6637b882.
//
// Solidity: function setDao(address _dao) returns()
func (_Nodedaopool *NodedaopoolTransactor) SetDao(opts *bind.TransactOpts, _dao common.Address) (*types.Transaction, error) {
	return _Nodedaopool.contract.Transact(opts, "setDao", _dao)
}

// SetDao is a paid mutator transaction binding the contract method 0x6637b882.
//
// Solidity: function setDao(address _dao) returns()
func (_Nodedaopool *NodedaopoolSession) SetDao(_dao common.Address) (*types.Transaction, error) {
	return _Nodedaopool.Contract.SetDao(&_Nodedaopool.TransactOpts, _dao)
}

// SetDao is a paid mutator transaction binding the contract method 0x6637b882.
//
// Solidity: function setDao(address _dao) returns()
func (_Nodedaopool *NodedaopoolTransactorSession) SetDao(_dao common.Address) (*types.Transaction, error) {
	return _Nodedaopool.Contract.SetDao(&_Nodedaopool.TransactOpts, _dao)
}

// SetRateManager is a paid mutator transaction binding the contract method 0xc49db0cd.
//
// Solidity: function setRateManager(address _rateManager) returns()
func (_Nodedaopool *NodedaopoolTransactor) SetRateManager(opts *bind.TransactOpts, _rateManager common.Address) (*types.Transaction, error) {
	return _Nodedaopool.contract.Transact(opts, "setRateManager", _rateManager)
}

// SetRateManager is a paid mutator transaction binding the contract method 0xc49db0cd.
//
// Solidity: function setRateManager(address _rateManager) returns()
func (_Nodedaopool *NodedaopoolSession) SetRateManager(_rateManager common.Address) (*types.Transaction, error) {
	return _Nodedaopool.Contract.SetRateManager(&_Nodedaopool.TransactOpts, _rateManager)
}

// SetRateManager is a paid mutator transaction binding the contract method 0xc49db0cd.
//
// Solidity: function setRateManager(address _rateManager) returns()
func (_Nodedaopool *NodedaopoolTransactorSession) SetRateManager(_rateManager common.Address) (*types.Transaction, error) {
	return _Nodedaopool.Contract.SetRateManager(&_Nodedaopool.TransactOpts, _rateManager)
}

// SetStrategyVault is a paid mutator transaction binding the contract method 0xe606768a.
//
// Solidity: function setStrategyVault(address _strategyVault) returns()
func (_Nodedaopool *NodedaopoolTransactor) SetStrategyVault(opts *bind.TransactOpts, _strategyVault common.Address) (*types.Transaction, error) {
	return _Nodedaopool.contract.Transact(opts, "setStrategyVault", _strategyVault)
}

// SetStrategyVault is a paid mutator transaction binding the contract method 0xe606768a.
//
// Solidity: function setStrategyVault(address _strategyVault) returns()
func (_Nodedaopool *NodedaopoolSession) SetStrategyVault(_strategyVault common.Address) (*types.Transaction, error) {
	return _Nodedaopool.Contract.SetStrategyVault(&_Nodedaopool.TransactOpts, _strategyVault)
}

// SetStrategyVault is a paid mutator transaction binding the contract method 0xe606768a.
//
// Solidity: function setStrategyVault(address _strategyVault) returns()
func (_Nodedaopool *NodedaopoolTransactorSession) SetStrategyVault(_strategyVault common.Address) (*types.Transaction, error) {
	return _Nodedaopool.Contract.SetStrategyVault(&_Nodedaopool.TransactOpts, _strategyVault)
}

// SetUnstakeAllowed is a paid mutator transaction binding the contract method 0xbba59090.
//
// Solidity: function setUnstakeAllowed(bool _unstakeAllowed) returns()
func (_Nodedaopool *NodedaopoolTransactor) SetUnstakeAllowed(opts *bind.TransactOpts, _unstakeAllowed bool) (*types.Transaction, error) {
	return _Nodedaopool.contract.Transact(opts, "setUnstakeAllowed", _unstakeAllowed)
}

// SetUnstakeAllowed is a paid mutator transaction binding the contract method 0xbba59090.
//
// Solidity: function setUnstakeAllowed(bool _unstakeAllowed) returns()
func (_Nodedaopool *NodedaopoolSession) SetUnstakeAllowed(_unstakeAllowed bool) (*types.Transaction, error) {
	return _Nodedaopool.Contract.SetUnstakeAllowed(&_Nodedaopool.TransactOpts, _unstakeAllowed)
}

// SetUnstakeAllowed is a paid mutator transaction binding the contract method 0xbba59090.
//
// Solidity: function setUnstakeAllowed(bool _unstakeAllowed) returns()
func (_Nodedaopool *NodedaopoolTransactorSession) SetUnstakeAllowed(_unstakeAllowed bool) (*types.Transaction, error) {
	return _Nodedaopool.Contract.SetUnstakeAllowed(&_Nodedaopool.TransactOpts, _unstakeAllowed)
}

// SetValidatorManager is a paid mutator transaction binding the contract method 0x45f34e92.
//
// Solidity: function setValidatorManager(address _validatorManager) returns()
func (_Nodedaopool *NodedaopoolTransactor) SetValidatorManager(opts *bind.TransactOpts, _validatorManager common.Address) (*types.Transaction, error) {
	return _Nodedaopool.contract.Transact(opts, "setValidatorManager", _validatorManager)
}

// SetValidatorManager is a paid mutator transaction binding the contract method 0x45f34e92.
//
// Solidity: function setValidatorManager(address _validatorManager) returns()
func (_Nodedaopool *NodedaopoolSession) SetValidatorManager(_validatorManager common.Address) (*types.Transaction, error) {
	return _Nodedaopool.Contract.SetValidatorManager(&_Nodedaopool.TransactOpts, _validatorManager)
}

// SetValidatorManager is a paid mutator transaction binding the contract method 0x45f34e92.
//
// Solidity: function setValidatorManager(address _validatorManager) returns()
func (_Nodedaopool *NodedaopoolTransactorSession) SetValidatorManager(_validatorManager common.Address) (*types.Transaction, error) {
	return _Nodedaopool.Contract.SetValidatorManager(&_Nodedaopool.TransactOpts, _validatorManager)
}

// SetValidatorRegistry is a paid mutator transaction binding the contract method 0x49773050.
//
// Solidity: function setValidatorRegistry(address _validatorRegistry) returns()
func (_Nodedaopool *NodedaopoolTransactor) SetValidatorRegistry(opts *bind.TransactOpts, _validatorRegistry common.Address) (*types.Transaction, error) {
	return _Nodedaopool.contract.Transact(opts, "setValidatorRegistry", _validatorRegistry)
}

// SetValidatorRegistry is a paid mutator transaction binding the contract method 0x49773050.
//
// Solidity: function setValidatorRegistry(address _validatorRegistry) returns()
func (_Nodedaopool *NodedaopoolSession) SetValidatorRegistry(_validatorRegistry common.Address) (*types.Transaction, error) {
	return _Nodedaopool.Contract.SetValidatorRegistry(&_Nodedaopool.TransactOpts, _validatorRegistry)
}

// SetValidatorRegistry is a paid mutator transaction binding the contract method 0x49773050.
//
// Solidity: function setValidatorRegistry(address _validatorRegistry) returns()
func (_Nodedaopool *NodedaopoolTransactorSession) SetValidatorRegistry(_validatorRegistry common.Address) (*types.Transaction, error) {
	return _Nodedaopool.Contract.SetValidatorRegistry(&_Nodedaopool.TransactOpts, _validatorRegistry)
}

// SetWithdrawalDelayBlocks is a paid mutator transaction binding the contract method 0x4d50f9a4.
//
// Solidity: function setWithdrawalDelayBlocks(uint256 _withdrawalDelayBlocks) returns()
func (_Nodedaopool *NodedaopoolTransactor) SetWithdrawalDelayBlocks(opts *bind.TransactOpts, _withdrawalDelayBlocks *big.Int) (*types.Transaction, error) {
	return _Nodedaopool.contract.Transact(opts, "setWithdrawalDelayBlocks", _withdrawalDelayBlocks)
}

// SetWithdrawalDelayBlocks is a paid mutator transaction binding the contract method 0x4d50f9a4.
//
// Solidity: function setWithdrawalDelayBlocks(uint256 _withdrawalDelayBlocks) returns()
func (_Nodedaopool *NodedaopoolSession) SetWithdrawalDelayBlocks(_withdrawalDelayBlocks *big.Int) (*types.Transaction, error) {
	return _Nodedaopool.Contract.SetWithdrawalDelayBlocks(&_Nodedaopool.TransactOpts, _withdrawalDelayBlocks)
}

// SetWithdrawalDelayBlocks is a paid mutator transaction binding the contract method 0x4d50f9a4.
//
// Solidity: function setWithdrawalDelayBlocks(uint256 _withdrawalDelayBlocks) returns()
func (_Nodedaopool *NodedaopoolTransactorSession) SetWithdrawalDelayBlocks(_withdrawalDelayBlocks *big.Int) (*types.Transaction, error) {
	return _Nodedaopool.Contract.SetWithdrawalDelayBlocks(&_Nodedaopool.TransactOpts, _withdrawalDelayBlocks)
}

// StakeETH is a paid mutator transaction binding the contract method 0xdceb986d.
//
// Solidity: function stakeETH() payable returns()
func (_Nodedaopool *NodedaopoolTransactor) StakeETH(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nodedaopool.contract.Transact(opts, "stakeETH")
}

// StakeETH is a paid mutator transaction binding the contract method 0xdceb986d.
//
// Solidity: function stakeETH() payable returns()
func (_Nodedaopool *NodedaopoolSession) StakeETH() (*types.Transaction, error) {
	return _Nodedaopool.Contract.StakeETH(&_Nodedaopool.TransactOpts)
}

// StakeETH is a paid mutator transaction binding the contract method 0xdceb986d.
//
// Solidity: function stakeETH() payable returns()
func (_Nodedaopool *NodedaopoolTransactorSession) StakeETH() (*types.Transaction, error) {
	return _Nodedaopool.Contract.StakeETH(&_Nodedaopool.TransactOpts)
}

// StrategyDeposit is a paid mutator transaction binding the contract method 0xdce590d5.
//
// Solidity: function strategyDeposit(uint256 _amount) returns()
func (_Nodedaopool *NodedaopoolTransactor) StrategyDeposit(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Nodedaopool.contract.Transact(opts, "strategyDeposit", _amount)
}

// StrategyDeposit is a paid mutator transaction binding the contract method 0xdce590d5.
//
// Solidity: function strategyDeposit(uint256 _amount) returns()
func (_Nodedaopool *NodedaopoolSession) StrategyDeposit(_amount *big.Int) (*types.Transaction, error) {
	return _Nodedaopool.Contract.StrategyDeposit(&_Nodedaopool.TransactOpts, _amount)
}

// StrategyDeposit is a paid mutator transaction binding the contract method 0xdce590d5.
//
// Solidity: function strategyDeposit(uint256 _amount) returns()
func (_Nodedaopool *NodedaopoolTransactorSession) StrategyDeposit(_amount *big.Int) (*types.Transaction, error) {
	return _Nodedaopool.Contract.StrategyDeposit(&_Nodedaopool.TransactOpts, _amount)
}

// StrategyReturn is a paid mutator transaction binding the contract method 0x553ffcbe.
//
// Solidity: function strategyReturn() payable returns()
func (_Nodedaopool *NodedaopoolTransactor) StrategyReturn(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nodedaopool.contract.Transact(opts, "strategyReturn")
}

// StrategyReturn is a paid mutator transaction binding the contract method 0x553ffcbe.
//
// Solidity: function strategyReturn() payable returns()
func (_Nodedaopool *NodedaopoolSession) StrategyReturn() (*types.Transaction, error) {
	return _Nodedaopool.Contract.StrategyReturn(&_Nodedaopool.TransactOpts)
}

// StrategyReturn is a paid mutator transaction binding the contract method 0x553ffcbe.
//
// Solidity: function strategyReturn() payable returns()
func (_Nodedaopool *NodedaopoolTransactorSession) StrategyReturn() (*types.Transaction, error) {
	return _Nodedaopool.Contract.StrategyReturn(&_Nodedaopool.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Nodedaopool *NodedaopoolTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Nodedaopool.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Nodedaopool *NodedaopoolSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Nodedaopool.Contract.TransferOwnership(&_Nodedaopool.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Nodedaopool *NodedaopoolTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Nodedaopool.Contract.TransferOwnership(&_Nodedaopool.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Nodedaopool *NodedaopoolTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nodedaopool.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Nodedaopool *NodedaopoolSession) Unpause() (*types.Transaction, error) {
	return _Nodedaopool.Contract.Unpause(&_Nodedaopool.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Nodedaopool *NodedaopoolTransactorSession) Unpause() (*types.Transaction, error) {
	return _Nodedaopool.Contract.Unpause(&_Nodedaopool.TransactOpts)
}

// UnstakeETH is a paid mutator transaction binding the contract method 0x62d53403.
//
// Solidity: function unstakeETH(uint256 _unstakeAmount) returns()
func (_Nodedaopool *NodedaopoolTransactor) UnstakeETH(opts *bind.TransactOpts, _unstakeAmount *big.Int) (*types.Transaction, error) {
	return _Nodedaopool.contract.Transact(opts, "unstakeETH", _unstakeAmount)
}

// UnstakeETH is a paid mutator transaction binding the contract method 0x62d53403.
//
// Solidity: function unstakeETH(uint256 _unstakeAmount) returns()
func (_Nodedaopool *NodedaopoolSession) UnstakeETH(_unstakeAmount *big.Int) (*types.Transaction, error) {
	return _Nodedaopool.Contract.UnstakeETH(&_Nodedaopool.TransactOpts, _unstakeAmount)
}

// UnstakeETH is a paid mutator transaction binding the contract method 0x62d53403.
//
// Solidity: function unstakeETH(uint256 _unstakeAmount) returns()
func (_Nodedaopool *NodedaopoolTransactorSession) UnstakeETH(_unstakeAmount *big.Int) (*types.Transaction, error) {
	return _Nodedaopool.Contract.UnstakeETH(&_Nodedaopool.TransactOpts, _unstakeAmount)
}

// UpdateApr is a paid mutator transaction binding the contract method 0x8552bf90.
//
// Solidity: function updateApr(uint256 _apr) returns()
func (_Nodedaopool *NodedaopoolTransactor) UpdateApr(opts *bind.TransactOpts, _apr *big.Int) (*types.Transaction, error) {
	return _Nodedaopool.contract.Transact(opts, "updateApr", _apr)
}

// UpdateApr is a paid mutator transaction binding the contract method 0x8552bf90.
//
// Solidity: function updateApr(uint256 _apr) returns()
func (_Nodedaopool *NodedaopoolSession) UpdateApr(_apr *big.Int) (*types.Transaction, error) {
	return _Nodedaopool.Contract.UpdateApr(&_Nodedaopool.TransactOpts, _apr)
}

// UpdateApr is a paid mutator transaction binding the contract method 0x8552bf90.
//
// Solidity: function updateApr(uint256 _apr) returns()
func (_Nodedaopool *NodedaopoolTransactorSession) UpdateApr(_apr *big.Int) (*types.Transaction, error) {
	return _Nodedaopool.Contract.UpdateApr(&_Nodedaopool.TransactOpts, _apr)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Nodedaopool *NodedaopoolTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Nodedaopool.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Nodedaopool *NodedaopoolSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Nodedaopool.Contract.UpgradeTo(&_Nodedaopool.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Nodedaopool *NodedaopoolTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Nodedaopool.Contract.UpgradeTo(&_Nodedaopool.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Nodedaopool *NodedaopoolTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Nodedaopool.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Nodedaopool *NodedaopoolSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Nodedaopool.Contract.UpgradeToAndCall(&_Nodedaopool.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Nodedaopool *NodedaopoolTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Nodedaopool.Contract.UpgradeToAndCall(&_Nodedaopool.TransactOpts, newImplementation, data)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Nodedaopool *NodedaopoolTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nodedaopool.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Nodedaopool *NodedaopoolSession) Receive() (*types.Transaction, error) {
	return _Nodedaopool.Contract.Receive(&_Nodedaopool.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Nodedaopool *NodedaopoolTransactorSession) Receive() (*types.Transaction, error) {
	return _Nodedaopool.Contract.Receive(&_Nodedaopool.TransactOpts)
}

// NodedaopoolAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Nodedaopool contract.
type NodedaopoolAdminChangedIterator struct {
	Event *NodedaopoolAdminChanged // Event containing the contract specifics and raw log

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
func (it *NodedaopoolAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodedaopoolAdminChanged)
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
		it.Event = new(NodedaopoolAdminChanged)
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
func (it *NodedaopoolAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodedaopoolAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodedaopoolAdminChanged represents a AdminChanged event raised by the Nodedaopool contract.
type NodedaopoolAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Nodedaopool *NodedaopoolFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*NodedaopoolAdminChangedIterator, error) {

	logs, sub, err := _Nodedaopool.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &NodedaopoolAdminChangedIterator{contract: _Nodedaopool.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Nodedaopool *NodedaopoolFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *NodedaopoolAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Nodedaopool.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodedaopoolAdminChanged)
				if err := _Nodedaopool.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_Nodedaopool *NodedaopoolFilterer) ParseAdminChanged(log types.Log) (*NodedaopoolAdminChanged, error) {
	event := new(NodedaopoolAdminChanged)
	if err := _Nodedaopool.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodedaopoolAprUpdatedIterator is returned from FilterAprUpdated and is used to iterate over the raw logs and unpacked data for AprUpdated events raised by the Nodedaopool contract.
type NodedaopoolAprUpdatedIterator struct {
	Event *NodedaopoolAprUpdated // Event containing the contract specifics and raw log

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
func (it *NodedaopoolAprUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodedaopoolAprUpdated)
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
		it.Event = new(NodedaopoolAprUpdated)
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
func (it *NodedaopoolAprUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodedaopoolAprUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodedaopoolAprUpdated represents a AprUpdated event raised by the Nodedaopool contract.
type NodedaopoolAprUpdated struct {
	OldApr *big.Int
	Apr    *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAprUpdated is a free log retrieval operation binding the contract event 0x782f84f1274a11befd10700001e41dfdbf825313fb93311d474b857dfd8f1c2b.
//
// Solidity: event AprUpdated(uint256 _oldApr, uint256 _apr)
func (_Nodedaopool *NodedaopoolFilterer) FilterAprUpdated(opts *bind.FilterOpts) (*NodedaopoolAprUpdatedIterator, error) {

	logs, sub, err := _Nodedaopool.contract.FilterLogs(opts, "AprUpdated")
	if err != nil {
		return nil, err
	}
	return &NodedaopoolAprUpdatedIterator{contract: _Nodedaopool.contract, event: "AprUpdated", logs: logs, sub: sub}, nil
}

// WatchAprUpdated is a free log subscription operation binding the contract event 0x782f84f1274a11befd10700001e41dfdbf825313fb93311d474b857dfd8f1c2b.
//
// Solidity: event AprUpdated(uint256 _oldApr, uint256 _apr)
func (_Nodedaopool *NodedaopoolFilterer) WatchAprUpdated(opts *bind.WatchOpts, sink chan<- *NodedaopoolAprUpdated) (event.Subscription, error) {

	logs, sub, err := _Nodedaopool.contract.WatchLogs(opts, "AprUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodedaopoolAprUpdated)
				if err := _Nodedaopool.contract.UnpackLog(event, "AprUpdated", log); err != nil {
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
func (_Nodedaopool *NodedaopoolFilterer) ParseAprUpdated(log types.Log) (*NodedaopoolAprUpdated, error) {
	event := new(NodedaopoolAprUpdated)
	if err := _Nodedaopool.contract.UnpackLog(event, "AprUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodedaopoolAssetsUpdatedIterator is returned from FilterAssetsUpdated and is used to iterate over the raw logs and unpacked data for AssetsUpdated events raised by the Nodedaopool contract.
type NodedaopoolAssetsUpdatedIterator struct {
	Event *NodedaopoolAssetsUpdated // Event containing the contract specifics and raw log

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
func (it *NodedaopoolAssetsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodedaopoolAssetsUpdated)
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
		it.Event = new(NodedaopoolAssetsUpdated)
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
func (it *NodedaopoolAssetsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodedaopoolAssetsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodedaopoolAssetsUpdated represents a AssetsUpdated event raised by the Nodedaopool contract.
type NodedaopoolAssetsUpdated struct {
	TotalUnderlyingAsset *big.Int
	EstimatedRewards     *big.Int
	BlockNumber          *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterAssetsUpdated is a free log retrieval operation binding the contract event 0x35a901c4413e585f9121eb5cf07e67760bd4ac498dd031249e5cd2cd225f74e4.
//
// Solidity: event AssetsUpdated(uint256 _totalUnderlyingAsset, uint256 _estimatedRewards, uint256 _blockNumber)
func (_Nodedaopool *NodedaopoolFilterer) FilterAssetsUpdated(opts *bind.FilterOpts) (*NodedaopoolAssetsUpdatedIterator, error) {

	logs, sub, err := _Nodedaopool.contract.FilterLogs(opts, "AssetsUpdated")
	if err != nil {
		return nil, err
	}
	return &NodedaopoolAssetsUpdatedIterator{contract: _Nodedaopool.contract, event: "AssetsUpdated", logs: logs, sub: sub}, nil
}

// WatchAssetsUpdated is a free log subscription operation binding the contract event 0x35a901c4413e585f9121eb5cf07e67760bd4ac498dd031249e5cd2cd225f74e4.
//
// Solidity: event AssetsUpdated(uint256 _totalUnderlyingAsset, uint256 _estimatedRewards, uint256 _blockNumber)
func (_Nodedaopool *NodedaopoolFilterer) WatchAssetsUpdated(opts *bind.WatchOpts, sink chan<- *NodedaopoolAssetsUpdated) (event.Subscription, error) {

	logs, sub, err := _Nodedaopool.contract.WatchLogs(opts, "AssetsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodedaopoolAssetsUpdated)
				if err := _Nodedaopool.contract.UnpackLog(event, "AssetsUpdated", log); err != nil {
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
func (_Nodedaopool *NodedaopoolFilterer) ParseAssetsUpdated(log types.Log) (*NodedaopoolAssetsUpdated, error) {
	event := new(NodedaopoolAssetsUpdated)
	if err := _Nodedaopool.contract.UnpackLog(event, "AssetsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodedaopoolBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the Nodedaopool contract.
type NodedaopoolBeaconUpgradedIterator struct {
	Event *NodedaopoolBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *NodedaopoolBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodedaopoolBeaconUpgraded)
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
		it.Event = new(NodedaopoolBeaconUpgraded)
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
func (it *NodedaopoolBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodedaopoolBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodedaopoolBeaconUpgraded represents a BeaconUpgraded event raised by the Nodedaopool contract.
type NodedaopoolBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Nodedaopool *NodedaopoolFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*NodedaopoolBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Nodedaopool.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &NodedaopoolBeaconUpgradedIterator{contract: _Nodedaopool.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Nodedaopool *NodedaopoolFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *NodedaopoolBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Nodedaopool.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodedaopoolBeaconUpgraded)
				if err := _Nodedaopool.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_Nodedaopool *NodedaopoolFilterer) ParseBeaconUpgraded(log types.Log) (*NodedaopoolBeaconUpgraded, error) {
	event := new(NodedaopoolBeaconUpgraded)
	if err := _Nodedaopool.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodedaopoolDaoChangedIterator is returned from FilterDaoChanged and is used to iterate over the raw logs and unpacked data for DaoChanged events raised by the Nodedaopool contract.
type NodedaopoolDaoChangedIterator struct {
	Event *NodedaopoolDaoChanged // Event containing the contract specifics and raw log

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
func (it *NodedaopoolDaoChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodedaopoolDaoChanged)
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
		it.Event = new(NodedaopoolDaoChanged)
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
func (it *NodedaopoolDaoChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodedaopoolDaoChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodedaopoolDaoChanged represents a DaoChanged event raised by the Nodedaopool contract.
type NodedaopoolDaoChanged struct {
	OldDao common.Address
	Dao    common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDaoChanged is a free log retrieval operation binding the contract event 0xfcde6c827a52b0870bc44ed9b10212272e18c9ea1725b772e9b493750afd8da4.
//
// Solidity: event DaoChanged(address _oldDao, address _dao)
func (_Nodedaopool *NodedaopoolFilterer) FilterDaoChanged(opts *bind.FilterOpts) (*NodedaopoolDaoChangedIterator, error) {

	logs, sub, err := _Nodedaopool.contract.FilterLogs(opts, "DaoChanged")
	if err != nil {
		return nil, err
	}
	return &NodedaopoolDaoChangedIterator{contract: _Nodedaopool.contract, event: "DaoChanged", logs: logs, sub: sub}, nil
}

// WatchDaoChanged is a free log subscription operation binding the contract event 0xfcde6c827a52b0870bc44ed9b10212272e18c9ea1725b772e9b493750afd8da4.
//
// Solidity: event DaoChanged(address _oldDao, address _dao)
func (_Nodedaopool *NodedaopoolFilterer) WatchDaoChanged(opts *bind.WatchOpts, sink chan<- *NodedaopoolDaoChanged) (event.Subscription, error) {

	logs, sub, err := _Nodedaopool.contract.WatchLogs(opts, "DaoChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodedaopoolDaoChanged)
				if err := _Nodedaopool.contract.UnpackLog(event, "DaoChanged", log); err != nil {
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
func (_Nodedaopool *NodedaopoolFilterer) ParseDaoChanged(log types.Log) (*NodedaopoolDaoChanged, error) {
	event := new(NodedaopoolDaoChanged)
	if err := _Nodedaopool.contract.UnpackLog(event, "DaoChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodedaopoolEthStakeIterator is returned from FilterEthStake and is used to iterate over the raw logs and unpacked data for EthStake events raised by the Nodedaopool contract.
type NodedaopoolEthStakeIterator struct {
	Event *NodedaopoolEthStake // Event containing the contract specifics and raw log

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
func (it *NodedaopoolEthStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodedaopoolEthStake)
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
		it.Event = new(NodedaopoolEthStake)
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
func (it *NodedaopoolEthStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodedaopoolEthStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodedaopoolEthStake represents a EthStake event raised by the Nodedaopool contract.
type NodedaopoolEthStake struct {
	Staker      common.Address
	StakeAmount *big.Int
	MintAmount  *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterEthStake is a free log retrieval operation binding the contract event 0x838d17987e57e587c458220b9b38723c41fbc3f397550b506712960a73ef19f9.
//
// Solidity: event EthStake(address _staker, uint256 _stakeAmount, uint256 _mintAmount)
func (_Nodedaopool *NodedaopoolFilterer) FilterEthStake(opts *bind.FilterOpts) (*NodedaopoolEthStakeIterator, error) {

	logs, sub, err := _Nodedaopool.contract.FilterLogs(opts, "EthStake")
	if err != nil {
		return nil, err
	}
	return &NodedaopoolEthStakeIterator{contract: _Nodedaopool.contract, event: "EthStake", logs: logs, sub: sub}, nil
}

// WatchEthStake is a free log subscription operation binding the contract event 0x838d17987e57e587c458220b9b38723c41fbc3f397550b506712960a73ef19f9.
//
// Solidity: event EthStake(address _staker, uint256 _stakeAmount, uint256 _mintAmount)
func (_Nodedaopool *NodedaopoolFilterer) WatchEthStake(opts *bind.WatchOpts, sink chan<- *NodedaopoolEthStake) (event.Subscription, error) {

	logs, sub, err := _Nodedaopool.contract.WatchLogs(opts, "EthStake")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodedaopoolEthStake)
				if err := _Nodedaopool.contract.UnpackLog(event, "EthStake", log); err != nil {
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
func (_Nodedaopool *NodedaopoolFilterer) ParseEthStake(log types.Log) (*NodedaopoolEthStake, error) {
	event := new(NodedaopoolEthStake)
	if err := _Nodedaopool.contract.UnpackLog(event, "EthStake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodedaopoolEthUnstakeIterator is returned from FilterEthUnstake and is used to iterate over the raw logs and unpacked data for EthUnstake events raised by the Nodedaopool contract.
type NodedaopoolEthUnstakeIterator struct {
	Event *NodedaopoolEthUnstake // Event containing the contract specifics and raw log

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
func (it *NodedaopoolEthUnstakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodedaopoolEthUnstake)
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
		it.Event = new(NodedaopoolEthUnstake)
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
func (it *NodedaopoolEthUnstakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodedaopoolEthUnstakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodedaopoolEthUnstake represents a EthUnstake event raised by the Nodedaopool contract.
type NodedaopoolEthUnstake struct {
	Sender        common.Address
	UnstakeAmount *big.Int
	EthAmount     *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterEthUnstake is a free log retrieval operation binding the contract event 0xdef4a00a06705bb4b80fd0c912337f4f4e01fb81d39a572514be68d328599abf.
//
// Solidity: event EthUnstake(address _sender, uint256 _unstakeAmount, uint256 _ethAmount)
func (_Nodedaopool *NodedaopoolFilterer) FilterEthUnstake(opts *bind.FilterOpts) (*NodedaopoolEthUnstakeIterator, error) {

	logs, sub, err := _Nodedaopool.contract.FilterLogs(opts, "EthUnstake")
	if err != nil {
		return nil, err
	}
	return &NodedaopoolEthUnstakeIterator{contract: _Nodedaopool.contract, event: "EthUnstake", logs: logs, sub: sub}, nil
}

// WatchEthUnstake is a free log subscription operation binding the contract event 0xdef4a00a06705bb4b80fd0c912337f4f4e01fb81d39a572514be68d328599abf.
//
// Solidity: event EthUnstake(address _sender, uint256 _unstakeAmount, uint256 _ethAmount)
func (_Nodedaopool *NodedaopoolFilterer) WatchEthUnstake(opts *bind.WatchOpts, sink chan<- *NodedaopoolEthUnstake) (event.Subscription, error) {

	logs, sub, err := _Nodedaopool.contract.WatchLogs(opts, "EthUnstake")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodedaopoolEthUnstake)
				if err := _Nodedaopool.contract.UnpackLog(event, "EthUnstake", log); err != nil {
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
func (_Nodedaopool *NodedaopoolFilterer) ParseEthUnstake(log types.Log) (*NodedaopoolEthUnstake, error) {
	event := new(NodedaopoolEthUnstake)
	if err := _Nodedaopool.contract.UnpackLog(event, "EthUnstake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodedaopoolInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Nodedaopool contract.
type NodedaopoolInitializedIterator struct {
	Event *NodedaopoolInitialized // Event containing the contract specifics and raw log

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
func (it *NodedaopoolInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodedaopoolInitialized)
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
		it.Event = new(NodedaopoolInitialized)
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
func (it *NodedaopoolInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodedaopoolInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodedaopoolInitialized represents a Initialized event raised by the Nodedaopool contract.
type NodedaopoolInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Nodedaopool *NodedaopoolFilterer) FilterInitialized(opts *bind.FilterOpts) (*NodedaopoolInitializedIterator, error) {

	logs, sub, err := _Nodedaopool.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &NodedaopoolInitializedIterator{contract: _Nodedaopool.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Nodedaopool *NodedaopoolFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *NodedaopoolInitialized) (event.Subscription, error) {

	logs, sub, err := _Nodedaopool.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodedaopoolInitialized)
				if err := _Nodedaopool.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Nodedaopool *NodedaopoolFilterer) ParseInitialized(log types.Log) (*NodedaopoolInitialized, error) {
	event := new(NodedaopoolInitialized)
	if err := _Nodedaopool.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodedaopoolOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Nodedaopool contract.
type NodedaopoolOwnershipTransferredIterator struct {
	Event *NodedaopoolOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *NodedaopoolOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodedaopoolOwnershipTransferred)
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
		it.Event = new(NodedaopoolOwnershipTransferred)
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
func (it *NodedaopoolOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodedaopoolOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodedaopoolOwnershipTransferred represents a OwnershipTransferred event raised by the Nodedaopool contract.
type NodedaopoolOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Nodedaopool *NodedaopoolFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*NodedaopoolOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Nodedaopool.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NodedaopoolOwnershipTransferredIterator{contract: _Nodedaopool.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Nodedaopool *NodedaopoolFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NodedaopoolOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Nodedaopool.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodedaopoolOwnershipTransferred)
				if err := _Nodedaopool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Nodedaopool *NodedaopoolFilterer) ParseOwnershipTransferred(log types.Log) (*NodedaopoolOwnershipTransferred, error) {
	event := new(NodedaopoolOwnershipTransferred)
	if err := _Nodedaopool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodedaopoolPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Nodedaopool contract.
type NodedaopoolPausedIterator struct {
	Event *NodedaopoolPaused // Event containing the contract specifics and raw log

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
func (it *NodedaopoolPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodedaopoolPaused)
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
		it.Event = new(NodedaopoolPaused)
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
func (it *NodedaopoolPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodedaopoolPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodedaopoolPaused represents a Paused event raised by the Nodedaopool contract.
type NodedaopoolPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Nodedaopool *NodedaopoolFilterer) FilterPaused(opts *bind.FilterOpts) (*NodedaopoolPausedIterator, error) {

	logs, sub, err := _Nodedaopool.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &NodedaopoolPausedIterator{contract: _Nodedaopool.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Nodedaopool *NodedaopoolFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *NodedaopoolPaused) (event.Subscription, error) {

	logs, sub, err := _Nodedaopool.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodedaopoolPaused)
				if err := _Nodedaopool.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Nodedaopool *NodedaopoolFilterer) ParsePaused(log types.Log) (*NodedaopoolPaused, error) {
	event := new(NodedaopoolPaused)
	if err := _Nodedaopool.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodedaopoolRateManagerChangedIterator is returned from FilterRateManagerChanged and is used to iterate over the raw logs and unpacked data for RateManagerChanged events raised by the Nodedaopool contract.
type NodedaopoolRateManagerChangedIterator struct {
	Event *NodedaopoolRateManagerChanged // Event containing the contract specifics and raw log

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
func (it *NodedaopoolRateManagerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodedaopoolRateManagerChanged)
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
		it.Event = new(NodedaopoolRateManagerChanged)
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
func (it *NodedaopoolRateManagerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodedaopoolRateManagerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodedaopoolRateManagerChanged represents a RateManagerChanged event raised by the Nodedaopool contract.
type NodedaopoolRateManagerChanged struct {
	OldAprManager common.Address
	AprManager    common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRateManagerChanged is a free log retrieval operation binding the contract event 0x6b8a1f2fd09d355f8419738a3593f646cbd0e7be553ffcce23694ce968cc6425.
//
// Solidity: event RateManagerChanged(address _oldAprManager, address _aprManager)
func (_Nodedaopool *NodedaopoolFilterer) FilterRateManagerChanged(opts *bind.FilterOpts) (*NodedaopoolRateManagerChangedIterator, error) {

	logs, sub, err := _Nodedaopool.contract.FilterLogs(opts, "RateManagerChanged")
	if err != nil {
		return nil, err
	}
	return &NodedaopoolRateManagerChangedIterator{contract: _Nodedaopool.contract, event: "RateManagerChanged", logs: logs, sub: sub}, nil
}

// WatchRateManagerChanged is a free log subscription operation binding the contract event 0x6b8a1f2fd09d355f8419738a3593f646cbd0e7be553ffcce23694ce968cc6425.
//
// Solidity: event RateManagerChanged(address _oldAprManager, address _aprManager)
func (_Nodedaopool *NodedaopoolFilterer) WatchRateManagerChanged(opts *bind.WatchOpts, sink chan<- *NodedaopoolRateManagerChanged) (event.Subscription, error) {

	logs, sub, err := _Nodedaopool.contract.WatchLogs(opts, "RateManagerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodedaopoolRateManagerChanged)
				if err := _Nodedaopool.contract.UnpackLog(event, "RateManagerChanged", log); err != nil {
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
func (_Nodedaopool *NodedaopoolFilterer) ParseRateManagerChanged(log types.Log) (*NodedaopoolRateManagerChanged, error) {
	event := new(NodedaopoolRateManagerChanged)
	if err := _Nodedaopool.contract.UnpackLog(event, "RateManagerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodedaopoolReceivedIterator is returned from FilterReceived and is used to iterate over the raw logs and unpacked data for Received events raised by the Nodedaopool contract.
type NodedaopoolReceivedIterator struct {
	Event *NodedaopoolReceived // Event containing the contract specifics and raw log

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
func (it *NodedaopoolReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodedaopoolReceived)
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
		it.Event = new(NodedaopoolReceived)
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
func (it *NodedaopoolReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodedaopoolReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodedaopoolReceived represents a Received event raised by the Nodedaopool contract.
type NodedaopoolReceived struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReceived is a free log retrieval operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address _sender, uint256 _amount)
func (_Nodedaopool *NodedaopoolFilterer) FilterReceived(opts *bind.FilterOpts) (*NodedaopoolReceivedIterator, error) {

	logs, sub, err := _Nodedaopool.contract.FilterLogs(opts, "Received")
	if err != nil {
		return nil, err
	}
	return &NodedaopoolReceivedIterator{contract: _Nodedaopool.contract, event: "Received", logs: logs, sub: sub}, nil
}

// WatchReceived is a free log subscription operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address _sender, uint256 _amount)
func (_Nodedaopool *NodedaopoolFilterer) WatchReceived(opts *bind.WatchOpts, sink chan<- *NodedaopoolReceived) (event.Subscription, error) {

	logs, sub, err := _Nodedaopool.contract.WatchLogs(opts, "Received")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodedaopoolReceived)
				if err := _Nodedaopool.contract.UnpackLog(event, "Received", log); err != nil {
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
func (_Nodedaopool *NodedaopoolFilterer) ParseReceived(log types.Log) (*NodedaopoolReceived, error) {
	event := new(NodedaopoolReceived)
	if err := _Nodedaopool.contract.UnpackLog(event, "Received", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodedaopoolRestakingPodAddedIterator is returned from FilterRestakingPodAdded and is used to iterate over the raw logs and unpacked data for RestakingPodAdded events raised by the Nodedaopool contract.
type NodedaopoolRestakingPodAddedIterator struct {
	Event *NodedaopoolRestakingPodAdded // Event containing the contract specifics and raw log

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
func (it *NodedaopoolRestakingPodAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodedaopoolRestakingPodAdded)
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
		it.Event = new(NodedaopoolRestakingPodAdded)
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
func (it *NodedaopoolRestakingPodAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodedaopoolRestakingPodAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodedaopoolRestakingPodAdded represents a RestakingPodAdded event raised by the Nodedaopool contract.
type NodedaopoolRestakingPodAdded struct {
	RestakingPod common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRestakingPodAdded is a free log retrieval operation binding the contract event 0x84cf405f0a1827114df681d11bd5006d56d6414d63c57e4cf6569fabd9a8e7ff.
//
// Solidity: event RestakingPodAdded(address _restakingPod)
func (_Nodedaopool *NodedaopoolFilterer) FilterRestakingPodAdded(opts *bind.FilterOpts) (*NodedaopoolRestakingPodAddedIterator, error) {

	logs, sub, err := _Nodedaopool.contract.FilterLogs(opts, "RestakingPodAdded")
	if err != nil {
		return nil, err
	}
	return &NodedaopoolRestakingPodAddedIterator{contract: _Nodedaopool.contract, event: "RestakingPodAdded", logs: logs, sub: sub}, nil
}

// WatchRestakingPodAdded is a free log subscription operation binding the contract event 0x84cf405f0a1827114df681d11bd5006d56d6414d63c57e4cf6569fabd9a8e7ff.
//
// Solidity: event RestakingPodAdded(address _restakingPod)
func (_Nodedaopool *NodedaopoolFilterer) WatchRestakingPodAdded(opts *bind.WatchOpts, sink chan<- *NodedaopoolRestakingPodAdded) (event.Subscription, error) {

	logs, sub, err := _Nodedaopool.contract.WatchLogs(opts, "RestakingPodAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodedaopoolRestakingPodAdded)
				if err := _Nodedaopool.contract.UnpackLog(event, "RestakingPodAdded", log); err != nil {
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
func (_Nodedaopool *NodedaopoolFilterer) ParseRestakingPodAdded(log types.Log) (*NodedaopoolRestakingPodAdded, error) {
	event := new(NodedaopoolRestakingPodAdded)
	if err := _Nodedaopool.contract.UnpackLog(event, "RestakingPodAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodedaopoolStrategyDepositedIterator is returned from FilterStrategyDeposited and is used to iterate over the raw logs and unpacked data for StrategyDeposited events raised by the Nodedaopool contract.
type NodedaopoolStrategyDepositedIterator struct {
	Event *NodedaopoolStrategyDeposited // Event containing the contract specifics and raw log

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
func (it *NodedaopoolStrategyDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodedaopoolStrategyDeposited)
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
		it.Event = new(NodedaopoolStrategyDeposited)
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
func (it *NodedaopoolStrategyDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodedaopoolStrategyDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodedaopoolStrategyDeposited represents a StrategyDeposited event raised by the Nodedaopool contract.
type NodedaopoolStrategyDeposited struct {
	Amount         *big.Int
	StrategyAmount *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterStrategyDeposited is a free log retrieval operation binding the contract event 0xf58adcbb0e084c416406ed1246914693d1bebcc911fa80eb933774c868034361.
//
// Solidity: event StrategyDeposited(uint256 _amount, uint256 _strategyAmount)
func (_Nodedaopool *NodedaopoolFilterer) FilterStrategyDeposited(opts *bind.FilterOpts) (*NodedaopoolStrategyDepositedIterator, error) {

	logs, sub, err := _Nodedaopool.contract.FilterLogs(opts, "StrategyDeposited")
	if err != nil {
		return nil, err
	}
	return &NodedaopoolStrategyDepositedIterator{contract: _Nodedaopool.contract, event: "StrategyDeposited", logs: logs, sub: sub}, nil
}

// WatchStrategyDeposited is a free log subscription operation binding the contract event 0xf58adcbb0e084c416406ed1246914693d1bebcc911fa80eb933774c868034361.
//
// Solidity: event StrategyDeposited(uint256 _amount, uint256 _strategyAmount)
func (_Nodedaopool *NodedaopoolFilterer) WatchStrategyDeposited(opts *bind.WatchOpts, sink chan<- *NodedaopoolStrategyDeposited) (event.Subscription, error) {

	logs, sub, err := _Nodedaopool.contract.WatchLogs(opts, "StrategyDeposited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodedaopoolStrategyDeposited)
				if err := _Nodedaopool.contract.UnpackLog(event, "StrategyDeposited", log); err != nil {
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

// ParseStrategyDeposited is a log parse operation binding the contract event 0xf58adcbb0e084c416406ed1246914693d1bebcc911fa80eb933774c868034361.
//
// Solidity: event StrategyDeposited(uint256 _amount, uint256 _strategyAmount)
func (_Nodedaopool *NodedaopoolFilterer) ParseStrategyDeposited(log types.Log) (*NodedaopoolStrategyDeposited, error) {
	event := new(NodedaopoolStrategyDeposited)
	if err := _Nodedaopool.contract.UnpackLog(event, "StrategyDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodedaopoolStrategyReturnIterator is returned from FilterStrategyReturn and is used to iterate over the raw logs and unpacked data for StrategyReturn events raised by the Nodedaopool contract.
type NodedaopoolStrategyReturnIterator struct {
	Event *NodedaopoolStrategyReturn // Event containing the contract specifics and raw log

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
func (it *NodedaopoolStrategyReturnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodedaopoolStrategyReturn)
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
		it.Event = new(NodedaopoolStrategyReturn)
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
func (it *NodedaopoolStrategyReturnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodedaopoolStrategyReturnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodedaopoolStrategyReturn represents a StrategyReturn event raised by the Nodedaopool contract.
type NodedaopoolStrategyReturn struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStrategyReturn is a free log retrieval operation binding the contract event 0x4fb0171001dc0ba5f8ca0996eb5a413cf67c611e3d6d1d6ae9ea1df128383410.
//
// Solidity: event StrategyReturn(uint256 _amount)
func (_Nodedaopool *NodedaopoolFilterer) FilterStrategyReturn(opts *bind.FilterOpts) (*NodedaopoolStrategyReturnIterator, error) {

	logs, sub, err := _Nodedaopool.contract.FilterLogs(opts, "StrategyReturn")
	if err != nil {
		return nil, err
	}
	return &NodedaopoolStrategyReturnIterator{contract: _Nodedaopool.contract, event: "StrategyReturn", logs: logs, sub: sub}, nil
}

// WatchStrategyReturn is a free log subscription operation binding the contract event 0x4fb0171001dc0ba5f8ca0996eb5a413cf67c611e3d6d1d6ae9ea1df128383410.
//
// Solidity: event StrategyReturn(uint256 _amount)
func (_Nodedaopool *NodedaopoolFilterer) WatchStrategyReturn(opts *bind.WatchOpts, sink chan<- *NodedaopoolStrategyReturn) (event.Subscription, error) {

	logs, sub, err := _Nodedaopool.contract.WatchLogs(opts, "StrategyReturn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodedaopoolStrategyReturn)
				if err := _Nodedaopool.contract.UnpackLog(event, "StrategyReturn", log); err != nil {
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

// ParseStrategyReturn is a log parse operation binding the contract event 0x4fb0171001dc0ba5f8ca0996eb5a413cf67c611e3d6d1d6ae9ea1df128383410.
//
// Solidity: event StrategyReturn(uint256 _amount)
func (_Nodedaopool *NodedaopoolFilterer) ParseStrategyReturn(log types.Log) (*NodedaopoolStrategyReturn, error) {
	event := new(NodedaopoolStrategyReturn)
	if err := _Nodedaopool.contract.UnpackLog(event, "StrategyReturn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodedaopoolStrategyVaultChangedIterator is returned from FilterStrategyVaultChanged and is used to iterate over the raw logs and unpacked data for StrategyVaultChanged events raised by the Nodedaopool contract.
type NodedaopoolStrategyVaultChangedIterator struct {
	Event *NodedaopoolStrategyVaultChanged // Event containing the contract specifics and raw log

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
func (it *NodedaopoolStrategyVaultChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodedaopoolStrategyVaultChanged)
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
		it.Event = new(NodedaopoolStrategyVaultChanged)
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
func (it *NodedaopoolStrategyVaultChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodedaopoolStrategyVaultChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodedaopoolStrategyVaultChanged represents a StrategyVaultChanged event raised by the Nodedaopool contract.
type NodedaopoolStrategyVaultChanged struct {
	OldStrategyVault common.Address
	StrategyVault    common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterStrategyVaultChanged is a free log retrieval operation binding the contract event 0x4ec60652ba5660c71f5d158ba76777060e4df45955126c42ffe26c0c447a781e.
//
// Solidity: event StrategyVaultChanged(address _oldStrategyVault, address _strategyVault)
func (_Nodedaopool *NodedaopoolFilterer) FilterStrategyVaultChanged(opts *bind.FilterOpts) (*NodedaopoolStrategyVaultChangedIterator, error) {

	logs, sub, err := _Nodedaopool.contract.FilterLogs(opts, "StrategyVaultChanged")
	if err != nil {
		return nil, err
	}
	return &NodedaopoolStrategyVaultChangedIterator{contract: _Nodedaopool.contract, event: "StrategyVaultChanged", logs: logs, sub: sub}, nil
}

// WatchStrategyVaultChanged is a free log subscription operation binding the contract event 0x4ec60652ba5660c71f5d158ba76777060e4df45955126c42ffe26c0c447a781e.
//
// Solidity: event StrategyVaultChanged(address _oldStrategyVault, address _strategyVault)
func (_Nodedaopool *NodedaopoolFilterer) WatchStrategyVaultChanged(opts *bind.WatchOpts, sink chan<- *NodedaopoolStrategyVaultChanged) (event.Subscription, error) {

	logs, sub, err := _Nodedaopool.contract.WatchLogs(opts, "StrategyVaultChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodedaopoolStrategyVaultChanged)
				if err := _Nodedaopool.contract.UnpackLog(event, "StrategyVaultChanged", log); err != nil {
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

// ParseStrategyVaultChanged is a log parse operation binding the contract event 0x4ec60652ba5660c71f5d158ba76777060e4df45955126c42ffe26c0c447a781e.
//
// Solidity: event StrategyVaultChanged(address _oldStrategyVault, address _strategyVault)
func (_Nodedaopool *NodedaopoolFilterer) ParseStrategyVaultChanged(log types.Log) (*NodedaopoolStrategyVaultChanged, error) {
	event := new(NodedaopoolStrategyVaultChanged)
	if err := _Nodedaopool.contract.UnpackLog(event, "StrategyVaultChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodedaopoolUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Nodedaopool contract.
type NodedaopoolUnpausedIterator struct {
	Event *NodedaopoolUnpaused // Event containing the contract specifics and raw log

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
func (it *NodedaopoolUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodedaopoolUnpaused)
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
		it.Event = new(NodedaopoolUnpaused)
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
func (it *NodedaopoolUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodedaopoolUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodedaopoolUnpaused represents a Unpaused event raised by the Nodedaopool contract.
type NodedaopoolUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Nodedaopool *NodedaopoolFilterer) FilterUnpaused(opts *bind.FilterOpts) (*NodedaopoolUnpausedIterator, error) {

	logs, sub, err := _Nodedaopool.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &NodedaopoolUnpausedIterator{contract: _Nodedaopool.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Nodedaopool *NodedaopoolFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *NodedaopoolUnpaused) (event.Subscription, error) {

	logs, sub, err := _Nodedaopool.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodedaopoolUnpaused)
				if err := _Nodedaopool.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_Nodedaopool *NodedaopoolFilterer) ParseUnpaused(log types.Log) (*NodedaopoolUnpaused, error) {
	event := new(NodedaopoolUnpaused)
	if err := _Nodedaopool.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodedaopoolUnstakeAllowedUpdatedIterator is returned from FilterUnstakeAllowedUpdated and is used to iterate over the raw logs and unpacked data for UnstakeAllowedUpdated events raised by the Nodedaopool contract.
type NodedaopoolUnstakeAllowedUpdatedIterator struct {
	Event *NodedaopoolUnstakeAllowedUpdated // Event containing the contract specifics and raw log

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
func (it *NodedaopoolUnstakeAllowedUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodedaopoolUnstakeAllowedUpdated)
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
		it.Event = new(NodedaopoolUnstakeAllowedUpdated)
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
func (it *NodedaopoolUnstakeAllowedUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodedaopoolUnstakeAllowedUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodedaopoolUnstakeAllowedUpdated represents a UnstakeAllowedUpdated event raised by the Nodedaopool contract.
type NodedaopoolUnstakeAllowedUpdated struct {
	OldUnstakeAllowed bool
	UnstakeAllowed    bool
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterUnstakeAllowedUpdated is a free log retrieval operation binding the contract event 0xf0700b7cce37d0ddc289445e07408f3709033f9785391e5e05b908cb9705748b.
//
// Solidity: event UnstakeAllowedUpdated(bool _oldUnstakeAllowed, bool _unstakeAllowed)
func (_Nodedaopool *NodedaopoolFilterer) FilterUnstakeAllowedUpdated(opts *bind.FilterOpts) (*NodedaopoolUnstakeAllowedUpdatedIterator, error) {

	logs, sub, err := _Nodedaopool.contract.FilterLogs(opts, "UnstakeAllowedUpdated")
	if err != nil {
		return nil, err
	}
	return &NodedaopoolUnstakeAllowedUpdatedIterator{contract: _Nodedaopool.contract, event: "UnstakeAllowedUpdated", logs: logs, sub: sub}, nil
}

// WatchUnstakeAllowedUpdated is a free log subscription operation binding the contract event 0xf0700b7cce37d0ddc289445e07408f3709033f9785391e5e05b908cb9705748b.
//
// Solidity: event UnstakeAllowedUpdated(bool _oldUnstakeAllowed, bool _unstakeAllowed)
func (_Nodedaopool *NodedaopoolFilterer) WatchUnstakeAllowedUpdated(opts *bind.WatchOpts, sink chan<- *NodedaopoolUnstakeAllowedUpdated) (event.Subscription, error) {

	logs, sub, err := _Nodedaopool.contract.WatchLogs(opts, "UnstakeAllowedUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodedaopoolUnstakeAllowedUpdated)
				if err := _Nodedaopool.contract.UnpackLog(event, "UnstakeAllowedUpdated", log); err != nil {
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
func (_Nodedaopool *NodedaopoolFilterer) ParseUnstakeAllowedUpdated(log types.Log) (*NodedaopoolUnstakeAllowedUpdated, error) {
	event := new(NodedaopoolUnstakeAllowedUpdated)
	if err := _Nodedaopool.contract.UnpackLog(event, "UnstakeAllowedUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodedaopoolUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Nodedaopool contract.
type NodedaopoolUpgradedIterator struct {
	Event *NodedaopoolUpgraded // Event containing the contract specifics and raw log

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
func (it *NodedaopoolUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodedaopoolUpgraded)
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
		it.Event = new(NodedaopoolUpgraded)
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
func (it *NodedaopoolUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodedaopoolUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodedaopoolUpgraded represents a Upgraded event raised by the Nodedaopool contract.
type NodedaopoolUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Nodedaopool *NodedaopoolFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*NodedaopoolUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Nodedaopool.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &NodedaopoolUpgradedIterator{contract: _Nodedaopool.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Nodedaopool *NodedaopoolFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *NodedaopoolUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Nodedaopool.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodedaopoolUpgraded)
				if err := _Nodedaopool.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_Nodedaopool *NodedaopoolFilterer) ParseUpgraded(log types.Log) (*NodedaopoolUpgraded, error) {
	event := new(NodedaopoolUpgraded)
	if err := _Nodedaopool.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodedaopoolValidatorManagerChangedIterator is returned from FilterValidatorManagerChanged and is used to iterate over the raw logs and unpacked data for ValidatorManagerChanged events raised by the Nodedaopool contract.
type NodedaopoolValidatorManagerChangedIterator struct {
	Event *NodedaopoolValidatorManagerChanged // Event containing the contract specifics and raw log

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
func (it *NodedaopoolValidatorManagerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodedaopoolValidatorManagerChanged)
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
		it.Event = new(NodedaopoolValidatorManagerChanged)
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
func (it *NodedaopoolValidatorManagerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodedaopoolValidatorManagerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodedaopoolValidatorManagerChanged represents a ValidatorManagerChanged event raised by the Nodedaopool contract.
type NodedaopoolValidatorManagerChanged struct {
	OldValidatorManager common.Address
	ValidatorManager    common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterValidatorManagerChanged is a free log retrieval operation binding the contract event 0xfcf16285f1e14b0b8a544860d3664317dd81073a42e811e33afe75a22886443e.
//
// Solidity: event ValidatorManagerChanged(address _oldValidatorManager, address _validatorManager)
func (_Nodedaopool *NodedaopoolFilterer) FilterValidatorManagerChanged(opts *bind.FilterOpts) (*NodedaopoolValidatorManagerChangedIterator, error) {

	logs, sub, err := _Nodedaopool.contract.FilterLogs(opts, "ValidatorManagerChanged")
	if err != nil {
		return nil, err
	}
	return &NodedaopoolValidatorManagerChangedIterator{contract: _Nodedaopool.contract, event: "ValidatorManagerChanged", logs: logs, sub: sub}, nil
}

// WatchValidatorManagerChanged is a free log subscription operation binding the contract event 0xfcf16285f1e14b0b8a544860d3664317dd81073a42e811e33afe75a22886443e.
//
// Solidity: event ValidatorManagerChanged(address _oldValidatorManager, address _validatorManager)
func (_Nodedaopool *NodedaopoolFilterer) WatchValidatorManagerChanged(opts *bind.WatchOpts, sink chan<- *NodedaopoolValidatorManagerChanged) (event.Subscription, error) {

	logs, sub, err := _Nodedaopool.contract.WatchLogs(opts, "ValidatorManagerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodedaopoolValidatorManagerChanged)
				if err := _Nodedaopool.contract.UnpackLog(event, "ValidatorManagerChanged", log); err != nil {
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
func (_Nodedaopool *NodedaopoolFilterer) ParseValidatorManagerChanged(log types.Log) (*NodedaopoolValidatorManagerChanged, error) {
	event := new(NodedaopoolValidatorManagerChanged)
	if err := _Nodedaopool.contract.UnpackLog(event, "ValidatorManagerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodedaopoolValidatorRegistrationIterator is returned from FilterValidatorRegistration and is used to iterate over the raw logs and unpacked data for ValidatorRegistration events raised by the Nodedaopool contract.
type NodedaopoolValidatorRegistrationIterator struct {
	Event *NodedaopoolValidatorRegistration // Event containing the contract specifics and raw log

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
func (it *NodedaopoolValidatorRegistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodedaopoolValidatorRegistration)
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
		it.Event = new(NodedaopoolValidatorRegistration)
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
func (it *NodedaopoolValidatorRegistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodedaopoolValidatorRegistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodedaopoolValidatorRegistration represents a ValidatorRegistration event raised by the Nodedaopool contract.
type NodedaopoolValidatorRegistration struct {
	Pubkeys [][]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterValidatorRegistration is a free log retrieval operation binding the contract event 0xe585eadb0042252d35431ecfed1027caf5672811fdf1db08cb49e32fee170505.
//
// Solidity: event ValidatorRegistration(bytes[] _pubkeys)
func (_Nodedaopool *NodedaopoolFilterer) FilterValidatorRegistration(opts *bind.FilterOpts) (*NodedaopoolValidatorRegistrationIterator, error) {

	logs, sub, err := _Nodedaopool.contract.FilterLogs(opts, "ValidatorRegistration")
	if err != nil {
		return nil, err
	}
	return &NodedaopoolValidatorRegistrationIterator{contract: _Nodedaopool.contract, event: "ValidatorRegistration", logs: logs, sub: sub}, nil
}

// WatchValidatorRegistration is a free log subscription operation binding the contract event 0xe585eadb0042252d35431ecfed1027caf5672811fdf1db08cb49e32fee170505.
//
// Solidity: event ValidatorRegistration(bytes[] _pubkeys)
func (_Nodedaopool *NodedaopoolFilterer) WatchValidatorRegistration(opts *bind.WatchOpts, sink chan<- *NodedaopoolValidatorRegistration) (event.Subscription, error) {

	logs, sub, err := _Nodedaopool.contract.WatchLogs(opts, "ValidatorRegistration")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodedaopoolValidatorRegistration)
				if err := _Nodedaopool.contract.UnpackLog(event, "ValidatorRegistration", log); err != nil {
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
func (_Nodedaopool *NodedaopoolFilterer) ParseValidatorRegistration(log types.Log) (*NodedaopoolValidatorRegistration, error) {
	event := new(NodedaopoolValidatorRegistration)
	if err := _Nodedaopool.contract.UnpackLog(event, "ValidatorRegistration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodedaopoolValidatorRegistryChangedIterator is returned from FilterValidatorRegistryChanged and is used to iterate over the raw logs and unpacked data for ValidatorRegistryChanged events raised by the Nodedaopool contract.
type NodedaopoolValidatorRegistryChangedIterator struct {
	Event *NodedaopoolValidatorRegistryChanged // Event containing the contract specifics and raw log

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
func (it *NodedaopoolValidatorRegistryChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodedaopoolValidatorRegistryChanged)
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
		it.Event = new(NodedaopoolValidatorRegistryChanged)
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
func (it *NodedaopoolValidatorRegistryChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodedaopoolValidatorRegistryChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodedaopoolValidatorRegistryChanged represents a ValidatorRegistryChanged event raised by the Nodedaopool contract.
type NodedaopoolValidatorRegistryChanged struct {
	OldValidatorManager common.Address
	ValidatorManager    common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterValidatorRegistryChanged is a free log retrieval operation binding the contract event 0x5f98eb9c39a016a522ab1ca3601f349c89b557a9b6471ce04afa8770d5589062.
//
// Solidity: event ValidatorRegistryChanged(address _oldValidatorManager, address _validatorManager)
func (_Nodedaopool *NodedaopoolFilterer) FilterValidatorRegistryChanged(opts *bind.FilterOpts) (*NodedaopoolValidatorRegistryChangedIterator, error) {

	logs, sub, err := _Nodedaopool.contract.FilterLogs(opts, "ValidatorRegistryChanged")
	if err != nil {
		return nil, err
	}
	return &NodedaopoolValidatorRegistryChangedIterator{contract: _Nodedaopool.contract, event: "ValidatorRegistryChanged", logs: logs, sub: sub}, nil
}

// WatchValidatorRegistryChanged is a free log subscription operation binding the contract event 0x5f98eb9c39a016a522ab1ca3601f349c89b557a9b6471ce04afa8770d5589062.
//
// Solidity: event ValidatorRegistryChanged(address _oldValidatorManager, address _validatorManager)
func (_Nodedaopool *NodedaopoolFilterer) WatchValidatorRegistryChanged(opts *bind.WatchOpts, sink chan<- *NodedaopoolValidatorRegistryChanged) (event.Subscription, error) {

	logs, sub, err := _Nodedaopool.contract.WatchLogs(opts, "ValidatorRegistryChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodedaopoolValidatorRegistryChanged)
				if err := _Nodedaopool.contract.UnpackLog(event, "ValidatorRegistryChanged", log); err != nil {
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
func (_Nodedaopool *NodedaopoolFilterer) ParseValidatorRegistryChanged(log types.Log) (*NodedaopoolValidatorRegistryChanged, error) {
	event := new(NodedaopoolValidatorRegistryChanged)
	if err := _Nodedaopool.contract.UnpackLog(event, "ValidatorRegistryChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodedaopoolWithdrawalDelayChangedIterator is returned from FilterWithdrawalDelayChanged and is used to iterate over the raw logs and unpacked data for WithdrawalDelayChanged events raised by the Nodedaopool contract.
type NodedaopoolWithdrawalDelayChangedIterator struct {
	Event *NodedaopoolWithdrawalDelayChanged // Event containing the contract specifics and raw log

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
func (it *NodedaopoolWithdrawalDelayChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodedaopoolWithdrawalDelayChanged)
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
		it.Event = new(NodedaopoolWithdrawalDelayChanged)
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
func (it *NodedaopoolWithdrawalDelayChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodedaopoolWithdrawalDelayChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodedaopoolWithdrawalDelayChanged represents a WithdrawalDelayChanged event raised by the Nodedaopool contract.
type NodedaopoolWithdrawalDelayChanged struct {
	OldWithdrawalDelayBlocks *big.Int
	WithdrawalDelayBlocks    *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalDelayChanged is a free log retrieval operation binding the contract event 0xab3f1d5eaee409b7067167f77f1fa3f8a863366d6fb2b88559cd4f9b8e03e182.
//
// Solidity: event WithdrawalDelayChanged(uint256 _oldWithdrawalDelayBlocks, uint256 _withdrawalDelayBlocks)
func (_Nodedaopool *NodedaopoolFilterer) FilterWithdrawalDelayChanged(opts *bind.FilterOpts) (*NodedaopoolWithdrawalDelayChangedIterator, error) {

	logs, sub, err := _Nodedaopool.contract.FilterLogs(opts, "WithdrawalDelayChanged")
	if err != nil {
		return nil, err
	}
	return &NodedaopoolWithdrawalDelayChangedIterator{contract: _Nodedaopool.contract, event: "WithdrawalDelayChanged", logs: logs, sub: sub}, nil
}

// WatchWithdrawalDelayChanged is a free log subscription operation binding the contract event 0xab3f1d5eaee409b7067167f77f1fa3f8a863366d6fb2b88559cd4f9b8e03e182.
//
// Solidity: event WithdrawalDelayChanged(uint256 _oldWithdrawalDelayBlocks, uint256 _withdrawalDelayBlocks)
func (_Nodedaopool *NodedaopoolFilterer) WatchWithdrawalDelayChanged(opts *bind.WatchOpts, sink chan<- *NodedaopoolWithdrawalDelayChanged) (event.Subscription, error) {

	logs, sub, err := _Nodedaopool.contract.WatchLogs(opts, "WithdrawalDelayChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodedaopoolWithdrawalDelayChanged)
				if err := _Nodedaopool.contract.UnpackLog(event, "WithdrawalDelayChanged", log); err != nil {
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
func (_Nodedaopool *NodedaopoolFilterer) ParseWithdrawalDelayChanged(log types.Log) (*NodedaopoolWithdrawalDelayChanged, error) {
	event := new(NodedaopoolWithdrawalDelayChanged)
	if err := _Nodedaopool.contract.UnpackLog(event, "WithdrawalDelayChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodedaopoolWithdrawalsClaimedIterator is returned from FilterWithdrawalsClaimed and is used to iterate over the raw logs and unpacked data for WithdrawalsClaimed events raised by the Nodedaopool contract.
type NodedaopoolWithdrawalsClaimedIterator struct {
	Event *NodedaopoolWithdrawalsClaimed // Event containing the contract specifics and raw log

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
func (it *NodedaopoolWithdrawalsClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodedaopoolWithdrawalsClaimed)
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
		it.Event = new(NodedaopoolWithdrawalsClaimed)
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
func (it *NodedaopoolWithdrawalsClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodedaopoolWithdrawalsClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodedaopoolWithdrawalsClaimed represents a WithdrawalsClaimed event raised by the Nodedaopool contract.
type NodedaopoolWithdrawalsClaimed struct {
	Receiver    common.Address
	RequestId   *big.Int
	ClaimAmount *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalsClaimed is a free log retrieval operation binding the contract event 0xf39bbe3e7fb2d887fe7e6e23dea5a53c1e720410dfb13e87567b873845136cf4.
//
// Solidity: event WithdrawalsClaimed(address _receiver, uint256 _requestId, uint256 _claimAmount)
func (_Nodedaopool *NodedaopoolFilterer) FilterWithdrawalsClaimed(opts *bind.FilterOpts) (*NodedaopoolWithdrawalsClaimedIterator, error) {

	logs, sub, err := _Nodedaopool.contract.FilterLogs(opts, "WithdrawalsClaimed")
	if err != nil {
		return nil, err
	}
	return &NodedaopoolWithdrawalsClaimedIterator{contract: _Nodedaopool.contract, event: "WithdrawalsClaimed", logs: logs, sub: sub}, nil
}

// WatchWithdrawalsClaimed is a free log subscription operation binding the contract event 0xf39bbe3e7fb2d887fe7e6e23dea5a53c1e720410dfb13e87567b873845136cf4.
//
// Solidity: event WithdrawalsClaimed(address _receiver, uint256 _requestId, uint256 _claimAmount)
func (_Nodedaopool *NodedaopoolFilterer) WatchWithdrawalsClaimed(opts *bind.WatchOpts, sink chan<- *NodedaopoolWithdrawalsClaimed) (event.Subscription, error) {

	logs, sub, err := _Nodedaopool.contract.WatchLogs(opts, "WithdrawalsClaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodedaopoolWithdrawalsClaimed)
				if err := _Nodedaopool.contract.UnpackLog(event, "WithdrawalsClaimed", log); err != nil {
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
func (_Nodedaopool *NodedaopoolFilterer) ParseWithdrawalsClaimed(log types.Log) (*NodedaopoolWithdrawalsClaimed, error) {
	event := new(NodedaopoolWithdrawalsClaimed)
	if err := _Nodedaopool.contract.UnpackLog(event, "WithdrawalsClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodedaopoolWithdrawalsRequestIterator is returned from FilterWithdrawalsRequest and is used to iterate over the raw logs and unpacked data for WithdrawalsRequest events raised by the Nodedaopool contract.
type NodedaopoolWithdrawalsRequestIterator struct {
	Event *NodedaopoolWithdrawalsRequest // Event containing the contract specifics and raw log

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
func (it *NodedaopoolWithdrawalsRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodedaopoolWithdrawalsRequest)
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
		it.Event = new(NodedaopoolWithdrawalsRequest)
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
func (it *NodedaopoolWithdrawalsRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodedaopoolWithdrawalsRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodedaopoolWithdrawalsRequest represents a WithdrawalsRequest event raised by the Nodedaopool contract.
type NodedaopoolWithdrawalsRequest struct {
	Receiver         common.Address
	WithdrawalAmount *big.Int
	BlockNumber      *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalsRequest is a free log retrieval operation binding the contract event 0x74ffedfd7821cf30dd556fac01944f4d077a2099ae55773bb89444cce29755f4.
//
// Solidity: event WithdrawalsRequest(address _receiver, uint256 _withdrawalAmount, uint256 _blockNumber)
func (_Nodedaopool *NodedaopoolFilterer) FilterWithdrawalsRequest(opts *bind.FilterOpts) (*NodedaopoolWithdrawalsRequestIterator, error) {

	logs, sub, err := _Nodedaopool.contract.FilterLogs(opts, "WithdrawalsRequest")
	if err != nil {
		return nil, err
	}
	return &NodedaopoolWithdrawalsRequestIterator{contract: _Nodedaopool.contract, event: "WithdrawalsRequest", logs: logs, sub: sub}, nil
}

// WatchWithdrawalsRequest is a free log subscription operation binding the contract event 0x74ffedfd7821cf30dd556fac01944f4d077a2099ae55773bb89444cce29755f4.
//
// Solidity: event WithdrawalsRequest(address _receiver, uint256 _withdrawalAmount, uint256 _blockNumber)
func (_Nodedaopool *NodedaopoolFilterer) WatchWithdrawalsRequest(opts *bind.WatchOpts, sink chan<- *NodedaopoolWithdrawalsRequest) (event.Subscription, error) {

	logs, sub, err := _Nodedaopool.contract.WatchLogs(opts, "WithdrawalsRequest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodedaopoolWithdrawalsRequest)
				if err := _Nodedaopool.contract.UnpackLog(event, "WithdrawalsRequest", log); err != nil {
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
func (_Nodedaopool *NodedaopoolFilterer) ParseWithdrawalsRequest(log types.Log) (*NodedaopoolWithdrawalsRequest, error) {
	event := new(NodedaopoolWithdrawalsRequest)
	if err := _Nodedaopool.contract.UnpackLog(event, "WithdrawalsRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
