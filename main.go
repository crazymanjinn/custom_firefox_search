package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
)

type OpenSearch struct {
	File  string        `-`
	Title string        `xml:"ShortName"`
	Raw   template.HTML `-`
}

var (
	xmlDir     string
	listenPort string
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	const (
		defaultDir  = "."
		dirUsage    = "directory containing OpenSearch XML files"
		defaultPort = "8080"
		portUsage   = "port to listen on"
	)
	flag.StringVar(&xmlDir, "dir", defaultDir, dirUsage)
	flag.StringVar(&xmlDir, "d", defaultDir, dirUsage+" (shorthand)")
	flag.StringVar(&listenPort, "port", defaultPort, portUsage)
	flag.StringVar(&listenPort, "p", defaultPort, portUsage+" (shorthand)")
}

func loadEngines(directory string) (*[]OpenSearch, error) {
	log.Printf("reading OpenSearch XML files from %s\n", directory)
	matches, err := filepath.Glob(filepath.Join(directory, "*.xml"))
	if err != nil {
		return nil, err
	}
	log.Printf("found OpenSearch files: %s\n", matches)

	engines := []OpenSearch{}
	for _, match := range matches {
		rawXml, err := ioutil.ReadFile(match)
		if err != nil {
			log.Printf("unable to read file %s: %+v\n", match, err)
			continue
		}
		opensearch := OpenSearch{
			File: filepath.Base(match),
			Raw:  template.HTML(rawXml),
		}
		if err := xml.Unmarshal(rawXml, &opensearch); err != nil {
			log.Printf("unable to unmarshal xml %s: %+v\n", match, err)
		}
		engines = append(engines, opensearch)
	}
	return &engines, nil
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	engines, err := loadEngines(xmlDir)
	if err != nil {
		fmt.Fprintf(w, "unable to list engines: %+v", err)
		return
	}

	t, err := template.New("index").Parse(indexTemplate)
	if err != nil {
		fmt.Fprintf(w, "unable to parse index.tmpl: %+v", err)
		return
	}
	t.Execute(w, engines)
}

func main() {
	flag.Parse()

	http.Handle("/xml/",
		http.StripPrefix("/xml/", http.FileServer(http.Dir(xmlDir))))

	http.HandleFunc("/", rootHandler)

	log.Printf("listening on port :%s\n", listenPort)
	log.Fatal(http.ListenAndServe(":"+listenPort, nil))
}
