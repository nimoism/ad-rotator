package cmd

import (

	// init postgres driver.
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/pressly/goose"
	"github.com/spf13/cobra"

	intsql "github.com/nimoism/ad-rotator/internal/repository/sql"
)

func init() {
	migrateCmd.Flags().StringVar(&dir, "dir", "migrations/", "Path to migrations dir")
}

var (
	dir        string
	migrateCmd = &cobra.Command{
		Use:   "migrate",
		Short: "Migrate DB",
		RunE:  migrate,
	}
)

func migrate(_ *cobra.Command, _ []string) error {
	conf, err := initConfig(configFile)
	if err != nil {
		return err
	}

	dbType, err := intsql.ParseDBType(conf.Storage.DSN)
	if err != nil {
		return err
	}

	var driver string
	switch dbType {
	case intsql.DBTypePostgres:
		driver = "postgres"
	default:
		return errors.New("no DB driver found")
	}
	db, err := goose.OpenDBWithDriver(driver, conf.Storage.DSN)
	if err != nil {
		return err
	}

	return goose.Up(db, dir)
}
