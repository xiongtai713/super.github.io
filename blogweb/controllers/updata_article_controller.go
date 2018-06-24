
package controllers

import (
"github.com/astaxie/beego"

_ "github.com/go-sql-driver/mysql"
"sanqi/blog/blogweb/models"
)



type UpdateArticleController struct {  //继承beego.Controller 可以使用他的方法
	beego.Controller
}

//当访问add路径的时候会触发AddArticleController的get方法,响应页面是通过TplName这个属性指定返回给客户
func (this *UpdateArticleController)Get()  {  //从写get方法   浏览器访问就会触发get
	id,_:=this.GetInt("id")
	println(id)

	art:=models.QueryArticleWithId(id) //通过id 找到文章
	this.Data["Title"]=art.Title
	this.Data["Tags"]=art.Tags
	this.Data["Short"]=art.Short
	this.Data["Content"]=art.Content
	this.Data["Id"]=art.Id
	this.TplName="write_article.html"   //显示写文章的html

}
//通过this.JSON()这个方法去返回json字符串
func (this *UpdateArticleController)Post()  {
	id,_:=this.GetInt("id")  //浏览器通过表单提交来过来内容 可以用GETInt或者GetString
	println("postid",id)

	//获取浏览器传输的数据
	title:=this.GetString("title")
	tags:=this.GetString("tags")
	short:=this.GetString("short")
	content:=this.GetString("content")

	art:=models.Article{id,title,tags,short,content,"",0}
	_, err := models.UpdateArticle(art)  //修改表

	var response map[string]interface{}
	if err==nil{
		response=map[string]interface{}{"code":1,"massage":"更新成功"}

	}else {
		response=map[string]interface{}{"code":0,"massage":"更新失败"}

	}

	//返回数据给浏览器
	this.Data["json"]=response  //serverjson()
	this.ServeJSON()
}
