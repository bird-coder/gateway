/*
 * @Author: yujiajie
 * @Date: 2024-03-18 11:01:10
 * @LastEditors: yujiajie
 * @LastEditTime: 2024-03-18 11:01:22
 * @FilePath: /gateway/core/constant/env.go
 * @Description:
 */
package constant

type Env string

const (
	Dev  Env = "development"
	Prod Env = "production"
)

func (e Env) String() string {
	return string(e)
}
