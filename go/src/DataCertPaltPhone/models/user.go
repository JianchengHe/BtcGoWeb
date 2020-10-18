package models

import (
	"DataCertPaltPhone/db_mysql"
	"DataCertPaltPhone/utils"
	"fmt"
)

type User struct {
	Id       int    `form:"id"`
	Phone    string `form:"phone"`
	Password string `form:"password"`
}

//保存用户的方法
/*
将用户的信息保存到数据库中
*/
//写法二
func (u User) AddUser() (int64, error) {
	//脱敏
	u.Password = utils.Md5HashString(u.Password)
	result, err := db_mysql.Db.Exec("insert into user (phone,password)values(?,?) ",
		u.Phone, u.Password)
	if err != nil { //保存数据遇到错误
		return -1, err
	}
	id, err := result.RowsAffected()
	if err != nil { //保存数据遇到错误
		return -1, err
	}
	//id值代表的是此次数据影响得到行数，id是一个整数的int64类型
	return id, nil
}

/*
查询用户信息
*/
func (u User) QueryUser() (*User, error) {
	//脱敏
	u.Password = utils.Md5HashString(u.Password)
	row := db_mysql.Db.QueryRow("select phone from user where phone = ? and password = ? ",
		u.Phone, u.Password)
	err := row.Scan(&u.Phone)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

//
func (u User) QueryUserByPhone() (*User, error) {
	fmt.Println("用户查询的phone信息：", u.Phone)
	row := db_mysql.Db.QueryRow("select id from user where phone = ?", u.Phone)
	var user User
	err := row.Scan(&user.Id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
