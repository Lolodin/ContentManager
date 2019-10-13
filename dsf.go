package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
	_"github.com/go-sql-driver/mysql"
)
var(
	users = make(map[int]bool)
	db, _ = sql.Open("mysql", "root:root@tcp(localhost:1994)/mydata")
)


func main()  {
	for true {
		// Получаем массив юзеров, сравниваем есть ли юзер в массиве users, если не

		time.Sleep(10 * time.Second)
		usersDB := getusersDB()
		for _, userID:= range usersDB  {
			if users[userID] == false {
				users[userID]=true
				go userServise(userID)
			}
		}

		//go userServise(user)
	}
}
func userServise(user int) {
	for true {
		time.Sleep(5* time.Second)
		//отправляем запрос на контент пользователя, получаем путь к файлу и текст
		var rowID int
		var fileID string
		var fileText string
		dataUser:= db.QueryRow("select id,filename, textfile from userfile where userid = ? order by id ASC limit 1", user)
		dataUser.Scan(&rowID,&fileID, &fileText)
		db.Exec("delete  from  userfile where id=?", rowID)
		fmt.Println(fileID, fileText)
		fmt.Println("Send Data, deleteROw")
		sendfile(fileID, fileText)
	}
}
func getusersDB() []int {
	var users[] int
	usersID,_ := db.Query("select distinct userid from userfile")
for usersID.Next() {
	var user int
	usersID.Scan(&user)
	users = append(users, user)
}
return users
}
func sendfile(fileID string, messageText string)  {
	file, err:= os.Open("SaveFile/"+ fileID)
	if err!= nil {
		fmt.Println("FileRead error")
	}


body:= &bytes.Buffer{}
writer:= multipart.NewWriter(body)
part, err:= writer.CreateFormFile("photo", filepath.Base(file.Name()))
if err!=nil {
	fmt.Println("Ошибка создания формы для отправки файла 67 строка")
}
io.Copy(part, file)
writer.Close()
	request, err := http.NewRequest("POST", "https://api.telegram.org/bot924816726:AAEbxecu_lDyh365UOyPOBCNRK3axT0j9R4/sendPhoto?chat_id=119596916", body)
	if err != nil {
		fmt.Println("error send file to telegramm")
	}
	request.Header.Add("Content-Type", writer.FormDataContentType())
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Ошибка получение ответ после запроса")
	}
	defer response.Body.Close()

	fmt.Println(response)
	fmt.Println("файл отправлен на сервер телеграмм" + messageText)
	file.Close()
defer deletefile(fileID)



}
func deletefile(fileID string)  {
	if strings.EqualFold(fileID, "") {
		return
	}
	err := os.Remove("SaveFile/"+ fileID)
	if err!= nil {
		fmt.Println("Ошибка удаления файла")
	}
}