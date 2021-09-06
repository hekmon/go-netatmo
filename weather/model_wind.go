package weather

// MeasureMacAddressNAModule2 struct for MeasureMacAddressNAModule2
type MeasureMacAddressNAModule2 struct {
	WindStrength float32 `json:"wind_strengh"` // yes the API has a typo on JSON key
	WindAngle    float32 `json:"wind_angle"`
	GustStrength float32 `json:"gust_strenght"` // yes the API has a typo on JSON key
	GustAngle    float32 `json:"gust_angle"`
	WindTimeUTC  float32 `json:"wind_timeutc"`
}
