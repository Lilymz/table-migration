package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"reflect"
)
import . "github.com/Lilymz/table-migration/v2/pkg/config"

func main() {
	useLog()
	fmt.Println("运行··")
}
func useLog() {
	DaoLog.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	DaoLog.WithFields(logrus.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")
	message := &SmsMessage{
		Id:      "001",
		Name:    "群发消息",
		Content: "这是一条通知消息",
	}
	var printLog Print
	printLog = message
	printLog = printLog.(Print)
	DaoLog.WithFields(printLog.toString()).Info("打印实体属性对象")
	DaoLog.WithFields(logrus.Fields{
		"omg":    true,
		"number": 100,
	}).Fatal("The ice breaks!")
}

type SmsMessage struct {
	Id      string
	Name    string
	Content string
}

func (sms *SmsMessage) toString() map[string]interface{} {
	fields := make(map[string]interface{}, 8)
	v := reflect.ValueOf(sms).Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldName := t.Field(i).Name
		fieldValue := field.Interface()
		fields[fieldName] = fieldValue
	}
	return fields
}

type Print interface {
	toString() map[string]interface{}
}
