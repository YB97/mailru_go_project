package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"./database"

	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/vision/v1"
	//"time"
	"github.com/julienschmidt/httprouter"
	"./handlers"
	"./configuration"
	"github.com/jinzhu/gorm"
	//"path/filepath"
)



func Router() {

	router := httprouter.New()

	router.POST("/login/", handlers.Login)
	router.GET("/recognition/", handlers.GetRecognitionMainPage)
	router.POST("/recognition/load_file/", handlers.LoadFileForRecoginiton)
//	router.GET("/", handlers.Index)
	router.GET("/registration/", handlers.RegPage)
	router.ServeFiles("/static/*filepath", http.Dir("./mailru_go_project/src/static"))
	http.ListenAndServe(":8080", router)
}



func MakeGoogleVisionRequest(config configuration.Config) {


	data, err := ioutil.ReadFile("./images")
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

func InitDatabaseConnection(conf configuration.Config)  {
	database_connection_arg := conf.Database.User + ":" + conf.Database.Password + "@/" + conf.Database.Name + ""
	db, err := gorm.Open("mysql", database_connection_arg)
//	handl := handlers.Handler{{db}}
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	if !db.HasTable(&database.User{}) {
		db.CreateTable(&database.User{})
	}
	if !db.HasTable(&database.Image{}) {
		db.CreateTable(&database.Image{})
	}
	if !db.HasTable(&database.Queue{}) {
		db.CreateTable(&database.Queue{})
	}
	db.AutoMigrate(&database.User{}, &database.Image{}, &database.Queue{})
}

func main() {
	//conf_path, err := filepath.Abs(filepath.Join("./mailru_go_project/src/configuration/config.json"))
	//if err!= nil{
	//	log.Fatal(err)
	//}
	//conf := configuration.LoadConfiguration(conf_path)
	//start := time.Now()
	//ch := make(chan int)

	Router()
	//InitDatabaseConnection(conf)
	//
	//
	//for range ch {
	//	go MakeGoogleVisionRequest(conf)
	//}

	//fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())

}
