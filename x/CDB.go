package x

import (
	"database/sql"
	_"github.com/go-sql-driver/mysql"
)
var(
	db, errDB = sql.Open("mysql", "quirky_muser_1119:16whGoxXlAXOfYYKDSaPrtks97IAmt@tcp(manager-288.opscaptain.com:2387)/manager")
)
func main() {
createFileDB()
createSсhemeDB()
//test()
//deletDB()
}
func createSсhemeDB () {
db.Exec("create table user (id int not null auto_increment, login varchar(20) not null,password varchar(20) not null, " +
	"email varchar(50) not null, primary key(id))")
db.Exec("create table setting (userid int not null, chatid varchar(20) , hours varchar(20) , minutes varchar(20) ,  primary key(userid))")
}
func createFileDB()  {
db.Exec("create table userfile (id int not null auto_increment" +
	" primary key , filename varchar(15) not null, textfile text ,userID int) ")
}
func test()  {
//	dataUser, err:= db.Query("select * from userfile where userID = ?", user)
}
func deletDB() {
	db.Exec("drop table user, userfile, setting")
}
