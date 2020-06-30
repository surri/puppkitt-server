package models

import (
	"time"
	"puppkitt.com/v1/lib/common"
)

type Place struct {
	Id			string 	`gorm:"column:Id;type:varchar(26);primary_key"`
	CreateAt	time.Time 	`gorm:"column:CreateAt;type:DATETIME"`
	UpdateAt	time.Time 	`gorm:"column:UpdateAt;type:DATETIME"`
	DeleteAt	time.Time 	`gorm:"column:DeleteAt;type:DATETIME"`
	Type		string 	`gorm:"column:Type;type:varchar(26)"`
	Name		string 	`gorm:"column:Name;type:varchar(64)"`
	Phone		string 	`gorm:"column:Phone;type:varchar(13)"`
}

// Serialize serializes post data
func (p Place) Serialize() common.JSON {
	return common.JSON{
		"Id":		p.Id,
		"CreateAt": p.CreateAt,
		"UpdateAt": p.UpdateAt,
		"DeleteAt": p.DeleteAt,
		"text":		p.Type,
		"Name":		p.Name,
		"Phone":	p.Phone,
	}
}