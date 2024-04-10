package depositData

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"os"
	"strings"
)

type DepositData struct {
	Pubkey                string `json:"pubkey"`
	WithdrawalCredentials string `json:"withdrawal_credentials"`
	Amount                int64  `json:"amount"`
	Signature             string `json:"signature"`
	DepositMessageRoot    string `json:"deposit_message_root"`
	DepositDataRoot       string `json:"deposit_data_root"`
	ForkVersion           string `json:"fork_version"`
	NetworkName           string `json:"network_name"`
	DepositCliVersion     string `json:"deposit_cli_version"`
}

func GenerateRegisterValidator(depositDataPath string) ([][]byte, [][]byte, [][32]byte, error) {
	data, err := os.ReadFile(depositDataPath)
	if err != nil {
		return nil, nil, nil, err
	}

	var depositDatas []DepositData
	err = json.Unmarshal(data, &depositDatas)
	if err != nil {
		return nil, nil, nil, err
	}

	var pubkeys, signatures, depositDataRoots string
	var pubkeyData [][]byte
	var signatureData [][]byte
	var depositDataRootData [][32]byte

	for i, deposit := range depositDatas {
		pubkeyData = append(pubkeyData, common.Hex2Bytes(deposit.Pubkey))
		signatureData = append(signatureData, common.Hex2Bytes(deposit.Signature))
		rootData := [32]byte{}
		copy(rootData[:], common.Hex2Bytes(deposit.Pubkey))
		depositDataRootData = append(depositDataRootData, rootData)

		if i == 0 {
			pubkeys = "\"0x" + deposit.Pubkey + "\""
			signatures = "\"0x" + deposit.Signature + "\""
			depositDataRoots = "\"0x" + deposit.DepositDataRoot + "\""
			continue
		}
		pubkeys = fmt.Sprintf("%s,%s", pubkeys, "\"0x"+deposit.Pubkey+"\"")
		signatures = fmt.Sprintf("%s,%s", signatures, "\"0x"+deposit.Signature+"\"")
		depositDataRoots = fmt.Sprintf("%s,%s", depositDataRoots, "\"0x"+deposit.DepositDataRoot+"\"")
	}

	if len(strings.Split(pubkeys, ",")) != len(strings.Split(signatures, ",")) || len(strings.Split(pubkeys, ",")) != len(strings.Split(depositDataRoots, ",")) {
		fmt.Println("WARN: pubkeys / signatures / depositDataRoots mismatch")
		os.Exit(1)
	}

	fmt.Printf("pubkeys: [%s]\n", pubkeys)
	fmt.Println("")
	fmt.Printf("signatures: [%s]\n", signatures)
	fmt.Println("")
	fmt.Printf("depositDataRoots: [%s]\n", depositDataRoots)

	return pubkeyData, signatureData, depositDataRootData, nil
}
