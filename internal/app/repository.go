package app

import (
	"context"
	"time"

	"github.com/pkg/errors"

	"github.com/nimoism/ad-rotator/internal/log"
	"github.com/nimoism/ad-rotator/internal/repository/sql"
	"github.com/nimoism/ad-rotator/internal/repository/sql/pg"
	"github.com/nimoism/ad-rotator/internal/service"
)

func createRepo(ctx context.Context, log log.Logger, conf StorageConf) (service.Repo, error) {
	dbType, err := sql.ParseStorageType(conf.DSN)
	if err != nil {
		return nil, err
	}
	attempts := 0
	var repo service.Repo

	for {
		switch dbType {
		case sql.StorageTypePostgres:
			repo, err = pg.NewRepo(ctx, log, conf.DSN, conf.ConnTimeout, conf.MaxConn, conf.MaxIdleConn)
		default:
			return nil, errors.New("unknown storage type")
		}

		if err != nil {
			if attempts < conf.Retry {
				log.Warnf("%v, retrying...", err)
				time.Sleep(conf.RetryInterval)
				attempts++
				continue
			} else {
				return nil, err
			}
		}

		return repo, nil
	}
}
