package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sort"
	"strconv"
	"sync/atomic"
	"time"
)

var test = []Message{}

var tempChat = map[string][]Message{}      // 这里面存的是最新的没有显示的消息
var oldChat = map[string]map[string]bool{} // 这里面存的是当前用户

var messageIdSequence = int64(0)

type ChatResponse struct {
	Response
	MessageList []Message `json:"message_list"`
}

// MessageAction no practical effect, just check if token is valid
func MessageAction(c *gin.Context) {
	fmt.Println("massageaction")

	token := c.Query("token")
	toUserId := c.Query("to_user_id")
	action_type := c.Query("action_type")
	content := c.Query("content")
	fmt.Println(action_type)
	fmt.Println(content)
	if users, exist := usersLoginInfo[token]; exist {
		userIdB, _ := strconv.Atoi(toUserId)
		chatKey := genChatKey(users.Id, int64(userIdB))
		atomic.AddInt64(&messageIdSequence, 1)
		time := time.Now().Unix()
		//curMessage := Message{
		//	Content:    content,
		//	CreateTime: time,
		//	FromUserId: users.Id,
		//	Id:         messageIdSequence,
		//	ToUserId:   int64(userIdB),
		//}
		InsertMessage(messageIdSequence, users.Id, int64(userIdB), content, time)
		if _, exists := tempChat[chatKey]; exists {

		} else {
			tempChat[chatKey] = []Message{}
		}
		c.JSON(http.StatusOK, Response{StatusCode: 0})
	} else {
		fmt.Println("action失败")
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})

	}
}

// MessageChat all users have same follow list
func MessageChat(c *gin.Context) {
	fmt.Println("messagechat")
	token := c.Query("token")
	toUserId := c.Query("to_user_id")

	if users, exist := usersLoginInfo[token]; exist {
		fmt.Println("成功进入")
		userIdB, _ := strconv.Atoi(toUserId)
		chatKey := genChatKey(users.Id, int64(userIdB))
		sort.Slice(tempChat[chatKey], func(i, j int) bool {
			return tempChat[chatKey][i].CreateTime < tempChat[chatKey][j].CreateTime
		})
		fmt.Println("聊天列表:", tempChat[chatKey])
		c.JSON(http.StatusOK, ChatResponse{Response: Response{StatusCode: 0}, MessageList: tempChat[chatKey]})
		tempChat[chatKey] = []Message{}
	} else {
		fmt.Println("请求失败")
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})

	}
}

func genChatKey(userIdA int64, userIdB int64) string {
	if userIdA > userIdB {
		return fmt.Sprintf("%d_%d", userIdB, userIdA)
	}
	return fmt.Sprintf("%d_%d", userIdA, userIdB)
}
