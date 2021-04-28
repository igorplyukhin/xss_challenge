package lvl1

import (
	"net/http"
	"strings"
	"text/template"
	"xss_challenge/cookies"
)

func IndexHandler(response http.ResponseWriter, request *http.Request){
	keys, ok := request.URL.Query()["query"]
	indexTempl := template.Must(template.ParseFiles("./lvl1/templ/index.html"))
	responseTempl := template.Must(template.ParseFiles("./lvl1/templ/response.html"))
	successResponseTempl := template.Must(template.ParseFiles("./lvl1/templ/successResponse.html"))
	if !ok || len(keys[0]) < 1 {
		indexTempl.Execute(response, nil)
	}else if strings.Contains(keys[0], "<script>") &&
		strings.Contains(keys[0], "alert") &&
		strings.Contains(keys[0], "</script>"){
		cookies.SetSession("lvl2","aboba", "access", response)
		successResponseTempl.Execute(response, keys[0])
	}else {
		responseTempl.Execute(response, keys[0])
	}
}
