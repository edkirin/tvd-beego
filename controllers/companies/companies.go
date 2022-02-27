package companies

import (
	"fmt"
	"tvd-beego/models"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

type CompaniesController struct {
	beego.Controller
}

func (c *CompaniesController) ListCompanies() {
	o := orm.NewOrm()
	var companies []*models.Company

	count, err := o.
		QueryTable(new(models.Company)).
		OrderBy("Id").
		Filter("Alive", true).
		Filter("Access__FiscalConfiguration", true).
		RelatedSel().
		All(&companies)

	if err == nil {
		c.Data["json"] = CompaniesResponseJson{
			Items: companies,
			Meta: MetaJson{
				ItemsCount: count,
			},
		}
		c.ServeJSON()
	} else {
		c.CustomAbort(500, fmt.Sprintf("Internal server error: %s", err))
	}
}

func (c *CompaniesController) ListMachines() {
	companyId := c.Ctx.Input.Param(":companyId")

	o := orm.NewOrm()
	var machines []*models.Machine

	count, err := o.
		QueryTable(new(models.Machine)).
		OrderBy("Id").
		Filter("Alive", true).
		Filter("Owner", companyId).
		RelatedSel("AuthorInfo", "AuthorInfo__CreatedBy", "AuthorInfo__ModifiedBy", "AuthorInfo__ModifiedBy").
		All(&machines)

	if err == nil {
		c.Data["json"] = MachinesResponseJson{
			Items: machines,
			Meta: MetaJson{
				ItemsCount: count,
			},
		}
		c.ServeJSON()
	} else {
		c.CustomAbort(500, fmt.Sprintf("Internal server error: %s", err))
	}
}
