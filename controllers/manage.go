package controllers

import (
	"campusMonitorSysterm/models"
	"fmt"
	"math"
	"time"

	"encoding/json"
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
	o := orm.NewOrm()
	setting := new(models.WarnSetting)
	var maps []orm.Params
	_, err := o.QueryTable(setting).Filter("Id", 1).Values(&maps)
	if err != nil {
		fmt.Println("get setting err", err.Error())
	}
	this.Data["m"] = maps
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

//设置预警阀值

func (this *ManageController) UpdateSetting() {
	fmt.Println("更新预警设置")
	o := orm.NewOrm()
	//list := make(map[string]interface{})
	var setting models.WarnSetting
	json.Unmarshal(this.Ctx.Input.RequestBody, &setting)
	fmt.Println("setting_info:", &setting)
	setting.Id = 1
	style := setting.Style
	switch style {
	case "设置预警通知方式":
		_, err := o.Update(&setting, "Tzfs")
		if err != nil {
			this.ajaxMsg("更新失败", MSG_ERR_Resources)
		}
	case "学业预警":
		_, err := o.Update(&setting, "XyyjZcj", "XyyjDkcj")
		if err != nil {
			this.ajaxMsg("更新失败", MSG_ERR_Resources)
		}
	case "一卡通消费预警":
		_, err := o.Update(&setting, "YktDbxf", "YktDrxf", "YktDyxf")
		if err != nil {
			this.ajaxMsg("更新失败", MSG_ERR_Resources)
		}
	case "上网预警":
		_, err := o.Update(&setting, "SwyjCs", "SwyjZsc", "SwyjDcswsc")
		if err != nil {
			this.ajaxMsg("更新失败", MSG_ERR_Resources)
		}
	case "图书借阅预警":
		_, err := o.Update(&setting, "TsjySl", "TsjySj")
		if err != nil {
			this.ajaxMsg("更新失败", MSG_ERR_Resources)
		}
	case "贫困生预警":
		_, err := o.Update(&setting, "PksYxf")
		if err != nil {
			this.ajaxMsg("更新失败", MSG_ERR_Resources)
		}
	case "图书馆爆满预警":
		_, err := o.Update(&setting, "TsgbmRs")
		if err != nil {
			this.ajaxMsg("更新失败", MSG_ERR_Resources)
		}
	case "挂科预警":
		_, err := o.Update(&setting, "Gkcs")
		if err != nil {
			this.ajaxMsg("更新失败", MSG_ERR_Resources)
		}
	default:
		this.ajaxMsg("请输入正确的类型", MSG_ERR_Param)
	}
	this.ajaxMsg("更新成功", MSG_OK)
}

//更新预警状态

func (this *ManageController) ChangeStatus() {
	fmt.Println("更新预警状态")
	//id
	id, err := this.GetInt("id")
	if err != nil {
		this.ajaxMsg("get id err", MSG_ERR_Param)
	}
	fmt.Println("id:", id)
	//status
	status := this.GetString("status")
	fmt.Println("status is", status)
	o := orm.NewOrm()
	warn := new(models.Warn)
	//updata status db
	num, err := o.QueryTable(warn).Filter("Id", id).Update(orm.Params{
		"Status": status,
	})
	if err != nil {
		fmt.Println("change status err", err.Error())
		this.ajaxMsg("change status err", MSG_ERR_Resources)
	}
	fmt.Println("num", num)
	this.ajaxMsg("处理成功", MSG_OK)
}

//更新预警状态

func (this *ManageController) ChangeRemark() {
	fmt.Println("更新预警")
	//id
	id, err := this.GetInt("id")
	if err != nil {
		this.ajaxMsg("get id err", MSG_ERR_Param)
	}
	fmt.Println("id:", id)
	//remark
	remark := this.GetString("remark")
	fmt.Println("remark is", remark)
	o := orm.NewOrm()
	warn := new(models.Warn)
	//updata status db
	num, err := o.QueryTable(warn).Filter("Id", id).Update(orm.Params{
		"Remark": remark,
	})
	if err != nil {
		fmt.Println("change status err", err.Error())
		this.ajaxMsg("change status err", MSG_ERR_Resources)
	}
	fmt.Println("num", num)
	this.ajaxMsg("处理成功", MSG_OK)
}
