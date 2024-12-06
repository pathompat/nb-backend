package dto

type CreatePriceRef struct {
	TierID       int     `json:"tierId" example:"1"`           // TierID
	ProductTitle string  `json:"productTitle" example:"cut8"`  // ProductTitle
	Plate        string  `json:"plate" example:"LARGE"`        // Plate
	Gram         int     `json:"gram" example:"12"`            // Gram
	Color        string  `json:"color" example:"1"`            // Color
	Page         int     `json:"page" example:"30"`            // Page
	Pattern      string  `json:"pattern" example:"TABLE"`      // Pattern
	HasReference bool    `json:"hasReference" example:"false"` // HasReference
	Price        float64 `json:"priceRef" example:"5.5"`       // Price
}

type PriceRefResponse struct {
	ProductTitle string  `json:"productTitle" example:"cut8"`  // ProductTitle
	Plate        string  `json:"plate" example:"LARGE"`        // Plate
	Gram         int     `json:"gram" example:"12"`            // Gram
	Color        string  `json:"color" example:"1"`            // Color
	Page         int     `json:"page" example:"30"`            // Page
	Pattern      string  `json:"pattern" example:"TABLE"`      // Pattern
	HasReference bool    `json:"hasReference" example:"false"` // HasReference
	Price        float64 `json:"priceRef" example:"5.5"`       // Price
}
