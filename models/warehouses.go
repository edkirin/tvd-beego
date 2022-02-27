package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type Warehouse struct {
	Id         int      `orm:"pk" json:"id"`
	Caption    string   `json:"caption"`
	ExternalId string   `json:"externalId"`
	Type       int      `json:"type"`
	Address    string   `json:"address"`
	Owner      *Company `orm:"rel(fk)" json:"-"`
}

func (u *Warehouse) TableName() string {
	return "administration_warehouse"
}

func init() {
	orm.RegisterModel(
		new(Warehouse),
	)
}
