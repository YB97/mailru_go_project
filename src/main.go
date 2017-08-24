package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"./project_database"
	"./recognition"
	"os"

	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/vision/v1"
	"time"
	"github.com/julienschmidt/httprouter"
	"./auth"
	"./registration"
)

// https://github.com/google/google-api-go-client/blob/master/GettingStarted.md

//const developerK
// ey = `AIzaSyA9QNmSSQNO0JF_HSQUnQqdqRTR6YWYyBo`


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

	router.POST("/login/", auth.Login)
	router.GET("/registration/", registration.Registrate)
	router.GET("/recognition/", recognition.GetRecognitionMainPage)
	router.GET("/", auth.Index)
	router.ServeFiles("/static/*filepath", http.Dir("/Users/yana/projects/mailru_go_project/src/static"))
	http.ListenAndServe(":8080", router)
}




func MakeGoogleVisionRequest(config Config) {


	data, err := ioutil.ReadFile("/Users/yana/projects/mailru_go_project/images")
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

func InitDatabaseConnection(conf Config, user project_database.User)  {
	project_database.StartConnection(conf.Database.Name, conf.Database.User, conf.Database.Password)
	//fmt.Println()
	//project_database.CheckExistAndCreate(conf.Database.Name, conf.Database.User, conf.Database.Password, &user)
}

func main() {

	conf := LoadConfiguration("/Users/yana/projects/mailru_go_project/config/config.json")
	start := time.Now()
	ch := make(chan int)

	Router()
	u := project_database.User{ "login", "passw"}
	InitDatabaseConnection(conf, u)


	for range ch {
		go MakeGoogleVisionRequest(conf)
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())

}