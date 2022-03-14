# Auth0 + Go Web App Sample with PKCE

This sample demonstrates how to add authentication to a Go web app using Auth0 with OAuth 2.0 PKCE flow.

Based on [Auth0 Go Example](https://github.com/auth0-samples/auth0-golang-web-app/tree/master/01-Login), read it to better understand this sample.

## Motivations

Because regular web apps are server-side apps where the source code is not publicly exposed, they can use the 
[Authorization Code Flow](https://auth0.com/docs/get-started/authentication-and-authorization-flow/authorization-code-flow),
which exchanges an Authorization Code for a token. Your app must be server-side because during this exchange, 
you must also pass along your application's Client Secret, which must always be kept secure, 
and you will have to store it in your client.

When public clients (e.g., native and single-page applications) request Access Tokens, 
some additional security concerns are posed that are not mitigated by the Authorization Code Flow alone. 
When cannot securely store a Client Secret you can use Given these situations, OAuth 2.0 provides a version of 
the Authorization Code Flow which makes use of a 
[Proof Key for Code Exchange (PKCE)](https://auth0.com/docs/get-started/authentication-and-authorization-flow/authorization-code-flow-with-proof-key-for-code-exchange-pkce).

The PKCE-enhanced Authorization Code Flow introduces a secret created by the calling application that can be verified by the authorization server; 
this secret is called the Code Verifier. Additionally, the calling app creates a transform value of the Code Verifier called the Code Challenge and 
sends this value over HTTPS to retrieve an Authorization Code. 
This way, a malicious attacker can only intercept the Authorization Code, and they cannot exchange it for a token without the Code Verifier.

## Auth0 App configuration
Follow [Auth0 Go Quickstart](https://auth0.com/docs/quickstart/webapp/golang) stepts and additionally configure in 
`Application Properties` section the `Token Endpoint Authentication Method` value with `None`.

![auth0_app_properties](https://github.com/facundoalarcon/oauth2-pkce-sample/blob/main/doc/resources/auth0_app_properties.jpg)

## Running the App

To run the app, make sure you have **go** installed.

Rename the `.env.example` file to `.env` and provide your Auth0 credentials.

```bash
# .env

AUTH0_CLIENT_ID={CLIENT_ID}
AUTH0_DOMAIN={DOMAIN}
AUTH0_CALLBACK_URL=http://localhost:3000/callback
```

Once you've set your Auth0 credentials in the `.env` file, run `go mod vendor` to download the Go dependencies.

Run `go run main.go` to start the app and navigate to [http://localhost:3000/](http://localhost:3000/).

Alternatively, you could build the docker image and run it.
```
docker build -t auth0-golang-web-app .
docker run --env-file .env -p 3000:3000 -it auth0-golang-web-app
```

## Related resources
- [Add Login Using the Authorization Code Flow with PKCE](https://auth0.com/docs/get-started/authentication-and-authorization-flow/add-login-using-the-authorization-code-flow-with-pkce)
