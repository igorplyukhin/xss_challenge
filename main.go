package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"xss_challenge/lvl1"
	"xss_challenge/lvl2"
	"xss_challenge/lvl3"
	"xss_challenge/lvl4"
	"xss_challenge/lvl5"
	"xss_challenge/lvl6"
	"xss_challenge/root"
)


func main() {
	var address = "127.0.0.1:8000"
	var r = mux.NewRouter()

	var rootRouter = r.PathPrefix("").Subrouter()
	rootRouter.HandleFunc("/", root.IndexHandler)

	var lvl1Router = r.PathPrefix("/lvl1").Subrouter()
	lvl1Router.HandleFunc("", lvl1.IndexHandler)

	var lvl2Router = r.PathPrefix("/lvl2").Subrouter()
	lvl2Router.HandleFunc("", lvl2.IndexHandler)

	var lvl3Router = r.PathPrefix("/lvl3").Subrouter()
	lvl3Router.HandleFunc("", lvl3.IndexHandler)

	var lvl4Router = r.PathPrefix("/lvl4").Subrouter()
	lvl4Router.HandleFunc("", lvl4.IndexHandler)

	var lvl5Router = r.PathPrefix("/lvl5").Subrouter()
	lvl5Router.HandleFunc("", lvl5.IndexHandler)

	var lvl6Router = r.PathPrefix("/lvl6").Subrouter()
	lvl6Router.HandleFunc("", lvl6.IndexHandler)



	http.Handle("/", r)
	http.Handle("/static/", http.StripPrefix("/static/" ,http.FileServer(http.Dir("./static"))))
	fmt.Printf("_____________________________________________________________________________________________" +
		"\nListening at %s", address)
	http.ListenAndServe(address, nil)
}
