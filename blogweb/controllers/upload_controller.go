package controllers

import (
	"fmt"


	"os"
	"time"
	"path/filepath"
	"io"
	"sanqi/blog/blogweb/utils"
	"sanqi/blog/blogweb/models"
)

type Uploadcontroller struct {
	BaseController
}

func (this *Uploadcontroller) Post() {
	fileData, fileheader, err := this.GetFile("upload")
	if err != nil {
		this.responseErr(err)
		return
	}
	filename:=fmt.Sprintf("%d-%s",utils.GetTimeStamp(),fileheader.Filename)
	fmt.Println(fileheader.Filename, fileheader.Size)
	fmt.Println(fileData)
	now := time.Now()
	fileType:="other"
	//判断后缀为图片的文件,如果是图片我们才存入到数据库中
	fileExt:=filepath.Ext(fileheader.Filename)
	if fileExt==".jpg"||fileExt==".png"||fileExt==".gif"||fileExt==".jpeg"{
		fileType="img"

	}

	//文件路径
	fileDir := fmt.Sprintf("static/upload/%s/%d/%d/%d",fileType, now.Year(), now.Month(), now.Day()) //拼接文件夹路径
	os.MkdirAll(fileDir, os.ModePerm)                                                  //创建文件夹  0777 超级权限
	filepatstr := filepath.Join(fileDir, filename)
	//拼接路径
	desfile, err := os.Create(filepatstr)  //创建文件
	if err != nil {
		this.responseErr(err)
		return
	}
	//将浏览器客户端上传的文件拷贝到本地路径的文件里面
	_,err=io.Copy(desfile,fileData) //把上传上来的文件, 写入到新建的文件里保存起来
	if err != nil {
		this.responseErr(err)
		return
	}

	if fileType=="img"{
		album:=models.Album{0,filepatstr,filename,0,utils.GetTimeStamp()}
		models.InsertAlbum(album)
	}



	this.Data["json"] = map[string]interface{}{"code": 1, "message": "上传成功"}
	this.ServeJSON()

}

func (this *Uploadcontroller)responseErr(err error)  {

	this.Data["json"]=map[string]interface{}{
		"code":0,
		"message":err,
	}
	this.ServeJSON()
}
