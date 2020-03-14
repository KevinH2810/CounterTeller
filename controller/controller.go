package Controller

import (
	"encoding/json"
	"net/http"

	. "CounterTeller/object"

	"github.com/simdb/db"
)

type (
	TellerObj Teller
	ChairObj  Chair
)

//getAmountOfOccupiedChairs ARE USED TO GET HOW MUCH CHAIRS ARE OCCUPIED
func getAmountOfOccupiedChairs(w http.ResponseWriter, r *http.Request) {
	driver, err := db.New("db")
	err = driver.Open(Chair{}).Where("chair_id", "=", "1").First().AsEntity(&AllChair)
	if err != nil {
		panic(err)
	}
	js, err := json.Marshal(&AllChair.Occupied)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

//getQueueNumber ARE USED TO GET THE QUEUE NUMBER SO
func getQueueNumber(w http.ResponseWriter, r *http.Request) {
	driver, err := db.New("db")
	err = driver.Open(Teller{}).Where("teller_id", "=", "1").First().AsEntity(&AllTeller)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	js, err := json.Marshal(&AllTeller.Queue)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

//THE TELLER WILL BE ABLE TO CLICK NEXT QUEUE TO GET ANY CUSTOMER IF THERE'S ANY AVAILABLE CUSTOMER ON THE CHAIR
//OTHERWISE THE CUSTOMER WILL GO SLEEP
func updateTellerNextQueue(w http.ResponseWriter, r *http.Request) {
	driver, err := db.New("db")
	//GET THE CHAIR'S DATA
	err = driver.Open(Chair{}).Where("Teller_ID", "=", "1").First().AsEntity(&AllChair)
	err = driver.Open(Teller{}).Where("Teller_ID", "=", "1").First().AsEntity(&AllTeller)
	if AllChair.Occupied == 0 {
		AllTeller.Status = false
		w.Write([]byte("NO MORE CUSTOMER, TELLER IS PUT TO SLEEP"))
	} else {
		AllChair.Occupied = AllChair.Occupied - 1
		AllChair.Available = AllChair.Available + 1
		AllTeller.Queue = AllTeller.Queue + 1
		w.Write([]byte("NEXT CUSTOMER PLEASE COME TO THE DESIGNATED COUNTER"))
	}
	err = driver.Update(AllTeller)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = driver.Update(AllChair)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

//Called when a new Customer enter the bank, it will add the occupied chairs by 1 if there is available chair
//OTHERWISE RETURN ERROR NO MORE AVAILABLE CHAIR
func updateChairNewCustomer(w http.ResponseWriter, r *http.Request) {
	driver, err := db.New("db")
	//GET THE CHAIR'S DATA
	err = driver.Open(Chair{}).Where("Teller_ID", "=", "1").First().AsEntity(&AllChair)
	err = driver.Open(Teller{}).Where("Teller_ID", "=", "1").First().AsEntity(&AllTeller)
	//CHECK FOR ANY AVAILABLE CHAIR LEFT
	if AllChair.Available == 0 {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - No More Empty Chair!"))
	}

	//CHECK IF TELLER IS ON SLEEP MODE OR NOT
	if AllTeller.Status == false {
		AllTeller.Status = true
		AllTeller.Queue = AllTeller.Queue + 1
		w.Write([]byte("CUSTOMER CAN CONTINUE TO DESIGNATED TELLER"))
	} else {
		AllChair.Occupied = AllChair.Occupied + 1
		AllChair.Available = AllChair.Available - 1
		w.Write([]byte("PLEASE WAIT FOR YOUR NUMBER TO BE CALLED"))
	}
	err = driver.Update(AllTeller)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = driver.Update(AllChair)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
