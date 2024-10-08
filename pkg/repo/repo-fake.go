package repo

import (
	"crypto/rsa"
	"crypto/tls"
	"database/sql"
	"strconv"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	goosepkg "github.com/powerman/goose/v2"
	"github.com/powerman/narada4d/schemaver"
	"github.com/powerman/sqlxx"
	"github.com/powerman/structlog"
)

// FakeDBConfig contains repo configuration.
type FakeDbConfig struct {
	FakeDb        *mysql.Config //*FakeConfig
	GooseMySQLDir string
	SchemaVersion int64
	Metric        Metrics
	ReturnErrs    []error // List of app.Err… returned by DAL methods.
}

type FakeConfig struct {
	User             string            // Username
	Passwd           string            // Password (requires User)
	Net              string            // Network type
	Addr             string            // Network address (requires Net)
	DBName           string            // Database name
	Params           map[string]string // Connection parameters
	Collation        string            // Connection collation
	Loc              *time.Location    // Location for time.Time values
	MaxAllowedPacket int               // Max packet size allowed
	ServerPubKey     string            // Server public key name
	pubKey           *rsa.PublicKey    // Server public key
	TLSConfig        string            // TLS configuration name
	tls              *tls.Config       // TLS configuration
	Timeout          time.Duration     // Dial timeout
	ReadTimeout      time.Duration     // I/O read timeout
	WriteTimeout     time.Duration     // I/O write timeout

	AllowAllFiles           bool // Allow all files to be used with LOAD DATA LOCAL INFILE
	AllowCleartextPasswords bool // Allows the cleartext client side plugin
	AllowNativePasswords    bool // Allows the native password authentication method
	AllowOldPasswords       bool // Allows the old insecure password method
	CheckConnLiveness       bool // Check connections for liveness before using them
	ClientFoundRows         bool // Return number of matching rows instead of rows changed
	ColumnsWithAlias        bool // Prepend table alias to column names
	InterpolateParams       bool // Interpolate placeholders into query string
	MultiStatements         bool // Allow multiple statements in one query
	ParseTime               bool // Parse time values to time.Time
	RejectReadOnly          bool // Reject read-only connections
}

func NewFakeDb(ctx Ctx, goose *goosepkg.Instance, cfg FakeDbConfig) (*Repo, error) {

	log := structlog.FromContext(ctx, nil)

	r := &Repo{
		DB:            sqlxx.NewDB(sqlx.NewDb(&sql.DB{}, "mysql")),
		SchemaVer:     &schemaver.SchemaVer{},
		schemaVersion: strconv.Itoa(int(cfg.SchemaVersion)),
		returnErrs:    cfg.ReturnErrs,
		metric:        cfg.Metric,
		log:           log,
		serialize:     fakeDbSerialize,
	}
	return r, nil
}

func fakeDbSerialize(doTx func() error) error {
	// TODO Implement auto-retries if transaction fails because of
	// serialization-related error.
	return doTx()
}
