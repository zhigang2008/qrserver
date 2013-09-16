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
	str := "SI30001051a003b1309010721SI30001051103758240309194311113262913090107213700001999045282660028426500000006b59"
	ret := du.CheckCRCCode(str)
	fmt.Println(ret)

	/*if string(ret) != "30f6" {
		t.Error("fail")

	} else {
		t.Log("OK")
	}
	*/
}
