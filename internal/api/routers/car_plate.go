package routers

import (
	"do/internal/api/controllers"
	"do/internal/api/services"

	"github.com/astaxie/beego"
)

func init() {
	carPlateService := services.NewCarPlateService()
	carPlateController := controllers.NewCarPlateController(carPlateService)
	beego.Router("/", carPlateController)
}
