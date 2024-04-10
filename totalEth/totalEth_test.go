package totalEth

import (
	"github.com/NodeDAO/nodedao-tool/conf"
	"testing"
)

func TestScanValidator(t *testing.T) {
	conf.Init("../conf/conf.yaml")
	validators, err := scanValidators(rnethPool)
	if err != nil {
		t.Fatal(err)
	}

	for _, pubkey := range validators {
		t.Log(pubkey)
	}
}
