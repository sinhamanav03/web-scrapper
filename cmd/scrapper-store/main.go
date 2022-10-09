package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/sinhamanav03/web-scrapper/internal/config"
	"github.com/sinhamanav03/web-scrapper/internal/db"

	scrapperStoreApi "github.com/sinhamanav03/web-scrapper/api/scrapper-store"
	scrapperStoreSvc "github.com/sinhamanav03/web-scrapper/services/scrapperstore"
	scrapperStoreRepo "github.com/sinhamanav03/web-scrapper/services/scrapperstore/repo"
)

const (
	defaultConfigFile = "./config.json"
)

func main() {
	var configFile string

	flag.StringVar(&configFile, "config", defaultConfigFile, "Config file to load")

	files := strings.Split(configFile, ",")

	conf, err := config.Load(files...)
	if err != nil {
		panic(err)
	}
	db, err := db.NewPostgresDB(conf.Database)
	if err != nil {
		log.Fatalf("error connecting db: %s", err)
	}

	router := mux.NewRouter()

	scrapperStoreApi.RegisterHandler(conf, scrapperStoreSvc.NewScrapperStoreService(scrapperStoreRepo.NewScrapperStoreServiceRepo(db)), router)

	srv := &http.Server{
		Addr:    ":" + fmt.Sprintf("%v", conf.Server.ScrapperStorePort),
		Handler: router,
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal("Server Error", err)
	}

}
