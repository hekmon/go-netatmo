package weather

// MeasureMacAddressNAModule3 struct for MeasureMacAddressNAModule3
type MeasureMacAddressNAModule3 struct {
	Rain60min   float32 `json:"rain_60min"`
	Rain24h     float32 `json:"rain_24h"`
	RainLive    float32 `json:"rain_live"`
	RainTimeUTC float32 `json:"rain_timeutc"`
}
