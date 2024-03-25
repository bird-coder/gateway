/*
 * @Description:
 * @Author: yujiajie
 * @Date: 2023-12-07 23:37:45
 * @LastEditTime: 2024-03-25 15:57:32
 * @LastEditors: yujiajie
 */
package middleware

import (
	"bytes"
	"errors"
	"gateway/core/codec"
	"gateway/core/container"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

const maxBytes = 1 << 20 // 1MiB

var errContentLengthExceeded = errors.New("content length exceeded")

func CryptionHanlder(key []byte) gin.HandlerFunc {
	return LimitCryptionHanlder(maxBytes, key)
}

func LimitCryptionHanlder(limitBytes int64, key []byte) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		LimitCryption(ctx, limitBytes, key)
	}
}

func LimitCryption(ctx *gin.Context, limitBytes int64, key []byte) {
	cw := newCryptionResponseWriter(ctx.Writer)
	ctx.Writer = cw
	defer cw.flush(key)

	if ctx.Request.ContentLength <= 0 {
		ctx.Next()
		return
	}

	if err := decryptBody(limitBytes, key, ctx); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ctx.Next()
}

func decryptBody(limitBytes int64, key []byte, ctx *gin.Context) error {
	if limitBytes > 0 && ctx.Request.ContentLength > limitBytes {
		return errContentLengthExceeded
	}

	var content []byte
	var err error
	if ctx.Request.ContentLength > 0 {
		content = make([]byte, ctx.Request.ContentLength)
		_, err = io.ReadFull(ctx.Request.Body, content)
	} else {
		content, err = io.ReadAll(io.LimitReader(ctx.Request.Body, maxBytes))
	}
	if err != nil {
		return err
	}

	output, err := codec.DecryptByEcb(string(content), key)
	if err != nil {
		return err
	}

	ctx.Request.Body = io.NopCloser(bytes.NewBuffer(output))

	return nil
}

type cryptionResponseWriter struct {
	gin.ResponseWriter
	buf *bytes.Buffer
}

func newCryptionResponseWriter(w gin.ResponseWriter) *cryptionResponseWriter {
	return &cryptionResponseWriter{
		ResponseWriter: w,
		buf:            new(bytes.Buffer),
	}
}

func (w *cryptionResponseWriter) Write(p []byte) (int, error) {
	return w.buf.Write(p)
}

func (w *cryptionResponseWriter) flush(key []byte) {
	if w.buf.Len() == 0 {
		return
	}

	content, err := codec.EncryptByEcb(key, w.buf.Bytes())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log := container.App.GetLogger("default")
	if n, err := io.WriteString(w.ResponseWriter, content); err != nil {
		log.Error("write response failed, error: %s", err)
	} else if n < len(content) {
		log.Error("actual bytes: %d, written bytes: %d", len(content), n)
	}
}
