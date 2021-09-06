package weather

import (
	"encoding/json"
	"fmt"
	"time"
)

// RainMeasures holds measures for the RainGauge module
type RainMeasures struct {
	Rain60min float64   `json:"rain_60min"`
	Rain24h   float64   `json:"rain_24h"`
	RainLive  float64   `json:"rain_live"`
	Time      time.Time `json:"time"` // not in this form on the orignal payload
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
