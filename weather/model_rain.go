package weather

// RainModule Weather - IN Module, for rain gauge module, getstationdata
type RainModule struct {
	ID             string                  `json:"_id"`
	Type           string                  `json:"type"`
	ModuleName     string                  `json:"module_name"`
	DataType       []string                `json:"data_type"`       // array of data measured by the device (e.g. \"Temperature\",\"Humidity\")
	LastSetup      float32                 `json:"last_setup"`      // timestamp of the last installation
	Reachable      bool                    `json:"reachable"`       // true if the station connected to Netatmo cloud within the last 4 hours
	DashboardData  RainModuleDashboardData `json:"dashboard_data"`  // values summary
	Firmware       float32                 `json:"firmware"`        // version of the software
	LastMessage    float32                 `json:"last_message"`    // timestamp of the last measure update
	LastSeen       float32                 `json:"last_seen"`       // timestamp of the last status update
	RfStatus       float32                 `json:"rf_status"`       // current radio status per module. (90=low, 60=highest)
	BatteryVp      float32                 `json:"battery_vp"`      // current battery status per module
	BatteryPercent float32                 `json:"battery_percent"` // percentage of battery remaining (10=low)
}

// RainModuleDashboardData struct for RainModuleDashboardData
type RainModuleDashboardData struct {
	TimeUTC   float32 `json:"time_utc"`    // timestamp when data was measured
	Rain      float32 `json:"Rain"`        // rain in mm
	SumRain24 float32 `json:"sum_rain_24"` // rain measured for past 24h(mm)
	SumRain1  float32 `json:"sum_rain_1"`  // rain measured for the last hour (mm)
}

// MeasureMacAddressNAModule3 struct for MeasureMacAddressNAModule3
type MeasureMacAddressNAModule3 struct {
	Rain60min   float32 `json:"rain_60min"`
	Rain24h     float32 `json:"rain_24h"`
	RainLive    float32 `json:"rain_live"`
	RainTimeUTC float32 `json:"rain_timeutc"`
}
