package weather

import (
	"encoding/json"
	"fmt"
	"time"
)

// RainMeasures holds measures for the RainGauge module
type RainMeasures struct {
	Time      time.Time `` // not in this form on the orignal payload
	Rain60min float64   `json:"rain_60min"`
	Rain24h   float64   `json:"rain_24h"`
	RainLive  float64   `json:"rain_live"`
}

// UnmarshalJSON allows to create a proper payloade on the fly during JSON unmarshaling
func (rm *RainMeasures) UnmarshalJSON(data []byte) (err error) {
	// Add tmp type
	type OriginalUnmarshal RainMeasures
	tmp := struct {
		RainTimestamp int `json:"rain_timeutc"`
		*OriginalUnmarshal
	}{
		OriginalUnmarshal: (*OriginalUnmarshal)(rm),
	}
	// Unmarshall into the tmp fields
	if err = json.Unmarshal(data, &tmp); err != nil {
		err = fmt.Errorf("failed to unmarshal data to the temporary RainMeasures struct: %w", err)
		return
	}
	// convert
	rm.Time = time.Unix(int64(tmp.RainTimestamp), 0)
	return
}

// RainModuleDashboardData struct for RainModuleDashboardData
type RainModuleDashboardData struct {
	Time      time.Time ``                   // date when data was measured
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
