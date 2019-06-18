package error

import "google.golang.org/grpc/codes"

const (
	// ErrOk 成功
	ErrOk codes.Code = 0
	// ErrUnknown       未知错误
	ErrUnknown codes.Code = 101
	// ErrArgsInvalid   参数异常
	ErrArgsInvalid codes.Code = 102
	// ErrArgsEmpty     参数为空
	ErrArgsEmpty codes.Code = 103
	// ErrSystem        系统错误
	ErrSystem codes.Code = 104
	// ErrDB            数据库错误
	ErrDB codes.Code = 105
	// ErrNoServe       未提供服务
	ErrNoServe codes.Code = 106
)

// Msg 异常消息映射表
var Msg = map[codes.Code]string{
	ErrOk:          "成功",
	ErrUnknown:     "未知错误",
	ErrArgsInvalid: "参数异常",
	ErrArgsEmpty:   "参数为空",
	ErrSystem:      "系统错误",
	ErrDB:          "数据库错误",
	ErrNoServe:     "未提供服务",
}
