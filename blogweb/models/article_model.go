package models

import (
	"sanqi/blog/blogweb/utils"
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
)

type Article struct {
	Id         int
	Title      string
	Tags       string
	Short      string
	Content    string
	Author     string
	CreateTime int64
}

//存储表的行数,只有自己可以更改,当文章被新增或者删除时需要更新这个值
var articleRowNum = 0

/*数据处理*/
//添加文章
func AddArticle(art Article) (int64, error) {
	i,err:=InsertArticle(art)
	SetArticleRowNum()  //添加文章的后 先判断文章总数
	return i,err

}
func DeleteArticle(atrid int) (int64,error) {
	i,err:=deleteAticleWithArtId(atrid)
	SetArticleRowNum()  //删除文章后 在统计文章总数
	return i,err
}




func FinArticleWithPage(page int) ([]Article, error) {
	num,_:=beego.AppConfig.Int("artNum")  //在配置文件里
	page--
	return QueryArticleWithPage(page,num)

}



//首次获取行数的时候才去统计表里面的行数
func GetArtcileRowNum() int{
	if articleRowNum==0{
		articleRowNum=QueryArticleRowNum()
	}
	return articleRowNum

}

func SetArticleRowNum()  {
	articleRowNum=QueryArticleRowNum()

}

/*
数据库操作
*/

func InsertArticle(art Article) (int64, error) {
	return utils.ModifyDB("insert into article(title,tags,short,content,author,createtime)values(?,?,?,?,?,?)",
		art.Title, art.Tags, art.Short, art.Content, art.Author, art.CreateTime)
}
//分页查询
//limit分页查询语句 limit m , n
//注意limit前面没有where
//m代表多少位开始获取,注意.与id值无关
//n代表获取多少条数据
func QueryArticleWithPage(page int,num int) ([]Article, error) {
	sqlString:=fmt.Sprintf("select id,title,tags,short,content,author,createtime from article limit %d,%d",page*num,num) //
	//utils.QueryDB可以取到db																					取0-10 条如果page是1(因为有page--) 就是0-10
									//																			如果page是2 就是从第条10开始,取10条纪录
	rows, err := utils.QueryDB(sqlString)
	if err != nil {  //
		return nil, err
	}
	var artlist []Article
	id := 0
	title := ""
	tags := ""
	short := ""
	content := ""
	author := ""
	var createtime int64

	for rows.Next() {
		rows.Scan(&id, &title, &tags, &short, &content, &author, &createtime)
		art := Article{id, title, tags, short, content, author, createtime}
		artlist = append(artlist, art)
	}
	return artlist, nil
}

func QueryArticleWithId(id int)Article  {  //整行查询

	row:=utils.QueryRowDB("select id,title,tags,short,content,author,createtime from article where id="+strconv.Itoa(id))

	title := ""
	tags := ""
	short := ""
	content := ""
	author := ""
	var createtime int64
	row.Scan(&id, &title, &tags, &short, &content, &author, &createtime)
	art := Article{id, title, tags, short, content, author, createtime}

	return art
}
func UpdateArticle(article Article) (int64,error) {  //更新数据
	return utils.ModifyDB("update article set title=?,tags=?,short=?,content=? where id=?",article.Title,article.Tags,article.Short,article.Content,article.Id)

}

func QueryArticleRowNum() int  {  //查询有多少篇文章
	row := utils.QueryRowDB("select count(id) from article")
	num:=0
	row.Scan(&num)
	return num

}

func QueryArticleWithParam(param string)[]string  {
	rows,_:=utils.QueryDB("select "+param+" from article")  //别忘加空格
	var paramlist []string
	for rows.Next() {  //循环取值
		arg:=""
		rows.Scan(&arg)  //把值写入arg
		paramlist=append(paramlist,arg)  //把arg存到切片
	}
	return paramlist

}
//通过标签查询
//有四种情况
/*
1.左右两边有&符号和其他符号
2.左有&和其他符合,同时右边没有任何符号
3.右有&和其他符合,同时左边没有任何符号
4.左右两边都没有符号
通过%去匹配任意多个字符,至少是一个


 */
func QueryArticlesWithTag(tag string)([]Article ,error)  {
	sqlString:="where tags like '%&"+tag+"&%'"   //模糊查询
	sqlString+="or tags like '%&"+tag+"'"   //模糊查询
	sqlString+="or tags like '"+tag+"&%'"   //模糊查询
	sqlString+="or tags like '"+tag+"'"   //模糊查询
	fmt.Println("#################",sqlString)
	return QueryArticlesWithCon(sqlString)
	//where tags like '%&AJXA&%'or tags like '%&AJXA'or tags like 'AJXA&%'or tags like 'AJXA'

}

func QueryArticlesWithCon(sqlString string) ([]Article, error) {
	sqlString = "select id, title, tags, short, content, author, createtime from article "+sqlString
	rows, err := utils.QueryDB(sqlString)
	if err != nil {
		return nil, err
	}
	var artList []Article
	for rows.Next() {
		id := 0
		title := ""
		tags := ""
		short := ""
		content := ""
		author := ""
		var createtime int64
		createtime = 0
		rows.Scan(&id, &title, &tags, &short, &content, &author, &createtime)
		art := Article{id, title, tags, short, content, author, createtime}

		artList = append(artList, art)
	}
	return artList,nil
}

func deleteAticleWithArtId(artID int)(int64, error){
	return utils.ModifyDB("delete from article where id=?",artID)

}