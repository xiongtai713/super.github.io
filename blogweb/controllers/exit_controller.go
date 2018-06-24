package controllers

type ExitController struct {
	BaseController //包括beego继承
}

func (this *ExitController)Get()  {
	//清除用户登录状态的数据
	this.DelSession("loginuesr")  //删除Session
	//重定向,不用从新设定页面参数
	this.Redirect("/",302)  //显示主页

}
