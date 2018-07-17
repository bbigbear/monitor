package controllers

import (
	"campusMonitorSysterm/models"
	"encoding/json"
	"fmt"
	//	"strconv"
	//	"math"

	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type LoginController struct {
	BaseController
}

func (this *LoginController) Login() {
	fmt.Println("点击登录")
	o := orm.NewOrm()
	l := new(models.Login)
	var user_info models.Login
	json.Unmarshal(this.Ctx.Input.RequestBody, &user_info)
	fmt.Println("user_info:", &user_info)
	n := user_info.Name
	p := user_info.Pwd
	query := o.QueryTable(l)
	exist := query.Filter("Name", n).Exist()
	if !exist {
		fmt.Println("用户不存在")
		this.ajaxMsg("用户不存在", MSG_ERR_Param)
	}
	err := query.Filter("Name", n).One(&user_info)
	if err != nil {
		this.ajaxMsg("登录失败", MSG_ERR_Resources)
	}
	pwd := user_info.Pwd
	if p != pwd {
		this.ajaxMsg("密码错误", MSG_ERR_Verified)
	}
	fmt.Println("用户名、密码正确")
	appkey := beego.AppConfig.String("appkey")
	//授权
	token, i := this.Create_token(n, appkey)
	fmt.Println("token&time", token, i)

	list := make(map[string]interface{})
	list["name"] = n
	list["token"] = token
	list["time"] = i
	list["pid"] = user_info.Pid
	this.ajaxList("登录成功", MSG_OK, 1, list)
}
