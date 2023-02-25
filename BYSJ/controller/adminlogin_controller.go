package controller

import(
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type AdminController struct{
	ctx iris.Context
}

func (c *AdminController) GetLogin() mvc.View {
	return mvc.View{
		Name: "admin/login.html",
	}
}

func (c *AdminController)PostLogin() {
	
}