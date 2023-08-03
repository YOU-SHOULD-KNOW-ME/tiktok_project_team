package controller

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/controller/cover"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"path/filepath"
	"strconv"
)

var Id int64 // 记录视频数量

type VideoListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	token := c.PostForm("token")
	title := c.PostForm("title")
	if _, exist := usersLoginInfo[token]; !exist {
		fmt.Println("没登陆")
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "请先登录~"})
		return
	}

	data, err := c.FormFile("data") // 接受视频数据
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	filename := filepath.Base(data.Filename)
	user = usersLoginInfo[token]
	finalName := fmt.Sprintf("%d_%s", user.Id, filename)
	fmt.Println(user.Id)

	saveFile := fmt.Sprintf("D:\\GOLAND\\gocode\\simple-demo\\controller\\video\\", finalName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	Id++
	photoname := cover.Get_cover(saveFile, token+strconv.FormatInt(Id, 10), user.Id)
	fmt.Println("*************" + photoname)
	file, err := data.Open()
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	defer file.Close()
	fmt.Println(finalName)
	ossvedio(finalName, file)
	var Newvideo Video
	Newvideo = Video{Id, user, "https://graphatk1141087952.oss-cn-beijing.aliyuncs.com/" + finalName, "https://graphatk1141087952.oss-cn-beijing.aliyuncs.com/" + photoname, 0, 0, false, title}
	InsertVideo(Id, int(user.Id), "https://graphatk1141087952.oss-cn-beijing.aliyuncs.com/"+finalName, "https://graphatk1141087952.oss-cn-beijing.aliyuncs.com/"+photoname, 0, 0, false, title)
	AddVideoCount(Id)
	user.Work_count++
	user1 := usersLoginInfo[token]
	user1.Work_count = user.Work_count
	usersLoginInfo[token] = user1

	Newvideo.Author.Work_count++
	fmt.Println("当前用户信息为:", user)
	go Updatework_count(Id, user.Work_count)
	PublishVideos = append(PublishVideos, Newvideo) //这里把投稿的视频信息添加到videolist中
	UserVideoList[token] = append(UserVideoList[token], Newvideo)
	//saveFile := filepath.Join("./public/", finalName)
	//if err := c.SaveUploadedFile(data, saveFile); err != nil {
	//	c.JSON(http.StatusOK, Response{
	//		StatusCode: 1,
	//		StatusMsg:  err.Error(),
	//	})
	//	return
	//}

	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) { // 用户发布过的视频

	token := c.Query("token")
	fmt.Println(token)
	if _, exist := usersLoginInfo[token]; exist {

	} else {
		fmt.Println("账号不存在,请重新登录")
	}
	fmt.Println("用户投稿列表")
	fmt.Println(UserVideoList[token])
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: UserVideoList[token], //这里传的也是PublishVideos
	})
}
