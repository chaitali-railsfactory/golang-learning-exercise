package urlshort

import (
	"encoding/json"
	"net/http"

	yaml "gopkg.in/yaml.v2"
)

type PathUrl struct {
	URL  string `yaml:"url"`
	Path string `yaml:"path"`
}

type PathToUrl struct {
	URL  string
	Path string
}

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
		}
		fallback.ServeHTTP(w, r)
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.

func parseYaml(ymlBytes []byte) ([]PathUrl, error) {
	var pathUrls []PathUrl
	err := yaml.Unmarshal(ymlBytes, &pathUrls)
	if err != nil {
		return nil, err
	}
	return pathUrls, nil
}

func buildMap(pathUrls []PathUrl) map[string]string {
	pathToUrls := make(map[string]string)
	for _, pu := range pathUrls {
		pathToUrls[pu.Path] = pu.URL
	}
	return pathToUrls
}

func YAMLHandler(ymlBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {
	pathUrls, err := parseYaml(ymlBytes)
	if err != nil {
		return nil, err
	}
	pathToUrls := buildMap(pathUrls)
	return MapHandler(pathToUrls, fallback), nil
}

func parseJSON(jsonData []byte) (pathsToURLs []PathUrl, err error) {
	err = json.Unmarshal(jsonData, &pathsToURLs)
	return
}

func JsonHandler(jsonData []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedJSON, err := parseJSON(jsonData)
	if err != nil {
		return nil, err
	}
	pathToUrls := buildMap(parsedJSON)
	return MapHandler(pathToUrls, fallback), nil
}
