// desc:
// @author renshiwei
// Date: 2023/4/6 20:29

package eth2

import (
	"context"
	"fmt"
	"github.com/NodeDAO/nodedao-tool/conf"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestGetConsensusInfo(t *testing.T) {
	conf.Init("../conf/config-goerli.dev.yaml")
	clAddr := conf.GetConfig().Eth2Rpc
	timeout := 1 * time.Minute

	background := context.Background()

	err := InitConsensusClient(background, clAddr, timeout)
	require.NoError(t, err)

	currentEpoch := ConsensusClient.ChainTimeService.CurrentEpoch()
	firstSlotOfCurrentEpoch := ConsensusClient.ChainTimeService.FirstSlotOfEpoch(currentEpoch)
	fmt.Printf("current epoch: %v\n", currentEpoch)
	fmt.Printf("first slot of current epoch: %v\n", firstSlotOfCurrentEpoch)
	require.Equal(t, currentEpoch, ConsensusClient.ChainTimeService.SlotToEpoch(firstSlotOfCurrentEpoch))
	require.Equal(t, 3852416, int(ConsensusClient.ChainTimeService.EpochToSlot(120388)))
}
