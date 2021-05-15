package weather

// WindModule Weather - IN Module, for wind gauge module, getstationdata
type WindModule struct {
	ID             string                  `json:"_id"`
	Type           string                  `json:"type"`
	ModuleName     string                  `json:"module_name"`
	DataType       []string                `json:"data_type"`       // Array of data measured by the device (e.g. \"Temperature\",\"Humidity\")
	LastSetup      float32                 `json:"last_setup"`      // timestamp of the last installation
	BatteryPercent float32                 `json:"battery_percent"` // Percentage of battery remaining (10=low)
	Reachable      bool                    `json:"reachable"`       // true if the station connected to Netatmo cloud within the last 4 hours
	Firmware       float32                 `json:"firmware"`        // version of the software
	LastMessage    float32                 `json:"last_message"`    // timestamp of the last measure update
	LastSeen       float32                 `json:"last_seen"`       // timestamp of the last status update
	RfStatus       float32                 `json:"rf_status"`       // Current radio status per module. (90=low, 60=highest)
	BatteryVp      float32                 `json:"battery_vp"`      // current battery status per module
	DashboardData  WindModuleDashboardData `json:"dashboard_data"`  // values summary
}

// WindModuleDashboardData struct for WindModuleDashboardData
type WindModuleDashboardData struct {
	TimeUTC        float32 `json:"time_utc"`     // timestamp when data was measured
	WindStrength   float32 `json:"WindStrength"` // wind strenght (km/h)
	WindAngle      float32 `json:"WindAngle"`    // wind angle
	GustStrength   float32 `json:"GustStrength"` // gust strengh (km/h)
	GustAngle      float32 `json:"GustAngle"`    // gust angle
	MaxWindStr     float32 `json:"max_wind_str"`
	MaxWindAngle   float32 `json:"max_wind_angle"`
	DateMaxWindStr float32 `json:"date_max_wind_str"`
}

// MeasureMacAddressNAModule2 struct for MeasureMacAddressNAModule2
type MeasureMacAddressNAModule2 struct {
	WindStrengh  float32 `json:"wind_strengh"`
	WindAngle    float32 `json:"wind_angle"`
	GustStrenght float32 `json:"gust_strenght"`
	GustAngle    float32 `json:"gust_angle"`
	WindTimeUTC  float32 `json:"wind_timeutc"`
}
