package report

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

var contentId string = "XXXXXX"
var confulenceUrl string = "https://<BASE_URL>.net/wiki/rest/api/content/" + contentId + "/child/attachment"
var confluenceApiToken = "Basic XXXXX"
var fileP string = "/Users/vijeth.ag/go/src/confluence/report.csv"

func UploadReportToConfluence() {

	method := "PUT"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	file, errFileOpen := os.Open(fileP)
	if errFileOpen != nil {
		log.Println("error opening file", errFileOpen)
	}

	defer file.Close()
	part1, errCreateFormFile := writer.CreateFormFile("file", filepath.Base(fileP))
	if errCreateFormFile != nil {
		log.Println("error creating form file", errCreateFormFile)
	}
	_, err2 := io.Copy(part1, file)
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, confulenceUrl, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", confluenceApiToken)

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
