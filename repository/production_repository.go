package repository

import (
	"errors"
	"notebook-backend/repository/model"

	"gorm.io/gorm"
)

type ProductionRepository interface {
	FindProductionByID(productionID uint) (model.Production, error)
	Create(production model.Production) (*model.Production, error)
	FindProductionItemByID(productionID uint, itemID uint) (model.ProductionItem, error)
	CountItemByStatus(userID *uint) ([]model.StatusCount, error)
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

func (r *productionRepository) Create(production model.Production) (*model.Production, error) {
	if err := r.db.Create(&production).Error; err != nil {
		return nil, err
	}
	return &production, nil
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

func (r *productionRepository) CountItemByStatus(userID *uint) ([]model.StatusCount, error) {
	var results []model.StatusCount
	db := r.db.Model(&model.ProductionItem{}).
		Select("status", "COUNT(*) as count")

	if userID != nil {
		db.Joins("JOIN productions ON productions.id = production_items.production_id").
			Where("user_id = ?", userID)
	}
	if err := db.Group("status").Scan(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

func (r *productionRepository) UpdateStatusItem(productionItem model.ProductionItem) (model.ProductionItem, error) {
	err := r.db.Save(&productionItem).Error
	if err != nil {
		return model.ProductionItem{}, err
	}
	return productionItem, nil
}
