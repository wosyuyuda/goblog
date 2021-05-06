package server

/*
 * @Description:
 * @Author: longfei
 * @FilePath: \go\server\session.go
 */

import (
	"github.com/gin-contrib/sessions"

	"github.com/gin-gonic/gin"
)

//这个页面暂弃，没啥用。后期再写一些功能放在这里
func Get(c *gin.Context) *sessions.Session {
	session := sessions.Default(c)
	return &session

}
