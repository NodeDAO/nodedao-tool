package totalEth

import (
	"context"
	"github.com/NodeDAO/nodedao-tool/conf"
	"github.com/NodeDAO/nodedao-tool/contract/nodedaopool"
	"github.com/NodeDAO/nodedao-tool/eth1"
	"github.com/NodeDAO/nodedao-tool/eth2"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
	logging "github.com/ipfs/go-log/v2"
	"math/big"
)

var log = logging.Logger("totalETH")

const startBlock = uint64(19516980)

var ValidatorRegistrationTopic = common.HexToHash("0xe585eadb0042252d35431ecfed1027caf5672811fdf1db08cb49e32fee170505")

func scanValidators(addr common.Address) ([]string, error) {
	config := conf.GetConfig()
	eth1Client, cancel, err := eth1.GetEthClient(config.Eth1Rpc)
	if err != nil {
		return nil, err
	}
	defer cancel()

	curBlock, err := eth1Client.BlockNumber(context.Background())
	if err != nil {
		return nil, err
	}

	var validators []string = make([]string, 0)
	for fromBlock := startBlock; fromBlock < curBlock; {
		nextBlock := fromBlock + 20000
		if nextBlock >= curBlock {
			nextBlock = curBlock
		}

		filter := ethereum.FilterQuery{
			FromBlock: big.NewInt(int64(fromBlock)),
			ToBlock:   big.NewInt(int64(nextBlock)),
			Addresses: []common.Address{addr},
			Topics:    [][]common.Hash{{ValidatorRegistrationTopic}},
		}

		log.Infow("scanValidators", "fromBlock", fromBlock, "nextBlock", nextBlock)
		fromBlock = nextBlock

		logs, err := eth1Client.FilterLogs(context.Background(), filter)
		if err != nil {
			log.Error(err)
			continue
		}

		for _, l := range logs {
			pubkeys, err := nodedaopool.NodeDaoPoolABI.Events["ValidatorRegistration"].Inputs.Unpack(l.Data)
			if err != nil {
				log.Warnw("err", err)
				continue
			}

			for _, pubkey := range pubkeys[0].([][]uint8) {
				validator := common.Bytes2Hex(pubkey)
				validators = append(validators, "0x"+validator)
				log.Infow("scanValidators", "pubkey", validator)
			}
		}
	}

	return validators, nil
}

func calcEth2Balance(pubkeys []string, refSlot string) (*big.Int, error) {
	validators, err := eth2.ConsensusClient.CustomizeBeaconService.ValidatorsByPubKey(context.Background(), refSlot, pubkeys)
	if err != nil {
		return nil, err
	}

	totalBeaconBalance := big.NewInt(0)
	for _, validator := range validators {
		totalBeaconBalance = big.NewInt(0).Add(totalBeaconBalance, big.NewInt(int64(validator.Balance)))
	}

	if len(validators) != len(pubkeys) {
		// Newly registered validator
		number := int64(len(pubkeys) - len(validators))
		balance := big.NewInt(32000000000)
		totalBeaconBalance = big.NewInt(0).Add(totalBeaconBalance, big.NewInt(0).Mul(big.NewInt(number), balance))
	}

	return GWEIToWEI(totalBeaconBalance), nil
}

func GWEIToWEI(value *big.Int) *big.Int {
	return new(big.Int).Mul(value, big.NewInt(params.GWei))
}

func WEIToEth(value *big.Int) *big.Float {
	return new(big.Float).Quo(new(big.Float).Quo(big.NewFloat(0).SetInt(value), big.NewFloat(params.GWei)), big.NewFloat(params.GWei))
}
