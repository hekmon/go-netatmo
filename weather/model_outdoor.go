package weather

import (
	"encoding/json"
	"fmt"
	"time"
)

// OutdoorModuleDashboardData struct for OutdoorModuleDashboardData
type OutdoorModuleDashboardData struct {
	Time        time.Time ``                   // date when data was measured
	Temperature float64   `json:"Temperature"` // temperature (in °C)
	Humidity    int64     `json:"Humidity"`    // humidity (in %)
	MinTemp     float64   `json:"min_temp"`    // minimum temperature measured
	MaxTemp     float64   `json:"max_temp"`    // maximum temperature measured
	DateMinTemp time.Time ``                   // date of minimum temperature measured
	DateMaxTemp time.Time ``                   // date of maximum temperature measured
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
