package app

import (
	gSQL "database/sql"
	"fmt"
	"github.com/s4kibs4mi/newschain-cache/config"
	"github.com/s4kibs4mi/newschain-cache/ent"
	"net/url"

	"github.com/facebook/ent/dialect"
	"github.com/facebook/ent/dialect/sql"
	_ "github.com/jackc/pgx/v4/stdlib"
)

var instance *ent.Client
var rawInstance *gSQL.DB

func ConnectSQLDB() error {
	dbUrl := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s", config.DB().Username,
		url.QueryEscape(config.DB().Password), config.DB().Host, config.DB().Port, config.DB().Name)

	db, err := gSQL.Open("pgx", dbUrl)
	if err != nil {
		return err
	}

	db.SetMaxIdleConns(config.DB().MaxIdleConnections)
	db.SetMaxOpenConns(config.DB().MaxOpenConnections)
	db.SetConnMaxLifetime(config.DB().MaxConnectionLifetime)
	rawInstance = db

	driver := sql.OpenDB(dialect.Postgres, db)
	entDriver := ent.Driver(driver)

	if config.App().Mode == "debug" {
		instance = ent.NewClient(entDriver).Debug()
	} else {
		instance = ent.NewClient(entDriver)
	}

	return nil
}

func DB() *ent.Client {
	return instance
}

func RawDB() *gSQL.DB {
	return rawInstance
}
