package sso

import (
	"sync"
)

// Http输出固定结构
type HttpRes struct {
	Code int         `json:"code" example:"9999"`             // 错误码，0表示正确，其他都为异常
	Msg  string      `json:"msg" example:"The Error Message"` // 错误信息，code非0时填充
	Data interface{} `json:"data"`                            // 业务返回内容，code非0时为nil，code为0时，基于业务填充
}

// SSO数据结构
type UserInfo struct {
	UserName string `json:"userName"`
	Type     int    `json:"type"`
}

// 缓存用户SSO状态的临时存储
var userCache sync.Map
