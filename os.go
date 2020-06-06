package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println("获得所有输入的参数列表xx。。。。。")
	for index, value := range args {
		fmt.Printf("%d ----- %v\n", index, value)
	}
}
