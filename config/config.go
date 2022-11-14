package config

import (
	"fmt"
	"goblog/model"
	"os"

	"github.com/spf13/viper"
)

var (
	Configv  *viper.Viper
	setpath  = "config/set.json"
	tempset  = "config/temp.json"
	IsMobile = 0 //是否支持移动端
)

func init() {
	Configv = viper.New()
	Configv.SetConfigType("json")
	//判断配置文件是否存在,如果不存在就用默认的(git提交时配置文件不迁移)
	exit, err := PathExists(setpath)
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
	if !exit {
		setpath = tempset
	}
	IsMobile = Configv.GetInt("ismobile") //看内置模板是否支持手机
	Configv.SetConfigFile(setpath)
	Configv.WatchConfig()
	err = Configv.ReadInConfig()
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

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
