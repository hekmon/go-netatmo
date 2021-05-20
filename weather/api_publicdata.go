package weather

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hekmon/go-netatmo"

	"github.com/google/go-querystring/query"
)

type GetPublicDataParameters struct {
	LatitudeNorthEast  float64 `url:"lat_ne"`                  // latitude of the north east corner of the requested area. -85 <= lat_ne <= 85 and lat_ne>lat_sw
	LongitudeNorthEast float64 `url:"lon_ne"`                  // Longitude of the north east corner of the requested area. -180 <= lon_ne <= 180 and lon_ne>lon_sw
	LatitudeSouthWest  float64 `url:"lat_sw"`                  // latitude of the south west corner of the requested area. -85 <= lat_sw <= 85
	LongitudeSouthWest float64 `url:"lon_sw"`                  // Longitude of the south west corner of the requested area. -180 <= lon_sw <= 180
	RequiredData       string  `url:"required_data,omitempty"` // To filter stations based on relevant measurements you want (e.g. rain will only return stations with rain gauges). Default is no filter.
	Filter             bool    `url:"filter,omitempty"`        // True to exclude station with abnormal temperature measures.
}

// GetPublicData retrieves publicly shared weather data from Outdoor Modules within a predefined area.
func (wc *Client) GetPublicData(ctx context.Context, params GetPublicDataParameters) (data json.RawMessage,
	headers http.Header, rs netatmo.RequestStats, err error) {
	urlValues, err := query.Values(params)
	if err != nil {
		err = fmt.Errorf("can not convert params as URL values: %w", err)
		return
	}
	headers, rs, err = wc.client.ExecuteNetatmoAPIRequest(ctx, "GET", "/getpublicdata", urlValues, nil, &data)
	return
}

type GetPublicDataAnswer struct {
}
