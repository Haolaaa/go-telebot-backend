package main

import (
	"admin/core"
	"admin/global"
	"admin/initialize"

	"go.uber.org/zap"
)

func main() {
	global.VP = core.Viper()
	initialize.OtherInit()
	global.LOG = core.Zap()
	zap.ReplaceGlobals(global.LOG)
	global.DB = initialize.Gorm()
	initialize.Timer()
	if global.DB != nil {
		initialize.RegisterTables()
		db, _ := global.DB.DB()
		defer db.Close()
	}
	core.RunWindowServer()
}
