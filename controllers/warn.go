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
	out := make(map[string]interface{})
	var maps []orm.Params
	var handleMaps []orm.Params
	//获取token
	token := this.Input().Get("token")
	if token == "" {
		fmt.Println("token 为空")
		this.ajaxMsg("token is not nil", MSG_ERR_Param)
	}
	appkey := beego.AppConfig.String("appkey")
	name, err := this.Token_auth(token, appkey)
	if err != nil {
		fmt.Println("token err", err)
		this.ajaxMsg("token err!", MSG_ERR_Verified)
	}
	fmt.Println("当前访问用户为:", name)
	nowtime := time.Now().Format("2006-01-02 15:04:05")

	style := this.Input().Get("style")
	switch style {
	case "day":
		_, err := o.Raw("select HOUR(warn_time) as hour,warn_name,COUNT(*) as count from warn WHERE warn_time between DATE_SUB(?,INTERVAL 1 DAY) and ? GROUP BY hour;", nowtime, nowtime).Values(&maps)
		if err != nil {
			fmt.Println("get warn 1 day err!", err.Error())
		}
		_, err1 := o.Raw("select HOUR(warn_time) as hour,warn_name,COUNT(*) as count from warn WHERE warn_time between DATE_SUB(?,INTERVAL 1 DAY) and ? and status = '已处理' GROUP BY hour;", nowtime, nowtime).Values(&handleMaps)
		if err1 != nil {
			fmt.Println("get handle warn 1 day err!", err.Error())
		}
	case "week":
		_, err := o.Raw("select DAY(warn_time) as day,warn_name,COUNT(*) as count from warn WHERE warn_time between DATE_SUB(?,INTERVAL 1 WEEK) and ? GROUP BY day;", nowtime, nowtime).Values(&maps)
		if err != nil {
			fmt.Println("get warn 1 week err!", err.Error())
		}
		_, err1 := o.Raw("select DAY(warn_time) as day,warn_name,COUNT(*) as count from warn WHERE warn_time between DATE_SUB(?,INTERVAL 1 WEEK) and ? and status = '已处理' GROUP BY day;", nowtime, nowtime).Values(&handleMaps)
		if err1 != nil {
			fmt.Println("get handle warn 1 week err!", err.Error())
		}
	case "month":
		_, err := o.Raw("select DAY(warn_time) as day,warn_name,COUNT(*) as count from warn WHERE warn_time between DATE_SUB(?,INTERVAL 1 MONTH) and ? GROUP BY day;", nowtime, nowtime).Values(&maps)
		if err != nil {
			fmt.Println("get warn 1 month err!", err.Error())
		}
		_, err1 := o.Raw("select DAY(warn_time) as day,warn_name,COUNT(*) as count from warn WHERE warn_time between DATE_SUB(?,INTERVAL 1 MONTH) and ? and status = '已处理' GROUP BY day;", nowtime, nowtime).Values(&handleMaps)
		if err1 != nil {
			fmt.Println("get handle month 1 day err!", err.Error())
		}
	case "year":
		_, err := o.Raw("select MONTH(warn_time) as month,warn_name,COUNT(*) as count from warn WHERE warn_time between DATE_SUB(?,INTERVAL 1 YEAR) and ? GROUP BY month;", nowtime, nowtime).Values(&maps)
		if err != nil {
			fmt.Println("get warn 1 year err!", err.Error())
		}
		_, err1 := o.Raw("select MONTH(warn_time) as month,warn_name,COUNT(*) as count from warn WHERE warn_time between DATE_SUB(?,INTERVAL 1 YEAR) and ? and status = '已处理' GROUP BY month;", nowtime, nowtime).Values(&handleMaps)
		if err1 != nil {
			fmt.Println("get handle warn 1 year err!", err.Error())
		}
	default:
		this.ajaxMsg("请输入正确的类型", MSG_ERR_Param)
	}
	fmt.Println("maps len", len(maps))
	//	for _, m := range maps {
	//		fmt.Println("status", m)
	//		fmt.Println("status", m["status"].(string))
	//		if m["status"].(string) == "已处理" {
	//			handleMaps = append(handleMaps, m)
	//		}
	//	}
	out["total"] = maps
	out["handle"] = handleMaps
	this.ajaxList("获取信息成功", MSG_OK, int64(len(maps)), out)
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
	appkey := beego.AppConfig.String("appkey")
	name, err := this.Token_auth(token, appkey)
	if err != nil {
		fmt.Println("token err", err)
		this.ajaxMsg("token err!", MSG_ERR_Verified)
	}
	fmt.Println("当前访问用户为:", name)
	//获取预警类型
	style := this.Input().Get("style")
	if style != "" {
		ws := new(models.WarnStyle)
		exist := o.QueryTable(ws).Filter("Name", style).Exist()
		if exist {
			query = query.Filter("WarnName", style)
		} else {
			this.ajaxMsg("不存在该类型", MSG_ERR_Param)
		}
	}
	//获取类型
	status := this.Input().Get("status")
	if status != "" {
		query = query.Filter("Status", status)
	}

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
	_, err1 := query.OrderBy("-WarnTime").Values(&maps)
	if err1 != nil {
		fmt.Println("获取预警信息失败")
		this.ajaxMsg("获取预警信息失败", MSG_ERR_Resources)
	}
	this.ajaxList("获取预警信息成功", MSG_OK, count, maps)
}

//获取前几位
func (this *WarnController) GetRank() {

	fmt.Println("获取排名")
	o := orm.NewOrm()
	warn := new(models.Warn)
	//query := o.QueryTable(warn)
	var maps []orm.Params

	//获取token
	token := this.Input().Get("token")
	if token == "" {
		fmt.Println("token 为空")
		this.ajaxMsg("token is not nil", MSG_ERR_Param)
	}
	appkey := beego.AppConfig.String("appkey")
	name, err := this.Token_auth(token, appkey)
	if err != nil {
		fmt.Println("token err", err)
		this.ajaxMsg("token err!", MSG_ERR_Verified)
	}
	fmt.Println("当前访问用户为:", name)

	//state 1[个人]2[类型]
	state := this.Input().Get("state")
	if state == "" {
		this.ajaxMsg("state不能为空", MSG_ERR_Param)
	}
	if state == "1" {
		//获取排名数量
		count, err := this.GetInt("count")
		if err != nil {
			this.ajaxMsg("count为int类型", MSG_ERR_Param)
		}
		if count == 0 {
			this.ajaxMsg("count不能为空", MSG_ERR_Param)
		}
		_, err1 := o.Raw("SELECT sname , warn_name , warn_info , count( * ) AS count FROM warn GROUP BY sname ORDER BY count DESC LIMIT ?", count).Values(&maps)
		if err1 != nil {
			fmt.Println("get student rank info err!", err.Error())
		}

	} else if state == "2" {
		_, err2 := o.Raw("SELECT warn_name , warn_info , count( * ) AS count FROM warn GROUP BY warn_name ORDER BY count DESC").Values(&maps)
		if err2 != nil {
			fmt.Println("get style rank info err!", err.Error())
		}
	} else {
		this.ajaxMsg("state err", MSG_ERR_Param)
	}

	//total
	total, err := o.QueryTable(warn).Count()
	if err != nil {
		fmt.Println("mysql get count err", err.Error())
		this.ajaxMsg("内部错误", MSG_ERR)
	}

	this.ajaxList("获取预警信息成功", MSG_OK, total, maps)
}

//获取预警类型
func (this *WarnController) GetWarnStyle() {

	fmt.Println("获取预警类型")
	o := orm.NewOrm()
	warn := new(models.WarnStyle)
	var maps orm.ParamsList

	//获取token
	token := this.Input().Get("token")
	if token == "" {
		fmt.Println("token 为空")
		this.ajaxMsg("token is not nil", MSG_ERR_Param)
	}
	appkey := beego.AppConfig.String("appkey")
	name, err := this.Token_auth(token, appkey)
	if err != nil {
		fmt.Println("token err", err)
		this.ajaxMsg("token err!", MSG_ERR_Verified)
	}
	fmt.Println("当前访问用户为:", name)

	num, err := o.QueryTable(warn).Distinct().ValuesFlat(&maps, "name")
	if err != nil {
		fmt.Println("get warn style err", err.Error())
	}
	this.ajaxList("获取预警类型成功", MSG_OK, num, maps)
}
