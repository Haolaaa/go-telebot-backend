package core

import (
	"admin/global"
	"admin/initialize"
	"fmt"
	"time"
)

type server interface {
	ListenAndServe() error
}

func RunWindowServer() {
	Router := initialize.Routers()

	address := fmt.Sprintf(":%d", global.CONFIG.System.Addr)

	s := initServer(address, Router)
	time.Sleep(10 * time.Microsecond)
	global.LOG.Info(fmt.Sprintf("%s: admin服务启动成功，端口为%s", "admin", address))
	global.LOG.Error(s.ListenAndServe().Error())
}
