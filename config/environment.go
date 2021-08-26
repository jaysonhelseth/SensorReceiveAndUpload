package config

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"os"
)

var (
	Websocket *websocket.Conn
)

func GetEnv(key string) string {
	val, ok := os.LookupEnv(key)

	if !ok {
		log.Fatal(fmt.Sprintf("Missing %s", key))
	}

	return val
}
