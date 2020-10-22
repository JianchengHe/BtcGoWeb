package main

import (
	"DataCertPaltPhone/blockchain"
	"DataCertPaltPhone/db_mysql"
	_ "DataCertPaltPhone/routers"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
	block0 := blockchain.CreateGenesisBlock() //创建创世区块
	//block1 := blockchain.NewBlock(block0.Height+1,block0.Hash,[]byte("a"))
	//fmt.Println(block1)
	block1 := blockchain.NewBlock(
		block0.Height+1,
		block0.Hash,
		[]byte{})
	fmt.Printf("block0的哈希：%x\n", block0.Hash)
	fmt.Printf("block1的哈希：%x\n", block1.Hash)
	fmt.Printf("block1的PrevHash哈希：%x\n", block1.PerviousHash)

	//序列化Marshal,只有序列化之后的数据才能传输;反序列化Unmarshal
	blockJsonBytes,_ := json.Marshal(block0)
	fmt.Println("通过json序列化以后的block:",string(blockJsonBytes))

	blockXml,_ := xml.Marshal(block0)
	fmt.Println("通过xml序列化的block：",string(blockXml))



	return
	//连接数据库
	db_mysql.Connect()
	//静态资源文件路径映设设置
	beego.SetStaticPath("/js", "./static/js")
	beego.SetStaticPath("/css", "./static/css")
	beego.SetStaticPath("/img", "./static/img")
	beego.Run()
}
