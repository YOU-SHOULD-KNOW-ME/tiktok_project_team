// ****
// ****这个包是用来测试mysql数据库的，并且封装了一些对于数据库基本操作的函数
// ****
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id            int64  `json:"id,omitempty"`             //用户id
	Name          string `json:"name,omitempty"`           //用户名称
	FollowCount   int64  `json:"follow_count,omitempty"`   //被关注总数
	FollowerCount int64  `json:"follower_count,omitempty"` // 关注总数
	IsFollow      bool   `json:"is_follow,omitempty"`      // 是否关注了这个用户
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

var Db *sql.DB

// 初始化数据库连接池并连接数据库
func InitDB() (err error) {
	dsn := "root:114108@tcp(116.204.90.201)/tiktok_project"
	Db, err = sql.Open("mysql", dsn) //这个函数不会校验用户名和密码是否正确
	if err != nil {
		fmt.Println(err)
		return
	}
	err = Db.Ping() // 检查用户名和密码是否正确
	if err != nil {
		fmt.Println("连接失败")
		return
	} else {
		fmt.Println("连接成功")
		return
	}
}

func Deleteall() {

	// 清除视频信息
	deletevideo, err := Db.Prepare("delete from video")
	if err != nil {
		fmt.Println(err)
	}
	defer deletevideo.Close()
	_, err = deletevideo.Exec()
	if err != nil {
		fmt.Println(err)
	}

	// 同时清除用户的注册信息
	deletepassword, err := Db.Prepare("delete from password;")
	if err != nil {
		fmt.Println(err)
	}
	defer deletepassword.Close()
	_, err = deletepassword.Exec()
	if err != nil {
		fmt.Println(err)
	}

	// 清除用户信息
	deleteuser, err := Db.Prepare("DELETE FROM user;")
	if err != nil {
		fmt.Println(err)
	}
	defer deleteuser.Close()
	_, err = deleteuser.Exec()
	if err != nil {
		fmt.Println(err)
	}

	// 将视频数量归零
	deletevideocount, err := Db.Prepare("update video_count set count=0;")
	if err != nil {
		fmt.Println(err)
	}
	defer deletevideocount.Close()
	_, err = deletevideocount.Exec()
	if err != nil {
		fmt.Println(err)
	}

	// 将用户数量归零
	deleteusercount, err := Db.Prepare("update user_count set count=0;")
	if err != nil {
		fmt.Println(err)
	}
	defer deleteusercount.Close()
	_, err = deleteusercount.Exec()
	if err != nil {
		fmt.Println(err)
	}

	// 将用户数量归零
	deletecomment, err := Db.Prepare("DELETE FROM comments;")
	if err != nil {
		fmt.Println(err)
	}
	defer deletecomment.Close()
	_, err = deletecomment.Exec()
	if err != nil {
		fmt.Println(err)
	}

	deletefollow, err := Db.Prepare("DELETE FROM follow;")
	if err != nil {
		fmt.Println(err)
	}
	defer deletefollow.Close()
	_, err = deletefollow.Exec()
	if err != nil {
		fmt.Println(err)
	}
	deletemessage, err := Db.Prepare("DELETE FROM message;")
	if err != nil {
		fmt.Println(err)
	}
	defer deletemessage.Close()
	_, err = deletemessage.Exec()
	if err != nil {
		fmt.Println(err)
	}
}

// ********************************************* user 功能区 ***********************************************************

// 查询单个user信息并返回一个user类型
func Queryuserone(id int) (user User) {
	query, err := Db.Prepare("select * from user where id = ?;") //使用预编译语句，防止sql注入
	if err != nil {
		fmt.Println(err)
	}
	defer query.Close()
	err = query.QueryRow(id).Scan(&user.Id, &user.Name, &user.FollowCount, &user.FollowerCount, &user.IsFollow) // 根据key:id来查询user信息并填入user变量中
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user)
	return user
}

// 查询多个user信息并使用循环返回多个user,这里的返回并不是指函数的返回，而是可以通过循环向user列表中逐个假如user
func Queryusermore(id int) {
	query, err := Db.Prepare("select * from user where id = ?;") //定义查询语句
	if err != nil {
		fmt.Println(err)
	}
	rows, err := query.Query(id) // 根据key:id来查询user信息
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close() // 记得关闭连接，要不然别人无法访问数据库
	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Name, &user.FollowCount, &user.FollowerCount, &user.IsFollow) //接收信息
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(user)
	}
}

// 插入数据
func Insertuser(Id int64, name string, FollowCount int64, FollowerCount int64, IsFollow bool) {

	insert, err := Db.Prepare("insert into user(id,name,follow_count,follower_count,is_follow) values(?,?,?,?,?);")
	if err != nil {
		fmt.Println(err)
	}
	defer insert.Close()
	result, err := insert.Exec(Id, name, FollowCount, FollowerCount, IsFollow)
	if err != nil {
		fmt.Println(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(id, "号user插入成功")
}

func Deleteuser(id int) {
	deletevar, err := Db.Prepare("delete from user where id=?")
	if err != nil {
		fmt.Println(err)
	}
	defer deletevar.Close()
	result, err := deletevar.Exec(id)
	if err != nil {
		fmt.Println(err)
	}

	n, err := result.RowsAffected()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)
}

// ********************************************* user 功能区 ***********************************************************

// ********************************************* video 功能区 ***********************************************************

func Insertvideo(id int64, author_id int, play_url string, cover_url string, favorite_count int64, comment_count int64, is_favorite bool, title string) {
	insert, err := Db.Prepare("insert into video(id,author_id,play_url,cover_url,favorite_count, comment_count, is_favorite, title) values(?,?,?,?,?,?,?,?);")
	if err != nil {
		fmt.Println(err)
	}
	defer insert.Close()
	result, err := insert.Exec(id, author_id, play_url, cover_url, favorite_count, comment_count, is_favorite, title)
	if err != nil {
		fmt.Println(err)
	}
	Id, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(Id, "号video插入成功")
}

func Queryvideoone(id int) (v Video) {
	query, err := Db.Prepare("select * from video where id = ?;") //定义查询语句
	if err != nil {
		fmt.Println(err)
	}
	defer query.Close()
	var I int
	query.QueryRow(id).Scan(&v.Id, &I, &v.PlayUrl, &v.CoverUrl, &v.FavoriteCount, &v.CommentCount, &v.IsFavorite, &v.Title) // 根据key:id来查询video信息并填入user变量中
	user := Queryuserone(I)
	v.Author = user
	fmt.Println(v)
	return v
}

// 删除一条指定视频数据
func Deletevideo(id int) {
	deletevar, err := Db.Prepare("delete from video where id=?")
	if err != nil {
		fmt.Println(err)
	}
	defer deletevar.Close()
	result, err := deletevar.Exec(id)
	if err != nil {
		fmt.Println(err)
	}

	n, err := result.RowsAffected()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)
	fmt.Println(id, "号删除成功!")
}

// ********************************************* video 功能区 ***********************************************************

// 更新数据
func Updateage(newage int, id int) {
	update := `update people set age=? where id = ?`
	reseult, err := Db.Exec(update, newage, id)
	if err != nil {
		fmt.Println(err)
		return
	}
	n, err := reseult.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(n)

}

// 更新名字
func Updatename(newname string, id int) {
	update := `update people set name=? where id = ?`
	result, err := Db.Exec(update, newname, id)
	if err != nil {
		fmt.Println(err)
	}
	n, err := result.RowsAffected()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)
}

func QueryIsFollow(id1 int64, id2 int64) (status bool) {
	query, err := Db.Prepare("select * from follow where user1 = ? and user2 = ?;") //定义查询语句
	if err != nil {
		return false
		fmt.Println(err)
	}
	defer query.Close()
	var s string
	err = query.QueryRow(id1, id2).Scan(&s)
	if err == sql.ErrNoRows {
		return false
	} else if err != nil {
		fmt.Println(err)
	}
	return true
}

func main() { //一会要把它放到主函数中
	err := InitDB()
	if err != nil {
		fmt.Println("连接失败")
	}
	Deleteall()
	//Insertvideo(1, 1, "https://graphatk1141087952.oss-cn-beijing.aliyuncs.com/cat.mp4", "https://s3.bmp.ovh/imgs/2023/07/29/6d499421835b1364.jpg", 0, 0, false, "test")
	//Queryvideoone(1)
	//
}
