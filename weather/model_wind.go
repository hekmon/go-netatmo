package weather

import (
	"encoding/json"
	"fmt"
	"time"
)

// WindMeasures holds measures for the anemometer module
type WindMeasures struct {
	WindStrength int       `json:"wind_strengh"` // yes the API has a typo on JSON key
	WindAngle    int       `json:"wind_angle"`
	GustStrength int       `json:"gust_strenght"` // yes the API has a typo on JSON key
	GustAngle    int       `json:"gust_angle"`
	Time         time.Time `json:"time"` // not in this form on the orignal payload
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
