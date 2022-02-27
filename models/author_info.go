package models

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type AuthorInfo struct {
	Id           int       `orm:"pk" json:"id"`
	CreatedBy    *User     `orm:"rel(fk)" json:"createdBy"`
	DateCreated  time.Time `json:"dateCreated"`
	ModifiedBy   *User     `orm:"rel(fk)" json:"modifiedBy"`
	DateModified time.Time `json:"dateModified"`
}

func (u *AuthorInfo) TableName() string {
	return "author_infos"
}

func init() {
	orm.RegisterModel(
		new(AuthorInfo),
	)
}
