package scrapperstore

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sinhamanav03/web-scrapper/internal/config"
	"github.com/sinhamanav03/web-scrapper/internal/utils"
	scrapperstore "github.com/sinhamanav03/web-scrapper/services/scrapperstore"
	"github.com/sinhamanav03/web-scrapper/services/scrapperstore/model"
)

type ScapperHandler struct {
	svc  scrapperstore.ScrapperStoreService
	conf *config.Config
}

func RegisterHandler(conf *config.Config, svc scrapperstore.ScrapperStoreService, router *mux.Router) {
	res := ScapperHandler{
		conf: conf,
		svc:  svc,
	}

	router.HandleFunc("/add", res.Add).Methods("POST")
	router.HandleFunc("/get/{id}", res.Get).Methods("GET")
	router.HandleFunc("/update/{id}", res.Update).Methods("PUT")
	router.HandleFunc("/delete/{id}", res.Delete).Methods("DELETE")
}

type status struct {
	Success bool `json:"success"`
}
type addRequest struct {
	URL     string     `json:"url"`
	Product model.Item `json:"product"`
}

type addResp struct {
	status
	Id int `json:"user_id,omitempty"`
}

func (handler *ScapperHandler) Add(w http.ResponseWriter, r *http.Request) {
	item := addRequest{}
	resp := addResp{}

	err := utils.DecodeRequest(r, &item)
	if err != nil {
		log.Println(err)
		resp.Success = false
		utils.JsonResponse(w, http.StatusInternalServerError, resp)
		return
	}

	id, err := handler.svc.Add(r.Context(), &item.Product)
	if err != nil {
		log.Println(err)
		resp.Success = false
		utils.JsonResponse(w, http.StatusInternalServerError, resp)
		return
	}
	resp.Success = true
	resp.Id = id
	utils.JsonResponse(w, http.StatusOK, resp)
}

func (handler *ScapperHandler) Get(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	resp := status{}
	idv, err := strconv.Atoi(id)

	if err != nil {
		resp.Success = false
		utils.JsonResponse(w, http.StatusInternalServerError, resp)
		return
	}
	item, err := handler.svc.Get(r.Context(), idv)

	if err != nil {
		resp.Success = false
		utils.JsonResponse(w, http.StatusInternalServerError, resp)
		return

	}

	utils.JsonResponse(w, http.StatusOK, item)
}
func (handler *ScapperHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	resp := status{}
	idv, err := strconv.Atoi(id)

	if err != nil {
		resp.Success = false
		utils.JsonResponse(w, http.StatusInternalServerError, resp)
		return
	}
	item := model.Item{}

	err = utils.DecodeRequest(r, &item)
	if err != nil {
		resp.Success = false
		utils.JsonResponse(w, http.StatusInternalServerError, resp)
		return
	}

	err = handler.svc.Update(r.Context(), idv, &item)
	if err != nil {
		resp.Success = false
		utils.JsonResponse(w, http.StatusInternalServerError, resp)
		return
	}

	resp.Success = true
	utils.JsonResponse(w, http.StatusOK, resp)

}
func (handler *ScapperHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	resp := status{}
	idv, err := strconv.Atoi(id)

	if err != nil {
		log.Println(err)
		resp.Success = false
		utils.JsonResponse(w, http.StatusInternalServerError, resp)
		return
	}

	err = handler.svc.Delete(r.Context(), idv)

	if err != nil {
		log.Println(err)
		resp.Success = false
		utils.JsonResponse(w, http.StatusInternalServerError, resp)
		return
	}
	resp.Success = true
	utils.JsonResponse(w, http.StatusOK, resp)

}
