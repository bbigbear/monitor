package main

import (
	"campusMonitorSysterm/models"
	_ "campusMonitorSysterm/routers"
	"fmt"

	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	DBConnection()
	RegisterModel()
}

func main() {
	orm.Debug = true
	orm.RunSyncdb("default", false, true)
	beego.Run()
}

func DBConnection() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	host := beego.AppConfig.String("host")
	db := beego.AppConfig.String("database")
	user := beego.AppConfig.String("user")
	passwd := beego.AppConfig.String("passwd")
	maxOpenConns, err := beego.AppConfig.Int("MaxOpenConns")
	if err != nil {
		fmt.Println("MaxOpenConns is nil", err)
	}
	maxIdleConns, err := beego.AppConfig.Int("MaxIdleConns")
	if err != nil {
		fmt.Println("MaxIdleConns is nil", err)
	}

	sql := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", user, passwd, host, db)
	orm.RegisterDataBase("default", "mysql", sql, maxIdleConns, maxOpenConns)
}

func RegisterModel() {
	orm.RegisterModel(new(models.Warn), new(models.Login), new(models.WarnStyle))
}
