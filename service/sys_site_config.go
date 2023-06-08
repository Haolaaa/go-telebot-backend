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
	return global.DB.Transaction(func(tx *gorm.DB) error {
		var siteConfig model.SysSiteConfig
		err := tx.Where("id = ?", req.ID).First(&siteConfig).Error
		if err != nil {
			return err
		}
		if siteConfig.SiteName != req.SiteName {
			var siteConfig model.SysSiteConfig
			if !errors.Is(tx.Where("site_name = ?", req.SiteName).First(&siteConfig).Error, gorm.ErrRecordNotFound) {
				return errors.New("存在相同站点名称")
			}
		}
		req.UpdatedAt = time.Now()
		err = tx.Where("id = ?", req.ID).Updates(&req).Error
		if err != nil {
			return err
		}
		return nil
	})
}

func (s *SiteConfigService) GetSiteConfigByID(id int, parentName string) (siteConfig model.SysSiteConfig, err error) {
	err = global.DB.Where("id = ? AND parent_name = ?", id, parentName).First(&siteConfig).Error
	return
}
