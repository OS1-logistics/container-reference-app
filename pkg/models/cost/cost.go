package cost

type Unit struct {
	Amount int 			`json:"amount"`
	Currency string `json:"currency"`
}

type Cost struct {
	Unit Unit 	`json:"unit"`
	Total Unit 	`json:"total"`
}