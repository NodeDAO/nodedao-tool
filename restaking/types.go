package restaking

import (
	"encoding/hex"
	"encoding/json"
	"os"
)

type StateRootProof struct {
	BeaconStateRoot string `json:"beaconStateRoot"`
	Proof           string `json:"proof"`
}

type WithdrawalProofsCallParams struct {
	WithdrawalProof        string `json:"withdrawalProof"`
	SlotProof              string `json:"slotProof"`
	ExecutionPayloadProof  string `json:"executionPayloadProof"`
	TimestampProof         string `json:"timestampProof"`
	HistoricalSummaryProof string `json:"historicalSummaryBlockRootProof"`
	BlockHeaderRootIndex   uint64 `json:"blockRootIndex"`
	HistoricalSummaryIndex uint64 `json:"historicalSummaryIndex"`
	WithdrawalIndex        uint64 `json:"withdrawalIndex"`
	BlockHeaderRoot        string `json:"blockRoot"`
	SlotRoot               string `json:"slotRoot"`
	TimestampRoot          string `json:"timestampRoot"`
	ExecutionPayloadRoot   string `json:"executionPayloadRoot"`
}

type VerifyAndProcessWithdrawalCallParams struct {
	StateRootProof        StateRootProof               `json:"stateRootProof"`
	WithdrawalProofs      []WithdrawalProofsCallParams `json:"withdrawalProofs"`
	ValidatorFieldsProofs []string                     `json:"validatorFieldsProofs"`
	ValidatorFields       [][]string                   `json:"validatorFields"`
	WithdrawalFields      [][]string                   `json:"WithdrawalFields"`
}

type VerifyWithdrawalCredentialsCallParams struct {
	StateRootProof        StateRootProof `json:"stateRootProof"`
	ValidatorIndices      []uint64       `json:"validatorIndices"`
	ValidatorFieldsProofs []string       `json:"validatorFieldsProofs"`
	ValidatorFields       [][]string     `json:"validatorFields"`
}

func ParseVerifyWithdrawalCredentialsCallParams(p string) (*VerifyWithdrawalCredentialsCallParams, error) {
	data, err := os.ReadFile(p)
	if err != nil {
		return nil, err
	}
	var cp VerifyWithdrawalCredentialsCallParams
	err = json.Unmarshal(data, &cp)
	if err != nil {
		return nil, err
	}

	return &cp, nil
}

func ParseVerifyAndProcessWithdrawalCallParams(p string) (*VerifyAndProcessWithdrawalCallParams, error) {
	data, err := os.ReadFile(p)
	if err != nil {
		return nil, err
	}
	var cp VerifyAndProcessWithdrawalCallParams
	err = json.Unmarshal(data, &cp)
	if err != nil {
		return nil, err
	}

	return &cp, nil
}

func hex2Bytes(s string) ([]byte, error) {
	if s[:2] == "0x" {
		return hex.DecodeString(s[2:])
	}

	return hex.DecodeString(s)
}

func hex2Byte32(s string) ([32]byte, error) {
	h, err := hex2Bytes(s)
	if err != nil {
		return [32]byte{}, err
	}
	var res = [32]byte{}
	copy(res[:], h)

	return res, nil
}
