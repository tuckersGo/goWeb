package main

import (
	"log"
	"net/http"

	"github.com/gorilla/pat"
	"github.com/gorilla/websocket"
	"github.com/urfave/negroni"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Message struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

func wshandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	for {
		m := &Message{}
		err := conn.ReadJSON(m)
		if err != nil {
			log.Println(err)
			return
		}

		err = conn.WriteJSON(m)
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func main() {
	mux := pat.New()
	mux.Get("/ws", wshandler)

	n := negroni.Classic()
	n.UseHandler(mux)

	http.ListenAndServe(":3000", n)
}
