package routers

import (
	"do/internal/api/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/api/carplates", &controllers.CarPlateController{}, "get:GetAll;post:Add;put:Update")
	beego.Router("/api/carplates/:id", &controllers.CarPlateController{}, "get:Get;delete:Delete")

}
