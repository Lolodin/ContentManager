package main

import (
//"crypto/md5"
	"database/sql"
	"fmt"

	//"fmt"
	_"github.com/go-sql-driver/mysql"
)
var(
	db, _= sql.Open("mysql", "root:root@tcp(localhost:1994)/mydata")
)
func main() {


//createShemeDB(db)

//result, err := db.Exec("insert into mydata.user (login, password, email) values ('Peeeixel 2', 'Google', 64000)")


func createShemeDB () {
db.Exec("create table user (id int not null auto_increment, login varchar(20) not null,password varchar(20) not null, " +
	"email varchar(50) not null, primary key(id))")
}
func createFileDB()  {
	db.Exec("create table userFile (id int not null auto_increment, filename varchar(35) not null, userID int) ")
}

