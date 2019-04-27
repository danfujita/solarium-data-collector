package app

import (
	"encoding/json"
	"github.com/influxdata/influxdb1-client/v2"
	"solarium-golang/internal/configs"
	"time"
)

type TelementryData struct {
	RealTimeClock time.Time `json:"real_time_clock"`
	Temperature   float32   `json:"temperature"`
	Pressure      float32   `json:"pressure"`
	Humidity      float32   `json:"humidity"`
	Accel         []float32 `json:"accel"`
	Gyro          []float32 `json:"gyro"`
	Mag           []float32 `json:"mag"`
	Voltages      []float32 `json:"voltages"`
}

func UploadTelemetryData(telementryData map[string]interface{}, userId string, deviceId string) error {
	myConfig := configs.Config()
	client.NewHTTPClient(client.HTTPConfig{ //TODO Handle Error
		Addr: myConfig.InfluxAddress,
	})
	tags := map[string]string{
		"userId":   userId,
		"deviceId": deviceId,
	}
	bp, _ := client.NewBatchPoints(client.BatchPointsConfig{ //TODO Handle Error
		Database:  myConfig.InfluxDatabase,
		Precision: "us",
	})
	pt, err := client.NewPoint(myConfig.InfluxTableName, tags, telementryData, time.Now()) //TODO Handle Error
	if err != nil {
		return err
	}
	bp.AddPoint(pt)
	return nil
}
func TelemetryData(deviceInfo map[string]string, data []byte) error {
	var telementryData TelementryData
	err := json.Unmarshal(data, &telementryData)
	if err != nil {
		return err
	}
	var telementryDataMap map[string]interface{}
	telementryDataJson, _ := json.Marshal(telementryData)
	json.Unmarshal(telementryDataJson, &telementryDataMap) //TODO fix line 43-49.
	err = UploadTelemetryData(telementryDataMap, deviceInfo["user_id"], deviceInfo["device_id"])
	if err != nil {
		return err
	}
	return nil
}
