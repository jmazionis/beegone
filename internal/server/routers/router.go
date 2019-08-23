package routers

import (
	"github.com/ICanHaz/beegone/internal/server/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("*", &controllers.MainController{}, "get:Index")
}
