package controllers

import (
	"do/internal/api/models"
	"do/internal/api/services"

	"github.com/astaxie/beego"
)

type CarplateController struct {
	beego.Controller
	services.CarPlateService
}

func (c *CarplateController) Get() {
	cars := models.GetCarPlates()
	c.Data["json"] = cars
	c.ServeJSON()
}
