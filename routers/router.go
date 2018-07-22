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
	//成长
	beego.Router("/api/v1/grow/growth_targets", &controllers.GrowController{}, "*:GrowthTargets")
	beego.Router("/api/v1/grow/growth_record", &controllers.GrowController{}, "*:GrowthRecord")
	beego.Router("/api/v1/grow/excellent_portrait", &controllers.GrowController{}, "*:ExcellentStudentPortrait")
	beego.Router("/api/v1/grow/get_student_info", &controllers.GrowController{}, "*:StudentInformation")
}
