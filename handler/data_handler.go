package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yenchu/go-demo/model"
	"github.com/yenchu/go-demo/repository"
)

func NewDataHandler() *DataHandler {

	repo := repository.NewDataRepository()

	return &DataHandler{
		repo: repo,
	}
}

type DataHandler struct {
	repo *repository.DataRepository
}

func (handler *DataHandler) FindAllData(w http.ResponseWriter, req *http.Request) {

	data, err := handler.repo.FindAllData()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(&data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintf(w, string(res))
}

func (handler *DataHandler) CreateData(w http.ResponseWriter, req *http.Request) {

	body, err := ioutil.ReadAll(req.Body)

	defer req.Body.Close()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	datum := &model.Data{}

	err = json.Unmarshal(body, datum)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = handler.repo.CreateData(datum)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
