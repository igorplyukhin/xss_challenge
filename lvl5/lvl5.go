package lvl5

import (
	"net/http"
	"text/template"
	"xss_challenge/checker"
	"xss_challenge/cookies"
)

var indexTempl = template.Must(template.ParseFiles("./lvl5/templ/index.html"))
var successTempl = template.Must(template.ParseFiles("./lvl5/templ/successResponse.html"))
var mock = template.Must(template.ParseFiles("./root/tmpl/mock.html"))

func IndexHandler(response http.ResponseWriter, request *http.Request) {
	a := cookies.GetCookie("lvl5", request)

	bot, botOk := request.URL.Query()["bot"]
	hash, err := request.Cookie("hash")
	hashVal := "123"

	if err == nil {
		hashVal = hash.Value
	}
	if !botOk && a != "access" {
		mock.Execute(response, nil)
		return
	}
	if botOk && len(bot[0]) > 0 {
		indexTempl.Execute(response, nil)
	} else if checker.PayloadWasExecuted(request, "#"+hashVal) {
		cookies.SetCookie("lvl6", "access", response)
		successTempl.Execute(response, nil)
	} else {
		indexTempl.Execute(response, nil)
	}
}
