/*
 * @Description:
 * @Author: yujiajie
 * @Date: 2024-03-17 23:11:56
 * @LastEditTime: 2024-03-22 10:43:39
 * @LastEditors: yujiajie
 */
package rest

import "github.com/gin-gonic/gin"

type Route struct {
	Group      string
	Path       string
	Handler    gin.HandlerFunc
	Middleware []gin.HandlerFunc
}
