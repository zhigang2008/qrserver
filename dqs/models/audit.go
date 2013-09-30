package models

import (
	"time"
)

const (
	ActLogin = iota
	ActLogout
	ActOperation
)

//审计信息
//服务端后台使用
type AuditLog struct {
	ActTime    time.Time //操作时间
	UserId     string    //操作用户id
	UserName   string    //操作用户名称
	ActType    int       //操作类型
	ActContent string    //操作内容
	Status     bool      //状态
	RemoteAddr string    //访问地址
}
