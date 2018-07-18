package routers

import (
	"campusMonitorSysterm/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/api/v1/login", &controllers.LoginController{}, "post:Login")
	beego.Router("/api/v1/warn/getwarndata", &controllers.WarnController{}, "*:GetWarnData")
	beego.Router("/api/v1/warn/getcontrastdata", &controllers.WarnController{}, "*:GetTotalWarnAndHandleWarnData")
	beego.Router("/api/v1/warn/getrankdata", &controllers.WarnController{}, "*:GetRank")
	beego.Router("/api/v1/warn/getwarnstyle", &controllers.WarnController{}, "*:GetWarnStyle")
}
