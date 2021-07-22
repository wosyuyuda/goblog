package controller

/*
 * @Description:
 * @Author: longfei
 * @FilePath: \go\controller\index.go
 */

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	d "goblog/model"
	"goblog/server"
	"goblog/util"

	"github.com/gin-gonic/gin"
)

/* type view struct {
	d.View
} */

//这里是详情页
func GetView(c *gin.Context) {
	id := c.Param("id")
	newList := server.GetViewlist("0", 1)  //最新的列表
	tuijian := server.GetViewlist("-1", 1) //推荐的列表
	vvv := util.GetView(id, 1)             //获取文章详情
	tp := GetTypeNew("0")                  //栏目分类
	c.HTML(http.StatusOK, "view.html", gin.H{
		"view":    vvv,
		"body":    template.HTML(vvv.Body),
		"types":   tp,
		"newlist": util.Imgsrc(newList),
		"tuijian": util.Imgsrc(tuijian),
	})
}

//这里是列表页
func Views(c *gin.Context) {

	id1 := c.Param("id")
	list := server.GetViewlist(id1, 1) //获取列表数据
	typeinfo := GetTypeNew(id1)
	newList := server.GetViewlist("0", 1)  //最新
	tuijian := server.GetViewlist("-4", 1) //推荐
	tp := GetTypeNew("0")                  //栏目全部分类
	page := c.DefaultQuery("page", "1")    //获取当前分页
	pagenum, _ := strconv.Atoi(page)       //获取分页数据
	i := util.GetTypeCount(id1)            //获取当前分类有多少条
	p := util.GetPage(i, pagenum)          //获取分页数据
	fmt.Printf("最新数据是%+v", typeinfo[0])
	fmt.Printf("分页信息%+v", p)
	c.HTML(http.StatusOK, "list.html", gin.H{
		"list":     util.Imgsrc(list),
		"typeinfo": typeinfo[0], //如果没有一条数据是会报错的
		"types":    tp,
		"newlist":  util.Imgsrc(newList),
		"tuijian":  util.Imgsrc(tuijian),
		"pageinfo": p,
	})
}

//这里是首页
func Lists1(c *gin.Context) {

	views := server.Findlist2("0") //最新
	jsonByte, _ := json.Marshal(views)
	fmt.Printf("首页的最新列表%+v", string(jsonByte))

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

	c.HTML(http.StatusOK, "index.html", gin.H{
		"list":    util.Imgsrc(views),
		"types":   tp,
		"tuijian": util.Imgsrc(tuijian),
		"tt":      tnew,
		"remen":   util.Imgsrc(remen),
		"swiper":  util.Imgsrc(swiper),
	})
}
