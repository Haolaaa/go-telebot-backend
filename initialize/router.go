package initialize

import (
	"admin/docs"
	"admin/global"
	"admin/middleware"
	"admin/router"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	Router.Use(middleware.Cors()) // 跨域

	docs.SwaggerInfo.BasePath = "/"
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.LOG.Info("register swagger handler")

	PublicGroup := Router.Group("")
	{
		PublicGroup.GET("/health", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, "ok")
		})
	}
	{
		router.RouterGroupApp.InitBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权
	}
	PrivateGroup := Router.Group("")
	// PrivateGroup.Use(middleware.Cors()).Use(middleware.JWTAuth())
	{
		router.RouterGroupApp.InitAuthorityRouter(PrivateGroup)
		router.RouterGroupApp.InitMenuRouter(PrivateGroup)
		router.RouterGroupApp.InitUserRouter(PrivateGroup)
		router.RouterGroupApp.InitSiteConfigRouter(PrivateGroup)
	}

	global.LOG.Info("router register success")
	return Router
}
