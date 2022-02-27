package base64

import (
	b64 "encoding/base64"
)

type Base64 interface {
	Encode(str string) string
	Decode(str string) (string, error)
}

type base64 struct{}

func NewBase64() Base64 {
	return &base64{}
}

func (b *base64) Encode(str string) string {
	return b64.StdEncoding.EncodeToString([]byte(str))
}

func (b *base64) Decode(str string) (string, error) {
	decodeStr, err := b64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", err
	}

	return string(decodeStr), nil
}
