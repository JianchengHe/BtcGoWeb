package main

import (
	"DataCertPaltPhone/blockchain"
	"DataCertPaltPhone/db_mysql"
	_ "DataCertPaltPhone/routers"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
	//1、创世区块
	bc := blockchain.NewBlockChain()//封装
	fmt.Printf("创世区块的哈希值:%x\n",bc.LastHash)
	//bc.SaveData([]byte("用户的数据"))

	//block0 := blockchain.CreateGenesisBlock() //创建创世区块
	////block1 := blockchain.NewBlock(block0.Height+1,block0.Hash,[]byte("a"))
	////fmt.Println(block1)
	//block1 := blockchain.NewBlock(
	//	block0.Height+1,
	//	block0.Hash,
	//	[]byte{})
	//fmt.Printf("block0的哈希：%x\n", block0.Hash)
	//fmt.Printf("block1的哈希：%x\n", block1.Hash)
	//fmt.Printf("block1的PrevHash哈希：%x\n", block1.PerviousHash)
	//
	//block0Bytes := block0.Serialize()
	//fmt.Println("创世区块gob序列化后：",block0Bytes)
	//deBlock0,err := blockchain.DeSerialize(block0Bytes)
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
	//fmt.Println("反序列化后的区块的高度是",deBlock0.Height)
	//fmt.Printf("反序列化的区块的哈希:%x\n ",deBlock0.Hash)
	////序列化Marshal,只有序列化之后的数据才能传输;反序列化Unmarshal
	//blockJsonBytes,_ := json.Marshal(block0)
	//fmt.Println("通过json序列化以后的block:",string(blockJsonBytes))
	//
	//blockXml,_ := xml.Marshal(block0)
	//fmt.Println("通过xml序列化的block：",string(blockXml))
	return
	//连接数据库
	db_mysql.Connect()
	//静态资源文件路径映设设置
	beego.SetStaticPath("/js", "./static/js")
	beego.SetStaticPath("/css", "./static/css")
	beego.SetStaticPath("/img", "./static/img")
	beego.Run()
}
