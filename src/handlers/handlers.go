package handlers

import (
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"html/template"
	"path"
	"log"
	"encoding/json"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"../configuration"
	"../database"

	"github.com/julienschmidt/httprouter"
	"path/filepath"
	"io/ioutil"
)

var (
	index_template = template.Must(template.ParseFiles(path.Join("/Users/yana/projects/mailru_go_project/src/template/", "layout.html")))
)

var (
	recognition_template = template.Must(template.ParseFiles(path.Join("/Users/yana/projects/mailru_go_project/src/template/", "recoginition.html")))
)

var (
	reg_template = template.Must(template.ParseFiles(path.Join("/Users/yana/projects/mailru_go_project/src/template/", "registration.html")))
)
type userData struct {
	Login string `json:"login"`
	Password string `json:"password"`
}


func Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if err := index_template.ExecuteTemplate(w, "layout", nil); err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
	}
}

func GetRecognitionMainPage(w http.ResponseWriter, r *http.Request, ps httprouter.Params)  {
	if err := recognition_template.ExecuteTemplate(w, "recognition", nil); err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
	}
}

func LoadFileForRecoginiton(w http.ResponseWriter, r *http.Request, ps httprouter.Params)  {
	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println(err)
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile("./mailru_go_project/images/"+handler.Filename, data, 0777)
	if err != nil {
		fmt.Println(err)
	}
	conf_path, err := filepath.Abs(filepath.Join("mailru_go_project/src/configuration/config.json"))
	if err!= nil{
		log.Fatal(err)
	}
	conf := configuration.LoadConfiguration(conf_path)
	fmt.Println(conf.Database.User)
	db, err := gorm.Open("mysql", conf.Database.User + ":" +
		conf.Database.Password + "@/" + conf.Database.Name + "")
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}
	NewPath:=database.Image{"./mailru_go_project/images/"+handler.Filename, ""}
	db.NewRecord(NewPath)
	db.Create(NewPath)


}

func RegPage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if err := reg_template.ExecuteTemplate(w, "registration", nil); err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
	}
}


func Login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	queryVal := r.URL.Query()
	jsonUserData := queryVal.Get("userData")
	var ud userData
	err := json.Unmarshal([]byte(jsonUserData), &ud)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Wrong json"))
		panic(err)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	//fmt.Println(ud)


}

func Register(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	queryVal := r.URL.Query()
	username := queryVal.Get("username")
	password := queryVal.Get("password")
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	if err!= nil{
		log.Fatal(err)
	}
	conf_path, err := filepath.Abs(filepath.Join("./src/configuration/config.json"))
	if err!= nil{
		log.Fatal(err)
	}
	conf := configuration.LoadConfiguration(conf_path)
	fmt.Println(username)

	if (username != "") || (password != ""){
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Wrong json"))
		fmt.Printf("Empty username or password field")
	} else {
		db, err := gorm.Open("mysql", conf.Database.User + ":" +
			conf.Database.Password + "@/" + conf.Database.Name + "")
		defer db.Close()

		if err != nil {
			log.Fatal(err)
		}
		NewUser := database.User{LOGIN:username, PASSWORD: string(hash)}
		db.NewRecord(NewUser)
		db.Create(&NewUser)
		w.WriteHeader(http.StatusOK)
	}

}