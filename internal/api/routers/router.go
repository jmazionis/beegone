package routers

import (
	"github.com/ICanHaz/beegone/internal/api/controllers"
	"github.com/ICanHaz/beegone/internal/api/services"
	"github.com/ICanHaz/beegone/internal/api/storages"

	"github.com/astaxie/beego"
)

var carplateStorage = storages.CarPlateDb()
var carPlateService = services.NewCarPlateService(carplateStorage)

func init() {
	beego.Router("/api/carplates", &controllers.CarPlateController{CarPlateService: carPlateService}, "get:GetAll;post:Add;put:Update")
	beego.Router("/api/carplates/:id", &controllers.CarPlateController{CarPlateService: carPlateService}, "get:Get;delete:Delete")
}
