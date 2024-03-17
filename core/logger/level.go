/*
 * @Description:
 * @Author: yujiajie
 * @Date: 2023-12-11 22:33:26
 * @LastEditTime: 2023-12-25 23:16:29
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
	logger.Log(InfoLevel, format, args...)
}

func Debug(format string, args ...interface{}) {
	logger.Log(DebugLevel, format, args...)
}

func Warn(format string, args ...interface{}) {
	logger.Log(WarnLevel, format, args...)
}

func Error(format string, args ...interface{}) {
	logger.Log(ErrorLevel, format, args...)
}

func Panic(format string, args ...interface{}) {
	logger.Log(PanicLevel, format, args...)
}

func DPanic(format string, args ...interface{}) {
	logger.Log(DPanicLevel, format, args...)
}

func Fatal(format string, args ...interface{}) {
	logger.Log(FatalLevel, format, args...)
}
