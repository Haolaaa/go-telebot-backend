package request

import "admin/model"

type Register struct {
	Username    string `json:"userName" example:"用户名"`
	Password    string `json:"passWord" example:"密码"`
	NickName    string `json:"nickName" example:"昵称"`
	HeaderImg   string `json:"headerImg" example:"头像链接"`
	AuthorityID uint   `json:"authorityID" swaggertype:"string" example:"int 角色id"`
	Enable      int    `json:"enable" swaggertype:"string" example:"int 是否启用"`
}

type Login struct {
	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 密码
}

type ChangePassword struct {
	ID          uint   `json:"-"`           // 从 JWT 中提取 user id，避免越权
	Password    string `json:"password"`    // 密码
	NewPassword string `json:"newPassword"` // 新密码
}

// Modify  user's auth structure
type SetUserAuthority struct {
	ID          uint
	AuthorityID uint `json:"authorityID"`
}

type ChangeUserInfo struct {
	ID          uint                 `gorm:"primarykey"`                                                                           // 主键ID
	NickName    string               `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`                                            // 用户昵称
	Phone       string               `json:"phone"  gorm:"comment:用户手机号"`                                                          // 用户手机号
	AuthorityID uint                 `json:"authorityIds" gorm:"-"`                                                                // 角色ID
	Email       string               `json:"email"  gorm:"comment:用户邮箱"`                                                           // 用户邮箱
	HeaderImg   string               `json:"headerImg" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:用户头像"` // 用户头像
	Enable      int                  `json:"enable" gorm:"comment:冻结用户"`                                                           //冻结用户
	Authorities []model.SysAuthority `json:"-" gorm:"many2many:sys_user_authority;"`
}
