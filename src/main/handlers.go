package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const PORT string = ":8080"

func main() {
	println("__lesson2Handler__")
	http.Handle("/", new(myHandler))
	http.ListenAndServe(PORT, nil)
}

type myHandler struct {
}

func hasSuffix(str, suffix string) bool {
	return strings.HasSuffix(str, suffix)
}

func getContentType(path string) string {
	if hasSuffix(path, ".css") {
		return "text/css"
	} else if hasSuffix(path, ".html") {
		return "text/html"
	} else if hasSuffix(path, ".js") {
		return "application/javascript"
	} else if hasSuffix(path, ".png") {
		return "image/png"
	} else if hasSuffix(path, ".svc") {
		return "image/svg"
	}
	return "text/plain"
}

func (h *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	path := r.URL.Path[1:]
	log.Println("PATH:" + path)
	path = "../../" + path
	data, err := ioutil.ReadFile(string(path))

	if err == nil {
		w.Header().Add("Content-Type", getContentType(path))
		w.Write(data)
	} else {
		log.Println(err)
		w.WriteHeader(404)
		w.Write([]byte("NOT FOUND BRO - " + http.StatusText(404)))
	}
}
