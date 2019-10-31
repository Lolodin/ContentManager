package main

import (
	"fmt"
	"os"
)

func getUserContent(userID string) map[int]string {
	//var arrUserConten  map[int]string
	arrUserContent:= make( map[int]string)
	userContent, err:= db.Query("select id, filename from userfile where userID = ?", userID)
	if err!=nil {
		fmt.Println("Ошибка запроса файлов пользователя")
	}

	for userContent.Next() {
		var filename string
		var contentID int
		userContent.Scan(&contentID ,&filename)
		arrUserContent[contentID] = filename
	}
	fmt.Println(arrUserContent)
	return arrUserContent
}
func deleteRow(idRow string, filename string) {
	err := os.Remove(con.SaveFile+ filename)
	if err!= nil {
		fmt.Println("Ошибка удаления файла")
	}
	_, err=db.Exec("delete  from userfile where id = ?", idRow)
	if err!=nil {
		fmt.Println("Ошибка удаления записи")
		return
	}
	fmt.Println("Запись удалена")
}
func updataUserContent(updata [][]string) {
	for _,arr := range updata {
		idContent := arr[0]
		contentFileName:= arr[1]
		_,err:= db.Exec("UPDATE userfile set filename = ? where  id =?", contentFileName, idContent)
		if err!=nil {
			fmt.Println("error updata DB")
		}
	}
}
func updateUserSetting(chatId, hours, minutes, userid string) {
	_,err:=db.Exec("UPDATE setting set chatid = ?, hours = ?, minutes = ? where userid = ?",chatId, hours, minutes, userid)
	if err!= nil {
		fmt.Println("Error write user setting to BD")
	}
}
func getSetting(userID string) map[string]string {
	var chatId, hours, minutes string
	var setting map[string]string
	setting = make(map[string]string)
	row :=db.QueryRow("SELECT chatid, hours, minutes FROM setting WHERE userid=?", userID)
	row.Scan(&chatId, &hours, &minutes)
	setting["chatid"] =chatId
	setting["hours"] = hours
	setting["minutes"] = minutes
	fmt.Println(setting, "getSetting")
	return setting


}