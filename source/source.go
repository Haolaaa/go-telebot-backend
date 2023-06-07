package source

import (
	"admin/model"
	"admin/utils"

	"github.com/gofrs/uuid"
)

type UserEntities []model.SysUser

type SiteConfigEntities []model.SysSiteConfig

type AuthoritiesEntities []model.SysAuthority

type MenuEntities []model.SysBaseMenu

func GetUserEntities() UserEntities {
	users := UserEntities{
		{
			UUID:      uuid.Must(uuid.NewV4()),
			Username:  "admin",
			Password:  utils.BcryptHash("123456"),
			NickName:  "管理员",
			HeaderImg: "",
		},
	}
	return users
}

func GetSiteConfigEntities() SiteConfigEntities {
	siteConfigs := SiteConfigEntities{
		{
			ParentName:    "桃红",
			SiteName:      "桃红",
			SiteKey:       "taohsj",
			DirectPlayUrl: "https://thplay.taoh2550.com",
			CFPlayUrl:     "https://thplay.3cej.com",
			CDNPlayUrl:    "https://aws-news-taohplay.3cej.com",
			VideoCover:    "https://thimg.3cej.com",
			SiteID:        4,
		},
		{
			ParentName:    "桃红",
			SiteName:      "汤姆叔叔",
			SiteKey:       "uncletom",
			DirectPlayUrl: "https://tmplay3.tomtv079.com",
			CFPlayUrl:     "https://tmplay3.3cej.com",
			CDNPlayUrl:    "https://aws-tmplay.3cej.com",
			VideoCover:    "https://tmpic3.3cej.com",
			SiteID:        1,
		},
		{
			ParentName:    "桃红",
			SiteName:      "骚虎",
			SiteKey:       "saohuold",
			DirectPlayUrl: "https://shpplay.2850saohu.com",
			CFPlayUrl:     "https://shpplay.3cej.com",
			CDNPlayUrl:    "https://aws-news-shpplay.3cej.cc",
			VideoCover:    "https://shpimg.3cej.com",
			SiteID:        3,
		},
		{
			ParentName:    "桃红",
			SiteName:      "四虎",
			SiteKey:       "dh365",
			DirectPlayUrl: "https://4hu.saohu687.com",
			CFPlayUrl:     "https://sihplay.3cej.com",
			CDNPlayUrl:    "https://aws-news-sihplay.3cej.cc",
			VideoCover:    "https://shpimg.3cej.com",
			SiteID:        17,
		},
		{
			ParentName:    "集团",
			SiteName:      "东方2",
			SiteKey:       "dongfang",
			DirectPlayUrl: "https://news-zlplay.971df.com",
			CFPlayUrl:     "https://news-df2play.3d9b.com",
			CDNPlayUrl:    "https://aws-news-df2play.c20df.cc",
			VideoCover:    "https://news-dfimg.3d9b.com",
			SiteID:        2,
		},
		{
			ParentName:    "集团",
			SiteName:      "汤姆",
			SiteKey:       "tom",
			DirectPlayUrl: "https://news-tmplay.5682tom.com",
			CFPlayUrl:     "https://news-tmplay.3d9b.com",
			CDNPlayUrl:    "https://aws-tmplay3.5693tom.com",
			VideoCover:    "https://news-tmimg.3d9b.com",
			SiteID:        1,
		},
		{
			ParentName:    "集团",
			SiteName:      "好大哥",
			SiteKey:       "hdg",
			DirectPlayUrl: "https://news-hdgplay.hdg013.vip",
			CFPlayUrl:     "https://news-hdgplay.3d9b.com",
			CDNPlayUrl:    "https://aws-news-hdgplay.hdg005.cc",
			VideoCover:    "https://news-hdgpic.3d9b.com",
			SiteID:        3,
		},
		{
			ParentName:    "集团",
			SiteName:      "365导航",
			SiteKey:       "dh365",
			DirectPlayUrl: "https://news-hdgplay.hdg013.vip",
			CFPlayUrl:     "https://news-365play.3d9b.com",
			CDNPlayUrl:    "https://aws-news-hdgplay.hdg005.cc",
			VideoCover:    "https://news-365pic.3d9b.com",
			SiteID:        4,
		},
		{
			ParentName:    "集团",
			SiteName:      "柠檬",
			SiteKey:       "nm",
			DirectPlayUrl: "https://news-lemonplay.nmsp662.com",
			CFPlayUrl:     "https://news-lemonplay.3d9b.com",
			CDNPlayUrl:    "https://aws-news-lemonplay.nm016.cc",
			VideoCover:    "https://news-lemonimg.3d9b.com",
			SiteID:        23,
		},
		{
			ParentName:    "集团",
			SiteName:      "樱桃",
			SiteKey:       "yt",
			DirectPlayUrl: "https://news-cherryplay.yt1350.com",
			CFPlayUrl:     "https://news-cherryplay.3d9b.com",
			CDNPlayUrl:    "https://aws-news-cherryplay.yt026.cc",
			VideoCover:    "https://news-cherryimg.3d9b.com",
			SiteID:        22,
		},
	}

	return siteConfigs
}

func GetMenuEntities() MenuEntities {
	menus := MenuEntities{
		{
			MenuLevel: 0,
			Hidden:    false,
			ParentID:  "0",
			Path:      "/home/index",
			Name:      "home",
			Component: "/home/index",
			Sort:      1,
			Meta: model.Meta{
				Title: "首页",
				Icon:  "HomeFilled",
			},
		},
		{
			MenuLevel: 0,
			Hidden:    false,
			ParentID:  "0",
			Path:      "/proTable",
			Name:      "proTable",
			Redirect:  "/proTable/useProTable",
			Sort:      2,
			Meta: model.Meta{
				Title: "站点管理",
				Icon:  "MessageBox",
			},
		},
		{
			MenuLevel: 0,
			Hidden:    false,
			ParentID:  "2",
			Path:      "/proTable/useProTable",
			Name:      "useProTable",
			Component: "/proTable/useProTable/index",
			Sort:      1,
			Meta: model.Meta{
				Title: "M3U8管理",
				Icon:  "Menu",
			},
		},
		{
			MenuLevel: 0,
			Hidden:    false,
			ParentID:  "0",
			Path:      "/system",
			Name:      "system",
			Redirect:  "/system/accountManage",
			Sort:      3,
			Meta: model.Meta{
				Title: "系统管理",
				Icon:  "Tools",
			},
		},
		{
			MenuLevel: 0,
			Hidden:    false,
			ParentID:  "4",
			Path:      "/system/menuManage",
			Name:      "menuManage",
			Component: "/system/menuManage/index",
			Sort:      1,
			Meta: model.Meta{
				Title: "菜单管理",
				Icon:  "Menu",
			},
		},
	}

	return menus
}
