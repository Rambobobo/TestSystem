package controller

import (
	"myweb/BYSJ/service"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)
/*
管理员控制器
*/
//将发送的请求的字段映射为指定字段
type AdminLogin struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}
type AdminController struct{
	ctx iris.Context
	Serive service.AdminService
}

func (ac *AdminController) PostLogin(ctx iris.Context) mvc.Result{
	var adminlogin AdminLogin
	ac.ctx.ReadJSON(&adminlogin)

	if adminlogin.UserName ==""||adminlogin.Password==""{
		return mvc.Response{
			Object: map[string]interface{}{
				"status":"0",
				"success":"登录失败",
				"message":"用户名为空",
			},
		}
	}
	admin,exist :=ac.Serive.GetByAdminNameAndPassword(adminlogin.UserName,adminlogin.Password)
	if !exist{
		return mvc.Response{
			Object: map[string]interface{}{
				"status":"1",
				"success":"登录失败",
				"message":"管理员不存在",
			},
		}
	}
}