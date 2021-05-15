package netatmo

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"golang.org/x/oauth2"
)

/*
	This is one exemple satisfying the Client interface required by the sub packages (weather, energy, etc...).
	You can use this one, or make your own and still being able tu use the bindings off the sub packages.
*/

// Controller can act as a netatmo API Client.
// Do not instantiate directly, use ExecuteNetatmoAPIReaquest(),
// NewClientWithClientCredentials() or NewClientWithToken() instead.
type Controller struct {
	ctx   context.Context
	conf  *oauth2.Config
	token *oauth2.Token
	http  *http.Client
}

// NewClientWithAuthorizationCode returns an initialized and ready to use Netatmo API client.
// authCode is generated by Netatmo once the client has accepted the access thru the auth URL (see GetOfflineAuthURL()).
// Correspond to the 4th step of the OAuth2 autorization code grant type: https://dev.netatmo.com/apidocumentation/oauth#authorization-code
// customClient can be nil
func NewClientWithAuthorizationCode(ctx context.Context, conf OAuth2BaseConfig, authCode string, customClient *http.Client) (client Client, err error) {
	// Prepare the oauth2 enabled client
	c := prepareController(ctx, conf, customClient)
	// Exchange auth code for access & refresh token
	if c.token, err = c.conf.Exchange(ctx, authCode); err != nil {
		client = nil
		err = fmt.Errorf("can not get oauth2 tokens with authorization code: %w", err)
		return
	}
	// Generate the oauth2 enabled http client
	c.http = c.conf.Client(ctx, c.token)
	// Return the initialized controller as Client
	client = c
	return
}

// NewClientWithAuthorizationCode returns an initialized and ready to use Netatmo API client.
// https://dev.netatmo.com/apidocumentation/oauth#client-credential
// customClient can be nil
func NewClientWithClientCredentials(ctx context.Context, conf OAuth2BaseConfig, username, password string, customClient *http.Client) (client Client, err error) {
	// Prepare the oauth2 enabled client
	c := prepareController(ctx, conf, customClient)
	// Exchange auth code for access & refresh token
	if c.token, err = c.conf.PasswordCredentialsToken(c.ctx, username, password); err != nil {
		client = nil
		err = fmt.Errorf("can not get oauth2 tokens with client credentials: %w", err)
		return
	}
	// Generate the oauth2 enabled http client
	c.http = c.conf.Client(c.ctx, c.token)
	// Return the initialized controller as Client
	client = c
	return
}

func NewClientWithToken(ctx context.Context, conf OAuth2BaseConfig, previousTokens *oauth2.Token, customClient *http.Client) (client Client, err error) {
	// Prepare the oauth2 enabled client
	c := prepareController(ctx, conf, customClient)
	// Restore previous auth
	if previousTokens == nil {
		err = errors.New("can not create a client with nil tokens")
		return
	}
	c.token = previousTokens
	// Generate the oauth2 enabled http client
	c.http = c.conf.Client(c.ctx, c.token)
	return
}

func prepareController(ctx context.Context, conf OAuth2BaseConfig, customClient *http.Client) (c *Controller) {
	// Generate OAuth2 config
	c = &Controller{
		conf: GenerateOAuth2Config(conf),
	}
	// Use custom client if any
	if customClient != nil {
		c.ctx = context.WithValue(ctx, oauth2.HTTPClient, customClient)
	} else {
		c.ctx = ctx
	}
	return
}

// GetTokens returns a copy of the client tokens.
// Might not be safe to use while the client is used, use to save state once
// the client is not used any more (ex: stopping/shutdown)
func (c *Controller) GetTokens() (tokens oauth2.Token) {
	return *c.token
}
