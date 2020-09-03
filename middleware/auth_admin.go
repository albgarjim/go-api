package middleware

import (
	out "go-api/api/v1/output"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

/*
AdminAuth contains the middleware functionality to authenticate an admin

The admin is authenticated using the BasicAuth method against the ADMIN_AUTH_USER and ADMIN_AUTH_PASS environment variables
*/
func BasicAuthAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(_w http.ResponseWriter, _r *http.Request) {

		log.Info("Starting basic auth middleware check")
		user, pass, authOK := _r.BasicAuth()

		if authOK == false {
			log.Info("Error with basic admin authentication")
			out.RespondWithError(_w, out.ErrWrongBasicAuth)
			return
		}

		if os.Getenv("ADMIN_AUTH_USER") != user || os.Getenv("ADMIN_AUTH_PASS") != pass {
			log.Info("Invalid username or password, unhautorized")
			_w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			out.RespondWithError(_w, out.ErrWrongBasicAuth)
			return
		}
		log.Info("Basic auth admin middleware success, next...")
		next.ServeHTTP(_w, _r)
	})
}
