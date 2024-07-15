package restaking

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"github.com/NodeDAO/nodedao-tool/conf"
	"github.com/NodeDAO/nodedao-tool/contract/restakingpod"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"math/big"
	"os"
)

func VerifyWithdrawalCredentials(gasPrice int64, gasLimit uint64, chainID int64, noSend bool, rpcHost string, oracleTimestamp uint64, vwp *VerifyWithdrawalCredentialsCallParams) error {
	key, err := getKeyFromENV()
	if err != nil {
		return err
	}

	eth1Client, err := getEthClient(rpcHost)
	if err != nil {
		return err
	}

	pod, err := restakingpod.NewRestakingpod(conf.RestakingPod, eth1Client)
	if err != nil {
		return err
	}
	proof, err := hex2Bytes(vwp.StateRootProof.Proof)
	if err != nil {
		return err
	}
	beaconStateRoot, err := hex2Byte32(vwp.StateRootProof.BeaconStateRoot)
	if err != nil {
		return err
	}
	var stateRootProof = restakingpod.BeaconChainProofsStateRootProof{
		BeaconStateRoot: beaconStateRoot,
		Proof:           proof,
	}

	var validatorIndices []*big.Int
	for _, wi := range vwp.ValidatorIndices {
		validatorIndices = append(validatorIndices, big.NewInt(int64(wi)))
	}

	var validatorFieldsProofs = [][]byte{}
	for _, vfp := range vwp.ValidatorFieldsProofs {
		vfProof, err := hex2Bytes(vfp)
		if err != nil {
			return err
		}
		validatorFieldsProofs = append(validatorFieldsProofs, vfProof)
	}

	var validatorFields = [][][32]byte{}
	for _, vfs := range vwp.ValidatorFields {
		var vField = [][32]byte{}
		for _, vf := range vfs {
			f, err := hex2Byte32(vf)
			if err != nil {
				return err
			}
			vField = append(vField, f)
		}
		validatorFields = append(validatorFields, vField)
	}

	addr := crypto.PubkeyToAddress(str2pri(key).PublicKey)
	fmt.Println("from: ", addr.String())
	opts := MakeTxOpts(addr, big.NewInt(gasPrice), gasLimit, chainID, key, noSend)

	tx, err := pod.VerifyWithdrawalCredentials(opts, oracleTimestamp, stateRootProof, validatorIndices, validatorFieldsProofs, validatorFields)
	if err != nil {
		return err
	}

	if noSend {
		fmt.Println("tx noSend")
	} else {
		fmt.Println("tx send success!")
	}

	fmt.Println("tx: ", tx.Hash().String())
	return nil
}

func VerifyAndProcessWithdrawal(gasPrice int64, gasLimit uint64, chainID int64, noSend bool, rpcHost string, oracleTimestamp uint64, vcp *VerifyAndProcessWithdrawalCallParams) error {
	key, err := getKeyFromENV()
	if err != nil {
		return err
	}

	eth1Client, err := getEthClient(rpcHost)
	if err != nil {
		return err
	}

	pod, err := restakingpod.NewRestakingpod(conf.RestakingPod, eth1Client)
	if err != nil {
		return err
	}

	proof, err := hex2Bytes(vcp.StateRootProof.Proof)
	if err != nil {
		return err
	}
	beaconStateRoot, err := hex2Byte32(vcp.StateRootProof.BeaconStateRoot)
	if err != nil {
		return err
	}
	var stateRootProof = restakingpod.BeaconChainProofsStateRootProof{
		BeaconStateRoot: beaconStateRoot,
		Proof:           proof,
	}

	var withdrawalProofs = []restakingpod.BeaconChainProofsWithdrawalProof{}
	for _, wp := range vcp.WithdrawalProofs {
		withdrawalProof, err := hex2Bytes(wp.WithdrawalProof)
		if err != nil {
			return err
		}
		slotProof, err := hex2Bytes(wp.SlotProof)
		if err != nil {
			return err
		}
		executionPayloadProof, err := hex2Bytes(wp.ExecutionPayloadProof)
		if err != nil {
			return err
		}
		timestampProof, err := hex2Bytes(wp.TimestampProof)
		if err != nil {
			return err
		}
		historicalSummaryProof, err := hex2Bytes(wp.HistoricalSummaryProof)
		if err != nil {
			return err
		}

		blockHeaderRoot, err := hex2Byte32(wp.BlockHeaderRoot)
		if err != nil {
			return err
		}
		slotRoot, err := hex2Byte32(wp.SlotRoot)
		if err != nil {
			return err
		}
		timestampRoot, err := hex2Byte32(wp.TimestampRoot)
		if err != nil {
			return err
		}
		executionPayloadRoot, err := hex2Byte32(wp.ExecutionPayloadRoot)
		if err != nil {
			return err
		}
		withdrawalProofs = append(withdrawalProofs, restakingpod.BeaconChainProofsWithdrawalProof{
			WithdrawalProof:                 withdrawalProof,
			SlotProof:                       slotProof,
			ExecutionPayloadProof:           executionPayloadProof,
			TimestampProof:                  timestampProof,
			HistoricalSummaryBlockRootProof: historicalSummaryProof,
			BlockRootIndex:                  wp.BlockHeaderRootIndex,
			HistoricalSummaryIndex:          wp.HistoricalSummaryIndex,
			WithdrawalIndex:                 wp.WithdrawalIndex,
			BlockRoot:                       blockHeaderRoot,
			SlotRoot:                        slotRoot,
			TimestampRoot:                   timestampRoot,
			ExecutionPayloadRoot:            executionPayloadRoot,
		})
	}

	var validatorFieldsProofs = [][]byte{}
	for _, vfp := range vcp.ValidatorFieldsProofs {
		vfProof, err := hex2Bytes(vfp)
		if err != nil {
			return err
		}
		validatorFieldsProofs = append(validatorFieldsProofs, vfProof)
	}

	var validatorFields = [][][32]byte{}
	for _, vfs := range vcp.ValidatorFields {
		var vField = [][32]byte{}
		for _, vf := range vfs {
			f, err := hex2Byte32(vf)
			if err != nil {
				return err
			}
			vField = append(vField, f)
		}
		validatorFields = append(validatorFields, vField)
	}

	var withdrawalFields [][][32]byte
	for _, wfs := range vcp.WithdrawalFields {
		var wField = [][32]byte{}
		for _, wf := range wfs {
			f, err := hex2Byte32(wf)
			if err != nil {
				return err
			}
			wField = append(wField, f)
		}
		withdrawalFields = append(withdrawalFields, wField)
	}

	addr := crypto.PubkeyToAddress(str2pri(key).PublicKey)
	fmt.Println("from: ", addr.String())
	opts := MakeTxOpts(addr, big.NewInt(gasPrice), gasLimit, chainID, key, noSend)
	tx, err := pod.VerifyAndProcessWithdrawals(opts, oracleTimestamp, stateRootProof, withdrawalProofs, validatorFieldsProofs, validatorFields, withdrawalFields)
	if err != nil {
		return err
	}
	if noSend {
		fmt.Println("tx noSend")
	} else {
		fmt.Println("tx send success!")
	}
	fmt.Println("tx: ", tx.Hash().String())

	return nil
}

func getEthClient(rpcHost string) (*ethclient.Client, error) {
	client, err := ethclient.Dial(rpcHost)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func getKeyFromENV() (string, error) {
	key, isExist := os.LookupEnv("PRIVATE_KEY")
	if !isExist {
		return "", errors.New("must set PRIVATE_KEY")
	}
	if key[:2] == "0x" {
		return key[2:], nil
	}

	return key, nil
}

func MakeTxOpts(from common.Address, gasPrice *big.Int, gasLimit uint64, chainID int64, key string, noSend bool) *bind.TransactOpts {
	txOpts := &bind.TransactOpts{
		From: from,
		Signer: func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			signedTx, err := SignTxFromPriKey(tx, chainID, key)
			if err != nil {
				return nil, err
			}
			fmt.Println("--inputdata--", hex.EncodeToString(tx.Data()))
			return signedTx, nil
		},
		GasPrice: gasPrice,
		GasLimit: gasLimit,
		Context:  context.Background(),
		NoSend:   noSend,
	}
	return txOpts
}

func SignTxFromPriKey(tx *types.Transaction, chainID int64, key string) (*types.Transaction, error) {
	signer := types.NewLondonSigner(big.NewInt(chainID))
	h := signer.Hash(tx)
	sign, err := crypto.Sign(h[:], str2pri(key))
	if err != nil {
		return nil, err
	}
	return tx.WithSignature(signer, sign)
}

func str2pri(pkStr string) *ecdsa.PrivateKey {
	privateKey, err := crypto.HexToECDSA(pkStr)
	if err != nil {
		return nil
	}
	return privateKey
}
