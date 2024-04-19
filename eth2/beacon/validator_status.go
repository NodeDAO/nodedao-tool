// description:
// @author renshiwei
// Date: 2023/6/30

package beacon

import (
	"context"
	consensusApi "github.com/attestantio/go-eth2-client/api/v1"
	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/ethereum/go-ethereum/params"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"math/big"
	"strconv"
)

// ValidatorIsFullExited Whether the validator has completed the lifecycle of the exit
func ValidatorIsFullExited(validator *consensusApi.Validator) bool {
	return validator.Status == consensusApi.ValidatorStateWithdrawalDone
}

// ValidatorIsRequestExiting Whether the validator request the exit
func ValidatorIsRequestExiting(validator *consensusApi.Validator) bool {
	return validator.Status == consensusApi.ValidatorStateActiveExiting ||
		validator.Status == consensusApi.ValidatorStateActiveSlashed ||
		validator.Status == consensusApi.ValidatorStateExitedUnslashed ||
		validator.Status == consensusApi.ValidatorStateExitedSlashed ||
		validator.Status == consensusApi.ValidatorStateWithdrawalPossible ||
		validator.Status == consensusApi.ValidatorStateWithdrawalDone
}

// ValidatorSlashedAmount 32ETh - ValidatorWithdrawAbleBalance
func (b *BeaconService) ValidatorSlashedAmount(ctx context.Context, validator *consensusApi.Validator) (*big.Int, error) {
	if !validator.Validator.Slashed {
		return nil, errors.New("validator not slashed.")
	}

	pubkey := validator.Validator.PublicKey.String()

	withdrawAbleSlot := epochToSlot(validator.Validator.WithdrawableEpoch)
	pubkeys := make([]string, 1)
	pubkeys[0] = pubkey
	slashedValidator, err := b.ValidatorsByPubKey(ctx, strconv.FormatInt(int64(withdrawAbleSlot), 10), pubkeys)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to get validators info.")
	}

	slashAmount := new(big.Int).Sub(ETH32().BigInt(), GweiToWei(slashedValidator[pubkey].Balance))

	return slashAmount, nil
}

func ETH32() decimal.Decimal {
	eth32, _ := decimal.NewFromString("32000000000000000000")
	return eth32
}

func GweiToWei(gwei phase0.Gwei) *big.Int {
	return new(big.Int).Mul(big.NewInt(int64(gwei)), big.NewInt(params.GWei))
}

func epochToSlot(epoch phase0.Epoch) phase0.Slot {
	return phase0.Slot(uint64(epoch) * 12)
}
