// Package api 测试相关接口
// @author: zxl
package api

import (
	"github.com/gin-gonic/gin"
	"go-David/result"
)

// Success 成功测试
// @Summary 成功测试
// @Tags 测试相关接口
// @Produce json
// @Description 成功测试接口
// @Success 200 {object} result.Result
// @router /api/success [get]
func Success(c *gin.Context) {
	result.Success(c, 200)
}

// Failed 失败测试
// @Summary 失败测试
// @Tags 测试相关接口
// @Produce json
// @Description 失败测试接口
// @Success 200 {object} result.Result
// @router /api/failed [get]
func Failed(c *gin.Context) {
	result.Failed(c, int(result.ApiCode.Failed), result.ApiCode.GetMessage(result.ApiCode.Failed))
}
