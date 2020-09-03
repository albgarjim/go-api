package middleware

import (
	"net/http"
	"strings"
	"sync"
	"time"

	out "go-api/api/v1/output"

	log "github.com/sirupsen/logrus"

	"golang.org/x/time/rate"
)

// Create a custom visitor struct which holds the rate limiter for each
// visitor and the last time that the visitor was seen.
type visitor struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

// Change the the map to hold values of the type visitor.
var visitors = make(map[string]*visitor)
var mtx sync.Mutex

// Run a background goroutine to remove old entries from the visitors map.
func init() {
	log.Info("Initializing rate limiter")
	go cleanupVisitors()
}

// Create a new rate limiter and add it to the visitors map, using the
// IP address as the key.
func addVisitor(_ip string) *rate.Limiter {
	limiter := rate.NewLimiter(2, 5)
	mtx.Lock()
	// Include the current time when creating a new visitor.
	visitors[_ip] = &visitor{limiter, time.Now()}
	mtx.Unlock()
	return limiter
}

// Retrieve and return the rate limiter for the current visitor if it
// already exists. Otherwise call the addVisitor function to add a
// new entry to the map.
func getVisitor(_ip string) *rate.Limiter {
	mtx.Lock()
	v, exists := visitors[_ip]
	if !exists {
		mtx.Unlock()
		return addVisitor(_ip)
	}

	// Update the last seen time for the visitor.
	v.lastSeen = time.Now()
	mtx.Unlock()
	return v.limiter
}

// Every minute check the map for visitors that haven't been seen for
// more than 3 minutes and delete the entries.
func cleanupVisitors() {
	for {
		time.Sleep(30 * time.Second)
		log.Info("Removing visitors with more than 3 minutes")
		mtx.Lock()
		for ip, v := range visitors {
			if time.Now().Sub(v.lastSeen) > 3*time.Minute {
				delete(visitors, ip)
			}
		}
		mtx.Unlock()
	}
}

/*
RateLimiter middleware sets up the request limit based on the user ip

The user ip gets stored on a map and has associated a rate, if the rate is exceeded, a TooManyRequests error is thrown

The map is checked every 3 minutes, the ips that have not called in that time are deleted
*/
func RateLimiter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(_w http.ResponseWriter, _r *http.Request) {
		// Call the getVisitor function to retreive the rate limiter for
		// the current user.
		limiter := getVisitor(strings.Split(_r.RemoteAddr, ":")[0])
		log.Info("Call from ", strings.Split(_r.RemoteAddr, ":")[0])
		if limiter.Allow() == false {
			log.Info("Too many requests from ip: ", strings.Split(_r.RemoteAddr, ":")[0])
			out.RespondWithError(_w, out.ErrTooManyRequests)
			return
		}

		next.ServeHTTP(_w, _r)
	})
}
