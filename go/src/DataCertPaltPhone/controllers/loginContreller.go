package controllers

import (
	"DataCertPaltPhone/models"
	"fmt"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}
/*
直接跳转登录页面
 */
func (l *LoginController) Get() {
	l.TplName = "login.html"
}
/*
post犯法处理用户的登录请求
 */
func (l *LoginController) Post() {
	//1、解析客户端用户提交的数据
	var user models.User
	err := l.ParseForm(&user)
	if err != nil {
		l.Ctx.WriteString("抱歉，用户登录信息解析失败")
		return
	}
	//2、根据解析到的数据，执行数据库查询操作
		u,err := user.QueryUser()
	//3、判断数据库查询结果
		if err != nil {
			fmt.Println(err.Error())
			l.Ctx.WriteString("抱歉，用户登录失败")
			return
	}
	//4、根据查询结果返回客户端相应的信息或页面跳转
	l.Data["Phone"] = u.Phone
	l.TplName = "home.html"//上传文件界面
}