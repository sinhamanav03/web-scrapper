package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gocolly/colly/v2"
	"github.com/gorilla/mux"
	scrapperApi "github.com/sinhamanav03/web-scrapper/api/scrapper"
	"github.com/sinhamanav03/web-scrapper/internal/config"
	scrapperSvc "github.com/sinhamanav03/web-scrapper/services/scrapper"
)

const (
	defaultConfigFile = "./config.json,.env"
)

func main() {
	var configFile string

	flag.StringVar(&configFile, "config", defaultConfigFile, "Cconfig file to load")

	files := strings.Split(configFile, ",")

	conf, err := config.Load(files...)
	if err != nil {
		panic(err)
	}

	router := mux.NewRouter()

	c := colly.NewCollector()

	scrapperApi.RegisterHandler(conf, scrapperSvc.NewScrapperService(conf, c), router)

	srv := &http.Server{
		Addr:    ":" + fmt.Sprintf("%v", conf.Server.ScrapperPort),
		Handler: router,
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal("Server Error", err)
	}

}
