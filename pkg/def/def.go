package def

import (
	"fmt"
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/powerman/getenv"
	"github.com/powerman/must"
	"github.com/powerman/sensitive"
	"github.com/powerman/sqlxx"
	"github.com/prometheus/client_golang/prometheus"
)

func init() { //nolint:gochecknoinits // Ensure time.Now() assigned to global vars uses UTC.
	// Make time.Now()==time.Now().UTC() https://github.com/golang/go/issues/19486
	time.Local = nil
}

func Init() error {

	http.DefaultServeMux = nil
	prometheus.DefaultRegisterer = nil
	prometheus.DefaultGatherer = nil

	must.AbortIf = must.PanicIf

	sqlx.NameMapper = sqlxx.ToSnake

	sensitive.Redact()

	// setupLog()

	if hostnameErr != nil {
		return fmt.Errorf("os.Hostname: %w", hostnameErr)
	}
	return getenv.LastErr()

}
