package util

import (
	"encoding/base64"
	"testing"
)

func TestBase64(t *testing.T) {
	ori := "spider"
	stdBase64Str := base64.StdEncoding.EncodeToString([]byte(ori))
	rawBase64Str := base64.RawStdEncoding.EncodeToString([]byte(ori))
	t.Log(stdBase64Str)
	t.Log(rawBase64Str)
}
