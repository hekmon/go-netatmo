package weather

import "github.com/hekmon/go-netatmo"

/*
	https://dev.netatmo.com/apidocumentation/weather
*/
type WeatherClient struct {
	client netatmo.AuthenticatedClient
}

func New(client netatmo.AuthenticatedClient) *WeatherClient {
	return &WeatherClient{
		client: client,
	}
}

