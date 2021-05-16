package netatmo

import (
	"context"
	"fmt"
	"net/http"

	cleanhttp "github.com/hashicorp/go-cleanhttp"
	"golang.org/x/oauth2"
)

const (
	NetatmoAPIAuthURL  = "https://api.netatmo.com/oauth2/authorize"
	NetatmoAPITokenURL = "https://api.netatmo.com/oauth2/token"
)

// OAuth2BaseConfig contains basic OAuth2 informations allowing to build the full conf with
// GenerateOAuth2Config().
type OAuth2BaseConfig struct {
	ClientID     string
	ClientSecret string
	Scopes       []string // see the string constants begenning with "ScopeXXX"
	// RedirectURL must match the one set on your application profil on the dev portal
	// mandatory if you are using authorization code workflow, leave empty for client credentials workflow
	RedirectURL string
}

// GenerateOAuth2Config generates a complete OAuth2 config for the Netatmo API
func GenerateOAuth2Config(conf OAuth2BaseConfig) (oac oauth2.Config) {
	return oauth2.Config{
		ClientID:     conf.ClientID,
		ClientSecret: conf.ClientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:   NetatmoAPIAuthURL,
			TokenURL:  NetatmoAPITokenURL,
			AuthStyle: oauth2.AuthStyleInParams,
		},
		RedirectURL: conf.RedirectURL,
		Scopes:      conf.Scopes,
	}
}

// GetUserAuthorizationURL returns the Netatmo OAuth2 standard URL for user authorization.
// Unfortunatly, Netatmo requires you to access it with POST, in order for the Netatmo servers to answer back
// the real URL with a 302. This can be annoying if you want to show/transmit your user a simple GET url:
// in that case, you can use the RetreiveUserRealAuthorizationURL() function in order to make the POST for
// you and retreive the real GET generated URL by the Netatmo servers. See step 1 of
// https://dev.netatmo.com/apidocumentation/oauth#authorization-code
func GetUserAuthorizationURL(oac oauth2.Config, uniqID string) (userAuthURL string) {
	return oac.AuthCodeURL(uniqID, oauth2.SetAuthURLParam("redirect_uri", oac.RedirectURL))
}

// RetreiveUserRealAuthorizationURL is an helper to retreive the real auth URL you must redirect your user
// to in order for him to allow your app and trigger your redirect URL set in your app profil.
// THe returned URL is GETabble. customClient can be nil.
func RetreiveUserRealAuthorizationURL(ctx context.Context, oac oauth2.Config, uniqID string,
	customClient *http.Client) (userRealAuthURL string, err error) {

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
	// Generate POST request expecting final URL to be on the 302 Location header response
	generateURL := GetUserAuthorizationURL(oac, uniqID)
	req, err := http.NewRequestWithContext(ctx, "POST", generateURL, nil)
	if err != nil {
		err = fmt.Errorf("failed to forge a HTTP request using '%s' as base URL: %w", generateURL, err)
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
	// Execute request
	resp, err := customClient.Do(req)
	if err != nil {
		err = fmt.Errorf("failed to recover the user auth URL using '%s' as base URL: %w", generateURL, err)
		return
	}
	defer resp.Body.Close()
	// Handle response
	if resp.StatusCode != http.StatusFound {
		err = fmt.Errorf("unexpected return code from '%s' was expecting %d: %s",
			generateURL, http.StatusFound, resp.Status)
		return
	}
	// Parse the URL returned in the Location header
	userRealAuthURL = resp.Header.Get("Location")
	return
}
