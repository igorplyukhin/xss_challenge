package lvl2

import (
	"net/http"
	"text/template"
	. "xss_challenge/cookies"
)

func LoginHandler(response http.ResponseWriter, request *http.Request) {
	name := request.FormValue("name")
	pass := request.FormValue("password")
	redirectTarget := "."
	if name != "" && pass != "" {
		SetSession("session", "name", name, response)
		redirectTarget = "./internal"
	}
	http.Redirect(response, request, redirectTarget, 302)
}

func LogoutHandler(response http.ResponseWriter, request *http.Request) {
	ClearSession(response)
	http.Redirect(response, request, ".", 302)
}

func IndexHandler(response http.ResponseWriter, request *http.Request) {
	a :=  GetCookie("lvl2","aboba", request)
	if a == ""{
		t := template.Must(template.ParseFiles("./root/tmpl/mock.html"))
		t.Execute(response, nil)
	}else {
		t := template.Must(template.ParseFiles("./lvl2/templ/index.html"))
		t.Execute(response, nil)
	}
}

func InternalHandler(response http.ResponseWriter, request *http.Request) {
	userName := GetCookie("session", "name", request)
	if userName != "" {
		t := template.Must(template.ParseFiles("./lvl2/templ/internal.html"))
		t.Execute(response, userName)
	} else {
		http.Redirect(response, request, "./", 302)
	}
}
