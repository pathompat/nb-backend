package repository

import (
	"notebook-backend/handler/dto"
	"notebook-backend/repository/model"

	"gorm.io/gorm"
)

type QuotationRepository interface {
	FindAll(userID *uint, filter dto.QuotationFilter) ([]model.Quotation, error)
	FindById(id uint) (*model.Quotation, error)
	FindItemByIDAndItemID(id uint, itemID uint) (*model.QuotationItem, error)
	CountByStatus(userID *uint) ([]model.StatusCount, error)
	Create(quotation model.Quotation) (*model.Quotation, error)
	Update(quotation model.Quotation) (*model.Quotation, error)
	UpdateItem(item model.QuotationItem) (*model.QuotationItem, error)
}

type quotationRepository struct {
	db *gorm.DB
}

func NewQuotationRepository(db *gorm.DB) QuotationRepository {
	return &quotationRepository{db: db}
}

func (r *quotationRepository) FindAll(userID *uint, filter dto.QuotationFilter) ([]model.Quotation, error) {
	var quotations []model.Quotation
	db := r.db.Preload("Items").Preload("User")

	if filter.IncludeProduction {
		db = db.Preload("Production", func(db *gorm.DB) *gorm.DB {
			return db.Preload("Items", func(db *gorm.DB) *gorm.DB {
				return db.Where("status != ?", "DONE")
			})
		})
	}

	if userID != nil {
		db.Where("user_id = ?", userID)
	}

	db = db.Where("status != ?", "CANCELED")

	if err := db.Find(&quotations).Error; err != nil {
		return nil, err
	}
	return quotations, nil
}

func (r *quotationRepository) FindById(id uint) (*model.Quotation, error) {
	var quotation model.Quotation
	db := r.db.Preload("Items").Preload("User")
	if err := db.Where("id = ?", id).First(&quotation).Error; err != nil {
		return nil, err
	}
	return &quotation, nil
}

func (r *quotationRepository) FindItemByIDAndItemID(id uint, itemID uint) (*model.QuotationItem, error) {
	var item model.QuotationItem
	err := r.db.Where("quotation_id = ?", id).Where("id = ?", itemID).First(&item).Error
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *quotationRepository) CountByStatus(userID *uint) ([]model.StatusCount, error) {
	var results []model.StatusCount

	db := r.db.Model(&model.Quotation{}).Select("status", "COUNT(*) as count")
	if userID != nil {
		db.Where("user_id = ?", userID)
	}

	if err := db.Group("status").Scan(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

func (r *quotationRepository) Create(quotation model.Quotation) (*model.Quotation, error) {
	if err := r.db.Create(&quotation).Error; err != nil {
		return nil, err
	}
	return &quotation, nil
}

func (r *quotationRepository) Update(quotation model.Quotation) (*model.Quotation, error) {
	if err := r.db.Save(&quotation).Error; err != nil {
		return nil, err
	}
	return &quotation, nil
}

func (r *quotationRepository) UpdateItem(item model.QuotationItem) (*model.QuotationItem, error) {
	err := r.db.Save(&item).Error
	if err != nil {
		return nil, err
	}
	return &item, nil
}
