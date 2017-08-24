package registration

import (
	"path"
	"html/template"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"log"
//	"../project_database"
//	"encoding/json"
)

//type userData struct {
//	Login string `json:"login"`
//	Password string `json:"password"`
//}

var (
	post_template = template.Must(template.ParseFiles(path.Join("/Users/yana/projects/mailru_go_project/src/template", "registration.html")))
)

func Registrate(w http.ResponseWriter, r *http.Request, ps httprouter.Params)  {
	if err := post_template.ExecuteTemplate(w, "registration", nil); err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
	}

}