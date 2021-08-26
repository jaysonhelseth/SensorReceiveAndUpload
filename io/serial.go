package io

import (
	"SensorReceiveAndUpload/config"
	"encoding/json"
	"github.com/gorilla/websocket"
	"time"
)

func ReadFromSerial() {
	// Todo: Put sensor read loop in here

	var temp float32 = 0.0
	for {
		dict := map[string]float32{
			"temp": temp,
		}
		msg, _ := json.Marshal(dict)

		config.Websocket.WriteMessage(websocket.TextMessage, msg)

		temp += 0.1
		time.Sleep(time.Second * 1)
	}
}
