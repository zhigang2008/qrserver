package quickserver

import (
	"errors"
	"fmt"
	"syscall"
	"unsafe"
)

//用来与dll进行交互的组件
var (
	DllUtil dllUtil
)

//包含调用的dll以及 其中的function句柄
type dllUtil struct {
	dll                   *syscall.DLL
	p_parseReadFlashParam *syscall.Proc
	p_parseReadSetParam   *syscall.Proc
	p_ParseDelParam       *syscall.Proc
	p_ParseSetParam       *syscall.Proc
	p_GenerateSetParam    *syscall.Proc
	p_parseFlashData      *syscall.Proc
	p_GenerateReadParam   *syscall.Proc
	p_sendStr             *syscall.Proc
}

//初始化数据处理器
func init() {
	DllUtil = dllUtil{}
	DllUtil.dll = syscall.MustLoadDLL("socket1.dll")
	DllUtil.p_parseReadFlashParam = DllUtil.dll.MustFindProc("parseReadFlashParam")
	DllUtil.p_parseReadSetParam = DllUtil.dll.MustFindProc("parseReadSetParam")
	DllUtil.p_ParseDelParam = DllUtil.dll.MustFindProc("ParseDelParam")
	DllUtil.p_ParseSetParam = DllUtil.dll.MustFindProc("ParseSetParam")
	DllUtil.p_GenerateSetParam = DllUtil.dll.MustFindProc("GenerateSetParam")
	DllUtil.p_parseFlashData = DllUtil.dll.MustFindProc("parseFlashData")
	DllUtil.p_GenerateReadParam = DllUtil.dll.MustFindProc("GenerateReadParam")
	//CRC校验用
	DllUtil.p_sendStr = DllUtil.dll.MustFindProc("sendStr")

}

//释放Dll资源
func (dp *dllUtil) FreeDLL() {
	dp.dll.Release()
}

//DLL解析接收的突发数据
func (dp *dllUtil) ParseReadFlashParam(rec []byte) (*FlashData, error) {
	flashData := FlashData{}

	ok, _, _ := dp.p_parseReadFlashParam.Call(
		uintptr(unsafe.Pointer(&rec[0])),
		uintptr(unsafe.Pointer(&flashData)))
	if ok != 1 {
		return nil, errors.New("DLL解析突发数据失败")
	}
	return &flashData, nil
}

//生成参数读取命令
func (dp *dllUtil) GenerateReadParam(param string) ([]byte, error) {
	p := []byte(param + "g")
	p = append(p, 0)

	ret := make([]byte, 20)

	fmt.Printf("ret b=%s\n", ret)

	ok, _, err := dp.p_GenerateReadParam.Call(
		uintptr(unsafe.Pointer(&p[0])),
		uintptr(unsafe.Pointer(&ret[0])))
	if ok == 1 {

		fmt.Printf("ret r=%s\n", ret)

		rett := []byte{}
		//截取真正的数据
		for _, v := range ret {
			if v == 0 {
				break
			}
			rett = append(rett, v)
		}
		return rett, nil
	} else {
		return nil, err //errors.New("调用dll解析读取参数失败")
	}
}

//DLL解析接收的读取参数命令返回的数据
func (dp *dllUtil) ParseReadSetParam(rec []byte) (*RetData, error) {
	retData := RetData{}

	ok, _, _ := dp.p_parseReadSetParam.Call(
		uintptr(unsafe.Pointer(&rec[0])),
		uintptr(unsafe.Pointer(&retData)))
	if ok != 1 {
		return nil, errors.New("DLL解析设备的设置参数失败")
	}
	return &retData, nil
}

//DLL解析删除设备参数是否成功
func (dp *dllUtil) ParseDelParam(rec []byte) bool {
	ok, _, _ := dp.p_ParseDelParam.Call(
		uintptr(unsafe.Pointer(&rec[0])))
	if ok == 1 {
		return true
	} else {
		return false
	}
}

//生成设置参数的指令

func (dp *dllUtil) GenerateSetParam(param string, retData *RetData) ([]byte, error) {
	p := []byte(param + "s")
	p = append(p, 0)

	ret := [1024]byte{}

	ok, _, err := dp.p_GenerateSetParam.Call(
		uintptr(unsafe.Pointer(&p[0])),
		uintptr(unsafe.Pointer(retData)),
		uintptr(unsafe.Pointer(&ret)))
	if ok == 1 {
		rett := []byte{}
		//截取真正的数据
		for _, v := range ret {
			if v == 0 {
				break
			}
			rett = append(rett, v)
		}
		return rett, nil
	} else {
		return nil, err //errors.New("调用dll解析读取参数失败")
	}
}

//DLL解析设备参数设置是否成功
func (dp *dllUtil) ParseSetParam(rec []byte) bool {
	ok, _, _ := dp.p_ParseSetParam.Call(
		uintptr(unsafe.Pointer(&rec[0])))
	if ok == 1 {
		return true
	} else {
		return false
	}
}

//DLL CRC校验
func (dp *dllUtil) SendStr(rec []byte) [5]byte {
	code := [5]byte{}
	rec = append(rec, 0)

	dp.p_sendStr.Call(
		uintptr(unsafe.Pointer(&rec[0])),
		uintptr(unsafe.Pointer(&code[0])))
	return code
}

//添加CRC校验码
func (dp *dllUtil) AppendCRCCode(rec []byte) []byte {
	code := dp.SendStr(rec)
	finalStr := append(rec, code[0], code[1], code[2], code[3])
	return finalStr
}

//校验字符串是否符合crc校验
func (dp *dllUtil) CheckCRCCode(str []byte) bool {
	str = []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9'}
	fmt.Printf("zifu=%d\n", str[5])
	fmt.Printf("str=%c\n", str)
	initCode := str[(len(str) - 5):]
	initStr := str[0:(len(str) - 4)]

	ccode := dp.SendStr([]byte(initStr))
	code := []byte{ccode[0], ccode[1], ccode[2], ccode[3]}
	fmt.Printf("initstr=%s \ninitCode=%d\nccode=%s\ncode=%s\n", initStr, initCode, ccode, code)
	if string(code) == string(initCode) {
		return true
	} else {
		return false
	}

}
