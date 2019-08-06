package routers

import (
	"do/internal/api/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/api/carplates/", &controllers.CarPlateController{}, "get:GetAll")
	beego.Router("/api/carplates/:id", &controllers.CarPlateController{}, "get:Get")
	beego.Router("/api/carplates/add", &controllers.CarPlateController{}, "post:Add")
	beego.Router("/api/carplates/update", &controllers.CarPlateController{}, "put:Update")
	beego.Router("/api/carplates/delete/:id", &controllers.CarPlateController{}, "delete:Delete")
}
