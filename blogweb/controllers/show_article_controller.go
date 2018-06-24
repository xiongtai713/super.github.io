package controllers

import (
	"github.com/astaxie/beego"

	"sanqi/blog/blogweb/models"
	"strconv"

	"sanqi/blog/blogweb/utils"
)

type ShowArticleController struct {
	beego.Controller
}

func (this *ShowArticleController)Get()  {
	idstr:= this.Ctx.Input.Param(":id")  //一定要加:
	//id,_:=this.GetInt(":id")  //匹配取id 要加:
	id,_:=strconv.Atoi(idstr)
	//通过id所对应的文章信息
	art:=models.QueryArticleWithId(id)  //通过id查询MySQL
	this.Data["Title"]=art.Title  //在页面上显示查询出来的标题
	this.Data["Content"]=utils.SwtichMaekdownToHtml(art.Content) //在页面上显示查询出来的内容
	this.TplName="show_article.html"

}
