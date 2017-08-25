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
)

var (
	index_template = template.Must(template.ParseFiles(path.Join("./src/template", "layout.html")))
)

var (
	recognition_template = template.Must(template.ParseFiles(path.Join("./src/template", "recoginition.html")))
)

var (
	reg_template = template.Must(template.ParseFiles(path.Join("./src/template", "registration.html")))
)
type userData struct {
	Login string `json:"login"`
	Password string `json:"password"`
}

type Handler struct{
	db_instance *gorm.DB
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

func RegPage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if err := reg_template.ExecuteTemplate(w, "registration", nil); err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
	}
}


func (h Handler) Login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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
		h.db_instance.First(&user)
		user.UUID = string(user_uuid)
		h.db_instance.Save(&user)
		cookie := &http.Cookie{Name: "test", Value: string(user_uuid), MaxAge: -1, Expires: time.Now().Add(-100 * time.Hour) }
		http.SetCookie(w, cookie)
		w.WriteHeader(http.StatusOK)
	}

}

func (h Handler) Register(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if (username != "") || (password != ""){
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Wrong json"))
		fmt.Printf("Empty username or password field")
	} else {

	//	user := database.User{LOGIN:username, PASSWORD: string(hash)}

		NewUser := database.User{LOGIN:username, PASSWORD: string(hash)}
		h.db_instance.NewRecord(NewUser)
		h.db_instance.Create(&NewUser)
		w.WriteHeader(http.StatusOK)
	}

}