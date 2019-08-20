package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/ICanHaz/beegone/internal/api/models"
	"github.com/ICanHaz/beegone/internal/api/services"

	"github.com/astaxie/beego"
	"github.com/segmentio/ksuid"
)

type CarPlateController struct {
	beego.Controller
	CarPlateService services.CarPlateService
}

func (c *CarPlateController) Get() {
	id := c.Ctx.Input.Param(":id")
	carplate, err := c.CarPlateService.Get(id)

	if err != nil {
		c.Ctx.Output.SetStatus(404)
		return
	}

	c.Data["json"] = carplate
	c.ServeJSON()
}

func (c *CarPlateController) GetAll() {
	cars := c.CarPlateService.GetAll()

	c.Data["json"] = cars
	c.ServeJSON()
}

func (c *CarPlateController) Add() {
	var carplate *models.CarPlate
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &carplate)

	if err != nil {
		beego.Error(err)
		c.Ctx.Output.SetStatus(400)
		return
	}

	carplate.ID = ksuid.New().String()
	if valid, validationSummary := carplate.Validate(); !valid {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = validationSummary
		c.ServeJSON()
		return
	}

	err = c.CarPlateService.Add(carplate)
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

	if valid, validationSummary := carplate.Validate(); !valid {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = validationSummary
		c.ServeJSON()
		return
	}

	err = c.CarPlateService.Update(carplate)
	if err != nil {
		c.Ctx.Output.SetStatus(404)
		beego.Error(err)
		return
	}

	c.Ctx.Output.SetStatus(204)
}

func (c *CarPlateController) Delete() {
	id := c.Ctx.Input.Param(":id")
	c.CarPlateService.Delete(id)
	c.Ctx.Output.SetStatus(204)
}
