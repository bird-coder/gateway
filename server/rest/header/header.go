/*
 * @Description:
 * @Author: yujiajie
 * @Date: 2023-12-10 21:08:41
 * @LastEditTime: 2024-03-25 10:52:18
 * @LastEditors: yujiajie
 */
package header

import "strings"

const (
	AllowOrigin      = "Access-Control-Allow-Origin"
	AllOrigins       = "*"
	AllowMethods     = "Access-Control-Allow-Methods"
	AllowHeaders     = "Access-Control-Allow-Headers"
	AllowCredentials = "Access-Control-Allow-Credentials"
	ExposeHeaders    = "Access-Control-Expose-Headers"
	RequestMethod    = "Access-Control-Request-Method"
	RequestHeaders   = "Access-Control-Request-Headers"
	AllowHeadersVal  = "Content-Type, Origin, X-CSRF-Token, Authorization, AccessToken, Token, Range"
	ExposeHeadersVal = "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers"
	Methods          = "GET, HEAD, POST, PATCH, PUT, DELETE, OPTIONS"
	AllowTrue        = "true"
	MaxAgeHeader     = "Access-Control-Max-Age"
	MaxAgeHeaderVal  = "86400"
	VaryHeader       = "Vary"
	OriginHeader     = "Origin"
)

const (
	ContentEncoding  = "Content-Encoding"
	ContentLength    = "Content-Length"
	ContentSecurity  = "X-Content-Security"
	RequestUriHeader = "X-Request-Uri"
	ApplicationJson  = "application/json"
	ContentType      = "Content-Type"
	JsonContentType  = "application/json; charset=utf-8"
	KeyField         = "key"
	SecretField      = "secret"
	TypeField        = "type"
	SignatureField   = "signature"
	TimeField        = "time"
	CryptionType     = 1
	GzipEncoding     = "gzip"
	AppIdField       = "appid"
	NonceField       = "nonce"
	ContentSignature = "X-Content-Signature"
)

const (
	formKey           = "form"
	pathKey           = "path"
	maxMemory         = 32 << 20 // 32MB
	maxBodyLen        = 8 << 20  // 8MB
	separator         = ";"
	tokensInAttribute = 2
)

const (
	// CodeSignaturePass means signature verification passed.
	CodeSignaturePass = iota
	// CodeSignatureInvalidHeader means invalid header in signature.
	CodeSignatureInvalidHeader
	// CodeSignatureWrongTime means wrong timestamp in signature.
	CodeSignatureWrongTime
	// CodeSignatureInvalidToken means invalid token in signature.
	CodeSignatureInvalidToken
)

func ParseHeader(headerVal string) map[string]string {
	ret := make(map[string]string)
	fields := strings.Split(headerVal, separator)

	for _, field := range fields {
		field = strings.TrimSpace(field)
		if len(field) == 0 {
			continue
		}

		kv := strings.SplitN(field, "=", tokensInAttribute)
		if len(kv) != tokensInAttribute {
			continue
		}

		ret[kv[0]] = kv[1]
	}

	return ret
}
