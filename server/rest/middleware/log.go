/*
 * @Description:
 * @Author: yujiajie
 * @Date: 2023-12-11 23:58:04
 * @LastEditTime: 2024-03-25 16:37:00
 * @LastEditors: yujiajie
 */
package middleware

import (
	"bytes"
	"fmt"
	"gateway/core/container"
	"gateway/core/iox"
	"gateway/core/timex"
	"io"
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	limitBodyBytes       = 1024
	defaultSlowThreshold = time.Millisecond * 500
)

func LogHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()

		lrw := newDetailLoggedResponseWriter(ctx.Writer)
		ctx.Writer = lrw

		var dup io.ReadCloser
		var buf bytes.Buffer
		ctx.Request.Body, dup = iox.DupReadCloser(ctx.Request.Body)
		io.Copy(&buf, ctx.Request.Body)
		body, _ := io.ReadAll(&buf)
		ctx.Request.Body = dup

		ctx.Next()

		duration := time.Since(start)

		logDetail(ctx, lrw, body, duration)
	}
}

func logDetail(ctx *gin.Context, response *detailLoggedResponseWriter, reqBody []byte, duration time.Duration) {
	var buf bytes.Buffer
	log := container.App.GetLogger("chain")

	request_id := ctx.GetInt64("RequestId")
	code := ctx.Writer.Status()
	buf.WriteString(fmt.Sprintf("[HTTP] %s - %d - %s - %s - %d\n=> %s\n",
		ctx.Request.Method, code, ctx.Request.RemoteAddr, timex.ReprOfDuration(duration), request_id, dumpRequest(ctx.Request)))

	if duration > defaultSlowThreshold {
		log.Info(fmt.Sprintf("[HTTP] %s - %d - %s - slowcall(%s) - %d\n=> %s\n", ctx.Request.Method, code, ctx.Request.RemoteAddr,
			timex.ReprOfDuration(duration), request_id, dumpRequest(ctx.Request)))
	}

	if len(reqBody) > 0 {
		buf.WriteString(fmt.Sprintf("%s\n", reqBody))
	}

	respBuf := response.buf.Bytes()
	if len(respBuf) > 0 {
		buf.WriteString(fmt.Sprintf("<= %s", respBuf))
	}

	if ctx.Writer.Status() < http.StatusInternalServerError {
		log.Info(buf.String())
	} else {
		log.Error(buf.String())
	}
}

func dumpRequest(r *http.Request) string {
	reqContent, err := httputil.DumpRequest(r, true)
	if err != nil {
		return err.Error()
	}

	return string(reqContent)
}

type detailLoggedResponseWriter struct {
	gin.ResponseWriter
	buf *bytes.Buffer
}

func newDetailLoggedResponseWriter(w gin.ResponseWriter) *detailLoggedResponseWriter {
	return &detailLoggedResponseWriter{
		ResponseWriter: w,
		buf:            new(bytes.Buffer),
	}
}

func (w *detailLoggedResponseWriter) Write(p []byte) (int, error) {
	w.buf.Write(p)
	return w.ResponseWriter.Write(p)
}
