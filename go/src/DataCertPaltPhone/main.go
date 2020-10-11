package main

import (
	"DataCertPaltPhone/db_mysql"
	_ "DataCertPaltPhone/routers"
	"github.com/astaxie/beego"
)

func main() {
	//静态资源文件路径映设设置
	beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/css","./static/css")
	beego.SetStaticPath("/img","./static/img")
	beego.Run()
	//连接数据库
	db_mysql.Connect()


}

