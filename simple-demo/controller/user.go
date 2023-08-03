package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync/atomic"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
var usersLoginInfo = map[string]User{ //存储用户注册信息
	//"user1114108": {
	//	Id:            1,
	//	Name:          "user1",
	//	FollowCount:   0,
	//	FollowerCount: 0,
	//	IsFollow:      true,
	//},
}
var UserVideoList = map[string][]Video{}

var userIdSequence = int64(1) // 用户序号，第几位用户

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User User `json:"user"`
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	fmt.Println("注册用户:", username, "注册密码:", password)
	token := username + password

	if _, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "该账号已存在!"},
		})
	} else {
		atomic.AddInt64(&userIdSequence, 1)                                                                                                                                                                                                //将用户数量加1
		InsertUser(userIdSequence, username, 0, 0, false, "https://graphatk1141087952.oss-cn-beijing.aliyuncs.com/init_avatar.jpg", "https://graphatk1141087952.oss-cn-beijing.aliyuncs.com/init_background.jpg", "hi,tiktok!", "0", 0, 0) //将user信息同步添加到数据库
		AddUserCount(userIdSequence)
		InsertPassword(token, userIdSequence)
		newUser := User{
			Id:               userIdSequence,
			Name:             username,
			Avatar:           "https://graphatk1141087952.oss-cn-beijing.aliyuncs.com/init_avatar.jpg",
			Background_image: "https://graphatk1141087952.oss-cn-beijing.aliyuncs.com/init_background.jpg",
			Signature:        "hi,tiktok!",
		}
		usersLoginInfo[token] = newUser // 将该注册用户添加索引
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   userIdSequence,
			Token:    username + password,
		})
	}
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	fmt.Println("用户名是:", username, "密码是:", password)
	token := username + password

	if user, exist := usersLoginInfo[token]; exist {
		fmt.Println("用户有效，登陆成功!")
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   user.Id,
			Token:    token,
		})
	} else {
		fmt.Println("该账户未注册!")
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "该账户未注册!"},
		})
	}
}

func UserInfo(c *gin.Context) { //这个函数在用户登录完成后调用
	token := c.Query("token")
	fmt.Println("用户状态检测")
	UserVideoList[token] = []Video{}
	QueryToken(token)
	fmt.Println(UserVideoList[token])
	user = usersLoginInfo[token]
	if _, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 0},
			User:     user,
		})

	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "账号检测失败,亲先注册或者登录哦~"},
		})
	}
}
