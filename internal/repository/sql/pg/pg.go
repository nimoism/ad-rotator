package pg

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"

	"github.com/nimoism/ad-rotator/internal/log"
	intsql "github.com/nimoism/ad-rotator/internal/repository/sql"
)

const driverName = "pgx"

type Repo struct {
	*BannerRepo
	*SlotRepo
	*UserGroupRepo
}

func NewRepo(ctx context.Context, log log.Logger, dsn string, timeout time.Duration, maxConn, maxIdleConn int) (*Repo, error) {
	driverConfig, err := pgx.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("pgx initializing error: %w", err)
	}
	driverConfig.Logger = &PgxLoggerAdapter{Logger: log}
	dsn = stdlib.RegisterConnConfig(driverConfig)
	db, err := intsql.NewConn(ctx, log, driverName, dsn, timeout, maxConn, maxIdleConn)
	if err != nil {
		return nil, fmt.Errorf("DB initialization error: %w", err)
	}
	log.Debug("DB initialized")

	return &Repo{
		BannerRepo:    NewBannerRepo(db),
		SlotRepo:      NewSlotRepo(db),
		UserGroupRepo: NewUserGroupRepo(db),
	}, nil
}
