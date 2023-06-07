package service

import (
	"admin/global"
	"admin/model"
	"admin/model/common/request"
	"errors"
	"strconv"

	"gorm.io/gorm"
)

type MenuService struct{}

var MenuServiceApp = new(MenuService)

func (menuService *MenuService) getMenuTreeMap(authorityID uint) (treeMap map[string][]model.SysMenu, err error) {
	var allMenus []model.SysMenu
	var baseMenu []model.SysBaseMenu
	var btns []model.SysAuthorityBtn
	treeMap = make(map[string][]model.SysMenu)

	var SysAuthorityMenus []model.SysAuthorityMenu
	err = global.DB.Where("sys_authority_authority_id = ?", authorityID).Find(&SysAuthorityMenus).Error
	if err != nil {
		return
	}

	var MenuIDs []string
	for i := range SysAuthorityMenus {
		MenuIDs = append(MenuIDs, SysAuthorityMenus[i].MenuID)
	}

	err = global.DB.Where("id in (?)", MenuIDs).Order("sort").Preload("Parameters").Find(&baseMenu).Error
	if err != nil {
		return
	}

	for i := range baseMenu {
		allMenus = append(allMenus, model.SysMenu{
			SysBaseMenu: baseMenu[i],
			AuthorityID: authorityID,
			MenuID:      strconv.Itoa(int(baseMenu[i].ID)),
			Parameters:  baseMenu[i].Parameters,
		})
	}

	err = global.DB.Where("authority_id = ?", authorityID).Preload("SysBaseMenuBtn").Find(&btns).Error
	if err != nil {
		return
	}
	var btnMap = make(map[uint]map[string]uint)
	for _, v := range btns {
		if btnMap[v.SysMenuID] == nil {
			btnMap[v.SysMenuID] = make(map[string]uint)
		}
		btnMap[v.SysMenuID][v.SysBaseMenuBtn.Name] = authorityID
	}
	for _, v := range allMenus {
		v.Btns = btnMap[v.ID]
		treeMap[v.ParentID] = append(treeMap[v.ParentID], v)
	}
	return treeMap, err
}

func (menuService *MenuService) getChildrenList(menu *model.SysMenu, treeMap map[string][]model.SysMenu) (err error) {
	menu.Children = treeMap[menu.MenuID]
	for i := 0; i < len(menu.Children); i++ {
		err = menuService.getChildrenList(&menu.Children[i], treeMap)
	}
	return err
}

func (menuService *MenuService) GetMenuTree(authorityID uint) (menus []model.SysMenu, err error) {
	menuTree, err := menuService.getMenuTreeMap(authorityID)
	menus = menuTree["0"]
	for i := 0; i < len(menus); i++ {
		err = menuService.getChildrenList(&menus[i], menuTree)
	}
	return menus, err
}

func (menuService *MenuService) getBaseMenuTreeMap() (treeMap map[string][]model.SysBaseMenu, err error) {
	var allMenus []model.SysBaseMenu
	treeMap = make(map[string][]model.SysBaseMenu)
	err = global.DB.Order("sort").Preload("MenuBtn").Preload("Parameters").Find(&allMenus).Error
	for _, v := range allMenus {
		treeMap[v.ParentID] = append(treeMap[v.ParentID], v)
	}
	return treeMap, err
}

func (menuService *MenuService) getBaseChildrenList(menu *model.SysBaseMenu, treeMap map[string][]model.SysBaseMenu) (err error) {
	menu.Children = treeMap[strconv.Itoa(int(menu.ID))]
	for i := 0; i < len(menu.Children); i++ {
		err = menuService.getBaseChildrenList(&menu.Children[i], treeMap)
	}
	return err
}

func (menuService *MenuService) GetInfoList() (list interface{}, total int64, err error) {
	var menuList []model.SysBaseMenu
	treeMap, err := menuService.getBaseMenuTreeMap()
	menuList = treeMap["0"]
	for i := 0; i < len(menuList); i++ {
		err = menuService.getBaseChildrenList(&menuList[i], treeMap)
	}
	return menuList, total, err
}

func (menuService *MenuService) AddBaseMenu(menu model.SysBaseMenu) error {
	if !errors.Is(global.DB.Where("name = ?", menu.Name).First(&model.SysBaseMenu{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在重复的name，请修改name")
	}
	return global.DB.Create(&menu).Error
}

func (menuService *MenuService) GetBaseMenuTree() (menus []model.SysBaseMenu, err error) {
	treeMap, err := menuService.getBaseMenuTreeMap()
	menus = treeMap["0"]
	for i := 0; i < len(menus); i++ {
		err = menuService.getBaseChildrenList(&menus[i], treeMap)
	}
	return menus, err
}

func (menuService *MenuService) AddMenuAuthority(menus []model.SysBaseMenu, authorityID uint) (err error) {
	var auth model.SysAuthority
	auth.AuthorityID = authorityID
	auth.SysBaseMenus = menus
	err = AuthorityServiceApp.SetMenuAuthority(&auth)
	return err
}

func (menuService *MenuService) GetMenuAuthority(info *request.GetAuthorityID) (menus []model.SysMenu, err error) {
	var baseMenu []model.SysBaseMenu
	var sysAuthorityMenus []model.SysAuthorityMenu
	err = global.DB.Where("sys_authority_authority_id = ?", info.AuthorityID).Find(&sysAuthorityMenus).Error
	if err != nil {
		return
	}

	var MenuIDs []string
	for i := range sysAuthorityMenus {
		MenuIDs = append(MenuIDs, sysAuthorityMenus[i].MenuID)
	}

	err = global.DB.Where("id in (?)", MenuIDs).Order("sort").Find(&baseMenu).Error

	for i := range baseMenu {
		menus = append(menus, model.SysMenu{
			SysBaseMenu: baseMenu[i],
			AuthorityID: info.AuthorityID,
			MenuID:      strconv.Itoa(int(baseMenu[i].ID)),
			Parameters:  baseMenu[i].Parameters,
		})
	}

	return menus, err
}
