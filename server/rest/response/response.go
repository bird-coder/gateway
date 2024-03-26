/*
 * @Author: yujiajie
 * @Date: 2024-03-26 11:06:08
 * @LastEditors: yujiajie
 * @LastEditTime: 2024-03-26 11:10:52
 * @FilePath: /gateway/server/rest/response/response.go
 * @Description:
 */
package response

import "net/http"

type ApiResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
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

func (res ApiResponse) Clone() Response {
	return &res
}

var Default = &ApiResponse{}

func Error(code int, msg string) Response {
	res := Default.Clone()
	if msg != "" {
		res.SetMsg(msg)
	}
	res.SetCode(code)
	return res
}

func OK(data interface{}, msg string) Response {
	res := Default.Clone()
	if msg == "" {
		msg = "success"
	}
	res.SetMsg(msg)
	res.SetCode(http.StatusOK)
	res.SetData(data)
	return res
}
