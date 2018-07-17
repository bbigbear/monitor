package controllers

import (
	"campusMonitorSysterm/models"
	"fmt"
	"math"
	"time"

	"github.com/astaxie/beego"

	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/astaxie/beego/orm"
)

type WarnController struct {
	BaseController
}

func (this *WarnController) GetTotalWarnAndHandleWarnData() {
	o := orm.NewOrm()
	style := this.Input().Get("style")
	var maps []orm.Params
	nowtime := time.Now().Format("2006-01-02 15:04:05")
	switch style {
	case "day":
		_, err := o.Raw("select * from warn WHERE warn_time between DATE_SUB(?,INTERVAL 1 DAY) and ?", nowtime, nowtime).Values(&maps)
		if err != nil {
			fmt.Println("get warn 1 day err!", err.Error())
		}
	case "week":
		_, err := o.Raw("select * from warn WHERE warn_time between DATE_SUB(?,INTERVAL 1 WEEK) and ?", nowtime, nowtime).Values(&maps)
		if err != nil {
			fmt.Println("get warn 1 week err!", err.Error())
		}
	case "mouth":
		_, err := o.Raw("select * from warn WHERE warn_time between DATE_SUB(?,INTERVAL 1 MONTH) and ?", nowtime, nowtime).Values(&maps)
		if err != nil {
			fmt.Println("get warn 1 day err!", err.Error())
		}
	case "year":
		_, err := o.Raw("select * from warn WHERE warn_time between DATE_SUB(?,INTERVAL 1 YEAR) and ?", nowtime, nowtime).Values(&maps)
		if err != nil {
			fmt.Println("get warn 1 day err!", err.Error())
		}
	default:
		this.ajaxMsg("请输入正确的类型", MSG_ERR_Param)
	}
	this.ajaxList("获取信息成功", MSG_OK, int64(len(maps)), maps)
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
