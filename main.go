package main

import (
	"encoding/json"
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
	//ws.WriteMessage(websocket.TextMessage, []byte("Sensor connected!"))

	var temp float32 = 0.0
	for {
		dict := map[string]float32{
			"temp": temp,
		}
		msg, _ := json.Marshal(dict)
		ws.WriteMessage(websocket.TextMessage, msg)

		temp += 0.1
		time.Sleep(time.Second * 1)
	}
}

func serveWs(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	var err error

	ws, err = upgrader.Upgrade(w, r, nil)
	if err != nil {
		if _, ok := err.(websocket.HandshakeError); !ok {
			log.Println(err)
		}
		return
	}

	// ws.WriteMessage(websocket.TextMessage, []byte("I'm connected!"))
	log.Println("I'm connected.")
	readFromSerial()
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", serveWs)

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}