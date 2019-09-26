package main

import (
//"crypto/md5"
	"database/sql"
	"fmt"

	//"fmt"
	_"github.com/go-sql-driver/mysql"
)

func main() {
	db, err:= sql.Open("mysql", "root:root@tcp(localhost:1994)/mydata")
	if err != nil {
		panic(err)
	}
	fmt.Println(db)

//createShemeDB(db)
	defer db.Close()
	result, err := db.Exec("insert into mydata.user (login, password, email) values ('Peeeixel 2', 'Google', 64000)")
	if err != nil{
		panic(err)
	}
	fmt.Println(result.LastInsertId())  // id добавленного объекта
	fmt.Println(result.RowsAffected())  // количество затронутых строк
}

func createShemeDB (d *sql.DB) {
d.Exec("create table user (id int not null auto_increment, login varchar(20) not null,password varchar(20) not null, " +
	"email varchar(50) not null, primary key(id))")
}

