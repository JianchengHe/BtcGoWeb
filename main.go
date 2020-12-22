package main

import (
	"BtcGoWeb/btc"
	_ "BtcGoWeb/routers"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
	fmt.Println("hello  world !")
	command, err := btc.GetMsgByCommand("getblock", 0)
	if err ！= nil {
		
	}
	beego.Run()

}

