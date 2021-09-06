package weather

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/hekmon/go-netatmo"

	"github.com/google/go-querystring/query"
)

const (
	minLat = -85
	maxLat = 85
	minLon = -180
	maxLon = 180
)

// GetPublicDataParameters represents the parameters for GetPublicData()
type GetPublicDataParameters struct {
	NorthEastLatitude  float64 `url:"lat_ne"`                  // latitude of the north east corner of the requested area. -85 <= lat_ne <= 85 and lat_ne>lat_sw
	NorthEastLongitude float64 `url:"lon_ne"`                  // Longitude of the north east corner of the requested area. -180 <= lon_ne <= 180 and lon_ne>lon_sw
	SouthWestLatitude  float64 `url:"lat_sw"`                  // latitude of the south west corner of the requested area. -85 <= lat_sw <= 85
	SouthWestLongitude float64 `url:"lon_sw"`                  // Longitude of the south west corner of the requested area. -180 <= lon_sw <= 180
	RequiredData       string  `url:"required_data,omitempty"` // To filter stations based on relevant measurements you want (e.g. rain will only return stations with rain gauges). Default is no filter.
	Filter             bool    `url:"filter,omitempty"`        // True to exclude station with abnormal temperature measures.
}

// GetPublicData retrieves publicly shared weather data from Outdoor Modules within a predefined area.
func (wc *Client) GetPublicData(ctx context.Context, params GetPublicDataParameters) (publicStations []PublicStationData, headers http.Header, rs netatmo.RequestStats, err error) {
	// verify
	if params.NorthEastLatitude < minLat || params.NorthEastLatitude > maxLat {
		err = errors.New("invalid latitude for North East corner")
		return
	}
	if params.NorthEastLongitude < minLon || params.NorthEastLongitude > maxLon {
		err = errors.New("invalid longitude for North East corner")
		return
	}
	if params.SouthWestLatitude < minLat || params.SouthWestLatitude > maxLat {
		err = errors.New("invalid latitude for South West corner")
		return
	}
	if params.SouthWestLongitude < minLon || params.SouthWestLongitude > maxLon {
		err = errors.New("invalid longitude for South West corner")
		return
	}
	if params.NorthEastLatitude <= params.SouthWestLatitude {
		err = errors.New("north east latitude must be greater than south west latitude")
		return
	}
	if params.NorthEastLongitude <= params.SouthWestLongitude {
		err = errors.New("north east longitude must be greater than south west longitude")
		return
	}
	// prepare parameters
	urlValues, err := query.Values(params)
	if err != nil {
		err = fmt.Errorf("can not convert params as URL values: %w", err)
		return
	}
	// query
	headers, rs, err = wc.client.ExecuteNetatmoAPIRequest(ctx, "GET", "/getpublicdata", urlValues, nil, &publicStations)
	return
}
