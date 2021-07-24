package util

import (
	"encoding/json"
	"fmt"
	"goblog/model"
	"io/ioutil"
	"os"
	"path"
)

const (
	BaseCache = "baseinfo"       //首页基础缓存
	ListCache = "listcache%d-%d" //列表缓存,分类-页码
)

//简单的文件缓存,缓存首页与列表信息,先不整过期时间
func SetCache(key string, data interface{}) {
	studata, _ := json.Marshal(data)
	//fmt.Println(string(studata))
	fileName := "cache/" + key
	dstFile, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer dstFile.Close()
	dstFile.WriteString(string(studata))
}

//获取缓存,stru请转结构体,如果获取不到err!=nil
func GetCache(key string, stru interface{}) (err error) {
	f, err := ioutil.ReadFile("cache/" + key)
	if err != nil {
		return
	}
	err = json.Unmarshal(f, stru)
	return
}

//设置列表页的缓存
func SetListCache(views *model.ListInfo) {
	str := fmt.Sprintf(ListCache, views.Page.ID, views.Page.Page)
	SetCache(str, views)
}

//获取列表的缓存
func GetListCache(views *model.ListInfo) (err error) {
	str := fmt.Sprintf(ListCache, views.Page.ID, views.Page.Page)
	err = GetCache(str, views)
	return
}

//删除所有缓存
func DelAll() {
	dir, _ := ioutil.ReadDir("cache")
	for _, d := range dir {
		os.RemoveAll(path.Join([]string{"cache", d.Name()}...))
	}
}
