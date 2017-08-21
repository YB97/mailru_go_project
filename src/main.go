package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"./project_database"
	"os"
	//"../config"


	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/vision/v1"
	"time"
	"github.com/julienschmidt/httprouter"
)

// https://github.com/google/google-api-go-client/blob/master/GettingStarted.md

//const developerKey = `AIzaSyA9QNmSSQNO0JF_HSQUnQqdqRTR6YWYyBo`


type Config struct {
	Database struct{
		Name 	   string  `json:"dbname"`
		User       string  `json:"user"`
		Password   string  `json:"password"`
	} `json:"database"`

	Key string `json:"key"`
}

func LoadConfiguration(file string) Config {
	var config Config
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}

	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}

func Router() {
	router := httprouter.New()
	router.GET("/reg/", auth.Registration)

	http.ListenAndServe(":8080", router)
	//log.Fatal(http.ListenAndServe(":8080", router))
}



func MakeGoogleVisionRequest(config Config) {

//	key := &Config{}
	//conf :=
	data, err := ioutil.ReadFile("images/cat.jpg")

	enc := base64.StdEncoding.EncodeToString(data)
	img := &vision.Image{Content: enc}

	feature := &vision.Feature{
		Type:       "LABEL_DETECTION",
		MaxResults: 10,
	}

	req := &vision.AnnotateImageRequest{
		Image:    img,
		Features: []*vision.Feature{feature},
	}

	batch := &vision.BatchAnnotateImagesRequest{
		Requests: []*vision.AnnotateImageRequest{req},
	}

	client := &http.Client{
		Transport: &transport.APIKey{Key: config.Key },
	}
	svc, err := vision.New(client)
	if err != nil {
		log.Fatal(err)
	}
	res, err := svc.Images.Annotate(batch).Do()
	if err != nil {
		log.Fatal(err)
	}

	body, err := json.Marshal(res.Responses[0].LabelAnnotations)
	fmt.Println(string(body))
}

func InitDatabaseConnection(conf Config)  {
	project_database.StartConnection(conf.Database.Name, conf.Database.User, conf.Database.Password)

}

func Router()  {
	fmt.Printf("%s", "here")
	router := httprouter.New()
	router.GET("/reg/", auth.Registration)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	conf:=LoadConfiguration("/Users/yana/projects/mailru_go_project/config/config.json")
	start := time.Now()
  
	ch := make(chan int)
	InitDatabaseConnection(conf)
	Router()
	for range ch {
		go MakeGoogleVisionRequest(conf)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())

}
