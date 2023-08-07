package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type FeedResponse struct {
	Response
	VideoList []Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	token := c.Query("token")
	fmt.Println("*****正在执行feed流函数")
	PublishVideos = PublishVideos[:0]
	UpdateVideoList(&PublishVideos, token)
	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: PublishVideos, // 这里传的是PublishVideos
		NextTime:  time.Now().Unix(),
	})

}
