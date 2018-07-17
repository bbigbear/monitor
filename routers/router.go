package routers

import (
	"gree/push_svr/campusMonitorSysterm/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/api/v1/login", &controllers.LoginController{}, "post:Login")
	beego.Router("/api/v1/warn/getdata", &controllers.WarnController{}, "*:GetWarnData")
}
