package controllers

import (
	"do/internal/api/models"
	"do/internal/api/services"
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/segmentio/ksuid"
)

type CarPlateController struct {
	beego.Controller
	service services.CarPlateService
}

func (c *CarPlateController) Prepare() {
	c.service = services.NewCarPlateService()
}

func (c *CarPlateController) Get() {
	id := c.Ctx.Input.Param(":id")
	carplate, _ := c.service.Get(id)

	c.Data["json"] = carplate
	c.ServeJSON()
}

func (c *CarPlateController) GetAll() {
	cars := c.service.GetAll()

	c.Data["json"] = cars
	c.ServeJSON()
}

func (c *CarPlateController) Add() {
	var carplate *models.CarPlate
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &carplate)
	carplate.ID = ksuid.New().String()

	if err != nil {
		beego.Error("Unable to unmarshal add carplate req")
		c.Abort("400")
	}

	err = c.service.Add(carplate)
	if err != nil {
		beego.Error("Unable to add carplate")
	}

	c.Data["json"] = c.service.GetAll()
	c.ServeJSON()
}

func (c *CarPlateController) Update() {
	var carplate *models.CarPlate
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &carplate)
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		beego.Error("Unable to unmarshal udpate carplate req")
	}

	err = c.service.Update(carplate)
	if err != nil {
		beego.Error("Unable to update carplate")
	}

	c.Data["json"] = carplate
	c.ServeJSON()
}

func (c *CarPlateController) Delete() {
	id := c.Ctx.Input.Param(":id")
	c.service.Delete(id)
	c.Data["json"] = id
	c.ServeJSON()
}
