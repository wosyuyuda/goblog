package controller

/*
 * @Description:
 * @Author: longfei
 * @LastEditors: ylf
 * @FilePath: \go\controller\user.go
 */

import (
	"fmt"
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
	fmt.Printf("接收到的数据是%+v", user)
	if err != nil {
		server.Fail(c)
		return
	}

	//仅更新用户名,如果密码有则修改密码
	if user.Pwd != "" { //修改账号跟密码
		user.Pwd = util.Md5jiayan(user.Pwd)
		err = dao.MDB.Where("id = ?", uid).Updates(user).Error
	} else { //仅修改账号
		err = dao.MDB.Where("id = ?", uid).Update("name", user.Name).Error
	}

	util.SetSession(c, "name", user.Name) //更新session里面的名称

	server.IfRes(err, c)
}

//获取当前用户名
func GetUserName(c *gin.Context) {
	server.OkWithData(util.GetSession(c, "name"), c)
}

//获取个人简介信息
func GetUserInfo(c *gin.Context) {
	vvv := util.GetView("0", 2) //获取文章详情
	server.OkWithData(vvv, c)
}

//编辑个人简介
func EditUserInfo(c *gin.Context) {
	var data d.View
	err := c.ShouldBind(&data)
	if err != nil {
		server.FailWithMessage(err.Error(), c)
		return
	}
	if data.Tuijian == 3 {
		data.Status = 3
	}
	if data.ID != 0 {
		err = dao.MDB.Where("id = ?", data.ID).Updates(data).Error
	} else {
		err = dao.MDB.Create(&data).Error
	}
	util.DelAll() //删除缓存
	server.ResDataError(data, err, c)
}
