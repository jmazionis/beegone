package routers

import (
	"do/internal/api/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/api/GetAll", &controllers.CarPlateController{}, "get:GetAll")
	beego.Router("/api/Add", &controllers.CarPlateController{}, "post:Add")
}
