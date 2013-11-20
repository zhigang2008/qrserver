package quickserver

import (
	"fmt"
	"testing"
)

/**/
/*
 */
func TestGenerateSetParam(t *testing.T) {

	du := DllUtil

	params := "SI12345678"
	ret := new(SensorInfo)
	ret.SensorId = "SI12345678"
	ret.SiteName = "haidian000"
	ret.Longitude = 15.82629
	ret.Latitude = 322.43544
	ret.SiteType = 1
	ret.ObserveObject = 0
	ret.Accelerometer = 1
	ret.Direction = 0
	ret.RangeType = 1
	ret.Period = 1
	ret.RegionCode = "510802"
	ret.Custom1 = "00000000"
	ret.Custom2 = "00000000"
	ret.PGATrigger = 1
	ret.PGATrgThreshold = 22
	ret.SITrigger = 0
	ret.SITrgThreshold = 3.2
	ret.CombTrigger = 0
	ret.ReserveTrigger = 0
	ret.ReserveTrgThreshold = 0
	ret.PGAAlert = 1
	ret.PGAAlertThreshold = 45
	ret.SIAlert = 0
	ret.SIAlertThreshold = 6.7
	ret.CombAlert = 0
	ret.ReserveAlert = 0
	ret.ReserveAlertThreshold = 0
	ret.DA1 = 0
	ret.DA2 = 1
	ret.IO1 = 1
	ret.IO2 = 1

	retdata := SensorInfo2RetData(ret)
	r, _ := du.GenerateSetParam(params, retdata)
	fmt.Printf("\n[%s]", r)
	fmt.Printf("\n[%s]\n", du.AppendCRCCode(r))

}

/*
func TestParseSetParam(t *testing.T) {

	du := DllUtil
	str := "SI30002012s0001019b3f"
	ret := du.ParseSetParam([]byte(str))
	fmt.Printf("%s", ret)
}

*/

/*
func TestSendStr(t *testing.T) {

	du := DllUtil
	str := "SI30002012s000101"
	ret := du.SendStr([]byte(str))
	fmt.Printf("%s", ret)
}
*/
/*
func TestParseFlashData(t *testing.T) {

	du := DllUtil
	str := "FFSI30001001rB700000020080512142801SI30001001117300000380600000130924080512142837012000000000000000000000B70000000f230f230f230f230f230f230f230f230f230f2301f230f230f230f230f230f230f230f230f230f2302f230f230f230f230f230f230f230f230f230f23012340D"
	ret, frame, err := du.ParseFlashData([]byte(str), "SI30001001")
	fmt.Println(frame)
	fmt.Println(ret)
	fmt.Println(err)

}
*/
