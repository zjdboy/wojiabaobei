package routers

import (
	"github.com/astaxie/beego"
	"wojiabaobei/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
