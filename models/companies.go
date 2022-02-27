package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type Company struct {
	Id         int            `orm:"pk" json:"id"`
	Alive      bool           `json:"alive"`
	Caption    string         `orm:"size(100)" json:"caption"`
	Email      string         `orm:"size(100)" json:"email"`
	Address    string         `orm:"size(100)" json:"address"`
	City       string         `orm:"size(100)" json:"city"`
	Phone      string         `orm:"size(100)" json:"phone"`
	Access     *CompanyAccess `orm:"rel(one)" json:"access"`
	AuthorInfo *AuthorInfo    `orm:"rel(fk)"`
}

type CompanyAccess struct {
	Id                  int  `orm:"pk" json:"id"`
	FiscalConfiguration bool `json:"fiscalConfiguration"`
	Fiscalization       bool `json:"fiscalization"`
	FiscalModuleB       bool `json:"fiscalModuleB"`
	FiscalModuleC       bool `json:"fiscalModuleC"`
	FiscalModuleD       bool `json:"fiscalModuleD"`
	FiscalModuleE       bool `json:"fiscalModuleE"`
	FiscalModuleF       bool `json:"fiscalModuleF"`
}

func (u *Company) TableName() string {
	return "vending_companies"
}

func (u *CompanyAccess) TableName() string {
	return "administration_companyaccess"
}

func init() {
	orm.RegisterModel(
		new(Company),
		new(CompanyAccess),
	)
}
