package model

import "time"

type SysAuthority struct {
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time    `sql:"index"`
	ParentId      *uint         `json:"parentId" gorm:"comment:父角色ID"`                                       // 父角色ID
	AuthorityID   uint          `json:"authorityID" gorm:"not null;unique;primary_key;comment:角色ID;size:90"` // 角色ID
	AuthorityName string        `json:"authorityName" gorm:"comment:角色名"`                                    // 角色名
	SysBaseMenus  []SysBaseMenu `json:"menus" gorm:"many2many:sys_authority_menus;"`
	Users         []SysUser     `json:"-" gorm:"many2many:sys_user_authority;"`
}

func (SysAuthority) tableName() string {
	return "sys_authorities"
}
