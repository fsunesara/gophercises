package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type StoryArc struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

func (a *StoryArc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := tmpl.Execute(w, a)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
}

type Story map[string]StoryArc

func parseJSON(rawJSON []byte) Story {
	jsonMap := Story{}
	json.Unmarshal(rawJSON, &jsonMap)
	return jsonMap
}

func registerHandle(story Story, mux *http.ServeMux) {
	for a := range story {
		path := "/"
		if a != "intro" {
			path += a
		}
		if arc, ok := story[a]; ok {
			mux.Handle(path, &arc)
		}
	}
}

var tmpl *template.Template

func main() {
	templateFile := flag.String("templateFile", "layout.html", "the html template file to use")
	jsonFile := flag.String("jsonFile", "gopher.json", "the json file of the story to use")
	flag.Parse()

	tmpl = template.Must(template.ParseFiles(*templateFile))
	rawJSON, err := os.ReadFile(*jsonFile)
	if err != nil {
		panic(err)
	}

	story := parseJSON(rawJSON)
	mux := http.NewServeMux()
	registerHandle(story, mux)

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", mux)
}
