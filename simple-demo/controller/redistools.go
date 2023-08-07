// 这个包是redis的工具，其中详细部署了各个操作对应的redis函数

package controller

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"strconv"
)

var rdb *redis.Client

// ************************************* 初始化 ************************************************************************
func Init_Redis() {
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
}

// ************************************* 点赞操作 **********************************************************************

// 点赞操作:将当前用户所点赞的视频唯一id添加到列表当中
func UserPushVideo_id(name string, item string) {
	fmt.Println("正在向", name, "添加", item)
	err := rdb.RPush(name, item)
	if err != nil {
		fmt.Println(err)
	}
}

// 取消点赞操作:将当前用户所取消点赞的视频唯一id在list当中删除
func UserRemvideo_id(name string, item string) {
	err := rdb.LRem(name, 1, item) //因为一个视频只能点赞一次，所以列表当中每个视频id至多有一个
	if err != nil {
		fmt.Println(err)
	}
}

func QueryFavoriteVideos(name string) (res map[int]bool) {
	res = make(map[int]bool)
	vals, err := rdb.LRange(name, 0, -1).Result()
	fmt.Println("喜欢的列表为：", vals)
	if err != nil {
		panic(err)
	}
	for _, val := range vals {
		valnew, err := strconv.Atoi(val)
		if err != nil {
			fmt.Println(err)
		}
		res[valnew] = true
	}
	return res
}
