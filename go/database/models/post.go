package models

import (
	"hotpler.com/v1/lib/common"
)

// Post data model
type Post struct {
	Id			string 	`gorm:"column:id;type:varchar(26);primary_key"`
	CreateAt	int64 	`gorm:"column:create_at;type:bigint(20)"`
	UpdateAt	int64 	`gorm:"column:update_at;type:bigint(20)"`
	DeleteAt	int64 	`gorm:"column:delete_at;type:bigint(20)"`
	Title   	string 	`gorm:"column:title;type:varchar(128)"`
	Text   		string 	`gorm:"column:text;type:Text"`
	User   		User   	`gorm:"foreignkey:user_id"`
	UserId 		string 	`gorm:"column:user_id;type:varchar(26)"`
}

// Serialize serializes post data
func (p Post) Serialize() common.JSON {
	return common.JSON{
		"id"		: p.Id,
		"create_at"	: p.CreateAt,
		"update_at"	: p.UpdateAt,
		"title"		: p.Title,
		"text"		: p.Text,
		"user"		: p.User.Serialize(),
	}
}
