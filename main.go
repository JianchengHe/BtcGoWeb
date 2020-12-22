package main

import (
	_ "BTCBeego/routers"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
	fmt.Println("hello this world")
	beego.Run()

}

