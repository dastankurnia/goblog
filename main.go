package main

import (
	"log"
	"net/http"

	"github.com/dastankurnia/goblog/common"
	"github.com/dastankurnia/goblog/modules/user"
	"github.com/gorilla/mux"
)

// NotFound : Handle not found route
func NotFound(w http.ResponseWriter, r *http.Request) {
	common.Response(w, r, 404, "Not found!", []string{})
}

// DefaultHandler : Handle default page
func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	common.Response(w, r, 200, "Welcome to jungle!", []string{})
}

func main() {

	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/register", user.Register).Methods("POST")
	api.HandleFunc("/login", user.Login).Methods("POST")
	api.HandleFunc("/user", user.GetUsers).Methods("GET")
	api.NotFoundHandler = r.NewRoute().HandlerFunc(NotFound).GetHandler()
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":9999", nil))
}
