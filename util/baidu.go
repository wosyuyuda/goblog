package util

import (
	"fmt"
	"goblog/config"
	"io/ioutil"
	"net/http"
	"strings"
)

//百度链接自动提交,传文章ID
func BaiduLinksubmit(id uint) {
	config := config.Configv
	info := config.GetStringMapString("baidu")
	connArgs := fmt.Sprintf(info["domain"]+"view/%d", id)
	client := &http.Client{}
	rsp, err := client.Post(info["url"], "text/plain", strings.NewReader(connArgs))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rsp.Body.Close()
	bt, _ := ioutil.ReadAll(rsp.Body)
	fmt.Println(string(bt))
}
