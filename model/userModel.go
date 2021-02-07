package model

type User struct {
	UserId      int64	`json:"userId" gorm:"primary_key;not_null;auto_increment"`
	UserName    string	`json:"userName" gorm:"unique;not_null"`
	UserPwd     string	`json:"userPwd"`
	UserAvatar  string	`json:"userAvatar"`
}
