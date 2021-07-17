package controller

import (
	"fmt"
	"time"

	pbpkg "github.com/tiny911/prelude/internal/common/proto"

	"golang.org/x/net/context"

	"github.com/tiny911/doraemon/log"
)

var _ pbpkg.PreludeServer = &PreludeServer{}

type PreludeServer struct {
	pbpkg.UnimplementedPreludeServer
}

// NewPreludeServer 生成prelude server
func NewPreludeServer() *PreludeServer {
	return &PreludeServer{}
}

// Ping demo
func (ss *PreludeServer) Ping(ctx context.Context, req *pbpkg.STPreludePingReq) (*pbpkg.STPreludePingRsp, error) {

	log.WithTrace(ctx).Info("test log")
	log.WithoutField().Info("test 2 log")
	log.Kit.Info("test 3")
	rsp := &pbpkg.STPreludePingRsp{
		Message: fmt.Sprintf(
			"BeiJing Time: %s, Welcome, %s!",
			time.Now().Format("2006-01-02 15:04:05"),
			req.Name,
		),
	}
	return rsp, nil
}
