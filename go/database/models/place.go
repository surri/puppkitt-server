package models

import (
	"hotpler.com/v1/lib/common"
)

type Place struct {
	Id			string 	`gorm:"column:id;type:varchar(26);primary_key"`
	CreateAt	int64 	`gorm:"column:create_at;type:bigint(20)"`
	UpdateAt	int64 	`gorm:"column:update_at;type:bigint(20)"`
	DeleteAt	int64 	`gorm:"column:delete_at;type:bigint(20)"`
	Type		string 	`gorm:"column:type;type:varchar(26)"`
	Name		string 	`gorm:"column:name;type:varchar(64)"`
	Phone		string 	`gorm:"column:phone;type:varchar(13)"`
}

// Serialize serializes post data
func (p Place) Serialize() common.JSON {
	return common.JSON{
		"id":		p.Id,
		"create_at": p.CreateAt,
		"update_at": p.UpdateAt,
		"delete_At": p.DeleteAt,
		"text":		p.Type,
		"name":		p.Name,
		"phone":	p.Phone,
	}
}