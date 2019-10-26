package x

import (
	"database/sql"
	_"github.com/go-sql-driver/mysql"
)
var(
	db, _= sql.Open("mysql", "root:root@tcp(localhost:1994)/mydata")
)
func main() {
//createFileDB()
//createSсhemeDB()
test()
//deletDB()
}
func createSсhemeDB () {
db.Exec("create table user (id int not null auto_increment, login varchar(20) not null,password varchar(20) not null, " +
	"email varchar(50) not null, primary key(id))")
}
func createFileDB()  {
db.Exec("create table userFile (id int not null auto_increment" +
	" primary key , filename varchar(15) not null, textfile text ,userID int) ")
}
func test()  {
	dataUser, err:= db.Query("select * from userfile where userID = ?", user)
}
func deletDB() {
	db.Exec("drop table user, userfile")
}
