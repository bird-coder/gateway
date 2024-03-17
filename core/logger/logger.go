package logger

import (
	"encoding/json"
	"fmt"
	"gateway/options"
	"io"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger Logger
)

const (
	Dev  = "dev"
	Prod = "prod"
)

type LoggerConfig struct {
	LogLevel   string `json:"level"`
	LogPath    string `json:"logpath"`
	MaxSize    int    `json:"maxsize"`
	MaxAge     int    `json:"age"`
	MaxBackups int    `json:"backups"`
	Compress   string `json:"compress"`
}

type zaplog struct {
	zap *zap.Logger
	al  *zap.AtomicLevel
}

func NewLogger(cfg *options.LoggerConfig, env string) Logger {
	var writer io.Writer

	level := toZapLevel(Level(cfg.LogLevel))

	var zapOptions []zap.Option
	if env == Dev {
		writer = os.Stdout
		zapOptions = append(zapOptions, zap.Development())
	} else {
		writer = NewRotateWriter(cfg)
	}
	zapOptions = append(zapOptions, zap.AddCaller(),
		zap.AddCallerSkip(1), zap.AddStacktrace(zap.WarnLevel))

	zl := New(writer, level, zapOptions...)
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
		break
	case InfoLevel:
		logLevel = zap.InfoLevel
		break
	case WarnLevel:
		logLevel = zap.WarnLevel
		break
	case ErrorLevel:
		logLevel = zap.ErrorLevel
		break
	case PanicLevel:
		logLevel = zap.PanicLevel
		break
	case DPanicLevel:
		logLevel = zap.DPanicLevel
		break
	case FatalLevel:
		logLevel = zap.FatalLevel
		break
	default:
		logLevel = zap.InfoLevel
		break
	}
	return logLevel
}

func formatConfig(configMap map[string]interface{}) *LoggerConfig {
	data, err := json.Marshal(configMap)
	if err != nil {
		fmt.Fprint(os.Stderr, "load logger config error, error: json marshal failed\n")
		os.Exit(1)
	}
	config := &LoggerConfig{
		LogLevel:   "debug",
		MaxSize:    128,
		MaxAge:     7,
		MaxBackups: 30,
		Compress:   "false",
	}
	if err := json.Unmarshal(data, config); err != nil {
		fmt.Fprintf(os.Stderr, "load logger config error, error: json unmarshal failed, data: %s\n", string(data))
		os.Exit(1)
	}
	return config
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
