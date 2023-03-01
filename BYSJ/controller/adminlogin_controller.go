package controller

import (
	"myweb/BYSJ/service"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type AdminController struct{
	ctx iris.Context
}

func (c *AdminController) PostLogin(ctx iris.Context) mvc.Result{
	iris.New().Logger().Info("admin login")
	username :=ctx.FormValue("username")
	password := ctx.FormValue("password")
	admin,err:=service.AdminLogin(username,password)
	if err!=nil{
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString(err.Error())
		return
	}
	if admin==nil {
		ctx.ViewData("Error","用户名密码错误")
		return
		
	}
}