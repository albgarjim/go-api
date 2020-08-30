package utils

import "github.com/dghubble/sessions"

type key int

const UserID key = 0

const (
	SessionName    = "example-facebook-app"
	SessionSecret  = "example cookie signing secret"
	SessionUserKey = "facebookID"
)

// sessionStore encodes and decodes session data stored in signed cookies
var SessionStore = sessions.NewCookieStore([]byte(SessionSecret), nil)

// Config configures the main ServeMux.
type Config struct {
	FacebookClientID     string
	FacebookClientSecret string
}
