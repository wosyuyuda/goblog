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

	d "test/model"
	"test/util"

	"github.com/gin-gonic/gin"
)

/* type view struct {
	d.View
} */

//这里是详情页
func GetView(c *gin.Context) {
	id := c.Param("id")
	newList := FindListNew("0", 1)  //最新的列表
	tuijian := FindListNew("-1", 1) //推荐的列表
	vvv := util.GetView(id, 1)      //获取文章详情
	tp := GetTypeNew("0")           //栏目分类
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
	list := FindListNew(id1, 1) //获取列表数据
	typeinfo := GetTypeNew(id1)
	newList := FindListNew("0", 1)      //最新
	tuijian := FindListNew("-4", 1)     //推荐
	tp := GetTypeNew("0")               //栏目全部分类
	page := c.DefaultQuery("page", "1") //获取当前分页
	pagenum, _ := strconv.Atoi(page)    //获取分页数据
	i := util.GetTypeCount(id1)         //获取当前分类有多少条
	p := util.GetPage(i, pagenum)       //获取分页数据
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

	views := FindListNew("0", 1)    //最新
	tuijian := FindListNew("-4", 1) //推荐
	tp := GetTypeNew("0")           //分类
	remen := FindListNew("-3", 1)   //热门
	swiper := FindListNew("-2", 1)  //轮播
	tp22 := GetTypeNew("0")         //获取分类
	tnew := []d.Tp{}
	for _, v := range tp22 {
		v.Views = Findlist2(strconv.Itoa(int(v.ID)))
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

//获取当前分类ID下面的文章 ,一页10条
func FindListNew(id string, page int) (vi []d.View) {
	db := d.LinkDb() //连接数据库模型
	//下面查询的字段去掉body数据，列表页不获取这个，减少内存的使用
	db = db.Select("id", "created_at", "typeid", "title", "click", "pic", "tuijian", "swiper", "content")
	num := 10 //一页默认10条
	if page < 1 {
		page = 1
	}
	page -= 1 //过滤掉的页面要减 1
	order := "created_at desc"
	switch id {
	case "0":
		db.Limit(num).Offset(page * num).Order(order).Preload("Tps").Find(&vi)
	case "-1":
		db.Where("tuijian = ?", 1).Limit(num).Order(order).Preload("Tps").Find(&vi)
	case "-2":
		db.Where("swiper = ?", 1).Limit(num / 2).Order(order).Find(&vi)
	case "-3":
		db.Limit(num).Offset(page * num).Order("click desc").Find(&vi)
	case "-4":
		db.Where("tuijian = ?", 1).Limit(3).Order(order).Find(&vi)
	default:
		db.Where("typeid = ?", id).Limit(num).Offset(page * num).Order(order).Preload("Tps").Find(&vi)
	}
	return
}

//获取当前分类下面的10条文章
func Findlist2(id string) (vi []d.View) {
	db := d.LinkDb() //连接数据库模型
	db.Where("typeid = ?", id).Limit(10).Order("created_at desc").Find(&vi)
	return
}
