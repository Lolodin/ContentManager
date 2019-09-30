package main

import (
	"database/sql"
	"fmt"
	"time"
	_"github.com/go-sql-driver/mysql"
)
var(users = make(map[int]int)
db, _ = sql.Open("mysql", "root:root@tcp(localhost:1994)/mydata"))
func main()  {
	for true {
		time.Sleep(10* time.Second)
		fmt.Println("tick")
	}
}

