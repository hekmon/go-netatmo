package netatmo

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	cleanhttp "github.com/hashicorp/go-cleanhttp"
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

// OAuth2BaseConfig contains basic OAuth2 related information
// in order to configure the full OAuth2 conf.
type OAuth2BaseConfig struct {
	ClientID     string
	ClientSecret string
	Scopes       Scopes
	RedirectURL  string // optional, if set it must match the one defined in the application profil on netatmo dev platform
}

func GetUserAuthorizationURL(ctx context.Context, conf OAuth2BaseConfig, uniqID string, customClient *http.Client) (userAuthURL *url.URL, err error) {
	// Spawn an http client if necessary or clone the given one
	if customClient == nil {
		customClient = cleanhttp.DefaultClient()
	} else {
		clone := *customClient
		customClient = &clone
	}
	// Deactivate redirect follow
	customClient.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}
	// POST request expecting final URL to be on the 302 Location header response
	generateURL := GenerateOAuth2Config(conf).AuthCodeURL(uniqID, oauth2.AccessTypeOffline)
	resp, err := customClient.PostForm(generateURL, nil)
	if err != nil {
		err = fmt.Errorf("failed to recover the user auth URL using '%s' as base URL: %w", generateURL, err)
		return
	}
	if resp.StatusCode != http.StatusFound {
		err = fmt.Errorf("unexpected return code from '%s' was expecting %d: %s", generateURL, http.StatusFound, resp.Status)
		return
	}
	return url.Parse(resp.Header.Get("Location"))
}

// GenerateOAuth2Config generates a prefilled OAuth2 conf for the Netatmo API
func GenerateOAuth2Config(conf OAuth2BaseConfig) (oac *oauth2.Config) {
	return &oauth2.Config{
		ClientID:     conf.ClientID,
		ClientSecret: conf.ClientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:   NetatmoAPIAuthURL,
			TokenURL:  NetatmoAPITokenURL,
			AuthStyle: oauth2.AuthStyleInParams,
		},
		RedirectURL: conf.RedirectURL,
		Scopes:      conf.Scopes.toStrSlice(),
	}
}
