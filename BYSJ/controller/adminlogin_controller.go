package controller

import(
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type AdminController struct{
	ctx iris.Context
}

func (c *AdminController) PostLogin() mvc.Result{
	iris.New().Logger().Info("admin login")
}