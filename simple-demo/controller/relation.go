package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var Friend_List []User

type UserListResponse struct {
	Response
	UserList []User `json:"user_list"`
}

func removeElement(slice []User, element User) []User {
	for i, v := range slice {
		if v.Id == element.Id {
			print("找到了")
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	token := c.Query("token")        // 获取的是当前用户的token
	user_id := c.Query("to_user_id") // 获取的是关注方用户的id
	action_type := c.Query("action_type")

	fmt.Print(token, user_id)
	u := usersLoginInfo[token]
	id1 := u.Id                       // 当前用户的id
	id2, err := strconv.Atoi(user_id) // 目标用户的id
	if err != nil {
		fmt.Println(err)
	}

	people := QueryUserOne(id2)
	if action_type == "1" {
		for i, v := range Follower_List {
			if v.Id == int64(id2) {
				Follower_List[i].IsFollow = true
			}
		}
		for i, v := range PublishVideos {
			if v.Author.Id == int64(id2) {
				PublishVideos[i].Author.IsFollow = true
			}
		}
		InsertFollow(id1, int64(id2))
		people.IsFollow = true
		Follow_List = append(Follow_List, people)
		UpdateFollowCount(id1, int64(len(Follow_List)))
		u1 := usersLoginInfo[token]
		u1.FollowCount++
		usersLoginInfo[token] = u1
		token2 := Userid_Query_Token(id2)
		u2 := usersLoginInfo[token2]
		u2.FollowerCount++
		usersLoginInfo[token2] = u2
		UpdateFollowerCount(int64(id2), people.FollowerCount+1)
	} else {
		for i, v := range Follower_List {
			if v.Id == int64(id2) {
				Follower_List[i].IsFollow = false
			}
		}
		for i, v := range PublishVideos {
			if v.Author.Id == int64(id2) {
				PublishVideos[i].Author.IsFollow = false
			}
		}
		DeleteFollow(id1, int64(id2))
		fmt.Println(people)
		Follow_List = removeElement(Follow_List, people)
		UpdateFollowCount(id1, int64(len(Follow_List)))
		UpdateFollowerCount(int64(id2), people.FollowerCount-1)
		u1 := usersLoginInfo[token]
		u1.FollowCount--
		usersLoginInfo[token] = u1
		token2 := Userid_Query_Token(id2)
		u2 := usersLoginInfo[token2]
		u2.FollowerCount--
		usersLoginInfo[token2] = u2
	}

	if _, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

func FollowList(c *gin.Context) {
	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: Follow_List,
	})
}

func FollowerList(c *gin.Context) {
	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: Follower_List,
	})
}

func FriendList(c *gin.Context) {
	Friend_List = []User{}
	fmt.Println("好友列表")
	user_id := c.Query("user_id")
	id2, err := strconv.Atoi(user_id)
	if err != nil {
		fmt.Println(err)
	}
	for _, u := range Follow_List {
		status := QueryIsFollow(u.Id, int64(id2))
		if status {
			Friend_List = append(Friend_List, u)
		}
	}
	for _, u := range Friend_List {
		chatKey := genChatKey(int64(id2), u.Id)
		tempChat[chatKey] = []Message{}
		QueryMessage(int64(id2), u.Id, chatKey)
		QueryMessage(u.Id, int64(id2), chatKey)
	}
	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: Friend_List,
	})
}
