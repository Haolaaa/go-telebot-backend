package api

import (
	"admin/global"
	"admin/model"
	"admin/model/common/request"
	"admin/model/common/response"
	"admin/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SiteConfigApi struct{}

func (s *SiteConfigApi) GetSiteConfigList(c *gin.Context) {
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
	siteConfigs, total, err := siteConfigService.GetSiteConfigList(pageInfo)
	if err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(
		response.PageResult{
			List:     siteConfigs,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		},
		"获取成功",
		c,
	)
}

func (s *SiteConfigApi) AddSiteConfig(c *gin.Context) {
	var siteConfig model.SysSiteConfig
	err := c.ShouldBindJSON(&siteConfig)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(siteConfig, utils.SiteConfigVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = siteConfigService.AddSiteConfig(siteConfig)
	if err != nil {
		global.LOG.Error("添加失败!", zap.Error(err))
		response.FailWithMessage("添加失败", c)
		return
	}
	response.OkWithMessage("添加成功", c)
}

func (s *SiteConfigApi) DeleteSiteConfig(c *gin.Context) {
	var siteConfig request.GetByID
	err := c.ShouldBindJSON(&siteConfig)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(siteConfig, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = siteConfigService.DeleteSiteConfig(siteConfig.ID)
	if err != nil {
		global.LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

func (s *SiteConfigApi) UpdateSiteConfig(c *gin.Context) {
	var siteConfig model.SysSiteConfig
	err := c.ShouldBindJSON(&siteConfig)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(siteConfig, utils.SiteConfigVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = siteConfigService.UpdateSiteConfig(siteConfig)
	if err != nil {
		global.LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}
