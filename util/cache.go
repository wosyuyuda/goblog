package util

import (
	"errors"
	"fmt"
	"goblog/model"
	"io/ioutil"
	"os"
	"path"
	"time"

	"github.com/patrickmn/go-cache"
)

const (
	BaseCache    = "baseinfo"       //首页基础缓存
	ListCache    = "listcache%d-%d" //列表缓存,分类-页码
	ViewCache    = "viewcache-%s"   //文章详情的缓存
	ViewCacheNum = "viewcache-%d"   //文章详情的缓存
)

var (
	MCA *cache.Cache
)

func init() {
	MCA = cache.New(180*time.Minute, 360*time.Minute)
}

//简单的文件缓存,缓存首页与列表信息,先不整过期时间
func SetCache(key string, data interface{}) {
	MCA.Set(key, data, cache.DefaultExpiration)
}

//获取缓存,stru请转结构体,如果获取不到err!=nil
func GetCache(key string) (stru interface{}, err error) {
	stru, found := MCA.Get(key)
	if !found {
		err = errors.New("不存在")
	}
	return
}

//设置列表页的缓存
func SetListCache(views *model.ListInfo) {
	str := fmt.Sprintf(ListCache, views.Page.ID, views.Page.Page)
	SetCache(str, views)
}

//获取列表的缓存
func GetListCache(page *model.PageList) (views *model.ListInfo, err error) {
	str := fmt.Sprintf(ListCache, views.Page.ID, views.Page.Page)
	fmt.Println("缓存文件是:", str)
	v1, found := MCA.Get(str)
	if !found {
		err = errors.New("未找到缓存")
		return
	}
	views = v1.(*model.ListInfo)
	//err = GetCache(str, views)
	return
}

//获取文章的缓存
func GetViewCache(id string) (v *model.View, err error) {
	str := fmt.Sprintf(ViewCache, id)
	v1, found := MCA.Get(str)
	if !found {
		err = errors.New("未找到缓存")
		return
	}
	v = v1.(*model.View)
	//err = GetCache(str, v)
	return
}

//设置文章详情的缓存
func SetViewCache(v *model.View) {
	str := fmt.Sprintf(ViewCacheNum, v.ID)
	if v.Tuijian == 3 {
		str = fmt.Sprintf(ViewCacheNum, 0)
	}
	SetCache(str, v)
}

//删除所有缓存
func DelAll() {
	dir, _ := ioutil.ReadDir("cache")
	for _, d := range dir {
		os.RemoveAll(path.Join([]string{"cache", d.Name()}...))
	}
}
