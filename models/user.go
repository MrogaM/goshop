package models

type User struct {
	Id          int    `gorm:"column:id; serial"`
	Username    string `gorm:"column:username; not null; type:varchar(150)"`
	Password    string `gorm:"column:password; not null; type:varchar(150)"`
	Email       string `gorm:"column:email; type:varchar(150)"`
	Roles       string `gorm:"column:roles;"`
	Permissions []Permissions
}