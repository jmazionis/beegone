package main

import (
	_ "github.com/ICanHaz/beegone/internal/server/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.BConfig.Listen.HTTPPort = 9000
	beego.SetStaticPath("/static", "web/app/build/static")
	beego.SetStaticPath("/", "web/app/build")
	beego.BConfig.WebConfig.ViewsPath = "web/app/build"
	beego.Run()
}
