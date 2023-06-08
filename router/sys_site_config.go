package router

import (
	"admin/api"

	"github.com/gin-gonic/gin"
)

type SiteConfigRouter struct{}

func (s *SiteConfigRouter) InitSiteConfigRouter(Router *gin.RouterGroup) {
	siteConfigRouter := Router.Group("siteConfig")
	baseApi := api.ApiGroupApp.SiteConfigApi
	{
		siteConfigRouter.POST("siteConfigList", baseApi.GetSiteConfigList)  // 获取站点配置列表
		siteConfigRouter.POST("addSiteConfig", baseApi.AddSiteConfig)       // 新增站点配置
		siteConfigRouter.POST("deleteSiteConfig", baseApi.DeleteSiteConfig) // 删除站点配置
		siteConfigRouter.POST("SetSiteConfig", baseApi.UpdateSiteConfig)    // 设置站点配置
		siteConfigRouter.POST("getByID", baseApi.GetSiteConfigByID)         // 根据ID获取站点配置
	}
}
