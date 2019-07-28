package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.com"
	c.Data["Email"] = "test@gmail.com"
	c.TplName = "index.tpl"
}
