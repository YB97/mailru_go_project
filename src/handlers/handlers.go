package handlers

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"html/template"
	"path"
	"log"
	"encoding/json"
	"fmt"

)

var (
	post_template = template.Must(template.ParseFiles(path.Join("./src/template", "layout.html")))
)

type userData struct {
	Login string `json:"login"`
	Password string `json:"password"`
}

func Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if err := post_template.ExecuteTemplate(w, "layout", nil); err != nil {
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

}

func Register(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	queryVal := r.URL.Query()
	username := queryVal.Get("username")
	password := queryVal.Get("password")

	if (username != "") || (password != ""){
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Wrong json"))
		fmt.Printf("Empty username or password field")
	} else {
		w.WriteHeader(http.StatusOK)
	}

}