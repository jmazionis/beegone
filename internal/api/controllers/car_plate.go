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
	carplate, err := c.service.Get(id)

	if err != nil {
		c.Ctx.Output.SetStatus(404)
		return
	}

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
		beego.Error(err)
		c.Ctx.Output.SetStatus(400)
		return
	}

	err = c.service.Add(carplate)
	if err != nil {
		beego.Error(err)
		c.Ctx.Output.SetStatus(400)
		return
	}

	c.Ctx.Output.SetStatus(201)
	c.Ctx.Output.Header("location", fmt.Sprintf("%s/%s", c.Ctx.Input.URL(), carplate.ID))
}

func (c *CarPlateController) Update() {
	var carplate *models.CarPlate
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &carplate)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		beego.Error(err)
		return
	}

	err = c.service.Update(carplate)
	if err != nil {
		c.Ctx.Output.SetStatus(404)
		beego.Error(err)
		return
	}

	c.Ctx.Output.SetStatus(204)
}

func (c *CarPlateController) Delete() {
	id := c.Ctx.Input.Param(":id")
	c.service.Delete(id)
	c.Ctx.Output.SetStatus(204)
}
