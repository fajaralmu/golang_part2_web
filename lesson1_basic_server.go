package main

import "net/http"

func lesson1Run() {
	println("__lesson1Run__")
	multiPlexor := http.NewServeMux()
	http.HandleFunc("/", basePage)
	http.ListenAndServe(PORT, multiPlexor)
}

func basePage(w http.ResponseWriter, res *http.Request) {
	w.Write([]byte("Hello World"))
}
