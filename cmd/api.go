package main

import (
	_ "github.com/ICanHaz/beegone/internal/api/routers"

	"github.com/astaxie/beego"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.BConfig.Listen.HTTPPort = 9090
	beego.BConfig.CopyRequestBody = true
	beego.Run()
}
