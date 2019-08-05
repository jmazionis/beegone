package controllers

import (
	"do/internal/api/models"
	"do/internal/api/services"
	"encoding/json"

	"github.com/astaxie/beego"
)

type CarPlateController struct {
	beego.Controller
	service services.CarPlateService
}

func (c *CarPlateController) Prepare() {
	c.service = services.NewCarPlateService()
}

func (c *CarPlateController) GetAll() {
	cars := c.service.GetAll()

	c.Data["json"] = cars
	c.ServeJSON()
}

func (c *CarPlateController) Add() {
	var carplate *models.CarPlate
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &carplate)
	if err != nil {
		beego.Error("Unable to unmarshal add carplate req")
	}

	c.service.Add(carplate)

	c.Data["json"] = c.service.GetAll() //map[string]interface{}{"id": carplate.ID}
	c.ServeJSON()

}
