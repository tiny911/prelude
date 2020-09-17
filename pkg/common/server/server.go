package server

import (
	"github.com/tiny911/prelude/pkg/dao"

	"github.com/tiny911/prelude/pkg/controller"

	"github.com/tiny911/prelude/pkg/common/proto"

	svrpkg "github.com/tiny911/doraemon/server"
)

var (
	httpRegisters = []svrpkg.HTTPRegister{
		proto.RegisterPreludeHandlerFromEndpoint,
	}

	rpcRegister = func() {
		proto.RegisterPreludeServer(svr.RPCSvr, controller.NewPreludeServer())
	}

	stopFunc = func() { // 停止服务执行的一些回调
		dao.Close() // 关闭数据库
	}

	svr = svrpkg.New()
)

// Run 运行
func Run(port int) {
	svr.SetRPCRegister(rpcRegister)
	svr.SetHTTPRegisters(httpRegisters)
	svr.SetPort(port)
	svr.Run()
}

// Stop 停止
func Stop() {
	svr.Stop()
	stopFunc()
}
