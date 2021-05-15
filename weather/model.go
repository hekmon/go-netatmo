package weather

// BodyMeasureGet struct for BodyMeasureGet
type BodyMeasureGet struct {
	Body []MeasureGet `json:"body,omitempty"`
}

// MeasureGet struct for MeasureGet
type MeasureGet struct {
	BegTime  float32   `json:"beg_time,omitempty"`
	StepTime float32   `json:"step_time,omitempty"`
	Value    []float32 `json:"value,omitempty"`
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

// IndoorModule Weather - Weather module indoor - getstationsdata
type IndoorModule struct {
	Id             string                    `json:"_id"`
	Type           string                    `json:"type"`
	ModuleName     string                    `json:"module_name"`
	DataType       []string                  `json:"data_type"`  // Array of data measured by the device (e.g. \"Temperature\",\"Humidity\")
	LastSetup      float32                   `json:"last_setup"` // timestamp of the last installation
	Reachable      bool                      `json:"reachable"`  // true if the station connected to Netatmo cloud within the last 4 hours
	DashboardData  IndoorModuleDashboardData `json:"dashboard_data"`
	Firmware       float32                   `json:"firmware"`        // version of the software
	LastMessage    float32                   `json:"last_message"`    // timestamp of the last measure update
	LastSeen       float32                   `json:"last_seen"`       // timestamp of the last status update
	RfStatus       float32                   `json:"rf_status"`       // Current radio status per module. (90=low, 60=highest)
	BatteryVp      float32                   `json:"battery_vp"`      // current battery status per module
	BatteryPercent float32                   `json:"battery_percent"` // Percentage of battery remaining (10=low)
}

// InvalidServerResponse struct for InvalidServerResponse
type InvalidServerResponse struct {
	Error InvalidServerResponseError `json:"error"`
}

// InvalidServerResponseError struct for InvalidServerResponseError
type InvalidServerResponseError struct {
	Code    float32 `json:"code"`
	Message string  `json:"message"`
}

// Measure struct for Measure
type Measure struct {
	MacAddressNAMain    *MeasureMacAddressNAMain    `json:"mac_address_NAMain,omitempty"`
	MacAddressNAModule1 *MeasureMacAddressNAModule1 `json:"mac_address_NAModule1,omitempty"`
	MacAddressNAModule3 *MeasureMacAddressNAModule3 `json:"mac_address_NAModule3,omitempty"`
	MacAddressNAModule2 *MeasureMacAddressNAModule2 `json:"mac_address_NAModule2,omitempty"`
}

// MeasureMacAddressNAMain struct for MeasureMacAddressNAMain
type MeasureMacAddressNAMain struct {
	Res  *MeasureMacAddressNAMainRes `json:"res,omitempty"`
	Type *string                     `json:"type,omitempty"`
}

// MeasureMacAddressNAMainRes struct for MeasureMacAddressNAMainRes
type MeasureMacAddressNAMainRes struct {
	TimeStamp *[]float32 `json:"time_stamp,omitempty"`
}

// MeasureMacAddressNAModule1 struct for MeasureMacAddressNAModule1
type MeasureMacAddressNAModule1 struct {
	Res  *MeasureMacAddressNAModule1Res `json:"res,omitempty"`
	Type *[]string                      `json:"type,omitempty"`
}

// MeasureMacAddressNAModule1Res struct for MeasureMacAddressNAModule1Res
type MeasureMacAddressNAModule1Res struct {
	TimeStamp *[]float32 `json:"time_stamp,omitempty"`
}

// MeasureMacAddressNAModule2 struct for MeasureMacAddressNAModule2
type MeasureMacAddressNAModule2 struct {
	WindStrengh  *float32 `json:"wind_strengh,omitempty"`
	WindAngle    *float32 `json:"wind_angle,omitempty"`
	GustStrenght *float32 `json:"gust_strenght,omitempty"`
	GustAngle    *float32 `json:"gust_angle,omitempty"`
	WindTimeutc  *float32 `json:"wind_timeutc,omitempty"`
}

// MeasureMacAddressNAModule3 struct for MeasureMacAddressNAModule3
type MeasureMacAddressNAModule3 struct {
	Rain60min   *float32 `json:"rain_60min,omitempty"`
	Rain24h     *float32 `json:"rain_24h,omitempty"`
	RainLive    *float32 `json:"rain_live,omitempty"`
	RainTimeutc *float32 `json:"rain_timeutc,omitempty"`
}

// OutdoorModule Weather - Weather module outdoor - getstationsdata
type OutdoorModule struct {
	Id             *string                     `json:"_id,omitempty"`
	Type           *string                     `json:"type,omitempty"`
	ModuleName     *string                     `json:"module_name,omitempty"`
	DataType       *[]string                   `json:"data_type,omitempty"`  // Array of data measured by the device (e.g. \"Temperature\",\"Humidity\")
	LastSetup      *float32                    `json:"last_setup,omitempty"` // timestamp of the last installation
	Reachable      *bool                       `json:"reachable,omitempty"`  // true if the station connected to Netatmo cloud within the last 4 hours
	DashboardData  *OutdoorModuleDashboardData `json:"dashboard_data,omitempty"`
	Firmware       *float32                    `json:"firmware,omitempty"`        // version of the software
	LastMessage    *float32                    `json:"last_message,omitempty"`    // timestamp of the last measure update
	LastSeen       *float32                    `json:"last_seen,omitempty"`       // timestamp of the last status update
	RfStatus       *float32                    `json:"rf_status,omitempty"`       // Current radio status per module. (90=low, 60=highest)
	BatteryVp      *float32                    `json:"battery_vp,omitempty"`      // current battery status per module
	BatteryPercent *float32                    `json:"battery_percent,omitempty"` // Percentage of battery remaining (10=low)
}

// OutdoorModuleDashboardData struct for OutdoorModuleDashboardData
type OutdoorModuleDashboardData struct {
	TimeUtc     *float32 `json:"time_utc,omitempty"`      // timestamp when data was measured
	Temperature *float32 `json:"Temperature,omitempty"`   // temperature (in °C)
	Humidity    *float32 `json:"Humidity,omitempty"`      // humidity (in %)
	MinTemp     *float32 `json:"min_temp,omitempty"`      // minimum temperature measured
	MaxTemp     *float32 `json:"max_temp,omitempty"`      // maximum temperature measured
	DateMinTemp *float32 `json:"date_min_temp,omitempty"` // date of minimum temperature measured
	DateMaxTemp *float32 `json:"date_max_temp,omitempty"` // date of maximum temperature measured
	TempTrend   *string  `json:"temp_trend,omitempty"`    // trend for the last 12h (up, down, stable)
}

// PublicData struct for PublicData
type PublicData struct {
	Body *[]PublicDataBody `json:"body,omitempty"`
}

// PublicDataBody struct for PublicDataBody
type PublicDataBody struct {
	Id          *string                   `json:"_id,omitempty"`
	Place       *Place                    `json:"place,omitempty"`
	Mark        *float32                  `json:"mark,omitempty"`
	Measures    *Measure                  `json:"measures,omitempty"`
	Modules     *[]string                 `json:"modules,omitempty"`
	ModuleTypes *[]map[string]interface{} `json:"module_types,omitempty"`
}

// Place struct for Place
type Place struct {
	Timezone *string        `json:"timezone,omitempty"` // Timezone
	Country  *string        `json:"country,omitempty"`  // Country
	Altitude *float32       `json:"altitude,omitempty"` // Altitude
	Location *[]interface{} `json:"location,omitempty"`
}

// RainModule Weather - IN Module, for rain gauge module, getstationdata
type RainModule struct {
	Id             *string                  `json:"_id,omitempty"`
	Type           *string                  `json:"type,omitempty"`
	ModuleName     *string                  `json:"module_name,omitempty"`
	DataType       *[]string                `json:"data_type,omitempty"`  // Array of data measured by the device (e.g. \"Temperature\",\"Humidity\")
	LastSetup      *float32                 `json:"last_setup,omitempty"` // timestamp of the last installation
	Reachable      *bool                    `json:"reachable,omitempty"`  // true if the station connected to Netatmo cloud within the last 4 hours
	DashboardData  *RainModuleDashboardData `json:"dashboard_data,omitempty"`
	Firmware       *float32                 `json:"firmware,omitempty"`        // version of the software
	LastMessage    *float32                 `json:"last_message,omitempty"`    // timestamp of the last measure update
	LastSeen       *float32                 `json:"last_seen,omitempty"`       // timestamp of the last status update
	RfStatus       *float32                 `json:"rf_status,omitempty"`       // Current radio status per module. (90=low, 60=highest)
	BatteryVp      *float32                 `json:"battery_vp,omitempty"`      // current battery status per module
	BatteryPercent *float32                 `json:"battery_percent,omitempty"` // Percentage of battery remaining (10=low)
}

// RainModuleDashboardData struct for RainModuleDashboardData
type RainModuleDashboardData struct {
	TimeUtc   *float32 `json:"time_utc,omitempty"`    // timestamp when data was measured
	Rain      *float32 `json:"Rain,omitempty"`        // rain in mm
	SumRain24 *float32 `json:"sum_rain_24,omitempty"` // rain measured for past 24h(mm)
	SumRain1  *float32 `json:"sum_rain_1,omitempty"`  // rain measured for the last hour (mm)
}

// ServerResponse struct for ServerResponse
type ServerResponse struct {
	Status     *string `json:"status,omitempty"`
	TimeExec   *string `json:"time_exec,omitempty"`
	TimeServer *string `json:"time_server,omitempty"`
}

// StationData struct for StationData
type StationData struct {
	Body *StationDataBody `json:"body,omitempty"`
}

// StationDataBody struct for StationDataBody
type StationDataBody struct {
	Devices *[]StationDataBodyDevices `json:"devices,omitempty"`
	User    *UserWeather              `json:"user,omitempty"`
}

// StationDataBodyDevices struct for StationDataBodyDevices
type StationDataBodyDevices struct {
	Id              *string                                               `json:"_id,omitempty"`               // mac address of the device
	DateSetup       *float32                                              `json:"date_setup,omitempty"`        // date when the weather station was set up
	LastSetup       *float32                                              `json:"last_setup,omitempty"`        // timestamp of the last installation
	Type            *string                                               `json:"type,omitempty"`              // type of the device
	LastStatusStore *float32                                              `json:"last_status_store,omitempty"` // timestamp of the last status update
	ModuleName      *string                                               `json:"module_name,omitempty"`       // name of the module
	Firmware        *float32                                              `json:"firmware,omitempty"`          // version of the software
	LastUpgrade     *float32                                              `json:"last_upgrade,omitempty"`      // timestamp of the last upgrade
	WifiStatus      *float32                                              `json:"wifi_status,omitempty"`       // wifi status per Base station. (86=bad, 56=good)
	Reachable       *bool                                                 `json:"reachable,omitempty"`         // true if the station connected to Netatmo cloud within the last 4 hours
	Co2Calibrating  *bool                                                 `json:"co2_calibrating,omitempty"`   // true if the station is calibrating
	StationName     *string                                               `json:"station_name,omitempty"`      // name of the station - DO NOT USE ANYMORE - use home_name and module_name instead
	DataType        *[]string                                             `json:"data_type,omitempty"`         // array of data measured by the device (e.g. \"Temperature\",\"Humidity\")
	Place           *Place                                                `json:"place,omitempty"`
	ReadOnly        *bool                                                 `json:"read_only,omitempty"` // true if the user owns the station, false if he is invited to a station
	HomeId          *string                                               `json:"home_id,omitempty"`   // id of the home where the station is placed
	HomeName        *string                                               `json:"home_name,omitempty"` // name of the home where the station is placed
	DashboardData   *DashboardDataWeatherstation                          `json:"dashboard_data,omitempty"`
	Modules         *[]OneOfIndoorModuleOutdoorModuleRainModuleWindModule `json:"modules,omitempty"`
}

// UserWeather struct for UserWeather
type UserWeather struct {
	Mail           *string                    `json:"mail,omitempty"`
	Administrative *UserWeatherAdministrative `json:"administrative,omitempty"`
}

// UserWeatherAdministrative struct for UserWeatherAdministrative
type UserWeatherAdministrative struct {
	RegLocale    *string  `json:"reg_locale,omitempty"`     // user regional preferences (used for displaying date)
	Lang         *string  `json:"lang,omitempty"`           // user locale
	Country      *string  `json:"country,omitempty"`        // country
	Unit         *float32 `json:"unit,omitempty"`           // 0 -> metric system, 1 -> imperial system
	Windunit     *float32 `json:"windunit,omitempty"`       // 0 -> kph, 1 -> mph, 2 -> ms, 3 -> beaufort, 4 -> knot
	Pressureunit *float32 `json:"pressureunit,omitempty"`   // 0 -> mbar, 1 -> inHg, 2 -> mmHg
	FeelLikeAlgo *float32 `json:"feel_like_algo,omitempty"` // algorithm used to compute feel like temperature, 0 -> humidex, 1 -> heat-index
}

// WindModule Weather - IN Module, for wind gauge module, getstationdata
type WindModule struct {
	Id             *string                  `json:"_id,omitempty"`
	Type           *string                  `json:"type,omitempty"`
	ModuleName     *string                  `json:"module_name,omitempty"`
	DataType       *[]string                `json:"data_type,omitempty"`       // Array of data measured by the device (e.g. \"Temperature\",\"Humidity\")
	LastSetup      *float32                 `json:"last_setup,omitempty"`      // timestamp of the last installation
	BatteryPercent *float32                 `json:"battery_percent,omitempty"` // Percentage of battery remaining (10=low)
	Reachable      *bool                    `json:"reachable,omitempty"`       // true if the station connected to Netatmo cloud within the last 4 hours
	Firmware       *float32                 `json:"firmware,omitempty"`        // version of the software
	LastMessage    *float32                 `json:"last_message,omitempty"`    // timestamp of the last measure update
	LastSeen       *float32                 `json:"last_seen,omitempty"`       // timestamp of the last status update
	RfStatus       *float32                 `json:"rf_status,omitempty"`       // Current radio status per module. (90=low, 60=highest)
	BatteryVp      *float32                 `json:"battery_vp,omitempty"`      // current battery status per module
	DashboardData  *WindModuleDashboardData `json:"dashboard_data,omitempty"`
}

// WindModuleDashboardData struct for WindModuleDashboardData
type WindModuleDashboardData struct {
	TimeUtc        *float32 `json:"time_utc,omitempty"`     // timestamp when data was measured
	WindStrength   *float32 `json:"WindStrength,omitempty"` // wind strenght (km/h)
	WindAngle      *float32 `json:"WindAngle,omitempty"`    // wind angle
	GustStrength   *float32 `json:"GustStrength,omitempty"` // gust strengh (km/h)
	GustAngle      *float32 `json:"GustAngle,omitempty"`    // gust angle
	MaxWindStr     *float32 `json:"max_wind_str,omitempty"`
	MaxWindAngle   *float32 `json:"max_wind_angle,omitempty"`
	DateMaxWindStr *float32 `json:"date_max_wind_str,omitempty"`
}
