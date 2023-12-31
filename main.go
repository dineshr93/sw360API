package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	lib "github.com/dineshr93/sw360api/lib"
	"github.com/dineshr93/sw360api/lib/model"
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

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
func main() {
	apicfg := lib.Sw360(configfile)

	fmt.Println("No of projects in server: ", apicfg.GetTotalProjects())

	fmt.Println("=================Testing FindProjectswithName=======================")
	no, names := apicfg.FindProjectswithName("ada")

	fmt.Println("No of projects matched is: ", no)

	for _, name := range names {
		fmt.Println(name)
	}

	fmt.Println("=================Testing Testing GetProjectDetails=======================")
	// err, data := apicfg.GetProjectDetails("ZF_Middleware_Program", "IOLITHFBL_0500")
	err, ProjectDetail := apicfg.GetProjectDetails("gradle_single", "1.0")
	if err != nil {
		log.Fatalln("Error getting details for given project")
	} else {

		fmt.Println(ProjectDetail.LinkedProjects)
	}
	fmt.Println("===================Testing GetLinkedReleases=====================")
	err, releases := apicfg.GetLinkedReleases("gradle_single", "1.0")
	if err != nil {
		log.Fatalln(err)
	} else {

		for i, release := range *releases {
			i++

			fmt.Println(strconv.Itoa(i), release.Name, " ", release.Version)
		}
	}
	fmt.Println("=================Testing GetLinkedProjects=======================")
	err, projects := apicfg.GetLinkedProjects("gradle_single", "1.0")
	if err != nil {
		log.Fatalln(err)
	} else {
		for i, project := range *projects {
			i++
			fmt.Println(strconv.Itoa(i), project.Name, " ", project.Version)
		}
	}
	fmt.Println("================Testing GetProjectlink========================")
	err, projectlink := apicfg.GetProjectlink("gradle_single", "1.0")
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println(projectlink)
	}
	fmt.Println("=================Testing GetLinkedReleasesTransitive=======================")
	err, releases = apicfg.GetLinkedReleasesTransitive("gradle_single", "1.0")
	if err != nil {
		log.Fatalln(err)
	} else {

		for i, release := range *releases {
			i++

			fmt.Println(strconv.Itoa(i), release.Name, " ", release.Version)
		}
	}
	fmt.Println("==================Testing CreateProject======================")
	data := model.ProjectCreationModel{Name: "fromDinesh", Version: "2.0", Description: "Sample Project"}

	err, msg := apicfg.CreateProject(&data)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(prettyPrint(msg))
	fmt.Println("=================Testing DeleteProject=======================")
	errormsg, msgdelete := apicfg.DeleteProject("fromDinesh", "2.0")
	if errormsg != nil {
		log.Fatalln(errormsg)
	}
	fmt.Println(msgdelete)
	fmt.Println("========================================")

	fmt.Println("========================================")
}
