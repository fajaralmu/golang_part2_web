package main

import (
	"net/http"
	"text/template"
)

const PORT string = ":8080"

const htmlTemplate = `
<!DOCTYPE html>
<html>
	<head><title>{{.Title}}</title>
	</head><body>
	<p>
	{{.Message}}
	</p>
	<p>
	ID: {{.ID}}
	{{if eq .ID 3}}
	id is 3
	{{else}}
	id not 3
	{{end}}
	</p>
	</body>
</html>
`

func main() {
	http.HandleFunc("/", oneFunc)
	http.ListenAndServe(PORT, nil)
}

type pageObject struct {
	ID      int
	Title   string
	Message string
}

func oneFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	html, err := template.New("templateName").Parse(htmlTemplate)
	if err == nil {
		_pageObject := pageObject{3, "Page One", "Hello Go"}
		html.Execute(w, _pageObject)
	}
}