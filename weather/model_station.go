package weather

// StationDataBodyDevices struct for StationDataBodyDevices
type StationDataBodyDevices struct {
	ID              string                      `json:"_id"`               // mac address of the device
	DateSetup       float32                     `json:"date_setup"`        // date when the weather station was set up
	LastSetup       float32                     `json:"last_setup"`        // timestamp of the last installation
	Type            string                      `json:"type"`              // type of the device
	LastStatusStore float32                     `json:"last_status_store"` // timestamp of the last status update
	ModuleName      string                      `json:"module_name"`       // name of the module
	Firmware        float32                     `json:"firmware"`          // version of the software
	LastUpgrade     float32                     `json:"last_upgrade"`      // timestamp of the last upgrade
	WifiStatus      float32                     `json:"wifi_status"`       // wifi status per Base station. (86=bad, 56=good)
	Reachable       bool                        `json:"reachable"`         // true if the station connected to Netatmo cloud within the last 4 hours
	Co2Calibrating  bool                        `json:"co2_calibrating"`   // true if the station is calibrating
	StationName     string                      `json:"station_name"`      // name of the station - DO NOT USE ANYMORE - use home_name and module_name instead
	DataType        []string                    `json:"data_type"`         // array of data measured by the device (e.g. \"Temperature\",\"Humidity\")
	Place           Place                       `json:"place"`             // informations about where the station is
	ReadOnly        bool                        `json:"read_only"`         // true if the user owns the station, false if he is invited to a station
	HomeId          string                      `json:"home_id"`           // id of the home where the station is placed
	HomeName        string                      `json:"home_name"`         // name of the home where the station is placed
	DashboardData   DashboardDataWeatherstation `json:"dashboard_data"`    // values summary
	//Modules         []OneOfIndoorModuleOutdoorModuleRainModuleWindModule `json:"modules"`
}

// DashboardDataWeatherstation Weather - Weather station, getstationdata
type DashboardDataWeatherstation struct {
	TimeUTC          float32 `json:"time_utc"`         // timestamp when data was measured
	Temperature      float32 `json:"Temperature"`      // temperature (in °C)
	CO2              float32 `json:"CO2"`              // CO2 level (in ppm)
	Humidity         float32 `json:"Humidity"`         // humidity (in %)
	Noise            float32 `json:"Noise"`            // noise level (in dB)
	Pressure         float32 `json:"Pressure"`         // surface pressure in mbar
	AbsolutePressure float32 `json:"AbsolutePressure"` // sea-level pressure in mbar
	MinTemp          float32 `json:"min_temp"`         // minimum temperature measured
	MaxTemp          float32 `json:"max_temp"`         // maximum temperature measured
	DateMinTemp      float32 `json:"date_min_temp"`    // date of minimum temperature measured
	DateMaxTemp      float32 `json:"date_max_temp"`    // date of maximum temperature measured
	TempTrend        string  `json:"temp_trend"`       // trend for the last 12h (up, down, stable)
	PressureTrend    string  `json:"pressure_trend"`   // trend for the last 12h (up, down, stable)
}
