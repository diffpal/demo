package orders

import "time"

type Order struct {
	ID             string    `json:"id"`
	UserID         string    `json:"user_id"`
	ProductID      string    `json:"product_id"`
	Quantity       int       `json:"quantity"`
	UnitPriceCents int       `json:"unit_price_cents"`
	TotalCents     int       `json:"total_cents"`
	CreatedAt      time.Time `json:"created_at"`
}

type CreateInput struct {
	UserID         string
	ProductID      string
	Quantity       int
	UnitPriceCents int
}
