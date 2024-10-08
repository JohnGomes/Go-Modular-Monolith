package auth

import (
	"context"

	"github.com/JohnGomes/Go-Modular-Monolith/ms/auth/internal/app"
	"github.com/JohnGomes/Go-Modular-Monolith/ms/auth/internal/config"
	"github.com/powerman/structlog"
	"github.com/spf13/cobra"
)

// Ctx is a synonym for convenience.
type Ctx = context.Context

type Service struct {
}

func (s *Service) Name() string { return app.ServiceName }

// Init implements main.embeddedService interface.
func (s *Service) Init(sharedCfg *config.SharedCfg, cmd, serveCmd *cobra.Command) error {

	return nil
}

// RunServe implements main.embeddedService interface.
func (s *Service) RunServe(ctxStartup, ctxShutdown Ctx, shutdown func()) (err error) {

	log := structlog.FromContext(ctxShutdown, nil)
	log.Err("failed to serve")
	// log := structlog.FromContext(ctxShutdown, nil)
	// if s.cfg == nil {
	// 	s.cfg, err = config.GetServe()
	// }
	// if err == nil {
	// 	s.ca, err = netx.LoadCACert(s.cfg.TLSCACert)
	// }
	// if err == nil {
	// 	s.cert, err = tls.LoadX509KeyPair(s.cfg.TLSCert, s.cfg.TLSKey)
	// }
	// if err == nil {
	// 	s.certInt, err = tls.LoadX509KeyPair(s.cfg.TLSCertInt, s.cfg.TLSKeyInt)
	// }
	// if err != nil {
	// 	return log.Err("failed to get config", "err", err)
	// }

	// err = concurrent.Setup(ctxStartup, map[interface{}]concurrent.SetupFunc{
	// 	&s.repo: s.connectRepo,
	// })
	// if err != nil {
	// 	return log.Err("failed to connect", "err", err)
	// }

	// if s.appl == nil {
	// 	s.appl = app.New(s.repo, app.Config{
	// 		Secret: s.cfg.Secret,
	// 	})
	// }

	// s.srv = grpc.NewServer(s.appl, grpc.Config{
	// 	CtxShutdown: ctxShutdown,
	// 	Cert:        &s.cert,
	// })
	// s.srvInt = grpc.NewServerInt(s.appl, grpc.Config{
	// 	CtxShutdown: ctxShutdown,
	// 	Cert:        &s.certInt,
	// })
	// s.mux, err = grpcgw.NewServer(grpcgw.Config{
	// 	CtxShutdown:      ctxShutdown,
	// 	Endpoint:         s.cfg.AuthAddr,
	// 	CA:               s.ca,
	// 	GRPCGWPattern:    "/",
	// 	OpenAPIPattern:   "/openapi/", // Also hardcoded in web/static/swagger-ui/index.html.
	// 	SwaggerUIPattern: "/swagger-ui/",
	// })
	// if err != nil {
	// 	return log.Err("failed to setup grpc-gateway", "err", err)
	// }

	// err = concurrent.Serve(ctxShutdown, shutdown,
	// 	s.serveMetrics,
	// 	s.serveGRPC,
	// 	s.serveGRPCInt,
	// 	s.serveGRPCGW,
	// )
	// if err != nil {
	// 	return log.Err("failed to serve", "err", err)
	// }
	return nil
}

// func (s *Service) connectRepo(ctx Ctx) (interface{}, error) {
// 	return dal.New(ctx, s.cfg.GoosePostgresDir, s.cfg.Postgres)
// }

// func (s *Service) serveMetrics(ctx Ctx) error {
// 	return serve.Metrics(ctx, s.cfg.BindMetricsAddr, reg)
// }

// func (s *Service) serveGRPC(ctx Ctx) error {
// 	return serve.GRPC(ctx, s.cfg.BindAddr, s.srv, "gRPC")
// }

// func (s *Service) serveGRPCInt(ctx Ctx) error {
// 	return serve.GRPC(ctx, s.cfg.BindAddrInt, s.srvInt, "gRPC internal")
// }

// func (s *Service) serveGRPCGW(ctx Ctx) error {
// 	tlsConfig := &tls.Config{
// 		Certificates: []tls.Certificate{s.cert},
// 		MinVersion:   tls.VersionTLS12,
// 	}
// 	return serve.HTTP(ctx, s.cfg.BindGRPCGWAddr, tlsConfig, s.mux, "grpc-gateway")
// }
