package middleware

import (
	"net/http"
	"os"

	out "goggers/api/v1/output"

	log "github.com/sirupsen/logrus"
)

/*
BasicAuthUser contains the middleware functionality to authenticate an admin

The admin is authenticated using the BasicAuth method against the BASIC_AUTH_USER and BASIC_AUTH_PASS environment variables
*/
func BasicAuthUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(_w http.ResponseWriter, _r *http.Request) {

		log.Info("Starting basic auth middleware check")
		user, pass, authOK := _r.BasicAuth()

		if authOK == false {
			log.Info("Error with basic user authentication")
			out.RespondWithError(_w, out.ErrWrongBasicAuth)
			return
		}

		if os.Getenv("BASIC_AUTH_USER") != user || os.Getenv("BASIC_AUTH_PASS") != pass {
			log.Info("Invalid username or password, unhautorized")
			_w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			out.RespondWithError(_w, out.ErrWrongBasicAuth)
			return
		}
		log.Info("Basic auth user middleware success, next...")
		next.ServeHTTP(_w, _r)
	})
}
