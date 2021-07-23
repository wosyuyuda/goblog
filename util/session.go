package util

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SetSession(c *gin.Context, key string, val interface{}) {
	session := sessions.Default(c)
	option := sessions.Options{MaxAge: 3600 * 8} //8小时后过期
	session.Options(option)
	session.Set(key, val) //把用户ID存进session,后面拿出来确认
	session.Save()
}
func GetSession(c *gin.Context, key string) interface{} {
	session := sessions.Default(c)
	return session.Get(key)
}

func NullSession(c *gin.Context, key string) {
	SetSession(c, key, 0) //笑死,清除不行,只能改为0
	session := sessions.Default(c)
	session.Delete(key)
	session.Clear()
}
