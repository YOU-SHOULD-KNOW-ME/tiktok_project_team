package controller

var PublishVideos = []Video{
	//{
	//	Id:            1,
	//	Author:        user1,                                                            // 这里是存的是这些视频的作者
	//	PlayUrl:       "https://graphatk1141087952.oss-cn-beijing.aliyuncs.com/cat.mp4", //"https://ugc.kizoa.app/klon1/o399790837_9974871.mp4",
	//	CoverUrl:      "https://s3.bmp.ovh/imgs/2023/07/29/6d499421835b1364.jpg",
	//	FavoriteCount: 9999, //视频点赞数
	//	CommentCount:  0,    //视频评论数
	//	IsFavorite:    true, //这个用户是否已经对这个视频点赞
	//},
	//{
	//	Id:            2,
	//	Author:        DemoUser, // 这里是存的是这些视频的作者
	//	PlayUrl:       "https://graphatk1141087952.oss-cn-beijing.aliyuncs.com/bear.mp4",
	//	CoverUrl:      "https://s3.bmp.ovh/imgs/2023/07/28/9fec015dfc7da741.jpg",
	//	FavoriteCount: 9999,  //视频点赞数
	//	CommentCount:  0,     //视频评论数
	//	IsFavorite:    false, //这个用户是否已经对这个视频点赞
	//},
}

var FavoriteVideos = []Video{ // 喜欢_视频列表

}

// 评论
var DemoComments = []Comment{
	{
		Id:         1,
		User:       DemoUser,
		Content:    "hello,抖音极简版",
		CreateDate: "05-01",
	},
}

var user1 = User{
	Id:            1,
	Name:          "user1",
	FollowCount:   1314,
	FollowerCount: 1000000,
	IsFollow:      true,
}

var DemoUser = User{
	Id:            0,
	Name:          "安铜凯",
	FollowCount:   0,
	FollowerCount: 0,
	IsFollow:      false,
}
