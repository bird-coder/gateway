/*
 * @Description:
 * @Author: yujiajie
 * @Date: 2023-12-11 22:33:51
 * @LastEditTime: 2023-12-21 23:29:14
 * @LastEditors: yujiajie
 */
package logger

type Logger interface {
	Log(level Level, format string, args ...interface{})
	String() string
}

func Sync() {
	if zl, ok := logger.(*zaplog); ok {
		zl.Sync()
	}
}
