package models

import (
	"encoding/json"
	"time"
)

type DataBin struct {
	Temp     float32 `json:"temp"`
	Humidity int32   `json:"humidity"`
	Pool     float32 `json:"pool"`
	Central  string  `json:"cental"`
}

func (dataBin *DataBin) GetTime(serialData []byte) {
	_ = json.Unmarshal(serialData, dataBin)
	dataBin.Central = time.Now().Format("2006-01-02 15:04:05")
}

func (dataBin *DataBin) Stringify() string {
	bytes, _ := json.Marshal(dataBin)
	return string(bytes)
}
