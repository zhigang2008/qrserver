package quickserver

import (
	"fmt"
	"testing"
)

/*func TestGenerateSetParam(t *testing.T) {

	du := dllUtil{}

	params := "SI12345678"
	ret := new(RetData)
	ret.Accelerometer = 1
	ret.CombAlert = 1

	r, err := du.GenerateSetParam(params, ret)
	if err != nil {
		if string(r) == "" {
			t.Log("ok")
		} else {
			t.Error("fail")
		}
	} else {
		t.Error("调用失败")
	}
}
*/
func TestCheckCRCCode(t *testing.T) {

	du := DllUtil
	str := "SI30001051a003b1309010721SI300010511037582403091943111132629130901072137000019990452826600284265000000055"
	ret := du.CheckCRCCode([]byte(str))
	fmt.Println(ret)

	/*if string(ret) != "30f6" {
		t.Error("fail")

	} else {
		t.Log("OK")
	}
	*/
}

func TestParseFlashData(t *testing.T) {

	du := DllUtil
	str := "FFSI30001001rB700000020080512142801SI30001001117300000380600000130924080512142837012000000000000000000000B70000000f230f230f230f230f230f230f230f230f230f2301f230f230f230f230f230f230f230f230f230f2302f230f230f230f230f230f230f230f230f230f23012340D"
	ret, frame, err := du.ParseFlashData([]byte(str), "SI30001001")
	fmt.Println(frame)
	fmt.Println(ret)
	fmt.Println(err)
	/*if string(ret) != "30f6" {
		t.Error("fail")

	} else {
		t.Log("OK")
	}
	*/
}
