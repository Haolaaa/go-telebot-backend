package service

type ServiceGroup struct {
	MenuService
	UserService
	BaseMenuService
	AuthorityService
	JwtService
	SiteConfigService
}

var ServiceGroupApp = new(ServiceGroup)
