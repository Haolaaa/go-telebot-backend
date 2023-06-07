package router

import (
	"admin/api"

	"github.com/gin-gonic/gin"
)

type BaseRouter struct{}

func (s *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	baseRouter := Router.Group("base")
	baseApi := api.ApiGroupApp.BaseApi
	{
		baseRouter.POST("login", baseApi.Login)
		baseRouter.POST("adminRegister", baseApi.Register) //管理员注册
	}
	return baseRouter
}
