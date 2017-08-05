package web

import (
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func Register(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	redirectURL := vars["redirect-url"]
	fmt.Fprintln(w, "Redirect URL:", redirectURL)
}