package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"xss_challenge/lvl1"
	"xss_challenge/lvl2"
	"xss_challenge/root"
)

func main() {
	var r = mux.NewRouter()

	var rootRouter = r.PathPrefix("").Subrouter()
	rootRouter.HandleFunc("/", root.IndexHandler)

	var lvl1Router = r.PathPrefix("/lvl1").Subrouter()
	lvl1Router.HandleFunc("", lvl1.IndexHandler)

	var lvl2Router = r.PathPrefix("/lvl2").Subrouter()
	lvl2Router.HandleFunc("", lvl2.IndexHandler)
	lvl2Router.HandleFunc("/internal", lvl2.InternalHandler)
	lvl2Router.HandleFunc("/login", lvl2.LoginHandler).Methods("POST")
	lvl2Router.HandleFunc("/logout", lvl2.LogoutHandler).Methods("POST")



	http.Handle("/", r)
	http.Handle("/static/", http.StripPrefix("/static/" ,http.FileServer(http.Dir("./static"))))
	http.ListenAndServe(":8000", nil)
}
