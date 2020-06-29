package models

import (
	"puppkitt.com/v1/lib/common"
)

// User data model
type User struct {
	ID        string `gorm:"primary_key"`
	Username		string
	NickName		string
	PasswordHash	string
}

// Serialize serializes user data
func (u *User) Serialize() common.JSON {
	return common.JSON{
		"id":           u.ID,
		"username":     u.Username,
		"nick_name":	u.NickName,
	}
}

func (u *User) Read(m common.JSON) {
	u.ID = m["id"].(string)
	u.Username = m["username"].(string)
	u.NickName = m["nick_name"].(string)
}
