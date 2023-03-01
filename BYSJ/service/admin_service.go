package service

import (
	"fmt"
	"myweb/BYSJ/model"

	"github.com/go-xorm/xorm"
)

type AdminService interface {
	GetByAdminNameAndPassword(username, password string) (model.Admin, bool)
}
//工厂函数
func NewAdminService(db *xorm.Engine) AdminService {
	return &adminService{
		engine: db,
	}
}
type adminService struct{
	engine *xorm.Engine
}
func (ac *adminService) GetByAdminNameAndPassword(username, password string) (model.Admin, bool) {
	var admin model.Admin

	ac.engine.Where("admin_name = ? and pwd = ? ", username, password).Get(&admin)

	fmt.Println(admin, "............", admin.AdminId != 0)

	return admin, admin.AdminId != 0
}
