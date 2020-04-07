package main

import (
	"net/http"

	"github.com/tuckersGo/goWeb/web4/myapp"
)

func main() {

	http.ListenAndServe(":3000", myapp.NewHttpHandler())
}
