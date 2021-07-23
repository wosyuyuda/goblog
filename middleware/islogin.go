package middleware

import (
	"fmt"
	"goblog/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Islogin(c *gin.Context) {
	uid := util.GetSession(c, "uid")
	fmt.Printf("login信息是%+v", uid)
	if uid == 0 || uid == "" || uid == nil { //跳转到登陆界面，这里后面应该放到中间件里面去
		//	c.Redirect(http.StatusMovedPermanently, "/admin/login")
		c.Redirect(http.StatusTemporaryRedirect, "/admin/login") //307跳转到登陆页面
		c.Abort()
		return
	}
}
