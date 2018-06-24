package controllers

import (
	"github.com/astaxie/beego"

	"sanqi/blog/blogweb/models"
)

type DeleteArticle struct {
	beego.Controller
}

//点击删除后 重定向首页
func (this *DeleteArticle)Get()  {
	artid,_:=this.GetInt("id")

	models.DeleteArticle(artid)  //在数据库删除文章

	this.Redirect("/",302)

}
