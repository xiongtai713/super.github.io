package controllers

import (
	"github.com/astaxie/beego"

	_ "github.com/go-sql-driver/mysql"
	"sanqi/blog/blogweb/models"
	"sanqi/blog/blogweb/utils"
)

type Blog struct {
	title string
	tags string
	short string
	content  string
}



type AddCrticleController struct {  //继承beego.Controller 可以使用他的方法
	beego.Controller
}

//当访问add路径的时候会触发AddArticleController的get方法,响应页面是通过TplName这个属性指定返回给客户
func (this *AddCrticleController)Get()  {  //从写get方法   浏览器访问就会触发get
	
	this.TplName="write_article.html"   //显示写文章的html
	//显示原始页面

}
//通过this.JSON()这个方法去返回json字符串
func (this *AddCrticleController)Post()  {


	//获取浏览器传输的数据
	title:=this.GetString("title")  //post表单里的数据
	tags:=this.GetString("tags")
	short:=this.GetString("short")
	content:=this.GetString("content")
	art:=models.Article{0,title,tags,short,content,"空",utils.GetTimeStamp()}
	_, err := models.AddArticle(art)  //把文章添加到数据库
	var response map[string]interface{}
	if err==nil{
		response=map[string]interface{}{"code":1,"massage":"ok"}

	}else {
		response=map[string]interface{}{"code":0,"massage":"error"}

	}

	//fmt.Printf("title:%s\ntags:%s\nshort:%s\ncontent:%s\n",art.Title,art.Tags,art.Short,art.Content)
	//返回数据给浏览器
	this.Data["json"]=response  //serverjson()
	this.ServeJSON()
}