package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
)

type BaseController struct {
	beego.Controller
	IsLogin bool
	LoginUser interface{}
}

func (this *BaseController)Prepare()  {  //调用Get和post之前调用,调用之前先确认有没有session
	loginuesr:=this.GetSession("loginuesr")  //获取session ,sessionname在登录的时候有设定
	fmt.Println("######",loginuesr)
	if loginuesr!=nil{  //判断Session是否存在  是否登录, 如果登录
		this.IsLogin=true   //变为true
		this.LoginUser=loginuesr   //用户名
	}else {
		this.IsLogin=false

	}
	this.Data["IsLogin"]=this.IsLogin

}
//tmp文件夹里 储存了每次登录储存的信息