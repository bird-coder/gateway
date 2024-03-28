/*
 * @Author: yujiajie
 * @Date: 2024-03-26 11:06:08
 * @LastEditors: yujiajie
 * @LastEditTime: 2024-03-28 16:07:32
 * @FilePath: /Gateway/server/rest/response/response.go
 * @Description:
 */
package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiResponse struct {
	RequestId int64       `json:"requestId"`
	Code      int         `json:"code"`
	Msg       string      `json:"msg"`
	Data      interface{} `json:"data"`
}

func (res *ApiResponse) SetData(data interface{}) {
	res.Data = data
}

func (res *ApiResponse) SetMsg(msg string) {
	res.Msg = msg
}

func (res *ApiResponse) SetCode(code int) {
	res.Code = code
}

func (res *ApiResponse) SetRequestId(request_id int64) {
	res.RequestId = request_id
}

func (res ApiResponse) Clone() Response {
	return &res
}

var Default = &ApiResponse{}

func Error(ctx *gin.Context, code int, msg string) Response {
	res := Default.Clone()
	if msg != "" {
		res.SetMsg(msg)
	}
	res.SetCode(code)
	res.SetRequestId(ctx.GetInt64("RequestId"))
	return res
}

func OK(ctx *gin.Context, data interface{}, msg string) Response {
	res := Default.Clone()
	if msg == "" {
		msg = "success"
	}
	res.SetMsg(msg)
	res.SetCode(http.StatusOK)
	res.SetData(data)
	res.SetRequestId(ctx.GetInt64("RequestId"))
	return res
}
