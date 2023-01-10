package server

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
	"gitlab.com/aucampia/eg/service-golang/internal/config"
	"gitlab.com/aucampia/eg/service-golang/internal/grpcsvc"
	"gitlab.com/aucampia/eg/service-golang/internal/version"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"

	greeterpb "gitlab.com/aucampia/eg/service-golang/internal/generated/proto/example/greeter/v1"
)

// Server holds the configuration and state for the server.
type Server struct {
	cfg            *config.ServerConfig
	grpcServer     *grpc.Server
	healthServicer *health.Server
}

// NewServer creates a new server object.
func NewServer(ctx context.Context, cfg *config.ServerConfig) (server *Server, err error) {
	serverOpts := []grpc.ServerOption{}
	if cfg.TLS.Enabled {
		creds, err := credentials.NewServerTLSFromFile(cfg.TLS.CertFile, cfg.TLS.KeyFile)
		if err != nil {
			return nil, fmt.Errorf("failed to generate server credentials: %w", err)
		}
		serverOpts = append(serverOpts, grpc.Creds(creds))
	}

	grpcServer := grpc.NewServer(serverOpts...)

	servicer, err := grpcsvc.NewService()
	if err != nil {
		return nil, err
	}

	reflection.Register(grpcServer)

	healthServicer := health.NewServer()
	healthpb.RegisterHealthServer(grpcServer, healthServicer)
	greeterpb.RegisterGreeterAPIServer(grpcServer, servicer)

	server = &Server{
		cfg:            cfg,
		grpcServer:     grpcServer,
		healthServicer: healthServicer,
	}

	return server, nil
}

// Run runs the server.
func (s *Server) Run() error {
	ver, err := version.Get()
	if err != nil {
		return err
	}
	logrus.WithField("version", ver).WithField("address", s.cfg.Address()).Infof("run")

	socket, err := net.Listen("tcp", s.cfg.Address())
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	s.healthServicer.SetServingStatus(
		greeterpb.GreeterAPI_ServiceDesc.ServiceName,
		healthpb.HealthCheckResponse_SERVING,
	)

	if err := s.grpcServer.Serve(socket); err != nil {
		return fmt.Errorf("failed to serve: %w", err)
	}

	logrus.WithField("served", s.cfg.Address()).Infof("run")

	return nil
}

// Shutdown attempts to stop the server gracefully within the shutdown timeout but reverts to a forced shutdown if graceful shutdown does not complete within this time.
func (s *Server) Shutdown() {
	stoppedch := make(chan struct{})
	timer := time.NewTimer(s.cfg.ShutdownTimeout)

	go func() {
		s.grpcServer.GracefulStop()
		close(stoppedch)
	}()

	select {
	case <-timer.C:
		logrus.WithFields(logrus.Fields{
			"timeout": s.cfg.ShutdownTimeout.String(),
		}).Warnf("graceful server shutdown timed out, forcing shutdown")
	case <-stoppedch:
		timer.Stop()
	}
}

// Start starts the server with signal handling.
func (s *Server) Start() error {
	signalch := make(chan os.Signal, 1)
	signal.Notify(signalch, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(signalch)

	errors := make(chan error)
	logrus.Info("running the server in a separate goroutine")
	go func() {
		err := s.Run()
		if err != nil {
			errors <- err
		}
	}()

	logrus.Info("waiting server termination")
	select {
	case msg := <-signalch:
		logrus.Infof("Received shutdown message %v", msg)
		s.Shutdown()
	case err := <-errors:
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Errorf("failed to start the server")
		return fmt.Errorf("failed to start the server: %w", err)
	}
	logrus.Infof("server completed")
	return nil
}
