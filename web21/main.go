package main

import (
	"log"
	"net/http"

	"github.com/tuckersGo/goWeb/web21/app"
)

func main() {
	m := app.MakeHandler("./test.db")
	defer m.Close()

	log.Println("Started App")
	err := http.ListenAndServe(":3000", m)
	if err != nil {
		panic(err)
	}
}
