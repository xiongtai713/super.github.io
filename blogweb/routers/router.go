package routers

import (
	"sanqi/blog/blogweb/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//主页 显示所有文章
    beego.Router("/", &controllers.HomeController{})  //从浏览器访问会触发get
    //添加文章
    beego.Router("/add",&controllers.AddCrticleController{})
    //显示文章
    beego.Router("/article/:id",&controllers.ShowArticleController{})
    //更新文章
    beego.Router("/article/update",&controllers.UpdateArticleController{})
    //删除文章
    beego.Router("/article/delete",&controllers.DeleteArticle{})
    //标签显示文章
    beego.Router("/tags",&controllers.TagsController{})
    //注册
    beego.Router("/register",&controllers.RegsiterController{})
    //登录
	beego.Router("/login",&controllers.LoginController{})
	//退出
	beego.Router("/exit",&controllers.ExitController{})
	//相册
	beego.Router("/album",&controllers.AlbumController{})
	//上传图片
	beego.Router("/upload",&controllers.Uploadcontroller{})


}
