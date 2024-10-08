// Package repo provide helpers for Data Access Layer.
package repo

import (
	"context"
	"errors"

	"github.com/powerman/narada4d/schemaver"
	"github.com/powerman/sqlxx"
	"github.com/powerman/structlog"
)

// Ctx is a synonym for convenience.
type Ctx = context.Context

// Errors.
var (
	ErrSchemaVer = errors.New("unsupported DB schema version")
)

// Repo provides access to storage.
type Repo struct {
	DB            *sqlxx.DB
	SchemaVer     *schemaver.SchemaVer
	schemaVersion string
	returnErrs    []error
	metric        Metrics
	log           *structlog.Logger
	serialize     func(doTx func() error) error
}

// Close closes connection to DB.
func (r *Repo) Close() {
	r.log.WarnIfFail(r.DB.Close)
	r.log.WarnIfFail(r.SchemaVer.Close)
}
