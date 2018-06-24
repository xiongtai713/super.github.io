package controllers

import (
	"github.com/astaxie/beego"

	"sanqi/blog/blogweb/models"
	"sanqi/blog/blogweb/utils"

)

type RegsiterController struct {
	beego.Controller
}

func (this *RegsiterController)Get()  {
	this.TplName="register.html"

}

func (this *RegsiterController)Post()  {
	name:=this.GetString("username")
	password:=this.GetString("password")

	//注册之前判断该用户名是否已经被注册,如果存在,不能注册
	//为啥id会大于0
	id:=models.QueryUserWithUsername(name)
	if id>0{
		this.Data["json"]=map[string]interface{}{"code":0,"message":"用户名已经存在"}
		this.ServeJSON()
		return
	}


	//存储的密码是md5后的数据,那么在登录验证的时候,也是需要将用户的密码md5之后和数据库里的密码进行判断
	user:=models.User{0,name,utils.MD5(password),0,utils.GetTimeStamp()}
	_,err:=models.InsertUser(user)
	if err!=nil{
		this.Data["json"]=map[string]interface{}{"code":0,"message":"注册失败"}
	}else {
		this.Data["json"]=map[string]interface{}{"code":1,"message":"注册成功"}

	}
	this.ServeJSON()
}