package main

import (
	"encoding/json"
	"fmt"
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

// temporary code to test json parsing and http serving
// TODO replace with html template
func (a *StoryArc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	arc, err := json.Marshal(a)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(arc)
	}
}

type Story map[string]StoryArc

func parseJSON(raw []byte) Story {
	jsonMap := Story{}
	json.Unmarshal(raw, &jsonMap)
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

func main() {
	rawJSON, err := os.ReadFile("gopher.json")
	if err != nil {
		panic(err)
	}

	story := parseJSON(rawJSON)
	mux := http.NewServeMux()
	registerHandle(story, mux)

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", mux)
}
