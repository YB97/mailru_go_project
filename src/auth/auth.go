package auth

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"fmt"
)

func Registration(w http.ResponseWriter, r *http.Request, ps httprouter.Params)  {



	fmt.Fprintf(w, "hello, %s!\n", "Vasya")
}
func RequestAPI(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}