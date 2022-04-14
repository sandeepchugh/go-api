package app

import (
	"github.com/gorilla/mux"
	"github.com/sandeepchugh/profileapi/domain"
	"github.com/sandeepchugh/profileapi/service"
	"log"
	"net/http"
)

func Start() {

	router := mux.NewRouter()

	// wiring
	ch := CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	// routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	// start server
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
