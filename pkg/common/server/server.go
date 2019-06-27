package server

import (
	"prelude/pkg/common/proto"
	"prelude/pkg/controller"
	"prelude/pkg/dao"

	svrpkg "github.com/tiny911/doraemon/server"
)

var (
	httpRegisters = []svrpkg.HTTPRegister{
		proto.RegisterSampleHandlerFromEndpoint,
	}

	rpcRegister = func() {
		proto.RegisterSampleServer(svr.RPCSvr, controller.NewSampleServer())
	}

	stopFunc = func() { //停止服务执行的一些回调
		dao.Close() //关闭数据库
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
