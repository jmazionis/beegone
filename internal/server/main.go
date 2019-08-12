package main

import (
	_ "github.com/ICanHaz/beegone/internal/server/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

