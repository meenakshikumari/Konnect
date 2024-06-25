package server

import (
"context"
	"github.com/sirupsen/logrus"
	"net/http"
"os"
"os/signal"
"syscall"
"time"

)

// New initializes the server with all the routes
func New(router http.Handler) *Server {
	srv := &Server{
		apiServer: &http.Server{
			Handler: router,
		},
	}
	return srv
}

// Server is a wrapper around http.Server and provides
// Serve method with graceful-shutdown enabled
type Server struct {
	apiServer *http.Server
}

// Serve starts the server and blocks until any termination
// signals and performs graceful shutdown.
func (srv Server) Serve(addr string) {
	srv.apiServer.Addr = addr
	go listenServer(srv.apiServer)
	waitForShutdown(srv.apiServer)
}

func listenServer(apiServer *http.Server) {
	err := apiServer.ListenAndServe()
	if err != http.ErrServerClosed {
		logrus.Fatalf("Failed to listen to server", err.Error())
	}
}

func waitForShutdown(apiServer *http.Server) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig,
		syscall.SIGINT,
		syscall.SIGTERM)

	_ = <-sig

	terminationGracePeriod := time.Duration(5 * int(time.Second))
	ctx, cancel := context.WithTimeout(context.Background(), terminationGracePeriod)
	defer cancel()
	logrus.Infof("API server shutting down")

	if err := apiServer.Shutdown(ctx); err != nil {
		logrus.Errorf("API Server Shutdown Failed: %+v", err)
		return
	}

	logrus.Infof("API Server Exited Properly")
}
