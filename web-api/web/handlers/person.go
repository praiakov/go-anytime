package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/praiakov/webapi/core"
)

func MakePersonHandlers(r *mux.Router, n *negroni.Negroni, service core.UseCase) {
	r.Handle("/v1/person", n.With(
		negroni.Wrap(getAllPeople(service))),
	).Methods("GET", "OPTIONS")

	r.Handle("/v1/person/{id}", n.With(
		negroni.Wrap(getPersonById(service)),
	)).Methods("GET", "OPTIONS")

	r.Handle("/v1/person", n.With(
		negroni.Wrap(createPerson(service)),
	)).Methods("POST", "OPTIONS")

	r.Handle("/v1/person/{id}", n.With(
		negroni.Wrap(updatePerson(service)),
	)).Methods("PUT", "OPTIONS")

	r.Handle("/v1/person/{id}", n.With(
		negroni.Wrap(removePerson(service)),
	)).Methods("DELETE", "OPTIONS")
}

// @Summary Get details of all people
// @Tags people
// @Accept  json
// @Produce  json
// @Success 200 {array} core.Person
// @Router /person [get]
func getAllPeople(service core.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		all, err := service.GetPeople()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(formatJSONError(err.Error()))
			return
		}
		err = json.NewEncoder(w).Encode(all)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(formatJSONError("Erro convertendo em JSON"))
			return
		}
	})
}

// @Summary Get person by id
// @Tags people
// @Accept  json
// @Produce  json
// @Param   id     path    string     true		"person id"
// @Success 200 {array} core.Person
// @Router /person/{id} [get]
func getPersonById(service core.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		vars := mux.Vars(r)
		id, err := strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(formatJSONError(err.Error()))
			return
		}
		p, err := service.GetPerson(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write(formatJSONError(err.Error()))
			return
		}

		err = json.NewEncoder(w).Encode(p)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(formatJSONError("Erro convertendo em JSON"))
			return
		}
	})
}

// @Summary Create person
// @Tags people
// @Accept json
// @Produce json
// @Param message body core.Person true "Person Info"
// @Success 200 {string} string "OK"
// @Router /person [post]
func createPerson(service core.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var p core.Person
		var a core.Address
		p.Address = &a

		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(formatJSONError(err.Error()))
			return
		}

		err = service.CreatePerson(&p)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(formatJSONError(err.Error()))
			return
		}
		w.WriteHeader(http.StatusCreated)
	})
}

// @Summary Delete person
// @Tags people
// @Accept json
// @Produce json
// @Param   id     path    string     true		"person id"
// @Success 200  {string} string "OK"
// @Router /person/{id} [delete]
func removePerson(service core.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		vars := mux.Vars(r)

		id, err := strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(formatJSONError(err.Error()))
			return
		}

		err = service.DeletePerson(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(formatJSONError(err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
	})
}

// @Summary Update a person
// @Tags people
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"OK"
// @Param   id     path    string     true		"person id"
// @Router /person/{id} [put]
func updatePerson(service core.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var p core.Person
		var a core.Address
		p.Address = &a
		err := json.NewDecoder(r.Body).Decode(&p)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(formatJSONError(err.Error()))
			return
		}
		vars := mux.Vars(r)
		id, err := strconv.ParseInt(vars["id"], 10, 64)
		p.ID = strconv.FormatInt(id, 10)

		err = service.UpdatePerson(&p)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(formatJSONError(err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
	})
}
