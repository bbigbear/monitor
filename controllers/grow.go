package controllers

import (
	//	"campusMonitorSysterm/models"
	"fmt"
	//	"math"
	//	"time"

	"github.com/astaxie/beego"

	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/astaxie/beego/orm"
)

type GrowController struct {
	BaseController
}

func (this *GrowController) GrowthTargets() {
	fmt.Printf("成长目标")
	o := orm.NewOrm()
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

	//获取style
	style := this.Input().Get("style")
	if style == "" {
		this.ajaxMsg("style不能为空", MSG_ERR_Param)
	}
	//获取人数
	count := this.Input().Get("count")
	switch style {
	case "成绩":
		_, err := o.Raw("SELECT sid,sname,SUM(CASE subject WHEN '高等数学' THEN grade ELSE 0 END) as math , SUM(CASE subject WHEN '体育' THEN grade ELSE 0 END) as pe , SUM(CASE subject WHEN '大学英语' THEN grade ELSE 0 END) as en , SUM(CASE subject WHEN '中国文学' THEN grade ELSE 0 END) as history , SUM(CASE subject WHEN '计算机基础' THEN grade ELSE 0 END) as computer , SUM(CASE subject WHEN '心理学' THEN grade ELSE 0 END) as psychology , SUM(grade) as sum FROM score GROUP BY sname ORDER BY SUM(grade) DESC LIMIT ?", count).Values(&maps)
		if err != nil {
			fmt.Println("get style score err!", err.Error())
		}
	case "阅读":
		_, err := o.Raw("SELECT sid,sname,SUM(CASE subject WHEN '高等数学' THEN grade ELSE 0 END) as math , SUM(CASE subject WHEN '体育' THEN grade ELSE 0 END) as pe , SUM(CASE subject WHEN '大学英语' THEN grade ELSE 0 END) as en , SUM(CASE subject WHEN '中国文学' THEN grade ELSE 0 END) as history , SUM(CASE subject WHEN '计算机基础' THEN grade ELSE 0 END) as computer , SUM(CASE subject WHEN '心理学' THEN grade ELSE 0 END) as psychology , SUM(grade) as sum FROM score GROUP BY sname ORDER BY SUM(grade) DESC").Values(&maps)
		if err != nil {
			fmt.Println("get style score err!", err.Error())
		}
	case "科研":
		_, err := o.Raw("SELECT sid,sname,SUM(CASE subject WHEN '高等数学' THEN grade ELSE 0 END) as math , SUM(CASE subject WHEN '体育' THEN grade ELSE 0 END) as pe , SUM(CASE subject WHEN '大学英语' THEN grade ELSE 0 END) as en , SUM(CASE subject WHEN '中国文学' THEN grade ELSE 0 END) as history , SUM(CASE subject WHEN '计算机基础' THEN grade ELSE 0 END) as computer , SUM(CASE subject WHEN '心理学' THEN grade ELSE 0 END) as psychology , SUM(grade) as sum FROM score GROUP BY sname ORDER BY SUM(grade) DESC").Values(&maps)
		if err != nil {
			fmt.Println("get style score err!", err.Error())
		}
	case "上网":
		_, err := o.Raw("SELECT sid,sname,SUM(CASE subject WHEN '高等数学' THEN grade ELSE 0 END) as math , SUM(CASE subject WHEN '体育' THEN grade ELSE 0 END) as pe , SUM(CASE subject WHEN '大学英语' THEN grade ELSE 0 END) as en , SUM(CASE subject WHEN '中国文学' THEN grade ELSE 0 END) as history , SUM(CASE subject WHEN '计算机基础' THEN grade ELSE 0 END) as computer , SUM(CASE subject WHEN '心理学' THEN grade ELSE 0 END) as psychology , SUM(grade) as sum FROM score GROUP BY sname ORDER BY SUM(grade) DESC").Values(&maps)
		if err != nil {
			fmt.Println("get style score err!", err.Error())
		}
	case "纪律":
		_, err := o.Raw("SELECT sid,sname,SUM(CASE subject WHEN '高等数学' THEN grade ELSE 0 END) as math , SUM(CASE subject WHEN '体育' THEN grade ELSE 0 END) as pe , SUM(CASE subject WHEN '大学英语' THEN grade ELSE 0 END) as en , SUM(CASE subject WHEN '中国文学' THEN grade ELSE 0 END) as history , SUM(CASE subject WHEN '计算机基础' THEN grade ELSE 0 END) as computer , SUM(CASE subject WHEN '心理学' THEN grade ELSE 0 END) as psychology , SUM(grade) as sum FROM score GROUP BY sname ORDER BY SUM(grade) DESC").Values(&maps)
		if err != nil {
			fmt.Println("get style score err!", err.Error())
		}
	case "消费":
		_, err := o.Raw("SELECT sid,sname,SUM(CASE subject WHEN '高等数学' THEN grade ELSE 0 END) as math , SUM(CASE subject WHEN '体育' THEN grade ELSE 0 END) as pe , SUM(CASE subject WHEN '大学英语' THEN grade ELSE 0 END) as en , SUM(CASE subject WHEN '中国文学' THEN grade ELSE 0 END) as history , SUM(CASE subject WHEN '计算机基础' THEN grade ELSE 0 END) as computer , SUM(CASE subject WHEN '心理学' THEN grade ELSE 0 END) as psychology , SUM(grade) as sum FROM score GROUP BY sname ORDER BY SUM(grade) DESC").Values(&maps)
		if err != nil {
			fmt.Println("get style score err!", err.Error())
		}
	default:
		this.ajaxMsg("请输入正确的类型", MSG_ERR_Param)
	}

	this.ajaxList("获取数据成功", MSG_OK, 1, maps)
}
