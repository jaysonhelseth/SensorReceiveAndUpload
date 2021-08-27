package main

import (
	"SensorReceiveAndUpload/io"
	"SensorReceiveAndUpload/models"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func sensorData(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.WriteHeader(http.StatusOK)
	fmt.Fprint(writer, models.CurrentData.Read())
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", sensorData)

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go io.ReadFromSerial()
	log.Fatal(srv.ListenAndServe())
}
