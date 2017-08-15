package main

import (
	"net/http"
	"fmt"
	"os"
	"log"
)

func main() {
	buf, file_err := os.Open(os.Args[1])
	if file_err != nil{
		log.Fatal("input error")
	}
	resp, err := http.Post("https://vision.googleapis.com/v1/images:annotate?key=AIzaSyD8AT_TbaiOotFmgMsdHZ1lifx1bIGQo-Q", "image/jpeg", &buf)
	if err!= nil {
		fmt.Println(resp)
	} else { log.Fatal(err)
	}
}