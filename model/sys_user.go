package model

import (
	"admin/global"

	"github.com/gofrs/uuid"
)

type SysUser struct {
	global.MODEL
	UUID        uuid.UUID    `json:"uuid" gorm:"comment:用户UUID"`                  // 用户UUID
	Username    string       `json:"userName" gorm:"comment:用户登录名"`               // 用户登录名
	Password    string       `json:"-"  gorm:"comment:用户登录密码"`                    // 用户登录密码
	NickName    string       `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`   // 用户昵称
	HeaderImg   string       `json:"headerImg" gorm:"comment:用户头像"`               // 用户头像
	AuthorityID uint         `json:"authorityID" gorm:"default:1;comment:用户角色ID"` // 用户角色ID
	Authority   SysAuthority `json:"authority" gorm:"foreignKey:AuthorityID;references:AuthorityID;comment:用户角色"`
	Enable      int          `json:"enable" gorm:"default:1;comment:用户是否被冻结 1正常 2冻结"` //用户是否被冻结 1正常 2冻结
}

func (SysUser) TableName() string {
	return "sys_users"
}
