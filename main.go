package main

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"rmsServer/route"
	"rmsServer/tools"
)
var cfg *tools.Config
func init() {
	fmt.Println("init config")
	cfg = tools.ParseConfig("./conf/conf.yaml")
	tools.InitOrmEngine(cfg.Mysql)
}
func main() {
	app:=route.InitRouter()
	_ = app.Run(iris.Addr(cfg.App.Host + ":" + cfg.App.Port),iris.WithConfiguration(iris.YAML("./conf/appSet.yaml")))
}
