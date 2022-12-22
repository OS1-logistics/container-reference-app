package container

import "container/pkg/models/item"

type Container struct {
	ID string 				`json:"id"`
	State string 			`json:"state"`
	IsReuseable bool 	`json:"isReusable"`
	Items []item.Item `json:"items"`
}