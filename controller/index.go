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

	"goblog/config"
	"goblog/dao"
	"goblog/model"
	d "goblog/model"
	"goblog/server"
	"goblog/util"

	"github.com/gin-gonic/gin"
)

//这里是详情页
func GetView(c *gin.Context) {
	vvv := util.GetView(c.Param("id"), 1) //获取文章详情
	fmt.Printf("111%+v", vvv)
	if vvv.ID == 0 {
		Not404(c)
		c.Abort()
		return
	}
	tp := GetTypeNew("0")                  //栏目分类
	newList := server.GetViewlist("0", 1)  //最新的列表
	tuijian := server.GetViewlist("-1", 1) //推荐的列表

	c.HTML(http.StatusOK, "view.html", gin.H{
		"view":    vvv,
		"body":    template.HTML(vvv.Body),
		"types":   tp,
		"newlist": util.Imgsrc(newList),
		"tuijian": util.Imgsrc(tuijian),
		"tdk":     config.GetTDK(),
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

//这里是列表页
func Lists(c *gin.Context) {

	id1 := c.Param("id")
	typeinfo := GetTypeNew(id1)

	list := server.GetViewlist(id1, 1) //获取列表数据
	//fmt.Printf("%+v", list)
	if len(list) == 0 || id1 == "0" {
		Not404(c)
		c.Abort()
		return
	}
	newList := server.GetViewlist("0", 1)  //最新
	tuijian := server.GetViewlist("-4", 1) //推荐
	tp := GetTypeNew("0")                  //栏目全部分类
	page := c.DefaultQuery("page", "1")    //获取当前分页
	pagenum, _ := strconv.Atoi(page)       //获取分页数据
	i := util.GetTypeCount(id1)            //获取当前分类有多少条
	p := util.GetPage(i, pagenum)          //获取分页数据
	//	fmt.Printf("最新数据是%+v", list)

	c.HTML(http.StatusOK, "list.html", gin.H{
		"list":     util.Imgsrc(list),
		"typeinfo": typeinfo[0], //如果没有一条数据是会报错的
		"types":    tp,
		"newlist":  util.Imgsrc(newList),
		"tuijian":  util.Imgsrc(tuijian),
		"pageinfo": p,
		"tdk":      config.GetTDK(),
	})
}

//这里是首页
func Index(c *gin.Context) {

	views := server.Findlist2("0")         //最新
	tuijian := server.GetViewlist("-4", 1) //推荐
	tp := GetTypeNew("0")                  //分类
	remen := server.GetViewlist("-3", 1)   //热门
	swiper := server.GetViewlist("-2", 1)  //轮播
	tp22 := GetTypeNew("0")                //获取分类
	tnew := []d.Tp{}
	for _, v := range tp22 {
		v.Views = server.Findlist2(strconv.Itoa(int(v.ID)))
		v.Views = util.Imgsrc(v.Views) //对图片进行批量替换，如果无图则设置默认
		tnew = append(tnew, v)
	}
	//fmt.Printf("转换后的数据%+v", tnew)
	var link []model.Link
	dao.MDB.Find(&link)
	//fmt.Printf("link%+v", link)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"list":    util.Imgsrc(views),
		"types":   tp,
		"tuijian": util.Imgsrc(tuijian),
		"tt":      tnew,
		"tdk":     config.GetTDK(),
		"remen":   util.Imgsrc(remen),
		"swiper":  util.Imgsrc(swiper),
		"link":    link,
	})
}
