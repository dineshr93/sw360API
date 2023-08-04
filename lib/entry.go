package lib

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"

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
	ascii   = `
	░██████╗░██╗░░░░░░░██╗██████╗░░█████╗░░█████╗░  ░█████╗░██████╗░██╗
	██╔════╝░██║░░██╗░░██║╚════██╗██╔═══╝░██╔══██╗  ██╔══██╗██╔══██╗██║
	╚█████╗░░╚██╗████╗██╔╝░█████╔╝██████╗░██║░░██║  ███████║██████╔╝██║
	░╚═══██╗░░████╔═████║░░╚═══██╗██╔══██╗██║░░██║  ██╔══██║██╔═══╝░██║
	██████╔╝░░╚██╔╝░╚██╔╝░██████╔╝╚█████╔╝╚█████╔╝  ██║░░██║██║░░░░░██║
	╚═════╝░░░░╚═╝░░░╚═╝░░╚═════╝░░╚════╝░░╚════╝░  ╚═╝░░╚═╝╚═╝░░░░░╚═╝`
)

func Sw360(configfile string) *Config {
	fmt.Println(ascii)

	cfg := LoadConfig(configfile)

	// API_URL := cfg.API
	// Token := cfg.Token
	// log.Println(API_URL)
	// log.Println(Token)

	return cfg
}
func LoadConfig(config string) *Config {
	var cfg Config = Config{}
	if _, err := os.Stat("./" + config); err == nil {
		log.Println("Using Config file " + config)
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

type Config struct {
	API   string `yaml:"api"`
	Token string `yaml:"token"`
}

func (c *Config) GetTotalProjects() int {

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

	// log.Println(len(project.Embedded.Sw360Projects))
	return len(project.Embedded.Sw360Projects)

}

func (c *Config) FindProjectswithName(pjnameguess string) (int, []string) {

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

	// log.Println(len(project.Embedded.Sw360Projects))
	Sw360Projects := project.Embedded.Sw360Projects
	var names []string

	count := 0

	for _, p := range Sw360Projects {
		if strings.Contains(strings.ToLower(p.Name), strings.ToLower(pjnameguess)) {

			count++
			nv := strconv.Itoa(count) + ". " + p.Name + " " + p.Version
			names = append(names, nv)

		}
	}

	return count, names

}

func (c *Config) GetProjectDetails(pjname string, version string) (error, *model.ProjectDetail) {

	req, err := http.NewRequest(http.MethodGet, c.API+"projects?allDetails=true", nil)

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

	// log.Println(len(project.Embedded.Sw360Projects))
	Sw360Projects := project.Embedded.Sw360Projects

	for _, p := range Sw360Projects {
		if strings.ToLower(p.Name) == strings.ToLower(pjname) && strings.ToLower(p.Version) == strings.ToLower(version) {
			fmt.Println(prettyPrint(p))
			return nil, &p
		}
	}

	pjd := &model.ProjectDetail{}
	return errors.New("No project details found for " + pjname + " " + version), pjd
}
func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
