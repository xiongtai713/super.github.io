package utils

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"fmt"
)

var db *sql.DB

func InitMysql()  {
	fmt.Println("InitMysql")
	if db == nil {
		db, _ = sql.Open("mysql", "root:liu713@tcp(localhost:3306)/blogweb?charset=utf8")
		CreateTableWithArticle()
		CreateTableWithUser()
		CreateTableWithAlbum()
	}
}

func CreateTableWithArticle() {  //保存博客表
	sql := `create table if not exists article(
		id int(4) primary key auto_increment not null,
		title varchar(30),
		author varchar(20),
		tags varchar(30),
		short varchar(255),
		content longtext,
		createtime int(10)
		);`
	ModifyDB(sql)
}
func CreateTableWithUser() {  //登录表
	sql := `create table if not exists users(
		id int(4) primary key auto_increment not null,
		username varchar(64),
		password varchar(64),
		status int(4),
		createtime int(10)
		);`
	ModifyDB(sql)
}

func CreateTableWithAlbum() {
	sql := `create table if not exists album(
		id int(4) primary key auto_increment not null,
		filepath varchar(255),
		filename varchar(64),
		status int(4),
		createtime int(10)
		);`
	ModifyDB(sql)
}

func ModifyDB(sqlString string, args ...interface{}) (int64, error) {
	result, err := db.Exec(sqlString, args...)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	//fmt.Println(count)

	return count, nil
}

func QueryDB(sqlstring string)(*sql.Rows,error) {
	return db.Query(sqlstring)

}

func QueryRowDB(sqlstring string) *sql.Row {
	return db.QueryRow(sqlstring)

}