package logger

import (
	"fmt"
	"gateway/core/constant"
	"gateway/options"
	"io"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger Logger
)

type zaplog struct {
	zap *zap.Logger
	al  *zap.AtomicLevel
}

func NewLogger(cfg *options.LoggerConfig, env string) Logger {
	var writer io.Writer

	level := toZapLevel(Level(cfg.LogLevel))

	var zapOptions []zap.Option
	if env == constant.Dev.String() {
		writer = os.Stdout
		zapOptions = append(zapOptions, zap.Development())
	} else {
		writer = NewRotateWriter(cfg)
	}
	zapOptions = append(zapOptions, zap.AddCaller(),
		zap.AddCallerSkip(1), zap.AddStacktrace(zap.WarnLevel))

	zl := New(writer, level, zapOptions...)
	logger = zl
	return zl
}

func New(out io.Writer, level zapcore.Level, opts ...zap.Option) *zaplog {
	if out == nil {
		out = os.Stdout
	}

	encodeConfig := zap.NewProductionEncoderConfig()
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encodeConfig.EncodeName = zapcore.FullNameEncoder

	al := zap.NewAtomicLevelAt(level)

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encodeConfig),
		zapcore.AddSync(out),
		al,
	)

	return &zaplog{zap: zap.New(core, opts...), al: &al}
}

func (zl *zaplog) SetLevel(level Level) {
	if zl.al != nil {
		zl.al.SetLevel(toZapLevel(level))
	}
}

func toZapLevel(level Level) zapcore.Level {
	var logLevel zapcore.Level
	switch level {
	case DebugLevel:
		logLevel = zap.DebugLevel
	case InfoLevel:
		logLevel = zap.InfoLevel
	case WarnLevel:
		logLevel = zap.WarnLevel
	case ErrorLevel:
		logLevel = zap.ErrorLevel
	case PanicLevel:
		logLevel = zap.PanicLevel
	case DPanicLevel:
		logLevel = zap.DPanicLevel
	case FatalLevel:
		logLevel = zap.FatalLevel
	default:
		logLevel = zap.InfoLevel
	}
	return logLevel
}

func (zl *zaplog) Log(level Level, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	switch level {
	case InfoLevel:
		zl.zap.Info(msg)
	case DebugLevel:
		zl.zap.Debug(msg)
	case WarnLevel:
		zl.zap.Warn(msg)
	case ErrorLevel:
		zl.zap.Error(msg)
	case PanicLevel:
		zl.zap.Panic(msg)
	case DPanicLevel:
		zl.zap.DPanic(msg)
	case FatalLevel:
		zl.zap.Fatal(msg)
	}
}

func (zl *zaplog) Sync() error {
	return zl.zap.Sync()
}

func (zl *zaplog) String() string {
	return "zap"
}
