/*
 * @Description:
 * @Author: yujiajie
 * @Date: 2023-12-10 19:50:59
 * @LastEditTime: 2024-03-25 15:57:45
 * @LastEditors: yujiajie
 */
package security

import (
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"gateway/core/codec"
	"gateway/core/container"
	"gateway/core/iox"
	"gateway/server/rest/header"
	"io"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	// ErrInvalidContentType is an error that indicates invalid content type.
	ErrInvalidContentType = errors.New("invalid content type")
	// ErrInvalidHeader is an error that indicates invalid X-Content-Security header.
	ErrInvalidHeader = errors.New("invalid X-Content-Security header")
	// ErrInvalidKey is an error that indicates invalid key.
	ErrInvalidKey = errors.New("invalid key")
	// ErrInvalidPublicKey is an error that indicates invalid public key.
	ErrInvalidPublicKey = errors.New("invalid public key")
	// ErrInvalidSecret is an error that indicates invalid secret.
	ErrInvalidSecret = errors.New("invalid secret")
)

type ContentSecurityHeader struct {
	Key         []byte
	Timestamp   string
	ContentType int
	Signature   string
}

func (h *ContentSecurityHeader) Encrypted() bool {
	return h.ContentType == header.CryptionType
}

func ParseContentSecurity(ctx *gin.Context, decrypters map[string]codec.RsaDecrypter) (*ContentSecurityHeader, error) {
	contentSecurity := ctx.GetHeader(header.ContentSecurity)
	attrs := header.ParseHeader(contentSecurity)
	fingerprint := attrs[header.KeyField]
	secret := attrs[header.SecretField]
	signature := attrs[header.SignatureField]

	if len(fingerprint) == 0 || len(secret) == 0 || len(signature) == 0 {
		return nil, ErrInvalidHeader
	}

	decrypter, ok := decrypters[fingerprint]
	if !ok {
		return nil, ErrInvalidPublicKey
	}

	decryptedSecret, err := decrypter.DecryptBase64(secret)
	if err != nil {
		return nil, ErrInvalidSecret
	}

	attrs = header.ParseHeader(string(decryptedSecret))
	base64Key := attrs[header.KeyField]
	timestamp := attrs[header.TimeField]
	contentType := attrs[header.TypeField]

	key, err := base64.StdEncoding.DecodeString(base64Key)
	if err != nil {
		return nil, err
	}

	cType, err := strconv.Atoi(contentType)
	if err != nil {
		return nil, err
	}

	return &ContentSecurityHeader{
		Key:         key,
		Timestamp:   timestamp,
		ContentType: cType,
		Signature:   signature,
	}, nil
}

func VerifySignature(ctx *gin.Context, securityHeader *ContentSecurityHeader, tolerance time.Duration) int {
	seconds, err := strconv.ParseInt(securityHeader.Timestamp, 10, 64)
	if err != nil {
		return header.CodeSignatureInvalidHeader
	}

	now := time.Now().Unix()
	toleranceSeconds := int64(tolerance.Seconds())
	if seconds+toleranceSeconds < now || now+toleranceSeconds < seconds {
		return header.CodeSignatureWrongTime
	}

	reqPath, reqQuery := getPathQuery(ctx)
	signContent := strings.Join([]string{
		securityHeader.Timestamp,
		ctx.Request.Method,
		reqPath,
		reqQuery, computeBodySignature(ctx),
	}, "\n")
	actualSignature := codec.HmacBase64(securityHeader.Key, signContent)

	if securityHeader.Signature == actualSignature {
		return header.CodeSignaturePass
	}

	log := container.App.GetLogger("default")
	log.Info("signature different, expect: %s, actual: %s",
		securityHeader.Signature, actualSignature)

	return header.CodeSignatureInvalidToken
}

func computeBodySignature(ctx *gin.Context) string {
	var dup io.ReadCloser
	ctx.Request.Body, dup = iox.DupReadCloser(ctx.Request.Body)
	sha := sha256.New()
	io.Copy(sha, ctx.Request.Body)
	ctx.Request.Body = dup
	return fmt.Sprintf("%x", sha.Sum(nil))
}

func getPathQuery(ctx *gin.Context) (string, string) {
	requestUri := ctx.GetHeader(header.RequestUriHeader)
	if len(requestUri) == 0 {
		return ctx.Request.URL.Path, ctx.Request.URL.RawQuery
	}

	uri, err := url.Parse(requestUri)
	if err != nil {
		return ctx.Request.URL.Path, ctx.Request.URL.RawQuery
	}

	return uri.Path, uri.RawQuery
}
