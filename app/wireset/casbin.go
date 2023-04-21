package wireset

import (
	"github.com/casbin/casbin/v2"
	adapter "github.com/casbin/ent-adapter"
	"github.com/one-meta/meta/app/entity/config"
	"path/filepath"
)

func InitCasbin() (*casbin.Enforcer, func(), error) {
	entConfig := config.CFG.Ent
	dbConfig := entConfig.DB
	backend := entConfig.Backend
	var dsn string
	switch backend {
	case "sqlite3":
		dsn = dbConfig.Sqlite3.DSN()
	case "mariadb":
		dsn = dbConfig.MariaDB.DSN()
	case "mysql":
		dsn = dbConfig.MySQL.DSN()
	case "postgres":
		dsn = dbConfig.PostGres.DSN()
	}
	newAdapter, err := adapter.NewAdapter(backend, dsn)
	if err != nil {
		return nil, nil, err
	}
	modelPath := config.CFG.Auth.Casbin.ModelPath
	if modelPath == "" {
		modelPath = "resource"
	}
	modelFile := filepath.Join(modelPath, "model.conf")
	enf, err := casbin.NewEnforcer(modelFile, newAdapter)
	if err != nil {
		return nil, nil, err
	}
	return enf, func() {}, nil
}
