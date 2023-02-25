package main

import "github.com/kataras/iris/v12"

/**
程序主入口
**/

func main() {
	app:= iris.New()
	app.Run(iris.Addr(":9000"))
}
