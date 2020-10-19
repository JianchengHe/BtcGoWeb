package main

import (
	"DataCertPaltPhone/blockchain"
	"DataCertPaltPhone/db_mysql"
	_ "DataCertPaltPhone/routers"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
	block0 := blockchain.CreateGenesisBlock()
	block1 := blockchain.NewBlock(block0.Height+1,block0.Hash,[]byte("a"))
	fmt.Println(block1)

	//连接数据库
	db_mysql.Connect()
	//静态资源文件路径映设设置
	beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/css","./static/css")
	beego.SetStaticPath("/img","./static/img")
	beego.Run()
}

