package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	cfenv "github.com/cloudfoundry-community/go-cfenv"
)

// Index holds fields displayed on the index.html template
type Index struct {
	AppName          string
	AppInstanceIndex int
	AppInstanceGUID  string
	Envars           []string
	Services         []Service
	SpaceName        string
}

// Service holds the name and label of a service instance
type Service struct {
	Name  string
	Label string
}

func main() {

	index := Index{"Unknown", -1, "Unknown", []string{}, []Service{}, "Unknown"}

	template := template.Must(template.ParseFiles("templates/index.html"))

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))

	if cfenv.IsRunningOnCF() {
		appEnv, err := cfenv.Current()
		if err != nil {
			log.Fatal(err)
		}
		if appEnv.Name != "" {
			index.AppName = appEnv.Name
		}
		if appEnv.Index > -1 {
			index.AppInstanceIndex = appEnv.Index
		}
		if appEnv.InstanceID != "" {
			index.AppInstanceGUID = appEnv.InstanceID
		}
		if appEnv.SpaceName != "" {
			index.SpaceName = appEnv.SpaceName
		}
		for _, svcs := range appEnv.Services {
			for _, svc := range svcs {
				index.Services = append(index.Services, Service{svc.Name, svc.Label})
			}
		}
		for _, envar := range os.Environ() {
			if strings.HasPrefix(envar, "TRAINING_") {
				index.Envars = append(index.Envars, envar)
			}
		}
	}

	//grow to 128Mb
	memSize := 128 * 1024 * 1024
	memory := make([]byte, memSize)
	i := 0

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate") // HTTP 1.1.
		w.Header().Set("Pragma", "no-cache")                                   // HTTP 1.0.
		w.Header().Set("Expires", "0")                                         // Proxies.
		if err := template.ExecuteTemplate(w, "index.html", index); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		memory[i] = 1
		fmt.Fprintf(os.Stderr, "Allocated %d MB of memory\n", i)
		for k := i; k < 1024*1024; k += 1024 {
			memory[i+k] = 1
		}
		i += 1020 * 1024

	})

	http.HandleFunc("/kill", func(w http.ResponseWriter, r *http.Request) {
		os.Exit(1)
	})

	var PORT string
	if PORT = os.Getenv("PORT"); PORT == "" {
		PORT = "8080"
	}

	fmt.Println(http.ListenAndServe(":"+PORT, nil))
}
