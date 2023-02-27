package model

import (
	"time"
)

type Admin struct {
	AdminId    int64     `xorm:"pk autoincr"`
	AdminName  string    `xorm:"varchar(32)"`
	CreateTime time.Time `xorm:"Datetime"`
	Pwd        string    `xorm:"varchar(255)"`
	Status     int64     `xorm:"default 0"`
}
