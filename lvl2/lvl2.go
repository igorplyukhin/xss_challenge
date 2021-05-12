package lvl2

import (
	"net/http"
	"text/template"
	"xss_challenge/checker"
)

var indexTempl = template.Must(template.ParseFiles("./lvl2/templ/index.html"))
var successTempl = template.Must(template.ParseFiles("./lvl2/templ/success.html"))

func IndexHandler(response http.ResponseWriter, request *http.Request) {
	//a :=  GetCookie("lvl2", request)
	//if a == "access"{
	bot, botOk := request.URL.Query()["bot"]
	hash, _ :=request.Cookie("hash")
	if botOk && len(bot[0]) > 0 {
		indexTempl.Execute(response, nil)
	} else if checker.PayloadWasExecuted(request, "#" + hash.Value) {
		successTempl.Execute(response, nil)
	} else {
		indexTempl.Execute(response, nil)
	}
	//}else {
	//	t := template.Must(template.ParseFiles("./root/tmpl/mock.html"))
	//	t.Execute(response, nil)
	//
	//}
}
