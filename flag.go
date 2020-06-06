package main

import (
	"flag"
	"fmt"
)

func main() {
	//输入示例
	//go run .\flag.go -loginName xiaowencheng -logPwd=7896542 -age 12 --married=true --delay=1h30m fdafdfad fdadfadfa fafadfadf
	// Args := flag.Args()
	//支持字符串 第一个参数为参数名 第二个为默认值 第三个为提示信息
	loginName := flag.String("loginName", "lixinghua", "登录用户名")
	logPwd := flag.String("logPwd", "123654789", "登录密码")
	age := flag.Int("age", 0, "年龄")
	//是否已婚
	married := flag.Bool("married", true, "是否已经结婚？")
	//时间类型
	delay := flag.Duration("delay", 0, "时间间隔")
	//解析
	flag.Parse()
	fmt.Printf("loginName = %s \n", *loginName)
	fmt.Printf("logPwd = %s \n", *logPwd)
	fmt.Printf("age = %d \n", *age)
	fmt.Printf("married = %v \n", *married)
	fmt.Printf("delay = %v \n", *delay)
	//返回命令行后的其他参数
	args := flag.Args()
	//返回命令行后面的参数格式
	count := flag.NArg()
	//返回使用的命令行参数个数
	countFlag := flag.NFlag()
	fmt.Printf("参数个数有 %d ,参数内容为：%v \n", count, args)
	//返回使用的命令行参数个数
	fmt.Println(countFlag)
}
