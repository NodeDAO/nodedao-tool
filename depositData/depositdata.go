package depositData

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"os"
	"path/filepath"
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

func SpiteDepositData(depositDataPath, outputDir string, validatorCount int) error {
	data, err := os.ReadFile(depositDataPath)
	if err != nil {
		return err
	}

	var depositDatas []DepositData
	err = json.Unmarshal(data, &depositDatas)
	if err != nil {
		return err
	}

	depositDatasLen := len(depositDatas)

	fileName := filepath.Base(depositDataPath)

	writeFunc := func(i int, d []byte) {
		fi, err := os.Create(filepath.Join(outputDir, fmt.Sprintf("%d-", i)+fileName))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		_, err = fi.Write(d)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fi.Close()
	}

	if len(depositDatas) <= validatorCount {
		writeFunc(1, data)
		return nil
	}

	count := len(depositDatas) / validatorCount
	if len(depositDatas)%validatorCount != 0 {
		count++
	}

	spiteLen := 0
	for i := 0; i < count; i++ {
		start := i * validatorCount
		end := start + validatorCount
		if i == count-1 {
			end = len(depositDatas)
		}
		tem := depositDatas[start:end]

		spiteLen += len(tem)

		d, err := json.Marshal(tem)
		if err != nil {
			return err
		}

		writeFunc(i+1, d)
	}

	if spiteLen != depositDatasLen {
		fmt.Println("Length mismatch after splitting")
		os.Exit(1)
	}

	return nil
}

func GenerateRegisterValidator(depositDataPath, format string) ([][]byte, [][]byte, [][32]byte, error) {
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
		copy(rootData[:], common.Hex2Bytes(deposit.DepositDataRoot))
		depositDataRootData = append(depositDataRootData, rootData)

		if format == "scan" {
			if i == 0 {
				pubkeys = "0x" + deposit.Pubkey
				signatures = "0x" + deposit.Signature
				depositDataRoots = "0x" + deposit.DepositDataRoot
				continue
			}
			pubkeys = fmt.Sprintf("%s,%s", pubkeys, "0x"+deposit.Pubkey)
			signatures = fmt.Sprintf("%s,%s", signatures, "0x"+deposit.Signature)
			depositDataRoots = fmt.Sprintf("%s,%s", depositDataRoots, "0x"+deposit.DepositDataRoot)
		} else {
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
