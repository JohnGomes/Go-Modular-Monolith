package dal

import (
	"github.com/JohnGomes/Go-Modular-Monolith/ms/example/internal/app"
	"github.com/JohnGomes/Go-Modular-Monolith/pkg/repo"
	"github.com/prometheus/client_golang/prometheus"
)

var metric repo.Metrics //nolint:gochecknoglobals // Metrics are global anyway.

// InitMetrics must be called once before using this package.
// It registers and initializes metrics used by this package.
func InitMetrics(reg *prometheus.Registry) {
	const subsystem = "dal_mysql"

	metric = repo.NewMetrics(reg, app.ServiceName, subsystem, new(app.Repo))
}
