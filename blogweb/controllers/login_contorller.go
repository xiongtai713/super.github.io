package controllers

import (
	"github.com/astaxie/beego"
	"sanqi/blog/blogweb/utils"
	"sanqi/blog/blogweb/models"

)

type LoginController struct {
	beego.Controller
}

func (this *LoginController)Get()  {
	this.TplName="login.html"
	
}

func (this *LoginController)Post()  {
	name:=this.GetString("username")
	password:=this.GetString("password")
	password=utils.MD5(password)
	id:=models.QueryUserWithUsernamePassword(name,password)
	if id>0{
		//设置了session后会将数据处理后设置到cookie,然后在浏览器进行网路请求的时候会自动带上cookie
		//因此我们可以通过获取这个cookie来判断用户是谁,这里我们使用的是session的方式进行设置
		this.SetSession("loginuesr",name)
		this.Data["json"]=map[string]interface{}{"code":1,"message":"登录成功"}
		this.ServeJSON()

	}else {
		this.Data["json"]=map[string]interface{}{"code":0,"message":"登录失败"}
		this.ServeJSON()
	}

}