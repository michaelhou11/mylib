package main

import (
	"fmt"

	// <----- 这里引入本项目的 utl 包
	// 路径与使用 go mod init初始化时的 module 名称一致
	"github.com/michaelhou11/mylib/pkg/util"
)

func main() {
	fmt.Println("Hello, mylib!")

	// <----- 这里调用本项目的 utl 包的方法
	util.SayHello()
}
