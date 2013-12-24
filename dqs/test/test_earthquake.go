package main

import (
	"bytes"
	"dqs/models"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("test...")
	e := models.EarthQuake{}
	e.Time = "2013-11-03 12:12:34"
	e.Level = 6
	e.Longitude = 105.23654
	e.Latitude = 32.1543
	content, err0 := xml.MarshalIndent(e, "  ", "    ") //Marshal(e)
	if err0 != nil {
		fmt.Printf("marshal xml 出错:%s\n", err0.Error())
		return
	}
	body := bytes.NewBuffer(content)
	//fmt.Printf("%s", content)
	r, err := http.Post("http://115.29.37.85:8083/earthquake", "text/xml", body)
	if err != nil {
		fmt.Printf("调用远程接口出错:%s\n", err.Error())
		return
	}
	response, err2 := ioutil.ReadAll(r.Body)
	if err2 != nil {
		fmt.Printf("解析response内容出错:%s\n", err2.Error())
		return
	}
	fmt.Println(string(response))
}
