package handler

import (
	"encoding/json"
	_ "errors"
	"io/ioutil"
	"net/http"
	_ "strconv"
	"survey/api/controllers"
	"survey/api/middlewares"
	"survey/api/models/dto"
	"survey/api/usecase"
	"survey/api/utils/httpParse"
	"survey/api/utils/httpResponse"
	"survey/api/utils/status"

	"survey/api/models"
	"survey/api/utils/formaterror"

	"github.com/gorilla/mux"
)

type RumahController struct {
	router    *mux.Router
	parseJson *httpParse.JsonParse
	responder httpResponse.IResponder
	service   usecase.IRumahUseCase
}

func NewRumahController(router *mux.Router, parse *httpParse.JsonParse, responder httpResponse.IResponder, service usecase.IRumahUseCase) controllers.IDelivery {
	return &RumahController{
		router, parse, responder, service,
	}
}

func (s *RumahController) InitRoute(mdw ...mux.MiddlewareFunc) {
	r := s.router.PathPrefix("/rumah").Subrouter()
	r.Use(mdw...)
	r.HandleFunc("/get", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.GetAllRumah))).Methods("GET")
	r.HandleFunc("/create", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.CreateRumah))).Methods("POST")
	r.HandleFunc("/edit", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.EditRumah))).Methods("PUT")

}

func (ru *RumahController) GetAllRumah(w http.ResponseWriter, r *http.Request) {
	page := r.Header.Get("page")
	size := r.Header.Get("size")
	order := r.Header.Get("order")
	wallets, err := ru.service.GetAllRumah(page, size, order)
	if err != nil {
		ru.responder.Error(w, http.StatusNotFound, err.Error())
		return
	}
	ru.responder.Data(w, status.Success, status.StatusText(http.StatusOK), wallets)
}

func (ru *RumahController) CreateRumah(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ru.responder.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	var rumah dto.RumahRequest
	err = json.Unmarshal(body, &rumah)
	if err != nil {
		ru.responder.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	userCreated, _ := ru.service.SaveRumah(rumah)

	if err != nil {

		formattedError := formaterror.FormatError(err.Error())

		ru.responder.Error(w, http.StatusInternalServerError, formattedError.Error())
		return
	}
	ru.responder.Data(w, http.StatusCreated, status.StatusText(status.CREATED), userCreated)
}

func (ru *RumahController) EditRumah(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ru.responder.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	var rumah models.Rumah

	err = json.Unmarshal(body, &rumah)
	if err != nil {
		ru.responder.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	userEdit, err := ru.service.UpdateInfo(&rumah)
	if err != nil {
		ru.responder.Error(w, http.StatusNotFound, err.Error())
		return
	}

	ru.responder.Data(w, http.StatusCreated, status.StatusText(status.CREATED), userEdit)
}
