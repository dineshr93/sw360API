package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/a", test)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))

}
func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}
