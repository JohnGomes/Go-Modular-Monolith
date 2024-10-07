package auth

import (
	"github.com/JohnGomes/Go-Modular-Monolith/ms/auth/internal/app"
	"github.com/JohnGomes/Go-Modular-Monolith/ms/auth/internal/config"
	"github.com/spf13/cobra"
)

type Service struct {
}

func (s *Service) Name() string { return app.ServiceName }

// Init implements main.embeddedService interface.
func (s *Service) Init(sharedCfg *config.SharedCfg, cmd, serveCmd *cobra.Command) error {

	return nil
}
