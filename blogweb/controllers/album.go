package controllers

import (
	"sanqi/blog/blogweb/models"
	"fmt"
)

type AlbumController struct {
	BaseController //高级继承 基础beego和sessno
}

func (this *AlbumController)Get()  {

	albums, _ := models.FindAllAlbums()
	fmt.Println(albums)
	this.Data["Album"]=albums

	this.TplName="album.html"

}
