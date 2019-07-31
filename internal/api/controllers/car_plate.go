package controllers

import (
	"do/internal/api/services"

	"github.com/astaxie/beego"
)

type CarplateController struct {
	beego.Controller
	service services.CarPlateService
}

func NewCarPlateController(carplateService services.CarPlateService) *CarplateController {
	return &CarplateController{
		service: carplateService,
	}
}

func (c *CarplateController) Prepare() {
	c.service = services.NewCarPlateService()
}

func (c *CarplateController) Get() {
	cars := c.service.GetAll()
	//cars := models.GetCarPlates()
	c.Data["json"] = cars
	c.ServeJSON()
}
