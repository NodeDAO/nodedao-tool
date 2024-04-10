package conf

import (
	"fmt"
	logging "github.com/ipfs/go-log/v2"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

var log = logging.Logger("config")

type Config struct {
	Eth1Rpc string `json:"eth1rpc"`
	Eth2Rpc string `json:"eth2rpc"`
}

var conf Config

func GetConfig() Config {
	return conf
}

func Init(p string) {
	v := viper.New()
	v.SetConfigName("conf")
	v.SetConfigType("yaml")
	if p != "" {
		v.AddConfigPath(filepath.Dir(p))
	} else {
		v.AddConfigPath("./conf")
		v.AddConfigPath("../config")
	}
	v.AddConfigPath(".")
	err := v.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = v.Unmarshal(&conf)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
