package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"google.golang.org/api/googleapi/transport"
	vision "google.golang.org/api/vision/v1"
	"os"
)

// https://github.com/google/google-api-go-client/blob/master/GettingStarted.md

//const developerKey = `AIzaSyA9QNmSSQNO0JF_HSQUnQqdqRTR6YWYyBo`

type Config struct {
	Database struct{
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

func ExampleGoogleCloudVisionAPI() {
	conf := LoadConfiguration("../config/config.json")

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
		Transport: &transport.APIKey{Key: conf.Key},
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

func main (){
	ExampleGoogleCloudVisionAPI()
}