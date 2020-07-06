package models

import (
	"puppkitt.com/v1/lib/common"
)

// User data model
type User struct {
	Id				string 	`gorm:"column:id;type:varchar(26);primary_key"`
	CreateAt		int64 	`gorm:"column:create_at;type:bigint(20)"`
	UpdateAt		int64 	`gorm:"column:update_at;type:bigint(20)"`
	DeleteAt		int64 	`gorm:"column:delete_at;type:bigint(20)"`
	Email			string	`gorm:"column:email;type:varchar(128)"`
	Phone			string	`gorm:"column:phone;type:varchar(128)"`
	DisplayName		string 	`gorm:"column:display_name;type:varchar(64)"`
	Password		string 	`gorm:"column:password;type:varchar(128)"`
}

// Serialize serializes user data
func (u *User) Serialize() common.JSON {
	return common.JSON{
		"id":	u.Id,
		"email":	u.Email,
		// "phone":	u.Phone,
		"display_name":	u.DisplayName,
	}
}

func (u *User) Read(m common.JSON) {
	u.Id		= m["id"].(string)
	u.Email	= m["first_name"].(string)
	// u.Phone	= m["phone"].(string)
	u.DisplayName	= m["display_name"].(string)
}
