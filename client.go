package netatmo

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	cleanhttp "github.com/hashicorp/go-cleanhttp"
	"golang.org/x/oauth2"
)

// AuthenticatedClient represents a Netatmo API client needed by the subpackages to query the API.
// This package provides a reference implementation, see below.
type AuthenticatedClient interface {
	ExecuteNetatmoAPIRequest(ctx context.Context, method, endpoint string, urlValues url.Values,
		body io.Reader, destination interface{}) (http.Header, RequestStats, error)
	GetTokens() oauth2.Token
}

/*
	Below is a reference OAuth2 netatmo api client wrapper satisfying the Client interface required by the
	sub packages (weather, energy, etc..). You can use this one, or make your own and still being able tu use
	the bindings off the sub packages as long as the interface is respected.
*/

const (
	devMode = true
)

var (
	netatmoAPIURL *url.URL
)

func init() {
	var err error
	if netatmoAPIURL, err = url.Parse(NetatmoAPIBaseURL); err != nil {
		panic(err)
	}
}

// Controller can act as a netatmo API Client.
// Do not instantiate directly, use ExecuteNetatmoAPIReaquest(),
// NewClientWithClientCredentials() or NewClientWithToken() instead.
type Controller struct {
	ctx   context.Context
	token *oauth2.Token
	http  *http.Client
}

// NewClientWithAuthorizationCode returns an initialized and ready to use Netatmo API client.
// authCode is generated by Netatmo once the client has accepted the access thru the auth URL
// (see GetOfflineAuthURL()). Matches to the 4th step of the OAuth2 autorization code grant type:
// https://dev.netatmo.com/apidocumentation/oauth#authorization-code . customClient can be nil.
func NewClientWithAuthorizationCode(ctx context.Context, oac oauth2.Config, authCode string,
	customClient *http.Client) (client AuthenticatedClient, err error) {
	// Spawn a clean client if necessary
	if customClient == nil {
		customClient = cleanhttp.DefaultClient()
	}
	// Prepare the oauth2 enabled client
	c := &Controller{
		ctx: context.WithValue(ctx, oauth2.HTTPClient, customClient),
	}
	// Exchange auth code for access & refresh token
	if c.token, err = oac.Exchange(c.ctx, authCode,
		oauth2.SetAuthURLParam("scope", strings.Join(oac.Scopes, " ")),
		oauth2.SetAuthURLParam("redirect_uri", oac.RedirectURL),
	); err != nil {
		err = fmt.Errorf("can not get oauth2 tokens with authorization code: %w", err)
		return
	}
	// Generate the oauth2 enabled http client
	c.http = oac.Client(c.ctx, c.token)
	// Return the initialized controller as Client
	client = c
	return
}

// NewClientWithClientCredentials returns an initialized and ready to use Netatmo API client.
// https://dev.netatmo.com/apidocumentation/oauth#client-credential customClient can be nil
func NewClientWithClientCredentials(ctx context.Context, oac oauth2.Config, username, password string,
	customClient *http.Client) (client AuthenticatedClient, err error) {
	// Spawn a clean client if necessary
	if customClient == nil {
		customClient = cleanhttp.DefaultClient()
	}
	// Prepare the oauth2 enabled client
	c := &Controller{
		ctx: context.WithValue(ctx, oauth2.HTTPClient, customClient),
	}
	// Get tokens with credentials loging
	if c.token, err = oac.PasswordCredentialsToken(c.ctx, username, password); err != nil {
		err = fmt.Errorf("can not get oauth2 tokens with client credentials: %w", err)
		return
	}
	// Generate the oauth2 enabled http client
	c.http = oac.Client(c.ctx, c.token)
	// Return the initialized controller as Client
	client = c
	return
}

// NewClientWithTokens allows to restore an already authenticated client with saved tokens.
// To check how to retrieve a client tokens, check GetTokens(). customClient can be nil
func NewClientWithTokens(ctx context.Context, oac oauth2.Config, previousTokens *oauth2.Token,
	customClient *http.Client) (client AuthenticatedClient, err error) {
	// Spawn a clean client if necessary
	if customClient == nil {
		customClient = cleanhttp.DefaultClient()
	}
	// Prepare the oauth2 enabled client
	c := &Controller{
		ctx: context.WithValue(ctx, oauth2.HTTPClient, customClient),
	}
	// Restore previous auth
	if previousTokens == nil {
		err = errors.New("can not create a client with nil tokens")
		return
	}
	c.token = previousTokens
	// Generate the oauth2 enabled http client
	c.http = oac.Client(c.ctx, c.token)
	// Return the initialized controller as Client
	client = c
	return
}

// GetTokens returns a copy of the client tokens.
// Might not be safe to call to use while the client is used, use to save state once
// the client is not used any more (ex: stopping/shutdown)
func (c *Controller) GetTokens() (tokens oauth2.Token) {
	return *c.token
}

// ExecuteNetatmoAPIRequest takes care of all the HTTP logic as well as JSON parsing and error handling
func (c *Controller) ExecuteNetatmoAPIRequest(ctx context.Context, method, endpoint string,
	urlValues url.Values, body io.Reader, destination interface{}) (headers http.Header,
	rs RequestStats, err error) {
	// If no meaningfull context is provided for this request, reuse the ctx used by the client
	if ctx == nil || ctx == context.TODO() || ctx == context.Background() {
		ctx = c.ctx
	}
	// Forge request
	reqUrl := *netatmoAPIURL
	reqUrl.Path += endpoint
	reqUrl.RawQuery = urlValues.Encode()
	req, err := http.NewRequestWithContext(ctx, method, reqUrl.String(), body)
	if err != nil {
		err = fmt.Errorf("can not forge HTTP request: %w", err)
		return
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "github.com/hekmon/go-netatmo")
	// Execute request
	resp, err := c.http.Do(req)
	if err != nil {
		err = fmt.Errorf("HTTP request execution failed: %w", err)
		return
	}
	defer resp.Body.Close()
	// Data extraction
	headers = resp.Header
	decoder := json.NewDecoder(resp.Body)
	if devMode {
		decoder.DisallowUnknownFields()
	}
	var unparsedBody json.RawMessage
	// Handle HTTP errors
	switch resp.StatusCode {
	case http.StatusOK:
		// Decode body
		receivedPayload := RequestStatusOKPayload{
			Body: &unparsedBody,
		}
		if err = decoder.Decode(&receivedPayload); err != nil {
			err = fmt.Errorf("encountered %s but failed to decode the body as expected: %w",
				resp.Status, err)
			return
		}
		// Extract stats
		rs.Status = receivedPayload.Status
		rs.TimeExec = time.Duration(receivedPayload.TimeExec * float64(time.Second))
		rs.TimeServer = time.Unix(receivedPayload.TimeServer, 0)
		// Handle errors within body
		if len(receivedPayload.Errors) > 0 {
			err = receivedPayload.Errors
			return
		}
		// continue if no errors
	case http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound,
		http.StatusNotAcceptable, http.StatusInternalServerError:
		receivedPayload := struct {
			Error HTTPStatusGenericError `json:"error"`
		}{
			Error: HTTPStatusGenericError{HTTPCode: resp.StatusCode},
		}
		if err = decoder.Decode(&receivedPayload); err != nil {
			err = fmt.Errorf("encountered %s but fail to decode the body as the expected error: %w",
				resp.Status, err)
			return
		}
		err = receivedPayload.Error
		return
	default:
		var respBody []byte
		if respBody, err = ioutil.ReadAll(resp.Body); err != nil {
			err = fmt.Errorf("failed to read %s body: %w", resp.Status, err)
			return
		}
		err = UnexpectedHTTPCode{
			HTTPCode: resp.StatusCode,
			Body:     respBody,
		}
		return
	}
	// Unmarshall body to dest
	if err = json.Unmarshal(unparsedBody, destination); err != nil {
		err = fmt.Errorf("request successful but can not parse body as JSON: %w", err)
	}
	return
}

// RequestStatusOKPayload is the generic payload received from the netatmo API servers
type RequestStatusOKPayload struct {
	Body       *json.RawMessage   `json:"body"`
	Errors     HTTPStatusOKErrors `json:"errors"`
	Status     string             `json:"status"`      // ex: "ok"
	TimeExec   float64            `json:"time_exec"`   // ex: 0.1116938591003418
	TimeServer int64              `json:"time_server"` // ex: 1621453691
}

// RequestStats contains the request stats sent by the netatmo API servers after a query
type RequestStats struct {
	Status     string
	TimeExec   time.Duration
	TimeServer time.Time
}
