package middleware

import (
	"context"
	"net/http"
	"os"
	"strings"

	out "go-api/api/v1/output"

	jwt "github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
)

type Token struct {
	UserId string
	jwt.StandardClaims
}

type Account struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

/*
JwtAuthentication with the jwt, first we set the paths allowed without token(none in this case)
Check if the response contains the token
Generates the access token using the secret key contained on the .env file and compares it with the received one, if they match, we grant access, if not, we deny it.

Returns 401 - "Unauthorized" if the request contains a malformed authorization token

Returns 403 - "Forbidden" if the request does not contain an authorization token
If the token is correct, we perform the request of the user
*/
func JwtAuthentication(next http.Handler) http.Handler {

	return http.HandlerFunc(func(_w http.ResponseWriter, _r *http.Request) {
		log.Info("Starting JWT middleware check")

		notAuth := []string{""}
		requestPath := _r.URL.Path

		//check paths that doesn't require authentication
		for _, value := range notAuth {
			if value == requestPath {
				next.ServeHTTP(_w, _r)
				return
			}
		}

		tokenHeader := _r.Header.Get("Authorization") //Grab the token from the header
		log.Info("Getting Authorization token")

		if tokenHeader == "" {
			log.Error("Missing Authorization token")
			out.RespondWithError(_w, out.ErrMissingToken)
			return
		}

		//The token normally comes in format `Bearer {token-body}`, we check if the retrieved token matched this requirement
		splitted := strings.Split(tokenHeader, " ")
		if len(splitted) != 2 {
			log.Error("Malformed Authorization token")
			out.RespondWithError(_w, out.ErrMalformedToken)
			return
		}

		tk := &Token{}
		//Splitted[1] contains the jwt
		token, err := jwt.ParseWithClaims(splitted[1], tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET_KEY")), nil
		})

		claims, ok := token.Claims.(*Token)

		//Token is invalid, maybe not signed on this server or expired(claims.ExpiresAt == 0) if it expired
		if !ok || !token.Valid {
			log.Error("Token not valid, it may be expired ", token, " this token expires at: ", claims.ExpiresAt)
			out.RespondWithError(_w, out.ErrInvalidToken)
			return
		}

		//Malformed token, returns with http code 403 as usual
		if err != nil {
			log.Error("Malformed authentication token", token)
			out.RespondWithError(_w, out.ErrMalformedToken)
			return
		}

		ctx := context.WithValue(_r.Context(), "user", tk.UserId)
		_r = _r.WithContext(ctx)

		log.Info("JWT  middleware success, next...")
		next.ServeHTTP(_w, _r)
	})

}
