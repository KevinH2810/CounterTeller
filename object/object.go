package object

type (

	//Teller status SLEEP = 0; Occupied = 1
	Teller struct {
		Teller_ID   string `json:"teller_id"`
		Teller_name string `json:"teller_name"`
		Queue       int    `json:"queue"`
		Status      bool   `json:"status"`
	}

	Chair struct {
		Chair_ID  string `json:"chair_id"`
		Available int    `json:"available"`
		Occupied  int    `json:"occupied"`
	}

	Tellers []Teller
	Chairs  []Chair
)

func (t Teller) ID() (jsonField string, value interface{}) {
	value = t.Teller_ID
	jsonField = "teller_id"
	return
}

func (c Chair) ID() (jsonField string, value interface{}) {
	value = c.Chair_ID
	jsonField = "chair_id"
	return
}

var (
	AllTeller Teller
	AllChair  Chair
	TellerArr Tellers
	ChairArr  Chairs
)
