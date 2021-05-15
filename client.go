package netatmo

import (
	"golang.org/x/oauth2"
)

const (
	NetatmoAPIAuthURL  = "https://api.netatmo.com/oauth2/authorize"
	NetatmoAPITokenURL = "https://api.netatmo.com/oauth2/token"
)

// Client represents the Netatmo API client needed by the subpackages to query the API.
// This package provides a reference implementation. See the Controller struct.
type Client interface {
	ExecuteNetatmoAPIReaquest()
	GetTokens() oauth2.Token
}

// OAuth2Config contains basic OAuth2 related information
// in order to configure a Controller as Client.
type OAuth2Config struct {
	ClientID     string
	ClientSecret string
	Scopes       Scopes
	RedirectURL  string // optional
}

// GetOfflineAuthURL returns a ready to use URL for oauth2 autorization against netatmo API
func GetOfflineAuthURL(conf OAuth2Config, uniqID string) (url string) {
	return GenerateOAuth2Config(conf).AuthCodeURL(uniqID, oauth2.AccessTypeOffline)
}

// GenerateOAuth2Config generates a prefilled OAuth2 conf for the Netatmo API
func GenerateOAuth2Config(conf OAuth2Config) (oac *oauth2.Config) {
	return &oauth2.Config{
		ClientID:     conf.ClientID,
		ClientSecret: conf.ClientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:  NetatmoAPIAuthURL,
			TokenURL: NetatmoAPITokenURL,
		},
		RedirectURL: conf.RedirectURL,
		Scopes:      conf.Scopes.toStrSlice(),
	}
}
