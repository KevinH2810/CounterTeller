package Controller

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Routers(router *mux.Router) {
	router.HandleFunc("/Teller", updateTellerNextQueue).Methods("PATCH")
	router.HandleFunc("/Chair", updateChairNewCustomer).Methods("PATCH")
	router.HandleFunc("/Chair", getAmountOfOccupiedChairs).Methods("GET")
	router.HandleFunc("/Queue", getQueueNumber).Methods("GET")

	http.ListenAndServe(":9090", router)
}
