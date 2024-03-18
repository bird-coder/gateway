/*
 * @Author: yujiajie
 * @Date: 2024-03-18 15:05:04
 * @LastEditors: yujiajie
 * @LastEditTime: 2024-03-18 15:05:20
 * @FilePath: /gateway/options/auth.go
 * @Description:
 */
package options

type AuthConfig struct {
	Secret     string
	PrevSecret string
}
