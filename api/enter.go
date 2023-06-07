package api

import "admin/service"

type ApiGroup struct {
	AuthorityApi
	AuthorityMenuApi
	BaseApi
	SiteConfigApi
}

var ApiGroupApp = new(ApiGroup)

var (
	authorityService  = service.ServiceGroupApp.AuthorityService
	baseMenuService   = service.ServiceGroupApp.BaseMenuService
	menuService       = service.ServiceGroupApp.MenuService
	userService       = service.ServiceGroupApp.UserService
	siteConfigService = service.ServiceGroupApp.SiteConfigService
)
