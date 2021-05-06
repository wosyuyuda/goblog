package controller

/*
 * @Description:
 * @Author: longfei
 * @Date: 2021-05-05 13:16:04
 * @LastEditTime: 2021-05-06 14:11:40
 * @LastEditors: longfei
 * @FilePath: \go\controller\admin.go
 */
import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/http"
	d "test/model"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Islogin(c *gin.Context) {
	session := sessions.Default(c)
	uid := session.Get("uid")
	fmt.Printf("aa%+v\n", uid)
	//手动删除了cookie后获取到的缓存是nil,没有找到用户信息的时候存的是0
	if uid == 0 || uid == "" || uid == nil { //跳转到登陆界面，这里后面应该放到中间件里面去
		c.Redirect(http.StatusMovedPermanently, "/admin/login")
		return
	}
}

//后台的列表页
func AdminIndex(c *gin.Context) {
	Islogin(c)
	c.HTML(http.StatusOK, "admin_index.html", gin.H{}) //进入管理首页
}

//后台的添加文章详情页
func AdminAddView(c *gin.Context) {
	Islogin(c)
	c.HTML(http.StatusOK, "admin_AddView.html", gin.H{}) //进入管理首页
}

//用户登陆提交的页面
func Login(c *gin.Context) {

	code := c.PostForm("code")
	if code != "1111" { //此处为验证码验证，后期再扩展
		c.JSON(200, gin.H{"msg": "验证码错误", "code": 400})
		return
	}
	name := c.PostForm("name")
	pwd := c.PostForm("pwd")
	conn := d.GetDb()
	u := new(User)
	conn.Where("name = ?", name).Find(&u)
	/* fmt.Printf("用户信息是%+v\n", u)
	fmt.Printf("传过来的账号是：%s 密码是：%s", name, pwd)
	fmt.Printf("后台的账号是：%s 密码是：%s，uid是 %d", u.Name, u.Pwd, u.Id) */

	h := md5.New()
	h.Write([]byte(pwd)) // 需要加密的字符串为 123456，如果密码放到前端加密，后端就可以不用
	if hex.EncodeToString(h.Sum(nil)) != u.Pwd || u.Id == 0 {
		c.JSON(200, gin.H{"msg": "账号不存在或者密码错误", "code": 400})
		return
	}
	session := sessions.Default(c)
	session.Set("uid", u.Id) //把用户ID存进session,后面拿出来确认
	session.Save()
	c.Redirect(http.StatusMovedPermanently, "/admin/")
	c.JSON(200, gin.H{"msg": "登陆成功", "code": 200})

}
