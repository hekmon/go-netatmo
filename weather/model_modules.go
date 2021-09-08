package weather

import (
	"encoding/json"
	"fmt"
	"net"
	"time"
)

// StationDataBodyDevices struct for StationDataBodyDevices
type StationDataBodyDevices struct {
	ID              string                      `json:"_id"`             // uniq ID of the module (MAC address)
	DateSetup       time.Time                   ``                       // date when the weather station was set up
	LastSetup       time.Time                   ``                       // timestamp of the last installation
	LastStatusStore time.Time                   ``                       // timestamp of the last status update
	LastUpgrade     time.Time                   ``                       // timestamp of the last upgrade
	Type            ModuleType                  `json:"type"`            // type of the device (should alway be ModuleTypeStation const value)
	ModuleName      string                      `json:"module_name"`     // name of the module
	Firmware        int                         `json:"firmware"`        // version of the software
	WifiStatus      WiFiQuality                 `json:"wifi_status"`     // wifi status per Base station. (86=bad, 56=good)
	Reachable       bool                        `json:"reachable"`       // true if the station connected to Netatmo cloud within the last 4 hours
	CO2Calibrating  bool                        `json:"co2_calibrating"` // true if the station is calibrating
	DataType        []ModuleDataType            `json:"data_type"`       // array of data measured by the device (see ModuleDataType const values)
	Place           Place                       `json:"place"`           // informations about where the station is
	ReadOnly        bool                        `json:"read_only"`       // true if the user owns the station, false if he is invited to a station
	HomeId          string                      `json:"home_id"`         // id of the home where the station is placed
	HomeName        string                      `json:"home_name"`       // name of the home where the station is placed
	DashboardData   DashboardDataWeatherstation `json:"dashboard_data"`  // values summary
	Modules         []Module                    `json:"modules"`
	// StationName     string                      `json:"station_name"`      // name of the station - DO NOT USE ANYMORE - use home_name and module_name instead
}

// UnmarshalJSON allows to create a proper payloade on the fly during JSON unmarshaling
func (sdbd *StationDataBodyDevices) UnmarshalJSON(data []byte) (err error) {
	// Add tmp type
	type OriginalUnmarshal StationDataBodyDevices
	tmp := struct {
		DateSetup       int64 `json:"date_setup"`        // date when the weather station was set up
		LastSetup       int64 `json:"last_setup"`        // timestamp of the last installation
		LastStatusStore int64 `json:"last_status_store"` // timestamp of the last status update
		LastUpgrade     int64 `json:"last_upgrade"`      // timestamp of the last upgrade
		*OriginalUnmarshal
	}{
		OriginalUnmarshal: (*OriginalUnmarshal)(sdbd),
	}
	// Unmarshall into the tmp fields
	if err = json.Unmarshal(data, &tmp); err != nil {
		err = fmt.Errorf("failed to unmarshal data to the temporary StationDataBodyDevices struct: %w", err)
		return
	}
	// Convert timestamps
	sdbd.DateSetup = time.Unix(tmp.DateSetup, 0)
	sdbd.LastSetup = time.Unix(tmp.LastSetup, 0)
	sdbd.LastStatusStore = time.Unix(tmp.LastStatusStore, 0)
	sdbd.LastUpgrade = time.Unix(tmp.LastUpgrade, 0)
	return
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

// Module contains information about any additionnal modules
type Module struct {
	ID                   net.HardwareAddr            `json:"-"`               // uniq ID of the module (MAC address)
	Type                 ModuleType                  `json:"type"`            // type of module (see ModuleType const values)
	ModuleName           string                      `json:"module_name"`     // user set name of the module
	DataType             []ModuleDataType            `json:"data_type"`       // array of data measured by the device (see ModuleDataType const values)
	LastSetup            time.Time                   `json:"-"`               // date of the last installation
	Reachable            bool                        `json:"reachable"`       // true if the station connected to Netatmo cloud within the last 4 hours
	Firmware             int64                       `json:"firmware"`        // version of the software
	LastMessage          time.Time                   `json:"-"`               // date of the last measure update
	LastSeen             time.Time                   `json:"-"`               // date of the last status update
	RfStatus             RadioQuality                `json:"rf_status"`       // current radio status per module (see RadioQuality const values)
	BatteryVp            int64                       `json:"battery_vp"`      // current battery status per module (legacy, see BatteryPercent)
	BatteryPercent       int64                       `json:"battery_percent"` // percentage of battery remaining (10=low)
	DashboardDataOutdoor *OutdoorModuleDashboardData `json:"-"`               // values summary if module type is outdoor and is reachable
	DashboardDataWind    *WindModuleDashboardData    `json:"-"`               // values summary if module type is wind and is reachable
	DashboardDataRain    *RainModuleDashboardData    `json:"-"`               // values summary if module type is rain and is reachable
	DashboardDataIndoor  *IndoorModuleDashboardData  `json:"-"`               // values summary if module type is indoor and is reachable
	DashboardDataRaw     json.RawMessage             `json:"-"`               // in case type auto detect has failed, raw dashboard will be kept here (module must still be reachable)
}

// UnmarshalJSON allows to automatically convert data to go types
func (m *Module) UnmarshalJSON(data []byte) (err error) {
	// Add tmp type
	type OriginalUnmarshal Module
	tmp := struct {
		ID               string          `json:"_id"`
		LastSetup        int64           `json:"last_setup"`
		LastMessage      int64           `json:"last_message"`
		LastSeen         int64           `json:"last_seen"`
		DashboardDataRaw json.RawMessage `json:"dashboard_data"`
		*OriginalUnmarshal
	}{
		OriginalUnmarshal: (*OriginalUnmarshal)(m),
	}
	// Unmarshall into the tmp fields
	if err = json.Unmarshal(data, &tmp); err != nil {
		err = fmt.Errorf("failed to unmarshal data to the temporary Module struct: %w", err)
		return
	}
	// Convert
	if m.ID, err = net.ParseMAC(tmp.ID); err != nil {
		err = fmt.Errorf("failed to parse ID as MAC Addr from the temporary Module struct: %w", err)
		return
	}
	m.LastSetup = time.Unix(tmp.LastSetup, 0)
	m.LastMessage = time.Unix(tmp.LastMessage, 0)
	m.LastSeen = time.Unix(tmp.LastSeen, 0)
	// Handle module type to select the right dashboard
	if m.Reachable {
		switch tmp.Type {
		case ModuleTypeOutdoor:
			m.DashboardDataOutdoor = new(OutdoorModuleDashboardData)
			if err = json.Unmarshal(tmp.DashboardDataRaw, m.DashboardDataOutdoor); err != nil {
				err = fmt.Errorf("failed to parse module name '%s' of type '%s' into Outdoor dashboard: %w",
					m.ModuleName, m.Type, err)
				return
			}
		case ModuleTypeAnemometer:
			m.DashboardDataWind = new(WindModuleDashboardData)
			if err = json.Unmarshal(tmp.DashboardDataRaw, m.DashboardDataWind); err != nil {
				err = fmt.Errorf("failed to parse module name '%s' of type '%s' into Wind dashboard: %w",
					m.ModuleName, m.Type, err)
				return
			}
		case ModuleTypeRainGauge:
			m.DashboardDataRain = new(RainModuleDashboardData)
			if err = json.Unmarshal(tmp.DashboardDataRaw, m.DashboardDataRain); err != nil {
				err = fmt.Errorf("failed to parse module name '%s' of type '%s' into Rain dashboard: %w",
					m.ModuleName, m.Type, err)
				return
			}
		case ModuleTypeIndoor:
			m.DashboardDataIndoor = new(IndoorModuleDashboardData)
			if err = json.Unmarshal(tmp.DashboardDataRaw, m.DashboardDataIndoor); err != nil {
				err = fmt.Errorf("failed to parse module name '%s' of type '%s' into Indoor dashboard: %w",
					m.ModuleName, m.Type, err)
				return
			}
		default:
			// keep the raw value for user to work around if necessary
			m.DashboardDataRaw = tmp.DashboardDataRaw
		}
	}
	return
}

// OutdoorModuleDashboardData struct for OutdoorModuleDashboardData
type OutdoorModuleDashboardData struct {
	Time        time.Time `json:"-"`           // date when data was measured
	Temperature float64   `json:"Temperature"` // temperature (in °C)
	Humidity    int64     `json:"Humidity"`    // humidity (in %)
	MinTemp     float64   `json:"min_temp"`    // minimum temperature measured
	MaxTemp     float64   `json:"max_temp"`    // maximum temperature measured
	DateMinTemp time.Time `json:"-"`           // date of minimum temperature measured
	DateMaxTemp time.Time `json:"-"`           // date of maximum temperature measured
	TempTrend   Trend     `json:"temp_trend"`  // trend for the last 12h (up, down, stable: see Trend const values)
}

// UnmarshalJSON allows to automatically convert data to go types
func (omdd *OutdoorModuleDashboardData) UnmarshalJSON(data []byte) (err error) {
	type OriginalUnmarshal OutdoorModuleDashboardData
	tmp := struct {
		TimeUTC     int64 `json:"time_utc"`      // timestamp when data was measured
		DateMinTemp int64 `json:"date_min_temp"` // timestamp of minimum temperature measured
		DateMaxTemp int64 `json:"date_max_temp"` // timestamp of maximum temperature measured
		*OriginalUnmarshal
	}{
		OriginalUnmarshal: (*OriginalUnmarshal)(omdd),
	}
	// Unmarshall into the tmp fields
	if err = json.Unmarshal(data, &tmp); err != nil {
		err = fmt.Errorf("can not unmarshall into Outdoor dashboard tmp struct: %w", err)
		return
	}
	// convert
	omdd.Time = time.Unix(tmp.TimeUTC, 0)
	omdd.DateMinTemp = time.Unix(tmp.DateMinTemp, 0)
	omdd.DateMaxTemp = time.Unix(tmp.DateMaxTemp, 0)
	return
}

// WindModuleDashboardData struct for WindModuleDashboardData
type WindModuleDashboardData struct {
	Time           time.Time `json:"-"`              // date when data was measured
	WindStrength   int64     `json:"WindStrength"`   // wind strength (km/h)
	WindAngle      int64     `json:"WindAngle"`      // wind angle
	GustStrength   int64     `json:"GustStrength"`   // gust strength (km/h)
	GustAngle      int64     `json:"GustAngle"`      // gust angle
	MaxWindStr     int64     `json:"max_wind_str"`   // max wind strength (km/h)
	MaxWindAngle   int64     `json:"max_wind_angle"` // max wind angle
	DateMaxWindStr time.Time `json:"-"`              // max wind date
}

// UnmarshalJSON allows to automatically convert data to go types
func (wmdd *WindModuleDashboardData) UnmarshalJSON(data []byte) (err error) {
	type OriginalUnmarshal WindModuleDashboardData
	tmp := struct {
		TimeUTC        int64 `json:"time_utc"`          // timestamp when data was measured
		DateMaxWindStr int64 `json:"date_max_wind_str"` // max wind date
		*OriginalUnmarshal
	}{
		OriginalUnmarshal: (*OriginalUnmarshal)(wmdd),
	}
	// Unmarshall into the tmp fields
	if err = json.Unmarshal(data, &tmp); err != nil {
		err = fmt.Errorf("can not unmarshall into Wind dashboard tmp struct: %w", err)
		return
	}
	// convert
	wmdd.Time = time.Unix(tmp.TimeUTC, 0)
	wmdd.DateMaxWindStr = time.Unix(tmp.DateMaxWindStr, 0)
	return
}

// RainModuleDashboardData struct for RainModuleDashboardData
type RainModuleDashboardData struct {
	Time      time.Time `json:"-"`           // date when data was measured
	Rain      float64   `json:"Rain"`        // rain in mm
	SumRain24 float64   `json:"sum_rain_24"` // rain measured for past 24h(mm)
	SumRain1  float64   `json:"sum_rain_1"`  // rain measured for the last hour (mm)
}

// UnmarshalJSON allows to automatically convert data to go types
func (rmdd *RainModuleDashboardData) UnmarshalJSON(data []byte) (err error) {
	type OriginalUnmarshal RainModuleDashboardData
	tmp := struct {
		TimeUTC int64 `json:"time_utc"` // timestamp when data was measured
		*OriginalUnmarshal
	}{
		OriginalUnmarshal: (*OriginalUnmarshal)(rmdd),
	}
	// Unmarshall into the tmp fields
	if err = json.Unmarshal(data, &tmp); err != nil {
		err = fmt.Errorf("can not unmarshall into Rain dashboard tmp struct: %w", err)
		return
	}
	// convert
	rmdd.Time = time.Unix(tmp.TimeUTC, 0)
	return
}

// IndoorModuleDashboardData struct for IndoorModuleDashboardData
type IndoorModuleDashboardData struct {
	Time        time.Time `json:"-"`           // date when data was measured
	Temperature float64   `json:"Temperature"` // temperature (in °C)
	CO2         int64     `json:"CO2"`         // CO2 level (in ppm)
	Humidity    int64     `json:"Humidity"`    // humidity (in %)
	MinTemp     float64   `json:"min_temp"`    // maximum temperature measured
	MaxTemp     float64   `json:"max_temp"`    // maximum temperature measured
	DateMinTemp time.Time `json:"-"`           // date of minimum temperature measured
	DateMaxTemp time.Time `json:"-"`           // date of maximum temperature measured
	TempTrend   Trend     `json:"temp_trend"`  // trend for the last 12h (up, down, stable: see Trend const values)
}

// UnmarshalJSON allows to automatically convert data to go types
func (imdd *IndoorModuleDashboardData) UnmarshalJSON(data []byte) (err error) {
	type OriginalUnmarshal IndoorModuleDashboardData
	tmp := struct {
		TimeUTC     int64 `json:"time_utc"`      // timestamp when data was measured
		DateMinTemp int64 `json:"date_min_temp"` // timestamp of minimum temperature measured
		DateMaxTemp int64 `json:"date_max_temp"` // timestamp of maximum temperature measured
		*OriginalUnmarshal
	}{
		OriginalUnmarshal: (*OriginalUnmarshal)(imdd),
	}
	// Unmarshall into the tmp fields
	if err = json.Unmarshal(data, &tmp); err != nil {
		err = fmt.Errorf("can not unmarshall into Indoor dashboard tmp struct: %w", err)
		return
	}
	// convert
	imdd.Time = time.Unix(tmp.TimeUTC, 0)
	imdd.DateMinTemp = time.Unix(tmp.DateMinTemp, 0)
	imdd.DateMaxTemp = time.Unix(tmp.DateMaxTemp, 0)
	return
}
