package main

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/controller"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
)

func main() {
	go service.RunMessageServer()

	err := controller.InitDB()
	if err != nil {
		fmt.Println("连接失败")
	} //初始化数据库连接池，连接云服务器数据库

	controller.Init() // 初始化数据池，其中包括video,user......

	r := gin.Default()

	InitRouter(r)
	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
