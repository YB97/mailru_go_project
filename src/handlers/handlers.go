package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"

	"golang.org/x/crypto/bcrypt"

	"time"

	"../database"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/julienschmidt/httprouter"
	"github.com/satori/go.uuid"
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
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Handler struct {
	DB_instance *gorm.DB
}

func Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if err := index_template.ExecuteTemplate(w, "layout", nil); err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
	}
}

func (h Handler) GetRecognitionMainPage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println(r.Cookie("logged in"))
	cookie, _ := r.Cookie("logged in")
	user := database.User{UUID: cookie.Value}
	h.DB_instance.First(&user)
	fmt.Println("!!!!!!")
	fmt.Println(user)
	fmt.Println(cookie.Value)
	if user.ID != 0 {
		if err := recognition_template.ExecuteTemplate(w, "recognition", nil); err != nil {
			log.Println(err.Error())
			http.Error(w, http.StatusText(500), 500)
		}
	} else {
		w.WriteHeader(http.StatusForbidden)
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
	jsonUserData := queryVal.Get("UserData")
	var ud userData
	err := json.Unmarshal([]byte(jsonUserData), &ud)
	fmt.Println(ud)
	username := ud.Login
	password := ud.Password
	fmt.Println(username, password)
	user_uuid := uuid.NewV4().String()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Wrong json"))
		fmt.Println("error")
	} else {
		user := database.User{LOGIN: username}
		h.DB_instance.Where("login = ?", username).First(&user)
		fmt.Println("123")
		if bcrypt.CompareHashAndPassword([]byte(user.PASSWORD), []byte(password)) == nil {
			fmt.Println("hash is ok")
			h.DB_instance.First(&user)
			user.UUID = user_uuid
			h.DB_instance.Save(&user)

			cookie := &http.Cookie{Name: "logged in", Value: user_uuid, MaxAge: -1, Expires: time.Now().Add(-100 * time.Hour)}
			http.SetCookie(w, cookie)
			w.WriteHeader(http.StatusOK)

			fmt.Println("cookie value is: ", string(cookie.Value))
		} else {
			cookie := &http.Cookie{Value: "False", MaxAge: -1, Expires: time.Now().Add(-100 * time.Hour)}
			http.SetCookie(w, cookie)
			w.WriteHeader(http.StatusForbidden)
			fmt.Println("response is: ", w.Header())
		}
	}

}

func (h Handler) Register(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	queryVal := r.URL.Query()
	jsonUserData := queryVal.Get("UserData")
	var ud userData
	err := json.Unmarshal([]byte(jsonUserData), &ud)
	fmt.Println(ud)
	username := ud.Login
	password := ud.Password

	hash, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	if err != nil {
		log.Fatal(err)
	}

	if (username == "") || (password == "") {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Wrong json"))
		fmt.Println(r.RequestURI)
		fmt.Println("Empty username or password field")
	} else {

		user := database.User{}
		h.DB_instance.Where("login = ?", username).First(&user)
		if user.ID == 0 {
			NewUser := database.User{LOGIN: username, PASSWORD: string(hash)}
			h.DB_instance.NewRecord(&NewUser)
			h.DB_instance.Create(&NewUser)
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}
