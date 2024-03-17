/*
 * @Description:
 * @Author: yujiajie
 * @Date: 2024-03-17 23:11:56
 * @LastEditTime: 2024-03-17 23:13:06
 * @LastEditors: yujiajie
 */
package rest

import "github.com/gin-gonic/gin"

type Route struct {
	Path    string
	Handler gin.HandlerFunc
}
