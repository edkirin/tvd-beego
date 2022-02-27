package routers

import (
	companies "tvd-beego/controllers/companies"
	machines "tvd-beego/controllers/machines"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	// beego.Router("/", &controllers.MainController{})
	beego.Router("/companies", &companies.CompaniesController{}, "get:ListCompanies")
	beego.Router("/companies/:companyId:int/machines", &companies.CompaniesController{}, "get:ListMachines")
	beego.Router("/machines", &machines.MachinesController{})
}
