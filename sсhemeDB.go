package main

import (
	"database/sql"
	_"github.com/go-sql-driver/mysql"
)
var(
	db, _= sql.Open("mysql", "root:root@tcp(localhost:1994)/mydata")
)
func main() {
//createFileDB()
//createShemeDB()
//test()
}
func createShemeDB () {
db.Exec("create table user (id int not null auto_increment, login varchar(20) not null,password varchar(20) not null, " +
	"email varchar(50) not null, primary key(id))")
}
func createFileDB()  {
db.Exec("create table userFile (id int not null auto_increment" +
	" primary key , filename varchar(35) not null, userID int) ")
}
func test()  {
	db.Exec("insert into user(login, password, email) values (?,?,?)", "test", "test", "test@test")
	db.Exec("insert into userFile(filename, userID) values (?, ?)", "testfile", 1)
}
func deletDB() {
	db.Exec("drop table user, userfile")
}
