// Package response
package response

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

const (
	ErrorTypeWarning = 1
	ErrorTypeInfo    = 2
	ErrorTypeError   = 3
	ErrorTypeNoAuth  = 4
)

// Json 标准返回结果数据结构封装。
// 返回固定数据结构的JSON:
// err:  错误码(0:成功, 1:失败, >1:错误码);
// msg:  请求结果信息;
// data: 请求结果,根据不同接口返回结果的数据结构不同;
func Json(r *ghttp.Request, err int, msg string, data ...interface{}) {
	responseData := interface{}(nil)
	var code int
	if len(data) > 0 {
		responseData = data[0]
		if len(data) > 1 {
			code = gconv.Int(data[1])
		}
	} else {
		code = 200
	}
	r.Response.WriteHeader(code)
	_ = r.Response.WriteJson(g.Map{
		"code":    200,
		"error":   err,
		"message": msg,
		"data":    responseData,
	})
}

// JsonExit 返回JSON数据并退出当前HTTP执行函数。
func JsonExit(r *ghttp.Request, err int, msg string, data ...interface{}) {
	Json(r, err, msg, data...)
	r.ExitAll()
}

// Error 直接返回错误信息
func Error(r *ghttp.Request, msg string, data ...interface{}) {
	JsonExit(r, ErrorTypeError, msg, data...)
}

// Success 直接返回成功信息
func Success(r *ghttp.Request, msg string, data ...interface{}) {
	JsonExit(r, 0, msg, data...)
}
func EmptyList(r *ghttp.Request) {
	data := g.Slice{
		g.Slice{},
	}
	JsonExit(r, ErrorTypeInfo, "暂无列表内容", data...)
}
func EmptyRow(r *ghttp.Request) {
	data := g.Slice{
		g.Map{},
	}
	JsonExit(r, ErrorTypeInfo, "暂无内容", data...)
}
func NoAuthority(r *ghttp.Request) {
	JsonExit(r, ErrorTypeNoAuth, "暂无操作权限")
}

// Warning 直接返回警告信息
func Warning(r *ghttp.Request, msg string, data ...interface{}) {
	JsonExit(r, ErrorTypeWarning, msg, data...)
}

// Info 直接返回消息信息
func Info(r *ghttp.Request, msg string, data ...interface{}) {
	JsonExit(r, ErrorTypeInfo, msg, data...)
}
