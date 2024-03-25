/*
 * @Description:
 * @Author: yujiajie
 * @Date: 2023-12-11 22:33:26
 * @LastEditTime: 2024-03-25 15:16:28
 * @LastEditors: yujiajie
 */
package logger

type Level string

const (
	InfoLevel   Level = "info"
	DebugLevel  Level = "debug"
	WarnLevel   Level = "warn"
	ErrorLevel  Level = "error"
	PanicLevel  Level = "panic"
	DPanicLevel Level = "dpanic"
	FatalLevel  Level = "fatal"
)

func Info(format string, args ...interface{}) {
	DefaultLogger.Log(InfoLevel, format, args...)
}

func Debug(format string, args ...interface{}) {
	DefaultLogger.Log(DebugLevel, format, args...)
}

func Warn(format string, args ...interface{}) {
	DefaultLogger.Log(WarnLevel, format, args...)
}

func Error(format string, args ...interface{}) {
	DefaultLogger.Log(ErrorLevel, format, args...)
}

func Panic(format string, args ...interface{}) {
	DefaultLogger.Log(PanicLevel, format, args...)
}

func DPanic(format string, args ...interface{}) {
	DefaultLogger.Log(DPanicLevel, format, args...)
}

func Fatal(format string, args ...interface{}) {
	DefaultLogger.Log(FatalLevel, format, args...)
}
