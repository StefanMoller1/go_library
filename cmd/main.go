package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/StefanMoller1/go_library/config"
	"github.com/StefanMoller1/go_library/models"
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

	db, err := models.Connect(conf, log)
	if err != nil {
		log.Panic(err, "failed to connect to datastore")
	}

	err = db.Migrate()
	if err != nil {
		log.Panic(err, "failed to migrate datastore")
	}

	log.Fatal(http.ListenAndServe(conf.Server.Host+":"+conf.Server.Port, router.StartRouter(db, log)))
}
