package controllers

import (
	"campusMonitorSysterm/models"
	"fmt"
	"math"
	"time"

	"strconv"
	"strings"

	"github.com/astaxie/beego"

	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/astaxie/beego/orm"
)

type ManageController struct {
	BaseController
}

func (this *ManageController) Get() {
	this.TplName = "login.tpl"
}

func (this *ManageController) WarnMonitor() {
	this.TplName = "warn_monitor.tpl"
}

func (this *ManageController) WarnHistroy() {
	this.TplName = "warn_histroy.tpl"
}

func (this *ManageController) Setting() {
	this.TplName = "warn_setting.tpl"
}

func (this *ManageController) GetWarnData() {

	fmt.Println("获取预警信息")
	o := orm.NewOrm()
	warn := new(models.Warn)
	query := o.QueryTable(warn)
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

	//获取类型
	status := this.Input().Get("status")
	if status != "" {
		query = query.Filter("Status", status)
	}
	//类型
	style := this.Input().Get("style")
	if style != "" {
		query = query.Filter("WarnName", style)
	}
	//index
	index, err := this.GetInt("page")
	if err != nil {
		fmt.Println("下标index不存在")
	}
	//pagemax  一页多少
	pagemax, err := this.GetInt("limit")
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
	for _, m := range maps {
		m["WarnTime"] = m["WarnTime"].(time.Time).Format("2006-01-02 15:04:05")
	}
	this.ajaxList("获取预警信息成功", 0, count, maps)
}

//批量删除
func (this *ManageController) DelMultiData() {
	fmt.Println("删除预警数据")
	o := orm.NewOrm()
	warn := new(models.Warn)
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
	//list := make(map[string]interface{})
	id := this.Input().Get("id")
	fmt.Println("del id:", id)
	idList := strings.Split(id, ",")
	fmt.Println("idList:", idList)
	id_len := len(idList) - 1
	var idIntList []int64
	for i := 0; i < id_len; i++ {
		idd, err := strconv.ParseInt(idList[i], 10, 64)
		if err != nil {
			fmt.Println("delmulti string转int 失败", err.Error())
		}
		idIntList = append(idIntList, idd)
	}
	fmt.Println("idIntList:", idIntList)
	num, err := o.QueryTable(warn).Filter("Id__in", idIntList).Delete()
	if err != nil {
		this.ajaxMsg("删除失败", MSG_ERR_Resources)
	}
	fmt.Println("del multidish reslut num:", num)
	if num == 0 {
		this.ajaxMsg("删除失败", MSG_ERR_Param)
	}
	//list["data"] = maps
	this.ajaxMsg("删除成功", MSG_OK)
	return
}
