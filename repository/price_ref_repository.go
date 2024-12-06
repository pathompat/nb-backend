package repository

import (
	"notebook-backend/repository/model"

	"gorm.io/gorm"
)

type PriceRefRepository interface {
	FindByTierID(TierID int) ([]model.PriceReference, error)
	CreatePriceRef(priceRefs []model.PriceReference) ([]model.PriceReference, error)
}

type priceRefRepository struct {
	db *gorm.DB
}

func NewPriceRefRepository(db *gorm.DB) PriceRefRepository {
	return &priceRefRepository{db: db}
}

func (r *priceRefRepository) FindByTierID(TierID int) ([]model.PriceReference, error) {
	var priceRef []model.PriceReference
	err := r.db.Unscoped().Where("tier_id = ?", TierID).Find(&priceRef).Error
	if err != nil {
		return []model.PriceReference{}, err
	}
	return priceRef, nil
}

func (r *priceRefRepository) CreatePriceRef(priceRefs []model.PriceReference) ([]model.PriceReference, error) {
	if err := r.db.Create(&priceRefs).Error; err != nil {
		return priceRefs, err
	}
	return priceRefs, nil
}
