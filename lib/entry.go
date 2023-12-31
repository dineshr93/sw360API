package lib

import (
	"bytes"
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

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

// ============================== API =======================================
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
			// fmt.Println(prettyPrint(p))
			return nil, &p
		}
	}

	// pjd := &model.ProjectDetail{}
	return errors.New("No project details found for " + pjname + " " + version), nil
}

func (c *Config) GetProjectlink(pjname string, version string) (error, string) {

	err, projectDetails := c.GetProjectDetails(pjname, version)
	if err != nil {
		return err, ""
	}

	projectSelfLink := projectDetails.Links.Self.Href

	if len(projectSelfLink) < 5 {
		return errors.New("Invalid projectSelfLink"), ""
	}
	return nil, projectSelfLink

}

func (c *Config) GetLinkedReleases(pjname string, version string) (error, *[]model.Release) {
	err, projectDetail := c.GetProjectDetails(pjname, version)
	if err != nil {
		log.Fatalln("Error while getting project details in GetLinkedReleases")
	}
	linkedReleases := projectDetail.LinkedReleases
	// Create an HTTP client
	client := &http.Client{}
	var releases []model.Release
	if len(linkedReleases) == 0 {
		return errors.New("No release details found for " + pjname + " " + version), nil
	}
	// loop through all the release
	for _, linkedRelease := range linkedReleases {
		releaselink := linkedRelease.Release
		req, err := http.NewRequest(http.MethodGet, releaselink, nil)
		if err != nil {
			log.Fatalln("error in releaselink")
		}
		req.Header.Add(contenttype, apphaljson)
		req.Header.Add(Authorization, c.Token)
		res, err := client.Do(req)
		if err != nil {
			log.Fatalln("error while client.Do(req) in GetLinkedReleases")
		}
		databytes, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatalln("Couln't read response body")
		}
		var release model.Release
		err = json.Unmarshal(databytes, &release)
		if err != nil {
			log.Fatalln("Error while unmarshalling json")
		}
		releases = append(releases, release)
	}
	return nil, &releases
}

func (c *Config) GetLinkedReleasesTransitive(pjname string, version string) (error, *[]model.Release) {
	err, projectlink := c.GetProjectlink(pjname, version)
	if err != nil {
		log.Fatalln(err)
		return errors.New("No projectlink found for " + pjname + " " + version), nil
	}

	req, err := http.NewRequest(http.MethodGet, projectlink+"/releases?transitive=true", nil)
	if err != nil {
		log.Fatalln("error in request preparation in GetLinkedReleasesTransitive")
	}
	req.Header.Add(contenttype, apphaljson)
	req.Header.Add(Authorization, c.Token)
	// Create an HTTP client
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatalln("error while client.Do(req) in GetLinkedReleasesTransitive")
	}
	databytes, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln("Couln't read response body")
	}
	var releaseshort model.SW360ReleaseShort
	err = json.Unmarshal(databytes, &releaseshort)
	if err != nil {
		log.Fatalln("Error while unmarshalling json")
	}

	linkedReleases := releaseshort.Embedded.Sw360Releases

	var releases []model.Release
	if len(linkedReleases) == 0 {
		return errors.New("No release details found for " + pjname + " " + version), nil
	}
	// loop through all the release
	for _, linkedRelease := range linkedReleases {
		releaselink := linkedRelease.Links.Self.Href
		req, err := http.NewRequest(http.MethodGet, releaselink, nil)
		if err != nil {
			log.Fatalln("error in releaselink")
		}
		req.Header.Add(contenttype, apphaljson)
		req.Header.Add(Authorization, c.Token)
		res, err := client.Do(req)
		if err != nil {
			log.Fatalln("error while client.Do(req) in GetLinkedReleases")
		}
		databytes, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatalln("Couln't read response body")
		}
		var release model.Release
		err = json.Unmarshal(databytes, &release)
		if err != nil {
			log.Fatalln("Error while unmarshalling json")
		}
		releases = append(releases, release)
	}
	return nil, &releases
}

func (c *Config) GetLinkedProjects(pjname string, version string) (error, *[]model.Sw360Project) {
	err, projectDetail := c.GetProjectDetails(pjname, version)
	if err != nil {
		log.Fatalln("Error while getting project details in GetLinkedProjects")
	}
	linkedProjects := projectDetail.LinkedProjects
	// Create an HTTP client
	client := &http.Client{}
	var projects []model.Sw360Project
	if len(linkedProjects) == 0 {
		return errors.New("No linked projecDetails found for " + pjname + " " + version), nil
	}
	// loop through all the release
	for _, linkedProject := range linkedProjects {
		projectlink := linkedProject.Project
		req, err := http.NewRequest(http.MethodGet, projectlink, nil)
		if err != nil {
			log.Fatalln("error in projectlink")
		}
		req.Header.Add(contenttype, apphaljson)
		req.Header.Add(Authorization, c.Token)
		res, err := client.Do(req)
		if err != nil {
			log.Fatalln("error while client.Do(req) in GetLinkedProjects")
		}
		databytes, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatalln("Couln't read response body")
		}
		var project model.Sw360Project
		err = json.Unmarshal(databytes, &project)
		if err != nil {
			log.Fatalln("Error while unmarshalling json")
		}
		projects = append(projects, project)
	}
	return nil, &projects
}

func (c *Config) DeleteProject(pjname string, version string) (error, string) {
	err, projectlink := c.GetProjectlink(pjname, version)
	if err != nil {
		log.Fatalln(err)
		return errors.New("No projectlink found for " + pjname + " " + version), ""
	}

	req, err := http.NewRequest(http.MethodDelete, projectlink, nil)
	if err != nil {
		log.Fatalln("error in request preparation in DeleteProject")
	}
	req.Header.Add(contenttype, apphaljson)
	req.Header.Add(Authorization, c.Token)
	// Create an HTTP client
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatalln("error while client.Do(req) in DeleteProject")
		return err, ""
	}
	if res.StatusCode != http.StatusOK {
		return errors.New("error while client.Do(req) in DeleteProject for " + pjname + " " + version), ""
	}
	return nil, pjname + " " + version + " deleted successfully"
}

func (c *Config) CreateProject(pjdata *model.ProjectCreationModel) (error, *model.ProjectCreated) {

	if pjdata.Name == "" && pjdata.Version == "" {
		return errors.New("Mandatory data missing: project name or version"), nil
	}

	// JSON body
	pjdatabody, err := json.Marshal(pjdata)
	if err != nil {
		return err, nil
	}

	req, err := http.NewRequest(http.MethodPost, c.API+"projects", bytes.NewBuffer(pjdatabody))
	if err != nil {
		log.Fatalln("error in request preparation in DeleteProject")
	}
	req.Header.Add(contenttype, apphaljson)
	req.Header.Add(Authorization, c.Token)
	// Create an HTTP client
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatalln("error while client.Do(req) in CreateProject")
		return err, nil
	}
	if res.StatusCode != http.StatusCreated {
		return errors.New("error while client.Do(req) in CreateProject"), nil
	}
	databytes, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln("Couln't read response body")
	}
	var projectCreated model.ProjectCreated
	err = json.Unmarshal(databytes, &projectCreated)
	if err != nil {
		log.Fatalln("Error while unmarshalling json")
		return errors.New("Error while unmarshalling json in CreateProject"), nil
	}
	log.Println("Successfully created Project: " + pjdata.Name + " " + pjdata.Version)
	return nil, &projectCreated
}
