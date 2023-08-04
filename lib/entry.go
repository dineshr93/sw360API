package lib

import (
	"encoding/json"
	"errors"
	"io"
	"log"

	"net/http"
	"os"

	"github.com/dineshr93/sw360api/lib/model"

	yaml "gopkg.in/yaml.v3"
)

var (
	configfile    = "config.yml"
	contenttype   = "Content-Type"
	Authorization = "Authorization"
	appjson       = "application/json"
	apphaljson    = "application/hal+json"
	appxml        = "application/xml"
	mulformdata   = "multipart/form-data"

	API_URL = ""
	Token   = ""
)

func Sw360(configfile string) *Config {
	cfg := LoadConfig(configfile)

	API_URL := cfg.API
	Token := cfg.Token
	log.Println(API_URL)
	log.Println(Token)

	return cfg
}

type Config struct {
	API   string `yaml:"api"`
	Token string `yaml:"token"`
}

func (c *Config) GetTotalProjects() {

	req, err := http.NewRequest(http.MethodGet, c.API+"projects", nil)

	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Add(contenttype, apphaljson)
	req.Header.Add(Authorization, c.Token)

	// Create an HTTP client
	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		log.Fatalln("error while client.Do(req)")
	}

	databytes, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln("Couln't read response body")
	}

	var project model.Project

	err = json.Unmarshal(databytes, &project)
	if err != nil {
		log.Fatalln("Error while unmarshalling json")
	}

	log.Println(len(project.Embedded.Sw360Projects))

}

func LoadConfig(config string) *Config {
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
