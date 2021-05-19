package weather

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

type GetStationDataParameters struct {
	DeviceID     string `url:"device_id,omitempty"`     // Weather station mac address
	GetFavorites bool   `url:"get_favorites,omitempty"` // To retrieve user's favorite weather stations
}

// GetStationData returns data from a user Weather Stations (measures and device specific data)
func (wc *Client) GetStationData(ctx context.Context, params GetStationDataParameters) (data json.RawMessage, headers http.Header, err error) {
	urlValues, err := query.Values(params)
	if err != nil {
		err = fmt.Errorf("can not convert params as URL values: %w", err)
		return
	}
	headers, err = wc.client.ExecuteNetatmoAPIRequest(ctx, "GET", "/getstationsdata", urlValues, nil, &data)
	return
}
