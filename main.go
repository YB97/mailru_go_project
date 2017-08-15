package main

import (
	"net/http"
	"fmt"
)

func main() {
	resp, err := http.Post("https://vision.googleapis.com/v1/images:annotate?key=api_key_would_be_here", "image/jpeg", &buf)
	if err!= nil {
		fmt.Println(resp)
	} else { fmt.Println(err)
	}
}