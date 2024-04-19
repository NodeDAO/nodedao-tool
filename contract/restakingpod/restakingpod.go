// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package restakingpod

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

// BeaconChainProofsStateRootProof is an auto generated low-level Go binding around an user-defined struct.
type BeaconChainProofsStateRootProof struct {
	BeaconStateRoot [32]byte
	Proof           []byte
}

// BeaconChainProofsWithdrawalProof is an auto generated low-level Go binding around an user-defined struct.
type BeaconChainProofsWithdrawalProof struct {
	WithdrawalProof                 []byte
	SlotProof                       []byte
	ExecutionPayloadProof           []byte
	TimestampProof                  []byte
	HistoricalSummaryBlockRootProof []byte
	BlockRootIndex                  uint64
	HistoricalSummaryIndex          uint64
	WithdrawalIndex                 uint64
	BlockRoot                       [32]byte
	SlotRoot                        [32]byte
	TimestampRoot                   [32]byte
	ExecutionPayloadRoot            [32]byte
}

// IDelegationManagerQueuedWithdrawalParams is an auto generated low-level Go binding around an user-defined struct.
type IDelegationManagerQueuedWithdrawalParams struct {
	Strategies []common.Address
	Shares     []*big.Int
	Withdrawer common.Address
}

// IDelegationManagerWithdrawal is an auto generated low-level Go binding around an user-defined struct.
type IDelegationManagerWithdrawal struct {
	Staker      common.Address
	DelegatedTo common.Address
	Withdrawer  common.Address
	Nonce       *big.Int
	StartBlock  uint32
	Strategies  []common.Address
	Shares      []*big.Int
}

// ISignatureUtilsSignatureWithExpiry is an auto generated low-level Go binding around an user-defined struct.
type ISignatureUtilsSignatureWithExpiry struct {
	Signature []byte
	Expiry    *big.Int
}

// RestakingpodMetaData contains all meta data concerning the Restakingpod contract.
var RestakingpodMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"activateRestaking\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claimDelayedWithdrawals\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"completeQueuedWithdrawals\",\"inputs\":[{\"name\":\"withdrawals\",\"type\":\"tuple[]\",\"internalType\":\"structIDelegationManager.Withdrawal[]\",\"components\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegatedTo\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"withdrawer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"shares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]},{\"name\":\"tokens\",\"type\":\"address[][]\",\"internalType\":\"contractIERC20[][]\"},{\"name\":\"middlewareTimesIndexes\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"receiveAsTokens\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"dao\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"delegateToEigenLayerOperator\",\"inputs\":[{\"name\":\"_eigenLayerOperator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"approverSignatureAndExpiry\",\"type\":\"tuple\",\"internalType\":\"structISignatureUtils.SignatureWithExpiry\",\"components\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"approverSalt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eigenLayerDelayedWithdrawalRouter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelayedWithdrawalRouter\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"eigenLayerDelegationManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDelegationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"eigenLayerEigenPod\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"eigenLayerEigenPodManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIEigenPodManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"eigenLayerOperator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"txGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllDelayedWithdrawals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getClaimableDelayedWithdrawals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"implementation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_ownerAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_dao\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_restakingPool\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_restakingPodManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_eigenLayerEigenPodManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_eigenLayerDelegationManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_eigenLayerDelayedWithdrawalRouter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"queueWithdrawals\",\"inputs\":[{\"name\":\"queuedWithdrawalParams\",\"type\":\"tuple[]\",\"internalType\":\"structIDelegationManager.QueuedWithdrawalParams[]\",\"components\":[{\"name\":\"strategies\",\"type\":\"address[]\",\"internalType\":\"contractIStrategy[]\"},{\"name\":\"shares\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"withdrawer\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"recoverTokens\",\"inputs\":[{\"name\":\"tokenList\",\"type\":\"address[]\",\"internalType\":\"contractIERC20[]\"},{\"name\":\"amountsToWithdraw\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"restakingPodManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"restakingPool\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setDao\",\"inputs\":[{\"name\":\"_dao\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRestakingPodManager\",\"inputs\":[{\"name\":\"_restakingPodManager\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setStakedButNotVerifiedEth\",\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stake\",\"inputs\":[{\"name\":\"_pubkey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_depositDataRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"stakedButNotVerifiedEth\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"undelegateEigenLayerOperator\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeTo\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"verifyAndProcessWithdrawals\",\"inputs\":[{\"name\":\"oracleTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"stateRootProof\",\"type\":\"tuple\",\"internalType\":\"structBeaconChainProofs.StateRootProof\",\"components\":[{\"name\":\"beaconStateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"withdrawalProofs\",\"type\":\"tuple[]\",\"internalType\":\"structBeaconChainProofs.WithdrawalProof[]\",\"components\":[{\"name\":\"withdrawalProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"slotProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"executionPayloadProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"timestampProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"historicalSummaryBlockRootProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blockRootIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"historicalSummaryIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"withdrawalIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"blockRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"slotRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"timestampRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"executionPayloadRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"validatorFieldsProofs\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"validatorFields\",\"type\":\"bytes32[][]\",\"internalType\":\"bytes32[][]\"},{\"name\":\"withdrawalFields\",\"type\":\"bytes32[][]\",\"internalType\":\"bytes32[][]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyBalanceUpdates\",\"inputs\":[{\"name\":\"oracleTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"validatorIndices\",\"type\":\"uint40[]\",\"internalType\":\"uint40[]\"},{\"name\":\"stateRootProof\",\"type\":\"tuple\",\"internalType\":\"structBeaconChainProofs.StateRootProof\",\"components\":[{\"name\":\"beaconStateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"validatorFieldsProofs\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"validatorFields\",\"type\":\"bytes32[][]\",\"internalType\":\"bytes32[][]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyWithdrawalCredentials\",\"inputs\":[{\"name\":\"oracleTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"stateRootProof\",\"type\":\"tuple\",\"internalType\":\"structBeaconChainProofs.StateRootProof\",\"components\":[{\"name\":\"beaconStateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"validatorIndices\",\"type\":\"uint40[]\",\"internalType\":\"uint40[]\"},{\"name\":\"validatorFieldsProofs\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"validatorFields\",\"type\":\"bytes32[][]\",\"internalType\":\"bytes32[][]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"withdrawBeforeRestaking\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawCredentials\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawNonBeaconChainETHBalanceWei\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amountToWithdraw\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AdminChanged\",\"inputs\":[{\"name\":\"previousAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BeaconUpgraded\",\"inputs\":[{\"name\":\"beacon\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DaoChanged\",\"inputs\":[{\"name\":\"_oldDao\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"_dao\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EigenLayerOperatorDelegated\",\"inputs\":[{\"name\":\"_delegateAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EigenLayerOperatorUndelegated\",\"inputs\":[{\"name\":\"_delegateAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Received\",\"inputs\":[{\"name\":\"_sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RestakingPodManagerChanged\",\"inputs\":[{\"name\":\"_oldPodManager\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"_podManager\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StakedButNotVerifiedEthChanged\",\"inputs\":[{\"name\":\"_oldAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_newAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ExecuteFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAddr\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidVersion\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidtypeId\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PermissionDenied\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TransferFailed\",\"inputs\":[]}]",
}

// RestakingpodABI is the input ABI used to generate the binding from.
// Deprecated: Use RestakingpodMetaData.ABI instead.
var RestakingpodABI = RestakingpodMetaData.ABI

// Restakingpod is an auto generated Go binding around an Ethereum contract.
type Restakingpod struct {
	RestakingpodCaller     // Read-only binding to the contract
	RestakingpodTransactor // Write-only binding to the contract
	RestakingpodFilterer   // Log filterer for contract events
}

// RestakingpodCaller is an auto generated read-only Go binding around an Ethereum contract.
type RestakingpodCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RestakingpodTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RestakingpodTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RestakingpodFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RestakingpodFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RestakingpodSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RestakingpodSession struct {
	Contract     *Restakingpod     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RestakingpodCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RestakingpodCallerSession struct {
	Contract *RestakingpodCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// RestakingpodTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RestakingpodTransactorSession struct {
	Contract     *RestakingpodTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// RestakingpodRaw is an auto generated low-level Go binding around an Ethereum contract.
type RestakingpodRaw struct {
	Contract *Restakingpod // Generic contract binding to access the raw methods on
}

// RestakingpodCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RestakingpodCallerRaw struct {
	Contract *RestakingpodCaller // Generic read-only contract binding to access the raw methods on
}

// RestakingpodTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RestakingpodTransactorRaw struct {
	Contract *RestakingpodTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRestakingpod creates a new instance of Restakingpod, bound to a specific deployed contract.
func NewRestakingpod(address common.Address, backend bind.ContractBackend) (*Restakingpod, error) {
	contract, err := bindRestakingpod(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Restakingpod{RestakingpodCaller: RestakingpodCaller{contract: contract}, RestakingpodTransactor: RestakingpodTransactor{contract: contract}, RestakingpodFilterer: RestakingpodFilterer{contract: contract}}, nil
}

// NewRestakingpodCaller creates a new read-only instance of Restakingpod, bound to a specific deployed contract.
func NewRestakingpodCaller(address common.Address, caller bind.ContractCaller) (*RestakingpodCaller, error) {
	contract, err := bindRestakingpod(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RestakingpodCaller{contract: contract}, nil
}

// NewRestakingpodTransactor creates a new write-only instance of Restakingpod, bound to a specific deployed contract.
func NewRestakingpodTransactor(address common.Address, transactor bind.ContractTransactor) (*RestakingpodTransactor, error) {
	contract, err := bindRestakingpod(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RestakingpodTransactor{contract: contract}, nil
}

// NewRestakingpodFilterer creates a new log filterer instance of Restakingpod, bound to a specific deployed contract.
func NewRestakingpodFilterer(address common.Address, filterer bind.ContractFilterer) (*RestakingpodFilterer, error) {
	contract, err := bindRestakingpod(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RestakingpodFilterer{contract: contract}, nil
}

// bindRestakingpod binds a generic wrapper to an already deployed contract.
func bindRestakingpod(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RestakingpodABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Restakingpod *RestakingpodRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Restakingpod.Contract.RestakingpodCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Restakingpod *RestakingpodRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Restakingpod.Contract.RestakingpodTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Restakingpod *RestakingpodRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Restakingpod.Contract.RestakingpodTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Restakingpod *RestakingpodCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Restakingpod.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Restakingpod *RestakingpodTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Restakingpod.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Restakingpod *RestakingpodTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Restakingpod.Contract.contract.Transact(opts, method, params...)
}

// Dao is a free data retrieval call binding the contract method 0x4162169f.
//
// Solidity: function dao() view returns(address)
func (_Restakingpod *RestakingpodCaller) Dao(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Restakingpod.contract.Call(opts, &out, "dao")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Dao is a free data retrieval call binding the contract method 0x4162169f.
//
// Solidity: function dao() view returns(address)
func (_Restakingpod *RestakingpodSession) Dao() (common.Address, error) {
	return _Restakingpod.Contract.Dao(&_Restakingpod.CallOpts)
}

// Dao is a free data retrieval call binding the contract method 0x4162169f.
//
// Solidity: function dao() view returns(address)
func (_Restakingpod *RestakingpodCallerSession) Dao() (common.Address, error) {
	return _Restakingpod.Contract.Dao(&_Restakingpod.CallOpts)
}

// EigenLayerDelayedWithdrawalRouter is a free data retrieval call binding the contract method 0x01349b52.
//
// Solidity: function eigenLayerDelayedWithdrawalRouter() view returns(address)
func (_Restakingpod *RestakingpodCaller) EigenLayerDelayedWithdrawalRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Restakingpod.contract.Call(opts, &out, "eigenLayerDelayedWithdrawalRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EigenLayerDelayedWithdrawalRouter is a free data retrieval call binding the contract method 0x01349b52.
//
// Solidity: function eigenLayerDelayedWithdrawalRouter() view returns(address)
func (_Restakingpod *RestakingpodSession) EigenLayerDelayedWithdrawalRouter() (common.Address, error) {
	return _Restakingpod.Contract.EigenLayerDelayedWithdrawalRouter(&_Restakingpod.CallOpts)
}

// EigenLayerDelayedWithdrawalRouter is a free data retrieval call binding the contract method 0x01349b52.
//
// Solidity: function eigenLayerDelayedWithdrawalRouter() view returns(address)
func (_Restakingpod *RestakingpodCallerSession) EigenLayerDelayedWithdrawalRouter() (common.Address, error) {
	return _Restakingpod.Contract.EigenLayerDelayedWithdrawalRouter(&_Restakingpod.CallOpts)
}

// EigenLayerDelegationManager is a free data retrieval call binding the contract method 0x23e127d5.
//
// Solidity: function eigenLayerDelegationManager() view returns(address)
func (_Restakingpod *RestakingpodCaller) EigenLayerDelegationManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Restakingpod.contract.Call(opts, &out, "eigenLayerDelegationManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EigenLayerDelegationManager is a free data retrieval call binding the contract method 0x23e127d5.
//
// Solidity: function eigenLayerDelegationManager() view returns(address)
func (_Restakingpod *RestakingpodSession) EigenLayerDelegationManager() (common.Address, error) {
	return _Restakingpod.Contract.EigenLayerDelegationManager(&_Restakingpod.CallOpts)
}

// EigenLayerDelegationManager is a free data retrieval call binding the contract method 0x23e127d5.
//
// Solidity: function eigenLayerDelegationManager() view returns(address)
func (_Restakingpod *RestakingpodCallerSession) EigenLayerDelegationManager() (common.Address, error) {
	return _Restakingpod.Contract.EigenLayerDelegationManager(&_Restakingpod.CallOpts)
}

// EigenLayerEigenPod is a free data retrieval call binding the contract method 0x6b0cbbd4.
//
// Solidity: function eigenLayerEigenPod() view returns(address)
func (_Restakingpod *RestakingpodCaller) EigenLayerEigenPod(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Restakingpod.contract.Call(opts, &out, "eigenLayerEigenPod")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EigenLayerEigenPod is a free data retrieval call binding the contract method 0x6b0cbbd4.
//
// Solidity: function eigenLayerEigenPod() view returns(address)
func (_Restakingpod *RestakingpodSession) EigenLayerEigenPod() (common.Address, error) {
	return _Restakingpod.Contract.EigenLayerEigenPod(&_Restakingpod.CallOpts)
}

// EigenLayerEigenPod is a free data retrieval call binding the contract method 0x6b0cbbd4.
//
// Solidity: function eigenLayerEigenPod() view returns(address)
func (_Restakingpod *RestakingpodCallerSession) EigenLayerEigenPod() (common.Address, error) {
	return _Restakingpod.Contract.EigenLayerEigenPod(&_Restakingpod.CallOpts)
}

// EigenLayerEigenPodManager is a free data retrieval call binding the contract method 0x2ffb004f.
//
// Solidity: function eigenLayerEigenPodManager() view returns(address)
func (_Restakingpod *RestakingpodCaller) EigenLayerEigenPodManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Restakingpod.contract.Call(opts, &out, "eigenLayerEigenPodManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EigenLayerEigenPodManager is a free data retrieval call binding the contract method 0x2ffb004f.
//
// Solidity: function eigenLayerEigenPodManager() view returns(address)
func (_Restakingpod *RestakingpodSession) EigenLayerEigenPodManager() (common.Address, error) {
	return _Restakingpod.Contract.EigenLayerEigenPodManager(&_Restakingpod.CallOpts)
}

// EigenLayerEigenPodManager is a free data retrieval call binding the contract method 0x2ffb004f.
//
// Solidity: function eigenLayerEigenPodManager() view returns(address)
func (_Restakingpod *RestakingpodCallerSession) EigenLayerEigenPodManager() (common.Address, error) {
	return _Restakingpod.Contract.EigenLayerEigenPodManager(&_Restakingpod.CallOpts)
}

// EigenLayerOperator is a free data retrieval call binding the contract method 0xf26668d7.
//
// Solidity: function eigenLayerOperator() view returns(address)
func (_Restakingpod *RestakingpodCaller) EigenLayerOperator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Restakingpod.contract.Call(opts, &out, "eigenLayerOperator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EigenLayerOperator is a free data retrieval call binding the contract method 0xf26668d7.
//
// Solidity: function eigenLayerOperator() view returns(address)
func (_Restakingpod *RestakingpodSession) EigenLayerOperator() (common.Address, error) {
	return _Restakingpod.Contract.EigenLayerOperator(&_Restakingpod.CallOpts)
}

// EigenLayerOperator is a free data retrieval call binding the contract method 0xf26668d7.
//
// Solidity: function eigenLayerOperator() view returns(address)
func (_Restakingpod *RestakingpodCallerSession) EigenLayerOperator() (common.Address, error) {
	return _Restakingpod.Contract.EigenLayerOperator(&_Restakingpod.CallOpts)
}

// GetAllDelayedWithdrawals is a free data retrieval call binding the contract method 0x2315c805.
//
// Solidity: function getAllDelayedWithdrawals() view returns(uint256)
func (_Restakingpod *RestakingpodCaller) GetAllDelayedWithdrawals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Restakingpod.contract.Call(opts, &out, "getAllDelayedWithdrawals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAllDelayedWithdrawals is a free data retrieval call binding the contract method 0x2315c805.
//
// Solidity: function getAllDelayedWithdrawals() view returns(uint256)
func (_Restakingpod *RestakingpodSession) GetAllDelayedWithdrawals() (*big.Int, error) {
	return _Restakingpod.Contract.GetAllDelayedWithdrawals(&_Restakingpod.CallOpts)
}

// GetAllDelayedWithdrawals is a free data retrieval call binding the contract method 0x2315c805.
//
// Solidity: function getAllDelayedWithdrawals() view returns(uint256)
func (_Restakingpod *RestakingpodCallerSession) GetAllDelayedWithdrawals() (*big.Int, error) {
	return _Restakingpod.Contract.GetAllDelayedWithdrawals(&_Restakingpod.CallOpts)
}

// GetClaimableDelayedWithdrawals is a free data retrieval call binding the contract method 0x01129226.
//
// Solidity: function getClaimableDelayedWithdrawals() view returns(uint256)
func (_Restakingpod *RestakingpodCaller) GetClaimableDelayedWithdrawals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Restakingpod.contract.Call(opts, &out, "getClaimableDelayedWithdrawals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetClaimableDelayedWithdrawals is a free data retrieval call binding the contract method 0x01129226.
//
// Solidity: function getClaimableDelayedWithdrawals() view returns(uint256)
func (_Restakingpod *RestakingpodSession) GetClaimableDelayedWithdrawals() (*big.Int, error) {
	return _Restakingpod.Contract.GetClaimableDelayedWithdrawals(&_Restakingpod.CallOpts)
}

// GetClaimableDelayedWithdrawals is a free data retrieval call binding the contract method 0x01129226.
//
// Solidity: function getClaimableDelayedWithdrawals() view returns(uint256)
func (_Restakingpod *RestakingpodCallerSession) GetClaimableDelayedWithdrawals() (*big.Int, error) {
	return _Restakingpod.Contract.GetClaimableDelayedWithdrawals(&_Restakingpod.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_Restakingpod *RestakingpodCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Restakingpod.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_Restakingpod *RestakingpodSession) Implementation() (common.Address, error) {
	return _Restakingpod.Contract.Implementation(&_Restakingpod.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_Restakingpod *RestakingpodCallerSession) Implementation() (common.Address, error) {
	return _Restakingpod.Contract.Implementation(&_Restakingpod.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Restakingpod *RestakingpodCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Restakingpod.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Restakingpod *RestakingpodSession) Owner() (common.Address, error) {
	return _Restakingpod.Contract.Owner(&_Restakingpod.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Restakingpod *RestakingpodCallerSession) Owner() (common.Address, error) {
	return _Restakingpod.Contract.Owner(&_Restakingpod.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Restakingpod *RestakingpodCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Restakingpod.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Restakingpod *RestakingpodSession) Paused() (bool, error) {
	return _Restakingpod.Contract.Paused(&_Restakingpod.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Restakingpod *RestakingpodCallerSession) Paused() (bool, error) {
	return _Restakingpod.Contract.Paused(&_Restakingpod.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Restakingpod *RestakingpodCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Restakingpod.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Restakingpod *RestakingpodSession) ProxiableUUID() ([32]byte, error) {
	return _Restakingpod.Contract.ProxiableUUID(&_Restakingpod.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Restakingpod *RestakingpodCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Restakingpod.Contract.ProxiableUUID(&_Restakingpod.CallOpts)
}

// RestakingPodManager is a free data retrieval call binding the contract method 0xc824475a.
//
// Solidity: function restakingPodManager() view returns(address)
func (_Restakingpod *RestakingpodCaller) RestakingPodManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Restakingpod.contract.Call(opts, &out, "restakingPodManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RestakingPodManager is a free data retrieval call binding the contract method 0xc824475a.
//
// Solidity: function restakingPodManager() view returns(address)
func (_Restakingpod *RestakingpodSession) RestakingPodManager() (common.Address, error) {
	return _Restakingpod.Contract.RestakingPodManager(&_Restakingpod.CallOpts)
}

// RestakingPodManager is a free data retrieval call binding the contract method 0xc824475a.
//
// Solidity: function restakingPodManager() view returns(address)
func (_Restakingpod *RestakingpodCallerSession) RestakingPodManager() (common.Address, error) {
	return _Restakingpod.Contract.RestakingPodManager(&_Restakingpod.CallOpts)
}

// RestakingPool is a free data retrieval call binding the contract method 0xdf604b28.
//
// Solidity: function restakingPool() view returns(address)
func (_Restakingpod *RestakingpodCaller) RestakingPool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Restakingpod.contract.Call(opts, &out, "restakingPool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RestakingPool is a free data retrieval call binding the contract method 0xdf604b28.
//
// Solidity: function restakingPool() view returns(address)
func (_Restakingpod *RestakingpodSession) RestakingPool() (common.Address, error) {
	return _Restakingpod.Contract.RestakingPool(&_Restakingpod.CallOpts)
}

// RestakingPool is a free data retrieval call binding the contract method 0xdf604b28.
//
// Solidity: function restakingPool() view returns(address)
func (_Restakingpod *RestakingpodCallerSession) RestakingPool() (common.Address, error) {
	return _Restakingpod.Contract.RestakingPool(&_Restakingpod.CallOpts)
}

// StakedButNotVerifiedEth is a free data retrieval call binding the contract method 0x397bfbac.
//
// Solidity: function stakedButNotVerifiedEth() view returns(uint256)
func (_Restakingpod *RestakingpodCaller) StakedButNotVerifiedEth(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Restakingpod.contract.Call(opts, &out, "stakedButNotVerifiedEth")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakedButNotVerifiedEth is a free data retrieval call binding the contract method 0x397bfbac.
//
// Solidity: function stakedButNotVerifiedEth() view returns(uint256)
func (_Restakingpod *RestakingpodSession) StakedButNotVerifiedEth() (*big.Int, error) {
	return _Restakingpod.Contract.StakedButNotVerifiedEth(&_Restakingpod.CallOpts)
}

// StakedButNotVerifiedEth is a free data retrieval call binding the contract method 0x397bfbac.
//
// Solidity: function stakedButNotVerifiedEth() view returns(uint256)
func (_Restakingpod *RestakingpodCallerSession) StakedButNotVerifiedEth() (*big.Int, error) {
	return _Restakingpod.Contract.StakedButNotVerifiedEth(&_Restakingpod.CallOpts)
}

// TypeId is a free data retrieval call binding the contract method 0x8f940f63.
//
// Solidity: function typeId() pure returns(bytes32)
func (_Restakingpod *RestakingpodCaller) TypeId(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Restakingpod.contract.Call(opts, &out, "typeId")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TypeId is a free data retrieval call binding the contract method 0x8f940f63.
//
// Solidity: function typeId() pure returns(bytes32)
func (_Restakingpod *RestakingpodSession) TypeId() ([32]byte, error) {
	return _Restakingpod.Contract.TypeId(&_Restakingpod.CallOpts)
}

// TypeId is a free data retrieval call binding the contract method 0x8f940f63.
//
// Solidity: function typeId() pure returns(bytes32)
func (_Restakingpod *RestakingpodCallerSession) TypeId() ([32]byte, error) {
	return _Restakingpod.Contract.TypeId(&_Restakingpod.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint8)
func (_Restakingpod *RestakingpodCaller) Version(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Restakingpod.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint8)
func (_Restakingpod *RestakingpodSession) Version() (uint8, error) {
	return _Restakingpod.Contract.Version(&_Restakingpod.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint8)
func (_Restakingpod *RestakingpodCallerSession) Version() (uint8, error) {
	return _Restakingpod.Contract.Version(&_Restakingpod.CallOpts)
}

// WithdrawCredentials is a free data retrieval call binding the contract method 0xb62be945.
//
// Solidity: function withdrawCredentials() view returns(bytes)
func (_Restakingpod *RestakingpodCaller) WithdrawCredentials(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Restakingpod.contract.Call(opts, &out, "withdrawCredentials")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// WithdrawCredentials is a free data retrieval call binding the contract method 0xb62be945.
//
// Solidity: function withdrawCredentials() view returns(bytes)
func (_Restakingpod *RestakingpodSession) WithdrawCredentials() ([]byte, error) {
	return _Restakingpod.Contract.WithdrawCredentials(&_Restakingpod.CallOpts)
}

// WithdrawCredentials is a free data retrieval call binding the contract method 0xb62be945.
//
// Solidity: function withdrawCredentials() view returns(bytes)
func (_Restakingpod *RestakingpodCallerSession) WithdrawCredentials() ([]byte, error) {
	return _Restakingpod.Contract.WithdrawCredentials(&_Restakingpod.CallOpts)
}

// ActivateRestaking is a paid mutator transaction binding the contract method 0x0cd4649e.
//
// Solidity: function activateRestaking() returns()
func (_Restakingpod *RestakingpodTransactor) ActivateRestaking(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Restakingpod.contract.Transact(opts, "activateRestaking")
}

// ActivateRestaking is a paid mutator transaction binding the contract method 0x0cd4649e.
//
// Solidity: function activateRestaking() returns()
func (_Restakingpod *RestakingpodSession) ActivateRestaking() (*types.Transaction, error) {
	return _Restakingpod.Contract.ActivateRestaking(&_Restakingpod.TransactOpts)
}

// ActivateRestaking is a paid mutator transaction binding the contract method 0x0cd4649e.
//
// Solidity: function activateRestaking() returns()
func (_Restakingpod *RestakingpodTransactorSession) ActivateRestaking() (*types.Transaction, error) {
	return _Restakingpod.Contract.ActivateRestaking(&_Restakingpod.TransactOpts)
}

// ClaimDelayedWithdrawals is a paid mutator transaction binding the contract method 0x28837c07.
//
// Solidity: function claimDelayedWithdrawals() returns()
func (_Restakingpod *RestakingpodTransactor) ClaimDelayedWithdrawals(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Restakingpod.contract.Transact(opts, "claimDelayedWithdrawals")
}

// ClaimDelayedWithdrawals is a paid mutator transaction binding the contract method 0x28837c07.
//
// Solidity: function claimDelayedWithdrawals() returns()
func (_Restakingpod *RestakingpodSession) ClaimDelayedWithdrawals() (*types.Transaction, error) {
	return _Restakingpod.Contract.ClaimDelayedWithdrawals(&_Restakingpod.TransactOpts)
}

// ClaimDelayedWithdrawals is a paid mutator transaction binding the contract method 0x28837c07.
//
// Solidity: function claimDelayedWithdrawals() returns()
func (_Restakingpod *RestakingpodTransactorSession) ClaimDelayedWithdrawals() (*types.Transaction, error) {
	return _Restakingpod.Contract.ClaimDelayedWithdrawals(&_Restakingpod.TransactOpts)
}

// CompleteQueuedWithdrawals is a paid mutator transaction binding the contract method 0x33404396.
//
// Solidity: function completeQueuedWithdrawals((address,address,address,uint256,uint32,address[],uint256[])[] withdrawals, address[][] tokens, uint256[] middlewareTimesIndexes, bool[] receiveAsTokens) returns()
func (_Restakingpod *RestakingpodTransactor) CompleteQueuedWithdrawals(opts *bind.TransactOpts, withdrawals []IDelegationManagerWithdrawal, tokens [][]common.Address, middlewareTimesIndexes []*big.Int, receiveAsTokens []bool) (*types.Transaction, error) {
	return _Restakingpod.contract.Transact(opts, "completeQueuedWithdrawals", withdrawals, tokens, middlewareTimesIndexes, receiveAsTokens)
}

// CompleteQueuedWithdrawals is a paid mutator transaction binding the contract method 0x33404396.
//
// Solidity: function completeQueuedWithdrawals((address,address,address,uint256,uint32,address[],uint256[])[] withdrawals, address[][] tokens, uint256[] middlewareTimesIndexes, bool[] receiveAsTokens) returns()
func (_Restakingpod *RestakingpodSession) CompleteQueuedWithdrawals(withdrawals []IDelegationManagerWithdrawal, tokens [][]common.Address, middlewareTimesIndexes []*big.Int, receiveAsTokens []bool) (*types.Transaction, error) {
	return _Restakingpod.Contract.CompleteQueuedWithdrawals(&_Restakingpod.TransactOpts, withdrawals, tokens, middlewareTimesIndexes, receiveAsTokens)
}

// CompleteQueuedWithdrawals is a paid mutator transaction binding the contract method 0x33404396.
//
// Solidity: function completeQueuedWithdrawals((address,address,address,uint256,uint32,address[],uint256[])[] withdrawals, address[][] tokens, uint256[] middlewareTimesIndexes, bool[] receiveAsTokens) returns()
func (_Restakingpod *RestakingpodTransactorSession) CompleteQueuedWithdrawals(withdrawals []IDelegationManagerWithdrawal, tokens [][]common.Address, middlewareTimesIndexes []*big.Int, receiveAsTokens []bool) (*types.Transaction, error) {
	return _Restakingpod.Contract.CompleteQueuedWithdrawals(&_Restakingpod.TransactOpts, withdrawals, tokens, middlewareTimesIndexes, receiveAsTokens)
}

// DelegateToEigenLayerOperator is a paid mutator transaction binding the contract method 0xd682e542.
//
// Solidity: function delegateToEigenLayerOperator(address _eigenLayerOperator, (bytes,uint256) approverSignatureAndExpiry, bytes32 approverSalt) returns()
func (_Restakingpod *RestakingpodTransactor) DelegateToEigenLayerOperator(opts *bind.TransactOpts, _eigenLayerOperator common.Address, approverSignatureAndExpiry ISignatureUtilsSignatureWithExpiry, approverSalt [32]byte) (*types.Transaction, error) {
	return _Restakingpod.contract.Transact(opts, "delegateToEigenLayerOperator", _eigenLayerOperator, approverSignatureAndExpiry, approverSalt)
}

// DelegateToEigenLayerOperator is a paid mutator transaction binding the contract method 0xd682e542.
//
// Solidity: function delegateToEigenLayerOperator(address _eigenLayerOperator, (bytes,uint256) approverSignatureAndExpiry, bytes32 approverSalt) returns()
func (_Restakingpod *RestakingpodSession) DelegateToEigenLayerOperator(_eigenLayerOperator common.Address, approverSignatureAndExpiry ISignatureUtilsSignatureWithExpiry, approverSalt [32]byte) (*types.Transaction, error) {
	return _Restakingpod.Contract.DelegateToEigenLayerOperator(&_Restakingpod.TransactOpts, _eigenLayerOperator, approverSignatureAndExpiry, approverSalt)
}

// DelegateToEigenLayerOperator is a paid mutator transaction binding the contract method 0xd682e542.
//
// Solidity: function delegateToEigenLayerOperator(address _eigenLayerOperator, (bytes,uint256) approverSignatureAndExpiry, bytes32 approverSalt) returns()
func (_Restakingpod *RestakingpodTransactorSession) DelegateToEigenLayerOperator(_eigenLayerOperator common.Address, approverSignatureAndExpiry ISignatureUtilsSignatureWithExpiry, approverSalt [32]byte) (*types.Transaction, error) {
	return _Restakingpod.Contract.DelegateToEigenLayerOperator(&_Restakingpod.TransactOpts, _eigenLayerOperator, approverSignatureAndExpiry, approverSalt)
}

// Execute is a paid mutator transaction binding the contract method 0x4d20b887.
//
// Solidity: function execute(uint256 value, address to, bytes data, uint256 txGas) returns(bool success)
func (_Restakingpod *RestakingpodTransactor) Execute(opts *bind.TransactOpts, value *big.Int, to common.Address, data []byte, txGas *big.Int) (*types.Transaction, error) {
	return _Restakingpod.contract.Transact(opts, "execute", value, to, data, txGas)
}

// Execute is a paid mutator transaction binding the contract method 0x4d20b887.
//
// Solidity: function execute(uint256 value, address to, bytes data, uint256 txGas) returns(bool success)
func (_Restakingpod *RestakingpodSession) Execute(value *big.Int, to common.Address, data []byte, txGas *big.Int) (*types.Transaction, error) {
	return _Restakingpod.Contract.Execute(&_Restakingpod.TransactOpts, value, to, data, txGas)
}

// Execute is a paid mutator transaction binding the contract method 0x4d20b887.
//
// Solidity: function execute(uint256 value, address to, bytes data, uint256 txGas) returns(bool success)
func (_Restakingpod *RestakingpodTransactorSession) Execute(value *big.Int, to common.Address, data []byte, txGas *big.Int) (*types.Transaction, error) {
	return _Restakingpod.Contract.Execute(&_Restakingpod.TransactOpts, value, to, data, txGas)
}

// Initialize is a paid mutator transaction binding the contract method 0x35876476.
//
// Solidity: function initialize(address _ownerAddr, address _dao, address _restakingPool, address _restakingPodManager, address _eigenLayerEigenPodManager, address _eigenLayerDelegationManager, address _eigenLayerDelayedWithdrawalRouter) returns()
func (_Restakingpod *RestakingpodTransactor) Initialize(opts *bind.TransactOpts, _ownerAddr common.Address, _dao common.Address, _restakingPool common.Address, _restakingPodManager common.Address, _eigenLayerEigenPodManager common.Address, _eigenLayerDelegationManager common.Address, _eigenLayerDelayedWithdrawalRouter common.Address) (*types.Transaction, error) {
	return _Restakingpod.contract.Transact(opts, "initialize", _ownerAddr, _dao, _restakingPool, _restakingPodManager, _eigenLayerEigenPodManager, _eigenLayerDelegationManager, _eigenLayerDelayedWithdrawalRouter)
}

// Initialize is a paid mutator transaction binding the contract method 0x35876476.
//
// Solidity: function initialize(address _ownerAddr, address _dao, address _restakingPool, address _restakingPodManager, address _eigenLayerEigenPodManager, address _eigenLayerDelegationManager, address _eigenLayerDelayedWithdrawalRouter) returns()
func (_Restakingpod *RestakingpodSession) Initialize(_ownerAddr common.Address, _dao common.Address, _restakingPool common.Address, _restakingPodManager common.Address, _eigenLayerEigenPodManager common.Address, _eigenLayerDelegationManager common.Address, _eigenLayerDelayedWithdrawalRouter common.Address) (*types.Transaction, error) {
	return _Restakingpod.Contract.Initialize(&_Restakingpod.TransactOpts, _ownerAddr, _dao, _restakingPool, _restakingPodManager, _eigenLayerEigenPodManager, _eigenLayerDelegationManager, _eigenLayerDelayedWithdrawalRouter)
}

// Initialize is a paid mutator transaction binding the contract method 0x35876476.
//
// Solidity: function initialize(address _ownerAddr, address _dao, address _restakingPool, address _restakingPodManager, address _eigenLayerEigenPodManager, address _eigenLayerDelegationManager, address _eigenLayerDelayedWithdrawalRouter) returns()
func (_Restakingpod *RestakingpodTransactorSession) Initialize(_ownerAddr common.Address, _dao common.Address, _restakingPool common.Address, _restakingPodManager common.Address, _eigenLayerEigenPodManager common.Address, _eigenLayerDelegationManager common.Address, _eigenLayerDelayedWithdrawalRouter common.Address) (*types.Transaction, error) {
	return _Restakingpod.Contract.Initialize(&_Restakingpod.TransactOpts, _ownerAddr, _dao, _restakingPool, _restakingPodManager, _eigenLayerEigenPodManager, _eigenLayerDelegationManager, _eigenLayerDelayedWithdrawalRouter)
}

// QueueWithdrawals is a paid mutator transaction binding the contract method 0x0dd8dd02.
//
// Solidity: function queueWithdrawals((address[],uint256[],address)[] queuedWithdrawalParams) returns(bytes32[])
func (_Restakingpod *RestakingpodTransactor) QueueWithdrawals(opts *bind.TransactOpts, queuedWithdrawalParams []IDelegationManagerQueuedWithdrawalParams) (*types.Transaction, error) {
	return _Restakingpod.contract.Transact(opts, "queueWithdrawals", queuedWithdrawalParams)
}

// QueueWithdrawals is a paid mutator transaction binding the contract method 0x0dd8dd02.
//
// Solidity: function queueWithdrawals((address[],uint256[],address)[] queuedWithdrawalParams) returns(bytes32[])
func (_Restakingpod *RestakingpodSession) QueueWithdrawals(queuedWithdrawalParams []IDelegationManagerQueuedWithdrawalParams) (*types.Transaction, error) {
	return _Restakingpod.Contract.QueueWithdrawals(&_Restakingpod.TransactOpts, queuedWithdrawalParams)
}

// QueueWithdrawals is a paid mutator transaction binding the contract method 0x0dd8dd02.
//
// Solidity: function queueWithdrawals((address[],uint256[],address)[] queuedWithdrawalParams) returns(bytes32[])
func (_Restakingpod *RestakingpodTransactorSession) QueueWithdrawals(queuedWithdrawalParams []IDelegationManagerQueuedWithdrawalParams) (*types.Transaction, error) {
	return _Restakingpod.Contract.QueueWithdrawals(&_Restakingpod.TransactOpts, queuedWithdrawalParams)
}

// RecoverTokens is a paid mutator transaction binding the contract method 0xdda3346c.
//
// Solidity: function recoverTokens(address[] tokenList, uint256[] amountsToWithdraw, address recipient) returns()
func (_Restakingpod *RestakingpodTransactor) RecoverTokens(opts *bind.TransactOpts, tokenList []common.Address, amountsToWithdraw []*big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Restakingpod.contract.Transact(opts, "recoverTokens", tokenList, amountsToWithdraw, recipient)
}

// RecoverTokens is a paid mutator transaction binding the contract method 0xdda3346c.
//
// Solidity: function recoverTokens(address[] tokenList, uint256[] amountsToWithdraw, address recipient) returns()
func (_Restakingpod *RestakingpodSession) RecoverTokens(tokenList []common.Address, amountsToWithdraw []*big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Restakingpod.Contract.RecoverTokens(&_Restakingpod.TransactOpts, tokenList, amountsToWithdraw, recipient)
}

// RecoverTokens is a paid mutator transaction binding the contract method 0xdda3346c.
//
// Solidity: function recoverTokens(address[] tokenList, uint256[] amountsToWithdraw, address recipient) returns()
func (_Restakingpod *RestakingpodTransactorSession) RecoverTokens(tokenList []common.Address, amountsToWithdraw []*big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Restakingpod.Contract.RecoverTokens(&_Restakingpod.TransactOpts, tokenList, amountsToWithdraw, recipient)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Restakingpod *RestakingpodTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Restakingpod.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Restakingpod *RestakingpodSession) RenounceOwnership() (*types.Transaction, error) {
	return _Restakingpod.Contract.RenounceOwnership(&_Restakingpod.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Restakingpod *RestakingpodTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Restakingpod.Contract.RenounceOwnership(&_Restakingpod.TransactOpts)
}

// SetDao is a paid mutator transaction binding the contract method 0x6637b882.
//
// Solidity: function setDao(address _dao) returns()
func (_Restakingpod *RestakingpodTransactor) SetDao(opts *bind.TransactOpts, _dao common.Address) (*types.Transaction, error) {
	return _Restakingpod.contract.Transact(opts, "setDao", _dao)
}

// SetDao is a paid mutator transaction binding the contract method 0x6637b882.
//
// Solidity: function setDao(address _dao) returns()
func (_Restakingpod *RestakingpodSession) SetDao(_dao common.Address) (*types.Transaction, error) {
	return _Restakingpod.Contract.SetDao(&_Restakingpod.TransactOpts, _dao)
}

// SetDao is a paid mutator transaction binding the contract method 0x6637b882.
//
// Solidity: function setDao(address _dao) returns()
func (_Restakingpod *RestakingpodTransactorSession) SetDao(_dao common.Address) (*types.Transaction, error) {
	return _Restakingpod.Contract.SetDao(&_Restakingpod.TransactOpts, _dao)
}

// SetRestakingPodManager is a paid mutator transaction binding the contract method 0x6d387265.
//
// Solidity: function setRestakingPodManager(address _restakingPodManager) returns()
func (_Restakingpod *RestakingpodTransactor) SetRestakingPodManager(opts *bind.TransactOpts, _restakingPodManager common.Address) (*types.Transaction, error) {
	return _Restakingpod.contract.Transact(opts, "setRestakingPodManager", _restakingPodManager)
}

// SetRestakingPodManager is a paid mutator transaction binding the contract method 0x6d387265.
//
// Solidity: function setRestakingPodManager(address _restakingPodManager) returns()
func (_Restakingpod *RestakingpodSession) SetRestakingPodManager(_restakingPodManager common.Address) (*types.Transaction, error) {
	return _Restakingpod.Contract.SetRestakingPodManager(&_Restakingpod.TransactOpts, _restakingPodManager)
}

// SetRestakingPodManager is a paid mutator transaction binding the contract method 0x6d387265.
//
// Solidity: function setRestakingPodManager(address _restakingPodManager) returns()
func (_Restakingpod *RestakingpodTransactorSession) SetRestakingPodManager(_restakingPodManager common.Address) (*types.Transaction, error) {
	return _Restakingpod.Contract.SetRestakingPodManager(&_Restakingpod.TransactOpts, _restakingPodManager)
}

// SetStakedButNotVerifiedEth is a paid mutator transaction binding the contract method 0x1fa23e85.
//
// Solidity: function setStakedButNotVerifiedEth(uint256 _amount) returns()
func (_Restakingpod *RestakingpodTransactor) SetStakedButNotVerifiedEth(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Restakingpod.contract.Transact(opts, "setStakedButNotVerifiedEth", _amount)
}

// SetStakedButNotVerifiedEth is a paid mutator transaction binding the contract method 0x1fa23e85.
//
// Solidity: function setStakedButNotVerifiedEth(uint256 _amount) returns()
func (_Restakingpod *RestakingpodSession) SetStakedButNotVerifiedEth(_amount *big.Int) (*types.Transaction, error) {
	return _Restakingpod.Contract.SetStakedButNotVerifiedEth(&_Restakingpod.TransactOpts, _amount)
}

// SetStakedButNotVerifiedEth is a paid mutator transaction binding the contract method 0x1fa23e85.
//
// Solidity: function setStakedButNotVerifiedEth(uint256 _amount) returns()
func (_Restakingpod *RestakingpodTransactorSession) SetStakedButNotVerifiedEth(_amount *big.Int) (*types.Transaction, error) {
	return _Restakingpod.Contract.SetStakedButNotVerifiedEth(&_Restakingpod.TransactOpts, _amount)
}

// Stake is a paid mutator transaction binding the contract method 0x9b4e4634.
//
// Solidity: function stake(bytes _pubkey, bytes _signature, bytes32 _depositDataRoot) payable returns()
func (_Restakingpod *RestakingpodTransactor) Stake(opts *bind.TransactOpts, _pubkey []byte, _signature []byte, _depositDataRoot [32]byte) (*types.Transaction, error) {
	return _Restakingpod.contract.Transact(opts, "stake", _pubkey, _signature, _depositDataRoot)
}

// Stake is a paid mutator transaction binding the contract method 0x9b4e4634.
//
// Solidity: function stake(bytes _pubkey, bytes _signature, bytes32 _depositDataRoot) payable returns()
func (_Restakingpod *RestakingpodSession) Stake(_pubkey []byte, _signature []byte, _depositDataRoot [32]byte) (*types.Transaction, error) {
	return _Restakingpod.Contract.Stake(&_Restakingpod.TransactOpts, _pubkey, _signature, _depositDataRoot)
}

// Stake is a paid mutator transaction binding the contract method 0x9b4e4634.
//
// Solidity: function stake(bytes _pubkey, bytes _signature, bytes32 _depositDataRoot) payable returns()
func (_Restakingpod *RestakingpodTransactorSession) Stake(_pubkey []byte, _signature []byte, _depositDataRoot [32]byte) (*types.Transaction, error) {
	return _Restakingpod.Contract.Stake(&_Restakingpod.TransactOpts, _pubkey, _signature, _depositDataRoot)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Restakingpod *RestakingpodTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Restakingpod.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Restakingpod *RestakingpodSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Restakingpod.Contract.TransferOwnership(&_Restakingpod.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Restakingpod *RestakingpodTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Restakingpod.Contract.TransferOwnership(&_Restakingpod.TransactOpts, newOwner)
}

// UndelegateEigenLayerOperator is a paid mutator transaction binding the contract method 0x18501962.
//
// Solidity: function undelegateEigenLayerOperator() returns()
func (_Restakingpod *RestakingpodTransactor) UndelegateEigenLayerOperator(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Restakingpod.contract.Transact(opts, "undelegateEigenLayerOperator")
}

// UndelegateEigenLayerOperator is a paid mutator transaction binding the contract method 0x18501962.
//
// Solidity: function undelegateEigenLayerOperator() returns()
func (_Restakingpod *RestakingpodSession) UndelegateEigenLayerOperator() (*types.Transaction, error) {
	return _Restakingpod.Contract.UndelegateEigenLayerOperator(&_Restakingpod.TransactOpts)
}

// UndelegateEigenLayerOperator is a paid mutator transaction binding the contract method 0x18501962.
//
// Solidity: function undelegateEigenLayerOperator() returns()
func (_Restakingpod *RestakingpodTransactorSession) UndelegateEigenLayerOperator() (*types.Transaction, error) {
	return _Restakingpod.Contract.UndelegateEigenLayerOperator(&_Restakingpod.TransactOpts)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Restakingpod *RestakingpodTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Restakingpod.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Restakingpod *RestakingpodSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Restakingpod.Contract.UpgradeTo(&_Restakingpod.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Restakingpod *RestakingpodTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Restakingpod.Contract.UpgradeTo(&_Restakingpod.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Restakingpod *RestakingpodTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Restakingpod.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Restakingpod *RestakingpodSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Restakingpod.Contract.UpgradeToAndCall(&_Restakingpod.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Restakingpod *RestakingpodTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Restakingpod.Contract.UpgradeToAndCall(&_Restakingpod.TransactOpts, newImplementation, data)
}

// VerifyAndProcessWithdrawals is a paid mutator transaction binding the contract method 0xe251ef52.
//
// Solidity: function verifyAndProcessWithdrawals(uint64 oracleTimestamp, (bytes32,bytes) stateRootProof, (bytes,bytes,bytes,bytes,bytes,uint64,uint64,uint64,bytes32,bytes32,bytes32,bytes32)[] withdrawalProofs, bytes[] validatorFieldsProofs, bytes32[][] validatorFields, bytes32[][] withdrawalFields) returns()
func (_Restakingpod *RestakingpodTransactor) VerifyAndProcessWithdrawals(opts *bind.TransactOpts, oracleTimestamp uint64, stateRootProof BeaconChainProofsStateRootProof, withdrawalProofs []BeaconChainProofsWithdrawalProof, validatorFieldsProofs [][]byte, validatorFields [][][32]byte, withdrawalFields [][][32]byte) (*types.Transaction, error) {
	return _Restakingpod.contract.Transact(opts, "verifyAndProcessWithdrawals", oracleTimestamp, stateRootProof, withdrawalProofs, validatorFieldsProofs, validatorFields, withdrawalFields)
}

// VerifyAndProcessWithdrawals is a paid mutator transaction binding the contract method 0xe251ef52.
//
// Solidity: function verifyAndProcessWithdrawals(uint64 oracleTimestamp, (bytes32,bytes) stateRootProof, (bytes,bytes,bytes,bytes,bytes,uint64,uint64,uint64,bytes32,bytes32,bytes32,bytes32)[] withdrawalProofs, bytes[] validatorFieldsProofs, bytes32[][] validatorFields, bytes32[][] withdrawalFields) returns()
func (_Restakingpod *RestakingpodSession) VerifyAndProcessWithdrawals(oracleTimestamp uint64, stateRootProof BeaconChainProofsStateRootProof, withdrawalProofs []BeaconChainProofsWithdrawalProof, validatorFieldsProofs [][]byte, validatorFields [][][32]byte, withdrawalFields [][][32]byte) (*types.Transaction, error) {
	return _Restakingpod.Contract.VerifyAndProcessWithdrawals(&_Restakingpod.TransactOpts, oracleTimestamp, stateRootProof, withdrawalProofs, validatorFieldsProofs, validatorFields, withdrawalFields)
}

// VerifyAndProcessWithdrawals is a paid mutator transaction binding the contract method 0xe251ef52.
//
// Solidity: function verifyAndProcessWithdrawals(uint64 oracleTimestamp, (bytes32,bytes) stateRootProof, (bytes,bytes,bytes,bytes,bytes,uint64,uint64,uint64,bytes32,bytes32,bytes32,bytes32)[] withdrawalProofs, bytes[] validatorFieldsProofs, bytes32[][] validatorFields, bytes32[][] withdrawalFields) returns()
func (_Restakingpod *RestakingpodTransactorSession) VerifyAndProcessWithdrawals(oracleTimestamp uint64, stateRootProof BeaconChainProofsStateRootProof, withdrawalProofs []BeaconChainProofsWithdrawalProof, validatorFieldsProofs [][]byte, validatorFields [][][32]byte, withdrawalFields [][][32]byte) (*types.Transaction, error) {
	return _Restakingpod.Contract.VerifyAndProcessWithdrawals(&_Restakingpod.TransactOpts, oracleTimestamp, stateRootProof, withdrawalProofs, validatorFieldsProofs, validatorFields, withdrawalFields)
}

// VerifyBalanceUpdates is a paid mutator transaction binding the contract method 0xa50600f4.
//
// Solidity: function verifyBalanceUpdates(uint64 oracleTimestamp, uint40[] validatorIndices, (bytes32,bytes) stateRootProof, bytes[] validatorFieldsProofs, bytes32[][] validatorFields) returns()
func (_Restakingpod *RestakingpodTransactor) VerifyBalanceUpdates(opts *bind.TransactOpts, oracleTimestamp uint64, validatorIndices []*big.Int, stateRootProof BeaconChainProofsStateRootProof, validatorFieldsProofs [][]byte, validatorFields [][][32]byte) (*types.Transaction, error) {
	return _Restakingpod.contract.Transact(opts, "verifyBalanceUpdates", oracleTimestamp, validatorIndices, stateRootProof, validatorFieldsProofs, validatorFields)
}

// VerifyBalanceUpdates is a paid mutator transaction binding the contract method 0xa50600f4.
//
// Solidity: function verifyBalanceUpdates(uint64 oracleTimestamp, uint40[] validatorIndices, (bytes32,bytes) stateRootProof, bytes[] validatorFieldsProofs, bytes32[][] validatorFields) returns()
func (_Restakingpod *RestakingpodSession) VerifyBalanceUpdates(oracleTimestamp uint64, validatorIndices []*big.Int, stateRootProof BeaconChainProofsStateRootProof, validatorFieldsProofs [][]byte, validatorFields [][][32]byte) (*types.Transaction, error) {
	return _Restakingpod.Contract.VerifyBalanceUpdates(&_Restakingpod.TransactOpts, oracleTimestamp, validatorIndices, stateRootProof, validatorFieldsProofs, validatorFields)
}

// VerifyBalanceUpdates is a paid mutator transaction binding the contract method 0xa50600f4.
//
// Solidity: function verifyBalanceUpdates(uint64 oracleTimestamp, uint40[] validatorIndices, (bytes32,bytes) stateRootProof, bytes[] validatorFieldsProofs, bytes32[][] validatorFields) returns()
func (_Restakingpod *RestakingpodTransactorSession) VerifyBalanceUpdates(oracleTimestamp uint64, validatorIndices []*big.Int, stateRootProof BeaconChainProofsStateRootProof, validatorFieldsProofs [][]byte, validatorFields [][][32]byte) (*types.Transaction, error) {
	return _Restakingpod.Contract.VerifyBalanceUpdates(&_Restakingpod.TransactOpts, oracleTimestamp, validatorIndices, stateRootProof, validatorFieldsProofs, validatorFields)
}

// VerifyWithdrawalCredentials is a paid mutator transaction binding the contract method 0x3f65cf19.
//
// Solidity: function verifyWithdrawalCredentials(uint64 oracleTimestamp, (bytes32,bytes) stateRootProof, uint40[] validatorIndices, bytes[] validatorFieldsProofs, bytes32[][] validatorFields) returns()
func (_Restakingpod *RestakingpodTransactor) VerifyWithdrawalCredentials(opts *bind.TransactOpts, oracleTimestamp uint64, stateRootProof BeaconChainProofsStateRootProof, validatorIndices []*big.Int, validatorFieldsProofs [][]byte, validatorFields [][][32]byte) (*types.Transaction, error) {
	return _Restakingpod.contract.Transact(opts, "verifyWithdrawalCredentials", oracleTimestamp, stateRootProof, validatorIndices, validatorFieldsProofs, validatorFields)
}

// VerifyWithdrawalCredentials is a paid mutator transaction binding the contract method 0x3f65cf19.
//
// Solidity: function verifyWithdrawalCredentials(uint64 oracleTimestamp, (bytes32,bytes) stateRootProof, uint40[] validatorIndices, bytes[] validatorFieldsProofs, bytes32[][] validatorFields) returns()
func (_Restakingpod *RestakingpodSession) VerifyWithdrawalCredentials(oracleTimestamp uint64, stateRootProof BeaconChainProofsStateRootProof, validatorIndices []*big.Int, validatorFieldsProofs [][]byte, validatorFields [][][32]byte) (*types.Transaction, error) {
	return _Restakingpod.Contract.VerifyWithdrawalCredentials(&_Restakingpod.TransactOpts, oracleTimestamp, stateRootProof, validatorIndices, validatorFieldsProofs, validatorFields)
}

// VerifyWithdrawalCredentials is a paid mutator transaction binding the contract method 0x3f65cf19.
//
// Solidity: function verifyWithdrawalCredentials(uint64 oracleTimestamp, (bytes32,bytes) stateRootProof, uint40[] validatorIndices, bytes[] validatorFieldsProofs, bytes32[][] validatorFields) returns()
func (_Restakingpod *RestakingpodTransactorSession) VerifyWithdrawalCredentials(oracleTimestamp uint64, stateRootProof BeaconChainProofsStateRootProof, validatorIndices []*big.Int, validatorFieldsProofs [][]byte, validatorFields [][][32]byte) (*types.Transaction, error) {
	return _Restakingpod.Contract.VerifyWithdrawalCredentials(&_Restakingpod.TransactOpts, oracleTimestamp, stateRootProof, validatorIndices, validatorFieldsProofs, validatorFields)
}

// WithdrawBeforeRestaking is a paid mutator transaction binding the contract method 0xbaa7145a.
//
// Solidity: function withdrawBeforeRestaking() returns()
func (_Restakingpod *RestakingpodTransactor) WithdrawBeforeRestaking(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Restakingpod.contract.Transact(opts, "withdrawBeforeRestaking")
}

// WithdrawBeforeRestaking is a paid mutator transaction binding the contract method 0xbaa7145a.
//
// Solidity: function withdrawBeforeRestaking() returns()
func (_Restakingpod *RestakingpodSession) WithdrawBeforeRestaking() (*types.Transaction, error) {
	return _Restakingpod.Contract.WithdrawBeforeRestaking(&_Restakingpod.TransactOpts)
}

// WithdrawBeforeRestaking is a paid mutator transaction binding the contract method 0xbaa7145a.
//
// Solidity: function withdrawBeforeRestaking() returns()
func (_Restakingpod *RestakingpodTransactorSession) WithdrawBeforeRestaking() (*types.Transaction, error) {
	return _Restakingpod.Contract.WithdrawBeforeRestaking(&_Restakingpod.TransactOpts)
}

// WithdrawNonBeaconChainETHBalanceWei is a paid mutator transaction binding the contract method 0xe2c83445.
//
// Solidity: function withdrawNonBeaconChainETHBalanceWei(address recipient, uint256 amountToWithdraw) returns()
func (_Restakingpod *RestakingpodTransactor) WithdrawNonBeaconChainETHBalanceWei(opts *bind.TransactOpts, recipient common.Address, amountToWithdraw *big.Int) (*types.Transaction, error) {
	return _Restakingpod.contract.Transact(opts, "withdrawNonBeaconChainETHBalanceWei", recipient, amountToWithdraw)
}

// WithdrawNonBeaconChainETHBalanceWei is a paid mutator transaction binding the contract method 0xe2c83445.
//
// Solidity: function withdrawNonBeaconChainETHBalanceWei(address recipient, uint256 amountToWithdraw) returns()
func (_Restakingpod *RestakingpodSession) WithdrawNonBeaconChainETHBalanceWei(recipient common.Address, amountToWithdraw *big.Int) (*types.Transaction, error) {
	return _Restakingpod.Contract.WithdrawNonBeaconChainETHBalanceWei(&_Restakingpod.TransactOpts, recipient, amountToWithdraw)
}

// WithdrawNonBeaconChainETHBalanceWei is a paid mutator transaction binding the contract method 0xe2c83445.
//
// Solidity: function withdrawNonBeaconChainETHBalanceWei(address recipient, uint256 amountToWithdraw) returns()
func (_Restakingpod *RestakingpodTransactorSession) WithdrawNonBeaconChainETHBalanceWei(recipient common.Address, amountToWithdraw *big.Int) (*types.Transaction, error) {
	return _Restakingpod.Contract.WithdrawNonBeaconChainETHBalanceWei(&_Restakingpod.TransactOpts, recipient, amountToWithdraw)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Restakingpod *RestakingpodTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Restakingpod.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Restakingpod *RestakingpodSession) Receive() (*types.Transaction, error) {
	return _Restakingpod.Contract.Receive(&_Restakingpod.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Restakingpod *RestakingpodTransactorSession) Receive() (*types.Transaction, error) {
	return _Restakingpod.Contract.Receive(&_Restakingpod.TransactOpts)
}

// RestakingpodAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Restakingpod contract.
type RestakingpodAdminChangedIterator struct {
	Event *RestakingpodAdminChanged // Event containing the contract specifics and raw log

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
func (it *RestakingpodAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RestakingpodAdminChanged)
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
		it.Event = new(RestakingpodAdminChanged)
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
func (it *RestakingpodAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RestakingpodAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RestakingpodAdminChanged represents a AdminChanged event raised by the Restakingpod contract.
type RestakingpodAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Restakingpod *RestakingpodFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*RestakingpodAdminChangedIterator, error) {

	logs, sub, err := _Restakingpod.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &RestakingpodAdminChangedIterator{contract: _Restakingpod.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Restakingpod *RestakingpodFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *RestakingpodAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Restakingpod.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RestakingpodAdminChanged)
				if err := _Restakingpod.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_Restakingpod *RestakingpodFilterer) ParseAdminChanged(log types.Log) (*RestakingpodAdminChanged, error) {
	event := new(RestakingpodAdminChanged)
	if err := _Restakingpod.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RestakingpodBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the Restakingpod contract.
type RestakingpodBeaconUpgradedIterator struct {
	Event *RestakingpodBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *RestakingpodBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RestakingpodBeaconUpgraded)
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
		it.Event = new(RestakingpodBeaconUpgraded)
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
func (it *RestakingpodBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RestakingpodBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RestakingpodBeaconUpgraded represents a BeaconUpgraded event raised by the Restakingpod contract.
type RestakingpodBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Restakingpod *RestakingpodFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*RestakingpodBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Restakingpod.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &RestakingpodBeaconUpgradedIterator{contract: _Restakingpod.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Restakingpod *RestakingpodFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *RestakingpodBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Restakingpod.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RestakingpodBeaconUpgraded)
				if err := _Restakingpod.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_Restakingpod *RestakingpodFilterer) ParseBeaconUpgraded(log types.Log) (*RestakingpodBeaconUpgraded, error) {
	event := new(RestakingpodBeaconUpgraded)
	if err := _Restakingpod.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RestakingpodDaoChangedIterator is returned from FilterDaoChanged and is used to iterate over the raw logs and unpacked data for DaoChanged events raised by the Restakingpod contract.
type RestakingpodDaoChangedIterator struct {
	Event *RestakingpodDaoChanged // Event containing the contract specifics and raw log

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
func (it *RestakingpodDaoChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RestakingpodDaoChanged)
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
		it.Event = new(RestakingpodDaoChanged)
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
func (it *RestakingpodDaoChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RestakingpodDaoChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RestakingpodDaoChanged represents a DaoChanged event raised by the Restakingpod contract.
type RestakingpodDaoChanged struct {
	OldDao common.Address
	Dao    common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDaoChanged is a free log retrieval operation binding the contract event 0xfcde6c827a52b0870bc44ed9b10212272e18c9ea1725b772e9b493750afd8da4.
//
// Solidity: event DaoChanged(address _oldDao, address _dao)
func (_Restakingpod *RestakingpodFilterer) FilterDaoChanged(opts *bind.FilterOpts) (*RestakingpodDaoChangedIterator, error) {

	logs, sub, err := _Restakingpod.contract.FilterLogs(opts, "DaoChanged")
	if err != nil {
		return nil, err
	}
	return &RestakingpodDaoChangedIterator{contract: _Restakingpod.contract, event: "DaoChanged", logs: logs, sub: sub}, nil
}

// WatchDaoChanged is a free log subscription operation binding the contract event 0xfcde6c827a52b0870bc44ed9b10212272e18c9ea1725b772e9b493750afd8da4.
//
// Solidity: event DaoChanged(address _oldDao, address _dao)
func (_Restakingpod *RestakingpodFilterer) WatchDaoChanged(opts *bind.WatchOpts, sink chan<- *RestakingpodDaoChanged) (event.Subscription, error) {

	logs, sub, err := _Restakingpod.contract.WatchLogs(opts, "DaoChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RestakingpodDaoChanged)
				if err := _Restakingpod.contract.UnpackLog(event, "DaoChanged", log); err != nil {
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
func (_Restakingpod *RestakingpodFilterer) ParseDaoChanged(log types.Log) (*RestakingpodDaoChanged, error) {
	event := new(RestakingpodDaoChanged)
	if err := _Restakingpod.contract.UnpackLog(event, "DaoChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RestakingpodEigenLayerOperatorDelegatedIterator is returned from FilterEigenLayerOperatorDelegated and is used to iterate over the raw logs and unpacked data for EigenLayerOperatorDelegated events raised by the Restakingpod contract.
type RestakingpodEigenLayerOperatorDelegatedIterator struct {
	Event *RestakingpodEigenLayerOperatorDelegated // Event containing the contract specifics and raw log

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
func (it *RestakingpodEigenLayerOperatorDelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RestakingpodEigenLayerOperatorDelegated)
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
		it.Event = new(RestakingpodEigenLayerOperatorDelegated)
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
func (it *RestakingpodEigenLayerOperatorDelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RestakingpodEigenLayerOperatorDelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RestakingpodEigenLayerOperatorDelegated represents a EigenLayerOperatorDelegated event raised by the Restakingpod contract.
type RestakingpodEigenLayerOperatorDelegated struct {
	DelegateAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterEigenLayerOperatorDelegated is a free log retrieval operation binding the contract event 0xc689767f2ea597f80695a2490ad872245da94c1ce1b7af9993be5f2ea2282fde.
//
// Solidity: event EigenLayerOperatorDelegated(address _delegateAddress)
func (_Restakingpod *RestakingpodFilterer) FilterEigenLayerOperatorDelegated(opts *bind.FilterOpts) (*RestakingpodEigenLayerOperatorDelegatedIterator, error) {

	logs, sub, err := _Restakingpod.contract.FilterLogs(opts, "EigenLayerOperatorDelegated")
	if err != nil {
		return nil, err
	}
	return &RestakingpodEigenLayerOperatorDelegatedIterator{contract: _Restakingpod.contract, event: "EigenLayerOperatorDelegated", logs: logs, sub: sub}, nil
}

// WatchEigenLayerOperatorDelegated is a free log subscription operation binding the contract event 0xc689767f2ea597f80695a2490ad872245da94c1ce1b7af9993be5f2ea2282fde.
//
// Solidity: event EigenLayerOperatorDelegated(address _delegateAddress)
func (_Restakingpod *RestakingpodFilterer) WatchEigenLayerOperatorDelegated(opts *bind.WatchOpts, sink chan<- *RestakingpodEigenLayerOperatorDelegated) (event.Subscription, error) {

	logs, sub, err := _Restakingpod.contract.WatchLogs(opts, "EigenLayerOperatorDelegated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RestakingpodEigenLayerOperatorDelegated)
				if err := _Restakingpod.contract.UnpackLog(event, "EigenLayerOperatorDelegated", log); err != nil {
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

// ParseEigenLayerOperatorDelegated is a log parse operation binding the contract event 0xc689767f2ea597f80695a2490ad872245da94c1ce1b7af9993be5f2ea2282fde.
//
// Solidity: event EigenLayerOperatorDelegated(address _delegateAddress)
func (_Restakingpod *RestakingpodFilterer) ParseEigenLayerOperatorDelegated(log types.Log) (*RestakingpodEigenLayerOperatorDelegated, error) {
	event := new(RestakingpodEigenLayerOperatorDelegated)
	if err := _Restakingpod.contract.UnpackLog(event, "EigenLayerOperatorDelegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RestakingpodEigenLayerOperatorUndelegatedIterator is returned from FilterEigenLayerOperatorUndelegated and is used to iterate over the raw logs and unpacked data for EigenLayerOperatorUndelegated events raised by the Restakingpod contract.
type RestakingpodEigenLayerOperatorUndelegatedIterator struct {
	Event *RestakingpodEigenLayerOperatorUndelegated // Event containing the contract specifics and raw log

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
func (it *RestakingpodEigenLayerOperatorUndelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RestakingpodEigenLayerOperatorUndelegated)
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
		it.Event = new(RestakingpodEigenLayerOperatorUndelegated)
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
func (it *RestakingpodEigenLayerOperatorUndelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RestakingpodEigenLayerOperatorUndelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RestakingpodEigenLayerOperatorUndelegated represents a EigenLayerOperatorUndelegated event raised by the Restakingpod contract.
type RestakingpodEigenLayerOperatorUndelegated struct {
	DelegateAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterEigenLayerOperatorUndelegated is a free log retrieval operation binding the contract event 0x4f26776a6a3608028cc6d042406ca04f7b35f736d614b41a4e6b18eac2ba7321.
//
// Solidity: event EigenLayerOperatorUndelegated(address _delegateAddress)
func (_Restakingpod *RestakingpodFilterer) FilterEigenLayerOperatorUndelegated(opts *bind.FilterOpts) (*RestakingpodEigenLayerOperatorUndelegatedIterator, error) {

	logs, sub, err := _Restakingpod.contract.FilterLogs(opts, "EigenLayerOperatorUndelegated")
	if err != nil {
		return nil, err
	}
	return &RestakingpodEigenLayerOperatorUndelegatedIterator{contract: _Restakingpod.contract, event: "EigenLayerOperatorUndelegated", logs: logs, sub: sub}, nil
}

// WatchEigenLayerOperatorUndelegated is a free log subscription operation binding the contract event 0x4f26776a6a3608028cc6d042406ca04f7b35f736d614b41a4e6b18eac2ba7321.
//
// Solidity: event EigenLayerOperatorUndelegated(address _delegateAddress)
func (_Restakingpod *RestakingpodFilterer) WatchEigenLayerOperatorUndelegated(opts *bind.WatchOpts, sink chan<- *RestakingpodEigenLayerOperatorUndelegated) (event.Subscription, error) {

	logs, sub, err := _Restakingpod.contract.WatchLogs(opts, "EigenLayerOperatorUndelegated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RestakingpodEigenLayerOperatorUndelegated)
				if err := _Restakingpod.contract.UnpackLog(event, "EigenLayerOperatorUndelegated", log); err != nil {
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

// ParseEigenLayerOperatorUndelegated is a log parse operation binding the contract event 0x4f26776a6a3608028cc6d042406ca04f7b35f736d614b41a4e6b18eac2ba7321.
//
// Solidity: event EigenLayerOperatorUndelegated(address _delegateAddress)
func (_Restakingpod *RestakingpodFilterer) ParseEigenLayerOperatorUndelegated(log types.Log) (*RestakingpodEigenLayerOperatorUndelegated, error) {
	event := new(RestakingpodEigenLayerOperatorUndelegated)
	if err := _Restakingpod.contract.UnpackLog(event, "EigenLayerOperatorUndelegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RestakingpodInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Restakingpod contract.
type RestakingpodInitializedIterator struct {
	Event *RestakingpodInitialized // Event containing the contract specifics and raw log

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
func (it *RestakingpodInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RestakingpodInitialized)
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
		it.Event = new(RestakingpodInitialized)
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
func (it *RestakingpodInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RestakingpodInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RestakingpodInitialized represents a Initialized event raised by the Restakingpod contract.
type RestakingpodInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Restakingpod *RestakingpodFilterer) FilterInitialized(opts *bind.FilterOpts) (*RestakingpodInitializedIterator, error) {

	logs, sub, err := _Restakingpod.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &RestakingpodInitializedIterator{contract: _Restakingpod.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Restakingpod *RestakingpodFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *RestakingpodInitialized) (event.Subscription, error) {

	logs, sub, err := _Restakingpod.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RestakingpodInitialized)
				if err := _Restakingpod.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Restakingpod *RestakingpodFilterer) ParseInitialized(log types.Log) (*RestakingpodInitialized, error) {
	event := new(RestakingpodInitialized)
	if err := _Restakingpod.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RestakingpodOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Restakingpod contract.
type RestakingpodOwnershipTransferredIterator struct {
	Event *RestakingpodOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RestakingpodOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RestakingpodOwnershipTransferred)
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
		it.Event = new(RestakingpodOwnershipTransferred)
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
func (it *RestakingpodOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RestakingpodOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RestakingpodOwnershipTransferred represents a OwnershipTransferred event raised by the Restakingpod contract.
type RestakingpodOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Restakingpod *RestakingpodFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RestakingpodOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Restakingpod.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RestakingpodOwnershipTransferredIterator{contract: _Restakingpod.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Restakingpod *RestakingpodFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RestakingpodOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Restakingpod.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RestakingpodOwnershipTransferred)
				if err := _Restakingpod.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Restakingpod *RestakingpodFilterer) ParseOwnershipTransferred(log types.Log) (*RestakingpodOwnershipTransferred, error) {
	event := new(RestakingpodOwnershipTransferred)
	if err := _Restakingpod.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RestakingpodPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Restakingpod contract.
type RestakingpodPausedIterator struct {
	Event *RestakingpodPaused // Event containing the contract specifics and raw log

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
func (it *RestakingpodPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RestakingpodPaused)
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
		it.Event = new(RestakingpodPaused)
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
func (it *RestakingpodPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RestakingpodPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RestakingpodPaused represents a Paused event raised by the Restakingpod contract.
type RestakingpodPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Restakingpod *RestakingpodFilterer) FilterPaused(opts *bind.FilterOpts) (*RestakingpodPausedIterator, error) {

	logs, sub, err := _Restakingpod.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &RestakingpodPausedIterator{contract: _Restakingpod.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Restakingpod *RestakingpodFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *RestakingpodPaused) (event.Subscription, error) {

	logs, sub, err := _Restakingpod.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RestakingpodPaused)
				if err := _Restakingpod.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Restakingpod *RestakingpodFilterer) ParsePaused(log types.Log) (*RestakingpodPaused, error) {
	event := new(RestakingpodPaused)
	if err := _Restakingpod.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RestakingpodReceivedIterator is returned from FilterReceived and is used to iterate over the raw logs and unpacked data for Received events raised by the Restakingpod contract.
type RestakingpodReceivedIterator struct {
	Event *RestakingpodReceived // Event containing the contract specifics and raw log

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
func (it *RestakingpodReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RestakingpodReceived)
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
		it.Event = new(RestakingpodReceived)
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
func (it *RestakingpodReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RestakingpodReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RestakingpodReceived represents a Received event raised by the Restakingpod contract.
type RestakingpodReceived struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReceived is a free log retrieval operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address _sender, uint256 _amount)
func (_Restakingpod *RestakingpodFilterer) FilterReceived(opts *bind.FilterOpts) (*RestakingpodReceivedIterator, error) {

	logs, sub, err := _Restakingpod.contract.FilterLogs(opts, "Received")
	if err != nil {
		return nil, err
	}
	return &RestakingpodReceivedIterator{contract: _Restakingpod.contract, event: "Received", logs: logs, sub: sub}, nil
}

// WatchReceived is a free log subscription operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address _sender, uint256 _amount)
func (_Restakingpod *RestakingpodFilterer) WatchReceived(opts *bind.WatchOpts, sink chan<- *RestakingpodReceived) (event.Subscription, error) {

	logs, sub, err := _Restakingpod.contract.WatchLogs(opts, "Received")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RestakingpodReceived)
				if err := _Restakingpod.contract.UnpackLog(event, "Received", log); err != nil {
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
func (_Restakingpod *RestakingpodFilterer) ParseReceived(log types.Log) (*RestakingpodReceived, error) {
	event := new(RestakingpodReceived)
	if err := _Restakingpod.contract.UnpackLog(event, "Received", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RestakingpodRestakingPodManagerChangedIterator is returned from FilterRestakingPodManagerChanged and is used to iterate over the raw logs and unpacked data for RestakingPodManagerChanged events raised by the Restakingpod contract.
type RestakingpodRestakingPodManagerChangedIterator struct {
	Event *RestakingpodRestakingPodManagerChanged // Event containing the contract specifics and raw log

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
func (it *RestakingpodRestakingPodManagerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RestakingpodRestakingPodManagerChanged)
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
		it.Event = new(RestakingpodRestakingPodManagerChanged)
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
func (it *RestakingpodRestakingPodManagerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RestakingpodRestakingPodManagerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RestakingpodRestakingPodManagerChanged represents a RestakingPodManagerChanged event raised by the Restakingpod contract.
type RestakingpodRestakingPodManagerChanged struct {
	OldPodManager common.Address
	PodManager    common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRestakingPodManagerChanged is a free log retrieval operation binding the contract event 0x5382acc51e7dbd7925f723745f575e071a00350da3e4a81672e6cd4e4711c224.
//
// Solidity: event RestakingPodManagerChanged(address _oldPodManager, address _podManager)
func (_Restakingpod *RestakingpodFilterer) FilterRestakingPodManagerChanged(opts *bind.FilterOpts) (*RestakingpodRestakingPodManagerChangedIterator, error) {

	logs, sub, err := _Restakingpod.contract.FilterLogs(opts, "RestakingPodManagerChanged")
	if err != nil {
		return nil, err
	}
	return &RestakingpodRestakingPodManagerChangedIterator{contract: _Restakingpod.contract, event: "RestakingPodManagerChanged", logs: logs, sub: sub}, nil
}

// WatchRestakingPodManagerChanged is a free log subscription operation binding the contract event 0x5382acc51e7dbd7925f723745f575e071a00350da3e4a81672e6cd4e4711c224.
//
// Solidity: event RestakingPodManagerChanged(address _oldPodManager, address _podManager)
func (_Restakingpod *RestakingpodFilterer) WatchRestakingPodManagerChanged(opts *bind.WatchOpts, sink chan<- *RestakingpodRestakingPodManagerChanged) (event.Subscription, error) {

	logs, sub, err := _Restakingpod.contract.WatchLogs(opts, "RestakingPodManagerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RestakingpodRestakingPodManagerChanged)
				if err := _Restakingpod.contract.UnpackLog(event, "RestakingPodManagerChanged", log); err != nil {
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

// ParseRestakingPodManagerChanged is a log parse operation binding the contract event 0x5382acc51e7dbd7925f723745f575e071a00350da3e4a81672e6cd4e4711c224.
//
// Solidity: event RestakingPodManagerChanged(address _oldPodManager, address _podManager)
func (_Restakingpod *RestakingpodFilterer) ParseRestakingPodManagerChanged(log types.Log) (*RestakingpodRestakingPodManagerChanged, error) {
	event := new(RestakingpodRestakingPodManagerChanged)
	if err := _Restakingpod.contract.UnpackLog(event, "RestakingPodManagerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RestakingpodStakedButNotVerifiedEthChangedIterator is returned from FilterStakedButNotVerifiedEthChanged and is used to iterate over the raw logs and unpacked data for StakedButNotVerifiedEthChanged events raised by the Restakingpod contract.
type RestakingpodStakedButNotVerifiedEthChangedIterator struct {
	Event *RestakingpodStakedButNotVerifiedEthChanged // Event containing the contract specifics and raw log

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
func (it *RestakingpodStakedButNotVerifiedEthChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RestakingpodStakedButNotVerifiedEthChanged)
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
		it.Event = new(RestakingpodStakedButNotVerifiedEthChanged)
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
func (it *RestakingpodStakedButNotVerifiedEthChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RestakingpodStakedButNotVerifiedEthChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RestakingpodStakedButNotVerifiedEthChanged represents a StakedButNotVerifiedEthChanged event raised by the Restakingpod contract.
type RestakingpodStakedButNotVerifiedEthChanged struct {
	OldAmount *big.Int
	NewAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStakedButNotVerifiedEthChanged is a free log retrieval operation binding the contract event 0x321780a20d81d42088113788f4254556c9c41214460d0fe6b156b1e7b4bf78c5.
//
// Solidity: event StakedButNotVerifiedEthChanged(uint256 _oldAmount, uint256 _newAmount)
func (_Restakingpod *RestakingpodFilterer) FilterStakedButNotVerifiedEthChanged(opts *bind.FilterOpts) (*RestakingpodStakedButNotVerifiedEthChangedIterator, error) {

	logs, sub, err := _Restakingpod.contract.FilterLogs(opts, "StakedButNotVerifiedEthChanged")
	if err != nil {
		return nil, err
	}
	return &RestakingpodStakedButNotVerifiedEthChangedIterator{contract: _Restakingpod.contract, event: "StakedButNotVerifiedEthChanged", logs: logs, sub: sub}, nil
}

// WatchStakedButNotVerifiedEthChanged is a free log subscription operation binding the contract event 0x321780a20d81d42088113788f4254556c9c41214460d0fe6b156b1e7b4bf78c5.
//
// Solidity: event StakedButNotVerifiedEthChanged(uint256 _oldAmount, uint256 _newAmount)
func (_Restakingpod *RestakingpodFilterer) WatchStakedButNotVerifiedEthChanged(opts *bind.WatchOpts, sink chan<- *RestakingpodStakedButNotVerifiedEthChanged) (event.Subscription, error) {

	logs, sub, err := _Restakingpod.contract.WatchLogs(opts, "StakedButNotVerifiedEthChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RestakingpodStakedButNotVerifiedEthChanged)
				if err := _Restakingpod.contract.UnpackLog(event, "StakedButNotVerifiedEthChanged", log); err != nil {
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

// ParseStakedButNotVerifiedEthChanged is a log parse operation binding the contract event 0x321780a20d81d42088113788f4254556c9c41214460d0fe6b156b1e7b4bf78c5.
//
// Solidity: event StakedButNotVerifiedEthChanged(uint256 _oldAmount, uint256 _newAmount)
func (_Restakingpod *RestakingpodFilterer) ParseStakedButNotVerifiedEthChanged(log types.Log) (*RestakingpodStakedButNotVerifiedEthChanged, error) {
	event := new(RestakingpodStakedButNotVerifiedEthChanged)
	if err := _Restakingpod.contract.UnpackLog(event, "StakedButNotVerifiedEthChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RestakingpodUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Restakingpod contract.
type RestakingpodUnpausedIterator struct {
	Event *RestakingpodUnpaused // Event containing the contract specifics and raw log

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
func (it *RestakingpodUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RestakingpodUnpaused)
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
		it.Event = new(RestakingpodUnpaused)
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
func (it *RestakingpodUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RestakingpodUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RestakingpodUnpaused represents a Unpaused event raised by the Restakingpod contract.
type RestakingpodUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Restakingpod *RestakingpodFilterer) FilterUnpaused(opts *bind.FilterOpts) (*RestakingpodUnpausedIterator, error) {

	logs, sub, err := _Restakingpod.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &RestakingpodUnpausedIterator{contract: _Restakingpod.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Restakingpod *RestakingpodFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *RestakingpodUnpaused) (event.Subscription, error) {

	logs, sub, err := _Restakingpod.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RestakingpodUnpaused)
				if err := _Restakingpod.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_Restakingpod *RestakingpodFilterer) ParseUnpaused(log types.Log) (*RestakingpodUnpaused, error) {
	event := new(RestakingpodUnpaused)
	if err := _Restakingpod.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RestakingpodUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Restakingpod contract.
type RestakingpodUpgradedIterator struct {
	Event *RestakingpodUpgraded // Event containing the contract specifics and raw log

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
func (it *RestakingpodUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RestakingpodUpgraded)
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
		it.Event = new(RestakingpodUpgraded)
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
func (it *RestakingpodUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RestakingpodUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RestakingpodUpgraded represents a Upgraded event raised by the Restakingpod contract.
type RestakingpodUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Restakingpod *RestakingpodFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*RestakingpodUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Restakingpod.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &RestakingpodUpgradedIterator{contract: _Restakingpod.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Restakingpod *RestakingpodFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *RestakingpodUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Restakingpod.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RestakingpodUpgraded)
				if err := _Restakingpod.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_Restakingpod *RestakingpodFilterer) ParseUpgraded(log types.Log) (*RestakingpodUpgraded, error) {
	event := new(RestakingpodUpgraded)
	if err := _Restakingpod.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
