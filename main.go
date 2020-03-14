package main

import (
	"CounterTeller/Controller"
	"CounterTeller/object"

	"github.com/gorilla/mux"
	"github.com/simdb/db"
)

func initialize() {
	driver, err := db.New("db")
	object.AllTeller = object.Teller{

		Teller_ID:   "1",
		Teller_name: "Silvia Romanov",
		Queue:       0,
		Status:      false,
	}

	object.AllChair = object.Chair{
		Chair_ID:  "1",
		Available: 20,
		Occupied:  0,
	}

	err = driver.Insert(object.AllTeller)
	if err != nil {
		panic(err)
	}

	err = driver.Insert(object.AllChair)
	if err != nil {
		panic(err)
	}
}

//	func (c Customer) ID() (jsonField string, value interface{}) {
//		value=c.CustID
//		jsonField="custid"
//		return
//	}

//BASED ON THE SCENARIO, TELLER WILL BE ONLY GIVEN OPTION TO UPDATE THEIR STATUS OCCUPIED/NOT
//BECAUSE OF THAT THE POST, GET AND DELETE ROUTE WILL NOT BE INLCUDED
//AND THE DEFAULT CHAIR WILL BE 20 FOR THE SAKE OF SIMPLICITY
//THE GIVEN INPUT WILL BE AMOUNT OF CHAIRS THAT THE CUSTOMER ABLE TO SIT AT
func main() {
	initialize()
	router := mux.NewRouter().StrictSlash(true)
	Controller.Routers(router)
}
