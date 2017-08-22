package auth

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	//"fmt"
	"encoding/json"
	"fmt"
)

type userData struct {
	login string `json:"login"`
	pass string `json:"pass"`
}

func Registration(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	jsonUserData := ps.ByName("userData")

	ud := userData{}
	json.Unmarshal([]byte(ud), &jsonUserData)



//
	fmt.Fprintf(w, "hello, %s!\n", "Vasya")
}


func RequestAPI(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}