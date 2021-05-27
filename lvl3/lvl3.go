package lvl3

import (
	"net/http"
	"strings"
	"text/template"
	"xss_challenge/checker"
	"xss_challenge/cookies"
)

var indexTempl = template.Must(template.ParseFiles("./lvl3/templ/index.html"))
var responseTempl = template.Must(template.ParseFiles("./lvl3/templ/response.html"))
var successResponseTempl = template.Must(template.ParseFiles("./lvl3/templ/successResponse.html"))
var mock = template.Must(template.ParseFiles("./root/tmpl/mock.html"))

func IndexHandler(response http.ResponseWriter, request *http.Request) {
	query, queryOk := request.URL.Query()["query"]
	bot, botOk := request.URL.Query()["bot"]
	access := cookies.GetCookie("lvl3", request)

	if !botOk && access != "access"{
		mock.Execute(response, nil)
		return
	}

	if queryOk && len(query[0]) > 0 {
		query := strings.ReplaceAll(strings.ReplaceAll(query[0], "<script>", ""), "</script>", "")
		if botOk && len(bot[0]) > 0 {
			responseTempl.Execute(response, query)
		} else if checker.PayloadWasExecuted(request, "") {
			cookies.SetCookie("lvl4", "access", response)
			successResponseTempl.Execute(response, query)
		} else {
			responseTempl.Execute(response, query)
		}
	} else {
		indexTempl.Execute(response, nil)
	}
}


