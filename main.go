package main

import (
	_ "BTCBeego/routers"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
	fmt.Println("hello this world")

	//静态资源映射文件
	beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/img","./static/img")
	beego.SetStaticPath("/css","./static/css")

	beego.Run()

}

