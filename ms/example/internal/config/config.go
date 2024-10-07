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

import "github.com/JohnGomes/Go-Modular-Monolith/internal/config"

type SharedCfg = config.Shared
