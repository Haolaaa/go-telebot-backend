package router

type RouterGroup struct {
	AuthorityRouter
	BaseRouter
	MenuRouter
	UserRouter
	SiteConfigRouter
}

var RouterGroupApp = new(RouterGroup)
