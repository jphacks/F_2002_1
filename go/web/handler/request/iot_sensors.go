package request

type iotSensorData struct {
	Psoc6DeviceID int     `json:"psoc6_device_id"`
	CultivationID int     `json:"cultivation_id"`
	Humidity      float32 `json:"humidity"`
	Temperature   float32 `json:"temperature"`
	Illuminance   float32 `json:"illuminance"`
	SoilMoisture  float32 `json:"soil_moisture"`
	Pressure      float32 `json:"pressure"`
}

type IotSensorsPost struct {
	UserID        int             `json:"user_id"`
	RaspiDeviceID int             `json:"raspi_device_id"`
	Data          []iotSensorData `json:"data"`
}
