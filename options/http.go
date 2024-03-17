/*
 * @Description:
 * @Author: yujiajie
 * @Date: 2023-11-15 23:28:26
 * @LastEditTime: 2024-03-17 14:09:52
 * @LastEditors: yujiajie
 */
package options

type HttpConfig struct {
	Addr           string
	ReadTimeout    int
	WriteTimeout   int
	MaxHeaderBytes int
}
