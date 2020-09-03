package core

import (
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	rdb "go-api/api/v1/models/rethinkdb"

	mdw "go-api/middleware"
)

/*
Server contains all the necessary information to initialize and use the server

Rouer is the base router, it SHOULDN'T be used, create a subrouter instead, e.g. User and admin subrouters

DB is a pointer to a struct that contains the database models
*/
type Server struct {
	Router *mux.Router
	User   *mux.Router
	Admin  *mux.Router
}

func (s *Server) addAdminMiddleware() {
	// s.Admin.Use(mdw.RateLimiter)
	// log.Info("Added RateLimiter middleware to admin router")

	s.Admin.Use(mdw.SecureAdmin.Handler)
	log.Info("Added SecureAdmin middleware to admin router")

	s.Admin.Use(mdw.BasicAuthAdmin)
	log.Info("Added BasicAuthAdmin middleware to admin router")
}

func (s *Server) addUserMiddleware() {
	// s.User.Use(mdw.RateLimiter)
	// log.Info("Added RateLimiter middleware to user router")

	s.User.Use(mdw.SecureUser.Handler)
	log.Info("Added SecureUser middleware to user router")

	s.User.Use(mdw.SecureAdmin.Handler)
	log.Info("Added SecureAdmin middleware to user router")
}

/*
Initialize does the server setup in the following order:

Router initialization, AddSelect of user and admin subroutes to the router, Connect to database, Database initialization, Cache initialization, Creation of routes

NOTE: the initialization of the database MUST be done before the cache initialization
*/
func (s *Server) Initialize(_user, _password, _dbHost string, _dbName string) {
	log.Info("Initializing server")

	s.Router = mux.NewRouter()
	s.User = s.Router.PathPrefix(AV1).Subrouter()
	s.Admin = s.Router.PathPrefix(AV1).Subrouter()

	rdb.InitializeRethinkDB()

	s.addUserMiddleware()
	s.addAdminMiddleware()
	s.CreateRoutes()

	log.Info("Server initialized correctly")
}
