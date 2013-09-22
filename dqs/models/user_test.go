package models

import (
	"fmt"
	"testing"
)

func TestCheckPwd(t *testing.T) {
	u := User{}
	u.SetPassword("abc")
	if u.CheckPwd("abc") {
		fmt.Printf("%s\n", u.Password)
		t.Log("ok")

	} else {
		t.Error("fail")
	}
}
