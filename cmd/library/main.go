package main

import (
	"context"
	"flag"
	"log"
	"net/http"

	"github.com/StefanMoller1/go_library/config"
	"github.com/StefanMoller1/go_library/pkg/psql"
	"github.com/StefanMoller1/go_library/repository"
	"github.com/StefanMoller1/go_library/router"
)

const (
	name = "server"
)

func main() {
	var (
		err  error
		conf *config.Config
	)

	// Load CLI configuration
	configFile := flag.String("c", "./etc/"+name+"/config.toml", "path to the config file")
	flag.Parse()

	log := log.Default()

	conf, err = config.NewConfig(*configFile, name)
	if err != nil {
		log.Panic(err, "failed to load config")
	}

	db, err := psql.Connect(conf.Database.DNS, log)
	if err != nil {
		log.Panic(err, "failed to connect to datastore")
	}

	defer db.Close(context.Background())

	err = db.Migrate(log)
	if err != nil {
		log.Panic(err, "failed to migrate datastore")
	}

	host := conf.Server.Host + ":" + conf.Server.Port

	routerManager := new(router.Manager)
	routerManager.Log = log
	routerManager.Library = repository.NewLibraryRepository(db.Conn, log)

	log.Fatal(http.ListenAndServe(host, routerManager.StartRouter(host)))
}
