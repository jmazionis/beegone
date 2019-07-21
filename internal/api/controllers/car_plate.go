package controllers

import (
	"do/internal/api/models"

	"github.com/astaxie/beego"
)

type CarplateController struct {
	beego.Controller
}

func (c *CarplateController) Get() {
	cars := models.GetCarPlates()
	c.Data["json"] = cars
	c.ServeJSON()
}
