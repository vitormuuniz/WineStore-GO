package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/vitormuuniz/winestore-go/api/models"
	"github.com/vitormuuniz/winestore-go/api/repositories"
	"github.com/vitormuuniz/winestore-go/api/utils"
)

type WineStoreController interface {
	Create(w http.ResponseWriter, r *http.Request)
	FindById(w http.ResponseWriter, r *http.Request)
	FindAll(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type wineStoreControllerImpl struct {
	wsRepository repositories.WineStoreRepository
}

func NewWineStoreController(ws repositories.WineStoreRepository) *wineStoreControllerImpl {
	return &wineStoreControllerImpl{ws}
}

func (c *wineStoreControllerImpl) Create(w http.ResponseWriter, r *http.Request) {
	if r.Body != nil {
		defer r.Body.Close()
	}

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	wineStore := &models.WineStore{}
	err = json.Unmarshal(bytes, wineStore)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	err = utils.ValidateFields(wineStore, c.wsRepository)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	err = utils.ValidateFieldsRange(wineStore)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	wineStore, err = c.wsRepository.Save(wineStore)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	buildCreatedResponse(w, buildLocation(r, wineStore.ID))
	utils.WriteAsJson(w, wineStore)
}

func (c *wineStoreControllerImpl) FindById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	wineStore_id, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	category, err := c.wsRepository.FindById(wineStore_id)
	if err != nil {
		utils.WriteError(w, err, http.StatusNotFound)
		return
	}

	utils.WriteAsJson(w, category)
}

func (c *wineStoreControllerImpl) FindAll(w http.ResponseWriter, r *http.Request) {
	categories, err := c.wsRepository.FindAll()
	if err != nil {
		utils.WriteError(w, err, http.StatusNotFound)
		return
	}

	utils.WriteAsJson(w, categories)
}

func (c *wineStoreControllerImpl) Update(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	if r.Body != nil {
		defer r.Body.Close()
	}

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	wineStore := &models.WineStore{}
	err = json.Unmarshal(bytes, wineStore)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	wineStore.ID = id

	wineStoreDB, err := c.wsRepository.FindById(wineStore.ID)
	if err != nil {
		utils.WriteError(w, err, http.StatusNotFound)
		return
	}

	wineStore.FillFieldsBeforeUpdate(wineStoreDB)

	err = utils.ValidateFields(wineStore, c.wsRepository)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	err = utils.ValidateFieldsRange(wineStore)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	err = c.wsRepository.Update(wineStore)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	utils.WriteAsJson(w, map[string]bool{"success": err == nil})
}

func (c *wineStoreControllerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	wineStore_id, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	_, err = c.wsRepository.FindById(wineStore_id)
	if err != nil {
		utils.WriteError(w, err, http.StatusNotFound)
		return
	}

	err = c.wsRepository.Delete(wineStore_id)
	if err != nil {
		utils.WriteError(w, err, http.StatusBadRequest)
		return
	}

	buildDeleteResponse(w, wineStore_id)
	utils.WriteAsJson(w, "{}")
}
