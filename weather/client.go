package weather

import "github.com/hekmon/go-netatmo"

/*
	https://dev.netatmo.com/apidocumentation/weather
*/

// Client holds the logic to perform weather api calls
type Client struct {
	client netatmo.AuthenticatedClient
}

// New returns a weather api capable client
func New(client netatmo.AuthenticatedClient) *Client {
	return &Client{
		client: client,
	}
}
