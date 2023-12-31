package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type CommentListResponse struct {
	Response
	CommentList []Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	Response
	Comment Comment `json:"comment,omitempty"`
}

//func index(str string) (res string) {
//	for k := range Find{
//
//	}
//}

func index(s string) (res string) {
	M := make(map[string]bool)
	M["nb"] = true
	M["牛逼"] = true
	M["tmd"] = true
	M["tm"] = true
	M["cnm"] = true
	M["nmd"] = true
	M["操你妈"] = true
	M["nmd"] = true
	for k := range M {
		i := strings.Index(s, k)
		if i != -1 {
			s = strings.Replace(s, k, "****", -1)
		}
	}
	return s
}

// CommentAction has practical effect, and check if token is valid
// CommentAction has practical effect, and check if token is valid
func CommentAction(c *gin.Context) {
	token := c.Query("token")
	actionType := c.Query("action_type")
	if user, exist := usersLoginInfo[token]; exist {
		if actionType == "1" {
			text := c.Query("comment_text")
			text = index(text)
			videoId, _ := strconv.Atoi(c.Query("video_id"))
			var num int64
			for i := range PublishVideos {
				if PublishVideos[i].Id == int64(videoId) {
					PublishVideos[i].CommentCount++
					num = PublishVideos[i].CommentCount
				}
			}
			createTime := time.Now().Format("2006-01-02 15:04:05")
			//数据库增加评论
			commentId := Insertcomment(int(user.Id), videoId, text, createTime)
			UpdateVideocommentcount(int64(videoId), num)
			videoId, _ = strconv.Atoi(c.Query("video_id"))
			createTime = time.Now().Format("2006-01-02 15:04:05")
			c.JSON(http.StatusOK, CommentActionResponse{Response: Response{StatusCode: 0},
				Comment: Comment{
					Id:         commentId,
					User:       user,
					Content:    text,
					CreateDate: createTime,
				}})
			return
		} else if actionType == "2" {
			//数据库删除对应评论ID的ID
			commentId, _ := strconv.Atoi(c.Query("comment_id"))
			videoId, _ := strconv.Atoi(c.Query("video_id"))
			var num int64
			for i := range PublishVideos {
				if PublishVideos[i].Id == int64(videoId) {
					PublishVideos[i].CommentCount--
					num = PublishVideos[i].CommentCount
				}
			}
			Deletecomment(commentId, videoId)
			UpdateVideocommentcount(int64(videoId), num)
		} else if actionType == "2" {
			//数据库删除对应评论ID的ID
			commentId, _ := strconv.Atoi(c.Query("comment_id"))
			videoId, _ := strconv.Atoi(c.Query("video_id"))
			Deletecomment(commentId, videoId)
			c.JSON(http.StatusOK, Response{StatusCode: 0})
		}
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	/*token := c.Query("token")
	if user,exist := usersLoginInfo[token]; exist{
		fmt.Print(user)
		videoId, _ := strconv.Atoi(c.Query("video_id"))
		var list []Comment
		Commentlist(videoId,&list)
		c.JSON(http.StatusOK, CommentListResponse{
			Response:    Response{StatusCode: 0},
			CommentList: list,
		})
	}*/
	videoId, _ := strconv.Atoi(c.Query("video_id"))
	var list []Comment
	Commentlist(videoId, &list)
	c.JSON(http.StatusOK, CommentListResponse{
		Response:    Response{StatusCode: 0},
		CommentList: list,
	})
}
