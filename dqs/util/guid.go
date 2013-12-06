package util

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"strconv"
	"time"
)

var hasher = md5.New()

func GUID() string {
	guid := GetMd5Hex(time.Now().Format(time.ANSIC) + strconv.Itoa(rand.Int()))
	return guid
}

//获取md5编码文件
func GetMd5Hex(input string) string {
	hasher.Reset()
	hasher.Write([]byte(input))
	return hex.EncodeToString(hasher.Sum(nil))
}
