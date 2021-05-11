package main

import (
	issue "GOPL/Ch4/practice/4.14/issues"
	"html/template"
	"log"
	"net/http"
)

const templ = `
<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>User</th>
  <th>Title</th>
</tr>
{{range .Items}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
  <td>{{.State}}</td>
  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`

func main() {
	http.HandleFunc("/", Handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func Handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Println(err)
	}
	q := r.FormValue("key")
	result, err := issue.SearchIssues(q)
	if err != nil {
		log.Println(err)
	}
	tmpl := template.Must(template.New("test").Parse(templ))
	if err := tmpl.Execute(w, result); err != nil {
		log.Println(err)
	}
}
