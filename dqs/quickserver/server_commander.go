package quickserver

import (
	"errors"
	//"fmt"
	log "github.com/cihub/seelog"
)

//发送设备参数读取指令
func CommandRead(id, remote string) error {
	connP := GetConnection(remote)
	if connP != nil {
		command, err := DllUtil.GenerateReadParam(id)
		if err == nil {
			//发送控制命令
			n, err0 := (*connP).Write(command)

			if err0 != nil {
				return err0
			} else {
				log.Infof("向[%s]设备发送参数读取指令成功:%d", id, n)

				//读取客户端反馈
				c := make(chan []byte)
				AddCommand(remote, "G", c)
				back := <-c

				log.Infof("back=%s\n", back)
				//取消控制命令
				DeleteCommand(remote, "G")

				//进行数据校验
				if len(back) < 11 {
					return errors.New("参数设置失败")
				}
				backid := back[0:10]
				if string(backid) != id || (back[10] != 'g' && back[10] != 'G') {
					log.Warn("接收到的数据与发送目标不匹配")
					//重新发送数据
					CommandRead(id, remote)
				} else {
					//进行数据处理
					err1 := dataProcessor.ProcessStatusData(back)
					if err1 != nil {
						return errors.New("读取数据更新失败:" + err1.Error())

					}
				}

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
func CommandSet(id, remote string, params *RetData) error {
	connP := GetConnection(remote)

	//log.Info(params)

	if connP != nil {
		command, err := DllUtil.GenerateSetParam(id, params)
		if err == nil {
			//添加CRC校验码
			commandCRC := DllUtil.AppendCRCCode(command)
			log.Infof("generated set command=[%s]\n", string(commandCRC))
			n, err0 := (*connP).Write(commandCRC)
			if err0 != nil {
				return err0
			} else {
				log.Infof("向[%s]设备发送参数设置指令成功:%d", id, n)

				//读取客户端反馈
				c := make(chan []byte)
				AddCommand(remote, "S", c)
				back := <-c

				log.Infof("back=%s\n", back)
				//取消控制命令
				DeleteCommand(remote, "S")

				//进行数据校验
				if len(back) < 11 {
					return errors.New("参数设置失败")
				}

				backid := back[0:10]
				if string(backid) != id || (back[10] != 's' && back[10] != 'S') {
					log.Warn("接收到的数据与发送目标不匹配")
					//重新发送数据
					CommandSet(id, remote, params)
				} else {
					//进行数据处理
					ok := DllUtil.ParseSetParam(back)
					if ok {
						return nil

					} else {
						return nil
						//return errors.New("参数设置失败")
					}
				}
			}

		} else {
			return errors.New("DLL操作失败[" + err.Error() + "]")
		}
	} else {
		return errors.New("服务器未与当前设备建立连接,或者该设备还未进行注册")
	}
	return nil
}

//发送波形图数据读取指令
func CommandFlashRead(id, remote string) error {
	connP := GetConnection(remote)
	if connP != nil {
		command, err := DllUtil.GenerateFlashReadParam(id)
		if err == nil {
			//发送控制命令
			n, err0 := (*connP).Write(command)

			if err0 != nil {
				return err0
			} else {
				log.Infof("向[%s]设备发送波形图读取指令成功:%d", id, n)

				//读取客户端反馈
				c := make(chan []byte)
				AddCommand(remote, "R", c)
				back := <-c

				//第一帧是报警信息
				log.Infof("back=%s\n", back)
				//进行处理第一帧报警信息,保存报警信息
				dataProcessor.ProcessWaveFlashData(back)

				//取消控制命令
				DeleteCommand(remote, "R")

				/*
					//进行数据校验
					if len(back) < 13 {
						return errors.New("参数设置失败")
					}
					backid := back[2:12]
					if string(backid) != id || (back[12] != 'r' && back[12] != 'R') {
						log.Warn("接收到的数据与发送目标不匹配")
						//重新发送数据
						CommandFlashRead(id, remote, params)
					} else {
						//进行数据处理
						ret, frame, err1 := DllUtil.ParseFlashData(back, string(backid))
						//TODO 进行波形图处理

						if err1 != nil {
							return errors.New("读取数据更新失败:" + err1.Error())

						} else {
							fmt.Println(frame)
							fmt.Println(ret)

						}

					}
				*/

			}

		} else {
			return errors.New("DLL操作失败[" + err.Error() + "]")
		}
	} else {
		return errors.New("服务器未与当前设备建立连接,或者该设备还未进行注册")
	}
	return nil
}
