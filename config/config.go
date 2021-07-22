package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var Configv *viper.Viper

func init() {
	Configv = viper.New()
	Configv.SetConfigType("json")
	Configv.SetConfigFile("config/set.json")
	Configv.WatchConfig()
	err := Configv.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
}
