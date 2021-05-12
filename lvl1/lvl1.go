package lvl1

import (
	"net/http"
	"text/template"
	"xss_challenge/checker"
	"xss_challenge/cookies"
)

var indexTempl = template.Must(template.ParseFiles("./lvl1/templ/index.html"))
var responseTempl = template.Must(template.ParseFiles("./lvl1/templ/response.html"))
var successResponseTempl = template.Must(template.ParseFiles("./lvl1/templ/successResponse.html"))

func IndexHandler(response http.ResponseWriter, request *http.Request) {
	query, queryOk := request.URL.Query()["query"]
	bot, botOk := request.URL.Query()["bot"]

	if queryOk && len(query[0]) > 0 {
		if botOk && len(bot[0]) > 0 {
			responseTempl.Execute(response, query[0])
		} else if checker.PayloadWasExecuted(request, "") {
			cookies.SetCookie("lvl2", "access", response)
			successResponseTempl.Execute(response, query[0])
		} else {
			responseTempl.Execute(response, query[0])
		}
	} else {
		indexTempl.Execute(response, nil)
	}
}
