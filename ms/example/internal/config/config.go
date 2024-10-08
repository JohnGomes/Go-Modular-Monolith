// Package config provides configurations for subcommands.
//
// It consists of both configuration values shared by all
// microservices and values specific to this microservice.
//
// Default values can be obtained from various sources (constants,
// environment variables, etc.) and then overridden by flags.
//
// As configuration is global you can get it only once for safety:
// you can call only one of Get… functions and call it just once.
package config

import (
	"fmt"

	"github.com/JohnGomes/Go-Modular-Monolith/internal/config"
	"github.com/JohnGomes/Go-Modular-Monolith/ms/example/internal/app"
	"github.com/JohnGomes/Go-Modular-Monolith/pkg/cobrax"
	"github.com/JohnGomes/Go-Modular-Monolith/pkg/def"
	"github.com/JohnGomes/Go-Modular-Monolith/pkg/netx"
	"github.com/powerman/appcfg"
	"github.com/spf13/pflag"
)

type SharedCfg = config.Shared

var shared *SharedCfg //nolint:gochecknoglobals // Config is global anyway.

// Own configurable values of the microservice.
//
// If microservice may runs in different ways (e.g. using CLI subcommands)
// then these subcommands may use subset of these values.
var own = &struct { //nolint:gochecknoglobals // Config is global anyway.
	MySQLUser     appcfg.NotEmptyString `env:"MYSQL_AUTH_LOGIN"`
	MySQLPass     appcfg.String         `env:"MYSQL_AUTH_PASS"`
	MySQLDBName   appcfg.NotEmptyString `env:"MYSQL_DB_NAME"`
	GooseMySQLDir appcfg.NotEmptyString
	Path          appcfg.NotEmptyString
}{ // Defaults, if any:
	MySQLUser:     appcfg.MustNotEmptyString(app.ServiceName),
	MySQLDBName:   appcfg.MustNotEmptyString(app.ServiceName),
	GooseMySQLDir: appcfg.MustNotEmptyString(fmt.Sprintf("ms/%s/internal/migrations", app.ServiceName)),
	Path:          appcfg.MustNotEmptyString("/rpc"),
}

// FlagSets for all CLI subcommands which use flags to set config values.
type FlagSets struct {
	Serve      *pflag.FlagSet
	GooseMySQL *pflag.FlagSet
}

var fs FlagSets //nolint:gochecknoglobals // Flags are global anyway.

func GetGooseMySQL() (c *cobrax.GooseFakeDbConfig, err error) {
	defer cleanup()

	c = &cobrax.GooseFakeDbConfig{
		MySQL: def.NewMySQLConfig(def.MySQLConfig{
			Addr:   netx.NewAddr(shared.XMySQLAddrHost.Value(&err), shared.XMySQLAddrPort.Value(&err)),
			User:   own.MySQLUser.Value(&err),
			Pass:   own.MySQLPass.Value(&err),
			DBName: own.MySQLDBName.Value(&err),
		}),
		GooseMySQLDir: own.GooseMySQLDir.Value(&err),
	}
	if err != nil {
		return nil, appcfg.WrapPErr(err, fs.GooseMySQL, own, shared)
	}
	return c, nil
}

// Cleanup must be called by all Get* functions to ensure second call to
// any of them will panic.
func cleanup() {
	own = nil
	shared = nil
}
