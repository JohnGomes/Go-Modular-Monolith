package cobrax

import (
	"context"
	"fmt"
	"strings"

	"github.com/JohnGomes/Go-Modular-Monolith/pkg/migrate"
	"github.com/go-sql-driver/mysql"
	goosepkg "github.com/powerman/goose/v2"
	"github.com/spf13/cobra"
)

// GooseMySQLConfig contain configuration for goose command.
type GooseFakeDbConfig struct {
	MySQL         *mysql.Config
	GooseMySQLDir string
}

// NewGooseMySQLCmd creates new goose command executed by run.
func NewGooseFakeDbCmd(ctx context.Context, goose *goosepkg.Instance, getCfg func() (*GooseFakeDbConfig, error)) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "goose-fakedb",
		Short: "Migrate FakeDb database schema",
		Args:  gooseArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			gooseCmd := strings.Join(args, " ")

			cfg, err := getCfg()
			if err != nil {
				return fmt.Errorf("failed to get config: %w", err)
			}

			connector := &migrate.FakeDb{Config: cfg.MySQL}
			err = migrate.Run(ctx, goose, cfg.GooseMySQLDir, gooseCmd, connector)
			if err != nil {
				return fmt.Errorf("failed to run goose %s: %w", gooseCmd, err)
			}
			return nil
		},
	}
	cmd.SetUsageTemplate(gooseUsageTemplate)
	return cmd
}
