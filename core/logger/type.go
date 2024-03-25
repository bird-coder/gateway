/*
 * @Description:
 * @Author: yujiajie
 * @Date: 2023-12-11 22:33:51
 * @LastEditTime: 2024-03-25 15:55:15
 * @LastEditors: yujiajie
 */
package logger

type Logger interface {
	Log(level Level, format string, args ...interface{})
	Info(format string, args ...interface{})
	Warn(format string, args ...interface{})
	Debug(format string, args ...interface{})
	Error(format string, args ...interface{})
	Panic(format string, args ...interface{})
	DPanic(format string, args ...interface{})
	Fatal(format string, args ...interface{})
	String() string
	Sync() error
}

func Sync() {
	if zl, ok := DefaultLogger.(*zaplog); ok {
		zl.Sync()
	}
}
