package api

import (
	"net/http"

	"github.com/didiegovieira/go-benchmark-api/pkg/route"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	Gin           *gin.Engine
	server        *http.Server
	WebServerPort string
}

func NewServer(serverPort string) *Server {
	s := &Server{
		Gin:           gin.Default(),
		WebServerPort: serverPort,
	}

	return s
}

func (s *Server) RegisterRoutes(routes []route.RouteInterface) {
	for _, routeInterface := range routes {
		getRoute := routeInterface.GetRoute()
		s.Gin.Handle(getRoute.Method, getRoute.Path, getRoute.Handlers...)
	}
}

func (s *Server) Start() {
	log.Infoln("Starting web server on port", s.WebServerPort)
	s.server = &http.Server{
		Addr:    ":" + s.WebServerPort,
		Handler: s.Gin,
	}
	err := s.server.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}
