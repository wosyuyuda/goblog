package config

import (
	"fmt"
	"goblog/model"

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

func GetTDK() (tdk model.Tdk) {
	tdk.Title = Configv.GetString("tdk.title")
	tdk.Keyword = Configv.GetString("tdk.keyword")
	tdk.Description = Configv.GetString("tdk.description")
	tdk.Beian = Configv.GetString("tdk.biean")
	tdk.Goan = Configv.GetString("tdk.goan")
	tdk.Goanurl = Configv.GetString("tdk.goanurl")
	tdk.Tongjiid = Configv.GetString("tdk.tonjiid")
	return
}
