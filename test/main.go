package main

import (
	"fmt"
	"syscall"
	"unsafe"
)

func main() {
	var readStr []byte = []byte("SI01234567g")
	var retStr string = ""

	dll, e := syscall.LoadDLL("socket1.dll")
	if e != nil {
		println(e)
	}
	defer dll.Release()
	p, e2 := dll.FindProc("MBCHAR2INT")
	if e2 != nil {
		println(e2)
	}
	println(p.Name)
	ret, _, _ := p.Call(
		uintptr(unsafe.Pointer(&readStr)))
	println("get the result:", ret)

	fmt.Println("dll return :" + retStr)
}
