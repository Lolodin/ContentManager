package x

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)



func main() {
		file, err := os.Open("testpng.png")
		if err != nil {
		log.Fatal(err)
	}
		defer file.Close()
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		part, err := writer.CreateFormFile("photo", filepath.Base(file.Name()))
		if err != nil {
		log.Fatal(err)
	}
		io.Copy(part, file)
		writer.Close()
		request, err := http.NewRequest("POST", "https://api.telegram.org/bot924816726:AAEbxecu_lDyh365UOyPOBCNRK3axT0j9R4/sendPhoto?chat_id=119596916", body)
		if err != nil {
		log.Fatal(err)
	}
		request.Header.Add("Content-Type", writer.FormDataContentType())
		client := &http.Client{}
		response, err := client.Do(request)
		if err != nil {
		log.Fatal(err)
	}
		defer response.Body.Close()

		_, err = ioutil.ReadAll(response.Body)

		if err != nil {
		log.Fatal(err)
	}
		fmt.Println(response)

}