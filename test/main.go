package main

import (
	server "dqs/quickserver"
	"fmt"
)

func main() {
	du := server.DllUtil
	str := []byte("SI30002345g005F")
	ret := du.SendStr(str)
	fmt.Printf("%s\n", ret)

}
