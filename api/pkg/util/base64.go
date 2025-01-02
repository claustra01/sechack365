package util

import (
	"bytes"
	"encoding/base64"
	"io"
	"strings"
)

func EncodeBase64(data []byte) string {
	var sb strings.Builder
	b64 := base64.NewEncoder(base64.StdEncoding, &sb)
	if _, err := b64.Write(data); err != nil {
		// NOTE: err should be nil
		panic(err)
	}
	b64.Close()
	return sb.String()
}

func DecodeBase64(s string) []byte {
	var buf bytes.Buffer
	b64 := base64.NewDecoder(base64.StdEncoding, strings.NewReader(s))
	if _, err := io.Copy(&buf, b64); err != nil {
		// NOTE: err should be nil
		panic(err)
	}
	return buf.Bytes()
}
