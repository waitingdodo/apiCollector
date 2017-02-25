package routers

import (
	"apiCollector/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/stopService", &controllers.MainController{},"get:StopService")

	beego.Router("/updateTimeLimit", &controllers.MainController{},"get:UpdateTimeLimit")
}
