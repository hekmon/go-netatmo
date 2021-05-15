package weather

import "github.com/hekmon/go-netatmo"

type WeatherClient struct {
	client netatmo.Client
}

func New(client netatmo.Client) *WeatherClient {
	return &WeatherClient{
		client: client,
	}
}
