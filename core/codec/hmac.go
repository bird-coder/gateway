/*
 * @Description:
 * @Author: yujiajie
 * @Date: 2023-12-07 23:48:44
 * @LastEditTime: 2023-12-07 23:53:26
 * @LastEditors: yujiajie
 */
package codec

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"io"
)

func Md5(body string) string {
	h := md5.New()
	io.WriteString(h, body)
	return hex.EncodeToString(h.Sum(nil))
}

func Hmac(key []byte, body string) []byte {
	h := hmac.New(sha256.New, key)
	io.WriteString(h, body)
	return h.Sum(nil)
}

func HmacBase64(key []byte, body string) string {
	return base64.StdEncoding.EncodeToString(Hmac(key, body))
}
