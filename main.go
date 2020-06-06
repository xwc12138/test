package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 完全读取
	// data, err := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(data))
	//  利用缓冲区来读取
	bytes := make([]byte, 1024)
	data := make([]byte, 0)
	//默认的缓冲区容量为4096
	read := bufio.NewReader(resp.Body)
	//重新制定缓冲区大小 一般使用默认的即可
	// read := bufio.NewReaderSize(resp.Body, 102400)
	for {
		count, err := read.Read(bytes)
		if err != nil {
			if err == io.EOF {
				fmt.Println("完整读取完毕！")
				fmt.Printf("err ====> %+v\n", err)
				break
			}
		}
		data = append(data, bytes[:count]...)
	}
	fmt.Print(string(data))
}
