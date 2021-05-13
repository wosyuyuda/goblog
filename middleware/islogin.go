package middleware

import (
	"net/http"
	"test/util"

	"github.com/gin-gonic/gin"
)

func Islogin(c *gin.Context) {
	uid := util.GetSession(c, "uid")

	if uid == 0 || uid == "" || uid == nil { //跳转到登陆界面，这里后面应该放到中间件里面去
		//	c.Redirect(http.StatusMovedPermanently, "/admin/login")
		c.Redirect(http.StatusTemporaryRedirect, "/admin/login") //时间跳转到登陆页面
		c.Abort()
		return
	}
}
