// ****
// **** 这个库是用来测试redis数据库的,并且封装了一些对于redis基本操作的函数
// ****

package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"strconv"
)

var rdb *redis.Client

// 向列表中添加元素
func RPush(name string, item int) {
	err := rdb.RPush(name, strconv.Itoa(item))
	if err != nil {
		fmt.Println(err)
	}
}

// 删除列表中的指定元素
func LRem(name string, item int) {
	err := rdb.LRem(name, 1, strconv.Itoa(item))
	if err != nil {
		fmt.Println(err)
	}

}

func main() {
	// 创建一个新的客户端实例
	rdb = redis.NewClient(&redis.Options{
		Addr: "116.204.90.201:6379", // 用您的云服务器IP地址和端口号替换
		DB:   0,                     // 使用默认数据库
	})
	ctx := context.Background()
	rdb = rdb.WithContext(ctx)
	// 测试连接
	pong, err := rdb.Ping().Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("连接成功", pong)
	//for i := 0; i < 3; i++ {
	//	RPush("key", i)
	//}
	RPush("user114108", 1)

}
