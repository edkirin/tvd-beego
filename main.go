package main

import (
	"fmt"
	"tvd-beego/models"
	_ "tvd-beego/routers"

	_ "github.com/lib/pq"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres",
		"user=televend password=televend host=localhost port=5432 dbname=televend sslmode=disable")
	orm.Debug = true
}

func listCompanies(o orm.Ormer) {
	var companies []*models.Company

	num, err := o.
		QueryTable(new(models.Company)).
		OrderBy("Id").
		RelatedSel().
		Filter("Alive", true).
		Filter("Access__FiscalConfiguration", true).
		All(&companies)

	if err == nil {
		for _, company := range companies {
			fmt.Println(
				company.Id,
				company.Caption,
				company.Access.FiscalConfiguration,
				company.Access.Fiscalization,
				company.Access.FiscalModuleB,
				company.Access.FiscalModuleC,
				company.Access.FiscalModuleD,
				company.Access.FiscalModuleE,
				company.Access.FiscalModuleF)
		}
		fmt.Println("Objects found:", num)
	} else {
		fmt.Println("Error:", err)
	}
}

func listMachines(o orm.Ormer) {
	var machines []*models.Machine

	num, err := o.
		QueryTable(new(models.Machine)).
		RelatedSel().
		Filter("Owner__Access__FiscalConfiguration", true).
		OrderBy("Id").
		Limit(50).
		All(&machines)

	if err == nil {
		for _, machine := range machines {
			fmt.Println(
				machine.Id,
				machine.Caption,
				machine.Owner.Caption,
				machine.Owner.Access.FiscalConfiguration)
		}
		fmt.Println("Objects found:", num)
	} else {
		fmt.Println("Error:", err)
	}
}

func main() {
	fmt.Println("-------------------------------------------------------------------")

	// o := orm.NewOrm()
	// listCompanies(o)
	// listMachines(o)

	beego.Run()
}
