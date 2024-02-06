package models

import "net/http"

// 基本响应结构
type BaseResponse struct {
	Code    int         `json:"code"`    // 自定义的业务状态码
	Message string      `json:"message"` // 响应消息
	Data    interface{} `json:"data"`    // 响应数据
}

// 成功响应
func NewSuccessResponse(data interface{}) *BaseResponse {
	return &BaseResponse{
		Code:    http.StatusOK, // 200 OK
		Message: "success",
		Data:    data,
	}
}

// 错误响应
func NewErrorResponse(code int, message string) *BaseResponse {
	return &BaseResponse{
		Code:    code,
		Message: message,
		Data:    nil,
	}
}
