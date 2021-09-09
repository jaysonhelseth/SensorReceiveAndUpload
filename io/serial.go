package io

import (
	"SensorReceiveAndUpload/config"
	"SensorReceiveAndUpload/models"
	"bufio"
	"encoding/json"
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

func read() ([]byte, error) {
	reader := bufio.NewReader(device)
	bytes, err := reader.ReadBytes('\x7d')

	if err != nil {
		return errorTempMsg(), err
	}

	return fixData(bytes), nil
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
		serialData, err := read()

		var msg string
		if err == nil {
			dataBin := models.DataBin{}
			dataBin.GetTime(serialData)
			msg = dataBin.Stringify()
		} else {
			msg = string(serialData)
		}

		log.Println(msg)
		models.CurrentData.Write(msg)
	}
}
