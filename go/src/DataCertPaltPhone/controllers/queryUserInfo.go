package controllers

import (
	"DataCertPaltPhone/db_mysql"
	"DataCertPaltPhone/models"
	"fmt"
	"github.com/astaxie/beego"
)

type QyeryUserInfo struct {
	beego.Controller
}
//到数据库查询信息
func (q *QyeryUserInfo)Post() {
	//解析用户信息注册的信息
	var user models.User
	err := q.ParseForm(&user)
	if err != nil {
		q.Ctx.WriteString("数据解析失败，请重试。")
		return
	}
	//到数据上匹配信息
	phone_num,err := db_mysql.PhoneNum(user.Phone)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	users,err := db_mysql.QueryUserIpone(user.Phone)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(users)
	if phone_num > 0 {
		q.TplName = "home.html"
	}
}