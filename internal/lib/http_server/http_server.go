package http_server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/oatsmoke/20250905/internal/handler"
	"github.com/oatsmoke/20250905/internal/lib/logger"
)

type Server struct {
	httpServer *http.Server
}

func New(port string, handlers *handler.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:    port,
			Handler: handlers.InitRoutes(),
		},
	}
}

func (s *Server) Run() {
	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	}()

	logger.Info(fmt.Sprintf("http server start on %s", s.httpServer.Addr))
}

func (s *Server) Stop(ctx context.Context) {
	if err := s.httpServer.Shutdown(ctx); err != nil {
		log.Println(err)
	}

	logger.Info("http server stop")
}
