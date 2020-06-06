package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestHttpClientGET(t *testing.T) {
	rsp, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println("err ===>", err)
		return
	}

	if rsp.StatusCode != 200 {
		fmt.Println("请求失败。。。")
		fmt.Println("statusCode == ", rsp.StatusCode)
		return
	}

	data, _ := ioutil.ReadAll(rsp.Body)
	fmt.Println(string(data))
}

//post 请求
func TestHttpClientPOST(t *testing.T) {
	var str = "hello"
	strread := strings.NewReader(str)
	//Content-type : application/json;charset=utf-8
	rsp, err := http.Post("http://www.baidu.com", "application/json;charset=utf-8", strread)
	// rsp, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println("err ===>", err)
		return
	}

	if rsp.StatusCode != 200 {
		fmt.Println("请求失败。。。")
		fmt.Println("statusCode == ", rsp.StatusCode)
		return
	}

	data, _ := ioutil.ReadAll(rsp.Body)
	fmt.Println(string(data))
}

func TestHttpClientComplex(t *testing.T) {
	obj := struct {
		name string
		age  int
	}{
		"zhangsan",
		12,
	}
	databytes, _ := json.Marshal(obj)
	requestReader := bytes.NewReader(databytes)
	request, _ := http.NewRequest("POST", "http://www.baidu.com", requestReader)
	client := http.DefaultClient
	rsp, err := client.Do(request)
	if err != nil {
		fmt.Println("err ===>", err)
		return
	}

	if rsp.StatusCode != 200 {
		fmt.Println("请求失败。。。")
		fmt.Println("statusCode == ", rsp.StatusCode)
		return
	}

	data, _ := ioutil.ReadAll(rsp.Body)
	fmt.Println(string(data))
}
