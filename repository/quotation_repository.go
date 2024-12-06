package repository

import (
	"notebook-backend/handler/dto"
	"notebook-backend/repository/model"

	"gorm.io/gorm"
)

type QuotationRepository interface {
	FindAll(filter dto.QuotationFilter) ([]model.Quotation, error)
}

type quotationRepository struct {
	db *gorm.DB
}

func NewQuotationRepository(db *gorm.DB) QuotationRepository {
	return &quotationRepository{db: db}
}

func (r *quotationRepository) FindAll(filter dto.QuotationFilter) ([]model.Quotation, error) {
	var quotations []model.Quotation
	db := r.db.Preload("Items").Preload("User")
	if filter.IncludeProduction {
		db.Preload("Production").Preload("Production.Items")
	}
	if err := db.Find(&quotations).Error; err != nil {
		return nil, err
	}
	return quotations, nil
}
