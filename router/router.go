//初始化路由及注册

package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go-David/api"
	"go-David/config"
)

func InitRouter() *gin.Engine {
	// 启动模式
	gin.SetMode(config.Config.System.Env)
	router := gin.New()
	// 跌机时恢复
	router.Use(gin.Recovery())
	// register 注册
	register(router)
	return router
}

// register 路由接口
func register(router *gin.Engine) {
	//todo 后续接口url
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/api/success", api.Success)
	router.GET("/api/failed", api.Failed)
	router.POST("/api/sysMenu/add", api.CreateSysMenu)

}
