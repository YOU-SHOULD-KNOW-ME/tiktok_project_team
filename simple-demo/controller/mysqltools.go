// ** 这个包存放连接数据库以及管理数据库的函数工具 **
// ** update time : 2023-08-02 **
// ** 更新信息:更新user与video的sql管控工具函数 **
// ** update time : 2023-08-02 **
// ** 更新信息:更新编译预处理，防止sql注入攻击 **
package controller

import (
	"database/sql"
	"fmt"
)

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

// 开始
// ********************************************* user_count 和 video_count 功能区 ***********************************************************

// 查询并同步已注册用户数量
func QueryUserCount() {
	query, err := Db.Prepare("select * from user_count;")
	if err != nil {
		fmt.Println(err)
		fmt.Println("用户数量同步失败，请检查用户数量同步函数")
	}
	defer query.Close()
	err = query.QueryRow().Scan(&userIdSequence) //将查询到的用户数量同步至本地
	if err != nil {
		fmt.Println(err)
	}
}

// 查询并同步视频数量
func QueryVideoCount() {
	query, err := Db.Prepare("select * from video_count;")
	if err != nil {
		fmt.Println(err)
		fmt.Println("视频数量同步失败,请检查视频同步函数")
	}
	defer query.Close()
	err = query.QueryRow().Scan(&Id)
	if err != nil {
		fmt.Println(err)
	}
}

func AddUserCount(n int64) {
	update, err := Db.Prepare("update user_count set count=?;")
	if err != nil {
		fmt.Println(err)
	}
	defer update.Close()
	_, err = update.Exec(n)
	if err != nil {
		fmt.Println(err)
	}
}

// 将视频数量+1同步到数据库
func AddVideoCount(n int64) {
	update, err := Db.Prepare("update video_count set count=?;")
	if err != nil {
		fmt.Println(err)
	}
	defer update.Close()
	_, err = update.Exec(n)
	if err != nil {
		fmt.Println(err)
	}
}

func Updatework_count(id int64, n int64) {
	update, err := Db.Prepare("update user set work_count=? where id=?;")
	if err != nil {
		fmt.Println(err)
	}
	defer update.Close()
	_, err = update.Exec(n, id)
	if err != nil {
		fmt.Println(err)
	}
}

// ********************************************* user_count 和 video_count 功能区 ***********************************************************
// 结束

// 开始
// ********************************************* user 功能区 ***********************************************************

// 查询单个user信息并返回一个user类型
func QueryUserOne(id int) (user User) {
	query, err := Db.Prepare("select * from user where id = ?;") //使用预编译语句，防止sql注入
	if err != nil {
		fmt.Println(err)
	}
	defer query.Close()
	err = query.QueryRow(id).Scan(&user.Id, &user.Name, &user.FollowCount, &user.FollowerCount, &user.IsFollow, &user.Avatar, &user.Background_image, &user.Signature, &user.Total_favorited, &user.Work_count, &user.Favorite_count) // 根据key:id来查询user信息并填入user变量中
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user)
	return user
}

func QueryUserOne_init(id int, user *User) {
	query, err := Db.Prepare("select * from user where id = ?;") //使用预编译语句，防止sql注入
	if err != nil {
		fmt.Println(err)
	}
	defer query.Close()
	err = query.QueryRow(id).Scan(&user.Id, &user.Name, &user.FollowCount, &user.FollowerCount, &user.IsFollow, &user.Avatar, &user.Background_image, &user.Signature, &user.Total_favorited, &user.Work_count, &user.Favorite_count) // 根据key:id来查询user信息并填入user变量中
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user)
}

// 查询多个user信息并使用循环返回多个user,这里的返回并不是指函数的返回，而是可以通过循环向user列表中逐个加入user
func QueryUserMore() {
	query, err := Db.Prepare("select * from password;") //定义查询语句
	if err != nil {
		fmt.Println(err)
	}
	rows, err := query.Query() // 根据key:id来查询user信息
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close() // 记得关闭连接，要不然别人无法访问数据库
	for rows.Next() {
		var token string
		var id int
		err := rows.Scan(&token, &id) //接收信息
		if err != nil {
			fmt.Println(err)
		}
		user := QueryUserOne(id)
		usersLoginInfo[token] = user
	}
}

// 插入数据
func InsertUser(Id int64, name string, FollowCount int64, FollowerCount int64, IsFollow bool, Avatar string, Background_image string, signature string, Total_favorited string, Work_count int64, favorite_count int64) {

	insert, err := Db.Prepare("insert into user(id,name,follow_count,follower_count,is_follow,avatar,background_image,signature,total_favorited,work_count,favorite_count) values(?,?,?,?,?,?,?,?,?,?,?);")
	if err != nil {
		fmt.Println(err)
	}
	defer insert.Close()
	result, err := insert.Exec(Id, name, FollowCount, FollowerCount, IsFollow, Avatar, Background_image, signature, Total_favorited, Work_count, favorite_count)
	if err != nil {
		fmt.Println(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(id, "号user插入成功")
}

func InsertPassword(password string, user_id int64) {
	insert, err := Db.Prepare("insert into password(password,user_id) value(?,?);")
	if err != nil {
		fmt.Println(err)
	}
	defer insert.Close()
	_, err = insert.Exec(password, user_id)
	if err != nil {
		fmt.Println(err)
	}
}

func DeleteUser(id int) {
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

func Updateuser_(newage int, id int) {
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

// ********************************************* user 功能区 ***********************************************************
// 结束

// 开始
// ********************************************* video 功能区 ***********************************************************

func InsertVideo(id int64, author_id int, play_url string, cover_url string, favorite_count int64, comment_count int64, is_favorite bool, title string) {
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

func QueryVideoOne(id int) (v Video) {
	query, err := Db.Prepare("select * from video where id = ?;") //定义查询语句
	if err != nil {
		fmt.Println(err)
	}
	defer query.Close()
	var I int
	query.QueryRow(id).Scan(&v.Id, &I, &v.PlayUrl, &v.CoverUrl, &v.FavoriteCount, &v.CommentCount, &v.IsFavorite, &v.Title) // 根据key:id来查询video信息并填入user变量中
	user := QueryUserOne(I)
	v.Author = user
	fmt.Println(v)
	return v
}

func QueryVideoMore(list *[]Video) {
	query, err := Db.Prepare("select * from video where id > 0;") //定义查询语句
	if err != nil {
		fmt.Println(err)
	}
	rows, err := query.Query() // 根据key:id来查询user信息
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close() // 记得关闭连接，要不然别人无法访问数据库
	for rows.Next() {
		var V Video
		var I int
		err := rows.Scan(&V.Id, &I, &V.PlayUrl, &V.CoverUrl, &V.FavoriteCount, &V.CommentCount, &V.IsFavorite, &V.Title) //接收信息
		if err != nil {
			fmt.Println(err)
		}
		user := QueryUserOne(I)
		V.Author = user
		fmt.Println(V)
		*list = append(*list, V)
	}
}

func QueryToken(token string) {
	query, err := Db.Prepare("select user_id from password where password = ?")
	if err != nil {
		fmt.Println(err)
	}
	defer query.Close()

	var user_id int64
	err = query.QueryRow(token).Scan(&user_id) // 根据key:id来查询user信息
	if err != nil {
		fmt.Println(err)
	}
	query, err = Db.Prepare("select * from video where author_id=?")
	if err != nil {
		fmt.Println(err)
	}
	defer query.Close()
	rows, err := query.Query(user_id)
	defer rows.Close() // 记得关闭连接，要不然别人无法访问数据库
	for rows.Next() {
		var v Video
		var I int
		err := rows.Scan(&v.Id, &I, &v.PlayUrl, &v.CoverUrl, &v.FavoriteCount, &v.CommentCount, &v.IsFavorite, &v.Title) //接收信息
		if err != nil {
			fmt.Println(err)
		}
		QueryUserOne_init(I, &user)
		v.Author = user
		UserVideoList[token] = append(UserVideoList[token], v)
	}
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
	fmt.Println(n, "号删除成功!")
}

// ********************************************* video 功能区 ***********************************************************
// 结束

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
