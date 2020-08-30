package middleware

import (
	"os"
	"strings"

	"github.com/unrolled/secure"
)

//SecureUser configuration for the user security middleware
var SecureUser = secure.New(secure.Options{
	AllowedHosts:          strings.Split(os.Getenv("CORS_ALLOWED_ORIGINS"), ","),
	HostsProxyHeaders:     []string{"X-Forwarded-Host"},
	SSLRedirect:           true,
	SSLHost:               "ssl.example.com",
	SSLProxyHeaders:       map[string]string{"X-Forwarded-Proto": "https"},
	STSSeconds:            315360000,
	STSIncludeSubdomains:  true,
	STSPreload:            true,
	FrameDeny:             true,
	ContentTypeNosniff:    true,
	BrowserXssFilter:      true,
	IsDevelopment:         true,
	ContentSecurityPolicy: "", //"script-src $NONCE",
	PublicKey:             `pin-sha256="base64+primary=="; pin-sha256="base64+backup=="; max-age=5184000; includeSubdomains; report-uri="https://www.example.com/hpkp-report"`,
})

//SecureAdmin configuration for the admin security middleware
var SecureAdmin = secure.New(secure.Options{
	AllowedHosts:          strings.Split(os.Getenv("CORS_ADMIN_ORIGINS"), ","),
	HostsProxyHeaders:     []string{"X-Forwarded-Host"},
	SSLRedirect:           true,
	SSLHost:               "ssl.example.com",
	SSLProxyHeaders:       map[string]string{"X-Forwarded-Proto": "https"},
	STSSeconds:            315360000,
	STSIncludeSubdomains:  true,
	STSPreload:            true,
	FrameDeny:             true,
	ContentTypeNosniff:    true,
	BrowserXssFilter:      true,
	IsDevelopment:         true,
	ContentSecurityPolicy: "", //"script-src $NONCE",
	PublicKey:             `pin-sha256="base64+primary=="; pin-sha256="base64+backup=="; max-age=5184000; includeSubdomains; report-uri="https://www.example.com/hpkp-report"`,
})
