package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"gophercises/urlshort"
)

func getFileBytes(fileName string) []byte {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Could not open file %s", fileName)
	}

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(f)
	if err != nil {
		log.Fatalf("Could not read file %s", fileName)
	}

	return buf.Bytes()
}

func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// yamlFileName := flag.String("yml", "main/sample.yml", "a yml file with url and path")
	jsonFileName := flag.String("json", "main/sample.json", "a json file with url and path")
	flag.Parse()

	// yaml := getFileBytes(*yamlFileName)
	// yamlHandler, err := urlshort.YAMLHandler(yaml, mapHandler)
	// if err != nil {
	// 	panic(err)
	// }

	json := getFileBytes(*jsonFileName)
	jsonHandler, err := urlshort.JsonHandler(json, mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", jsonHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
