package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/pat"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)

var rd *render.Render

type User struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func getUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	user := User{Name: "tucker", Email: "tucker@naver.com"}

	rd.JSON(w, http.StatusOK, user)
}

func addUserHandler(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		rd.Text(w, http.StatusBadRequest, err.Error())
		return
	}
	user.CreatedAt = time.Now()
	rd.JSON(w, http.StatusOK, user)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	user := User{Name: "tucker", Email: "tucker@naver.com"}
	rd.HTML(w, http.StatusOK, "body", user)
}

func main() {
	rd = render.New(render.Options{
		Directory:  "template",
		Extensions: []string{".html", ".tmpl"},
		Layout:     "hello",
	})
	mux := pat.New()

	mux.Get("/users", getUserInfoHandler)
	mux.Post("/users", addUserHandler)
	mux.Get("/hello", helloHandler)

	n := negroni.Classic()
	n.UseHandler(mux)
	http.ListenAndServe(":3000", n)
}
