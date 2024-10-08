// Package dal implements Data Access Layer using MySQL DB.
package dal

import (
	"context"
	"time"

	"github.com/JohnGomes/Go-Modular-Monolith/ms/example/internal/app"
	"github.com/JohnGomes/Go-Modular-Monolith/ms/example/internal/migrations"
	"github.com/JohnGomes/Go-Modular-Monolith/pkg/repo"
	"github.com/go-sql-driver/mysql"
)

const (
	schemaVersion  = 4
	dbMaxOpenConns = 0 // Unlimited.
	dbMaxIdleConns = 5 // A bit more than default (2).
)

// Ctx is a synonym for convenience.
type Ctx = context.Context

// Repo provides access to storage.
type Repo struct {
	*repo.Repo
}

// New creates and returns new Repo.
// It will also run required DB migrations and connects to DB.
func New(ctx Ctx, dir string, cfg *mysql.Config) (_ *Repo, err error) {
	returnErrs := []error{ // List of app.Err… returned by Repo methods.
		app.ErrNotFound,
	}

	r := &Repo{}
	r.Repo, err = repo.NewFakeDb(ctx, migrations.Goose(), repo.FakeDbConfig{
		FakeDb:        cfg,
		GooseMySQLDir: dir,
		SchemaVersion: schemaVersion,
		Metric:        metric,
		ReturnErrs:    returnErrs,
	})

	if err != nil {
		return nil, err
	}
	r.DB.SetMaxOpenConns(dbMaxOpenConns)
	r.DB.SetMaxIdleConns(dbMaxIdleConns)
	r.SchemaVer.HoldSharedLock(ctx, time.Second)
	return r, nil
}
