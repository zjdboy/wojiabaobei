package main

import (
	"github.com/astaxie/beego"
	_ "wojiabaobei/routers"
)

func main() {
	beego.Run()
	beego.AutoRender()
}
