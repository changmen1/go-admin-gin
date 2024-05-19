// Package api 菜单相关接口
// @author:zxl
package api

import (
	"github.com/gin-gonic/gin"
	. "go-David/core"
	"go-David/model"
	"go-David/result"
	"go-David/util"
	"time"
)

// CreateSysMenu 新增菜单
// @Summary:新增菜单
// @Tags 菜单相关接口
// @Produce json
// @Description 新增菜单
// @Param data body model.AddSysMenuDto true "data"
// @Success 200 {object} result.Result
// @router /api/sysMenu/add [post]
func CreateSysMenu(c *gin.Context) {
	// 绑定参数
	var dto model.AddSysMenuDto
	_ = c.BindJSON(&dto)
	// 查询菜单不能重复
	sysMenuByName := GetSysMenuByName(dto.MenuName)
	if sysMenuByName.ID != 0 {
		result.Failed(c, int(result.ApiCode.SysMenuIsExist), result.ApiCode.GetMessage(result.ApiCode.SysMenuIsExist))
		return
	}
	// 目录
	if dto.MenuType == 1 {
		sysMenu := model.SysMenu{
			ParentId:   dto.ParentId,
			MenuName:   dto.MenuName,
			Icon:       dto.Icon,
			MenuType:   dto.MenuType,
			Url:        dto.Url,
			MenuStatus: dto.MenuStatus,
			Sort:       dto.Sort,
			CreateTime: util.HTime{Time: time.Now()},
		}
		Db.Create(&sysMenu)
	} else if dto.MenuType == 2 {
		sysMenu := model.SysMenu{
			ParentId:   dto.ParentId,
			MenuName:   dto.MenuName,
			Icon:       dto.Icon,
			MenuType:   dto.MenuType,
			Url:        dto.Url,
			Value:      dto.Value,
			MenuStatus: dto.MenuStatus,
			Sort:       dto.Sort,
			CreateTime: util.HTime{Time: time.Now()},
		}
		Db.Create(&sysMenu)
	} else if dto.MenuType == 3 {
		sysMenu := model.SysMenu{
			ParentId:   dto.ParentId,
			MenuName:   dto.MenuName,
			MenuType:   dto.MenuType,
			MenuStatus: dto.MenuStatus,
			Value:      dto.Value,
			Sort:       dto.Sort,
			CreateTime: util.HTime{Time: time.Now()},
		}
		Db.Create(&sysMenu)
	}
	result.Success(c, true)
}

// GetSysMenuByName 根据菜单名称查询菜单
func GetSysMenuByName(menuName string) (sysMenu model.SysMenu) {
	Db.Where("menu_name = ?", menuName).First(&sysMenu)
	return sysMenu
}
