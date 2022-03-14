package login

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/facundoalarcon/oauth2-pkce-sample/platform/authenticator"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	cv "github.com/nirasan/go-oauth-pkce-code-verifier"
	"golang.org/x/oauth2"
	"net/http"
)

// Handler for our login.
func Handler(auth *authenticator.Authenticator) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		state, err := generateRandomState()
		if err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}

		codeVerifier, err := cv.CreateCodeVerifier()
		if err!= nil {
			ctx.String(http.StatusBadRequest, "Could not get code verifier")
			return
		}

		strCodeVerif := codeVerifier.String()

		// Create code_challenge with S256 method
		codeChallenge := codeVerifier.CodeChallengeS256()

		session := sessions.Default(ctx)
		session.Set("state", state)
		session.Set("code_verifier", strCodeVerif)

		// save session values
		if err := session.Save(); err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}

		authCodeOptions := []oauth2.AuthCodeOption {
			oauth2.SetAuthURLParam("code_challenge", codeChallenge),
			oauth2.SetAuthURLParam("code_challenge_method", "S256"),
			oauth2.SetAuthURLParam("response_type", "code"),
		}

		authorizeUrl := auth.AuthCodeURL(
			state,
			authCodeOptions...
		)

		ctx.Redirect(http.StatusTemporaryRedirect, authorizeUrl)
	}
}

func generateRandomState() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	state := base64.StdEncoding.EncodeToString(b)

	return state, nil
}
