package root

import (
	"net/http"
	"text/template"
)

func IndexHandler(response http.ResponseWriter, request *http.Request){
	t := template.Must(template.ParseFiles("./root/tmpl/index.html"))
	t.Execute(response, nil)
}
