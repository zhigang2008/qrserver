package quickserver

import (
	"errors"
	"fmt"
	log "github.com/cihub/seelog"
)

//发送设备参数读取指令
func CommandRead(id string) error {
	connP := ConnecitonPool[id]
	if connP != nil {
		command, err := DllUtil.GenerateReadParam(id)
		if err == nil {
			n, err0 := (*connP).Write(command)
			if err0 != nil {
				return err0
			} else {
				log.Infof("向[%s]设备发送参数读取指令成功:%d", id, n)
				return nil
			}

		} else {
			return errors.New("DLL操作失败[" + err.Error() + "]")
		}
	} else {
		return errors.New("服务器未与当前设备建立连接,或者该设备还未进行注册")
	}
	return nil
}

//发送参数设置指令
func CommandSet(id string, params *RetData) error {
	connP := ConnecitonPool[id]
	fmt.Println(params)

	if connP != nil {
		command, err := DllUtil.GenerateSetParam(id, params)
		if err == nil {
			fmt.Println(string(command))
			n, err0 := (*connP).Write(command)
			if err0 != nil {
				return err0
			} else {
				log.Infof("向[%s]设备发送参数设置指令成功:%d", id, n)
				return nil
			}

		} else {
			return errors.New("DLL操作失败[" + err.Error() + "]")
		}
	} else {
		return errors.New("服务器未与当前设备建立连接,或者该设备还未进行注册")
	}
	return nil
}
