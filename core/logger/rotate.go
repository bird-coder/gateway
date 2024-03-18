/*
 * @Description:
 * @Author: yujiajie
 * @Date: 2023-12-21 22:59:20
 * @LastEditTime: 2024-03-18 10:42:30
 * @LastEditors: yujiajie
 */
package logger

import (
	"gateway/options"
	"io"
	"strings"

	"gopkg.in/natefinch/lumberjack.v2"
)

func NewRotateWriter(cfg *options.LoggerConfig) io.Writer {
	return &lumberjack.Logger{
		Filename:   cfg.LogPath,
		MaxSize:    cfg.MaxSize,
		MaxAge:     cfg.MaxAge,
		MaxBackups: cfg.MaxBackups,
		Compress:   strings.ToLower(cfg.Compress) == "true",
	}
}
