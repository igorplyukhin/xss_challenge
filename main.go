package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"xss_challenge/lvl1"
	"xss_challenge/lvl2"
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


	http.Handle("/", r)
	http.Handle("/static/", http.StripPrefix("/static/" ,http.FileServer(http.Dir("./static"))))
	fmt.Printf("_____________________________________________________________________________________________" +
		"\nListening at %s", address)
	http.ListenAndServe(address, nil)
}
