package server

import (
	"context"
	"net/http"
	"time"

	"github.com/LambdaTest/photon/pkg/lumber"
	"golang.org/x/sync/errgroup"
)

// A Server defines parameters for running an HTTP server.
type Server struct {
	Handler http.Handler
	Addr    string
	Logger  lumber.Logger
}

const timeoutGracefulShutdown = 3 * time.Second

// ListenAndServe initializes a server to respond to HTTP network requests.
func (s Server) ListenAndServe(ctx context.Context) error {
	err := s.listenAndServe(ctx)
	if err == http.ErrServerClosed {
		return nil
	}
	s.Logger.Errorf("Server Shutdown: error %v", err)
	return err
}

func (s Server) listenAndServe(ctx context.Context) error {
	g, ctx := errgroup.WithContext(ctx)

	srv := &http.Server{
		Addr:    s.Addr,
		Handler: s.Handler,
	}
	g.Go(func() error {
		<-ctx.Done()
		s.Logger.Infof("Caller has requested graceful shutdown. shutting down the server")
		ctxShutdown, cancelFunc := context.WithTimeout(context.Background(), timeoutGracefulShutdown)
		defer cancelFunc()
		return srv.Shutdown(ctxShutdown)
	})
	g.Go(srv.ListenAndServe)
	return g.Wait()
}
