package lvl6

import (
	"fmt"
	"html/template"
	"net/http"
	"xss_challenge/checker"
	"xss_challenge/cookies"
)

var indexTempl = template.Must(template.ParseFiles("./lvl6/templ/index.html"))
var successResponseTempl = template.Must(template.ParseFiles("./lvl6/templ/successResponse.html"))
var mock = template.Must(template.ParseFiles("./root/tmpl/mock.html"))

func IndexHandler(response http.ResponseWriter, request *http.Request) {
	a := cookies.GetCookie("lvl6", request)
	query, queryOk := request.URL.Query()["query"]
	bot, botOk := request.URL.Query()["bot"]

	if !botOk && a != "access" {
		mock.Execute(response, nil)
		return
	}

	if queryOk && len(query[0]) > 0 {
		respTmpl := fmt.Sprintf(`
<!doctype html>
<html>
  <head>
  </head>
  <body id="level1">
      <div>
        Ой, ничего не найдено по запросу: <b>%s</b>. <a href='?'>Попробовать еще раз</a>.
    </div>
  </body>
</html>
`, query[0])
		r, _ := template.New("resp").Parse(respTmpl)
		if botOk && len(bot[0]) > 0 {
			r.Execute(response, nil)
		} else if checker.PayloadWasExecuted(request, "") {
			successResponseTempl.Execute(response, query[0])
		} else {
			r.Execute(response, nil)
		}
	} else {
		indexTempl.Execute(response, nil)
	}
}
