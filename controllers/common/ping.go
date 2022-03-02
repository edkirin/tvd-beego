package common

import (
	beego "github.com/beego/beego/v2/server/web"
)

type CommonController struct {
	beego.Controller
}

func (c *CommonController) Get() {
	c.Resp(200)
}
