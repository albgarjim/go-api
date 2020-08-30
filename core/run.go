package core

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"context"
	"flag"
	"os/signal"
	"time"

	"github.com/gorilla/handlers"
	log "github.com/sirupsen/logrus"
)

func (s *Server) Run(_addr string) {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	addrStr := fmt.Sprintf(":%s", _addr)
	log.Info("Server running on port: ", _addr)

	allowedHeaders := handlers.AllowedHeaders([]string{"Authorization", "Content-Type"})
	allowedOrigins := handlers.AllowedOrigins(strings.Split(os.Getenv("CORS_ALLOWED_ORIGINS"), ","))
	allowedCredentials := handlers.AllowCredentials()
	allowedMethods := handlers.AllowedMethods([]string{http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodDelete})

	handler := handlers.CORS(allowedHeaders, allowedOrigins, allowedMethods, allowedCredentials)(s.Router)
	srv := &http.Server{
		Addr:    addrStr,
		Handler: handler,
		// Good practice: enforce timeouts for servers you create, it protects against slowloris attacks
		WriteTimeout: 150 * time.Second,
		ReadTimeout:  150 * time.Second,
		IdleTimeout:  time.Second * 600,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Info("shutting down")
	os.Exit(0)
}
