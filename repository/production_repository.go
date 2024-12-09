package repository

import (
	"errors"
	"notebook-backend/repository/model"

	"gorm.io/gorm"
)

type ProductionRepository interface {
	FindProductionByID(productionID uint) (model.Production, error)
	FindProductionItemByID(productionID uint, itemID uint) (model.ProductionItem, error)
	UpdateStatusItem(productionItem model.ProductionItem) (model.ProductionItem, error)
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

func (r *productionRepository) FindProductionItemByID(productionID uint, itemID uint) (model.ProductionItem, error) {
	var productionItem model.ProductionItem
	err := r.db.Where("production_id = ? AND id = ?", productionID, itemID).First(&productionItem).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.ProductionItem{}, err
		}
		return model.ProductionItem{}, err
	}
	return productionItem, nil
}

func (r *productionRepository) UpdateStatusItem(productionItem model.ProductionItem) (model.ProductionItem, error) {
	err := r.db.Save(&productionItem).Error
	if err != nil {
		return model.ProductionItem{}, err
	}
	return productionItem, nil
}
