package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/williamaldevino/websocket/server"
)

func main() {
	fmt.Println("Listening port 3000...")

	hub := server.CreateHub()
	go hub.Run()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		server.ServeWs(hub, w, r)
	})

	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		log.Fatal("Erro", err)
	}

}
