package models

import "sanqi/blog/blogweb/utils"

type User struct {
	Id       int
	Username string
	Password string
	Status   int  //0正常状态,  //1删除
	Createtime int64
}

func InsertUser(user User) (int64,error) {
	return utils.ModifyDB("insert into users (username,password,status,createtime) values (?,?,?,?)",
		user.Username,user.Password,user.Status,user.Createtime)


}


//注册验证
func QueryUserWithUsername(username string)int  {
	//通过用户名检查id 如果查到了id 给id赋值,说明数据库里有重复的用户, 如果没查到 id返回0 说明可以创建用户
	row:=utils.QueryRowDB("select id from users where username='"+username+"'")
	id:=0
	row.Scan(&id)
	return id

}


func QueryUserWithUsernamePassword(username string,password string)int  {
	//通过用户名检查id 如果查到了id 给id赋值,说明数据库里有重复的用户, 如果没查到 id返回0 说明可以创建用户
	row:=utils.QueryRowDB("select id from users where username='"+username+"' and password='"+password+"'")
	id:=0
	row.Scan(&id)
	return id

}
