package domain

import (
	"time"
)

type Domain struct {
	Id      int    `json:"id"`
	Status  string `json:"status"`
	Grade   string `json:"D_Free"`
	GroupId string `json:"group_id"`
	Name    string `json:"name"`
}
type DomainList struct {
	Status struct {
		Code      string    `json:"code"`
		Message   string    `json:"message"`
		CreatedAt time.Time `json:"message"`
	} `json:"status"`

	Domains []Domain `json:"domain"`
}
