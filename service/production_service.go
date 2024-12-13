package service

import (
	"notebook-backend/handler/dto"
	"notebook-backend/repository"
	"time"
)

type ProductionService interface {
	GetProductionByID(productionID uint) (dto.ProductionResponse, error)
	UpdateStatusProductionByID(productionID uint, itemID uint, statusInput dto.UpdateStatusItemProduction) (dto.ProductionItemResponse, error)
}

type productionService struct {
	productionRepo repository.ProductionRepository
}

func NewProductionService(productionRepo repository.ProductionRepository) ProductionService {
	return &productionService{productionRepo: productionRepo}
}

func (s *productionService) GetProductionByID(productionID uint) (dto.ProductionResponse, error) {
	production, err := s.productionRepo.FindProductionByID(productionID)
	if err != nil {
		return dto.ProductionResponse{}, err
	}

	productionItemMap := []dto.ProductionItem{}
	for _, item := range production.Items {
		productionItemMap = append(productionItemMap, dto.ProductionItem{
			Category:     item.Category,
			Plate:        item.Plate,
			Gram:         item.Gram,
			Color:        item.Color,
			Page:         item.Page,
			Pattern:      item.Pattern,
			HasReference: item.HasReference,
			Quantity:     item.Quantity,
			Status:       item.Status,
			CreatedAt:    item.CreatedAt,
			UpdatedAt:    item.UpdatedAt,
		})
	}

	return dto.ProductionResponse{
		ID:              production.ID,
		UserID:          production.User.UserID,
		Username:        production.User.Username,
		StoreName:       production.User.StoreName,
		SchoolName:      production.School.Name,
		SchoolAddress:   production.School.Address,
		SchoolTelephone: production.School.Telephone,
		Remark:          production.Remark,
		Items:           productionItemMap,
	}, nil
}

func (s *productionService) UpdateStatusProductionByID(productionID uint, itemID uint, statusInput dto.UpdateStatusItemProduction) (dto.ProductionItemResponse, error) {
	production, err := s.productionRepo.FindProductionItemByID(productionID, itemID)
	if err != nil {
		return dto.ProductionItemResponse{}, err
	}

	production.Status = statusInput.Status
	production.UpdatedAt = time.Now()

	productionItem, err := s.productionRepo.UpdateStatusItem(production)
	if err != nil {
		return dto.ProductionItemResponse{}, err
	}

	return dto.ProductionItemResponse{
		ID:     productionID,
		ItemID: int(itemID),
		Status: productionItem.Status,
	}, nil
}
