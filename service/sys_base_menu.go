package service

import (
	"admin/global"
	"admin/model"
	"errors"

	"gorm.io/gorm"
)

type BaseMenuService struct{}

func (b *BaseMenuService) DeleteBaseMenu(id int) (err error) {
	err = global.DB.Preload("MenuBtn").Preload("Parameters").Where("parent_id = ?", id).First(&model.SysBaseMenu{}).Error
	if err != nil {
		var menu model.SysBaseMenu
		db := global.DB.Preload("SysAuthorities").Where("id = ?", id).First(&menu).Delete(&menu)
		err = global.DB.Delete(&model.SysBaseMenuParameter{}, "sys_base_menu_id = ?", id).Error
		err = global.DB.Delete(&model.SysBaseMenuBtn{}, "sys_base_menu_id = ?", id).Error
		err = global.DB.Delete(&model.SysAuthorityBtn{}, "sys_menu_id = ?", id).Error
		if err != nil {
			return err
		}
		if len(menu.SysAuthorities) > 0 {
			err = global.DB.Model(&menu).Association("SysAuthorities").Delete(&menu.SysAuthorities)
		} else {
			err = db.Error
			if err != nil {
				return
			}
		}
	} else {
		return errors.New("此菜单存在子菜单不可删除")
	}

	return err
}

func (b *BaseMenuService) UpdateBaseMenu(menu model.SysBaseMenu) (err error) {
	var oldMenu model.SysBaseMenu
	updateMap := make(map[string]interface{})
	updateMap["keep_alive"] = menu.KeepAlive
	updateMap["close_tab"] = menu.CloseTab
	updateMap["parent_id"] = menu.ParentID
	updateMap["path"] = menu.Path
	updateMap["component"] = menu.Component
	updateMap["name"] = menu.Name
	updateMap["hidden"] = menu.Hidden
	updateMap["active_name"] = menu.ActiveName
	updateMap["title"] = menu.Title
	updateMap["icon"] = menu.Icon
	updateMap["sort"] = menu.Sort

	err = global.DB.Transaction(func(tx *gorm.DB) error {
		db := tx.Where("id = ?", menu.ID).First(&oldMenu)
		if oldMenu.Name != menu.Name {
			if !errors.Is(tx.Where("id <> ? AND name = ?", menu.ID, menu.Name).First(&model.SysBaseMenu{}).Error, gorm.ErrRecordNotFound) {
				global.LOG.Debug("存在相同name, 修改失败")
				return errors.New("存在相同name, 修改失败")
			}
		}

		txErr := tx.Unscoped().Delete(&model.SysBaseMenuParameter{}, "sys_base_menu_id = ?", menu.ID).Error
		if txErr != nil {
			global.LOG.Debug(txErr.Error())
			return txErr
		}
		txErr = tx.Unscoped().Delete(&model.SysBaseMenuBtn{}, "sys_base_menu_id", menu.ID).Error
		if txErr != nil {
			global.LOG.Debug(txErr.Error())
			return txErr
		}
		if len(menu.Parameters) > 0 {
			for k := range menu.Parameters {
				menu.Parameters[k].SysBaseMenuID = menu.ID
			}
			txErr = tx.Create(&menu.Parameters).Error
			if txErr != nil {
				global.LOG.Debug(txErr.Error())
				return txErr
			}
		}

		if len(menu.MenuBtn) > 0 {
			for k := range menu.MenuBtn {
				menu.MenuBtn[k].SysBaseMenuID = menu.ID
			}
			txErr = tx.Create(&menu.MenuBtn).Error
			if txErr != nil {
				global.LOG.Debug(txErr.Error())
				return txErr
			}
		}

		txErr = db.Updates(updateMap).Error
		if txErr != nil {
			global.LOG.Debug(txErr.Error())
			return txErr
		}
		return nil
	})

	return err
}

func (b *BaseMenuService) GetBaseMenuByID(id int) (menu model.SysBaseMenu, err error) {
	err = global.DB.Preload("MenuBtn").Preload("Parameters").Where("id = ?", id).First(&menu).Error
	return
}
