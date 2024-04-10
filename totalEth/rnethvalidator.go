package totalEth

import "github.com/NodeDAO/nodedao-tool/conf"

func scanRNETHValidator() ([]string, error) {
	return scanValidators(conf.RNethPool)
}
