package weather

import "github.com/hekmon/go-netatmo"

/*
	https://dev.netatmo.com/apidocumentation/weather
*/
type Client struct {
	client netatmo.AuthenticatedClient
}

func New(client netatmo.AuthenticatedClient) *Client {
	return &Client{
		client: client,
	}
}
