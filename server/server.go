package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"os"
)

type Server struct {
	listenAddr string
	log        *logrus.Logger
	r          *mux.Router
}

func New(port string) *Server {
	srv := &Server{
		listenAddr: fmt.Sprintf(":%s", port),
		r:          mux.NewRouter(),
		log: &logrus.Logger{
			Out: os.Stdout,
			Formatter: &logrus.JSONFormatter{
				TimestampFormat: "2006-01-02 15:04:05",
			},
			Level: logrus.DebugLevel,
		},
	}
	srv.Healthz()
	return srv
}

func (srv *Server) Start() {
	srv.log.Infof("server starting. listening on port %s", srv.listenAddr)
	log.Fatalln(http.ListenAndServe(srv.listenAddr, srv.r))
}

func (srv *Server) AddHandler(path string, handler http.HandlerFunc) {
	srv.r.HandleFunc(path, handler)
}

func (srv *Server) Healthz() {
	srv.r.HandleFunc("/healtz", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("ok"))
	})
}
