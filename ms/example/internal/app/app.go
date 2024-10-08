//go:generate -command mockgen sh -c "$(git rev-parse --show-toplevel)/.gobincache/$DOLLAR{DOLLAR}0 \"$DOLLAR{DOLLAR}@\"" mockgen
//go:generate mockgen -package=$GOPACKAGE -source=$GOFILE -destination=mock.$GOFILE

// Package app provides business logic.

package app

import (
	"context"
	"errors"
	"time"
)

// Ctx is a synonym for convenience.
type Ctx = context.Context

// ServiceName provides name of this microservice for logs/metrics.
const ServiceName = "example"

// Errors.
var (
	ErrAccessDenied = errors.New("access denied")
	ErrNotFound     = errors.New("not found")
)

// Repo provides data storage.
type Repo interface {
	// Example returns ...
	// Errors: ErrNotFound.
	Example(Ctx, dom.UserName) (*Example, error)
	// IncExample creates or increments ...
	// Errors: none.
	IncExample(Ctx, dom.UserName) error
}

type (
	// Example describes ...
	Example struct {
		Counter int
		Mtime   time.Time
	}
)
