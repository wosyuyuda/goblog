package server

/*
 * @Description:
 * @Author: longfei
 * @FilePath: \go\server\session.go
 */

import (
	"encoding/json"
	"fmt"
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

//获取数据配置信息,后面配置了redis后直接存进redis这样获取速度会更快一些.
func GetDbConfig() Contact {
	var c Contact
	var content []byte
	var err error

	content, err = Get_contents("config/db.json")
	if err != nil {
		fmt.Println("open file error: " + err.Error())
		return c
	}
	err = json.Unmarshal([]byte(content), &c)
	if err != nil {
		fmt.Println("ERROR: ", err.Error())
		return c
	}
	//fmt.Printf("获取到配置文件的数据是%+v", c)
	return c
}
