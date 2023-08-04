package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/a", test).Methods(http.MethodGet)

	// restrict the url param to 0 to 9
	r.HandleFunc("/a/{b:[0-9]+}", test2).Methods(http.MethodGet)

	r.HandleFunc("/a/{b:[0-9]+}", test3).Methods(http.MethodPost)

	// http.HandleFunc("/a", test)
	log.Fatal(http.ListenAndServe(":80", r))

}
func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}
func test2(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, vars["b"])
}
func test3(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "Post "+vars["b"])
}
