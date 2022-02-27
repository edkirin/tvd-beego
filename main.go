package main

import (
	"fmt"
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

func main() {
	fmt.Println("-------------------------------------------------------------------")
	beego.Run()
}
