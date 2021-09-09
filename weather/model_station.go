package weather

import (
	"encoding/json"
	"fmt"
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
	DashboardData   DashboardDataWeatherStation `json:"dashboard_data"`  // values summary
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

// WiFiQuality represents the WiFi strength signal
type WiFiQuality int

const (
	// WiFiQualityBad represents a bad level for WiFi reception
	WiFiQualityBad WiFiQuality = 86
	// WiFiQualityAverage represents an average level for WiFi reception
	WiFiQualityAverage WiFiQuality = 71
	// WiFiQUalityGood represents a good level for WiFi reception
	WiFiQUalityGood WiFiQuality = 56
)

// String implements the https://golang.org/pkg/fmt/#Stringer interface
func (wq WiFiQuality) String() string {
	switch {
	case wq <= WiFiQUalityGood:
		return "good"
	case wq <= WiFiQualityAverage:
		return "average"
	case wq <= WiFiQualityBad:
		return "bad"
	default:
		return "very bad"
	}
}

// GoString implements the https://golang.org/pkg/fmt/#GoStringer interface
func (wq WiFiQuality) GoString() string {
	return fmt.Sprintf("%s (%d)", wq, wq)
}

// Place struct for Place
type Place struct {
	Timezone *time.Location ``                // Timezone
	Country  string         `json:"country"`  // Country
	Altitude float64        `json:"altitude"` // Altitude
	Location []float64      `json:"location"` // Lat, Long
}

// UnmarshalJSON allows to automatically convert data to go types
func (p *Place) UnmarshalJSON(data []byte) (err error) {
	type OriginalUnmarshal Place
	tmp := struct {
		Timezone string `json:"timezone"` // Timezone
		*OriginalUnmarshal
	}{
		OriginalUnmarshal: (*OriginalUnmarshal)(p),
	}
	// Unmarshall into the tmp fields
	if err = json.Unmarshal(data, &tmp); err != nil {
		err = fmt.Errorf("can not unmarshall into Indoor dashboard tmp struct: %w", err)
		return
	}
	// Convert
	if p.Timezone, err = time.LoadLocation(tmp.Timezone); err != nil {
		return fmt.Errorf("can not parse the Timezone: %w", err)
	}
	return
}

// DashboardDataWeatherstation Weather - Weather station, getstationdata
type DashboardDataWeatherStation struct {
	Time             time.Time ``                        // time when data was measured
	Temperature      float64   `json:"Temperature"`      // temperature (in °C)
	CO2              int       `json:"CO2"`              // CO2 level (in ppm)
	Humidity         int       `json:"Humidity"`         // humidity (in %)
	Noise            int       `json:"Noise"`            // noise level (in dB)
	Pressure         float64   `json:"Pressure"`         // surface pressure in mbar
	AbsolutePressure float64   `json:"AbsolutePressure"` // sea-level pressure in mbar
	TempMin          float64   `json:"min_temp"`         // minimum temperature measured
	TempMinDate      time.Time ``                        // date of minimum temperature measured
	TempMax          float64   `json:"max_temp"`         // maximum temperature measured
	TempMaxDate      time.Time ``                        // date of maximum temperature measured
	TempTrend        Trend     `json:"temp_trend"`       // trend for the last 12h (up, down, stable)
	PressureTrend    Trend     `json:"pressure_trend"`   // trend for the last 12h (up, down, stable)
}

// UnmarshalJSON allows to create a proper payloade on the fly during JSON unmarshaling
func (ddws *DashboardDataWeatherStation) UnmarshalJSON(data []byte) (err error) {
	// Add tmp type
	type OriginalUnmarshal DashboardDataWeatherStation
	tmp := struct {
		TimeUTC     int64   `json:"time_utc"`      // timestamp when data was measured
		DateMinTemp float32 `json:"date_min_temp"` // date of minimum temperature measured
		DateMaxTemp float32 `json:"date_max_temp"` // date of maximum temperature measured
		*OriginalUnmarshal
	}{
		OriginalUnmarshal: (*OriginalUnmarshal)(ddws),
	}
	// Unmarshall into the tmp fields
	if err = json.Unmarshal(data, &tmp); err != nil {
		err = fmt.Errorf("failed to unmarshal data to the temporary StationDataBodyDevices struct: %w", err)
		return
	}
	// Convert timestamps
	ddws.Time = time.Unix(tmp.TimeUTC, 0)
	ddws.TempMinDate = time.Unix(int64(tmp.DateMinTemp), 0)
	ddws.TempMaxDate = time.Unix(int64(tmp.DateMaxTemp), 0)
	return
}
