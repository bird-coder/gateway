/*
 * @Description:
 * @Author: yujiajie
 * @Date: 2023-12-12 23:23:31
 * @LastEditTime: 2023-12-12 23:24:46
 * @LastEditors: yujiajie
 */
package iox

import (
	"bytes"
	"io"
)

func DupReadCloser(reader io.ReadCloser) (io.ReadCloser, io.ReadCloser) {
	var buf bytes.Buffer
	tee := io.TeeReader(reader, &buf)
	return io.NopCloser(tee), io.NopCloser(&buf)
}
