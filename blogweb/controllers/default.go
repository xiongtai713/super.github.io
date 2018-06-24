package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

type User struct {
	UserName string
	Age int
	Sex bool
}

func (c *MainController) Get() {
	c.Data["Website"] = "kongyixueyuan.com"  //官方模板
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["myname"]="liujiao"
	user:=User{"1803",1,true}
	c.TplName = "index.html"
	user2:=User{"1802",2,false}
	user3:=User{"1801",3,true}
	c.Data["Users"]=[]User{user,user2,user3}
	c.Data["User"]=user
}


type HomeMainController struct {
	beego.Controller
}

func (c *HomeMainController) Get() {
	c.Data["Website"] = "kongyixueyuan.com"  //官方模板
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["myname"]="liujiao"
	c.TplName = "04-导航.html"
}

