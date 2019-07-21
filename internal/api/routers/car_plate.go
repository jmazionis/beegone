package routers

import (
	"do/internal/api/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.CarplateController{})
}
