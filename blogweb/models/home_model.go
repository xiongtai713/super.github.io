package models

import (
	"fmt"
	"html/template"
	"bytes"
	"strconv"
	"sanqi/blog/blogweb/utils"
	"github.com/astaxie/beego"

	"strings"
)

type TagLink struct {
	TagName string
	TagUrl  string
}

type HomeBlockParam struct {
	Id         int
	Title      string
	Tags       []TagLink
	Short      string
	Content    string
	Author     string
	CreateTime string
	Link       string //查看文章的地址
	UpdateLink string //修改文章的地址
	DeleteLink string //删除文章的地址
	IsLogin    bool   //纪录是否登录就
}

type HomeFooterPageCode struct {
	HasPre   bool   //上一页是否能点
	HasNext  bool   //下一页是否能点
	ShowPage string //当前显示页
	PreLink  string //上一页的地址
	NextLink string //下一页的地址
}

func MakeHomeBlock(articles []Article,islogin bool) template.HTML { //解析MySQL中的数据 解析成HTML格式
	htmlHome := ""

	for _, art := range articles {
		//buffer, _ := ioutil.ReadFile("views/block/block.html")
		//htmlHome+=string(buffer)

		htmlParam := HomeBlockParam{}
		htmlParam.Id = art.Id
		htmlParam.Tags = creatTaglinks(art.Tags)
		htmlParam.Title = art.Title
		htmlParam.CreateTime = utils.SwitchTimeStampToData(art.CreateTime)
		htmlParam.Author = art.Author
		htmlParam.Content = art.Content
		htmlParam.Short = art.Short
		htmlParam.Link = "/article/" + strconv.Itoa(art.Id)                 //添加
		htmlParam.UpdateLink = "/article/update?id=" + strconv.Itoa(art.Id) //?后面添加
		htmlParam.DeleteLink = "/article/delete?id=" + strconv.Itoa(art.Id) //?后面添加 id
		htmlParam.IsLogin=islogin
		//处理变量

		files, _ := template.ParseFiles("views/block/block.html") //解析网页
		buffer := bytes.Buffer{}                                  //创建一个实现了writer的结构体
		//替换html文件里面的变量替换为传递进去的书库
		files.Execute(&buffer, htmlParam) //替换解析完的网页内容
		htmlHome += buffer.String()

	}
	return template.HTML(htmlHome)

}

//page 是当前的页数
func ConfigHomeFooterPageCode(page int) HomeFooterPageCode {
	pageCode := HomeFooterPageCode{}                           //分页结构体
	num := GetArtcileRowNum()                                  //文章总数
	pagRow, _ := beego.AppConfig.Int("artNum")                 //每页显示多少篇文章
	allpageNum := ((num - 1) / pagRow) + 1                     //分页 (40篇-1)/10=3.9=取整3 +1=4页
	pageCode.ShowPage = fmt.Sprintf("%d/%d", page, allpageNum) //page当前页数/allpageNum总页数

	//如果当前页数小于1 那么上一页不能点击
	if page <= 1 {
		pageCode.HasPre = false //上一页不能点击
	} else {
		pageCode.HasPre = true //上一页可以点击
	}
	//如果页数等于最后一页,那么下一页不能点击
	if page == allpageNum {
		pageCode.HasNext = false

	} else {
		pageCode.HasNext = true
	}
	pageCode.PreLink = "/?page=" + strconv.Itoa(page-1)  //上一页
	pageCode.NextLink = "/?page=" + strconv.Itoa(page+1) //下一页

	return pageCode

}

//把标签分成标签和地址
func creatTaglinks(tags string) []TagLink {
	var taglink []TagLink
	tagPamar := strings.Split(tags, "&")
	for _, tag := range tagPamar {
		taglink = append(taglink, TagLink{tag, "/?tag=" + tag})
	}
	return taglink
}
