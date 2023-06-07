package api

import (
	"admin/global"
	"admin/model"
	"admin/model/common/request"
	"admin/model/common/response"
	menuReq "admin/model/request"
	menuRes "admin/model/response"
	"admin/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthorityMenuApi struct{}

func (a *AuthorityMenuApi) GetMenu(c *gin.Context) {
	menus, err := menuService.GetMenuTree(utils.GetUserAuthorityId(c))
	if err != nil {
		global.LOG.Error("error getting menu list", zap.Error(err))
		response.FailWithMessage("error getting menu list", c)
		return
	}
	if menus == nil {
		menus = []model.SysMenu{}
	}
	response.OkWithDetailed(menuRes.SysMenuResponse{Menus: menus}, "success", c)
}

func (a *AuthorityMenuApi) GetBaseMenuTree(c *gin.Context) {
	menus, err := menuService.GetBaseMenuTree()
	if err != nil {
		global.LOG.Error("error getting menu tree", zap.Error(err))
		response.FailWithMessage("error getting menu tree", c)
		return
	}
	response.OkWithDetailed(menuRes.SysBaseMenusResponse{Menus: menus}, "success", c)
}

func (a *AuthorityMenuApi) AddMenuAuthority(c *gin.Context) {
	var authorityMenu menuReq.AddMenuAuthorityInfo
	err := c.ShouldBindJSON(&authorityMenu)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := utils.Verify(authorityMenu, utils.AuthorityIdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := menuService.AddMenuAuthority(authorityMenu.Menus, authorityMenu.AuthorityID); err != nil {
		global.LOG.Error("failed to add menu auth", zap.Error(err))
		response.FailWithMessage("failed to add menu auth", c)
		return
	}
	response.OkWithMessage("success", c)
}

func (a *AuthorityMenuApi) GetMenuAuthority(c *gin.Context) {
	var param request.GetAuthorityID
	err := c.ShouldBindJSON(&param)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(param, utils.AuthorityIdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	menus, err := menuService.GetMenuAuthority(&param)
	if err != nil {
		global.LOG.Error("failed to get menu auth", zap.Error(err))
		response.FailWithDetailed(menuRes.SysMenuResponse{Menus: menus}, "failed operation", c)
		return
	}
	response.OkWithDetailed(gin.H{"menus": menus}, "success", c)
}

func (a *AuthorityMenuApi) AddBaseMenu(c *gin.Context) {
	var menu model.SysBaseMenu
	err := c.ShouldBindJSON(&menu)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(menu, utils.MenuMetaVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = menuService.AddBaseMenu(menu)
	if err != nil {
		global.LOG.Error("failed to add new base menu", zap.Error(err))
		response.FailWithMessage("failed operation", c)
		return
	}
	response.OkWithMessage("success", c)
}

func (a *AuthorityMenuApi) DeleteBaseMenu(c *gin.Context) {
	var menu request.GetByID
	err := c.ShouldBindJSON(&menu)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(menu, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = baseMenuService.DeleteBaseMenu(menu.ID)
	if err != nil {
		global.LOG.Error("failed to delete menu", zap.Error(err))
		response.FailWithMessage("failed operation", c)
		return
	}
	response.OkWithMessage("success", c)
}

func (a *AuthorityMenuApi) UpdateBaseMenu(c *gin.Context) {
	var menu model.SysBaseMenu
	err := c.ShouldBindJSON(&menu)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(menu, utils.MenuMetaVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = baseMenuService.UpdateBaseMenu(menu)
	if err != nil {
		global.LOG.Error("failed to update", zap.Error(err))
		response.FailWithMessage("failed operation", c)
		return
	}
	response.OkWithMessage("success", c)
}

func (a *AuthorityMenuApi) GetBaseMenuById(c *gin.Context) {
	var idInfo request.GetByID
	err := c.ShouldBindJSON(&idInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(idInfo, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	menu, err := baseMenuService.GetBaseMenuByID(idInfo.ID)
	if err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(menuRes.SysBaseMenuResponse{Menu: menu}, "获取成功", c)
}

func (a *AuthorityMenuApi) GetMenuList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(pageInfo, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	menuList, total, err := menuService.GetInfoList()
	if err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     menuList,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}
