package server

/*
 * @Description:
 * @Author: longfei
 * @FilePath: \go\server\session.go
 */

import (
	"encoding/json"
	"fmt"

	//	"log"
	//	"net/http"
	"io/ioutil"
	"os"
)

type Contact struct {
	User string `json:"user"`
	Pwd  string `json:"pwd"`
	Host string `json:"host"`
	Db   string `json:"db"`
	Port int    `json:"port"`
}

func Get_contents(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(f)
}

func GetDbConfig() {
	var c Contact
	var content []byte
	var err error
	content, err = Get_contents("/config/config.json")
	if err != nil {
		fmt.Println("open file error: " + err.Error())
		return
	}
	err = json.Unmarshal([]byte(content), &c)
	if err != nil {
		fmt.Println("ERROR: ", err.Error())
		return
	}
	fmt.Println(c)
	fmt.Printf("获取到配置文件的数据是%+v", c)
}
