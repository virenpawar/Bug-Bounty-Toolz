package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", Forwarder).Methods(http.MethodGet)
	return r
}

func Forwarder(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()["url"]
	http.Redirect(w, r, q[0], http.StatusFound)
}

func main() {
	log.Fatal(http.ListenAndServe("0.0.0.0:88", NewRouter()))
}
