package repository

import (
	// "notebook-backend/handler/dto"
	"notebook-backend/repository/model"

	"gorm.io/gorm"
)

type ProductionRepository interface {
	FindProductionByID(productionID uint) (model.Production, error)
	// Create(quotation model.Quotation) (*model.Quotation, error)
}

type productionRepository struct {
	db *gorm.DB
}

func NewProductionRepository(db *gorm.DB) ProductionRepository {
	return &productionRepository{db: db}
}

func (r *productionRepository) FindProductionByID(productionID uint) (model.Production, error) {
	var production model.Production
	err := r.db.Preload("User").Preload("School").Preload("Items").First(&production, "id = ?", productionID).Error
	if err != nil {
		return model.Production{}, err
	}
	return production, nil
}
