package service

import (
	"myweb/BYSJ/datasource"
	"myweb/BYSJ/model"

	"github.com/go-xorm/xorm"
)



func AdminLogin(username string,password string) (model.Admin,error){
	admin:=&model.Admin{AdminName: username,Pwd: password}
	has,err :=datasource.engine.Get(admin)
	if err!=nil {
		return nil,err
	}
	if !has {
		return nil,nil		
	}
	return admin,nil
}