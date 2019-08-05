package routers

import (
	"do/internal/api/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/api/carplate/GetAll", &controllers.CarPlateController{}, "get:GetAll")
	beego.Router("/api/carplate/Get/:id", &controllers.CarPlateController{}, "get:Get")
	beego.Router("/api/carplate/Add", &controllers.CarPlateController{}, "post:Add")
	beego.Router("/api/carplate/Update", &controllers.CarPlateController{}, "put:Update")
	beego.Router("/api/carplate/Delete/:id", &controllers.CarPlateController{}, "delete:Delete")
}
