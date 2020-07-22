package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func lesson2Handler() {
	println("__lesson2Handler__")
	http.Handle("/", new(myHandler))
	http.ListenAndServe(PORT, nil)
}

type myHandler struct {
}

func (h *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	path := r.URL.Path[1:]
	log.Println("PATH:" + path)

	data, err := ioutil.ReadFile(string(path))

	if err == nil {
		w.Write(data)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("NOT FOUND BRO - " + http.StatusText(404)))
	}
}
