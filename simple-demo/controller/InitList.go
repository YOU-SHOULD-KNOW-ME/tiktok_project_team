// 这个包用来初始化所有数据当用户开启软件时

package controller

func InitPublishVideos() {
	QueryVideoMore(&PublishVideos) // 这个函数使用来同步数据库中的video信息，在初始化函数中，用来同步数据库中已有的的视频信息到视频列表中
}

func InitUsersLoginInfo() { // 同步所有已注册用户信息
	QueryUserMore() // 这个函数来同步所有用户的注册信息，填入userLoginInfo
}

func Init() {
	QueryUserCount()     // 启动自动同步已有用户数量
	QueryVideoCount()    // 启动自动同步已有视频数量
	InitUsersLoginInfo() // 启动自动同步用户注册信息
	InitPublishVideos()  // 启动自动同步视频列表
}
