package app

import (
	"context"
	"fmt"

	"github.com/nimoism/ad-rotator/internal/log"
	"github.com/nimoism/ad-rotator/internal/repository/sql"
	"github.com/nimoism/ad-rotator/internal/repository/sql/pg"
	"github.com/nimoism/ad-rotator/internal/service"
)

func createRepo(ctx context.Context, log log.Logger, conf StorageConf) (service.Repo, error) {
	dbType, err := sql.ParseDBType(conf.DSN)
	if err != nil {
		return nil, fmt.Errorf("DB type error: %w", err)
	}
	switch dbType {
	case sql.DBTypePostgres:
		return pg.NewRepo(ctx, log, conf.DSN, conf.ConnTimeout, conf.MaxConn, conf.MaxIdleConn)
	default:
		return nil, fmt.Errorf("unknown DB type")
	}
}
