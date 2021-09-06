# Netatmo API Golang bindings

[![Go Reference](https://pkg.go.dev/badge/github.com/hekmon/go-netatmo.svg)](https://pkg.go.dev/github.com/hekmon/go-netatmo) [![Go report card](https://goreportcard.com/badge/github.com/hekmon/transmissionrpc)](https://goreportcard.com/report/github.com/hekmon/go-netatmo)

*WIP DO NOT USE*

Netatmo API version: 1.1.2

This lib is splitted into several packages in order for the user to only import what matters to him. The main package (this one) handles everything linked to oauth2 (authentification and token refreshing): it allows to create a client which automatically handles the authentification workflow but does not know any Netatmo product APIs. This main client is then to be used with sub package specific products API:

* [weather](https://github.com/hekmon/go-netatmo/tree/main/weather#readme)
* [energy](https://github.com/hekmon/go-netatmo/tree/main/energy#readme)
* [security](https://github.com/hekmon/go-netatmo/tree/main/security#readme)
* [aircaire](https://github.com/hekmon/go-netatmo/tree/main/aircaire#readme)

The first thing you should do is read the [Netatmo developers guidelines](https://dev.netatmo.com/guideline) and [create an app](https://dev.netatmo.com/apps/createanapp#form) in order to retreive a `client id`, a `client secret` but also setup the application `redirect URI` if you plan to use the standard oauth2 workflow.

## Getting started

First you need to know how you will authenticate against the Netatmo API. There is 2 methods supported:

* The ["authorization code"](https://dev.netatmo.com/apidocumentation/oauth#authorization-code) workflow which is the standard but also the most complex. You will need a reachable webserver (your go program ?) in order to catch the callback URL once the user has been authentified on Netatmo website and allowed your app to access your users data on their behalf. When you plan multi-tenancy (handling multiples users) by opening your service on the internet, this is the one you should consider.
* The ["client credentials"](https://dev.netatmo.com/apidocumentation/oauth#client-credential) workflow, simplier, which allow to get an oauth2 token directly by submiting a user/password combination. Perfect for personnal projects, you should still consider to not use your admin/main account but instead create a secondary account with limited privileges, invite it to your main account and then use the secondary account credentials for your program.

Both methods require an oauth2 configuration which you configure like this:

```golang
oauthConfig := netatmo.GenerateOAuth2Config(netatmo.OAuth2BaseConfig{
    ClientID:     ClientID, // retreived in your developers app netatmo page
    ClientSecret: SecretID, // retreived in your developers app netatmo page
    Scopes: []string{
        netatmo.ScopeStationRead,
        netatmo.ScopeThermostatRead,
    }, // see the the scopes.go file for all availables scopes
    RedirectURL: "https://yourdomain.tld/netatmocallback", // set up in your developers app netatmo page
})
```

If you do not plan to use the authorization code workflow, I still recommend you to setup an arbitraty redirect URL (ideally on a domain you own, even if the endpoint is fake) and set it up on your oauth config as well in order for netatmo to match your app profil on theirs servers (I got silent fails during testing because of it).

### Client credentials

Simply init the client with username/password:

```golang
client, err := netatmo.NewClientWithClientCredentials(context.TODO(), config, username, password, nil)
if err != nil {
    panic(err)
}
```

Keep in mind that should never used/store your admin account credentials ! (prefer a low privileges, secondary account)

### Authorization code

As described by the [Netatmo website](https://dev.netatmo.com/apidocumentation/oauth#authorization-code), oauth2 three-legged workflow needs several steps.

#### Getting the authorization URL

To get valid tokens to act as your user, you will need to redirect him on the oauth2 authorization page at netatmo servers:

```golang
userRandomID := "somethingrandomthatyouwillstore"
authURL := netatmo.GetUserAuthorizationURL(oac oauth2.Config, userRandomID string)
```

This will yield you an URL in order to your user to validate on the netatmo servers that it allows your app to act on its behalf for the scopes you declared on your oauth configuration. Unfortunatly, the URL is not reachable by GET but POST only; this POST will then yield the real GETable URL in the `Location` HTTP header by issuing a 302 response. This can be fine if you want to redirect your user on the netatmo servers using browser side javascript but can be anoying if you need to recover (and for example print) the real GETable URL yourself. This lib contains an helper function which will do it for you if you need it to:

```golang
userRandomID := "somethingrandomthatyouwillstore"
authURL, err := netatmo.RetreiveUserRealAuthorizationURL(context.TODO(), oauthConfig, userRandomID, nil)
if err != nil {
    panic(err)
}
fmt.Println(authURL)
```

#### Receiving the auth code from Netatmo

Once your user validates the scopes and your app on the Netatmo Oauth2 page, the Netatmo servers will send an authentication code back to your redirect URL along with the userID you setup (allowing you to match which user this auth code belongs to). You will need an HTTP handler to catch them as they are passed as query parameters: `[YOUR_REDIRECT_URI]?state=[USER_RANDOM_ID_YOU_PROVIDED]&code=[NETATMO_GENERATED_CODE]`

Once retreived, you can use the uniq ID in the `state` query parameter to match the user you assigned it to and use the auth code to retreive the oauth2 tokens and finally have the authenticated netatmo client:

```golang
authedClient, err := netatmo.NewClientWithAuthorizationCode(context.TODO(), oauthConfig, netatmoGeneratedCode, nil)
if err != nil {
    panic(err)
}
```

The lib will handle the oauth2 process, retreives the oauth2 tokens from this auth code and yield an authenticated client.

### Restoring a client from tokens

Once initialized, no matter if it was from the client credentials or authorization code workflow, you can (and should) store the oauth2 tokens somewhere in order for you to re init an auth client without redoing the whole auth process.

To retreive a client's oauth2 tokens simply do:

```golang
tokens := authedClient.GetTokens()
```

This will yield you a struct with JSON tags. Store it anyway you like. Then in order to restore a client, restore this struct and directly init the client:

```golang
tokens := retreivedSavedTokens(user)
authedClient, err := netatmo.NewClientWithTokens(context.TODO(), oauthConfig, tokens, nil)
if err != nil {
    panic(err)
}
```

## I have my authenticated client, now what ?

You can now init products API clients using the `authedClient` that will handle API requests authentication and oauth2 tokens auto refresh.

* [weather](https://github.com/hekmon/go-netatmo/tree/main/weather#readme)
* [energy](https://github.com/hekmon/go-netatmo/tree/main/energy#readme)
* [security](https://github.com/hekmon/go-netatmo/tree/main/security#readme)
* [aircaire](https://github.com/hekmon/go-netatmo/tree/main/aircaire#readme)
