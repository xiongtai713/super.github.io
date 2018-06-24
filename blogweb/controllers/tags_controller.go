package controllers

import (
	"sanqi/blog/blogweb/models"

)

type TagsController struct {
	BaseController    ////继承 里面有beego和新增的两个功能
}

func (this *TagsController)Get()  {
	tags:=models.QueryArticleWithParam("Tags")  //通过Tags在MySQL中查找数据 返回数据切片
	this.Data["Tags"]=models.HandleTagsListData(tags)  //切片传入 返回相同的字段分别有多少 返回map
	//遍历 并且自动创建标签
	this.TplName="tags.html"

}


