package error

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

//NewGrpcErr 创建grpc风格的错误
func NewGrpcErr(c codes.Code) error {
	return grpc.Errorf(c, Msg[c])
}
