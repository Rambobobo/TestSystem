package service

import "myweb/BYSJ/model"

//定义AdminService接口
/*
用于通过管理员用户名和密码查询管理员实体，如果查询到则返回管理员实体和 true，否则返回 nil 和 false。
*/
type AdminService interface {
	GetByAdminNameAndPassword(username,password string)(model.Admin,bool)
}