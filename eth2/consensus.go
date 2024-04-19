// desc:
// @author renshiwei
// Date: 2023/4/6 20:02

package eth2

import (
	"context"
	"github.com/NodeDAO/nodedao-tool/eth2/beacon"
	"github.com/NodeDAO/nodedao-tool/eth2/chaintime/standard"
	eth2client "github.com/attestantio/go-eth2-client"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"time"
)

type Consensus struct {
	ConsensusClient        eth2client.Service
	ChainTimeService       *standard.Service
	CustomizeBeaconService *beacon.BeaconService

	Timeout time.Duration
}

var ConsensusClient *Consensus

func InitConsensusClient(ctx context.Context, addr string, timeout time.Duration) error {
	beaconNode, err := beacon.ConnectToBeaconNode(ctx, addr, timeout, true)
	if err != nil {
		return errors.Wrap(err, "")
	}

	beaconService, err := beacon.New(ctx, addr, timeout)
	if err != nil {
		return errors.Wrapf(err, "Failed to connect to beacon service")
	}

	chainTimeService, err := standard.New(ctx,
		standard.WithGenesisTimeProvider(beaconNode.(eth2client.GenesisTimeProvider)),
		standard.WithSpecProvider(beaconNode.(eth2client.SpecProvider)),
		standard.WithLogLevel(zerolog.Disabled),
	)
	if err != nil {
		return errors.Wrap(err, "failed to create chaintime service")
	}

	ConsensusClient = &Consensus{
		ConsensusClient:        beaconNode,
		ChainTimeService:       chainTimeService,
		CustomizeBeaconService: beaconService,
		Timeout:                timeout,
	}
	return nil
}
