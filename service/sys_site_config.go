package service

import (
	"admin/global"
	"admin/model"
	"admin/model/common/request"
	"errors"
	"time"

	"gorm.io/gorm"
)

type SiteConfigService struct{}

func (s *SiteConfigService) GetSiteConfigList(info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.DB.Model(&model.SysSiteConfig{})
	var siteConfigs []model.SysSiteConfig
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&siteConfigs).Error
	return siteConfigs, total, err
}

func (s *SiteConfigService) AddSiteConfig(param model.SysSiteConfig) (err error) {
	var siteConfig model.SysSiteConfig
	if !errors.Is(global.DB.Where("site_name = ?", param.SiteName).First(&siteConfig).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同站点名称")
	}
	err = global.DB.Create(&param).Error
	return err
}

func (s *SiteConfigService) DeleteSiteConfig(id int) (err error) {
	var siteConfig model.SysSiteConfig
	err = global.DB.Where("id = ?", id).Delete(&siteConfig).Error

	return
}

func (s *SiteConfigService) UpdateSiteConfig(req model.SysSiteConfig) error {
	return global.DB.Model(&model.SysSiteConfig{}).
		Select("updated_at", "parent_name", "site_name", "site_key", "site_id", "direct_play_url", "cf_play_url", "cdn_play_url", "video_cover", "download_url").
		Updates(map[string]interface{}{
			"updated_at":      time.Now(),
			"parent_name":     req.ParentName,
			"site_name":       req.SiteName,
			"site_key":        req.SiteKey,
			"site_id":         req.SiteID,
			"direct_play_url": req.DirectPlayUrl,
			"cf_play_url":     req.CFPlayUrl,
			"cdn_play_url":    req.CDNPlayUrl,
			"video_cover":     req.VideoCover,
			"download_url":    req.DownloadUrl,
		}).
		Error
}
