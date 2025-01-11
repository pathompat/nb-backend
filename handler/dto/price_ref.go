package dto

type CreatePriceRef struct {
	TierID       int     `json:"tierId" example:"1"`           // TierID
	CategoryID   uint    `json:"category" example:"cut8"`      // Category
	Plate        string  `json:"plate" example:"LARGE"`        // Plate
	Gram         int     `json:"gram" example:"12"`            // Gram
	Color        string  `json:"color" example:"1"`            // Color
	Page         int     `json:"page" example:"30"`            // Page
	Pattern      string  `json:"pattern" example:"TABLE"`      // Pattern
	HasReference bool    `json:"hasReference" example:"false"` // HasReference
	Price        float64 `json:"priceRef" example:"5.5"`       // Price
}

type PriceRefResponse struct {
	CategoryID   uint    `json:"category" example:"cut8"`      // Category
	Plate        string  `json:"plate" example:"LARGE"`        // Plate
	Gram         int     `json:"gram" example:"12"`            // Gram
	Color        string  `json:"color" example:"1"`            // Color
	Page         int     `json:"page" example:"30"`            // Page
	Pattern      string  `json:"pattern" example:"TABLE"`      // Pattern
	HasReference bool    `json:"hasReference" example:"false"` // HasReference
	Price        float64 `json:"priceRef" example:"5.5"`       // Price
}
