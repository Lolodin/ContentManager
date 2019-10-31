package main

import (
	"bytes"
	"fmt"
	//	_ "github.com/go-sql-driver/mysql"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)
var(
	users = make(map[int]bool)
	//db, _ = sql.Open("mysql", "root:root@tcp(localhost:1994)/mydata")
)


func daemonSendFile()  {
	fmt.Println("Daemon Run")
	for true {
		// Получаем массив юзеров, сравниваем есть ли юзер в массиве users, если неt

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
	chatId, timeloop := getTimeLoopAndChatID(user)
		if chatId == "null"{
			time.Sleep(time.Duration(timeloop)*time.Minute)
			continue
		}
		time.Sleep(time.Duration(timeloop)*time.Minute)
		//отправляем запрос на контент пользователя, получаем путь к файлу и текст
		var rowID int
		var fileID string
		var fileText string
		dataUser:= db.QueryRow("select id,filename, textfile from userfile where userid = ? order by id ASC limit 1", user)

		err:=dataUser.Scan(&rowID,&fileID, &fileText)
		if err != nil {
			fmt.Println("not data")
			time.Sleep(time.Duration(timeloop)*time.Minute)
			continue
		}
		fmt.Println(fileID, fileText)

		send := sendfile(fileID, fileText, chatId)
		if send {
			fmt.Println("Send Data, deleteROw")
			db.Exec("delete  from  userfile where id=?", rowID)
		}
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
func sendfile(fileID string, messageText string, chatId string) bool  {
	file, err:= os.Open(con.SaveFile+ fileID)
	if err!= nil {
		fmt.Println("FileRead error")
		return false
	}


body:= &bytes.Buffer{}
writer:= multipart.NewWriter(body)
part, err:= writer.CreateFormFile("photo", filepath.Base(file.Name()))
if err!=nil {
	fmt.Println("Ошибка создания формы для отправки файла 67 строка")
	return false
}
io.Copy(part, file)
writer.Close()

	request, err := http.NewRequest("POST", "https://api.telegram.org/bot924816726:AAEbxecu_lDyh365UOyPOBCNRK3axT0j9R4/sendPhoto?chat_id="+chatId, body)
	if err != nil {
		fmt.Println("error send file to telegramm")
		return false
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

	return true
}
func deletefile(fileID string)  {
	if strings.EqualFold(fileID, "") {
		return
	}
	err := os.Remove(con.SaveFile+ fileID)
	if err!= nil {
		fmt.Println("Ошибка удаления файла")
	}
}
func getTimeLoopAndChatID(userId int) (string, int) {
row:= db.QueryRow("SELECT hours, minutes, chatid from setting where userid = ?", userId)
var timeloop int
var hours, minutes int
var chatId string
err:=row.Scan(&hours, &minutes, &chatId)
	if err!=nil {
		fmt.Println(hours, minutes, chatId, "data string")
		fmt.Println("Error get data user.setting",userId)
		return "null", 10
	}
fmt.Println(hours, minutes, chatId)
	if hours != 0{
		timeloop = hours * 60 + minutes
	} else {
		timeloop = minutes
	}
fmt.Println(timeloop)
return chatId, timeloop
}