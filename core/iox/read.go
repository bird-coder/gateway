/*
 * @Description:
 * @Author: yujiajie
 * @Date: 2023-12-12 23:23:31
 * @LastEditTime: 2024-03-26 14:05:45
 * @LastEditors: yujiajie
 */
package iox

import (
	"bytes"
	"io"
)

// 由于go语言buffer的特性(buffer无法重复读取)，这里需要通过teereader，在读取第一个buffer时同步写入到第二个buffer
func DupReadCloser(reader io.ReadCloser) (io.ReadCloser, io.ReadCloser) {
	var buf bytes.Buffer
	tee := io.TeeReader(reader, &buf)
	return io.NopCloser(tee), io.NopCloser(&buf)
}
