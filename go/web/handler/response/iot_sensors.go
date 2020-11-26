package response

import "time"

type IotSensorData struct {
	Value  float32 `json:"value"`
	Status string  `json:"status"`
}

type IotSensorsPost struct {
	CultivationID int           `json:"cultivation_id"`
	CreatedAt     time.Time     `json:"created_at"`
	Humidity      IotSensorData `json:"humidity"`
	Temperature   IotSensorData `json:"temperature"`
	Illuminance   IotSensorData `json:"illuminance"`
	SoilMoisture  IotSensorData `json:"soil_moisture"`
	Pressure      IotSensorData `json:"pressure"`
}
