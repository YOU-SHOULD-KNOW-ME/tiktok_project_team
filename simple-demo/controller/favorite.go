package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	token := c.Query("token")
	video_id := c.Query("video_id")
	action_type := c.Query("action_type") // 1是点赞，2是取消点赞
	fmt.Println("该视频的唯一视频号为:", video_id)
	fmt.Println(action_type)

	if _, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, Response{StatusCode: 0})
		if action_type == "1" {
			UserPushVideo_id(token, video_id)
			v_id, err := strconv.Atoi(video_id)
			if err != nil {
				fmt.Println(err)
			}
			v := QueryVideoOne(v_id)

			FavoriteVideos = append(FavoriteVideos, v)
			for i := range PublishVideos {
				if PublishVideos[i].Id == int64(v_id) {
					PublishVideos[i].FavoriteCount++
					p := PublishVideos[i].Author.Total_favorited
					p1, er := strconv.Atoi(p)
					if er != nil {
						fmt.Println(er)
					}
					p1++
					p2 := strconv.Itoa(p1)
					PublishVideos[i].Author.Total_favorited = p2
					Updateuser_total_favorite(PublishVideos[i].Author.Id, int64(p1))
					PublishVideos[i].IsFavorite = true
					break
				}
			}

			// 将当前用户的喜欢数加1
			users := usersLoginInfo[token]
			users.Favorite_count++
			usersLoginInfo[token] = users
			Updateuser_favorite_count(usersLoginInfo[token].Id, usersLoginInfo[token].Favorite_count)
			// 将目标视频的用户获赞数加1
			u := v.Author
			query_token := Userid_Query_Token(int(u.Id))
			user1 := usersLoginInfo[query_token]
			user1.Favorite_count++
			usersLoginInfo[query_token] = user1
			UpdateVideoFavoriteCount(int64(v_id), v.FavoriteCount+1)
		} else if action_type == "2" {
			UserRemvideo_id(token, video_id)
			v_id, err := strconv.Atoi(video_id)
			if err != nil {
				fmt.Println(err)
			}
			for i, v := range FavoriteVideos {
				if v.Id == int64(v_id) {
					FavoriteVideos = append(FavoriteVideos[:i], FavoriteVideos[i+1:]...)
					break
				}
			}
			// 查找feed流中的对应视频，将对应视频的获赞数减1，并且将是否点赞变为false
			for i := range PublishVideos {
				if PublishVideos[i].Id == int64(v_id) {
					PublishVideos[i].FavoriteCount--
					p := PublishVideos[i].Author.Total_favorited
					p1, er := strconv.Atoi(p)
					if er != nil {
						fmt.Println(er)
					}
					p1--
					p2 := strconv.Itoa(p1)
					PublishVideos[i].Author.Total_favorited = p2
					Updateuser_total_favorite(PublishVideos[i].Author.Id, int64(p1))
					PublishVideos[i].IsFavorite = false
					break
				}
			}
			users := usersLoginInfo[token]
			users.Favorite_count--
			usersLoginInfo[token] = users
			Updateuser_favorite_count(usersLoginInfo[token].Id, usersLoginInfo[token].Favorite_count)

			v := QueryVideoOne(v_id)
			// 将目标视频的用户获赞数减1
			u := v.Author
			query_token := Userid_Query_Token(int(u.Id))
			user1 := usersLoginInfo[query_token]
			user1.Favorite_count--
			usersLoginInfo[query_token] = user1
			UpdateVideoFavoriteCount(int64(v_id), v.FavoriteCount-1)
		}

	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	fmt.Println("喜爱的")
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: FavoriteVideos,
	})
}
