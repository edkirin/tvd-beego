package models

import "github.com/beego/beego/v2/client/orm"

type User struct {
	Id        int    `orm:"pk" json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

func (u *User) TableName() string {
	return "auth_user"
}

type CustomUser struct {
	// Id      int   `orm:"pk" json:"id"`
	Tel      string `json:"tel"`
	TimeZone string `json:"timeZone"`
}

func (u *CustomUser) TableName() string {
	return "custom_users"
}

func init() {
	orm.RegisterModel(
		new(User),
		new(CustomUser),
	)
}
