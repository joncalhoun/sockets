package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/joncalhoun/webjack"
)

var server *webjack.Server

type Message struct {
	Message string `schema:"message"`
	Channel string `schema:"channel"`
}

func main() {
	server = webjack.NewServer()
	http.Handle("/ws", server.GetHandler())
	http.HandleFunc("/msg", messageHandler)
	http.HandleFunc("/stats", statsHandler)
	http.Handle("/", http.FileServer(http.Dir("public")))
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func messageHandler(w http.ResponseWriter, r *http.Request) {
	key := r.Header.Get("Authorization")
	if key == "fWb7CIPUyTx3jZlE5VX9npgmMuQo" {
		err := r.ParseForm()

		if err != nil {
			// Handle error
		}

		decoder := schema.NewDecoder()
		msg := new(Message)
		fmt.Println(msg)
		err = decoder.Decode(msg, r.PostForm)
		server.SendChannel(msg.Message, msg.Channel)
	}
}

func statsHandler(w http.ResponseWriter, r *http.Request) {
	out, _ := json.Marshal(server.ConnectionsPerChannel())
	w.Write(out)
}
