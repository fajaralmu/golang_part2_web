package main

import (
	"net/http"
	"text/template"
)

const PORT string = ":8080"

const htmlTemplate = `
{{template "header" .Title}}
<body>
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
	<hr/>
	<p>Articles</p>
	<ul>
		{{range .Articles}}
		<li>{{.}}</li> 
		{{end}}
	</ul>
	{{template "footer"}}
`

const head = `
<!DOCTYPE html>
<html>
	<head><title>{{.}}</title>
	</head>
`

const foot = `
<div>FOOTER</div>
</body>
</html>
`

func main() {
	http.HandleFunc("/", oneFunc)
	http.ListenAndServe(PORT, nil)
}

type pageObject struct {
	ID       int
	Title    string
	Message  string
	Articles []string
}

func oneFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	//html, err := template.New("templateName").Parse(htmlTemplate)

	html := template.New("myTemplate")
	// if err == nil {
	html.New("page").Parse(htmlTemplate)
	html.New("header").Parse(head)
	html.New("footer").Parse(foot)
	_pageObject := pageObject{3, "Page One", "Hello Go", []string{
		"News #1", "News #2", "News #3",
	}}
	html.Lookup("page").Execute(w, _pageObject)
	// html.Execute(w, _pageObject)
	// }
}
