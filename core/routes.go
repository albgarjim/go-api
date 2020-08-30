package core

import (
	. "goggers/api/v1/handlers"

	log "github.com/sirupsen/logrus"

	"net/http/pprof"
)

const AV1 = "/api/v1"

func (s *Server) CreateRoutes() {
	s.User.HandleFunc("/example", UpdateItem).Methods("PUT")
	s.User.HandleFunc("/example", DeleteItem).Methods("DELETE")
	s.User.HandleFunc("/example", GetItem).Methods("GET")
	s.User.HandleFunc("/example", InsertItem).Methods("POST")

	s.Admin.HandleFunc("/debug/pprof/", pprof.Index)
	s.Admin.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	s.Admin.HandleFunc("/debug/pprof/profile", pprof.Profile)
	s.Admin.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	s.Admin.HandleFunc("/debug/pprof/trace", pprof.Trace)

	s.Admin.Handle("/debug/pprof/block", pprof.Handler("block"))
	s.Admin.Handle("/debug/pprof/goroutine", pprof.Handler("goroutine"))
	s.Admin.Handle("/debug/pprof/heap", pprof.Handler("heap"))
	s.Admin.Handle("/debug/pprof/threadcreate", pprof.Handler("threadcreate"))

	log.Info("------------------------------------------------")
}
