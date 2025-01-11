package dto

type CreatePriceRef struct {
	TierID     int      `json:"tierId" example:"1"`              // TierID
	CategoryID uint     `json:"category" example:"cut8"`         // Category
	Gram       int      `json:"gram" example:"12"`               // Gram
	Color      *string  `json:"color" example:"1"`               // Color
	Page       int      `json:"page" example:"30"`               // Page
	Pattern    []string `json:"pattern" example:"[TABLE, cut9]"` // Pattern
	Price      float64  `json:"priceRef" example:"5.5"`          // Price
}

type PriceRefResponse struct {
	CategoryID     uint     `json:"category_id" example:"1"`       // Category id
	CategoryNameTH string   `json:"categoryName" example:"รายงาน"` // Category name
	Options        []Option `json:"options"`                       // List of options
}

type Option struct {
	Gram    int      `json:"gram" example:"55"`                       // Gram
	Pattern []string `json:"pattern" example:"[SINGLE, HALF, TABLE]"` // Pattern
	Page    int      `json:"page" example:"30"`                       // Page
	Color   *string  `json:"color" example:"1"`                       // Color
	Price   float64  `json:"price" example:"3.5"`                     // Price
}
