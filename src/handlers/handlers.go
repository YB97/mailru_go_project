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
	"../database"
		"github.com/julienschmidt/httprouter"
	"github.com/satori/go.uuid"
	"time"
	"io/ioutil"
)

var (
	login_template = template.Must(template.ParseFiles(path.Join("./mailru_go_project/src/template", "auth.html")))
)

var (
	recognition_template = template.Must(template.ParseFiles(path.Join("./mailru_go_project/src/template", "recoginition.html")))
)

var (
	response_template = template.Must(template.ParseFiles(path.Join("./mailru_go_project/src/template", "response.html")))
)
type userData struct {
	Login string `json:"login"`
	Password string `json:"password"`
}

type Handler struct{
	DB_instance *gorm.DB
}

func GetLoginPage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if err := login_template.ExecuteTemplate(w, "auth", nil); err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
	}
}

func GetRecognitionPage(w http.ResponseWriter, r *http.Request, ps httprouter.Params)  {
	if err := recognition_template.ExecuteTemplate(w, "recognition", nil); err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
	}
}

func GetResponse(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if err := response_template.ExecuteTemplate(w, "response", nil); err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
	}
}

func (h Handler) LoginUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	queryVal := r.URL.Query()
	jsonUserData := queryVal.Get("userData")
	var ud userData
	err := json.Unmarshal([]byte(jsonUserData), &ud)
	username := ud.Login
	password := ud.Password
	user_uuid := uuid.NewV4()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Wrong json"))
		panic(err)
	} else {
		user := database.User{LOGIN: username, PASSWORD: password}
		h.DB_instance.First(&user)
		user.UUID = user_uuid.String()
		h.DB_instance.Save(&user)
		cookie := &http.Cookie{Name: "test", Value: user_uuid.String(), MaxAge: -1, Expires: time.Now().Add(-100 * time.Hour) }
		http.SetCookie(w, cookie)
		w.WriteHeader(http.StatusOK)
	}

}

func (h Handler) CreateNewUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	queryVal := r.URL.Query()
	jsonUserData := queryVal.Get("userData")
	var ud userData
	err := json.Unmarshal([]byte(jsonUserData), &ud)
	username := ud.Login
	password := ud.Password

	hash, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	if err!= nil{
		log.Fatal(err)
	}

	if (username == "") || (password == ""){
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Wrong json"))
		fmt.Printf("Empty username or password field")
	} else {


		user := database.User{}
		h.DB_instance.Where("login = ? AND password = ?", username, hash)
		if user.ID == 0{
			NewUser := database.User{LOGIN:username, PASSWORD: string(hash)}
			h.DB_instance.NewRecord(NewUser)
			h.DB_instance.Create(&NewUser)
			w.WriteHeader(http.StatusOK)
		} else{
			w.WriteHeader(http.StatusBadRequest)
		}


	}

}


func (h Handler) GetImage(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
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
	//image := database.Image{}

	NewImage:=database.Image{ FILENAME:handler.Filename}
	h.DB_instance.NewRecord(NewImage)
	h.DB_instance.Create(&NewImage)

	//cookie, _ := r.Cookie("logged in")
	//
	//
	//h.DB_instance.Where("uuid = ?", cookie.Value)
	user:=database.User{}
	NewQueue:=database.Queue{User:user, Image: NewImage}
	h.DB_instance.NewRecord(NewQueue)
	h.DB_instance.Create(&NewQueue)
}