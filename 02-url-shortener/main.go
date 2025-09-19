// copied from https://github.com/gophercises/urlshort/blob/master/main/main.go

package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	// "github.com/gophercises/urlshort"
)

func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the fallback
	yamlFile := flag.String("yamlFile", "paths.yaml", "a yaml file with the paths as a list of path/url pairings in the following format:\n- path: /some-path\n  url: example.com/some-url\n")
	flag.Parse()

	yaml, err := os.ReadFile(*yamlFile)
	if err != nil {
		panic(err)
	}

	yamlHandler, err := YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
