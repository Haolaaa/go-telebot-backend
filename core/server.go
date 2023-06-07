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

	s := initServer(":8082", Router)
	time.Sleep(10 * time.Microsecond)
	global.LOG.Info(fmt.Sprintf("%s: admin服务启动成功，端口为%s", "admin", ":8082"))
	global.LOG.Error(s.ListenAndServe().Error())
}
