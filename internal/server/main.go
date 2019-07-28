package main

import (
	_ "do/internal/server/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

