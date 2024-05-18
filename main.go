// 启动程序
// @author zxl
// 全局安装swag 初始化swag文档使用 swag init
// go install github.com/swaggo/swag/cmd/swag@latest
package main

import (
	"fmt"
	"go-David/config"
	"go-David/core"
	_ "go-David/docs"
	"go-David/global"
	"go-David/router"
)

// @title 中央监护系统
// @version 1.0
// @description 中央监护系统API接口文档
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	//初始化logger
	global.Log = core.InitLogger()
	//初始化mysql
	core.MysqlInit()
	//初始化redis
	core.RedisInit()
	//初始化路由
	router := router.InitRouter()
	address := fmt.Sprintf("%s:%d", config.Config.System.Host, config.Config.System.Port)
	global.Log.Infof("系统启动成功，运行在：%s", address)
	router.Run(address)
}
