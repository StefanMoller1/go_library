package test

import (
	"context"
	"log"
	"net/http"

	"github.com/StefanMoller1/go_library/config"
	"github.com/StefanMoller1/go_library/pkg/psql"
	"github.com/StefanMoller1/go_library/repository"
	"github.com/StefanMoller1/go_library/router"
)

func setupTestServer(log *log.Logger) {
	conf, err := config.NewConfig("../etc/server/config.toml", "testing")
	if err != nil {
		log.Panic(err, "failed to load config")
	}

	db, err := psql.Connect(conf.Database.DNS, log)
	if err != nil {
		log.Panic(err, "failed to connect to datastore")
	}

	err = db.Migrate(log)
	if err != nil {
		log.Panic(err, "failed to migrate datastore")
	}

	host := conf.Server.Host + ":" + conf.Server.Port

	routerManager := new(router.Manager)
	routerManager.Log = log
	routerManager.Library = repository.NewLibraryRepository(db.Conn, log)

	go http.ListenAndServe(host, routerManager.StartRouter(host))
}

func TearDownTestServer(log *log.Logger) {
	conf, err := config.NewConfig("../etc/server/config.toml", "testing")
	if err != nil {
		log.Panic(err, "failed to load config")
	}

	db, err := psql.Connect(conf.Database.DNS, log)
	if err != nil {
		log.Panic(err, "failed to connect to datastore")
	}

	db.Conn.Exec(context.Background(), "DROP TABLE books;")
}
