package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	type resultInfo struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			Fault struct {
				Code struct {
					Value string
				}
				Reason struct {
					Text string
				}
			}
			LDdatainputResponse struct {
				LDdatainputResult string
			}
		}
	}

	datastr := `<?xml version="1.0" encoding="utf-8"?>
  <soap:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
   <soap:Body>
    <LDdatainput xmlns="http://tempuri.org/">
      <dzid>%s</dzid>
      <dzxl>%s</dzxl>
      <xml>%s</xml>
    </LDdatainput>
   </soap:Body>
  </soap:Envelope>`

	sendStr := fmt.Sprintf(datastr, "20150101002", "001", "<DataList></DataList>")
	fmt.Println(sendStr)
	client := &http.Client{}
	reqest, _ := http.NewRequest("POST", "http://218.6.242.153/Service.asmx", strings.NewReader(sendStr))

	reqest.Header.Add("Content-Type", "text/xml; charset=utf-8")
	reqest.Header.Add("SOAPAction", "http://tempuri.org/LDdatainput")
	//reqest.Header.Set("Accept", "application/xml;q=0.9,*/*;q=0.8")
	/*reqest.Header.Set("Accept-Charset", "utf-8;q=0.7,*;q=0.3")
	reqest.Header.Set("Accept-Encoding", "gzip,deflate,sdch")
	reqest.Header.Set("Accept-Language", "zh-CN,zh;q=0.8")
	reqest.Header.Set("Cache-Control", "max-age=0")
	reqest.Header.Set("Connection", "keep-alive")
	*/

	response, _ := client.Do(reqest)
	if response.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response.Body)

		result := resultInfo{}
		err := xml.Unmarshal(body, &result)
		if err == nil {
			if result.Body.LDdatainputResponse.LDdatainputResult == "0" {
				fmt.Println("调用成功")
			}
		}
		//bodystr := string(body)
		//fmt.Println("success:" + bodystr)
	} else {
		fmt.Println(response.StatusCode)
		body, _ := ioutil.ReadAll(response.Body)
		result := resultInfo{}
		err := xml.Unmarshal(body, &result)
		if err == nil {
			fmt.Println(result.Body.Fault.Reason)
		}
		bodystr := string(body)
		fmt.Println("error:" + bodystr)
	}
}
