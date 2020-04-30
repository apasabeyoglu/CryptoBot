package main

import (
	"cryptobot/helpers"
	_ "cryptobot/routers"
	"cryptobot/tasks"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	defaultdb := beego.AppConfig.String("defaultdb")
	timezone := beego.AppConfig.String("timezone")
	enableDBDebug, _ := beego.AppConfig.Bool("debugdbqueries")
	location, _ := time.LoadLocation(timezone)
	orm.DefaultTimeLoc = location
	orm.Debug = enableDBDebug

	initLogging()

	if defaultdb == "sqlite3" {
		initSqlite3()
	}
	orm.SetDataBaseTZ("default", location)

	initTemplateFuncs()

	tasks.Start()
}

func initLogging() {
	runmode := beego.AppConfig.String("runmode")
	logonprod, _ := beego.AppConfig.Bool("logonprod")

	beego.SetLogFuncCall(true)
	logs.Async()

	if runmode == "prod" && logonprod {
		logpath := beego.AppConfig.String("logpath")
		maxdays := beego.AppConfig.String("logmaxdays")

		beego.SetLogger(logs.AdapterFile, `{"filename": "`+logpath+`", "maxdays": `+maxdays+`}`)
		beego.SetLevel(beego.LevelInformational)
		beego.BeeLogger.DelLogger("console")
	} else if runmode == "dev" || runmode == "test" {
		beego.SetLevel(beego.LevelDebug)
	}
}

func initSqlite3() {
	sqlitefilepath := beego.AppConfig.String("sqlitefilepath")
	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", sqlitefilepath)
}

func initTemplateFuncs() {
	beego.AddFuncMap("fConvertCheck", helpers.ConvertCheck)
}

func main() {
	beego.Run()
}
