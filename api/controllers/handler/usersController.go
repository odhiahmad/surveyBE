package handler

import (
	"survey/api/controllers"
	"survey/api/middlewares"
	"survey/api/models/dto"
	"survey/api/usecase"
	"survey/api/utils/httpParse"
	"survey/api/utils/httpResponse"
	"survey/api/utils/status"
	"encoding/json"
	_ "errors"
	"io/ioutil"
	"net/http"
	_ "strconv"

	"survey/api/models"
	"survey/api/utils/formaterror"

	"github.com/gorilla/mux"
)

type UserController struct {
	router    *mux.Router
	parseJson *httpParse.JsonParse
	responder httpResponse.IResponder
	service   usecase.IUserUseCase
}

func NewUserController(router *mux.Router, parse *httpParse.JsonParse, responder httpResponse.IResponder, service usecase.IUserUseCase) controllers.IDelivery {
	return &UserController{
		router, parse, responder, service,
	}
}

func (s *UserController) InitRoute(mdw ...mux.MiddlewareFunc) {
	u := s.router.PathPrefix("/users").Subrouter()
	u.Use(mdw...)
	u.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")
	u.HandleFunc("/create", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	u.HandleFunc("/edit", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.EditUser))).Methods("PUT")

}



func (u *UserController) Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		u.responder.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	var user dto.Login
	err = json.Unmarshal(body, &user)
	if err != nil {
		u.responder.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	token, err := u.service.(&user)
	if err != nil {
		u.responder.Error(w, http.StatusNotFound, err.Error())
		return
	}
	u.responder.Data(w, http.StatusOK, status.StatusText(status.Success), token)
}

func (u *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		u.responder.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	var user models.UserRequest
	err = json.Unmarshal(body, &user)
	if err != nil {
		u.responder.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	err = user.Validate("")
	if err != nil {
		u.responder.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	userCreated, _ := u.service.SaveUser(&user)

	if err != nil {

		formattedError := formaterror.FormatError(err.Error())

		u.responder.Error(w, http.StatusInternalServerError, formattedError.Error())
		return
	}
	u.responder.Data(w, http.StatusCreated, status.StatusText(status.CREATED), userCreated)
}

func (u *UserController) EditUser(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		u.responder.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	var user models.User

	err = json.Unmarshal(body, &user)
	if err != nil {
		u.responder.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	userEdit, err := u.service.UpdateInfo(&user)
	if err != nil {
		u.responder.Error(w, http.StatusNotFound, err.Error())
		return
	}

	u.responder.Data(w, http.StatusCreated, status.StatusText(status.CREATED), userEdit)
}
