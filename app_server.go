package main

import (
	"crypto/md5"
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	_"github.com/go-sql-driver/mysql"
	"time"
)
var (
	user map[string]string
	db, _ = sql.Open("mysql", "root:root@tcp(localhost:1994)/mydata")
)

type appRoute struct {
	Page string `json:"page"`
	Login string `json:"login"`
}


func main()  {
	//route

	http.HandleFunc("/regfunc", regFunc)
	http.HandleFunc("/authFunc", authFunc)
	http.HandleFunc("/test", indexTest)
	http.HandleFunc("/send/", sendMessage)
	http.HandleFunc("/update", getUpdate)
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/formHandler", formHandler)
	//checkAuth
	http.HandleFunc("/checkAuth", checkAuth)
	http.HandleFunc("/apiController", apiController)
	//static
	http.Handle("/static/js/", http.StripPrefix("/static/js/", http.FileServer(http.Dir("./my-app/build/static/js"))))
	http.Handle("/static/css/", http.StripPrefix("/static/css/", http.FileServer(http.Dir("./my-app/build/static/css"))))
	http.Handle("/icon/", http.StripPrefix("/icon/", http.FileServer(http.Dir("./my-app/build"))))
	http.Handle("/static/media/", http.StripPrefix("/static/media/", http.FileServer(http.Dir("./my-app/build/static/media"))))
	//static/media/
	http.ListenAndServe(":8080", nil)
}
func indexHandler (w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("my-app/build/index.html")
	 err := t.Execute(w, "index")
	 if err != nil {
	 	fmt.Println(err.Error())
	 }

}
func sendMessage (w http.ResponseWriter, r *http.Request) {


}
func authFunc (w http.ResponseWriter, r *http.Request) {
	cookieUser, err := r.Cookie("user")

	if err == nil{
		jsResponse:= appRoute{"mainPage",  cookieUser.Value}
		js, e := json.Marshal(jsResponse)
		if e != nil {
			fmt.Fprintf(w, e.Error())
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)

		return
	}
login, password:= r.PostFormValue("login"),r.PostFormValue("Password")
ID := db.QueryRow("SELECT id FROM user WHERE login = ? && Password = ?", login, password)
fmt.Println(ID)
var userID string
err = ID.Scan(&userID)
if err!=nil {
	jsResponse:= appRoute{"mainPage",  "false"}
	js, _ := json.Marshal(jsResponse)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
	return
}
	expiration := time.Now().Add(365 * 24 * time.Hour)
	cookie:= http.Cookie{Name: "user", Value: userID, Expires:expiration}
	http.SetCookie(w, &cookie)
	jsResponse:= appRoute{"mainPage",  userID}
	js, _ := json.Marshal(jsResponse)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
func regFunc (w http.ResponseWriter, r *http.Request) {
login, password, email := r.PostFormValue("login"),r.PostFormValue("Password"),r.PostFormValue("Email")
result, err := db.Exec("insert into mydata.user (login, password, email) values (?, ?, ?)", login, password, email)
	if err != nil {
	fmt.Println("Error query")
	}
ID, _ := result.LastInsertId()
userID := strconv.FormatInt(ID, 16)
expiration := time.Now().Add(365 * 24 * time.Hour)

cookie:= http.Cookie{Name: "user", Value: userID, Expires:expiration}
http.SetCookie(w, &cookie)
fmt.Fprintf(w, "registration  OK")

}
func indexTest (w http.ResponseWriter, r *http.Request) {
	//https://api.telegram.org/bot924816726:AAEbxecu_lDyh365UOyPOBCNRK3axT0j9R4/getMe
	//var chat_id, text string;
	a, e := http.Get("https://api.telegram.org/bot924816726:AAEbxecu_lDyh365UOyPOBCNRK3axT0j9R4/getMe")
	if e!= nil {
		fmt.Fprintf(w, "error")
	}
	err := a.Write(w)
	if err!=nil {
		fmt.Fprintf(w, "error")
}
//fmt.Fprintf(w , "%s", a)
}
func getUpdate (w http.ResponseWriter, r *http.Request) {
	//https://api.telegram.org/bot<YourBOTToken>/getUpdates
//var chat_id, text string;
a, e := http.Get("https://api.telegram.org/bot924816726:AAEbxecu_lDyh365UOyPOBCNRK3axT0j9R4/getUpdates")
if e!= nil {
fmt.Fprintf(w, "error")
}
err := a.Write(w)
if err!=nil {
fmt.Fprintf(w, "error")
}
//fmt.Fprintf(w , "%s", a)
}
func formHandler(w http.ResponseWriter, r *http.Request)  {
	r.ParseMultipartForm(5* 1024 * 1024)
file,Handler,err:= r.FormFile("testFile")
	if err!=nil {
	fmt.Errorf("%s", err)
	fmt.Println(err)
		return
	}

defer file.Close()
	data:= make([]byte, Handler.Size)
	for {
		_, err := file.Read(data)
		if err == io.EOF { // если конец файла
			break // выходим из цикла
		}
	}

hashSum := md5.Sum(data)
fileName := ""
for _, val := range hashSum {
	fileName += strconv.FormatInt(int64(val), 16)
}
suff := strings.Split(Handler.Filename, ".")[1]
f, _ := os.Create("SaveFile/"+ fileName +"."+ suff)


f.Write(data)
		defer f.Close()
	fmt.Println(Handler.Filename)
	fmt.Println(Handler.Header)

}
func checkAuth(w http.ResponseWriter, r *http.Request) {
	cookieUser, err := r.Cookie("user")

	if err != nil{
		jsResponse := appRoute{"mainPage", "false"}
		js, e := json.Marshal(jsResponse)
		if e != nil {
			fmt.Fprintf(w," e.Error()")
		}
		fmt.Println("error")
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
		return

	}
	jsResponse:= appRoute{"mainPage",  cookieUser.Value}
	js, e := json.Marshal(jsResponse)
	fmt.Println("not error")
	if e != nil {
		fmt.Fprintf(w," e.Error()")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)





}
func apiController (w http.ResponseWriter, r *http.Request) {


}