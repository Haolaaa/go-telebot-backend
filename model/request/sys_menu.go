package request

import "admin/model"

type AddMenuAuthorityInfo struct {
	Menus       []model.SysBaseMenu `json:"menus"`
	AuthorityID uint                `json:"authorityID"`
}
