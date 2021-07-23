package controller

/*
 * @Description:
 * @Author: longfei
 * @LastEditors: ylf
 * @FilePath: \go\controller\user.go
 */

import (
	"goblog/dao"
	d "goblog/model"
	"goblog/server"
	"goblog/util"

	"github.com/gin-gonic/gin"
)

func EditUser(c *gin.Context) {
	uid := util.GetSession(c, "uid")
	var user d.User
	err := c.ShouldBind(&user)
	if err != nil {
		server.Fail(c)
		return
	}
	if user.Id == uid {
		//仅更新用户名,如果密码有则修改密码
		if user.Pwd != "" {
			user.Pwd = util.Md5jiayan(user.Pwd)
		}
		err = dao.MDB.Where("id = ?", user.Id).Updates(user).Error
		util.SetSession(c, "name", user.Name) //更新session里面的名称
	}
	server.IfRes(err, c)
}

//获取当前用户名
func GetUserName(c *gin.Context) {
	server.OkWithData(util.GetSession(c, "name"), c)
}
