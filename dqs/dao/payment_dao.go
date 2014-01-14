package dao

import (
	"dqs/models"
	//"labix.org/v2/mgo"
	//"labix.org/v2/mgo/bson"
	//"time"
)

//获取系统参数
func GetAllPayment() ([]models.DevicePayment, error) {
	c := GetSession().DB(DatabaseName).C(PaymentCollection)
	payments := []models.DevicePayment{}

	err := c.Find(nil).All(&payments)
	if err != nil {
		payments = []models.DevicePayment{}
		return payments, err
	}
	return payments, nil
}
