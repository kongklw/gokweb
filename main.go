package main

import (
	// "gokweb/dao"
	"gokweb/router"
)

func main() {
	// dao.Init()
	r := router.Router()

	r.Run() // listen and serve on 0.0.0.0:8080
	// r.Run(":8000") // set self define port
	//window go build run in linux refer https://developer.aliyun.com/article/873494

	// 先运行下列命令window环境下打包可以在linux中运行的go 文件
	// set GOARCH=amd64
	// go env -w GOARCH=amd64
	// set GOOS=linux
	// go env -w GOOS=linux
	// go build

	// 打包后再重新设置回来
	// go env -w GOARCH=amd64
	// go env -w GOOS=windows

}
