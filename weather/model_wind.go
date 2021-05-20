package weather

// MeasureMacAddressNAModule2 struct for MeasureMacAddressNAModule2
type MeasureMacAddressNAModule2 struct {
	WindStrengh  float32 `json:"wind_strengh"`
	WindAngle    float32 `json:"wind_angle"`
	GustStrenght float32 `json:"gust_strenght"`
	GustAngle    float32 `json:"gust_angle"`
	WindTimeUTC  float32 `json:"wind_timeutc"`
}
