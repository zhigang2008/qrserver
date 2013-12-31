package util

import (
	"code.google.com/p/go.text/encoding/simplifiedchinese"
)

func UTF8ToGBK(text string) ([]byte, error) {
	dst := make([]byte, len(text)*2)
	tr := simplifiedchinese.GBK.NewEncoder()
	nDst, _, err := tr.Transform(dst, []byte(text), true)
	if err != nil {
		return []byte(text), err
	}
	return dst[:nDst], nil
}
func UTF8ToGBK18030(text string) ([]byte, error) {
	dst := make([]byte, len(text)*2)
	tr := simplifiedchinese.GB18030.NewEncoder()
	nDst, _, err := tr.Transform(dst, []byte(text), true)
	if err != nil {
		return []byte(text), err
	}
	return dst[:nDst], nil
}

func UTF8ToGB2312(text string) ([]byte, error) {
	dst := make([]byte, len(text)*2)
	tr := simplifiedchinese.HZGB2312.NewEncoder()
	nDst, _, err := tr.Transform(dst, []byte(text), true)
	if err != nil {
		return []byte(text), err
	}
	return dst[:nDst], nil
}
