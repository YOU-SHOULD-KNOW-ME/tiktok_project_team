package controller

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB //这个是数据库连接池，用来管理数据库的各种链接

var user User

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type Video struct {
	Id            int64  `json:"id,omitempty"`
	Author        User   `json:"author"`
	PlayUrl       string `json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
	Title         string `json:"video title"`
}

type Comment struct {
	Id         int64  `json:"id,omitempty"`
	User       User   `json:"user"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
}

type User struct {
	Id               int64  `json:"id,omitempty"`             // 用户id
	Name             string `json:"name,omitempty"`           // 用户名称
	FollowCount      int64  `json:"follow_count,omitempty"`   // 被关注总数
	FollowerCount    int64  `json:"follower_count,omitempty"` // 关注总数
	IsFollow         bool   `json:"is_follow,omitempty"`      // 是否关注了这个用户
	Avatar           string `json:"string"`                   // 头像
	Background_image string `json:"background"`               // 首页用户页顶部大图
	Signature        string `json:"signature"`                // 个人简介
	Total_favorited  string `json:"total_favorite"`           // 获赞数量
	Work_count       int64  `json:"work_count"`               // 作品数
	Favorite_count   int64  `json:"favorite_Count"`           // 喜欢数
}

type Message struct {
	Id         int64  `json:"id,omitempty"`
	ToUserId   int64  `json:"to_user_id,reception"`
	FromUserId int64  `json:"from_user_id,from_id"`
	Content    string `json:"content,omitempty"`
	CreateTime int64  `json:"create_time,omitempty"`
}

type MessageSendEvent struct {
	UserId     int64  `json:"user_id,omitempty"`
	ToUserId   int64  `json:"to_user_id,omitempty"`
	MsgContent string `json:"msg_content,omitempty"`
}

type MessagePushEvent struct {
	FromUserId int64  `json:"user_id,omitempty"`
	MsgContent string `json:"msg_content,omitempty"`
}
