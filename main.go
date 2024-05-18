// 启动程序
// @author zxl
package main

import (
	"go-David/core"
	"go-David/global"
)

func main() {
	//fmt.Println("gin单体项目-中央监护系统")
	//fmt.Println("系统配置文件：", config.Config.System)
	//fmt.Println("logger日志：", config.Config.Logger)

	//初始化logger
	global.Log = core.InitLogger()
	//初始化mysql
	core.MysqlInit()
}
