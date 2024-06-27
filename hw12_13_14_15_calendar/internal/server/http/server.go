package internalhttp

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/fixme_my_friend/hw12_13_14_15_calendar/internal/app"
)

type Server struct {
	logger app.Logger
	server *http.Server
}

func NewServer(logger app.Logger, address string, handler http.Handler) *Server {
	return &Server{
		logger: logger,
		server: &http.Server{
			Addr:           address,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
			Handler:        handler,
		},
	}
}

func (s *Server) Start(ctx context.Context) error {
	go func() {
		if err := s.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.logger.Error(err.Error())
		}
	}()
	<-ctx.Done()
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	s.logger.Error("Server stoped")
	if err := s.server.Shutdown(ctx); err != nil {
		s.logger.Error("shutdown error: " + err.Error())
	}
	return nil
}
