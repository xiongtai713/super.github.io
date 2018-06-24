package utils

import (
	"time"

	"html/template"
	"github.com/russross/blackfriday"
	"github.com/PuerkitoBio/goquery"
	"bytes"
	"github.com/sourcegraph/syntaxhighlight"
	"crypto/md5"
	"fmt"
)

func GetTimeStamp()  int64{ //获取时间
	now:=time.Now()

	return now.Unix()

}

func SwitchTimeStampToData(timeStamp int64)string  {
	t:=time.Unix(timeStamp,0)
	//时间戳转换格式"2006-01-02 15:04:05"
	return t.Format("2006-01-02 15:04:05")

}


//md5数据传入的不一样,那么md5加密后的32位数据肯等不一样
func MD5(string string)string  {
	md5string := fmt.Sprintf("%x",md5.Sum([]byte(string)))
	return md5string
}



func SwtichMaekdownToHtml(content string)template.HTML  {  //
	common := blackfriday.MarkdownCommon([]byte(content))
	//获取到html文档
	document, _ := goquery.NewDocumentFromReader(bytes.NewReader(common))
	//对docment进行查询,选择器css的语法一样
	//i是查询到的元素,selection就是查询到的元素
	document.Find("code").Each(func(i int, selection *goquery.Selection) {
		light, _ := syntaxhighlight.AsHTML([]byte(selection.Text()))
		selection.SetHtml(string(light))
		//fmt.Println("light",string(light))
	})
	htmlstring, _ := document.Html()
	return template.HTML(htmlstring)

}