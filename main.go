package main

import (
	"net/http"
	"fmt"
)

func main() {
	resp, err := http.Post("https://vision.googleapis.com/v1/images:annotate?key=AIzaSyD8AT_TbaiOotFmgMsdHZ1lifx1bIGQo-Q", "image/jpeg", &buf)
	if err!= nil {
		fmt.Println(resp)
	} else { fmt.Println(err)
	}
}