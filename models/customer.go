package models

type Customer struct {
	CustomerID int    `json:"customer_id"`
	Name       string `json:"name"`
	Active     int    `json:"active"`
	PriceChild int    `json:"price_child"`
	PriceAdult int    `json:"price_adult"`
}
