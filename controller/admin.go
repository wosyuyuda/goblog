package controller

/*
 * @Description:
 * @Author: longfei
 * @FilePath: \go\controller\admin.go
 */
import (
	"fmt"
	"goblog/dao"
	d "goblog/model"
	"goblog/server"
	"goblog/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//后台的首页，再在其它页面添加一个文章列表的功能
func AdminIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "admin_index.html", gin.H{}) //进入管理首页
}

//获取当前文章，并且获取分类信息接口
func AdminGetId(c *gin.Context) {
	id := c.Param("id")
	view := util.GetView(id, 0)
	tp := GetTypeNew("0")
	c.JSON(200, gin.H{"msg": "获取成功", "code": 200, "data": view, "type": tp})
}

//后台的文章的管理页面
func AdminList(c *gin.Context) {
	pagenum := util.PageNum(c)               //获取当前第几页
	list := server.GetViewlist("0", pagenum) //列表数据
	i := util.GetTypeCount("0")              //获取总共的文章数量
	p := util.GetPage(i, pagenum)
	c.HTML(http.StatusOK, "admin_list.html", gin.H{
		"list":     list,
		"page":     pagenum,
		"pageinfo": p,
	})
}

//后台的添加文章详情页,笑死，gin模板语法跟vue.js模板语法冲突，不能直接用，直接加载静态页面吧
func AdminAddView(c *gin.Context) {
	c.HTML(http.StatusOK, "admin_addview.html", gin.H{}) //进入管理首页
}

//这里加一个接收前端数据的再返回数据就好啦，应该再加一个是否登陆判断
func AddView(c *gin.Context) {
	var data d.View
	var err error
	c.ShouldBind(&data)

	//typeid没办法直接拿过来，只好再单独获取了
	typeid := c.PostForm("typeid")
	t, _ := strconv.Atoi(typeid)
	data.Typeid = t
	conn := dao.MDB
	msg := "创建成功"
	if data.ID > 0 {
		err = conn.Model(&d.View{}).Where("id = ?", data.ID).Updates(&data).Error
		msg = "更新成功"
	} else {
		err = conn.Create(&data).Error
	}
	if err != nil {
		c.JSON(200, gin.H{"msg": "更新失败", "code": 200})
		fmt.Println("更新失败")
		c.Abort()
		return
	}
	util.BaiduLinksubmit(data.ID) //自动提交到百度收录
	util.DelAll()                 //删除缓存
	util.SetViewCache(&data)      //更新或者添加view都设置缓存
	c.JSON(200, gin.H{"msg": msg, "code": 200})
}

//用户登陆提交的页面
func Login(c *gin.Context) {
	id := c.PostForm("id")
	code := c.PostForm("code")
	if !server.CaptVerify(id, code) { //此处为验证码验证，后期再扩展
		c.JSON(200, gin.H{"msg": "验证码错误", "code": 400})
		c.Abort()
		return
	}
	name := c.PostForm("name")
	pwd := c.PostForm("pwd")
	conn := dao.MDB
	var u d.User
	conn.Where("name = ?", name).Find(&u)
	if util.Md5jiayan(pwd) != u.Pwd || u.Id == 0 {
		c.JSON(200, gin.H{"msg": "账号不存在或者密码错误", "code": 400})
		c.Abort()
		return
	}

	util.SetSession(c, "uid", u.Id)    //把用户的ID存进session
	util.SetSession(c, "name", u.Name) //把用户名存进session
	c.JSON(200, gin.H{"msg": "登陆成功", "code": 200})

}

func Logins(c *gin.Context) {
	c.HTML(http.StatusOK, "login1.html", gin.H{})
}

func Loginout(c *gin.Context) {
	util.NullSession(c, "uid")
	c.JSON(200, gin.H{"msg": "退出成功", "code": 200})
}

func GetCode(c *gin.Context) {
	id, b64s, err := server.CaptMake()
	if err != nil {
		fmt.Printf("错误信息是%+v", err)
	}
	c.JSON(200, gin.H{"msg": "获取成功", "code": 400, "id": id, "img": b64s})
}
