package models

import (
	"BeegoDemo2/util"
	"database/sql"
	"fmt"
	"log"
)

type User struct {
	Id         int
	Username   string
	Password   string
	Status     int // 0 正常状态， 1删除
	Createtime int64
}

//--------------数据库操作-----------------

//插入
func InsertUser(user User) (int64, error) {
	return util.ModifyDB("insert into users(username,password,status,createtime) values (?,?,?,?)",
		user.Username, user.Password, user.Status, user.Createtime)
}

//按条件查询
func QueryUserWightCon(con string) int {
	_sql := fmt.Sprintf("select id from users %s", con)
	fmt.Println(_sql)
	id := 0
	err := util.QueryRowDB(_sql).Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows { //如果数据不存在会有这个err
			log.Println("There is not row")
		} else {
			log.Println(err)
		}
	}
	return id
}

//根据用户名查询id
func QueryUserWithUsername(username string) int {
	sql := fmt.Sprintf("where username='%s'", username)
	return QueryUserWightCon(sql)
}

//根据用户名和密码，查询id
func QueryUserWithParam(username, password string) int {
	sql := fmt.Sprintf("where username='%s' and password='%s'", username, password)
	return QueryUserWightCon(sql)
}

//查询符合条件的全部用户数据
func QueryUsersWightCon(con string) []string {
	_sql := fmt.Sprintf("select username from users %s", con)
	fmt.Println(_sql)

	result := []string{} //切片
	rows, err := util.Query(_sql)
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close() //记得关闭游标

	for rows.Next() {
		var s string
		err = rows.Scan(&s)
		if err != nil {
			log.Fatalln(err)
		}
		log.Printf("found row containing %q", s)
		result = append(result, s)
	}
	rows.Close()
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return result

}
