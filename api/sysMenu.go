// Package api 菜单相关接口
// @author:zxl
package api

import (
	"github.com/gin-gonic/gin"
	. "go-David/core"
	"go-David/model"
	"go-David/result"
	"go-David/util"
	"strconv"
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

// GetSysMenuList 查询菜单列表
// @Summary:查询菜单列表
// @Tags 菜单相关接口
// @Produce json
// @Description 查询菜单列表
// @Param menuName query string false "菜单名称"
// @Param menuStatus query string false "菜单状态"
// @Success 200 {object} result.Result
// @router /api/sysMenu/list [get]
func GetSysMenuList(c *gin.Context) {
	// 从请求上下文中获取 "menuName" 查询参数
	MenuName := c.Query("menuName")
	// 从请求上下文中获取 "menuStatus" 查询参数
	MenuStatus := c.Query("menuStatus")
	// 定义一个切片来存放查询结果
	var sysMenu []model.SysMenu
	// 初始化数据库查询，并按 "sort" 字段排序
	curDb := Db.Table("sys_menu").Order("sort")
	// 如果 "menuName" 参数不为空，添加相应的查询条件
	if MenuName != "" {
		curDb = curDb.Where("menu_name = ?", MenuName)
	}
	// 如果 "menuStatus" 参数不为空，添加相应的查询条件
	if MenuStatus != "" {
		curDb = curDb.Where("menu_status = ?", MenuStatus)
	}
	// 执行查询并将结果存储到 sysMenu 切片中
	curDb.Find(&sysMenu)
	// 返回查询结果
	result.Success(c, sysMenu)
}

// GetSysMenu 根据id查询菜单
// @Summary:根据id查询菜单
// @Tags 菜单相关接口
// @Produce json
// @Description 根据id查询菜单
// @Param id query int true "菜单id"
// @Success 200 {object} result.Result
// @router /api/sysMenu/info [get]
func GetSysMenu(c *gin.Context) {
	// 从请求上下文中获取 "id" 查询参数，并将其转换为整数类型
	Id, _ := strconv.Atoi(c.Query("id"))
	// 定义一个 SysMenu 结构体变量来存放查询结果
	var sysMenu model.SysMenu
	// 使用 GORM 的 First 方法根据 ID 查询 sys_menu 表中的第一条记录
	Db.First(&sysMenu, Id)
	// 返回查询结果，使用 result.Success 方法将结果发送到客户端
	result.Success(c, sysMenu)
}

// UpdateSysMenu 修改菜单
// @Summary:修改菜单
// @Tags 菜单相关接口
// @Produce json
// @Description 修改菜单
// @Param data body model.UpdateSysMenuDto true "data"
// @Success 200 {object} result.Result
// @router /api/sysMenu/update [put]
func UpdateSysMenu(c *gin.Context) {
	var dto model.UpdateSysMenuDto
	_ = c.BindJSON(&dto)
	var sysMenu model.SysMenu
	Db.First(&sysMenu, dto.ID)
	sysMenu.ParentId = dto.ParentId
	sysMenu.MenuName = dto.MenuName
	sysMenu.Icon = dto.Icon
	sysMenu.Value = dto.Value
	sysMenu.MenuType = dto.MenuType
	sysMenu.Url = dto.Url
	sysMenu.MenuStatus = dto.MenuStatus
	sysMenu.Sort = dto.Sort
	Db.Save(&sysMenu)
	result.Success(c, true)
}

// GetSysMenuByName 根据菜单名称查询菜单
func GetSysMenuByName(menuName string) (sysMenu model.SysMenu) {
	Db.Where("menu_name = ?", menuName).First(&sysMenu)
	return sysMenu
}
