package main

import (
	"net/http"
	"fmt"
	"os"
	"log"
	"bytes"
	"mime/multipart"
)

func main() {
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	buf, file_err := os.Open(os.Args[1])
	if file_err != nil{
		log.Fatal("input error")
	}
	defer buf.Close()
	fw, err := w.CreateFormFile("image", os.Args[1])
	if fw, err = w.CreateFormField("key"); err != nil {
		fmt.Println(fw)
	}
	req, err := http.NewRequest("POST", "https://vision.googleapis.com/v1/images:annotate?key=AIzaSyA9QNmSSQNO0JF_HSQUnQqdqRTR6YWYyBo", &b)
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", w.FormDataContentType())
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return
	}

	fmt.Println(res)
}