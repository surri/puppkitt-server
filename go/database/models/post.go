package models

import (
	"time"
	"puppkitt.com/v1/lib/common"
)

// Post data model
type Post struct {
	Id			string 	`gorm:"column:Id;type:varchar(26);primary_key"`
	CreateAt	time.Time 	`gorm:"column:CreateAt;type:DATETIME"`
	UpdateAt	time.Time 	`gorm:"column:UpdateAt;type:DATETIME"`
	DeleteAt	time.Time 	`gorm:"column:DeleteAt;type:DATETIME"`
	Title   	string 	`gorm:"column:Title;type:varchar(128)"`
	Text   		string 	`gorm:"column:Text;type:Text"`
	User   		User   	`gorm:"foreignkey:UserId"`
	UserId 		string 	`gorm:"column:UserId;type:varchar(26)"`
}

// Serialize serializes post data
func (p Post) Serialize() common.JSON {
	return common.JSON{
		"Id"		: p.Id,
		"CreateAt"	: p.CreateAt,
		"UpdateAt"	: p.UpdateAt,
		"Title"		: p.Title,
		"Text"		: p.Text,
		"User"		: p.User.Serialize(),
	}
}
