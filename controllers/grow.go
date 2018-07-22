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

func (this *GrowController) ExcellentStudentPortrait() {
	fmt.Println("获取优秀学生画像")
	o := orm.NewOrm()
	filters := make([]interface{}, 0)
	var school_maps []orm.Params
	var excellent_maps []orm.Params
	var total_maps []orm.Params
	var reslut1 []map[string]interface{}

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
	count := 1
	//成绩
	_, _err := o.Raw("SELECT sid,sname,SUM(grade) as sum FROM score GROUP BY sname ORDER BY SUM(grade) DESC LIMIT ?", count).Values(&excellent_maps)
	if _err != nil {
		fmt.Println("get style score err!", err.Error())
	}
	_, err1 := o.Raw("SELECT AVG(avg) as count FROM (SELECT sname,SUM(grade) as avg FROM score GROUP BY sname) t1;").Values(&school_maps)
	if err1 != nil {
		fmt.Println("get style score err!", err.Error())
	}
	_, err111 := o.Raw("SELECT COUNT(DISTINCT subject)*100 as count FROM score;").Values(&total_maps)
	if err111 != nil {
		fmt.Println("get style score err!", err.Error())
	}
	filters = append(filters, "score", excellent_maps, school_maps, total_maps)
	//阅读
	_, err2 := o.Raw("SELECT sid,sname,COUNT(*)/1 as count FROM monitor.read GROUP BY sname ORDER BY count(*) DESC LIMIT ?;", count).Values(&excellent_maps)
	if err2 != nil {
		fmt.Println("get style score err!", err.Error())
	}
	_, err3 := o.Raw("SELECT COUNT(*)/COUNT(DISTINCT sid) as count FROM monitor.read;").Values(&school_maps)
	if err3 != nil {
		fmt.Println("get style score err!", err.Error())
	}
	_, err33 := o.Raw("SELECT COUNT(*)/1 as count FROM monitor.read;").Values(&total_maps)
	if err33 != nil {
		fmt.Println("get style score err!", err.Error())
	}
	filters = append(filters, "read", excellent_maps, school_maps, total_maps)
	//科研
	_, err4 := o.Raw("SELECT sid,sname,COUNT(*)/1 as count FROM research GROUP BY sname ORDER BY count(*) DESC LIMIT ?;", count).Values(&excellent_maps)
	if err4 != nil {
		fmt.Println("get research err!", err.Error())
	}
	_, err5 := o.Raw("SELECT COUNT(*)/COUNT(DISTINCT sid) as count FROM research;").Values(&school_maps)
	if err5 != nil {
		fmt.Println("get research err!", err.Error())
	}
	_, err55 := o.Raw("SELECT COUNT(*)/1 as count FROM research;").Values(&total_maps)
	if err55 != nil {
		fmt.Println("get style score err!", err.Error())
	}
	filters = append(filters, "research", excellent_maps, school_maps, total_maps)
	//上网
	_, err6 := o.Raw("SELECT sid,sname,COUNT(*)/1 as count FROM surfer GROUP BY sname ORDER BY count(*) LIMIT ?;", count).Values(&excellent_maps)
	if err6 != nil {
		fmt.Println("get surfer err!", err.Error())
	}
	_, err7 := o.Raw("SELECT COUNT(*)/COUNT(DISTINCT sid) as count FROM surfer;").Values(&school_maps)
	if err7 != nil {
		fmt.Println("get surfer err!", err.Error())
	}
	_, err77 := o.Raw("SELECT COUNT(*)/1 as count FROM surfer;").Values(&total_maps)
	if err77 != nil {
		fmt.Println("get style score err!", err.Error())
	}
	filters = append(filters, "surfer", excellent_maps, school_maps, total_maps)
	//纪律
	_, err8 := o.Raw("SELECT sid,sname,COUNT(*)/1 as count FROM discipline GROUP BY sname ORDER BY count(*) LIMIT ?;", count).Values(&excellent_maps)
	if err8 != nil {
		fmt.Println("get discipline err!", err.Error())
	}
	_, err9 := o.Raw("SELECT COUNT(*)/COUNT(DISTINCT sid) as count FROM discipline;").Values(&school_maps)
	if err9 != nil {
		fmt.Println("get discipline err!", err.Error())
	}
	_, err99 := o.Raw("SELECT COUNT(*)/1 as count FROM discipline;").Values(&total_maps)
	if err99 != nil {
		fmt.Println("get style score err!", err.Error())
	}
	filters = append(filters, "discipline", excellent_maps, school_maps, total_maps)
	//消费
	_, err10 := o.Raw("SELECT sid,sname,SUM(money) as count FROM consume GROUP BY sname ORDER BY SUM(money) DESC LIMIT ?;", count).Values(&excellent_maps)
	if err10 != nil {
		fmt.Println("get consume err!", err.Error())
	}
	_, err11 := o.Raw("SELECT AVG(avg) as count FROM (SELECT sname,SUM(money) as avg FROM consume GROUP BY sname) t1;").Values(&school_maps)
	if err11 != nil {
		fmt.Println("get consume err!", err.Error())
	}
	_, err12 := o.Raw("SELECT sum(money) as count FROM consume;").Values(&total_maps)
	if err12 != nil {
		fmt.Println("get style score err!", err.Error())
	}
	filters = append(filters, "consume", excellent_maps, school_maps, total_maps)
	l := len(filters)
	for k := 0; k < l; k += 4 {
		out := make(map[string]interface{})
		//out1 := make(map[string]interface{})
		out["style"] = filters[k]
		out["excellent_student"] = filters[k+1]
		out["shcool_student"] = filters[k+2]
		out["total"] = filters[k+3]
		//out1[filters[k]] = out
		reslut1 = append(reslut1, out)
	}
	this.ajaxList("获取信息成功", MSG_OK, int64(len(reslut1)), reslut1)
}

func (this *GrowController) GrowthRecord() {
	fmt.Println("获取成长记录")
	o := orm.NewOrm()
	filters := make([]interface{}, 0)
	var school_maps []orm.Params
	var my_maps []orm.Params
	var total_maps []orm.Params
	var reslut1 []map[string]interface{}
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
	sid := "001"
	//成绩
	_, _err := o.Raw("SELECT sid,sname,SUM(grade) as sum FROM score WHERE sid = ?", sid).Values(&my_maps)
	if _err != nil {
		fmt.Println("get style score err!", err.Error())
	}
	_, err1 := o.Raw("SELECT AVG(avg) as count FROM (SELECT sname,SUM(grade) as avg FROM score GROUP BY sname) t1;").Values(&school_maps)
	if err1 != nil {
		fmt.Println("get style score err!", err.Error())
	}
	_, err111 := o.Raw("SELECT COUNT(DISTINCT subject)*100 as count FROM score;").Values(&total_maps)
	if err111 != nil {
		fmt.Println("get style score err!", err.Error())
	}
	filters = append(filters, "score", my_maps, school_maps, total_maps)
	//阅读
	_, err2 := o.Raw("SELECT sid,sname,COUNT(*)/1 as count FROM monitor.read WHERE sid = ?", sid).Values(&my_maps)
	if err2 != nil {
		fmt.Println("get style score err!", err.Error())
	}
	_, err3 := o.Raw("SELECT COUNT(*)/COUNT(DISTINCT sid) as count FROM monitor.read;").Values(&school_maps)
	if err3 != nil {
		fmt.Println("get style score err!", err.Error())
	}
	_, err33 := o.Raw("SELECT COUNT(*)/1 as count FROM monitor.read;").Values(&total_maps)
	if err33 != nil {
		fmt.Println("get style score err!", err.Error())
	}
	filters = append(filters, "read", my_maps, school_maps, total_maps)
	//科研
	_, err4 := o.Raw("SELECT sid,sname,COUNT(*)/1 as count FROM research WHERE sid = ?", sid).Values(&my_maps)
	if err4 != nil {
		fmt.Println("get research err!", err.Error())
	}
	_, err5 := o.Raw("SELECT COUNT(*)/COUNT(DISTINCT sid) as count FROM research;").Values(&school_maps)
	if err5 != nil {
		fmt.Println("get research err!", err.Error())
	}
	_, err55 := o.Raw("SELECT COUNT(*)/1 as count FROM research;").Values(&total_maps)
	if err55 != nil {
		fmt.Println("get style score err!", err.Error())
	}
	filters = append(filters, "research", my_maps, school_maps, total_maps)
	//上网
	_, err6 := o.Raw("SELECT sid,sname,COUNT(*)/1 as count FROM surfer WHERE sid = ?", sid).Values(&my_maps)
	if err6 != nil {
		fmt.Println("get surfer err!", err.Error())
	}
	_, err7 := o.Raw("SELECT COUNT(*)/COUNT(DISTINCT sid) as count FROM surfer;").Values(&school_maps)
	if err7 != nil {
		fmt.Println("get surfer err!", err.Error())
	}
	_, err77 := o.Raw("SELECT COUNT(*)/1 as count FROM surfer;").Values(&total_maps)
	if err77 != nil {
		fmt.Println("get style score err!", err.Error())
	}
	filters = append(filters, "surfer", my_maps, school_maps, total_maps)
	//纪律
	_, err8 := o.Raw("SELECT sid,sname,COUNT(*)/1 as count FROM discipline WHERE sid = ?", sid).Values(&my_maps)
	if err8 != nil {
		fmt.Println("get discipline err!", err.Error())
	}
	_, err9 := o.Raw("SELECT COUNT(*)/COUNT(DISTINCT sid) as count FROM discipline;").Values(&school_maps)
	if err9 != nil {
		fmt.Println("get discipline err!", err.Error())
	}
	_, err99 := o.Raw("SELECT COUNT(*)/1 as count FROM discipline;").Values(&total_maps)
	if err99 != nil {
		fmt.Println("get style score err!", err.Error())
	}
	filters = append(filters, "discipline", my_maps, school_maps, total_maps)
	//消费
	_, err10 := o.Raw("SELECT sid,sname,SUM(money) as count FROM consume WHERE sid = ?", sid).Values(&my_maps)
	if err10 != nil {
		fmt.Println("get consume err!", err.Error())
	}
	_, err11 := o.Raw("SELECT AVG(avg) as count FROM (SELECT sname,SUM(money) as avg FROM consume GROUP BY sname) t1;").Values(&school_maps)
	if err11 != nil {
		fmt.Println("get consume err!", err.Error())
	}
	_, err12 := o.Raw("SELECT sum(money) as count FROM consume;").Values(&total_maps)
	if err12 != nil {
		fmt.Println("get style score err!", err.Error())
	}
	filters = append(filters, "consume", my_maps, school_maps, total_maps)
	//fmt.Println("filters", filters)
	l := len(filters)
	for k := 0; k < l; k += 4 {
		out := make(map[string]interface{})
		//out1 := make(map[string]interface{})
		out["style"] = filters[k]
		out["mine"] = filters[k+1]
		out["shcool_student"] = filters[k+2]
		out["total"] = filters[k+3]
		//out1[filters[k]] = out
		reslut1 = append(reslut1, out)
	}
	//	fmt.Println("result1", reslut1)
	this.ajaxList("获取信息成功", MSG_OK, int64(len(reslut1)), reslut1)
}

func (this *GrowController) GrowthTargets() {
	fmt.Printf("成长目标")
	o := orm.NewOrm()
	out := make(map[string]interface{})
	var maps []orm.Params
	var my_maps []orm.Params
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
	count, err := this.GetInt("count")
	if err != nil {
		this.ajaxMsg("count为int类型", MSG_ERR_Param)
	}
	if count == 0 {
		this.ajaxMsg("count不能为空", MSG_ERR_Param)
	}
	switch style {
	case "成绩":
		_, err := o.Raw("SELECT sid,sname,SUM(CASE subject WHEN '高等数学' THEN grade ELSE 0 END) as math , SUM(CASE subject WHEN '体育' THEN grade ELSE 0 END) as pe , SUM(CASE subject WHEN '大学英语' THEN grade ELSE 0 END) as en , SUM(CASE subject WHEN '中国文学' THEN grade ELSE 0 END) as history , SUM(CASE subject WHEN '计算机基础' THEN grade ELSE 0 END) as computer , SUM(CASE subject WHEN '心理学' THEN grade ELSE 0 END) as psychology , SUM(grade) as sum FROM score GROUP BY sname ORDER BY SUM(grade) DESC LIMIT ?", count).Values(&maps)
		if err != nil {
			fmt.Println("get style score err!", err.Error())
		}
		_, err1 := o.Raw("SELECT sid,sname,SUM(CASE subject WHEN '高等数学' THEN grade ELSE 0 END) as math , SUM(CASE subject WHEN '体育' THEN grade ELSE 0 END) as pe , SUM(CASE subject WHEN '大学英语' THEN grade ELSE 0 END) as en , SUM(CASE subject WHEN '中国文学' THEN grade ELSE 0 END) as history , SUM(CASE subject WHEN '计算机基础' THEN grade ELSE 0 END) as computer , SUM(CASE subject WHEN '心理学' THEN grade ELSE 0 END) as psychology , SUM(grade) as sum FROM score WHERE sid='001';").Values(&my_maps)
		if err1 != nil {
			fmt.Println("get style score err!", err.Error())
		}
	case "阅读":
		_, err := o.Raw("SELECT sid,sname,COUNT(*) as count FROM read GROUP BY sname ORDER BY count(*) DESC LIMIT ?;", count).Values(&maps)
		if err != nil {
			fmt.Println("get style score err!", err.Error())
		}
		_, err1 := o.Raw("SELECT sid,sname,COUNT(*) as count FROM read WHERE sid='001';").Values(&my_maps)
		if err1 != nil {
			fmt.Println("get style score err!", err.Error())
		}
	case "科研":
		_, err := o.Raw("SELECT sid,sname,COUNT(*) as count FROM research GROUP BY sname ORDER BY count(*) DESC LIMIT ?;", count).Values(&maps)
		if err != nil {
			fmt.Println("get style score err!", err.Error())
		}
		_, err1 := o.Raw("SELECT sid,sname,COUNT(*) as count FROM research WHERE sid='001';").Values(&my_maps)
		if err1 != nil {
			fmt.Println("get style score err!", err.Error())
		}
	case "上网":
		_, err := o.Raw("SELECT sid,sname, SUM(duration) as duration FROM surfer GROUP BY sname ORDER BY SUM(duration) LIMIT ?;", count).Values(&maps)
		if err != nil {
			fmt.Println("get style score err!", err.Error())
		}
		_, err1 := o.Raw("SELECT sid,sname, SUM(duration) as duration FROM surfer WHERE sid='001';").Values(&my_maps)
		if err1 != nil {
			fmt.Println("get style score err!", err.Error())
		}
	case "纪律":
		_, err := o.Raw("SELECT sid,sname,COUNT(*) as count FROM discipline GROUP BY sname ORDER BY count(*) LIMIT ?;", count).Values(&maps)
		if err != nil {
			fmt.Println("get style score err!", err.Error())
		}
		_, err1 := o.Raw("SELECT sid,sname,COUNT(*) as count FROM discipline WHERE sid='001';").Values(&my_maps)
		if err1 != nil {
			fmt.Println("get style score err!", err.Error())
		}
	case "消费":
		_, err := o.Raw("SELECT sid,sname, SUM(money) as sum FROM consume GROUP BY sname ORDER BY SUM(money) LIMIT ?;", count).Values(&maps)
		if err != nil {
			fmt.Println("get style score err!", err.Error())
		}
		_, err1 := o.Raw("SELECT sid,sname, SUM(money) as sum FROM consume WHERE sid='001';").Values(&my_maps)
		if err1 != nil {
			fmt.Println("get style score err!", err.Error())
		}
	default:
		this.ajaxMsg("请输入正确的类型", MSG_ERR_Param)
	}
	out["excellent_students"] = maps
	out["mine"] = my_maps
	this.ajaxList("获取信息成功", MSG_OK, int64(len(maps)), out)
}
func (this *GrowController) StudentInformation() {
	fmt.Println("获取学生信息")
	o := orm.NewOrm()
	var my_maps []orm.Params
	//	var reslut []map[string]interface{}
	out := make(map[string]interface{})
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
	sid := "001"
	//学生信息
	_, err11 := o.Raw("SELECT *  FROM student WHERE sid = ?", sid).Values(&my_maps)
	if err11 != nil {
		fmt.Println("get style score err!", err.Error())
	}
	out["student_info"] = my_maps
	//out["sid"] = sid
	//成绩
	_, _err := o.Raw("SELECT SUM(CASE subject WHEN '高等数学' THEN grade ELSE 0 END) as math , SUM(CASE subject WHEN '体育' THEN grade ELSE 0 END) as pe , SUM(CASE subject WHEN '大学英语' THEN grade ELSE 0 END) as en , SUM(CASE subject WHEN '中国文学' THEN grade ELSE 0 END) as history , SUM(CASE subject WHEN '计算机基础' THEN grade ELSE 0 END) as computer , SUM(CASE subject WHEN '心理学' THEN grade ELSE 0 END) as psychology , SUM(grade) as sum FROM score WHERE sid = ?", sid).Values(&my_maps)
	if _err != nil {
		fmt.Println("get style score err!", err.Error())
	}
	out["score"] = my_maps
	//阅读
	_, err2 := o.Raw("SELECT book_type, COUNT(*) as count FROM monitor.read WHERE sid = ? GROUP BY book_type DESC", sid).Values(&my_maps)
	if err2 != nil {
		fmt.Println("get style score err!", err.Error())
	}
	out["read"] = my_maps
	//科研
	_, err4 := o.Raw("SELECT project_level, COUNT(*) as count FROM research WHERE sid = ? GROUP BY project_level DESC", sid).Values(&my_maps)
	if err4 != nil {
		fmt.Println("get research err!", err.Error())
	}
	out["research"] = my_maps
	//上网
	_, err6 := o.Raw("SELECT  AVG(duration) as duration,COUNT(*) as count FROM surfer WHERE sid = ?", sid).Values(&my_maps)
	if err6 != nil {
		fmt.Println("get surfer err!", err.Error())
	}
	out["surfer"] = my_maps
	//纪律
	_, err8 := o.Raw("SELECT style,COUNT(*) as count FROM discipline WHERE sid = ? GROUP BY style DESC", sid).Values(&my_maps)
	if err8 != nil {
		fmt.Println("get discipline err!", err.Error())
	}
	out["discipline"] = my_maps
	//消费
	_, err10 := o.Raw("SELECT SUM(money) as count FROM consume WHERE sid = ?", sid).Values(&my_maps)
	if err10 != nil {
		fmt.Println("get consume err!", err.Error())
	}
	out["consume"] = my_maps
	this.ajaxList("获取学生信息成功", MSG_OK, int64(len(out)), out)
}
