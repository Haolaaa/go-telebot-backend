package service

import (
	"admin/global"
	"admin/model"
	"admin/model/common/request"
	"errors"

	"gorm.io/gorm"
)

var ErrRoleExists = errors.New("角色已存在")

type AuthorityService struct{}

var AuthorityServiceApp = new(AuthorityService)

func (a *AuthorityService) CreateAuthority(auth model.SysAuthority) (authority model.SysAuthority, err error) {
	var authorityBox model.SysAuthority
	if !errors.Is(global.DB.Where("authority_id = ?", auth.AuthorityID).First(&authorityBox).Error, gorm.ErrRecordNotFound) {
		return auth, ErrRoleExists
	}

	err = global.DB.Create(&auth).Error
	return auth, err
}

func (a *AuthorityService) UpdateAuthority(auth model.SysAuthority) (authority model.SysAuthority, err error) {
	err = global.DB.Where("authority_id = ?", auth.AuthorityID).First(&model.SysAuthority{}).Updates(&auth).Error
	return auth, err
}

func (a *AuthorityService) DeleteAuthority(auth *model.SysAuthority) (err error) {
	if errors.Is(global.DB.Debug().Preload("Users").First(&auth).Error, gorm.ErrRecordNotFound) {
		return errors.New("角色不存在")
	}
	if len(auth.Users) != 0 {
		return errors.New("此角色有用户正在使用禁止删除")
	}
	if !errors.Is(global.DB.Where("authority_id = ?", auth.AuthorityID).First(&model.SysUser{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("此角色有用户正在使用禁止删除")
	}
	db := global.DB.Preload("SysBaseMenus").Where("authority_id = ?", auth.AuthorityID).First(auth)
	err = db.Unscoped().Delete(auth).Error
	if err != nil {
		return
	}

	if len(auth.SysBaseMenus) > 0 {
		err = global.DB.Model(auth).Association("SysBaseMenus").Delete(auth.SysBaseMenus)
		if err != nil {
			return
		}
	}

	err = global.DB.Delete(&[]model.SysUserAuthority{}, "sys_authority_authority_id = ?", auth.AuthorityID).Error
	if err != nil {
		return
	}
	err = global.DB.Delete(&[]model.SysAuthorityBtn{}, "authority_id = ?", auth.AuthorityID).Error
	if err != nil {
		return
	}

	return err
}

func (a *AuthorityService) GetAuthorityInfoList(info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.DB.Model(&model.SysAuthority{})
	if err = db.Where("parent_id = ?", "0").Count(&total).Error; total == 0 || err != nil {
		return
	}
	var authority []model.SysAuthority
	err = db.Limit(limit).Offset(offset).Where("parent_id = ?", "0").Find(&authority).Error

	return authority, total, err
}

func (a *AuthorityService) SetMenuAuthority(auth *model.SysAuthority) error {
	var s model.SysAuthority
	global.DB.Preload("SysBaseMenus").First(&s, "authority_id = ?", auth.AuthorityID)
	err := global.DB.Model(&s).Association("SysBaseMenus").Replace(&auth.SysBaseMenus)
	return err
}
