package controller

import (
	"fmt"
	"time"

	pbpkg "prelude/common/proto"

	"golang.org/x/net/context"
)

var _ pbpkg.SampleServer = &sampleServer{}

type sampleServer struct{}

// NewSampleServer 生成sample server
func NewSampleServer() *sampleServer {
	return &sampleServer{}
}

// Ping demo
func (ss *sampleServer) Ping(ctx context.Context, req *pbpkg.STSamplePingReq) (*pbpkg.STSamplePingRsp, error) {
	rsp := &pbpkg.STSamplePingRsp{
		Message: fmt.Sprintf(
			"BeiJing Time: %s, Welcome, %s!",
			time.Now().Format("2006-01-02 15:04:05"),
			req.Name,
		),
	}
	return rsp, nil
}
