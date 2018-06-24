package main

import (
	_ "sanqi/blog/blogweb/routers"
	"github.com/astaxie/beego"
	"sanqi/blog/blogweb/utils"
	"fmt"
)

func main() {
	utils.InitMysql()
	fmt.Println("链接成功")
	beego.Run()
}

