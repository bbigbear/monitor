package controllers

import (
	"campusMonitorSysterm/models"
	"fmt"

	"github.com/astaxie/beego"

	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/astaxie/beego/orm"
)

type WarnController struct {
	BaseController
}

func (this *WarnController) GetTotalWarnAndHandleWarnData() {
	o := orm.NewOrm()
	warn := new(models.Warn)
	query := o.QueryTable(warn)
	style := this.Input().Get("style")
	var maps []orm.Params
	switch style {
	case "day":

	case "week":
	case "mouth":
	case "year":
	default:
		this.ajaxMsg("请输入正确的类型", MSG_ERR_Param)
	}
}

func (this *WarnController) GetWarnData() {

	fmt.Println("获取预警信息")
	o := orm.NewOrm()
	warn := new(models.Warn)
	query := o.QueryTable(warn)

	//获取token
	token := this.Input().Get("token")
	if token == "" {
		fmt.Println("token 为空")
		this.ajaxMsg("token is not nil", MSG_ERR_Param)
	}
	//获取预警类型
	style := this.Input().Get("style")
	if style != "" {
		ws := new(models.WarnStyle)
		exist := o.QueryTable(ws).Filter("Style", style).Exist()
		if exist {
			query = query.Filter("Style", style)
		} else {
			this.ajaxMsg("不存在该类型", MSG_ERR_Param)
		}
	}
	appkey := beego.AppConfig.String("appkey")
	name, err := this.Token_auth(token, appkey)
	if err != nil {
		fmt.Println("token err", err)
		this.ajaxMsg("token err!", MSG_ERR_Verified)
	}
	fmt.Println("当前访问用户为:", name)

	var maps []orm.Params

	//index
	index, err := this.GetInt("index")
	if err != nil {
		fmt.Println("下标index不存在")
	}
	//pagemax  一页多少
	pagemax, err := this.GetInt("pagemax")
	if err != nil {
		fmt.Println("每页数量不存在")
	}
	//count
	count, err := query.Count()
	if err != nil {
		fmt.Println("获取数据总数为空")
		this.ajaxMsg("服务未知错误", MSG_ERR)
	}
	if pagemax != 0 {
		pagenum := int(math.Ceil(float64(count) / float64(pagemax)))

		if index > pagenum {
			//index = pagenum
			this.ajaxMsg("无法翻页了", MSG_ERR_Param)
		}
		fmt.Println("index&pagemax&pagenum", index, pagemax, pagenum)
	}
	query = query.Limit(pagemax, (index-1)*pagemax)
	_, err1 := query.OrderBy("-CreatTime").Values(&maps)
	if err1 != nil {
		fmt.Println("获取预警信息失败")
		this.ajaxMsg("获取预警信息失败", MSG_ERR_Resources)
	}
	this.ajaxList("获取预警信息成功", MSG_OK, count, maps)
}
