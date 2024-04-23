// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ssv

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

// ISSVNetworkCoreCluster is an auto generated low-level Go binding around an user-defined struct.
type ISSVNetworkCoreCluster struct {
	ValidatorCount  uint32
	NetworkFeeIndex uint64
	Index           uint64
	Active          bool
	Balance         *big.Int
}

// SsvMetaData contains all meta data concerning the Ssv contract.
var SsvMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ApprovalNotWithinTimeframe\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotWhitelisted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ClusterAlreadyEnabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ClusterDoesNotExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ClusterIsLiquidated\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ClusterNotLiquidatable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExceedValidatorLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeExceedsIncreaseLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeIncreaseNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectClusterState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectValidatorState\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"}],\"name\":\"IncorrectValidatorStateWithData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidOperatorIdsLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPublicKeyLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxValueExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewBlockPeriodIsBelowMinimum\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoFeeDeclared\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAuthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperatorAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperatorDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperatorsListNotUnique\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PublicKeysSharesLengthMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SameFeeChangeNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TargetModuleDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsortedOperatorsList\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatorAlreadyExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"}],\"name\":\"ValidatorAlreadyExistsWithData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatorDoesNotExist\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64[]\",\"name\":\"operatorIds\",\"type\":\"uint64[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"validatorCount\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"networkFeeIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structISSVNetworkCore.Cluster\",\"name\":\"cluster\",\"type\":\"tuple\"}],\"name\":\"ClusterDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64[]\",\"name\":\"operatorIds\",\"type\":\"uint64[]\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"validatorCount\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"networkFeeIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structISSVNetworkCore.Cluster\",\"name\":\"cluster\",\"type\":\"tuple\"}],\"name\":\"ClusterLiquidated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64[]\",\"name\":\"operatorIds\",\"type\":\"uint64[]\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"validatorCount\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"networkFeeIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structISSVNetworkCore.Cluster\",\"name\":\"cluster\",\"type\":\"tuple\"}],\"name\":\"ClusterReactivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64[]\",\"name\":\"operatorIds\",\"type\":\"uint64[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"validatorCount\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"networkFeeIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structISSVNetworkCore.Cluster\",\"name\":\"cluster\",\"type\":\"tuple\"}],\"name\":\"ClusterWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"value\",\"type\":\"uint64\"}],\"name\":\"DeclareOperatorFeePeriodUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"value\",\"type\":\"uint64\"}],\"name\":\"ExecuteOperatorFeePeriodUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipientAddress\",\"type\":\"address\"}],\"name\":\"FeeRecipientAddressUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"value\",\"type\":\"uint64\"}],\"name\":\"LiquidationThresholdPeriodUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"MinimumLiquidationCollateralUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"NetworkEarningsWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newFee\",\"type\":\"uint256\"}],\"name\":\"NetworkFeeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"operatorId\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"OperatorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"operatorId\",\"type\":\"uint64\"}],\"name\":\"OperatorFeeDeclarationCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"operatorId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"OperatorFeeDeclared\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"operatorId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"OperatorFeeExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"value\",\"type\":\"uint64\"}],\"name\":\"OperatorFeeIncreaseLimitUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"maxFee\",\"type\":\"uint64\"}],\"name\":\"OperatorMaximumFeeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"operatorId\",\"type\":\"uint64\"}],\"name\":\"OperatorRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"operatorId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"whitelisted\",\"type\":\"address\"}],\"name\":\"OperatorWhitelistUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"operatorId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"OperatorWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64[]\",\"name\":\"operatorIds\",\"type\":\"uint64[]\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"shares\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"validatorCount\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"networkFeeIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structISSVNetworkCore.Cluster\",\"name\":\"cluster\",\"type\":\"tuple\"}],\"name\":\"ValidatorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64[]\",\"name\":\"operatorIds\",\"type\":\"uint64[]\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"}],\"name\":\"ValidatorExited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64[]\",\"name\":\"operatorIds\",\"type\":\"uint64[]\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"validatorCount\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"networkFeeIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structISSVNetworkCore.Cluster\",\"name\":\"cluster\",\"type\":\"tuple\"}],\"name\":\"ValidatorRemoved\",\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"publicKeys\",\"type\":\"bytes[]\"},{\"internalType\":\"uint64[]\",\"name\":\"operatorIds\",\"type\":\"uint64[]\"}],\"name\":\"bulkExitValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"publicKeys\",\"type\":\"bytes[]\"},{\"internalType\":\"uint64[]\",\"name\":\"operatorIds\",\"type\":\"uint64[]\"},{\"internalType\":\"bytes[]\",\"name\":\"sharesData\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"validatorCount\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"networkFeeIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"internalType\":\"structISSVNetworkCore.Cluster\",\"name\":\"cluster\",\"type\":\"tuple\"}],\"name\":\"bulkRegisterValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"publicKeys\",\"type\":\"bytes[]\"},{\"internalType\":\"uint64[]\",\"name\":\"operatorIds\",\"type\":\"uint64[]\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"validatorCount\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"networkFeeIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"internalType\":\"structISSVNetworkCore.Cluster\",\"name\":\"cluster\",\"type\":\"tuple\"}],\"name\":\"bulkRemoveValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"operatorId\",\"type\":\"uint64\"}],\"name\":\"cancelDeclaredOperatorFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"operatorId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"declareOperatorFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"clusterOwner\",\"type\":\"address\"},{\"internalType\":\"uint64[]\",\"name\":\"operatorIds\",\"type\":\"uint64[]\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"validatorCount\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"networkFeeIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"internalType\":\"structISSVNetworkCore.Cluster\",\"name\":\"cluster\",\"type\":\"tuple\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"operatorId\",\"type\":\"uint64\"}],\"name\":\"executeOperatorFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64[]\",\"name\":\"operatorIds\",\"type\":\"uint64[]\"}],\"name\":\"exitValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token_\",\"type\":\"address\"},{\"internalType\":\"contractISSVOperators\",\"name\":\"ssvOperators_\",\"type\":\"address\"},{\"internalType\":\"contractISSVClusters\",\"name\":\"ssvClusters_\",\"type\":\"address\"},{\"internalType\":\"contractISSVDAO\",\"name\":\"ssvDAO_\",\"type\":\"address\"},{\"internalType\":\"contractISSVViews\",\"name\":\"ssvViews_\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"minimumBlocksBeforeLiquidation_\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"minimumLiquidationCollateral_\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"validatorsPerOperatorLimit_\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"declareOperatorFeePeriod_\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"executeOperatorFeePeriod_\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"operatorMaxFeeIncrease_\",\"type\":\"uint64\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"clusterOwner\",\"type\":\"address\"},{\"internalType\":\"uint64[]\",\"name\":\"operatorIds\",\"type\":\"uint64[]\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"validatorCount\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"networkFeeIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"internalType\":\"structISSVNetworkCore.Cluster\",\"name\":\"cluster\",\"type\":\"tuple\"}],\"name\":\"liquidate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64[]\",\"name\":\"operatorIds\",\"type\":\"uint64[]\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"validatorCount\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"networkFeeIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"internalType\":\"structISSVNetworkCore.Cluster\",\"name\":\"cluster\",\"type\":\"tuple\"}],\"name\":\"reactivate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"operatorId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"reduceOperatorFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"registerOperator\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64[]\",\"name\":\"operatorIds\",\"type\":\"uint64[]\"},{\"internalType\":\"bytes\",\"name\":\"sharesData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"validatorCount\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"networkFeeIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"internalType\":\"structISSVNetworkCore.Cluster\",\"name\":\"cluster\",\"type\":\"tuple\"}],\"name\":\"registerValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"operatorId\",\"type\":\"uint64\"}],\"name\":\"removeOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint64[]\",\"name\":\"operatorIds\",\"type\":\"uint64[]\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"validatorCount\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"networkFeeIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"internalType\":\"structISSVNetworkCore.Cluster\",\"name\":\"cluster\",\"type\":\"tuple\"}],\"name\":\"removeValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipientAddress\",\"type\":\"address\"}],\"name\":\"setFeeRecipientAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"operatorId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"whitelisted\",\"type\":\"address\"}],\"name\":\"setOperatorWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"timeInSeconds\",\"type\":\"uint64\"}],\"name\":\"updateDeclareOperatorFeePeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"timeInSeconds\",\"type\":\"uint64\"}],\"name\":\"updateExecuteOperatorFeePeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"blocks\",\"type\":\"uint64\"}],\"name\":\"updateLiquidationThresholdPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"maxFee\",\"type\":\"uint64\"}],\"name\":\"updateMaximumOperatorFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"updateMinimumLiquidationCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumSSVModules\",\"name\":\"moduleId\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"moduleAddress\",\"type\":\"address\"}],\"name\":\"updateModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"updateNetworkFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"percentage\",\"type\":\"uint64\"}],\"name\":\"updateOperatorFeeIncreaseLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64[]\",\"name\":\"operatorIds\",\"type\":\"uint64[]\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"validatorCount\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"networkFeeIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"internalType\":\"structISSVNetworkCore.Cluster\",\"name\":\"cluster\",\"type\":\"tuple\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"operatorId\",\"type\":\"uint64\"}],\"name\":\"withdrawAllOperatorEarnings\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawNetworkEarnings\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"operatorId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawOperatorEarnings\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// SsvABI is the input ABI used to generate the binding from.
// Deprecated: Use SsvMetaData.ABI instead.
var SsvABI = SsvMetaData.ABI

// Ssv is an auto generated Go binding around an Ethereum contract.
type Ssv struct {
	SsvCaller     // Read-only binding to the contract
	SsvTransactor // Write-only binding to the contract
	SsvFilterer   // Log filterer for contract events
}

// SsvCaller is an auto generated read-only Go binding around an Ethereum contract.
type SsvCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SsvTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SsvTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SsvFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SsvFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SsvSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SsvSession struct {
	Contract     *Ssv              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SsvCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SsvCallerSession struct {
	Contract *SsvCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SsvTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SsvTransactorSession struct {
	Contract     *SsvTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SsvRaw is an auto generated low-level Go binding around an Ethereum contract.
type SsvRaw struct {
	Contract *Ssv // Generic contract binding to access the raw methods on
}

// SsvCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SsvCallerRaw struct {
	Contract *SsvCaller // Generic read-only contract binding to access the raw methods on
}

// SsvTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SsvTransactorRaw struct {
	Contract *SsvTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSsv creates a new instance of Ssv, bound to a specific deployed contract.
func NewSsv(address common.Address, backend bind.ContractBackend) (*Ssv, error) {
	contract, err := bindSsv(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ssv{SsvCaller: SsvCaller{contract: contract}, SsvTransactor: SsvTransactor{contract: contract}, SsvFilterer: SsvFilterer{contract: contract}}, nil
}

// NewSsvCaller creates a new read-only instance of Ssv, bound to a specific deployed contract.
func NewSsvCaller(address common.Address, caller bind.ContractCaller) (*SsvCaller, error) {
	contract, err := bindSsv(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SsvCaller{contract: contract}, nil
}

// NewSsvTransactor creates a new write-only instance of Ssv, bound to a specific deployed contract.
func NewSsvTransactor(address common.Address, transactor bind.ContractTransactor) (*SsvTransactor, error) {
	contract, err := bindSsv(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SsvTransactor{contract: contract}, nil
}

// NewSsvFilterer creates a new log filterer instance of Ssv, bound to a specific deployed contract.
func NewSsvFilterer(address common.Address, filterer bind.ContractFilterer) (*SsvFilterer, error) {
	contract, err := bindSsv(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SsvFilterer{contract: contract}, nil
}

// bindSsv binds a generic wrapper to an already deployed contract.
func bindSsv(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SsvABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ssv *SsvRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ssv.Contract.SsvCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ssv *SsvRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ssv.Contract.SsvTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ssv *SsvRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ssv.Contract.SsvTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ssv *SsvCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ssv.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ssv *SsvTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ssv.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ssv *SsvTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ssv.Contract.contract.Transact(opts, method, params...)
}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() pure returns(string version)
func (_Ssv *SsvCaller) GetVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Ssv.contract.Call(opts, &out, "getVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() pure returns(string version)
func (_Ssv *SsvSession) GetVersion() (string, error) {
	return _Ssv.Contract.GetVersion(&_Ssv.CallOpts)
}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() pure returns(string version)
func (_Ssv *SsvCallerSession) GetVersion() (string, error) {
	return _Ssv.Contract.GetVersion(&_Ssv.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ssv *SsvCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ssv.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ssv *SsvSession) Owner() (common.Address, error) {
	return _Ssv.Contract.Owner(&_Ssv.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ssv *SsvCallerSession) Owner() (common.Address, error) {
	return _Ssv.Contract.Owner(&_Ssv.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Ssv *SsvCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ssv.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Ssv *SsvSession) PendingOwner() (common.Address, error) {
	return _Ssv.Contract.PendingOwner(&_Ssv.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Ssv *SsvCallerSession) PendingOwner() (common.Address, error) {
	return _Ssv.Contract.PendingOwner(&_Ssv.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Ssv *SsvCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Ssv.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Ssv *SsvSession) ProxiableUUID() ([32]byte, error) {
	return _Ssv.Contract.ProxiableUUID(&_Ssv.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Ssv *SsvCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Ssv.Contract.ProxiableUUID(&_Ssv.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Ssv *SsvTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ssv.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Ssv *SsvSession) AcceptOwnership() (*types.Transaction, error) {
	return _Ssv.Contract.AcceptOwnership(&_Ssv.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Ssv *SsvTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Ssv.Contract.AcceptOwnership(&_Ssv.TransactOpts)
}

// BulkExitValidator is a paid mutator transaction binding the contract method 0x32afd02f.
//
// Solidity: function bulkExitValidator(bytes[] publicKeys, uint64[] operatorIds) returns()
func (_Ssv *SsvTransactor) BulkExitValidator(opts *bind.TransactOpts, publicKeys [][]byte, operatorIds []uint64) (*types.Transaction, error) {
	return _Ssv.contract.Transact(opts, "bulkExitValidator", publicKeys, operatorIds)
}

// BulkExitValidator is a paid mutator transaction binding the contract method 0x32afd02f.
//
// Solidity: function bulkExitValidator(bytes[] publicKeys, uint64[] operatorIds) returns()
func (_Ssv *SsvSession) BulkExitValidator(publicKeys [][]byte, operatorIds []uint64) (*types.Transaction, error) {
	return _Ssv.Contract.BulkExitValidator(&_Ssv.TransactOpts, publicKeys, operatorIds)
}

// BulkExitValidator is a paid mutator transaction binding the contract method 0x32afd02f.
//
// Solidity: function bulkExitValidator(bytes[] publicKeys, uint64[] operatorIds) returns()
func (_Ssv *SsvTransactorSession) BulkExitValidator(publicKeys [][]byte, operatorIds []uint64) (*types.Transaction, error) {
	return _Ssv.Contract.BulkExitValidator(&_Ssv.TransactOpts, publicKeys, operatorIds)
}

// BulkRegisterValidator is a paid mutator transaction binding the contract method 0x22f18bf5.
//
// Solidity: function bulkRegisterValidator(bytes[] publicKeys, uint64[] operatorIds, bytes[] sharesData, uint256 amount, (uint32,uint64,uint64,bool,uint256) cluster) returns()
func (_Ssv *SsvTransactor) BulkRegisterValidator(opts *bind.TransactOpts, publicKeys [][]byte, operatorIds []uint64, sharesData [][]byte, amount *big.Int, cluster ISSVNetworkCoreCluster) (*types.Transaction, error) {
	return _Ssv.contract.Transact(opts, "bulkRegisterValidator", publicKeys, operatorIds, sharesData, amount, cluster)
}

// BulkRegisterValidator is a paid mutator transaction binding the contract method 0x22f18bf5.
//
// Solidity: function bulkRegisterValidator(bytes[] publicKeys, uint64[] operatorIds, bytes[] sharesData, uint256 amount, (uint32,uint64,uint64,bool,uint256) cluster) returns()
func (_Ssv *SsvSession) BulkRegisterValidator(publicKeys [][]byte, operatorIds []uint64, sharesData [][]byte, amount *big.Int, cluster ISSVNetworkCoreCluster) (*types.Transaction, error) {
	return _Ssv.Contract.BulkRegisterValidator(&_Ssv.TransactOpts, publicKeys, operatorIds, sharesData, amount, cluster)
}

// BulkRegisterValidator is a paid mutator transaction binding the contract method 0x22f18bf5.
//
// Solidity: function bulkRegisterValidator(bytes[] publicKeys, uint64[] operatorIds, bytes[] sharesData, uint256 amount, (uint32,uint64,uint64,bool,uint256) cluster) returns()
func (_Ssv *SsvTransactorSession) BulkRegisterValidator(publicKeys [][]byte, operatorIds []uint64, sharesData [][]byte, amount *big.Int, cluster ISSVNetworkCoreCluster) (*types.Transaction, error) {
	return _Ssv.Contract.BulkRegisterValidator(&_Ssv.TransactOpts, publicKeys, operatorIds, sharesData, amount, cluster)
}

// BulkRemoveValidator is a paid mutator transaction binding the contract method 0x5aed1142.
//
// Solidity: function bulkRemoveValidator(bytes[] publicKeys, uint64[] operatorIds, (uint32,uint64,uint64,bool,uint256) cluster) returns()
func (_Ssv *SsvTransactor) BulkRemoveValidator(opts *bind.TransactOpts, publicKeys [][]byte, operatorIds []uint64, cluster ISSVNetworkCoreCluster) (*types.Transaction, error) {
	return _Ssv.contract.Transact(opts, "bulkRemoveValidator", publicKeys, operatorIds, cluster)
}

// BulkRemoveValidator is a paid mutator transaction binding the contract method 0x5aed1142.
//
// Solidity: function bulkRemoveValidator(bytes[] publicKeys, uint64[] operatorIds, (uint32,uint64,uint64,bool,uint256) cluster) returns()
func (_Ssv *SsvSession) BulkRemoveValidator(publicKeys [][]byte, operatorIds []uint64, cluster ISSVNetworkCoreCluster) (*types.Transaction, error) {
	return _Ssv.Contract.BulkRemoveValidator(&_Ssv.TransactOpts, publicKeys, operatorIds, cluster)
}

// BulkRemoveValidator is a paid mutator transaction binding the contract method 0x5aed1142.
//
// Solidity: function bulkRemoveValidator(bytes[] publicKeys, uint64[] operatorIds, (uint32,uint64,uint64,bool,uint256) cluster) returns()
func (_Ssv *SsvTransactorSession) BulkRemoveValidator(publicKeys [][]byte, operatorIds []uint64, cluster ISSVNetworkCoreCluster) (*types.Transaction, error) {
	return _Ssv.Contract.BulkRemoveValidator(&_Ssv.TransactOpts, publicKeys, operatorIds, cluster)
}

// CancelDeclaredOperatorFee is a paid mutator transaction binding the contract method 0x23d68a6d.
//
// Solidity: function cancelDeclaredOperatorFee(uint64 operatorId) returns()
func (_Ssv *SsvTransactor) CancelDeclaredOperatorFee(opts *bind.TransactOpts, operatorId uint64) (*types.Transaction, error) {
	return _Ssv.contract.Transact(opts, "cancelDeclaredOperatorFee", operatorId)
}

// CancelDeclaredOperatorFee is a paid mutator transaction binding the contract method 0x23d68a6d.
//
// Solidity: function cancelDeclaredOperatorFee(uint64 operatorId) returns()
func (_Ssv *SsvSession) CancelDeclaredOperatorFee(operatorId uint64) (*types.Transaction, error) {
	return _Ssv.Contract.CancelDeclaredOperatorFee(&_Ssv.TransactOpts, operatorId)
}

// CancelDeclaredOperatorFee is a paid mutator transaction binding the contract method 0x23d68a6d.
//
// Solidity: function cancelDeclaredOperatorFee(uint64 operatorId) returns()
func (_Ssv *SsvTransactorSession) CancelDeclaredOperatorFee(operatorId uint64) (*types.Transaction, error) {
	return _Ssv.Contract.CancelDeclaredOperatorFee(&_Ssv.TransactOpts, operatorId)
}

// DeclareOperatorFee is a paid mutator transaction binding the contract method 0xb317c35f.
//
// Solidity: function declareOperatorFee(uint64 operatorId, uint256 fee) returns()
func (_Ssv *SsvTransactor) DeclareOperatorFee(opts *bind.TransactOpts, operatorId uint64, fee *big.Int) (*types.Transaction, error) {
	return _Ssv.contract.Transact(opts, "declareOperatorFee", operatorId, fee)
}

// DeclareOperatorFee is a paid mutator transaction binding the contract method 0xb317c35f.
//
// Solidity: function declareOperatorFee(uint64 operatorId, uint256 fee) returns()
func (_Ssv *SsvSession) DeclareOperatorFee(operatorId uint64, fee *big.Int) (*types.Transaction, error) {
	return _Ssv.Contract.DeclareOperatorFee(&_Ssv.TransactOpts, operatorId, fee)
}

// DeclareOperatorFee is a paid mutator transaction binding the contract method 0xb317c35f.
//
// Solidity: function declareOperatorFee(uint64 operatorId, uint256 fee) returns()
func (_Ssv *SsvTransactorSession) DeclareOperatorFee(operatorId uint64, fee *big.Int) (*types.Transaction, error) {
	return _Ssv.Contract.DeclareOperatorFee(&_Ssv.TransactOpts, operatorId, fee)
}

// Deposit is a paid mutator transaction binding the contract method 0xbc26e7e5.
//
// Solidity: function deposit(address clusterOwner, uint64[] operatorIds, uint256 amount, (uint32,uint64,uint64,bool,uint256) cluster) returns()
func (_Ssv *SsvTransactor) Deposit(opts *bind.TransactOpts, clusterOwner common.Address, operatorIds []uint64, amount *big.Int, cluster ISSVNetworkCoreCluster) (*types.Transaction, error) {
	return _Ssv.contract.Transact(opts, "deposit", clusterOwner, operatorIds, amount, cluster)
}

// Deposit is a paid mutator transaction binding the contract method 0xbc26e7e5.
//
// Solidity: function deposit(address clusterOwner, uint64[] operatorIds, uint256 amount, (uint32,uint64,uint64,bool,uint256) cluster) returns()
func (_Ssv *SsvSession) Deposit(clusterOwner common.Address, operatorIds []uint64, amount *big.Int, cluster ISSVNetworkCoreCluster) (*types.Transaction, error) {
	return _Ssv.Contract.Deposit(&_Ssv.TransactOpts, clusterOwner, operatorIds, amount, cluster)
}

// Deposit is a paid mutator transaction binding the contract method 0xbc26e7e5.
//
// Solidity: function deposit(address clusterOwner, uint64[] operatorIds, uint256 amount, (uint32,uint64,uint64,bool,uint256) cluster) returns()
func (_Ssv *SsvTransactorSession) Deposit(clusterOwner common.Address, operatorIds []uint64, amount *big.Int, cluster ISSVNetworkCoreCluster) (*types.Transaction, error) {
	return _Ssv.Contract.Deposit(&_Ssv.TransactOpts, clusterOwner, operatorIds, amount, cluster)
}

// ExecuteOperatorFee is a paid mutator transaction binding the contract method 0x8932cee0.
//
// Solidity: function executeOperatorFee(uint64 operatorId) returns()
func (_Ssv *SsvTransactor) ExecuteOperatorFee(opts *bind.TransactOpts, operatorId uint64) (*types.Transaction, error) {
	return _Ssv.contract.Transact(opts, "executeOperatorFee", operatorId)
}

// ExecuteOperatorFee is a paid mutator transaction binding the contract method 0x8932cee0.
//
// Solidity: function executeOperatorFee(uint64 operatorId) returns()
func (_Ssv *SsvSession) ExecuteOperatorFee(operatorId uint64) (*types.Transaction, error) {
	return _Ssv.Contract.ExecuteOperatorFee(&_Ssv.TransactOpts, operatorId)
}

// ExecuteOperatorFee is a paid mutator transaction binding the contract method 0x8932cee0.
//
// Solidity: function executeOperatorFee(uint64 operatorId) returns()
func (_Ssv *SsvTransactorSession) ExecuteOperatorFee(operatorId uint64) (*types.Transaction, error) {
	return _Ssv.Contract.ExecuteOperatorFee(&_Ssv.TransactOpts, operatorId)
}

// ExitValidator is a paid mutator transaction binding the contract method 0x3877322b.
//
// Solidity: function exitValidator(bytes publicKey, uint64[] operatorIds) returns()
func (_Ssv *SsvTransactor) ExitValidator(opts *bind.TransactOpts, publicKey []byte, operatorIds []uint64) (*types.Transaction, error) {
	return _Ssv.contract.Transact(opts, "exitValidator", publicKey, operatorIds)
}

// ExitValidator is a paid mutator transaction binding the contract method 0x3877322b.
//
// Solidity: function exitValidator(bytes publicKey, uint64[] operatorIds) returns()
func (_Ssv *SsvSession) ExitValidator(publicKey []byte, operatorIds []uint64) (*types.Transaction, error) {
	return _Ssv.Contract.ExitValidator(&_Ssv.TransactOpts, publicKey, operatorIds)
}

// ExitValidator is a paid mutator transaction binding the contract method 0x3877322b.
//
// Solidity: function exitValidator(bytes publicKey, uint64[] operatorIds) returns()
func (_Ssv *SsvTransactorSession) ExitValidator(publicKey []byte, operatorIds []uint64) (*types.Transaction, error) {
	return _Ssv.Contract.ExitValidator(&_Ssv.TransactOpts, publicKey, operatorIds)
}

// Initialize is a paid mutator transaction binding the contract method 0xc626c3c6.
//
// Solidity: function initialize(address token_, address ssvOperators_, address ssvClusters_, address ssvDAO_, address ssvViews_, uint64 minimumBlocksBeforeLiquidation_, uint256 minimumLiquidationCollateral_, uint32 validatorsPerOperatorLimit_, uint64 declareOperatorFeePeriod_, uint64 executeOperatorFeePeriod_, uint64 operatorMaxFeeIncrease_) returns()
func (_Ssv *SsvTransactor) Initialize(opts *bind.TransactOpts, token_ common.Address, ssvOperators_ common.Address, ssvClusters_ common.Address, ssvDAO_ common.Address, ssvViews_ common.Address, minimumBlocksBeforeLiquidation_ uint64, minimumLiquidationCollateral_ *big.Int, validatorsPerOperatorLimit_ uint32, declareOperatorFeePeriod_ uint64, executeOperatorFeePeriod_ uint64, operatorMaxFeeIncrease_ uint64) (*types.Transaction, error) {
	return _Ssv.contract.Transact(opts, "initialize", token_, ssvOperators_, ssvClusters_, ssvDAO_, ssvViews_, minimumBlocksBeforeLiquidation_, minimumLiquidationCollateral_, validatorsPerOperatorLimit_, declareOperatorFeePeriod_, executeOperatorFeePeriod_, operatorMaxFeeIncrease_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc626c3c6.
//
// Solidity: function initialize(address token_, address ssvOperators_, address ssvClusters_, address ssvDAO_, address ssvViews_, uint64 minimumBlocksBeforeLiquidation_, uint256 minimumLiquidationCollateral_, uint32 validatorsPerOperatorLimit_, uint64 declareOperatorFeePeriod_, uint64 executeOperatorFeePeriod_, uint64 operatorMaxFeeIncrease_) returns()
func (_Ssv *SsvSession) Initialize(token_ common.Address, ssvOperators_ common.Address, ssvClusters_ common.Address, ssvDAO_ common.Address, ssvViews_ common.Address, minimumBlocksBeforeLiquidation_ uint64, minimumLiquidationCollateral_ *big.Int, validatorsPerOperatorLimit_ uint32, declareOperatorFeePeriod_ uint64, executeOperatorFeePeriod_ uint64, operatorMaxFeeIncrease_ uint64) (*types.Transaction, error) {
	return _Ssv.Contract.Initialize(&_Ssv.TransactOpts, token_, ssvOperators_, ssvClusters_, ssvDAO_, ssvViews_, minimumBlocksBeforeLiquidation_, minimumLiquidationCollateral_, validatorsPerOperatorLimit_, declareOperatorFeePeriod_, executeOperatorFeePeriod_, operatorMaxFeeIncrease_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc626c3c6.
//
// Solidity: function initialize(address token_, address ssvOperators_, address ssvClusters_, address ssvDAO_, address ssvViews_, uint64 minimumBlocksBeforeLiquidation_, uint256 minimumLiquidationCollateral_, uint32 validatorsPerOperatorLimit_, uint64 declareOperatorFeePeriod_, uint64 executeOperatorFeePeriod_, uint64 operatorMaxFeeIncrease_) returns()
func (_Ssv *SsvTransactorSession) Initialize(token_ common.Address, ssvOperators_ common.Address, ssvClusters_ common.Address, ssvDAO_ common.Address, ssvViews_ common.Address, minimumBlocksBeforeLiquidation_ uint64, minimumLiquidationCollateral_ *big.Int, validatorsPerOperatorLimit_ uint32, declareOperatorFeePeriod_ uint64, executeOperatorFeePeriod_ uint64, operatorMaxFeeIncrease_ uint64) (*types.Transaction, error) {
	return _Ssv.Contract.Initialize(&_Ssv.TransactOpts, token_, ssvOperators_, ssvClusters_, ssvDAO_, ssvViews_, minimumBlocksBeforeLiquidation_, minimumLiquidationCollateral_, validatorsPerOperatorLimit_, declareOperatorFeePeriod_, executeOperatorFeePeriod_, operatorMaxFeeIncrease_)
}

// Liquidate is a paid mutator transaction binding the contract method 0xbf0f2fb2.
//
// Solidity: function liquidate(address clusterOwner, uint64[] operatorIds, (uint32,uint64,uint64,bool,uint256) cluster) returns()
func (_Ssv *SsvTransactor) Liquidate(opts *bind.TransactOpts, clusterOwner common.Address, operatorIds []uint64, cluster ISSVNetworkCoreCluster) (*types.Transaction, error) {
	return _Ssv.contract.Transact(opts, "liquidate", clusterOwner, operatorIds, cluster)
}

// Liquidate is a paid mutator transaction binding the contract method 0xbf0f2fb2.
//
// Solidity: function liquidate(address clusterOwner, uint64[] operatorIds, (uint32,uint64,uint64,bool,uint256) cluster) returns()
func (_Ssv *SsvSession) Liquidate(clusterOwner common.Address, operatorIds []uint64, cluster ISSVNetworkCoreCluster) (*types.Transaction, error) {
	return _Ssv.Contract.Liquidate(&_Ssv.TransactOpts, clusterOwner, operatorIds, cluster)
}

// Liquidate is a paid mutator transaction binding the contract method 0xbf0f2fb2.
//
// Solidity: function liquidate(address clusterOwner, uint64[] operatorIds, (uint32,uint64,uint64,bool,uint256) cluster) returns()
func (_Ssv *SsvTransactorSession) Liquidate(clusterOwner common.Address, operatorIds []uint64, cluster ISSVNetworkCoreCluster) (*types.Transaction, error) {
	return _Ssv.Contract.Liquidate(&_Ssv.TransactOpts, clusterOwner, operatorIds, cluster)
}

// Reactivate is a paid mutator transaction binding the contract method 0x5fec6dd0.
//
// Solidity: function reactivate(uint64[] operatorIds, uint256 amount, (uint32,uint64,uint64,bool,uint256) cluster) returns()
func (_Ssv *SsvTransactor) Reactivate(opts *bind.TransactOpts, operatorIds []uint64, amount *big.Int, cluster ISSVNetworkCoreCluster) (*types.Transaction, error) {
	return _Ssv.contract.Transact(opts, "reactivate", operatorIds, amount, cluster)
}

// Reactivate is a paid mutator transaction binding the contract method 0x5fec6dd0.
//
// Solidity: function reactivate(uint64[] operatorIds, uint256 amount, (uint32,uint64,uint64,bool,uint256) cluster) returns()
func (_Ssv *SsvSession) Reactivate(operatorIds []uint64, amount *big.Int, cluster ISSVNetworkCoreCluster) (*types.Transaction, error) {
	return _Ssv.Contract.Reactivate(&_Ssv.TransactOpts, operatorIds, amount, cluster)
}

// Reactivate is a paid mutator transaction binding the contract method 0x5fec6dd0.
//
// Solidity: function reactivate(uint64[] operatorIds, uint256 amount, (uint32,uint64,uint64,bool,uint256) cluster) returns()
func (_Ssv *SsvTransactorSession) Reactivate(operatorIds []uint64, amount *big.Int, cluster ISSVNetworkCoreCluster) (*types.Transaction, error) {
	return _Ssv.Contract.Reactivate(&_Ssv.TransactOpts, operatorIds, amount, cluster)
}

// ReduceOperatorFee is a paid mutator transaction binding the contract method 0x190d82e4.
//
// Solidity: function reduceOperatorFee(uint64 operatorId, uint256 fee) returns()
func (_Ssv *SsvTransactor) ReduceOperatorFee(opts *bind.TransactOpts, operatorId uint64, fee *big.Int) (*types.Transaction, error) {
	return _Ssv.contract.Transact(opts, "reduceOperatorFee", operatorId, fee)
}

// ReduceOperatorFee is a paid mutator transaction binding the contract method 0x190d82e4.
//
// Solidity: function reduceOperatorFee(uint64 operatorId, uint256 fee) returns()
func (_Ssv *SsvSession) ReduceOperatorFee(operatorId uint64, fee *big.Int) (*types.Transaction, error) {
	return _Ssv.Contract.ReduceOperatorFee(&_Ssv.TransactOpts, operatorId, fee)
}

// ReduceOperatorFee is a paid mutator transaction binding the contract method 0x190d82e4.
//
// Solidity: function reduceOperatorFee(uint64 operatorId, uint256 fee) returns()
func (_Ssv *SsvTransactorSession) ReduceOperatorFee(operatorId uint64, fee *big.Int) (*types.Transaction, error) {
	return _Ssv.Contract.ReduceOperatorFee(&_Ssv.TransactOpts, operatorId, fee)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0xff212c5c.
//
// Solidity: function registerOperator(bytes publicKey, uint256 fee) returns(uint64 id)
func (_Ssv *SsvTransactor) RegisterOperator(opts *bind.TransactOpts, publicKey []byte, fee *big.Int) (*types.Transaction, error) {
	return _Ssv.contract.Transact(opts, "registerOperator", publicKey, fee)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0xff212c5c.
//
// Solidity: function registerOperator(bytes publicKey, uint256 fee) returns(uint64 id)
func (_Ssv *SsvSession) RegisterOperator(publicKey []byte, fee *big.Int) (*types.Transaction, error) {
	return _Ssv.Contract.RegisterOperator(&_Ssv.TransactOpts, publicKey, fee)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0xff212c5c.
//
// Solidity: function registerOperator(bytes publicKey, uint256 fee) returns(uint64 id)
func (_Ssv *SsvTransactorSession) RegisterOperator(publicKey []byte, fee *big.Int) (*types.Transaction, error) {
	return _Ssv.Contract.RegisterOperator(&_Ssv.TransactOpts, publicKey, fee)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0x06e8fb9c.
//
// Solidity: function registerValidator(bytes publicKey, uint64[] operatorIds, bytes sharesData, uint256 amount, (uint32,uint64,uint64,bool,uint256) cluster) returns()
func (_Ssv *SsvTransactor) RegisterValidator(opts *bind.TransactOpts, publicKey []byte, operatorIds []uint64, sharesData []byte, amount *big.Int, cluster ISSVNetworkCoreCluster) (*types.Transaction, error) {
	return _Ssv.contract.Transact(opts, "registerValidator", publicKey, operatorIds, sharesData, amount, cluster)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0x06e8fb9c.
//
// Solidity: function registerValidator(bytes publicKey, uint64[] operatorIds, bytes sharesData, uint256 amount, (uint32,uint64,uint64,bool,uint256) cluster) returns()
func (_Ssv *SsvSession) RegisterValidator(publicKey []byte, operatorIds []uint64, sharesData []byte, amount *big.Int, cluster ISSVNetworkCoreCluster) (*types.Transaction, error) {
	return _Ssv.Contract.RegisterValidator(&_Ssv.TransactOpts, publicKey, operatorIds, sharesData, amount, cluster)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0x06e8fb9c.
//
// Solidity: function registerValidator(bytes publicKey, uint64[] operatorIds, bytes sharesData, uint256 amount, (uint32,uint64,uint64,bool,uint256) cluster) returns()
func (_Ssv *SsvTransactorSession) RegisterValidator(publicKey []byte, operatorIds []uint64, sharesData []byte, amount *big.Int, cluster ISSVNetworkCoreCluster) (*types.Transaction, error) {
	return _Ssv.Contract.RegisterValidator(&_Ssv.TransactOpts, publicKey, operatorIds, sharesData, amount, cluster)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0x2e168e0e.
//
// Solidity: function removeOperator(uint64 operatorId) returns()
func (_Ssv *SsvTransactor) RemoveOperator(opts *bind.TransactOpts, operatorId uint64) (*types.Transaction, error) {
	return _Ssv.contract.Transact(opts, "removeOperator", operatorId)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0x2e168e0e.
//
// Solidity: function removeOperator(uint64 operatorId) returns()
func (_Ssv *SsvSession) RemoveOperator(operatorId uint64) (*types.Transaction, error) {
	return _Ssv.Contract.RemoveOperator(&_Ssv.TransactOpts, operatorId)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0x2e168e0e.
//
// Solidity: function removeOperator(uint64 operatorId) returns()
func (_Ssv *SsvTransactorSession) RemoveOperator(operatorId uint64) (*types.Transaction, error) {
	return _Ssv.Contract.RemoveOperator(&_Ssv.TransactOpts, operatorId)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x12b3fc19.
//
// Solidity: function removeValidator(bytes publicKey, uint64[] operatorIds, (uint32,uint64,uint64,bool,uint256) cluster) returns()
func (_Ssv *SsvTransactor) RemoveValidator(opts *bind.TransactOpts, publicKey []byte, operatorIds []uint64, cluster ISSVNetworkCoreCluster) (*types.Transaction, error) {
	return _Ssv.contract.Transact(opts, "removeValidator", publicKey, operatorIds, cluster)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x12b3fc19.
//
// Solidity: function removeValidator(bytes publicKey, uint64[] operatorIds, (uint32,uint64,uint64,bool,uint256) cluster) returns()
func (_Ssv *SsvSession) RemoveValidator(publicKey []byte, operatorIds []uint64, cluster ISSVNetworkCoreCluster) (*types.Transaction, error) {
	return _Ssv.Contract.RemoveValidator(&_Ssv.TransactOpts, publicKey, operatorIds, cluster)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x12b3fc19.
//
// Solidity: function removeValidator(bytes publicKey, uint64[] operatorIds, (uint32,uint64,uint64,bool,uint256) cluster) returns()
func (_Ssv *SsvTransactorSession) RemoveValidator(publicKey []byte, operatorIds []uint64, cluster ISSVNetworkCoreCluster) (*types.Transaction, error) {
	return _Ssv.Contract.RemoveValidator(&_Ssv.TransactOpts, publicKey, operatorIds, cluster)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ssv *SsvTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ssv.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ssv *SsvSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ssv.Contract.RenounceOwnership(&_Ssv.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ssv *SsvTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ssv.Contract.RenounceOwnership(&_Ssv.TransactOpts)
}

// SetFeeRecipientAddress is a paid mutator transaction binding the contract method 0xdbcdc2cc.
//
// Solidity: function setFeeRecipientAddress(address recipientAddress) returns()
func (_Ssv *SsvTransactor) SetFeeRecipientAddress(opts *bind.TransactOpts, recipientAddress common.Address) (*types.Transaction, error) {
	return _Ssv.contract.Transact(opts, "setFeeRecipientAddress", recipientAddress)
}

// SetFeeRecipientAddress is a paid mutator transaction binding the contract method 0xdbcdc2cc.
//
// Solidity: function setFeeRecipientAddress(address recipientAddress) returns()
func (_Ssv *SsvSession) SetFeeRecipientAddress(recipientAddress common.Address) (*types.Transaction, error) {
	return _Ssv.Contract.SetFeeRecipientAddress(&_Ssv.TransactOpts, recipientAddress)
}

// SetFeeRecipientAddress is a paid mutator transaction binding the contract method 0xdbcdc2cc.
//
// Solidity: function setFeeRecipientAddress(address recipientAddress) returns()
func (_Ssv *SsvTransactorSession) SetFeeRecipientAddress(recipientAddress common.Address) (*types.Transaction, error) {
	return _Ssv.Contract.SetFeeRecipientAddress(&_Ssv.TransactOpts, recipientAddress)
}

// SetOperatorWhitelist is a paid mutator transaction binding the contract method 0xc90a7eab.
//
// Solidity: function setOperatorWhitelist(uint64 operatorId, address whitelisted) returns()
func (_Ssv *SsvTransactor) SetOperatorWhitelist(opts *bind.TransactOpts, operatorId uint64, whitelisted common.Address) (*types.Transaction, error) {
	return _Ssv.contract.Transact(opts, "setOperatorWhitelist", operatorId, whitelisted)
}

// SetOperatorWhitelist is a paid mutator transaction binding the contract method 0xc90a7eab.
//
// Solidity: function setOperatorWhitelist(uint64 operatorId, address whitelisted) returns()
func (_Ssv *SsvSession) SetOperatorWhitelist(operatorId uint64, whitelisted common.Address) (*types.Transaction, error) {
	return _Ssv.Contract.SetOperatorWhitelist(&_Ssv.TransactOpts, operatorId, whitelisted)
}

// SetOperatorWhitelist is a paid mutator transaction binding the contract method 0xc90a7eab.
//
// Solidity: function setOperatorWhitelist(uint64 operatorId, address whitelisted) returns()
func (_Ssv *SsvTransactorSession) SetOperatorWhitelist(operatorId uint64, whitelisted common.Address) (*types.Transaction, error) {
	return _Ssv.Contract.SetOperatorWhitelist(&_Ssv.TransactOpts, operatorId, whitelisted)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ssv *SsvTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ssv.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ssv *SsvSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ssv.Contract.TransferOwnership(&_Ssv.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ssv *SsvTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ssv.Contract.TransferOwnership(&_Ssv.TransactOpts, newOwner)
}

// UpdateDeclareOperatorFeePeriod is a paid mutator transaction binding the contract method 0x79e3e4e4.
//
// Solidity: function updateDeclareOperatorFeePeriod(uint64 timeInSeconds) returns()
func (_Ssv *SsvTransactor) UpdateDeclareOperatorFeePeriod(opts *bind.TransactOpts, timeInSeconds uint64) (*types.Transaction, error) {
	return _Ssv.contract.Transact(opts, "updateDeclareOperatorFeePeriod", timeInSeconds)
}

// UpdateDeclareOperatorFeePeriod is a paid mutator transaction binding the contract method 0x79e3e4e4.
//
// Solidity: function updateDeclareOperatorFeePeriod(uint64 timeInSeconds) returns()
func (_Ssv *SsvSession) UpdateDeclareOperatorFeePeriod(timeInSeconds uint64) (*types.Transaction, error) {
	return _Ssv.Contract.UpdateDeclareOperatorFeePeriod(&_Ssv.TransactOpts, timeInSeconds)
}

// UpdateDeclareOperatorFeePeriod is a paid mutator transaction binding the contract method 0x79e3e4e4.
//
// Solidity: function updateDeclareOperatorFeePeriod(uint64 timeInSeconds) returns()
func (_Ssv *SsvTransactorSession) UpdateDeclareOperatorFeePeriod(timeInSeconds uint64) (*types.Transaction, error) {
	return _Ssv.Contract.UpdateDeclareOperatorFeePeriod(&_Ssv.TransactOpts, timeInSeconds)
}

// UpdateExecuteOperatorFeePeriod is a paid mutator transaction binding the contract method 0xeb608022.
//
// Solidity: function updateExecuteOperatorFeePeriod(uint64 timeInSeconds) returns()
func (_Ssv *SsvTransactor) UpdateExecuteOperatorFeePeriod(opts *bind.TransactOpts, timeInSeconds uint64) (*types.Transaction, error) {
	return _Ssv.contract.Transact(opts, "updateExecuteOperatorFeePeriod", timeInSeconds)
}

// UpdateExecuteOperatorFeePeriod is a paid mutator transaction binding the contract method 0xeb608022.
//
// Solidity: function updateExecuteOperatorFeePeriod(uint64 timeInSeconds) returns()
func (_Ssv *SsvSession) UpdateExecuteOperatorFeePeriod(timeInSeconds uint64) (*types.Transaction, error) {
	return _Ssv.Contract.UpdateExecuteOperatorFeePeriod(&_Ssv.TransactOpts, timeInSeconds)
}

// UpdateExecuteOperatorFeePeriod is a paid mutator transaction binding the contract method 0xeb608022.
//
// Solidity: function updateExecuteOperatorFeePeriod(uint64 timeInSeconds) returns()
func (_Ssv *SsvTransactorSession) UpdateExecuteOperatorFeePeriod(timeInSeconds uint64) (*types.Transaction, error) {
	return _Ssv.Contract.UpdateExecuteOperatorFeePeriod(&_Ssv.TransactOpts, timeInSeconds)
}

// UpdateLiquidationThresholdPeriod is a paid mutator transaction binding the contract method 0x6512447d.
//
// Solidity: function updateLiquidationThresholdPeriod(uint64 blocks) returns()
func (_Ssv *SsvTransactor) UpdateLiquidationThresholdPeriod(opts *bind.TransactOpts, blocks uint64) (*types.Transaction, error) {
	return _Ssv.contract.Transact(opts, "updateLiquidationThresholdPeriod", blocks)
}

// UpdateLiquidationThresholdPeriod is a paid mutator transaction binding the contract method 0x6512447d.
//
// Solidity: function updateLiquidationThresholdPeriod(uint64 blocks) returns()
func (_Ssv *SsvSession) UpdateLiquidationThresholdPeriod(blocks uint64) (*types.Transaction, error) {
	return _Ssv.Contract.UpdateLiquidationThresholdPeriod(&_Ssv.TransactOpts, blocks)
}

// UpdateLiquidationThresholdPeriod is a paid mutator transaction binding the contract method 0x6512447d.
//
// Solidity: function updateLiquidationThresholdPeriod(uint64 blocks) returns()
func (_Ssv *SsvTransactorSession) UpdateLiquidationThresholdPeriod(blocks uint64) (*types.Transaction, error) {
	return _Ssv.Contract.UpdateLiquidationThresholdPeriod(&_Ssv.TransactOpts, blocks)
}

// UpdateMaximumOperatorFee is a paid mutator transaction binding the contract method 0xe39c6744.
//
// Solidity: function updateMaximumOperatorFee(uint64 maxFee) returns()
func (_Ssv *SsvTransactor) UpdateMaximumOperatorFee(opts *bind.TransactOpts, maxFee uint64) (*types.Transaction, error) {
	return _Ssv.contract.Transact(opts, "updateMaximumOperatorFee", maxFee)
}

// UpdateMaximumOperatorFee is a paid mutator transaction binding the contract method 0xe39c6744.
//
// Solidity: function updateMaximumOperatorFee(uint64 maxFee) returns()
func (_Ssv *SsvSession) UpdateMaximumOperatorFee(maxFee uint64) (*types.Transaction, error) {
	return _Ssv.Contract.UpdateMaximumOperatorFee(&_Ssv.TransactOpts, maxFee)
}

// UpdateMaximumOperatorFee is a paid mutator transaction binding the contract method 0xe39c6744.
//
// Solidity: function updateMaximumOperatorFee(uint64 maxFee) returns()
func (_Ssv *SsvTransactorSession) UpdateMaximumOperatorFee(maxFee uint64) (*types.Transaction, error) {
	return _Ssv.Contract.UpdateMaximumOperatorFee(&_Ssv.TransactOpts, maxFee)
}

// UpdateMinimumLiquidationCollateral is a paid mutator transaction binding the contract method 0xb4c9c408.
//
// Solidity: function updateMinimumLiquidationCollateral(uint256 amount) returns()
func (_Ssv *SsvTransactor) UpdateMinimumLiquidationCollateral(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Ssv.contract.Transact(opts, "updateMinimumLiquidationCollateral", amount)
}

// UpdateMinimumLiquidationCollateral is a paid mutator transaction binding the contract method 0xb4c9c408.
//
// Solidity: function updateMinimumLiquidationCollateral(uint256 amount) returns()
func (_Ssv *SsvSession) UpdateMinimumLiquidationCollateral(amount *big.Int) (*types.Transaction, error) {
	return _Ssv.Contract.UpdateMinimumLiquidationCollateral(&_Ssv.TransactOpts, amount)
}

// UpdateMinimumLiquidationCollateral is a paid mutator transaction binding the contract method 0xb4c9c408.
//
// Solidity: function updateMinimumLiquidationCollateral(uint256 amount) returns()
func (_Ssv *SsvTransactorSession) UpdateMinimumLiquidationCollateral(amount *big.Int) (*types.Transaction, error) {
	return _Ssv.Contract.UpdateMinimumLiquidationCollateral(&_Ssv.TransactOpts, amount)
}

// UpdateModule is a paid mutator transaction binding the contract method 0xe3e324b0.
//
// Solidity: function updateModule(uint8 moduleId, address moduleAddress) returns()
func (_Ssv *SsvTransactor) UpdateModule(opts *bind.TransactOpts, moduleId uint8, moduleAddress common.Address) (*types.Transaction, error) {
	return _Ssv.contract.Transact(opts, "updateModule", moduleId, moduleAddress)
}

// UpdateModule is a paid mutator transaction binding the contract method 0xe3e324b0.
//
// Solidity: function updateModule(uint8 moduleId, address moduleAddress) returns()
func (_Ssv *SsvSession) UpdateModule(moduleId uint8, moduleAddress common.Address) (*types.Transaction, error) {
	return _Ssv.Contract.UpdateModule(&_Ssv.TransactOpts, moduleId, moduleAddress)
}

// UpdateModule is a paid mutator transaction binding the contract method 0xe3e324b0.
//
// Solidity: function updateModule(uint8 moduleId, address moduleAddress) returns()
func (_Ssv *SsvTransactorSession) UpdateModule(moduleId uint8, moduleAddress common.Address) (*types.Transaction, error) {
	return _Ssv.Contract.UpdateModule(&_Ssv.TransactOpts, moduleId, moduleAddress)
}

// UpdateNetworkFee is a paid mutator transaction binding the contract method 0x1f1f9fd5.
//
// Solidity: function updateNetworkFee(uint256 fee) returns()
func (_Ssv *SsvTransactor) UpdateNetworkFee(opts *bind.TransactOpts, fee *big.Int) (*types.Transaction, error) {
	return _Ssv.contract.Transact(opts, "updateNetworkFee", fee)
}

// UpdateNetworkFee is a paid mutator transaction binding the contract method 0x1f1f9fd5.
//
// Solidity: function updateNetworkFee(uint256 fee) returns()
func (_Ssv *SsvSession) UpdateNetworkFee(fee *big.Int) (*types.Transaction, error) {
	return _Ssv.Contract.UpdateNetworkFee(&_Ssv.TransactOpts, fee)
}

// UpdateNetworkFee is a paid mutator transaction binding the contract method 0x1f1f9fd5.
//
// Solidity: function updateNetworkFee(uint256 fee) returns()
func (_Ssv *SsvTransactorSession) UpdateNetworkFee(fee *big.Int) (*types.Transaction, error) {
	return _Ssv.Contract.UpdateNetworkFee(&_Ssv.TransactOpts, fee)
}

// UpdateOperatorFeeIncreaseLimit is a paid mutator transaction binding the contract method 0x3631983f.
//
// Solidity: function updateOperatorFeeIncreaseLimit(uint64 percentage) returns()
func (_Ssv *SsvTransactor) UpdateOperatorFeeIncreaseLimit(opts *bind.TransactOpts, percentage uint64) (*types.Transaction, error) {
	return _Ssv.contract.Transact(opts, "updateOperatorFeeIncreaseLimit", percentage)
}

// UpdateOperatorFeeIncreaseLimit is a paid mutator transaction binding the contract method 0x3631983f.
//
// Solidity: function updateOperatorFeeIncreaseLimit(uint64 percentage) returns()
func (_Ssv *SsvSession) UpdateOperatorFeeIncreaseLimit(percentage uint64) (*types.Transaction, error) {
	return _Ssv.Contract.UpdateOperatorFeeIncreaseLimit(&_Ssv.TransactOpts, percentage)
}

// UpdateOperatorFeeIncreaseLimit is a paid mutator transaction binding the contract method 0x3631983f.
//
// Solidity: function updateOperatorFeeIncreaseLimit(uint64 percentage) returns()
func (_Ssv *SsvTransactorSession) UpdateOperatorFeeIncreaseLimit(percentage uint64) (*types.Transaction, error) {
	return _Ssv.Contract.UpdateOperatorFeeIncreaseLimit(&_Ssv.TransactOpts, percentage)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Ssv *SsvTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Ssv.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Ssv *SsvSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Ssv.Contract.UpgradeTo(&_Ssv.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Ssv *SsvTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Ssv.Contract.UpgradeTo(&_Ssv.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Ssv *SsvTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Ssv.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Ssv *SsvSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Ssv.Contract.UpgradeToAndCall(&_Ssv.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Ssv *SsvTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Ssv.Contract.UpgradeToAndCall(&_Ssv.TransactOpts, newImplementation, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x686e682c.
//
// Solidity: function withdraw(uint64[] operatorIds, uint256 amount, (uint32,uint64,uint64,bool,uint256) cluster) returns()
func (_Ssv *SsvTransactor) Withdraw(opts *bind.TransactOpts, operatorIds []uint64, amount *big.Int, cluster ISSVNetworkCoreCluster) (*types.Transaction, error) {
	return _Ssv.contract.Transact(opts, "withdraw", operatorIds, amount, cluster)
}

// Withdraw is a paid mutator transaction binding the contract method 0x686e682c.
//
// Solidity: function withdraw(uint64[] operatorIds, uint256 amount, (uint32,uint64,uint64,bool,uint256) cluster) returns()
func (_Ssv *SsvSession) Withdraw(operatorIds []uint64, amount *big.Int, cluster ISSVNetworkCoreCluster) (*types.Transaction, error) {
	return _Ssv.Contract.Withdraw(&_Ssv.TransactOpts, operatorIds, amount, cluster)
}

// Withdraw is a paid mutator transaction binding the contract method 0x686e682c.
//
// Solidity: function withdraw(uint64[] operatorIds, uint256 amount, (uint32,uint64,uint64,bool,uint256) cluster) returns()
func (_Ssv *SsvTransactorSession) Withdraw(operatorIds []uint64, amount *big.Int, cluster ISSVNetworkCoreCluster) (*types.Transaction, error) {
	return _Ssv.Contract.Withdraw(&_Ssv.TransactOpts, operatorIds, amount, cluster)
}

// WithdrawAllOperatorEarnings is a paid mutator transaction binding the contract method 0x4bc93b64.
//
// Solidity: function withdrawAllOperatorEarnings(uint64 operatorId) returns()
func (_Ssv *SsvTransactor) WithdrawAllOperatorEarnings(opts *bind.TransactOpts, operatorId uint64) (*types.Transaction, error) {
	return _Ssv.contract.Transact(opts, "withdrawAllOperatorEarnings", operatorId)
}

// WithdrawAllOperatorEarnings is a paid mutator transaction binding the contract method 0x4bc93b64.
//
// Solidity: function withdrawAllOperatorEarnings(uint64 operatorId) returns()
func (_Ssv *SsvSession) WithdrawAllOperatorEarnings(operatorId uint64) (*types.Transaction, error) {
	return _Ssv.Contract.WithdrawAllOperatorEarnings(&_Ssv.TransactOpts, operatorId)
}

// WithdrawAllOperatorEarnings is a paid mutator transaction binding the contract method 0x4bc93b64.
//
// Solidity: function withdrawAllOperatorEarnings(uint64 operatorId) returns()
func (_Ssv *SsvTransactorSession) WithdrawAllOperatorEarnings(operatorId uint64) (*types.Transaction, error) {
	return _Ssv.Contract.WithdrawAllOperatorEarnings(&_Ssv.TransactOpts, operatorId)
}

// WithdrawNetworkEarnings is a paid mutator transaction binding the contract method 0xd2231741.
//
// Solidity: function withdrawNetworkEarnings(uint256 amount) returns()
func (_Ssv *SsvTransactor) WithdrawNetworkEarnings(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Ssv.contract.Transact(opts, "withdrawNetworkEarnings", amount)
}

// WithdrawNetworkEarnings is a paid mutator transaction binding the contract method 0xd2231741.
//
// Solidity: function withdrawNetworkEarnings(uint256 amount) returns()
func (_Ssv *SsvSession) WithdrawNetworkEarnings(amount *big.Int) (*types.Transaction, error) {
	return _Ssv.Contract.WithdrawNetworkEarnings(&_Ssv.TransactOpts, amount)
}

// WithdrawNetworkEarnings is a paid mutator transaction binding the contract method 0xd2231741.
//
// Solidity: function withdrawNetworkEarnings(uint256 amount) returns()
func (_Ssv *SsvTransactorSession) WithdrawNetworkEarnings(amount *big.Int) (*types.Transaction, error) {
	return _Ssv.Contract.WithdrawNetworkEarnings(&_Ssv.TransactOpts, amount)
}

// WithdrawOperatorEarnings is a paid mutator transaction binding the contract method 0x35f63767.
//
// Solidity: function withdrawOperatorEarnings(uint64 operatorId, uint256 amount) returns()
func (_Ssv *SsvTransactor) WithdrawOperatorEarnings(opts *bind.TransactOpts, operatorId uint64, amount *big.Int) (*types.Transaction, error) {
	return _Ssv.contract.Transact(opts, "withdrawOperatorEarnings", operatorId, amount)
}

// WithdrawOperatorEarnings is a paid mutator transaction binding the contract method 0x35f63767.
//
// Solidity: function withdrawOperatorEarnings(uint64 operatorId, uint256 amount) returns()
func (_Ssv *SsvSession) WithdrawOperatorEarnings(operatorId uint64, amount *big.Int) (*types.Transaction, error) {
	return _Ssv.Contract.WithdrawOperatorEarnings(&_Ssv.TransactOpts, operatorId, amount)
}

// WithdrawOperatorEarnings is a paid mutator transaction binding the contract method 0x35f63767.
//
// Solidity: function withdrawOperatorEarnings(uint64 operatorId, uint256 amount) returns()
func (_Ssv *SsvTransactorSession) WithdrawOperatorEarnings(operatorId uint64, amount *big.Int) (*types.Transaction, error) {
	return _Ssv.Contract.WithdrawOperatorEarnings(&_Ssv.TransactOpts, operatorId, amount)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_Ssv *SsvTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Ssv.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_Ssv *SsvSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Ssv.Contract.Fallback(&_Ssv.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_Ssv *SsvTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Ssv.Contract.Fallback(&_Ssv.TransactOpts, calldata)
}

// SsvAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Ssv contract.
type SsvAdminChangedIterator struct {
	Event *SsvAdminChanged // Event containing the contract specifics and raw log

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
func (it *SsvAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvAdminChanged)
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
		it.Event = new(SsvAdminChanged)
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
func (it *SsvAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvAdminChanged represents a AdminChanged event raised by the Ssv contract.
type SsvAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Ssv *SsvFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*SsvAdminChangedIterator, error) {

	logs, sub, err := _Ssv.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &SsvAdminChangedIterator{contract: _Ssv.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Ssv *SsvFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *SsvAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Ssv.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvAdminChanged)
				if err := _Ssv.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_Ssv *SsvFilterer) ParseAdminChanged(log types.Log) (*SsvAdminChanged, error) {
	event := new(SsvAdminChanged)
	if err := _Ssv.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsvBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the Ssv contract.
type SsvBeaconUpgradedIterator struct {
	Event *SsvBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *SsvBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvBeaconUpgraded)
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
		it.Event = new(SsvBeaconUpgraded)
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
func (it *SsvBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvBeaconUpgraded represents a BeaconUpgraded event raised by the Ssv contract.
type SsvBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Ssv *SsvFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*SsvBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Ssv.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &SsvBeaconUpgradedIterator{contract: _Ssv.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Ssv *SsvFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *SsvBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Ssv.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvBeaconUpgraded)
				if err := _Ssv.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_Ssv *SsvFilterer) ParseBeaconUpgraded(log types.Log) (*SsvBeaconUpgraded, error) {
	event := new(SsvBeaconUpgraded)
	if err := _Ssv.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsvClusterDepositedIterator is returned from FilterClusterDeposited and is used to iterate over the raw logs and unpacked data for ClusterDeposited events raised by the Ssv contract.
type SsvClusterDepositedIterator struct {
	Event *SsvClusterDeposited // Event containing the contract specifics and raw log

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
func (it *SsvClusterDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvClusterDeposited)
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
		it.Event = new(SsvClusterDeposited)
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
func (it *SsvClusterDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvClusterDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvClusterDeposited represents a ClusterDeposited event raised by the Ssv contract.
type SsvClusterDeposited struct {
	Owner       common.Address
	OperatorIds []uint64
	Value       *big.Int
	Cluster     ISSVNetworkCoreCluster
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterClusterDeposited is a free log retrieval operation binding the contract event 0x2bac1912f2481d12f0df08647c06bee174967c62d3a03cbc078eb215dc1bd9a2.
//
// Solidity: event ClusterDeposited(address indexed owner, uint64[] operatorIds, uint256 value, (uint32,uint64,uint64,bool,uint256) cluster)
func (_Ssv *SsvFilterer) FilterClusterDeposited(opts *bind.FilterOpts, owner []common.Address) (*SsvClusterDepositedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Ssv.contract.FilterLogs(opts, "ClusterDeposited", ownerRule)
	if err != nil {
		return nil, err
	}
	return &SsvClusterDepositedIterator{contract: _Ssv.contract, event: "ClusterDeposited", logs: logs, sub: sub}, nil
}

// WatchClusterDeposited is a free log subscription operation binding the contract event 0x2bac1912f2481d12f0df08647c06bee174967c62d3a03cbc078eb215dc1bd9a2.
//
// Solidity: event ClusterDeposited(address indexed owner, uint64[] operatorIds, uint256 value, (uint32,uint64,uint64,bool,uint256) cluster)
func (_Ssv *SsvFilterer) WatchClusterDeposited(opts *bind.WatchOpts, sink chan<- *SsvClusterDeposited, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Ssv.contract.WatchLogs(opts, "ClusterDeposited", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvClusterDeposited)
				if err := _Ssv.contract.UnpackLog(event, "ClusterDeposited", log); err != nil {
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

// ParseClusterDeposited is a log parse operation binding the contract event 0x2bac1912f2481d12f0df08647c06bee174967c62d3a03cbc078eb215dc1bd9a2.
//
// Solidity: event ClusterDeposited(address indexed owner, uint64[] operatorIds, uint256 value, (uint32,uint64,uint64,bool,uint256) cluster)
func (_Ssv *SsvFilterer) ParseClusterDeposited(log types.Log) (*SsvClusterDeposited, error) {
	event := new(SsvClusterDeposited)
	if err := _Ssv.contract.UnpackLog(event, "ClusterDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsvClusterLiquidatedIterator is returned from FilterClusterLiquidated and is used to iterate over the raw logs and unpacked data for ClusterLiquidated events raised by the Ssv contract.
type SsvClusterLiquidatedIterator struct {
	Event *SsvClusterLiquidated // Event containing the contract specifics and raw log

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
func (it *SsvClusterLiquidatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvClusterLiquidated)
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
		it.Event = new(SsvClusterLiquidated)
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
func (it *SsvClusterLiquidatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvClusterLiquidatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvClusterLiquidated represents a ClusterLiquidated event raised by the Ssv contract.
type SsvClusterLiquidated struct {
	Owner       common.Address
	OperatorIds []uint64
	Cluster     ISSVNetworkCoreCluster
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterClusterLiquidated is a free log retrieval operation binding the contract event 0x1fce24c373e07f89214e9187598635036111dbb363e99f4ce498488cdc66e688.
//
// Solidity: event ClusterLiquidated(address indexed owner, uint64[] operatorIds, (uint32,uint64,uint64,bool,uint256) cluster)
func (_Ssv *SsvFilterer) FilterClusterLiquidated(opts *bind.FilterOpts, owner []common.Address) (*SsvClusterLiquidatedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Ssv.contract.FilterLogs(opts, "ClusterLiquidated", ownerRule)
	if err != nil {
		return nil, err
	}
	return &SsvClusterLiquidatedIterator{contract: _Ssv.contract, event: "ClusterLiquidated", logs: logs, sub: sub}, nil
}

// WatchClusterLiquidated is a free log subscription operation binding the contract event 0x1fce24c373e07f89214e9187598635036111dbb363e99f4ce498488cdc66e688.
//
// Solidity: event ClusterLiquidated(address indexed owner, uint64[] operatorIds, (uint32,uint64,uint64,bool,uint256) cluster)
func (_Ssv *SsvFilterer) WatchClusterLiquidated(opts *bind.WatchOpts, sink chan<- *SsvClusterLiquidated, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Ssv.contract.WatchLogs(opts, "ClusterLiquidated", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvClusterLiquidated)
				if err := _Ssv.contract.UnpackLog(event, "ClusterLiquidated", log); err != nil {
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

// ParseClusterLiquidated is a log parse operation binding the contract event 0x1fce24c373e07f89214e9187598635036111dbb363e99f4ce498488cdc66e688.
//
// Solidity: event ClusterLiquidated(address indexed owner, uint64[] operatorIds, (uint32,uint64,uint64,bool,uint256) cluster)
func (_Ssv *SsvFilterer) ParseClusterLiquidated(log types.Log) (*SsvClusterLiquidated, error) {
	event := new(SsvClusterLiquidated)
	if err := _Ssv.contract.UnpackLog(event, "ClusterLiquidated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsvClusterReactivatedIterator is returned from FilterClusterReactivated and is used to iterate over the raw logs and unpacked data for ClusterReactivated events raised by the Ssv contract.
type SsvClusterReactivatedIterator struct {
	Event *SsvClusterReactivated // Event containing the contract specifics and raw log

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
func (it *SsvClusterReactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvClusterReactivated)
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
		it.Event = new(SsvClusterReactivated)
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
func (it *SsvClusterReactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvClusterReactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvClusterReactivated represents a ClusterReactivated event raised by the Ssv contract.
type SsvClusterReactivated struct {
	Owner       common.Address
	OperatorIds []uint64
	Cluster     ISSVNetworkCoreCluster
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterClusterReactivated is a free log retrieval operation binding the contract event 0xc803f8c01343fcdaf32068f4c283951623ef2b3fa0c547551931356f456b6859.
//
// Solidity: event ClusterReactivated(address indexed owner, uint64[] operatorIds, (uint32,uint64,uint64,bool,uint256) cluster)
func (_Ssv *SsvFilterer) FilterClusterReactivated(opts *bind.FilterOpts, owner []common.Address) (*SsvClusterReactivatedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Ssv.contract.FilterLogs(opts, "ClusterReactivated", ownerRule)
	if err != nil {
		return nil, err
	}
	return &SsvClusterReactivatedIterator{contract: _Ssv.contract, event: "ClusterReactivated", logs: logs, sub: sub}, nil
}

// WatchClusterReactivated is a free log subscription operation binding the contract event 0xc803f8c01343fcdaf32068f4c283951623ef2b3fa0c547551931356f456b6859.
//
// Solidity: event ClusterReactivated(address indexed owner, uint64[] operatorIds, (uint32,uint64,uint64,bool,uint256) cluster)
func (_Ssv *SsvFilterer) WatchClusterReactivated(opts *bind.WatchOpts, sink chan<- *SsvClusterReactivated, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Ssv.contract.WatchLogs(opts, "ClusterReactivated", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvClusterReactivated)
				if err := _Ssv.contract.UnpackLog(event, "ClusterReactivated", log); err != nil {
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

// ParseClusterReactivated is a log parse operation binding the contract event 0xc803f8c01343fcdaf32068f4c283951623ef2b3fa0c547551931356f456b6859.
//
// Solidity: event ClusterReactivated(address indexed owner, uint64[] operatorIds, (uint32,uint64,uint64,bool,uint256) cluster)
func (_Ssv *SsvFilterer) ParseClusterReactivated(log types.Log) (*SsvClusterReactivated, error) {
	event := new(SsvClusterReactivated)
	if err := _Ssv.contract.UnpackLog(event, "ClusterReactivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsvClusterWithdrawnIterator is returned from FilterClusterWithdrawn and is used to iterate over the raw logs and unpacked data for ClusterWithdrawn events raised by the Ssv contract.
type SsvClusterWithdrawnIterator struct {
	Event *SsvClusterWithdrawn // Event containing the contract specifics and raw log

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
func (it *SsvClusterWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvClusterWithdrawn)
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
		it.Event = new(SsvClusterWithdrawn)
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
func (it *SsvClusterWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvClusterWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvClusterWithdrawn represents a ClusterWithdrawn event raised by the Ssv contract.
type SsvClusterWithdrawn struct {
	Owner       common.Address
	OperatorIds []uint64
	Value       *big.Int
	Cluster     ISSVNetworkCoreCluster
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterClusterWithdrawn is a free log retrieval operation binding the contract event 0x39d1320bbda24947e77f3560661323384aa0a1cb9d5e040e617e5cbf50b6dbe0.
//
// Solidity: event ClusterWithdrawn(address indexed owner, uint64[] operatorIds, uint256 value, (uint32,uint64,uint64,bool,uint256) cluster)
func (_Ssv *SsvFilterer) FilterClusterWithdrawn(opts *bind.FilterOpts, owner []common.Address) (*SsvClusterWithdrawnIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Ssv.contract.FilterLogs(opts, "ClusterWithdrawn", ownerRule)
	if err != nil {
		return nil, err
	}
	return &SsvClusterWithdrawnIterator{contract: _Ssv.contract, event: "ClusterWithdrawn", logs: logs, sub: sub}, nil
}

// WatchClusterWithdrawn is a free log subscription operation binding the contract event 0x39d1320bbda24947e77f3560661323384aa0a1cb9d5e040e617e5cbf50b6dbe0.
//
// Solidity: event ClusterWithdrawn(address indexed owner, uint64[] operatorIds, uint256 value, (uint32,uint64,uint64,bool,uint256) cluster)
func (_Ssv *SsvFilterer) WatchClusterWithdrawn(opts *bind.WatchOpts, sink chan<- *SsvClusterWithdrawn, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Ssv.contract.WatchLogs(opts, "ClusterWithdrawn", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvClusterWithdrawn)
				if err := _Ssv.contract.UnpackLog(event, "ClusterWithdrawn", log); err != nil {
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

// ParseClusterWithdrawn is a log parse operation binding the contract event 0x39d1320bbda24947e77f3560661323384aa0a1cb9d5e040e617e5cbf50b6dbe0.
//
// Solidity: event ClusterWithdrawn(address indexed owner, uint64[] operatorIds, uint256 value, (uint32,uint64,uint64,bool,uint256) cluster)
func (_Ssv *SsvFilterer) ParseClusterWithdrawn(log types.Log) (*SsvClusterWithdrawn, error) {
	event := new(SsvClusterWithdrawn)
	if err := _Ssv.contract.UnpackLog(event, "ClusterWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsvDeclareOperatorFeePeriodUpdatedIterator is returned from FilterDeclareOperatorFeePeriodUpdated and is used to iterate over the raw logs and unpacked data for DeclareOperatorFeePeriodUpdated events raised by the Ssv contract.
type SsvDeclareOperatorFeePeriodUpdatedIterator struct {
	Event *SsvDeclareOperatorFeePeriodUpdated // Event containing the contract specifics and raw log

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
func (it *SsvDeclareOperatorFeePeriodUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvDeclareOperatorFeePeriodUpdated)
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
		it.Event = new(SsvDeclareOperatorFeePeriodUpdated)
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
func (it *SsvDeclareOperatorFeePeriodUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvDeclareOperatorFeePeriodUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvDeclareOperatorFeePeriodUpdated represents a DeclareOperatorFeePeriodUpdated event raised by the Ssv contract.
type SsvDeclareOperatorFeePeriodUpdated struct {
	Value uint64
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterDeclareOperatorFeePeriodUpdated is a free log retrieval operation binding the contract event 0x5fbd75d987b37490f91aa1909db948e7ff14c6ffb495b2f8e0b2334da9b192f1.
//
// Solidity: event DeclareOperatorFeePeriodUpdated(uint64 value)
func (_Ssv *SsvFilterer) FilterDeclareOperatorFeePeriodUpdated(opts *bind.FilterOpts) (*SsvDeclareOperatorFeePeriodUpdatedIterator, error) {

	logs, sub, err := _Ssv.contract.FilterLogs(opts, "DeclareOperatorFeePeriodUpdated")
	if err != nil {
		return nil, err
	}
	return &SsvDeclareOperatorFeePeriodUpdatedIterator{contract: _Ssv.contract, event: "DeclareOperatorFeePeriodUpdated", logs: logs, sub: sub}, nil
}

// WatchDeclareOperatorFeePeriodUpdated is a free log subscription operation binding the contract event 0x5fbd75d987b37490f91aa1909db948e7ff14c6ffb495b2f8e0b2334da9b192f1.
//
// Solidity: event DeclareOperatorFeePeriodUpdated(uint64 value)
func (_Ssv *SsvFilterer) WatchDeclareOperatorFeePeriodUpdated(opts *bind.WatchOpts, sink chan<- *SsvDeclareOperatorFeePeriodUpdated) (event.Subscription, error) {

	logs, sub, err := _Ssv.contract.WatchLogs(opts, "DeclareOperatorFeePeriodUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvDeclareOperatorFeePeriodUpdated)
				if err := _Ssv.contract.UnpackLog(event, "DeclareOperatorFeePeriodUpdated", log); err != nil {
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

// ParseDeclareOperatorFeePeriodUpdated is a log parse operation binding the contract event 0x5fbd75d987b37490f91aa1909db948e7ff14c6ffb495b2f8e0b2334da9b192f1.
//
// Solidity: event DeclareOperatorFeePeriodUpdated(uint64 value)
func (_Ssv *SsvFilterer) ParseDeclareOperatorFeePeriodUpdated(log types.Log) (*SsvDeclareOperatorFeePeriodUpdated, error) {
	event := new(SsvDeclareOperatorFeePeriodUpdated)
	if err := _Ssv.contract.UnpackLog(event, "DeclareOperatorFeePeriodUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsvExecuteOperatorFeePeriodUpdatedIterator is returned from FilterExecuteOperatorFeePeriodUpdated and is used to iterate over the raw logs and unpacked data for ExecuteOperatorFeePeriodUpdated events raised by the Ssv contract.
type SsvExecuteOperatorFeePeriodUpdatedIterator struct {
	Event *SsvExecuteOperatorFeePeriodUpdated // Event containing the contract specifics and raw log

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
func (it *SsvExecuteOperatorFeePeriodUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvExecuteOperatorFeePeriodUpdated)
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
		it.Event = new(SsvExecuteOperatorFeePeriodUpdated)
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
func (it *SsvExecuteOperatorFeePeriodUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvExecuteOperatorFeePeriodUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvExecuteOperatorFeePeriodUpdated represents a ExecuteOperatorFeePeriodUpdated event raised by the Ssv contract.
type SsvExecuteOperatorFeePeriodUpdated struct {
	Value uint64
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterExecuteOperatorFeePeriodUpdated is a free log retrieval operation binding the contract event 0xf6b8a2b45d0a60381de51a7b980c4660d9e5b82db6e07a4d342bfc17a6ff96bf.
//
// Solidity: event ExecuteOperatorFeePeriodUpdated(uint64 value)
func (_Ssv *SsvFilterer) FilterExecuteOperatorFeePeriodUpdated(opts *bind.FilterOpts) (*SsvExecuteOperatorFeePeriodUpdatedIterator, error) {

	logs, sub, err := _Ssv.contract.FilterLogs(opts, "ExecuteOperatorFeePeriodUpdated")
	if err != nil {
		return nil, err
	}
	return &SsvExecuteOperatorFeePeriodUpdatedIterator{contract: _Ssv.contract, event: "ExecuteOperatorFeePeriodUpdated", logs: logs, sub: sub}, nil
}

// WatchExecuteOperatorFeePeriodUpdated is a free log subscription operation binding the contract event 0xf6b8a2b45d0a60381de51a7b980c4660d9e5b82db6e07a4d342bfc17a6ff96bf.
//
// Solidity: event ExecuteOperatorFeePeriodUpdated(uint64 value)
func (_Ssv *SsvFilterer) WatchExecuteOperatorFeePeriodUpdated(opts *bind.WatchOpts, sink chan<- *SsvExecuteOperatorFeePeriodUpdated) (event.Subscription, error) {

	logs, sub, err := _Ssv.contract.WatchLogs(opts, "ExecuteOperatorFeePeriodUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvExecuteOperatorFeePeriodUpdated)
				if err := _Ssv.contract.UnpackLog(event, "ExecuteOperatorFeePeriodUpdated", log); err != nil {
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

// ParseExecuteOperatorFeePeriodUpdated is a log parse operation binding the contract event 0xf6b8a2b45d0a60381de51a7b980c4660d9e5b82db6e07a4d342bfc17a6ff96bf.
//
// Solidity: event ExecuteOperatorFeePeriodUpdated(uint64 value)
func (_Ssv *SsvFilterer) ParseExecuteOperatorFeePeriodUpdated(log types.Log) (*SsvExecuteOperatorFeePeriodUpdated, error) {
	event := new(SsvExecuteOperatorFeePeriodUpdated)
	if err := _Ssv.contract.UnpackLog(event, "ExecuteOperatorFeePeriodUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsvFeeRecipientAddressUpdatedIterator is returned from FilterFeeRecipientAddressUpdated and is used to iterate over the raw logs and unpacked data for FeeRecipientAddressUpdated events raised by the Ssv contract.
type SsvFeeRecipientAddressUpdatedIterator struct {
	Event *SsvFeeRecipientAddressUpdated // Event containing the contract specifics and raw log

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
func (it *SsvFeeRecipientAddressUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvFeeRecipientAddressUpdated)
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
		it.Event = new(SsvFeeRecipientAddressUpdated)
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
func (it *SsvFeeRecipientAddressUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvFeeRecipientAddressUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvFeeRecipientAddressUpdated represents a FeeRecipientAddressUpdated event raised by the Ssv contract.
type SsvFeeRecipientAddressUpdated struct {
	Owner            common.Address
	RecipientAddress common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterFeeRecipientAddressUpdated is a free log retrieval operation binding the contract event 0x259235c230d57def1521657e7c7951d3b385e76193378bc87ef6b56bc2ec3548.
//
// Solidity: event FeeRecipientAddressUpdated(address indexed owner, address recipientAddress)
func (_Ssv *SsvFilterer) FilterFeeRecipientAddressUpdated(opts *bind.FilterOpts, owner []common.Address) (*SsvFeeRecipientAddressUpdatedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Ssv.contract.FilterLogs(opts, "FeeRecipientAddressUpdated", ownerRule)
	if err != nil {
		return nil, err
	}
	return &SsvFeeRecipientAddressUpdatedIterator{contract: _Ssv.contract, event: "FeeRecipientAddressUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeRecipientAddressUpdated is a free log subscription operation binding the contract event 0x259235c230d57def1521657e7c7951d3b385e76193378bc87ef6b56bc2ec3548.
//
// Solidity: event FeeRecipientAddressUpdated(address indexed owner, address recipientAddress)
func (_Ssv *SsvFilterer) WatchFeeRecipientAddressUpdated(opts *bind.WatchOpts, sink chan<- *SsvFeeRecipientAddressUpdated, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Ssv.contract.WatchLogs(opts, "FeeRecipientAddressUpdated", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvFeeRecipientAddressUpdated)
				if err := _Ssv.contract.UnpackLog(event, "FeeRecipientAddressUpdated", log); err != nil {
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

// ParseFeeRecipientAddressUpdated is a log parse operation binding the contract event 0x259235c230d57def1521657e7c7951d3b385e76193378bc87ef6b56bc2ec3548.
//
// Solidity: event FeeRecipientAddressUpdated(address indexed owner, address recipientAddress)
func (_Ssv *SsvFilterer) ParseFeeRecipientAddressUpdated(log types.Log) (*SsvFeeRecipientAddressUpdated, error) {
	event := new(SsvFeeRecipientAddressUpdated)
	if err := _Ssv.contract.UnpackLog(event, "FeeRecipientAddressUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsvInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Ssv contract.
type SsvInitializedIterator struct {
	Event *SsvInitialized // Event containing the contract specifics and raw log

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
func (it *SsvInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvInitialized)
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
		it.Event = new(SsvInitialized)
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
func (it *SsvInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvInitialized represents a Initialized event raised by the Ssv contract.
type SsvInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Ssv *SsvFilterer) FilterInitialized(opts *bind.FilterOpts) (*SsvInitializedIterator, error) {

	logs, sub, err := _Ssv.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SsvInitializedIterator{contract: _Ssv.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Ssv *SsvFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SsvInitialized) (event.Subscription, error) {

	logs, sub, err := _Ssv.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvInitialized)
				if err := _Ssv.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Ssv *SsvFilterer) ParseInitialized(log types.Log) (*SsvInitialized, error) {
	event := new(SsvInitialized)
	if err := _Ssv.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsvLiquidationThresholdPeriodUpdatedIterator is returned from FilterLiquidationThresholdPeriodUpdated and is used to iterate over the raw logs and unpacked data for LiquidationThresholdPeriodUpdated events raised by the Ssv contract.
type SsvLiquidationThresholdPeriodUpdatedIterator struct {
	Event *SsvLiquidationThresholdPeriodUpdated // Event containing the contract specifics and raw log

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
func (it *SsvLiquidationThresholdPeriodUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvLiquidationThresholdPeriodUpdated)
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
		it.Event = new(SsvLiquidationThresholdPeriodUpdated)
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
func (it *SsvLiquidationThresholdPeriodUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvLiquidationThresholdPeriodUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvLiquidationThresholdPeriodUpdated represents a LiquidationThresholdPeriodUpdated event raised by the Ssv contract.
type SsvLiquidationThresholdPeriodUpdated struct {
	Value uint64
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLiquidationThresholdPeriodUpdated is a free log retrieval operation binding the contract event 0x42af14411036d7a50e5e92daf825781450fc8fac8fb65cbdb04720ff08efb84f.
//
// Solidity: event LiquidationThresholdPeriodUpdated(uint64 value)
func (_Ssv *SsvFilterer) FilterLiquidationThresholdPeriodUpdated(opts *bind.FilterOpts) (*SsvLiquidationThresholdPeriodUpdatedIterator, error) {

	logs, sub, err := _Ssv.contract.FilterLogs(opts, "LiquidationThresholdPeriodUpdated")
	if err != nil {
		return nil, err
	}
	return &SsvLiquidationThresholdPeriodUpdatedIterator{contract: _Ssv.contract, event: "LiquidationThresholdPeriodUpdated", logs: logs, sub: sub}, nil
}

// WatchLiquidationThresholdPeriodUpdated is a free log subscription operation binding the contract event 0x42af14411036d7a50e5e92daf825781450fc8fac8fb65cbdb04720ff08efb84f.
//
// Solidity: event LiquidationThresholdPeriodUpdated(uint64 value)
func (_Ssv *SsvFilterer) WatchLiquidationThresholdPeriodUpdated(opts *bind.WatchOpts, sink chan<- *SsvLiquidationThresholdPeriodUpdated) (event.Subscription, error) {

	logs, sub, err := _Ssv.contract.WatchLogs(opts, "LiquidationThresholdPeriodUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvLiquidationThresholdPeriodUpdated)
				if err := _Ssv.contract.UnpackLog(event, "LiquidationThresholdPeriodUpdated", log); err != nil {
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

// ParseLiquidationThresholdPeriodUpdated is a log parse operation binding the contract event 0x42af14411036d7a50e5e92daf825781450fc8fac8fb65cbdb04720ff08efb84f.
//
// Solidity: event LiquidationThresholdPeriodUpdated(uint64 value)
func (_Ssv *SsvFilterer) ParseLiquidationThresholdPeriodUpdated(log types.Log) (*SsvLiquidationThresholdPeriodUpdated, error) {
	event := new(SsvLiquidationThresholdPeriodUpdated)
	if err := _Ssv.contract.UnpackLog(event, "LiquidationThresholdPeriodUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsvMinimumLiquidationCollateralUpdatedIterator is returned from FilterMinimumLiquidationCollateralUpdated and is used to iterate over the raw logs and unpacked data for MinimumLiquidationCollateralUpdated events raised by the Ssv contract.
type SsvMinimumLiquidationCollateralUpdatedIterator struct {
	Event *SsvMinimumLiquidationCollateralUpdated // Event containing the contract specifics and raw log

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
func (it *SsvMinimumLiquidationCollateralUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvMinimumLiquidationCollateralUpdated)
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
		it.Event = new(SsvMinimumLiquidationCollateralUpdated)
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
func (it *SsvMinimumLiquidationCollateralUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvMinimumLiquidationCollateralUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvMinimumLiquidationCollateralUpdated represents a MinimumLiquidationCollateralUpdated event raised by the Ssv contract.
type SsvMinimumLiquidationCollateralUpdated struct {
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMinimumLiquidationCollateralUpdated is a free log retrieval operation binding the contract event 0xd363ab4392efaf967a89d8616cba1ff0c6f05a04c2f214671be365f0fab05960.
//
// Solidity: event MinimumLiquidationCollateralUpdated(uint256 value)
func (_Ssv *SsvFilterer) FilterMinimumLiquidationCollateralUpdated(opts *bind.FilterOpts) (*SsvMinimumLiquidationCollateralUpdatedIterator, error) {

	logs, sub, err := _Ssv.contract.FilterLogs(opts, "MinimumLiquidationCollateralUpdated")
	if err != nil {
		return nil, err
	}
	return &SsvMinimumLiquidationCollateralUpdatedIterator{contract: _Ssv.contract, event: "MinimumLiquidationCollateralUpdated", logs: logs, sub: sub}, nil
}

// WatchMinimumLiquidationCollateralUpdated is a free log subscription operation binding the contract event 0xd363ab4392efaf967a89d8616cba1ff0c6f05a04c2f214671be365f0fab05960.
//
// Solidity: event MinimumLiquidationCollateralUpdated(uint256 value)
func (_Ssv *SsvFilterer) WatchMinimumLiquidationCollateralUpdated(opts *bind.WatchOpts, sink chan<- *SsvMinimumLiquidationCollateralUpdated) (event.Subscription, error) {

	logs, sub, err := _Ssv.contract.WatchLogs(opts, "MinimumLiquidationCollateralUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvMinimumLiquidationCollateralUpdated)
				if err := _Ssv.contract.UnpackLog(event, "MinimumLiquidationCollateralUpdated", log); err != nil {
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

// ParseMinimumLiquidationCollateralUpdated is a log parse operation binding the contract event 0xd363ab4392efaf967a89d8616cba1ff0c6f05a04c2f214671be365f0fab05960.
//
// Solidity: event MinimumLiquidationCollateralUpdated(uint256 value)
func (_Ssv *SsvFilterer) ParseMinimumLiquidationCollateralUpdated(log types.Log) (*SsvMinimumLiquidationCollateralUpdated, error) {
	event := new(SsvMinimumLiquidationCollateralUpdated)
	if err := _Ssv.contract.UnpackLog(event, "MinimumLiquidationCollateralUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsvNetworkEarningsWithdrawnIterator is returned from FilterNetworkEarningsWithdrawn and is used to iterate over the raw logs and unpacked data for NetworkEarningsWithdrawn events raised by the Ssv contract.
type SsvNetworkEarningsWithdrawnIterator struct {
	Event *SsvNetworkEarningsWithdrawn // Event containing the contract specifics and raw log

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
func (it *SsvNetworkEarningsWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvNetworkEarningsWithdrawn)
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
		it.Event = new(SsvNetworkEarningsWithdrawn)
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
func (it *SsvNetworkEarningsWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvNetworkEarningsWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvNetworkEarningsWithdrawn represents a NetworkEarningsWithdrawn event raised by the Ssv contract.
type SsvNetworkEarningsWithdrawn struct {
	Value     *big.Int
	Recipient common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNetworkEarningsWithdrawn is a free log retrieval operation binding the contract event 0x370342c3bb9245e20bffe6dced02ba2fceca979701f881d5adc72d838e83f1c5.
//
// Solidity: event NetworkEarningsWithdrawn(uint256 value, address recipient)
func (_Ssv *SsvFilterer) FilterNetworkEarningsWithdrawn(opts *bind.FilterOpts) (*SsvNetworkEarningsWithdrawnIterator, error) {

	logs, sub, err := _Ssv.contract.FilterLogs(opts, "NetworkEarningsWithdrawn")
	if err != nil {
		return nil, err
	}
	return &SsvNetworkEarningsWithdrawnIterator{contract: _Ssv.contract, event: "NetworkEarningsWithdrawn", logs: logs, sub: sub}, nil
}

// WatchNetworkEarningsWithdrawn is a free log subscription operation binding the contract event 0x370342c3bb9245e20bffe6dced02ba2fceca979701f881d5adc72d838e83f1c5.
//
// Solidity: event NetworkEarningsWithdrawn(uint256 value, address recipient)
func (_Ssv *SsvFilterer) WatchNetworkEarningsWithdrawn(opts *bind.WatchOpts, sink chan<- *SsvNetworkEarningsWithdrawn) (event.Subscription, error) {

	logs, sub, err := _Ssv.contract.WatchLogs(opts, "NetworkEarningsWithdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvNetworkEarningsWithdrawn)
				if err := _Ssv.contract.UnpackLog(event, "NetworkEarningsWithdrawn", log); err != nil {
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

// ParseNetworkEarningsWithdrawn is a log parse operation binding the contract event 0x370342c3bb9245e20bffe6dced02ba2fceca979701f881d5adc72d838e83f1c5.
//
// Solidity: event NetworkEarningsWithdrawn(uint256 value, address recipient)
func (_Ssv *SsvFilterer) ParseNetworkEarningsWithdrawn(log types.Log) (*SsvNetworkEarningsWithdrawn, error) {
	event := new(SsvNetworkEarningsWithdrawn)
	if err := _Ssv.contract.UnpackLog(event, "NetworkEarningsWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsvNetworkFeeUpdatedIterator is returned from FilterNetworkFeeUpdated and is used to iterate over the raw logs and unpacked data for NetworkFeeUpdated events raised by the Ssv contract.
type SsvNetworkFeeUpdatedIterator struct {
	Event *SsvNetworkFeeUpdated // Event containing the contract specifics and raw log

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
func (it *SsvNetworkFeeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvNetworkFeeUpdated)
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
		it.Event = new(SsvNetworkFeeUpdated)
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
func (it *SsvNetworkFeeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvNetworkFeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvNetworkFeeUpdated represents a NetworkFeeUpdated event raised by the Ssv contract.
type SsvNetworkFeeUpdated struct {
	OldFee *big.Int
	NewFee *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNetworkFeeUpdated is a free log retrieval operation binding the contract event 0x8f49a76c5d617bd72673d92d3a019ff8f04f204536aae7a3d10e7ca85603f3cc.
//
// Solidity: event NetworkFeeUpdated(uint256 oldFee, uint256 newFee)
func (_Ssv *SsvFilterer) FilterNetworkFeeUpdated(opts *bind.FilterOpts) (*SsvNetworkFeeUpdatedIterator, error) {

	logs, sub, err := _Ssv.contract.FilterLogs(opts, "NetworkFeeUpdated")
	if err != nil {
		return nil, err
	}
	return &SsvNetworkFeeUpdatedIterator{contract: _Ssv.contract, event: "NetworkFeeUpdated", logs: logs, sub: sub}, nil
}

// WatchNetworkFeeUpdated is a free log subscription operation binding the contract event 0x8f49a76c5d617bd72673d92d3a019ff8f04f204536aae7a3d10e7ca85603f3cc.
//
// Solidity: event NetworkFeeUpdated(uint256 oldFee, uint256 newFee)
func (_Ssv *SsvFilterer) WatchNetworkFeeUpdated(opts *bind.WatchOpts, sink chan<- *SsvNetworkFeeUpdated) (event.Subscription, error) {

	logs, sub, err := _Ssv.contract.WatchLogs(opts, "NetworkFeeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvNetworkFeeUpdated)
				if err := _Ssv.contract.UnpackLog(event, "NetworkFeeUpdated", log); err != nil {
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

// ParseNetworkFeeUpdated is a log parse operation binding the contract event 0x8f49a76c5d617bd72673d92d3a019ff8f04f204536aae7a3d10e7ca85603f3cc.
//
// Solidity: event NetworkFeeUpdated(uint256 oldFee, uint256 newFee)
func (_Ssv *SsvFilterer) ParseNetworkFeeUpdated(log types.Log) (*SsvNetworkFeeUpdated, error) {
	event := new(SsvNetworkFeeUpdated)
	if err := _Ssv.contract.UnpackLog(event, "NetworkFeeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsvOperatorAddedIterator is returned from FilterOperatorAdded and is used to iterate over the raw logs and unpacked data for OperatorAdded events raised by the Ssv contract.
type SsvOperatorAddedIterator struct {
	Event *SsvOperatorAdded // Event containing the contract specifics and raw log

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
func (it *SsvOperatorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvOperatorAdded)
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
		it.Event = new(SsvOperatorAdded)
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
func (it *SsvOperatorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvOperatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvOperatorAdded represents a OperatorAdded event raised by the Ssv contract.
type SsvOperatorAdded struct {
	OperatorId uint64
	Owner      common.Address
	PublicKey  []byte
	Fee        *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOperatorAdded is a free log retrieval operation binding the contract event 0xd839f31c14bd632f424e307b36abff63ca33684f77f28e35dc13718ef338f7f4.
//
// Solidity: event OperatorAdded(uint64 indexed operatorId, address indexed owner, bytes publicKey, uint256 fee)
func (_Ssv *SsvFilterer) FilterOperatorAdded(opts *bind.FilterOpts, operatorId []uint64, owner []common.Address) (*SsvOperatorAddedIterator, error) {

	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Ssv.contract.FilterLogs(opts, "OperatorAdded", operatorIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &SsvOperatorAddedIterator{contract: _Ssv.contract, event: "OperatorAdded", logs: logs, sub: sub}, nil
}

// WatchOperatorAdded is a free log subscription operation binding the contract event 0xd839f31c14bd632f424e307b36abff63ca33684f77f28e35dc13718ef338f7f4.
//
// Solidity: event OperatorAdded(uint64 indexed operatorId, address indexed owner, bytes publicKey, uint256 fee)
func (_Ssv *SsvFilterer) WatchOperatorAdded(opts *bind.WatchOpts, sink chan<- *SsvOperatorAdded, operatorId []uint64, owner []common.Address) (event.Subscription, error) {

	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Ssv.contract.WatchLogs(opts, "OperatorAdded", operatorIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvOperatorAdded)
				if err := _Ssv.contract.UnpackLog(event, "OperatorAdded", log); err != nil {
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

// ParseOperatorAdded is a log parse operation binding the contract event 0xd839f31c14bd632f424e307b36abff63ca33684f77f28e35dc13718ef338f7f4.
//
// Solidity: event OperatorAdded(uint64 indexed operatorId, address indexed owner, bytes publicKey, uint256 fee)
func (_Ssv *SsvFilterer) ParseOperatorAdded(log types.Log) (*SsvOperatorAdded, error) {
	event := new(SsvOperatorAdded)
	if err := _Ssv.contract.UnpackLog(event, "OperatorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsvOperatorFeeDeclarationCancelledIterator is returned from FilterOperatorFeeDeclarationCancelled and is used to iterate over the raw logs and unpacked data for OperatorFeeDeclarationCancelled events raised by the Ssv contract.
type SsvOperatorFeeDeclarationCancelledIterator struct {
	Event *SsvOperatorFeeDeclarationCancelled // Event containing the contract specifics and raw log

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
func (it *SsvOperatorFeeDeclarationCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvOperatorFeeDeclarationCancelled)
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
		it.Event = new(SsvOperatorFeeDeclarationCancelled)
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
func (it *SsvOperatorFeeDeclarationCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvOperatorFeeDeclarationCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvOperatorFeeDeclarationCancelled represents a OperatorFeeDeclarationCancelled event raised by the Ssv contract.
type SsvOperatorFeeDeclarationCancelled struct {
	Owner      common.Address
	OperatorId uint64
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOperatorFeeDeclarationCancelled is a free log retrieval operation binding the contract event 0x5055fa347441172447637c015e80a3ee748b9382212ceb5dca5a3683298fd6f3.
//
// Solidity: event OperatorFeeDeclarationCancelled(address indexed owner, uint64 indexed operatorId)
func (_Ssv *SsvFilterer) FilterOperatorFeeDeclarationCancelled(opts *bind.FilterOpts, owner []common.Address, operatorId []uint64) (*SsvOperatorFeeDeclarationCancelledIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _Ssv.contract.FilterLogs(opts, "OperatorFeeDeclarationCancelled", ownerRule, operatorIdRule)
	if err != nil {
		return nil, err
	}
	return &SsvOperatorFeeDeclarationCancelledIterator{contract: _Ssv.contract, event: "OperatorFeeDeclarationCancelled", logs: logs, sub: sub}, nil
}

// WatchOperatorFeeDeclarationCancelled is a free log subscription operation binding the contract event 0x5055fa347441172447637c015e80a3ee748b9382212ceb5dca5a3683298fd6f3.
//
// Solidity: event OperatorFeeDeclarationCancelled(address indexed owner, uint64 indexed operatorId)
func (_Ssv *SsvFilterer) WatchOperatorFeeDeclarationCancelled(opts *bind.WatchOpts, sink chan<- *SsvOperatorFeeDeclarationCancelled, owner []common.Address, operatorId []uint64) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _Ssv.contract.WatchLogs(opts, "OperatorFeeDeclarationCancelled", ownerRule, operatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvOperatorFeeDeclarationCancelled)
				if err := _Ssv.contract.UnpackLog(event, "OperatorFeeDeclarationCancelled", log); err != nil {
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

// ParseOperatorFeeDeclarationCancelled is a log parse operation binding the contract event 0x5055fa347441172447637c015e80a3ee748b9382212ceb5dca5a3683298fd6f3.
//
// Solidity: event OperatorFeeDeclarationCancelled(address indexed owner, uint64 indexed operatorId)
func (_Ssv *SsvFilterer) ParseOperatorFeeDeclarationCancelled(log types.Log) (*SsvOperatorFeeDeclarationCancelled, error) {
	event := new(SsvOperatorFeeDeclarationCancelled)
	if err := _Ssv.contract.UnpackLog(event, "OperatorFeeDeclarationCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsvOperatorFeeDeclaredIterator is returned from FilterOperatorFeeDeclared and is used to iterate over the raw logs and unpacked data for OperatorFeeDeclared events raised by the Ssv contract.
type SsvOperatorFeeDeclaredIterator struct {
	Event *SsvOperatorFeeDeclared // Event containing the contract specifics and raw log

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
func (it *SsvOperatorFeeDeclaredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvOperatorFeeDeclared)
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
		it.Event = new(SsvOperatorFeeDeclared)
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
func (it *SsvOperatorFeeDeclaredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvOperatorFeeDeclaredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvOperatorFeeDeclared represents a OperatorFeeDeclared event raised by the Ssv contract.
type SsvOperatorFeeDeclared struct {
	Owner       common.Address
	OperatorId  uint64
	BlockNumber *big.Int
	Fee         *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorFeeDeclared is a free log retrieval operation binding the contract event 0x796204296f2eb56d7432fa85961e9750d0cb21741873ebf7077e28263e327358.
//
// Solidity: event OperatorFeeDeclared(address indexed owner, uint64 indexed operatorId, uint256 blockNumber, uint256 fee)
func (_Ssv *SsvFilterer) FilterOperatorFeeDeclared(opts *bind.FilterOpts, owner []common.Address, operatorId []uint64) (*SsvOperatorFeeDeclaredIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _Ssv.contract.FilterLogs(opts, "OperatorFeeDeclared", ownerRule, operatorIdRule)
	if err != nil {
		return nil, err
	}
	return &SsvOperatorFeeDeclaredIterator{contract: _Ssv.contract, event: "OperatorFeeDeclared", logs: logs, sub: sub}, nil
}

// WatchOperatorFeeDeclared is a free log subscription operation binding the contract event 0x796204296f2eb56d7432fa85961e9750d0cb21741873ebf7077e28263e327358.
//
// Solidity: event OperatorFeeDeclared(address indexed owner, uint64 indexed operatorId, uint256 blockNumber, uint256 fee)
func (_Ssv *SsvFilterer) WatchOperatorFeeDeclared(opts *bind.WatchOpts, sink chan<- *SsvOperatorFeeDeclared, owner []common.Address, operatorId []uint64) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _Ssv.contract.WatchLogs(opts, "OperatorFeeDeclared", ownerRule, operatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvOperatorFeeDeclared)
				if err := _Ssv.contract.UnpackLog(event, "OperatorFeeDeclared", log); err != nil {
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

// ParseOperatorFeeDeclared is a log parse operation binding the contract event 0x796204296f2eb56d7432fa85961e9750d0cb21741873ebf7077e28263e327358.
//
// Solidity: event OperatorFeeDeclared(address indexed owner, uint64 indexed operatorId, uint256 blockNumber, uint256 fee)
func (_Ssv *SsvFilterer) ParseOperatorFeeDeclared(log types.Log) (*SsvOperatorFeeDeclared, error) {
	event := new(SsvOperatorFeeDeclared)
	if err := _Ssv.contract.UnpackLog(event, "OperatorFeeDeclared", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsvOperatorFeeExecutedIterator is returned from FilterOperatorFeeExecuted and is used to iterate over the raw logs and unpacked data for OperatorFeeExecuted events raised by the Ssv contract.
type SsvOperatorFeeExecutedIterator struct {
	Event *SsvOperatorFeeExecuted // Event containing the contract specifics and raw log

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
func (it *SsvOperatorFeeExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvOperatorFeeExecuted)
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
		it.Event = new(SsvOperatorFeeExecuted)
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
func (it *SsvOperatorFeeExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvOperatorFeeExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvOperatorFeeExecuted represents a OperatorFeeExecuted event raised by the Ssv contract.
type SsvOperatorFeeExecuted struct {
	Owner       common.Address
	OperatorId  uint64
	BlockNumber *big.Int
	Fee         *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorFeeExecuted is a free log retrieval operation binding the contract event 0x513e931ff778ed01e676d55880d8db185c29b0094546ff2b3e9f5b6920d16bef.
//
// Solidity: event OperatorFeeExecuted(address indexed owner, uint64 indexed operatorId, uint256 blockNumber, uint256 fee)
func (_Ssv *SsvFilterer) FilterOperatorFeeExecuted(opts *bind.FilterOpts, owner []common.Address, operatorId []uint64) (*SsvOperatorFeeExecutedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _Ssv.contract.FilterLogs(opts, "OperatorFeeExecuted", ownerRule, operatorIdRule)
	if err != nil {
		return nil, err
	}
	return &SsvOperatorFeeExecutedIterator{contract: _Ssv.contract, event: "OperatorFeeExecuted", logs: logs, sub: sub}, nil
}

// WatchOperatorFeeExecuted is a free log subscription operation binding the contract event 0x513e931ff778ed01e676d55880d8db185c29b0094546ff2b3e9f5b6920d16bef.
//
// Solidity: event OperatorFeeExecuted(address indexed owner, uint64 indexed operatorId, uint256 blockNumber, uint256 fee)
func (_Ssv *SsvFilterer) WatchOperatorFeeExecuted(opts *bind.WatchOpts, sink chan<- *SsvOperatorFeeExecuted, owner []common.Address, operatorId []uint64) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _Ssv.contract.WatchLogs(opts, "OperatorFeeExecuted", ownerRule, operatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvOperatorFeeExecuted)
				if err := _Ssv.contract.UnpackLog(event, "OperatorFeeExecuted", log); err != nil {
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

// ParseOperatorFeeExecuted is a log parse operation binding the contract event 0x513e931ff778ed01e676d55880d8db185c29b0094546ff2b3e9f5b6920d16bef.
//
// Solidity: event OperatorFeeExecuted(address indexed owner, uint64 indexed operatorId, uint256 blockNumber, uint256 fee)
func (_Ssv *SsvFilterer) ParseOperatorFeeExecuted(log types.Log) (*SsvOperatorFeeExecuted, error) {
	event := new(SsvOperatorFeeExecuted)
	if err := _Ssv.contract.UnpackLog(event, "OperatorFeeExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsvOperatorFeeIncreaseLimitUpdatedIterator is returned from FilterOperatorFeeIncreaseLimitUpdated and is used to iterate over the raw logs and unpacked data for OperatorFeeIncreaseLimitUpdated events raised by the Ssv contract.
type SsvOperatorFeeIncreaseLimitUpdatedIterator struct {
	Event *SsvOperatorFeeIncreaseLimitUpdated // Event containing the contract specifics and raw log

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
func (it *SsvOperatorFeeIncreaseLimitUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvOperatorFeeIncreaseLimitUpdated)
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
		it.Event = new(SsvOperatorFeeIncreaseLimitUpdated)
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
func (it *SsvOperatorFeeIncreaseLimitUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvOperatorFeeIncreaseLimitUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvOperatorFeeIncreaseLimitUpdated represents a OperatorFeeIncreaseLimitUpdated event raised by the Ssv contract.
type SsvOperatorFeeIncreaseLimitUpdated struct {
	Value uint64
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterOperatorFeeIncreaseLimitUpdated is a free log retrieval operation binding the contract event 0x2fff7e5a48a4befc2c2be4d77e141f6d97907798977ce452429ec55c2658a342.
//
// Solidity: event OperatorFeeIncreaseLimitUpdated(uint64 value)
func (_Ssv *SsvFilterer) FilterOperatorFeeIncreaseLimitUpdated(opts *bind.FilterOpts) (*SsvOperatorFeeIncreaseLimitUpdatedIterator, error) {

	logs, sub, err := _Ssv.contract.FilterLogs(opts, "OperatorFeeIncreaseLimitUpdated")
	if err != nil {
		return nil, err
	}
	return &SsvOperatorFeeIncreaseLimitUpdatedIterator{contract: _Ssv.contract, event: "OperatorFeeIncreaseLimitUpdated", logs: logs, sub: sub}, nil
}

// WatchOperatorFeeIncreaseLimitUpdated is a free log subscription operation binding the contract event 0x2fff7e5a48a4befc2c2be4d77e141f6d97907798977ce452429ec55c2658a342.
//
// Solidity: event OperatorFeeIncreaseLimitUpdated(uint64 value)
func (_Ssv *SsvFilterer) WatchOperatorFeeIncreaseLimitUpdated(opts *bind.WatchOpts, sink chan<- *SsvOperatorFeeIncreaseLimitUpdated) (event.Subscription, error) {

	logs, sub, err := _Ssv.contract.WatchLogs(opts, "OperatorFeeIncreaseLimitUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvOperatorFeeIncreaseLimitUpdated)
				if err := _Ssv.contract.UnpackLog(event, "OperatorFeeIncreaseLimitUpdated", log); err != nil {
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

// ParseOperatorFeeIncreaseLimitUpdated is a log parse operation binding the contract event 0x2fff7e5a48a4befc2c2be4d77e141f6d97907798977ce452429ec55c2658a342.
//
// Solidity: event OperatorFeeIncreaseLimitUpdated(uint64 value)
func (_Ssv *SsvFilterer) ParseOperatorFeeIncreaseLimitUpdated(log types.Log) (*SsvOperatorFeeIncreaseLimitUpdated, error) {
	event := new(SsvOperatorFeeIncreaseLimitUpdated)
	if err := _Ssv.contract.UnpackLog(event, "OperatorFeeIncreaseLimitUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsvOperatorMaximumFeeUpdatedIterator is returned from FilterOperatorMaximumFeeUpdated and is used to iterate over the raw logs and unpacked data for OperatorMaximumFeeUpdated events raised by the Ssv contract.
type SsvOperatorMaximumFeeUpdatedIterator struct {
	Event *SsvOperatorMaximumFeeUpdated // Event containing the contract specifics and raw log

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
func (it *SsvOperatorMaximumFeeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvOperatorMaximumFeeUpdated)
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
		it.Event = new(SsvOperatorMaximumFeeUpdated)
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
func (it *SsvOperatorMaximumFeeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvOperatorMaximumFeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvOperatorMaximumFeeUpdated represents a OperatorMaximumFeeUpdated event raised by the Ssv contract.
type SsvOperatorMaximumFeeUpdated struct {
	MaxFee uint64
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterOperatorMaximumFeeUpdated is a free log retrieval operation binding the contract event 0x38552bed8df52ac76c5de6da688eafcda7d7b070f6c987f391a07dd69986d783.
//
// Solidity: event OperatorMaximumFeeUpdated(uint64 maxFee)
func (_Ssv *SsvFilterer) FilterOperatorMaximumFeeUpdated(opts *bind.FilterOpts) (*SsvOperatorMaximumFeeUpdatedIterator, error) {

	logs, sub, err := _Ssv.contract.FilterLogs(opts, "OperatorMaximumFeeUpdated")
	if err != nil {
		return nil, err
	}
	return &SsvOperatorMaximumFeeUpdatedIterator{contract: _Ssv.contract, event: "OperatorMaximumFeeUpdated", logs: logs, sub: sub}, nil
}

// WatchOperatorMaximumFeeUpdated is a free log subscription operation binding the contract event 0x38552bed8df52ac76c5de6da688eafcda7d7b070f6c987f391a07dd69986d783.
//
// Solidity: event OperatorMaximumFeeUpdated(uint64 maxFee)
func (_Ssv *SsvFilterer) WatchOperatorMaximumFeeUpdated(opts *bind.WatchOpts, sink chan<- *SsvOperatorMaximumFeeUpdated) (event.Subscription, error) {

	logs, sub, err := _Ssv.contract.WatchLogs(opts, "OperatorMaximumFeeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvOperatorMaximumFeeUpdated)
				if err := _Ssv.contract.UnpackLog(event, "OperatorMaximumFeeUpdated", log); err != nil {
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

// ParseOperatorMaximumFeeUpdated is a log parse operation binding the contract event 0x38552bed8df52ac76c5de6da688eafcda7d7b070f6c987f391a07dd69986d783.
//
// Solidity: event OperatorMaximumFeeUpdated(uint64 maxFee)
func (_Ssv *SsvFilterer) ParseOperatorMaximumFeeUpdated(log types.Log) (*SsvOperatorMaximumFeeUpdated, error) {
	event := new(SsvOperatorMaximumFeeUpdated)
	if err := _Ssv.contract.UnpackLog(event, "OperatorMaximumFeeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsvOperatorRemovedIterator is returned from FilterOperatorRemoved and is used to iterate over the raw logs and unpacked data for OperatorRemoved events raised by the Ssv contract.
type SsvOperatorRemovedIterator struct {
	Event *SsvOperatorRemoved // Event containing the contract specifics and raw log

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
func (it *SsvOperatorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvOperatorRemoved)
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
		it.Event = new(SsvOperatorRemoved)
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
func (it *SsvOperatorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvOperatorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvOperatorRemoved represents a OperatorRemoved event raised by the Ssv contract.
type SsvOperatorRemoved struct {
	OperatorId uint64
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOperatorRemoved is a free log retrieval operation binding the contract event 0x0e0ba6c2b04de36d6d509ec5bd155c43a9fe862f8052096dd54f3902a74cca3e.
//
// Solidity: event OperatorRemoved(uint64 indexed operatorId)
func (_Ssv *SsvFilterer) FilterOperatorRemoved(opts *bind.FilterOpts, operatorId []uint64) (*SsvOperatorRemovedIterator, error) {

	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _Ssv.contract.FilterLogs(opts, "OperatorRemoved", operatorIdRule)
	if err != nil {
		return nil, err
	}
	return &SsvOperatorRemovedIterator{contract: _Ssv.contract, event: "OperatorRemoved", logs: logs, sub: sub}, nil
}

// WatchOperatorRemoved is a free log subscription operation binding the contract event 0x0e0ba6c2b04de36d6d509ec5bd155c43a9fe862f8052096dd54f3902a74cca3e.
//
// Solidity: event OperatorRemoved(uint64 indexed operatorId)
func (_Ssv *SsvFilterer) WatchOperatorRemoved(opts *bind.WatchOpts, sink chan<- *SsvOperatorRemoved, operatorId []uint64) (event.Subscription, error) {

	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _Ssv.contract.WatchLogs(opts, "OperatorRemoved", operatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvOperatorRemoved)
				if err := _Ssv.contract.UnpackLog(event, "OperatorRemoved", log); err != nil {
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

// ParseOperatorRemoved is a log parse operation binding the contract event 0x0e0ba6c2b04de36d6d509ec5bd155c43a9fe862f8052096dd54f3902a74cca3e.
//
// Solidity: event OperatorRemoved(uint64 indexed operatorId)
func (_Ssv *SsvFilterer) ParseOperatorRemoved(log types.Log) (*SsvOperatorRemoved, error) {
	event := new(SsvOperatorRemoved)
	if err := _Ssv.contract.UnpackLog(event, "OperatorRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsvOperatorWhitelistUpdatedIterator is returned from FilterOperatorWhitelistUpdated and is used to iterate over the raw logs and unpacked data for OperatorWhitelistUpdated events raised by the Ssv contract.
type SsvOperatorWhitelistUpdatedIterator struct {
	Event *SsvOperatorWhitelistUpdated // Event containing the contract specifics and raw log

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
func (it *SsvOperatorWhitelistUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvOperatorWhitelistUpdated)
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
		it.Event = new(SsvOperatorWhitelistUpdated)
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
func (it *SsvOperatorWhitelistUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvOperatorWhitelistUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvOperatorWhitelistUpdated represents a OperatorWhitelistUpdated event raised by the Ssv contract.
type SsvOperatorWhitelistUpdated struct {
	OperatorId  uint64
	Whitelisted common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorWhitelistUpdated is a free log retrieval operation binding the contract event 0x29f72634ccb172103f8c542da23de7f6cf9bce724c5bb91bd6f3a516b14c63fe.
//
// Solidity: event OperatorWhitelistUpdated(uint64 indexed operatorId, address whitelisted)
func (_Ssv *SsvFilterer) FilterOperatorWhitelistUpdated(opts *bind.FilterOpts, operatorId []uint64) (*SsvOperatorWhitelistUpdatedIterator, error) {

	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _Ssv.contract.FilterLogs(opts, "OperatorWhitelistUpdated", operatorIdRule)
	if err != nil {
		return nil, err
	}
	return &SsvOperatorWhitelistUpdatedIterator{contract: _Ssv.contract, event: "OperatorWhitelistUpdated", logs: logs, sub: sub}, nil
}

// WatchOperatorWhitelistUpdated is a free log subscription operation binding the contract event 0x29f72634ccb172103f8c542da23de7f6cf9bce724c5bb91bd6f3a516b14c63fe.
//
// Solidity: event OperatorWhitelistUpdated(uint64 indexed operatorId, address whitelisted)
func (_Ssv *SsvFilterer) WatchOperatorWhitelistUpdated(opts *bind.WatchOpts, sink chan<- *SsvOperatorWhitelistUpdated, operatorId []uint64) (event.Subscription, error) {

	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _Ssv.contract.WatchLogs(opts, "OperatorWhitelistUpdated", operatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvOperatorWhitelistUpdated)
				if err := _Ssv.contract.UnpackLog(event, "OperatorWhitelistUpdated", log); err != nil {
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

// ParseOperatorWhitelistUpdated is a log parse operation binding the contract event 0x29f72634ccb172103f8c542da23de7f6cf9bce724c5bb91bd6f3a516b14c63fe.
//
// Solidity: event OperatorWhitelistUpdated(uint64 indexed operatorId, address whitelisted)
func (_Ssv *SsvFilterer) ParseOperatorWhitelistUpdated(log types.Log) (*SsvOperatorWhitelistUpdated, error) {
	event := new(SsvOperatorWhitelistUpdated)
	if err := _Ssv.contract.UnpackLog(event, "OperatorWhitelistUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsvOperatorWithdrawnIterator is returned from FilterOperatorWithdrawn and is used to iterate over the raw logs and unpacked data for OperatorWithdrawn events raised by the Ssv contract.
type SsvOperatorWithdrawnIterator struct {
	Event *SsvOperatorWithdrawn // Event containing the contract specifics and raw log

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
func (it *SsvOperatorWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvOperatorWithdrawn)
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
		it.Event = new(SsvOperatorWithdrawn)
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
func (it *SsvOperatorWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvOperatorWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvOperatorWithdrawn represents a OperatorWithdrawn event raised by the Ssv contract.
type SsvOperatorWithdrawn struct {
	Owner      common.Address
	OperatorId uint64
	Value      *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOperatorWithdrawn is a free log retrieval operation binding the contract event 0x178bf78bdd8914b8483d640b4a4f84e20943b5eb6b639b7474286364c7651d60.
//
// Solidity: event OperatorWithdrawn(address indexed owner, uint64 indexed operatorId, uint256 value)
func (_Ssv *SsvFilterer) FilterOperatorWithdrawn(opts *bind.FilterOpts, owner []common.Address, operatorId []uint64) (*SsvOperatorWithdrawnIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _Ssv.contract.FilterLogs(opts, "OperatorWithdrawn", ownerRule, operatorIdRule)
	if err != nil {
		return nil, err
	}
	return &SsvOperatorWithdrawnIterator{contract: _Ssv.contract, event: "OperatorWithdrawn", logs: logs, sub: sub}, nil
}

// WatchOperatorWithdrawn is a free log subscription operation binding the contract event 0x178bf78bdd8914b8483d640b4a4f84e20943b5eb6b639b7474286364c7651d60.
//
// Solidity: event OperatorWithdrawn(address indexed owner, uint64 indexed operatorId, uint256 value)
func (_Ssv *SsvFilterer) WatchOperatorWithdrawn(opts *bind.WatchOpts, sink chan<- *SsvOperatorWithdrawn, owner []common.Address, operatorId []uint64) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorIdRule []interface{}
	for _, operatorIdItem := range operatorId {
		operatorIdRule = append(operatorIdRule, operatorIdItem)
	}

	logs, sub, err := _Ssv.contract.WatchLogs(opts, "OperatorWithdrawn", ownerRule, operatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvOperatorWithdrawn)
				if err := _Ssv.contract.UnpackLog(event, "OperatorWithdrawn", log); err != nil {
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

// ParseOperatorWithdrawn is a log parse operation binding the contract event 0x178bf78bdd8914b8483d640b4a4f84e20943b5eb6b639b7474286364c7651d60.
//
// Solidity: event OperatorWithdrawn(address indexed owner, uint64 indexed operatorId, uint256 value)
func (_Ssv *SsvFilterer) ParseOperatorWithdrawn(log types.Log) (*SsvOperatorWithdrawn, error) {
	event := new(SsvOperatorWithdrawn)
	if err := _Ssv.contract.UnpackLog(event, "OperatorWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsvOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the Ssv contract.
type SsvOwnershipTransferStartedIterator struct {
	Event *SsvOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *SsvOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvOwnershipTransferStarted)
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
		it.Event = new(SsvOwnershipTransferStarted)
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
func (it *SsvOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the Ssv contract.
type SsvOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Ssv *SsvFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SsvOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ssv.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SsvOwnershipTransferStartedIterator{contract: _Ssv.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Ssv *SsvFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *SsvOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ssv.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvOwnershipTransferStarted)
				if err := _Ssv.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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

// ParseOwnershipTransferStarted is a log parse operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Ssv *SsvFilterer) ParseOwnershipTransferStarted(log types.Log) (*SsvOwnershipTransferStarted, error) {
	event := new(SsvOwnershipTransferStarted)
	if err := _Ssv.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsvOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ssv contract.
type SsvOwnershipTransferredIterator struct {
	Event *SsvOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SsvOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvOwnershipTransferred)
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
		it.Event = new(SsvOwnershipTransferred)
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
func (it *SsvOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvOwnershipTransferred represents a OwnershipTransferred event raised by the Ssv contract.
type SsvOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ssv *SsvFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SsvOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ssv.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SsvOwnershipTransferredIterator{contract: _Ssv.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ssv *SsvFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SsvOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ssv.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvOwnershipTransferred)
				if err := _Ssv.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Ssv *SsvFilterer) ParseOwnershipTransferred(log types.Log) (*SsvOwnershipTransferred, error) {
	event := new(SsvOwnershipTransferred)
	if err := _Ssv.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsvUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Ssv contract.
type SsvUpgradedIterator struct {
	Event *SsvUpgraded // Event containing the contract specifics and raw log

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
func (it *SsvUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvUpgraded)
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
		it.Event = new(SsvUpgraded)
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
func (it *SsvUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvUpgraded represents a Upgraded event raised by the Ssv contract.
type SsvUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Ssv *SsvFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*SsvUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Ssv.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &SsvUpgradedIterator{contract: _Ssv.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Ssv *SsvFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *SsvUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Ssv.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvUpgraded)
				if err := _Ssv.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_Ssv *SsvFilterer) ParseUpgraded(log types.Log) (*SsvUpgraded, error) {
	event := new(SsvUpgraded)
	if err := _Ssv.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsvValidatorAddedIterator is returned from FilterValidatorAdded and is used to iterate over the raw logs and unpacked data for ValidatorAdded events raised by the Ssv contract.
type SsvValidatorAddedIterator struct {
	Event *SsvValidatorAdded // Event containing the contract specifics and raw log

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
func (it *SsvValidatorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvValidatorAdded)
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
		it.Event = new(SsvValidatorAdded)
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
func (it *SsvValidatorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvValidatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvValidatorAdded represents a ValidatorAdded event raised by the Ssv contract.
type SsvValidatorAdded struct {
	Owner       common.Address
	OperatorIds []uint64
	PublicKey   []byte
	Shares      []byte
	Cluster     ISSVNetworkCoreCluster
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterValidatorAdded is a free log retrieval operation binding the contract event 0x48a3ea0796746043948f6341d17ff8200937b99262a0b48c2663b951ed7114e5.
//
// Solidity: event ValidatorAdded(address indexed owner, uint64[] operatorIds, bytes publicKey, bytes shares, (uint32,uint64,uint64,bool,uint256) cluster)
func (_Ssv *SsvFilterer) FilterValidatorAdded(opts *bind.FilterOpts, owner []common.Address) (*SsvValidatorAddedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Ssv.contract.FilterLogs(opts, "ValidatorAdded", ownerRule)
	if err != nil {
		return nil, err
	}
	return &SsvValidatorAddedIterator{contract: _Ssv.contract, event: "ValidatorAdded", logs: logs, sub: sub}, nil
}

// WatchValidatorAdded is a free log subscription operation binding the contract event 0x48a3ea0796746043948f6341d17ff8200937b99262a0b48c2663b951ed7114e5.
//
// Solidity: event ValidatorAdded(address indexed owner, uint64[] operatorIds, bytes publicKey, bytes shares, (uint32,uint64,uint64,bool,uint256) cluster)
func (_Ssv *SsvFilterer) WatchValidatorAdded(opts *bind.WatchOpts, sink chan<- *SsvValidatorAdded, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Ssv.contract.WatchLogs(opts, "ValidatorAdded", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvValidatorAdded)
				if err := _Ssv.contract.UnpackLog(event, "ValidatorAdded", log); err != nil {
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

// ParseValidatorAdded is a log parse operation binding the contract event 0x48a3ea0796746043948f6341d17ff8200937b99262a0b48c2663b951ed7114e5.
//
// Solidity: event ValidatorAdded(address indexed owner, uint64[] operatorIds, bytes publicKey, bytes shares, (uint32,uint64,uint64,bool,uint256) cluster)
func (_Ssv *SsvFilterer) ParseValidatorAdded(log types.Log) (*SsvValidatorAdded, error) {
	event := new(SsvValidatorAdded)
	if err := _Ssv.contract.UnpackLog(event, "ValidatorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsvValidatorExitedIterator is returned from FilterValidatorExited and is used to iterate over the raw logs and unpacked data for ValidatorExited events raised by the Ssv contract.
type SsvValidatorExitedIterator struct {
	Event *SsvValidatorExited // Event containing the contract specifics and raw log

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
func (it *SsvValidatorExitedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvValidatorExited)
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
		it.Event = new(SsvValidatorExited)
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
func (it *SsvValidatorExitedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvValidatorExitedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvValidatorExited represents a ValidatorExited event raised by the Ssv contract.
type SsvValidatorExited struct {
	Owner       common.Address
	OperatorIds []uint64
	PublicKey   []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterValidatorExited is a free log retrieval operation binding the contract event 0xb4b20ffb2eb1f020be3df600b2287914f50c07003526d3a9d89a9dd12351828c.
//
// Solidity: event ValidatorExited(address indexed owner, uint64[] operatorIds, bytes publicKey)
func (_Ssv *SsvFilterer) FilterValidatorExited(opts *bind.FilterOpts, owner []common.Address) (*SsvValidatorExitedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Ssv.contract.FilterLogs(opts, "ValidatorExited", ownerRule)
	if err != nil {
		return nil, err
	}
	return &SsvValidatorExitedIterator{contract: _Ssv.contract, event: "ValidatorExited", logs: logs, sub: sub}, nil
}

// WatchValidatorExited is a free log subscription operation binding the contract event 0xb4b20ffb2eb1f020be3df600b2287914f50c07003526d3a9d89a9dd12351828c.
//
// Solidity: event ValidatorExited(address indexed owner, uint64[] operatorIds, bytes publicKey)
func (_Ssv *SsvFilterer) WatchValidatorExited(opts *bind.WatchOpts, sink chan<- *SsvValidatorExited, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Ssv.contract.WatchLogs(opts, "ValidatorExited", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvValidatorExited)
				if err := _Ssv.contract.UnpackLog(event, "ValidatorExited", log); err != nil {
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

// ParseValidatorExited is a log parse operation binding the contract event 0xb4b20ffb2eb1f020be3df600b2287914f50c07003526d3a9d89a9dd12351828c.
//
// Solidity: event ValidatorExited(address indexed owner, uint64[] operatorIds, bytes publicKey)
func (_Ssv *SsvFilterer) ParseValidatorExited(log types.Log) (*SsvValidatorExited, error) {
	event := new(SsvValidatorExited)
	if err := _Ssv.contract.UnpackLog(event, "ValidatorExited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsvValidatorRemovedIterator is returned from FilterValidatorRemoved and is used to iterate over the raw logs and unpacked data for ValidatorRemoved events raised by the Ssv contract.
type SsvValidatorRemovedIterator struct {
	Event *SsvValidatorRemoved // Event containing the contract specifics and raw log

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
func (it *SsvValidatorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsvValidatorRemoved)
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
		it.Event = new(SsvValidatorRemoved)
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
func (it *SsvValidatorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsvValidatorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsvValidatorRemoved represents a ValidatorRemoved event raised by the Ssv contract.
type SsvValidatorRemoved struct {
	Owner       common.Address
	OperatorIds []uint64
	PublicKey   []byte
	Cluster     ISSVNetworkCoreCluster
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterValidatorRemoved is a free log retrieval operation binding the contract event 0xccf4370403e5fbbde0cd3f13426479dcd8a5916b05db424b7a2c04978cf8ce6e.
//
// Solidity: event ValidatorRemoved(address indexed owner, uint64[] operatorIds, bytes publicKey, (uint32,uint64,uint64,bool,uint256) cluster)
func (_Ssv *SsvFilterer) FilterValidatorRemoved(opts *bind.FilterOpts, owner []common.Address) (*SsvValidatorRemovedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Ssv.contract.FilterLogs(opts, "ValidatorRemoved", ownerRule)
	if err != nil {
		return nil, err
	}
	return &SsvValidatorRemovedIterator{contract: _Ssv.contract, event: "ValidatorRemoved", logs: logs, sub: sub}, nil
}

// WatchValidatorRemoved is a free log subscription operation binding the contract event 0xccf4370403e5fbbde0cd3f13426479dcd8a5916b05db424b7a2c04978cf8ce6e.
//
// Solidity: event ValidatorRemoved(address indexed owner, uint64[] operatorIds, bytes publicKey, (uint32,uint64,uint64,bool,uint256) cluster)
func (_Ssv *SsvFilterer) WatchValidatorRemoved(opts *bind.WatchOpts, sink chan<- *SsvValidatorRemoved, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Ssv.contract.WatchLogs(opts, "ValidatorRemoved", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsvValidatorRemoved)
				if err := _Ssv.contract.UnpackLog(event, "ValidatorRemoved", log); err != nil {
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

// ParseValidatorRemoved is a log parse operation binding the contract event 0xccf4370403e5fbbde0cd3f13426479dcd8a5916b05db424b7a2c04978cf8ce6e.
//
// Solidity: event ValidatorRemoved(address indexed owner, uint64[] operatorIds, bytes publicKey, (uint32,uint64,uint64,bool,uint256) cluster)
func (_Ssv *SsvFilterer) ParseValidatorRemoved(log types.Log) (*SsvValidatorRemoved, error) {
	event := new(SsvValidatorRemoved)
	if err := _Ssv.contract.UnpackLog(event, "ValidatorRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
