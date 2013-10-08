package util

import (
	"fmt"
	"testing"
)

func TestSendMail(t *testing.T) {
	m := InitMailer("smtp.163.com", "25", "china_dqs@163.com", true, "china_dqs", "dqs123")
	err := m.SendMail([]string{"wangzhg@minshenglife.com", "zhigang78@163.com"}, "test", "<h1>body message</h1>")
	if err != nil {
		fmt.Printf(err.Error())
		t.Error(err.Error())
	}
}
