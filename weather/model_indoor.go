package weather

// IndoorModule Weather - Weather module indoor - getstationsdata
type IndoorModule struct {
	ID             string                    `json:"_id"`
	Type           string                    `json:"type"`
	ModuleName     string                    `json:"module_name"`
	DataType       []string                  `json:"data_type"`       // array of data measured by the device (e.g. \"Temperature\",\"Humidity\")
	LastSetup      float32                   `json:"last_setup"`      // timestamp of the last installation
	Reachable      bool                      `json:"reachable"`       // true if the station connected to Netatmo cloud within the last 4 hours
	DashboardData  IndoorModuleDashboardData `json:"dashboard_data"`  // values summary
	Firmware       float32                   `json:"firmware"`        // version of the software
	LastMessage    float32                   `json:"last_message"`    // timestamp of the last measure update
	LastSeen       float32                   `json:"last_seen"`       // timestamp of the last status update
	RfStatus       float32                   `json:"rf_status"`       // current radio status per module. (90=low, 60=highest)
	BatteryVp      float32                   `json:"battery_vp"`      // current battery status per module
	BatteryPercent float32                   `json:"battery_percent"` // percentage of battery remaining (10=low)
}

// IndoorModuleDashboardData struct for IndoorModuleDashboardData
type IndoorModuleDashboardData struct {
	TimeUTC          float32 `json:"time_utc"`         // timestamp when data was measured
	Temperature      float32 `json:"Temperature"`      // temperature (in °C)
	CO2              float32 `json:"CO2"`              // CO2 level (in ppm)
	Humidity         float32 `json:"Humidity"`         // humidity (in %)
	Pressure         float32 `json:"Pressure"`         // surface pressure in mbar
	AbsolutePressure float32 `json:"AbsolutePressure"` // sea-level pressure in mbar
	MinTemp          float32 `json:"min_temp"`         // maximum temperature measured
	MaxTemp          float32 `json:"max_temp"`         // maximum temperature measured
	DateMinTemp      float32 `json:"date_min_temp"`    // date of minimum temperature measured
	DateMaxTemp      float32 `json:"date_max_temp"`    // date of maximum temperature measured
	TempTrend        string  `json:"temp_trend"`       // trend for the last 12h (up, down, stable)
}
