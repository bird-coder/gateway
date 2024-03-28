/*
 * @Author: yujiajie
 * @Date: 2024-03-26 11:05:17
 * @LastEditors: yujiajie
 * @LastEditTime: 2024-03-28 16:07:37
 * @FilePath: /Gateway/server/rest/response/type.go
 * @Description:
 */
package response

type Response interface {
	SetRequestId(int64)
	SetCode(int)
	SetMsg(string)
	SetData(interface{})
	Clone() Response
}
