package models

import (
	"time"
	"puppkitt.com/v1/lib/common"
)

// User data model
type User struct {
	Id				string 	`gorm:"column:Id;type:varchar(26);primary_key"`
	CreateAt		time.Time 	`gorm:"column:CreateAt;type:DATETIME"`
	UpdateAt		time.Time 	`gorm:"column:UpdateAt;type:DATETIME"`
	DeleteAt		time.Time 	`gorm:"column:DeleteAt;type:DATETIME"`
	UserName		string 	`gorm:"column:UserName;type:varchar(64)"`
	NickName		string 	`gorm:"column:NickName;type:varchar(64)"`
	Password		string 	`gorm:"column:Password;type:varchar(128)"`
}

// Serialize serializes user data
func (u *User) Serialize() common.JSON {
	return common.JSON{
		"Id":	u.Id,
		"UserName":	u.UserName,
		"NickName":	u.NickName,
	}
}

func (u *User) Read(m common.JSON) {
	u.Id		= m["Id"].(string)
	u.UserName	= m["UserName"].(string)
	u.NickName	= m["NickName"].(string)
}
