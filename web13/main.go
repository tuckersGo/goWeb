package main

import (
	"log"
	"net/http"

	"github.com/gorilla/pat"
	"github.com/urfave/negroni"
)

func postMessageHandler(w http.ResponseWriter, r *http.Request) {
	msg := r.FormValue("msg")
	name := r.FormValue("name")
	log.Println("postMessageHandler ", msg, name)
}

func main() {
	mux := pat.New()
	mux.Post("/messages", postMessageHandler)

	n := negroni.Classic()
	n.UseHandler(mux)

	http.ListenAndServe(":3000", n)
}
