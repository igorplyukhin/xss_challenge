package lvl2

import (
	"net/http"
	"text/template"
	"xss_challenge/checker"
	"xss_challenge/cookies"
)

var indexTempl = template.Must(template.ParseFiles("./lvl2/templ/index.html"))
var responseTempl = template.Must(template.ParseFiles("./lvl2/templ/response.html"))
var successResponseTempl = template.Must(template.ParseFiles("./lvl2/templ/successResponse.html"))
var mock = template.Must(template.ParseFiles("./root/tmpl/mock.html"))

func IndexHandler(response http.ResponseWriter, request *http.Request) {
	query, queryOk := request.URL.Query()["query"]
	bot, botOk := request.URL.Query()["bot"]
	access := cookies.GetCookie("lvl2", request)

	if !botOk && access != "access"{
		mock.Execute(response, nil)
		return
	}

	if queryOk && len(query[0]) > 0 {
		if botOk && len(bot[0]) > 0 {
			responseTempl.Execute(response, query[0])
		} else if checker.PayloadWasExecuted(request, "") {
			cookies.SetCookie("lvl3", "access", response)
			successResponseTempl.Execute(response, query[0])
		} else {
			responseTempl.Execute(response, query[0])
		}
	} else {
		indexTempl.Execute(response, nil)
	}
}

