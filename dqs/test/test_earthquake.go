package main

import (
	"bytes"
	"dqs/models"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func testQuake() {
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

func main() {
	fmt.Println("test...")
	v := url.Values{}
	v.Add("BBOX", "98.413391367189,27.90917609375,109.10308863281,33.92968390625")
	v.Add("DATA_ARRAY", "103.758240-30.919430-6,103.758240-30.919430-6")
	v.Add("FORMAT", "image/jpeg")
	v.Add("HEIGHT", "548")
	v.Add("INTERPOLATION_STRATEGY", "2")
	v.Add("INTERVALS[]", "0,1,2,3,4,5,6,7,8,9,10,11,12")
	v.Add("INTERVALS_COLORS[]", "0xffffff00,0xff8633cc,0xffad33cc,0xffdd33cc,0xffe233cc,0xfff533cc,0xf3ff33cc,0x9fff33cc,0x72ff33cc,0x33f33dcc,0x33d35dcc,0x3340f0cc,0xff3333cc")
	v.Add("LAYERS", "dqs_layers")
	v.Add("RADIUS", "50")
	v.Add("RENDERER_TYPE", "2")
	v.Add("REQUEST", "GetMap")
	v.Add("SERVICE", "AMS")
	v.Add("SIMPLIFY_METHOD", "1")
	v.Add("SIMPLIFY_SIZE", "2")
	v.Add("SRS", "EPSG:4326")
	v.Add("STYLES", "")
	v.Add("TRANSPARENT", "TRUE")
	v.Add("VERSION", "1.1.1")
	v.Add("WIDTH", "973")

	r, err := http.PostForm("http://127.0.0.1:8080/geoserver/dqs/wms", v)
	if err != nil {
		fmt.Printf("调用远程接口出错:%s\n", err.Error())
		return
	}
	response, err2 := ioutil.ReadAll(r.Body)
	if err2 != nil {
		fmt.Printf("解析response内容出错:%s\n", err2.Error())
		return
	}
	ioutil.WriteFile("test.jpg", response, 0777)
	//fmt.Println(string(response))
}
