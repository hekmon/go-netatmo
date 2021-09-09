package weather

import (
	"encoding/json"
	"fmt"
	"time"
)

// WindMeasures holds measures for the anemometer module
type WindMeasures struct {
	Time         time.Time ``                    // not in this form on the orignal payload
	WindStrength int       `json:"wind_strengh"` // yes the API has a typo on JSON key
	WindAngle    int       `json:"wind_angle"`
	GustStrength int       `json:"gust_strenght"` // yes the API has a typo on JSON key
	GustAngle    int       `json:"gust_angle"`
}

// UnmarshalJSON allows to create a proper payloade on the fly during JSON unmarshaling
func (wm *WindMeasures) UnmarshalJSON(data []byte) (err error) {
	// Add tmp type
	type OriginalUnmarshal WindMeasures
	tmp := struct {
		WindTimestamp int `json:"wind_timeutc"`
		*OriginalUnmarshal
	}{
		OriginalUnmarshal: (*OriginalUnmarshal)(wm),
	}
	// Unmarshall into the tmp fields
	if err = json.Unmarshal(data, &tmp); err != nil {
		err = fmt.Errorf("failed to unmarshal data to the temporary WindMeasures struct: %w", err)
		return
	}
	// convert
	wm.Time = time.Unix(int64(tmp.WindTimestamp), 0)
	return
}

// WindModuleDashboardData struct for WindModuleDashboardData
type WindModuleDashboardData struct {
	Time           time.Time ``                      // date when data was measured
	WindStrength   int       `json:"WindStrength"`   // wind strength (km/h)
	WindAngle      int       `json:"WindAngle"`      // wind angle
	GustStrength   int       `json:"GustStrength"`   // gust strength (km/h)
	GustAngle      int       `json:"GustAngle"`      // gust angle
	MaxWindStr     int       `json:"max_wind_str"`   // max wind strength (km/h)
	MaxWindAngle   int       `json:"max_wind_angle"` // max wind angle
	DateMaxWindStr time.Time ``                      // max wind date
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

// AnemometerBatteryStatus represents the battery status of the anemometer battery
type AnemometerBatteryStatus int

const (
	// AnemometerBatteryMax represents the maximum level of the anemometer battery
	AnemometerBatteryMax AnemometerBatteryStatus = 6000
	// AnemometerBatteryFull represents the full level of the anemometer battery
	AnemometerBatteryFull AnemometerBatteryStatus = 5590
	// AnemometerBatteryHigh represents the high level of the anemometer battery
	AnemometerBatteryHigh AnemometerBatteryStatus = 5180
	// AnemometerBatteryMedium represents the medium level of the anemometer battery
	AnemometerBatteryMedium AnemometerBatteryStatus = 4770
	// AnemometerBatteryLow represents the low level of the anemometer battery
	AnemometerBatteryLow AnemometerBatteryStatus = 4360
)

// String implements the https://golang.org/pkg/fmt/#Stringer interface
func (abs AnemometerBatteryStatus) String() string {
	switch {
	case abs >= AnemometerBatteryMax:
		return "max"
	case abs >= AnemometerBatteryFull:
		return "full"
	case abs >= AnemometerBatteryHigh:
		return "high"
	case abs >= AnemometerBatteryMedium:
		return "medium"
	case abs >= AnemometerBatteryLow:
		return "low"
	default:
		return "very low"
	}
}

// GoString implements the https://golang.org/pkg/fmt/#GoStringer interface
func (abs AnemometerBatteryStatus) GoString() string {
	return fmt.Sprintf("%s (%d)", abs, abs)
}
