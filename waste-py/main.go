package main

import (
	"github.com/yakaa/log4g"
	"waste-py/config"
	"waste-py/router"
)

func main() {
	// 初始化配置
	config.ConfigInit()
	// 初始化服务
	config.InitRrsl()
	// 启用路由
	app := router.NewRouter()
	if config.Conf.App.AppMode == "release" {
		log4g.Error(app.Run(config.Conf.App.AppUrl))
	} else {
		_ = app.Run(config.Conf.App.AppUrl)
	}
}