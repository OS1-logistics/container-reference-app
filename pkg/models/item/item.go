package item

import "container/pkg/models/cost"

type Item struct {
	Name     string 		`json:"name"`
	Code     string 		`json:"code"`
	Cost     cost.Cost	`json:"cost"`
	Quantity int    		`json:"quantity"`
}