package config

import (
	"fmt"
	"log"
	"os"
)

func GetEnv(key string) string {
	val, ok := os.LookupEnv(key)

	if !ok {
		log.Fatal(fmt.Sprintf("Missing %s", key))
	}

	return val
}
