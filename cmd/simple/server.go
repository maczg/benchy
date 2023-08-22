package simple

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/massimo-gollo/benchy/pkg/log"
	"github.com/sirupsen/logrus"
	"net/http"
)

const ServiceName = "simple"

type Server struct {
	name       string
	logger     *logrus.Logger
	listenAddr string
	router     *mux.Router
}

func (srv *Server) Start() {
	srv.logger.Infof("%s server starting. listening on %s", srv.name, srv.listenAddr)
	srv.logger.Fatalln(http.ListenAndServe(srv.listenAddr, srv.router))
}

func (srv *Server) AddHandler(path string, handler http.HandlerFunc) {
	srv.router.HandleFunc(path, handler)
}

func (srv *Server) AddMiddleware(mw mux.MiddlewareFunc) {
	srv.router.Use(mw)
}

func (srv *Server) Logger() *logrus.Logger {
	return srv.logger
}

func NewServer(serviceName, port string) *Server {
	srv := &Server{
		name:       serviceName,
		listenAddr: fmt.Sprintf(":%s", port),
		logger:     log.NewDefaultLogger(),
		router:     mux.NewRouter(),
	}
	srv.router.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("ok\n"))
	})
	return srv
}
