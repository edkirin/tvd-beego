package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type Machine struct {
	Id         int         `orm:"pk" json:"id"`
	Alive      bool        `json:"alive"`
	Caption    string      `json:"caption"`
	Owner      *Company    `orm:"rel(fk)" json:"-"`
	Warehouse  *Warehouse  `orm:"rel(fk)"`
	AuthorInfo *AuthorInfo `orm:"rel(fk)"`
}

func (u *Machine) TableName() string {
	return "machines"
}

func init() {
	orm.RegisterModel(
		new(Machine),
	)
}
