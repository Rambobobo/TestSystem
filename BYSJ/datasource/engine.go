package datasource

import (
	"myweb/BYSJ/model"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var (
	engine *xorm.Engine
)

func NewMysqlEngine() *xorm.Engine {
	//创建数据库引擎对象
	engine, err := xorm.NewEngine("mysql", "root:root@tcp(127.0.0.1:3306)/database_name?charset=utf8mb4")

	//同步数据库表格
	err = engine.Sync2(new(model.Admin))

	if err != nil {
		panic(err.Error())
	}
	//显示sql语句
	engine.ShowSQL(true)
	engine.SetMaxOpenConns(10)

	return engine
}
