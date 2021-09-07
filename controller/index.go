package controller

/*
 * @Description:
 * @Author: longfei
 * @FilePath: \go\controller\index.go
 */

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"goblog/config"
	"goblog/dao"
	d "goblog/model"
	"goblog/server"
	"goblog/util"

	"github.com/gin-gonic/gin"
)

//这里是详情页
func GetView(c *gin.Context) {
	vvv := util.GetView(c.Param("id"), 1) //获取文章详情
	if vvv.ID == 0 {
		Not404(c)
		c.Abort()
		return
	}
	tm := time.Unix(int64(vvv.CreatedAt), 0)
	vvv.Ctime = tm.Format("2006-01-02 15:04:05")

	baseinfo, err := server.Getinfo()
	if err != nil {
		server.Fail(c)
		return
	}
	//fmt.Printf("wyth%+v", baseinfo)
	c.HTML(http.StatusOK, "view.html", gin.H{
		"view": vvv,
		"body": template.HTML(vvv.Body),
		"base": baseinfo,
	})
}

//获取个人简介信息
func About(c *gin.Context) {
	vvv := util.GetView("0", 2) //个人简介
	if vvv.ID == 0 {
		Not404(c)
		c.Abort()
		return
	}
	tm := time.Unix(int64(vvv.CreatedAt), 0)
	vvv.Ctime = tm.Format("2006-01-02 15:04:05")

	baseinfo, err := server.Getinfo()
	if err != nil {
		server.Fail(c)
		return
	}
	//fmt.Printf("wyth%+v", baseinfo)
	c.HTML(http.StatusOK, "about.html", gin.H{
		"view": vvv,
		"body": template.HTML(vvv.Body),
		"base": baseinfo,
	})
}

//关联查询测试用
func Not404(c *gin.Context) {
	tp := GetTypeNew("0") //栏目分类
	c.HTML(http.StatusNotFound, "404.html", gin.H{
		"types": tp,
		"tdk":   config.GetTDK(),
	})
}

//文章搜索服务
/* func Search(c *gin.Context) {
	var searchinfo d.ListInfo
	err := c.ShouldBindQuery(&searchinfo)
	if err != nil {
		server.Fail(c)
		return
	}
	err = server.SearchView(&searchinfo)
	server.ResDataError(searchinfo, err, c)
} */

//新的获取列表功能,搜索页也在这
func NewList(c *gin.Context) {
	var searchinfo d.ListInfo
	err := c.ShouldBindQuery(&searchinfo)
	if err != nil {
		server.Fail(c)
		return
	}
	id1, _ := strconv.Atoi(c.Param("id"))
	if id1 != 0 {
		searchinfo.Page.ID = id1
	}
	fmt.Printf("info%+v %d", searchinfo, id1)
	err = server.GetListV(&searchinfo)
	if err != nil {
		server.Fail(c)
		return
	}
	//fmt.Printf("列表信息是%+v\n", searchinfo)
	baseinfo, err := server.Getinfo()
	if err != nil {
		server.Fail(c)
		return
	}
	//如果没有分类信息,那么是搜索
	if searchinfo.Page.ID == 0 {
		searchinfo.Listinfo.Name = searchinfo.Page.Keyword
		searchinfo.Listinfo.Info = "搜索:" + searchinfo.Page.Keyword + "的搜索结果"
	}
	//fmt.Printf("base基础信息的内容是 %+v \n", baseinfo)
	//server.ResDataError(searchinfo, err, c)
	c.HTML(http.StatusOK, "newlist.html", gin.H{
		"list":     searchinfo, //分类信息,文章列表,当前分类信息
		"baseinfo": baseinfo,   //最新,推荐,tdk,全部分类信息
	})
}

//这里是首页
func Index(c *gin.Context) {
	remen := server.GetViewlist("-3", 1)  //热门
	swiper := server.GetViewlist("-2", 1) //轮播

	baseinfo, err := server.Getinfo()
	if err != nil {
		server.Fail(c)
		return
	}

	//每个分类的循环获取
	tnew := []d.Tp{}
	for _, v := range baseinfo.Typeinfo {
		v.Views = server.Findlist2(strconv.Itoa(int(v.ID)))
		v.Views = util.Imgsrc(v.Views) //对图片进行批量替换，如果无图则设置默认
		tnew = append(tnew, v)
	}
	//fmt.Printf("转换后的数据%+v", tnew)
	//友情链接
	var link []d.Link
	dao.MDB.Find(&link)
	//fmt.Printf("link%+v", link)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"base":   baseinfo, //基础信息,可替换上面四条
		"tt":     tnew,
		"remen":  util.Imgsrc(remen),
		"swiper": util.Imgsrc(swiper),
		"link":   link,
	})
}

func Test(c *gin.Context) {
	util.DelAll()
}
