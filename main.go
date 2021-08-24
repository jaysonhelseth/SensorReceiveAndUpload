package main

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

// Maybe move this to a state file somewhere else.
// That way other files can access it.
var (
	ws *websocket.Conn
)

func readFromSerial() {
	// Todo: Put sensor read loop in here
	// Maybe move this code to a new file....
	ws.WriteMessage(websocket.TextMessage, []byte("Sensor connected!"))
}

func serveWs(w http.ResponseWriter, r *http.Request) {
	// Todo: need cors accept code in here

	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	var err error

	ws, err = upgrader.Upgrade(w, r, nil)
	if err != nil {
		if _, ok := err.(websocket.HandshakeError); !ok {
			log.Println(err)
		}
		return
	}

	ws.WriteMessage(websocket.TextMessage, []byte("I'm connected!"))
	readFromSerial()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", serveWs)

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
