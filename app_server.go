package main

import (
	"crypto/md5"
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
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
	Error string `json:"error"`
}
type myJson struct {
	Request string `json:"request"`
	User string `json:"user"`
}
type userResponse struct {
	Content map[int]string `json:"content"`

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
	http.Handle("/image/", http.StripPrefix("/image/", http.FileServer(http.Dir("./SaveFile"))))
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
		jsResponse:= appRoute{Page:"mainPage",  Login:cookieUser.Value}
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
	jsResponse:= appRoute{Page:"mainPage", Login: "false", Error: "user not found"}
	js, _ := json.Marshal(jsResponse)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
	return
}

	expiration := time.Now().Add(365 * 24 * time.Hour) // время работы куки
	cookie:= http.Cookie{Name: "user", Value: userID, Expires:expiration}
	http.SetCookie(w, &cookie)
	jsResponse:= appRoute{Page:"mainPage", Login: userID}
	js, _ := json.Marshal(jsResponse)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func regFunc (w http.ResponseWriter, r *http.Request) {
login, password, email := r.PostFormValue("login"),r.PostFormValue("Password"),r.PostFormValue("Email")
result, err := db.Exec("insert into mydata.user (login, password, email) values (?, ?, ?)", login, password, email)
	if err != nil {
	js:= appRoute{Error:"Login"}
	mjs,_:= json.Marshal(js)
	w.Write(mjs)
		return
	}
ID, _ := result.LastInsertId()
userID := strconv.FormatInt(ID, 16)
expiration := time.Now().Add(365 * 24 * time.Hour)

cookie:= http.Cookie{Name: "user", Value: userID, Expires:expiration}
http.SetCookie(w, &cookie)
//сделать json ответ
js := appRoute{Login:userID, Page:"mainPage"}
w.Header().Set("Content-Type", "application/json")
wjs,e:= json.Marshal(js)
if e!=nil {
	fmt.Println("error jsonendecode")
}
w.Write(wjs)


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
	userID, userError := r.Cookie("user")
	if userError != nil{
		jsResponse:= appRoute{Page:"mainPage", Login: "false"}
		js, e := json.Marshal(jsResponse)
		fmt.Println("not error 163")
		if e != nil {
			fmt.Fprintf(w," e.Error()")
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
		return
	}

	r.ParseMultipartForm(5* 1024 * 1024)
file,Handler,err:= r.FormFile("testFile")
text:=r.FormValue("testText")
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
f, _ := os.Create("SaveFile/"+ fileName[:10] +"."+ suff)
_, errDB :=db.Exec("insert into userfile(filename, userID, textfile) values (?, ?, ?)", fileName[:10] +"."+ suff, userID.Value, text)
if errDB != nil {
	fmt.Println("error sedfile to DB: " + errDB.Error())
}
f.Write(data)
		defer f.Close()
	jsResponse:= appRoute{Page:"mainPage",Login: userID.Value}
	js, e := json.Marshal(jsResponse)
	if e != nil {
		fmt.Fprintf(w," e.Error()")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

	fmt.Println(Handler.Filename)
	fmt.Println(Handler.Header)

}
func checkAuth(w http.ResponseWriter, r *http.Request) {
	cookieUser, err := r.Cookie("user")

	if err != nil{
		jsResponse := appRoute{Page:"mainPage", Login:"false"}
		js, e := json.Marshal(jsResponse)
		if e != nil {
			fmt.Fprintf(w," e.Error()")
		}
		fmt.Println("error")
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
		return

	}
	jsResponse:= appRoute{Page:"mainPage",  Login:cookieUser.Value}
	js, e := json.Marshal(jsResponse)
	fmt.Println("not error 233")
	if e != nil {
		fmt.Println(" e.Error()")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)





}
func apiController (w http.ResponseWriter, r *http.Request) {
	var js myJson

	userID, userError := r.Cookie("user")
	if userError != nil{
		jsResponse:= appRoute{Page:"mainPage", Login: "false"}
		ejs, e := json.Marshal(jsResponse)
		fmt.Println("UserNotFound")
		if e != nil {
			fmt.Fprintf(w," e.Error()")
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(ejs)
		return
	}
testbody, err:=ioutil.ReadAll(r.Body)
	if err!=nil {
		fmt.Println("Error read Body")
	}
js.User = userID.Value
err = json.Unmarshal(testbody, &js)
	if err!=nil {
		fmt.Println("error parse json")
	}
userData := getUserContent(js.User)
userR:= userResponse{Content:userData}
answer,_:= json.Marshal(userR)
w.Header().Set("Content-Type", "application/json")
w.Write(answer)

}

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