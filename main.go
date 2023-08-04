package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	yaml "gopkg.in/yaml.v3"
)

type Config struct {
	API   string `yaml:"api"`
	Token string `yaml:"token"`
}

var (
	configfile  = "config.yml"
	contenttype = "Content-Type"
	appjson     = "application/json"
	apphaljson  = "application/hal+json"
	appxml      = "application/xml"
	mulformdata = "multipart/form-data"

	API_KEY = ""
	Token   = ""
)

func main() {
	cfg := loadConfig(configfile)

	API_KEY := cfg.API
	Token := cfg.Token
	log.Println(API_KEY)
	log.Println(Token)
	r := mux.NewRouter()
	r.HandleFunc("/a", test).Methods(http.MethodGet)

	// restrict the url param to 0 to 9
	r.HandleFunc("/a/{b:[0-9]+}", test2).Methods(http.MethodGet)

	r.HandleFunc("/a/{b:[0-9]+}", test3).Methods(http.MethodPost)

	// http.HandleFunc("/a", test)
	log.Fatal(http.ListenAndServe(":8080", r))

}

func loadConfig(config string) *Config {
	var cfg Config = Config{}
	if _, err := os.Stat("./" + config); err == nil {
		log.Println("Config file " + config + " exists")
		// if file exists
		f, err := os.Open("./" + config)
		if err != nil {
			log.Println(err)
		}
		defer f.Close()

		decoder := yaml.NewDecoder(f)
		err = decoder.Decode(&cfg)
		if err != nil {
			log.Fatalln(err)
		}

	} else if errors.Is(err, os.ErrNotExist) {
		// path/to/whatever does *not* exist
		log.Fatalln("./" + config + " File does *not* exist")

	}

	return &cfg
}

func test(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello")
}

type Name struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func test2(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// fmt.Fprintf(w, vars["b"])

	//
	w.Header().Add(contenttype, appjson)

	// xml encoding
	// xml.NewEncoder(w).Encode()
	n := Name{vars["b"], "Dinesh"}
	// json encoding
	json.NewEncoder(w).Encode(n)
	// fmt.Fprintf(w, r.Header.Get(ct))
	// fmt.Fprintf(w, strings.Join(, ","))
}
func test3(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "Post "+vars["b"])
}
