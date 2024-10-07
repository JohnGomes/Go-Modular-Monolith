// Package mono provides embedded microservice.
package mono

import (
	"github.com/JohnGomes/Go-Modular-Monolith/internal/config"
	"github.com/spf13/cobra"
)

type Service struct {
}

// Name implements main.embeddedService interface.
func (s *Service) Name() string { return "mono" }

// Init implements main.embeddedService interface.
func (s *Service) Init(sharedCfg *config.Shared, _, serveCmd *cobra.Command) error {
	return nil
}
