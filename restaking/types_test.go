package restaking

import (
	"encoding/json"
	"github.com/NodeDAO/nodedao-tool/contract/restakingpod"
	"math/big"
	"testing"
)

func TestCompleteQueuedWithdrawals(t *testing.T) {
	withdrawals := `[{"staker":"0x9efdeb695f1e3f0ad3bcfb1e07d2749b50a3b75c","delegatedTo":"0xf626508e2f65614841cde2f3052cdc214eb13685","withdrawer":"0x9efdeb695f1e3f0ad3bcfb1e07d2749b50a3b75c","nonce":"0","startBlock":20280055,"strategies":["0xbeac0eeeeeeeeeeeeeeeeeeeeeeeeeeeeeebeac0"],"shares":["32000000000000000000"]}]`
	var datas = []Withdrawal{}
	err := json.Unmarshal([]byte(withdrawals), &datas)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(datas)

	var withdrawalCallParams []restakingpod.IDelegationManagerWithdrawal
	for _, data := range datas {
		nonce, ok := big.NewInt(0).SetString(data.Nonce, 10)
		if !ok {
			t.Fatal(err)
		}
		var shares []*big.Int
		for _, share := range data.Shares {
			amount, ok := big.NewInt(0).SetString(share, 10)
			if !ok {
				t.Fatal(err)
			}
			shares = append(shares, amount)
		}
		withdrawalCallParams = append(withdrawalCallParams, restakingpod.IDelegationManagerWithdrawal{
			Staker:      data.Staker,
			DelegatedTo: data.DelegatedTo,
			Withdrawer:  data.Withdrawer,
			Nonce:       nonce,
			StartBlock:  data.StartBlock,
			Strategies:  data.Strategies,
			Shares:      shares,
		})
	}

	t.Log(datas)
}
