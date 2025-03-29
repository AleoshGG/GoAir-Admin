package entities

type SensorType string

const (
	AirQuality  SensorType = "air_quality"
	Temperature SensorType = "temperature"
	Humidity    SensorType = "humidity"
)

type Sensor struct {
	Id_sensor         string
	Id_place          int
	Sensor_type       SensorType
	Model             string
	Installation_date string
}