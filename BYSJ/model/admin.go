package model

import (
	"time"
)

type Admin struct {
	AdminId    int64     `xorm:"pk autoincr"`
	AdminName  string    `xorm:"varchar(32)"`
	CreateTime time.Time `xorm:"Datetime"`
	Status     int64     `xorm:"default 0"`
	Avatar     string    `xorm:"varchar(255)"`
	Pwd        string    `xorm:"varchar(255)"`
	CityName   string    `xorm:"varchar(12)"`
	CityId     int64     `xorm:"index"`
}
