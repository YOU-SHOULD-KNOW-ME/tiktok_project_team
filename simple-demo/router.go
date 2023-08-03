package main

import (
	"github.com/RaymondCode/simple-demo/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {

	// public directory is used to serve static resources
	r.Static("/static", "./public")
	//视频接口,在你的url后面加上/static/视频文件名字.mp4 就可以访问视频

	apiRouter := r.Group("/douyin")
	//Group 是 gin.Engine 类型的一个方法，它用于创建一个新的路由组。路由组是一种组织共享相同中间件或具有相同路径前缀的路由的方式。例如，您可以将所有需要授权中间件的路由分组在一起。这样可以更容易地管理和维护您的路由。
	//在您提供的代码中，apiRouter := r.Group("/douyin") 这一行使用 Group 方法创建了一个新的路由组，其 URL 路径前缀为 /douyin。这意味着您可以在 apiRouter 路由组中添加以 /douyin 开头的路由。例如，您可以使用 apiRouter.GET("/example", exampleHandler) 来添加一个处理 /douyin/example URL 的路由。

	// basic apis
	apiRouter.GET("/feed/", controller.Feed)
	apiRouter.GET("/user/", controller.UserInfo)
	apiRouter.POST("/user/register/", controller.Register)
	apiRouter.POST("/user/login/", controller.Login)
	apiRouter.POST("/publish/action/", controller.Publish)
	apiRouter.GET("/publish/list/", controller.PublishList)

	// extra apis - I
	apiRouter.POST("/favorite/action/", controller.FavoriteAction)
	apiRouter.GET("/favorite/list/", controller.FavoriteList)
	apiRouter.POST("/comment/action/", controller.CommentAction)
	apiRouter.GET("/comment/list/", controller.CommentList)

	// extra apis - II
	apiRouter.POST("/relation/action/", controller.RelationAction)
	apiRouter.GET("/relation/follow/list/", controller.FollowList)
	apiRouter.GET("/relation/follower/list/", controller.FollowerList)
	apiRouter.GET("/relation/friend/list/", controller.FriendList)
	apiRouter.GET("/message/chat/", controller.MessageChat)
	apiRouter.POST("/message/action/", controller.MessageAction)
}
