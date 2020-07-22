package main

import "net/http"

func lesson1Run() {
	println("__lesson1Run__")
	// http.HandleFunc("/basepage", basePage)
	oneSampleTwo()
}

func oneSampleOne() {
	http.HandleFunc("/basepage", basePage)
}
func oneSampleTwo() {
	theUser := &user{name: "Fajar"}

	http.ListenAndServe(PORT, theUser)
}

type user struct {
	name string
}

//implemented method(ServeHTTP)
func (u *user) ServeHTTP(w http.ResponseWriter, res *http.Request) {
	w.Write([]byte("Name: " + u.name))
}

func basePage(w http.ResponseWriter, res *http.Request) {
	w.Write([]byte("Hello World"))
}
