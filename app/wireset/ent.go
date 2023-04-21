package wireset

import (
	"context"
	"database/sql"
	"log"
	"github.com/one-meta/meta/app/entity/config"
	"time"

	"github.com/one-meta/meta/app/ent"

	entsql "entgo.io/ent/dialect/sql"
	"github.com/casbin/ent-adapter/ent/migrate"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

func InitEnt() (*ent.Client, func(), error) {
	var (
		entConfig  = config.CFG.Ent
		client     *ent.Client
		cleanUn    func()
		driverName string
		dsn        string
		err        error
	)

	backend := entConfig.Backend
	if backend == "sqlite3" {
		client, err = ent.Open("sqlite3", entConfig.DB.Sqlite3.DSN())
		cleanUn = func() {
			err := client.Close()
			if err != nil {
				return
			}
		}
		if err != nil {
			log.Printf("failed opening connection to sqlite: %v", err)
		}
	} else {
		backend, driverName, dsn = configParser(config.CFG, driverName)
		db, err := sql.Open(driverName, dsn)
		if err != nil {
			return nil, nil, err
		}
		db.SetMaxIdleConns(entConfig.DB.MaxIdleConns)
		db.SetMaxOpenConns(entConfig.DB.MaxOpenConns)
		db.SetConnMaxLifetime(time.Hour * time.Duration(entConfig.DB.ConnMaxLifetime))
		drv := entsql.OpenDB(backend, db)
		client = ent.NewClient(ent.Driver(drv))
		cleanUn = func() {
			err := db.Close()
			if err != nil {
				return
			}
			err = drv.Close()
			if err != nil {
				return
			}
			err = client.Close()
			if err != nil {
				return
			}
		}
	}

	if entConfig.DebugMode {
		client = client.Debug()
	}
	if entConfig.AutoMigrate {
		ctx := context.Background()
		if err := client.Schema.Create(
			ctx,
			migrate.WithDropIndex(entConfig.WithDropIndex),
			migrate.WithDropColumn(entConfig.WithDropColumn),
			migrate.WithGlobalUniqueID(entConfig.WithGlobalUniqueID),
		); err != nil {
			log.Fatalf("failed creating schema resources: %v", err)
			// log.Printf("failed creating schema resources: %v", err)
		}
	}
	return client, cleanUn, nil
}

func configParser(config *config.Config, driverName string) (string, string, string) {
	var dsn string
	backend := config.Ent.Backend
	dbConfig := config.Ent.DB
	switch backend {
	case "mysql":
		driverName = backend
		dsn = dbConfig.MySQL.DSN()
	case "mariadb":
		driverName = "mysql"
		backend = "mysql"
		dsn = dbConfig.MariaDB.DSN()
	case "postgres":
		driverName = "pgx"
		dsn = dbConfig.PostGres.DSN()
	default:
		log.Fatal("database backend not supported")
	}
	return backend, driverName, dsn
}
