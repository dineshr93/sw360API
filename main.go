package main

import (
	"fmt"

	lib "github.com/dineshr93/sw360api/lib"
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

func main() {
	apicfg := lib.Sw360(configfile)

	fmt.Println("No of projects in server: ", apicfg.GetTotalProjects())

	no, names := apicfg.FindProjectswithName("ada")

	fmt.Println("No of projects matched is: ", no)

	for _, name := range names {
		fmt.Println(name)
	}

	v := apicfg.GetProjectDetails("ZF_Middleware_Program", "IOLITHFBL_0500")

	fmt.Println(v.)

}
