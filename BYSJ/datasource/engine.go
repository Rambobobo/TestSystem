package datasource

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var Engine *xorm.Engine

func init() {
	//数据库引擎
	var err error
	Engine, err = xorm.NewEngine("mysql", "root:root@tcp(127.0.0.1:3306)/database_name?charset=utf8mb4")
	if err != nil {
		panic(err)
	}
}
