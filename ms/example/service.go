package example

import (
	"github.com/JohnGomes/Go-Modular-Monolith/ms/example/internal/app"
	"github.com/JohnGomes/Go-Modular-Monolith/ms/example/internal/config"
	"github.com/spf13/cobra"
)

type Service struct {
}

// Name implements main.embeddedService interface.
func (s *Service) Name() string { return app.ServiceName }

// Init implements main.embeddedService interface.
func (s *Service) Init(sharedCfg *config.SharedCfg, cmd, serveCmd *cobra.Command) error {
	return nil
}
