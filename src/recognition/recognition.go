package recognition

import (
	"path"
	"html/template"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"log"
)

var (
	post_template = template.Must(template.ParseFiles(path.Join("./src/template", "recoginition.html")))
)

func GetRecognitionMainPage(w http.ResponseWriter, r *http.Request, ps httprouter.Params)  {
	if err := post_template.ExecuteTemplate(w, "recognition", nil); err != nil {
	log.Println(err.Error())
	http.Error(w, http.StatusText(500), 500)
	}
}
