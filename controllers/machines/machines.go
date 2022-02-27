package machines

import (
	"fmt"
	"tvd-beego/models"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

type MachinesController struct {
	beego.Controller
}

type metaJson struct {
	ItemsCount int64 `json:"itemsCount"`
}

type responseJson struct {
	Items []*models.Machine `json:"items"`
	Meta  metaJson          `json:"meta"`
}

func (c *MachinesController) Get() {
	o := orm.NewOrm()

	var machines []*models.Machine

	num, err := o.
		QueryTable(new(models.Machine)).
		OrderBy("Id").
		Filter("Alive", true).
		RelatedSel("AuthorInfo", "AuthorInfo__CreatedBy", "AuthorInfo__ModifiedBy", "AuthorInfo__ModifiedBy").
		All(&machines)

	if err == nil {
		c.Data["json"] = responseJson{
			Items: machines,
			Meta: metaJson{
				ItemsCount: num,
			},
		}
		c.ServeJSON()
	} else {
		fmt.Println("Error:", err)
		c.CustomAbort(500, fmt.Sprintf("Internal server error: %s", err))
	}
}
