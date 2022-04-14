package app

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sandeepchugh/profileapi/service"
	"net/http"
)

type CustomerHandler struct {
	service service.CustomerService
}

func (ch *CustomerHandler) getAllCustomers(writer http.ResponseWriter, request *http.Request) {
	customers, err := ch.service.GetAllCustomers()

	if err != nil {
		writeResponse(writer, err.Code, err.AsMessage())
	} else {
		writeResponse(writer, http.StatusOK, customers)
	}
}

func (ch *CustomerHandler) getCustomer(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["customer_id"]
	customer, err := ch.service.GetCustomer(id)
	if err != nil {
		writeResponse(writer, err.Code, err.AsMessage())
	} else {
		writeResponse(writer, http.StatusOK, customer)
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		panic(err)
	}
}
