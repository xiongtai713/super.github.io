package controllers

import (
	"sanqi/blog/blogweb/models"
	"fmt"
)

type HomeController struct {
	//继承 里面有beego和新增的两个功能
	BaseController
}



//可以通过翻页来获取该网页,也可以通过tag标签来获取
//传page参数代表翻页,传tag参数代表标签
//首先判断page有值,那么就是翻页,否则判断tag就是标签,否则就是默认的第一页

func (this *HomeController) Get() {
	fmt.Println(this.IsLogin,this.LoginUser)


	tag := this.GetString("tag")  //取得主页的tag 标签
	page, _ := this.GetInt("page")  //取得主页的page 页码
	var articles []models.Article
	if len(tag) > 0 {  //如果标签大于0 就是按照标签在数据库中查询
		articles,_=models.QueryArticlesWithTag(tag) //按照标签去数据库中查询数据保存到articles里
		this.Data["HasFooter"]=false
		//未完成
	} else {
		if page <= 0 {  //否则就显示主页
			page = 1
		}
		articles, _ = models.FinArticleWithPage(page)  //按照页码去数据库中查询数据保存到articles里
		this.Data["PageCode"] = models.ConfigHomeFooterPageCode(page) // 把总页数和上一页地址和下一页地址传到主页
		this.Data["HasFooter"]=true


	}

	this.Data["Content"] = models.MakeHomeBlock(articles,this.IsLogin) //把数据库中取到的数据传到主页上

	this.TplName = "home.html"

}
