package app

import (
	"encoding/json"
	"encoding/xml"
	"github.com/sandeepchugh/profileapi/service"
	"net/http"
)

type Customer struct {
	Name    string `json: "name"`
	City    string `json: "city"`
	ZipCode string `json: "zipCode"`
}

type CustomerHandler struct {
	service service.CustomerService
}

func (ch *CustomerHandler) getAllCustomers(writer http.ResponseWriter, request *http.Request) {
	//customers := []Customer{
	//	{"Ashish", "New Delhi", "110075"},
	//	{"Rob", "New Delhi", "110075"},
	//}

	customers, _ := ch.service.GetAllCustomers()

	if request.Header.Get("Content-Type") == "application/xml" {
		writer.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(writer).Encode(customers)
	} else {
		writer.Header().Add("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(customers)
	}
}
