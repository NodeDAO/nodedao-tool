// desc:
// @author renshiwei
// Date: 2023/4/6 13:01

package beacon

import (
	"context"
	"github.com/NodeDAO/nodedao-tool/conf"
	eth2client "github.com/attestantio/go-eth2-client"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

var clAddr string

func initClAddr() {
	conf.Init("../../conf/config-dev.yaml")
	clAddr = conf.GetConfig().Eth2Rpc
}

func newBeaconClientTest() eth2client.Service {
	initClAddr()
	beaconNode, _ := ConnectToBeaconNode(context.Background(), clAddr, 2*time.Minute, true)
	return beaconNode
}

func TestConnectToBeaconNode(t *testing.T) {
	beaconNode := newBeaconClientTest()
	require.Equal(t, "Standard (HTTP)", beaconNode.Name())
}
