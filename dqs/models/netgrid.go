package models

import "fmt"

type NetGrid struct {
	Longitude float32 //震中位置
	Latitude  float32
	Value     int     //值
	PGAValue  float32 //值
	SIValue   float32 //值
}

//格式化输出
func (n *NetGrid) String() string {
	return fmt.Sprintf("%f-%f-%d", n.Longitude, n.Latitude, n.Value)
}

//格式化输出
func (n *NetGrid) StringPGA() string {
	return fmt.Sprintf("%f-%f-%f", n.Longitude, n.Latitude, n.PGAValue)
}

//格式化输出
func (n *NetGrid) StringSI() string {
	return fmt.Sprintf("%f-%f-%f", n.Longitude, n.Latitude, n.SIValue)
}
