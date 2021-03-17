package models

type Roles struct {
	Id   int    `gorm:"column:id; serial"`
	Name string `gorm:"column:name;type:varchar(150)"`
	Permissions []Permissions
}

type Permissions struct {
	Id int `gorm:"column:id; serial"`
	UserID int
	RolesID int
}