package main

import (
	"context"
	"fmt"
	"os"

	"github.com/JohnGomes/Go-Modular-Monolith/internal/config"
	"github.com/JohnGomes/Go-Modular-Monolith/ms/auth"
	"github.com/JohnGomes/Go-Modular-Monolith/ms/example"
	"github.com/JohnGomes/Go-Modular-Monolith/ms/mono"
	"github.com/JohnGomes/Go-Modular-Monolith/pkg/cobrax"
	"github.com/JohnGomes/Go-Modular-Monolith/pkg/def"
	"github.com/spf13/cobra"

	"github.com/powerman/appcfg"
	"github.com/powerman/structlog"
)

// Ctx is a synonym for convenience.
type Ctx = context.Context

type emeddedService interface {
	Name() string
	Init(cfg *config.Shared, cmd, serveCmd *cobra.Command) error
	// RunServe(ctxStartup, ctxShutdown Ctx, shutdown func()) error
}

var (
	embeddedServices = []emeddedService{&mono.Service{}, &auth.Service{}, &example.Service{}}
	log              = structlog.New(structlog.KeyUnit, "main")
	logLevel         = appcfg.MustOneOfString("debug", []string{"debug", "info", "warn", "err"})

	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "Starts embedded microservices",
		Args:  cobra.NoArgs,
		RunE:  runServeWithGracefulShutdown,
	}
)

func main() {
	_ = os.Unsetenv("GO_TEST_DISABLE_SENSITIVE")

	err := def.Init()
	if err != nil {
		log.Fatalf("failed to get defaults: %s", err)
	}

	cfg, err := config.Get()
	if err != nil {
		log.Fatalf("failed to init config: %s", err)
	}

	seen := make(map[string]bool)
	for _, service := range embeddedServices {
		name := service.Name()
		if seen[name] {
			panic(fmt.Sprintf("duplicate service: %s", name))
		}

		seen[name] = true

		cmd := &cobra.Command{
			Use:   name,
			Short: fmt.Sprintf("Run %s microservice's command", name),
			RunE:  cobrax.RequireFlagOrCommand,
		}

		err := service.Init(cfg, cmd, serveCmd)
		if err != nil {
			log.Fatalf("failed to init service: %s: %s", name, err)
		}
		println(name)
	}

}

func runServeWithGracefulShutdown(_ *cobra.Command, _ []string) error {

	return nil
}
