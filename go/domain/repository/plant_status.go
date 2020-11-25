package repository

// PlantStatus はセンサーデータに対する植物の状態を参照するリポジトリです。
type PlantStatus interface {
	FindHumidityStatusBySensorData(plantID int, x float32) (string error)
	FindIlluminanceStatusBySensorData(plantID int, x float32) (string error)
	FindPressureStatusBySensorData(plantID int, x float32) (string error)
	FindSoilMoistureStatusBySensorData(plantID int, x float32) (string error)
	FindTemperatureStatusBySensorData(plantID int, x float32) (string error)
}
