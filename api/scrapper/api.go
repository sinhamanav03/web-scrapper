package scrapper

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sinhamanav03/web-scrapper/internal/config"
	"github.com/sinhamanav03/web-scrapper/internal/utils"
	scrapperSvc "github.com/sinhamanav03/web-scrapper/services/scrapper"
)

type ScapperHandler struct {
	conf *config.Config
	svc  scrapperSvc.ScrapperService
}

func RegisterHandler(conf *config.Config, svc scrapperSvc.ScrapperService, router *mux.Router) {
	res := ScapperHandler{
		conf: conf,
		svc:  svc,
	}

	router.HandleFunc("/scrape", res.Scrape).Methods("POST")
}

type ScrapeRequest struct {
	Url string `json:"url"`
}

func (res *ScapperHandler) Scrape(w http.ResponseWriter, r *http.Request) {

	req := ScrapeRequest{}

	err := utils.DecodeRequest(r, &req)
	if err != nil {
		utils.JsonResponse(w, http.StatusInternalServerError, nil)
		return
	}
	resp, err := res.svc.ScrapeUrl(r.Context(), req.Url)
	if err != nil {
		utils.JsonResponse(w, http.StatusInternalServerError, nil)
		return
	}

	utils.JsonResponse(w, http.StatusOK, resp)
}
