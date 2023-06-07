package model

type SysMenu struct {
	SysBaseMenu
	MenuID      string                 `json:"menuID" gorm:"comment:菜单ID"`
	AuthorityID uint                   `json:"-" gorm:"comment:角色ID"`
	Children    []SysMenu              `json:"children" gorm:"-"`
	Parameters  []SysBaseMenuParameter `json:"parameters" gorm:"foreignKey:SysBaseMenuID;references:MenuID"`
	Btns        map[string]uint        `json:"btns" gorm:"-"`
}

type SysAuthorityMenu struct {
	MenuID      string `json:"menuID" gorm:"comment:菜单ID;column:sys_base_menu_id"`
	AuthorityID string `json:"-" gorm:"comment:角色ID;column:sys_authority_authority_id"`
}

func (s SysAuthorityMenu) TableName() string {
	return "sys_authority_menus"
}
