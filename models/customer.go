package models

type Customer struct {
	CustomerID int    `json:"customer_id"`
	Name       string `json:"name"`
	Active     int    `json:"active"`
	PriceChild int    `json:"price_child"`
	PriceAdult int    `json:"price_adult"`
}
type MenuItem struct {
	ID              int      `json:"id"`
	CategoryID      *int     `json:"category_id,omitempty"`      // nullable
	OptionIDs       *string  `json:"option_ids,omitempty"`       // nullable longtext
	Name            *string  `json:"name,omitempty"`             // nullable longtext
	Price           *float64 `json:"price,omitempty"`            // nullable decimal
	DiscountedPrice *float64 `json:"discounted_price,omitempty"` // nullable decimal
	IsOrderable     *bool    `json:"is_orderable,omitempty"`     // nullable tinyint(1)
	Status          *bool    `json:"status,omitempty"`           // nullable tinyint(1)
	Description     *string  `json:"description,omitempty"`      // nullable text
}
type MenuCategory struct {
	ID   int               `json:"id"`
	Name map[string]string `json:"name"`
}
