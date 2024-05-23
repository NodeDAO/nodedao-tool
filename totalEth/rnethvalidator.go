package totalEth

import (
	"context"
	"fmt"
	"github.com/NodeDAO/nodedao-tool/conf"
	"github.com/NodeDAO/nodedao-tool/contract/elreward"
	"github.com/NodeDAO/nodedao-tool/contract/nodedaopool"
	"github.com/NodeDAO/nodedao-tool/contract/restakingpod"
	"github.com/NodeDAO/nodedao-tool/eth1"
	"github.com/NodeDAO/nodedao-tool/eth2"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"
)

func scanRNETHValidator() ([]string, error) {
	return scanValidators(conf.RNethPool)
}

func calcRnethEth1Balance(blockNumber int64) (*big.Int, *big.Int, error) {
	config := conf.GetConfig()
	eth1Client, cancel, err := eth1.GetEthClient(config.Eth1Rpc)
	if err != nil {
		return nil, nil, err
	}
	defer cancel()

	elVault, err := elreward.NewElreward(conf.RNEthElVault, eth1Client)
	if err != nil {
		return nil, nil, err
	}

	elReward, err := elVault.GetPoolRewards(&bind.CallOpts{
		BlockNumber: big.NewInt(blockNumber),
	})
	if err != nil {
		return nil, nil, err
	}

	log.Infow("rneth", "elReward", elReward.String())

	rnethPoolBalance, err := eth1Client.BalanceAt(context.Background(), conf.RNethPool, big.NewInt(blockNumber))
	if err != nil {
		return nil, nil, err
	}

	log.Infow("rneth", "rnethPoolBalance", rnethPoolBalance.String())

	rnethPoolContract, err := nodedaopool.NewPool(conf.RNethPool, eth1Client)
	if err != nil {
		return nil, nil, err
	}

	totalAssets, err := rnethPoolContract.TotalAssets(&bind.CallOpts{
		BlockNumber: big.NewInt(blockNumber),
	})
	if err != nil {
		return nil, nil, err
	}

	restakingPods, err := rnethPoolContract.GetRestakingPods(&bind.CallOpts{
		BlockNumber: big.NewInt(blockNumber),
	})
	if err != nil {
		return nil, nil, err
	}

	var clVaultBalance = big.NewInt(0)
	for _, pod := range restakingPods {
		podContract, err := restakingpod.NewRestakingpod(pod, eth1Client)
		if err != nil {
			return nil, nil, err
		}

		withdrawals, err := podContract.GetAllDelayedWithdrawals(&bind.CallOpts{
			BlockNumber: big.NewInt(blockNumber),
		})
		if err != nil {
			return nil, nil, err
		}
		log.Infow("rneth", "restakingpod", pod.String(), "withdrawals", withdrawals.String())
		clVaultBalance = big.NewInt(0).Add(clVaultBalance, withdrawals)

		eigenPod, err := podContract.EigenLayerEigenPod(nil)
		if err != nil {
			return nil, nil, err
		}

		eigenPodBalance, err := eth1Client.BalanceAt(context.Background(), eigenPod, big.NewInt(blockNumber))
		if err != nil {
			return nil, nil, err
		}
		log.Infow("rneth", "eigenPod", eigenPod.String(), "eigenPodBalance", eigenPodBalance.String())

		clVaultBalance = big.NewInt(0).Add(clVaultBalance, eigenPodBalance)

	}
	log.Infow("rneth", "clVaultBalance", clVaultBalance.String())
	return big.NewInt(0).Add(big.NewInt(0).Add(elReward, rnethPoolBalance), clVaultBalance), totalAssets, nil
}

func GetRnethValidatorIndex() ([]string, []uint64, error) {
	refSlot, err := eth2.ConsensusClient.CustomizeBeaconService.HeadSlot(context.Background())
	if err != nil {
		return nil, nil, err
	}

	pubkeys, err := scanRNETHValidator()
	if err != nil {
		return nil, nil, err
	}

	indexs, err := getValidatorIndex(pubkeys, refSlot.String())
	if err != nil {
		return nil, nil, err
	}

	return pubkeys, indexs, nil
}

func CalcRnethTotalEth() error {
	refSlot, err := eth2.ConsensusClient.CustomizeBeaconService.HeadSlot(context.Background())
	if err != nil {
		return err
	}
	executionBlock, err := eth2.ConsensusClient.CustomizeBeaconService.ExecutionBlock(context.Background(), refSlot.String())
	if err != nil {
		return err
	}

	pubkeys, err := scanRNETHValidator()
	if err != nil {
		return err
	}

	eth1Balance, totalAssets, err := calcRnethEth1Balance(executionBlock.BlockNumber.Int64())
	if err != nil {
		return err
	}
	log.Infow("rneth", "eth1Balance", eth1Balance.String())

	eth2Balance, err := calcEth2Balance(pubkeys, refSlot.String())
	if err != nil {
		return err
	}
	log.Infow("rneth", "eth2Balance", eth2Balance.String())

	totalETH := big.NewInt(0).Add(eth1Balance, eth2Balance)

	fmt.Println("===========================================================")
	fmt.Printf("totalETH : %f ETH (%s WEI) \n", WEIToEth(totalETH), totalETH.String())
	fmt.Printf("totalAsset : %f ETH (%s WEI) \n", WEIToEth(totalAssets), totalAssets.String())
	fmt.Printf("totalAsset - totalETH: %f ETH \n", WEIToEth(big.NewInt(0).Sub(totalAssets, totalETH)))
	fmt.Println("===========================================================")
	return nil
}
