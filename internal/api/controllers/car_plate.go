package controllers

import (
	"do/internal/api/models"
	"do/internal/api/services"
	"encoding/json"
	"fmt"

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
		beego.Error("unable to unmarshal add carplate req")
		c.Ctx.Output.SetStatus(400)
		return
	}

	err = c.service.Add(carplate)
	if err != nil {
		beego.Error("unable to add carplate")
		c.Ctx.Output.SetStatus(400)
		return
	}

	c.Ctx.Output.SetStatus(201)
	c.Data["json"] = fmt.Sprintf("%s/%s", c.Ctx.Input.URL(), carplate.ID)
	c.ServeJSON()
}

func (c *CarPlateController) Update() {
	var carplate *models.CarPlate
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &carplate)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		beego.Error("unable to unmarshal update carplate req")
		return
	}

	err = c.service.Update(carplate)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		beego.Error("unable to update carplate")
		return
	}

	c.Ctx.Output.SetStatus(204)
	c.Data["json"] = carplate
	c.ServeJSON()
}

func (c *CarPlateController) Delete() {
	id := c.Ctx.Input.Param(":id")
	c.service.Delete(id)
	c.Ctx.Output.SetStatus(204)
}
