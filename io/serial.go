package io

import (
	"SensorReceiveAndUpload/config"
	"bufio"
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/tarm/serial"
	"log"
	"strings"
)

var (
	deviceConfig serial.Config
	device       *serial.Port
)

func setup() {
	deviceConfig = serial.Config{Name: config.GetEnv("XBEE_RECEIVER"), Baud: 9600}

	var err error
	device, err = serial.OpenPort(&deviceConfig)

	if err != nil {
		log.Fatal("Failed to setup listening device.")
	}
}

func read() []byte {
	reader := bufio.NewReader(device)
	bytes, err := reader.ReadBytes('\x7d')

	if err != nil {
		errorTempMsg()
	}

	return fixData(bytes)
}

func fixData(bytes []byte) []byte {
	str := string(bytes)
	index := strings.IndexRune(str, '{')

	sub := str[index:]
	return []byte(sub)
}

func errorTempMsg() []byte {
	dict := map[string]float32{
		"temp": 0.0,
	}
	msg, _ := json.Marshal(dict)
	return msg
}

func ReadFromSerial() {
	setup()

	for {
		msg := read()

		log.Println(string(msg))
		config.Websocket.WriteMessage(websocket.TextMessage, msg)
	}
}
