package netatmo

import (
	"golang.org/x/oauth2"
)

// Client represents the Netatmo API client needed by the subpackages to query the API.
// This package provides a reference implementation. See the Controller struct.
type Client interface {
	ExecuteNetatmoAPIRequest()
	GetTokens() oauth2.Token
}
