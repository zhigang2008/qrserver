package util

import (
	"crypto/md5"
	"encoding/hex"
)

const (
	salt = "@Steven@"
)

var h = md5.New()

func EncodePwd(pwd string) string {
	h.Reset()
	h.Write([]byte(pwd + salt))
	return hex.EncodeToString(h.Sum(nil))
}
