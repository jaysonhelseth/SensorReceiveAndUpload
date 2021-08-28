package io

import (
	"SensorReceiveAndUpload/config"
	"SensorReceiveAndUpload/models"
	"context"
	"github.com/jackc/pgx/v4"
	"log"
	"time"
)

func store() {
	conn, connectErr := pgx.Connect(context.Background(), config.GetEnv("DB_URL"))
	if connectErr != nil {
		log.Printf("Unable to connect to database: %v\n", connectErr)
	}
	defer conn.Close(context.Background())

	data := models.CurrentData.Read()
	cmdTag, execErr := conn.Exec(context.Background(), "insert into sensors (info) values ($1)", data)
	if execErr != nil {
		log.Printf("Execution error: %v\n", execErr)
	}

	log.Printf("%v\n", cmdTag)
}

func StoreData() {
	for {
		store()
		time.Sleep(time.Minute * 3)
	}
}
