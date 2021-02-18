package sql

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"
	"time"

	"github.com/nimoism/ad-rotator/internal/log"
)

const (
	defaultMaxConn = 5
)

const (
	StorageTypeUnknown StorageType = iota
	StorageTypePostgres
	StorageTypeMySQL
)

type StorageType int

func ParseStorageType(dsn string) (StorageType, error) {
	parsedURL, err := url.Parse(dsn)
	if err != nil {
		return StorageTypeUnknown, fmt.Errorf("DSN parsing error: %w", err)
	}
	switch parsedURL.Scheme {
	case "postgres", "postgresql":
		return StorageTypePostgres, nil
	case "mysql":
		return StorageTypeMySQL, nil
	}
	return StorageTypeUnknown, nil
}

func NewConn(ctx context.Context, log log.Logger, driverName, dsn string, timeout time.Duration, maxConn, maxIdleConn int) (*sql.DB, error) {
	if timeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, timeout)
		defer cancel()
	}
	if maxConn == 0 {
		maxConn = defaultMaxConn
	}
	if maxIdleConn == 0 {
		maxIdleConn = maxConn
	}

	var db *sql.DB
	log.Debugf("opening DB (%s, %s)", driverName, dsn)
	db, err := sql.Open(driverName, dsn)
	if err != nil {
		return nil, fmt.Errorf("DB opening error: %w", err)
	}
	db.SetConnMaxIdleTime(0)
	db.SetConnMaxLifetime(0)
	db.SetMaxOpenConns(maxConn)
	db.SetMaxIdleConns(maxIdleConn)
	if err = db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("DB connecting error: %w", err)
	}
	return db, nil
}
