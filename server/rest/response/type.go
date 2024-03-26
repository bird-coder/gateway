/*
 * @Author: yujiajie
 * @Date: 2024-03-26 11:05:17
 * @LastEditors: yujiajie
 * @LastEditTime: 2024-03-26 11:06:00
 * @FilePath: /gateway/server/rest/response/type.go
 * @Description:
 */
package response

type Response interface {
	SetCode(int)
	SetMsg(string)
	SetData(interface{})
	Clone() Response
}
