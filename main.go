package main

import (
	"context"
	"fmt"
	"github.com/NodeDAO/nodedao-tool/conf"
	depositContract "github.com/NodeDAO/nodedao-tool/contract/depositcontract"
	"github.com/NodeDAO/nodedao-tool/contract/nodedaopool"
	"github.com/NodeDAO/nodedao-tool/depositData"
	"github.com/NodeDAO/nodedao-tool/eth1"
	"github.com/NodeDAO/nodedao-tool/eth2"
	"github.com/NodeDAO/nodedao-tool/totalEth"
	"github.com/ethereum/go-ethereum/common"
	logging "github.com/ipfs/go-log/v2"
	"github.com/spf13/cobra"
	"time"
)

var log = logging.Logger("main")
var (
	configPath          string
	depositDataJsonPath string
	outputDir           string
	spiteSize           int
)

var rootCmd = &cobra.Command{
	Use:   "nodedao-tool",
	Short: "nodedao-tool",
	Long:  `nodedao tool`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "./conf/config.yaml", "path to configuration file")
	registerValidatorCmd.PersistentFlags().StringVarP(&depositDataJsonPath, "depositDataJsonPath", "d", "", "deposit data file path")
	spiteDepositDataCmd.PersistentFlags().StringVarP(&depositDataJsonPath, "depositDataJsonPath", "d", "", "deposit data file path")
	spiteDepositDataCmd.PersistentFlags().StringVarP(&outputDir, "outputDir", "o", "", "output dir")
	spiteDepositDataCmd.PersistentFlags().IntVarP(&spiteSize, "spiteSize", "s", 10, "spite size")
}

func main() {
	_ = logging.SetLogLevel("*", "INFO")

	rootCmd.AddCommand(registerValidatorCmd, spiteDepositDataCmd, totalEThNethCmd, totalEThRNethCmd)

	_ = rootCmd.Execute()
}

var totalEThNethCmd = &cobra.Command{
	Use:     "totaleth-neth",
	Short:   "totaleth-neth",
	Example: "./nodedao-tool totaleth-neth",
	Run: func(cmd *cobra.Command, args []string) {
		conf.Init(configPath)
		err := eth2.InitConsensusClient(context.Background(), conf.GetConfig().Eth2Rpc, 1*time.Minute)
		if err != nil {
			log.Error(err)
			return
		}
		err = totalEth.CalcNethTotalEth()
		if err != nil {
			log.Error(err)
			return
		}
	},
}

var totalEThRNethCmd = &cobra.Command{
	Use:     "totaleth-rneth",
	Short:   "totaleth-rneth",
	Example: "./nodedao-tool totaleth-rneth",
	Run: func(cmd *cobra.Command, args []string) {
		conf.Init(configPath)
		err := eth2.InitConsensusClient(context.Background(), conf.GetConfig().Eth2Rpc, 1*time.Minute)
		if err != nil {
			log.Error(err)
			return
		}
		err = totalEth.CalcRnethTotalEth()
		if err != nil {
			log.Error(err)
			return
		}
	},
}

var spiteDepositDataCmd = &cobra.Command{
	Use:     "spite-depositdata",
	Short:   "spite-depositdata",
	Example: "./nodedao-tool spite-depositdata",
	Run: func(cmd *cobra.Command, args []string) {
		err := depositData.SpiteDepositData(depositDataJsonPath, outputDir, spiteSize)
		if err != nil {
			log.Error(err)
			return
		}
	},
}

var registerValidatorCmd = &cobra.Command{
	Use:     "register-validator",
	Short:   "register-validator",
	Example: "./nodedao-tool register-validator",
	Run: func(cmd *cobra.Command, args []string) {
		conf.Init(configPath)
		config := conf.GetConfig()
		client, cancel, err := eth1.GetEthClient(config.Eth1Rpc)
		if err != nil {
			log.Error(err)
			return
		}
		defer cancel()

		deposit, err := depositContract.NewDepositContract(conf.DepositContract, client)
		if err != nil {
			log.Error(err)
			return
		}
		root, err := deposit.GetDepositRoot(nil)
		if err != nil {
			log.Error(err)
			return
		}

		fmt.Printf("depositContractRoot: %s \n", "0x"+common.Bytes2Hex(root[:]))
		fmt.Println("")
		fmt.Printf("restakingPod: %s\n", conf.RestakingPod.String())
		fmt.Println("")
		pubkeyData, signatureData, depositDataRootData, err := depositData.GenerateRegisterValidator(depositDataJsonPath)
		if err != nil {
			log.Error(err)
			return
		}

		inputData, err := nodedaopool.NodeDaoPoolABI.Pack("registerValidator", root, conf.RestakingPod, pubkeyData, signatureData, depositDataRootData)
		if err != nil {
			log.Error(err)
			return
		}

		fmt.Println("")
		fmt.Println("inputData: ", "0x"+common.Bytes2Hex(inputData))
		fmt.Println("")
		fmt.Println(" !!! Check the output inputData and tx inputData are exactly the same. e.g. https://text-compare.com/")
	},
}
