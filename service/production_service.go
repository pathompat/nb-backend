package service

import (
	"notebook-backend/handler/dto"
	"notebook-backend/repository"
)

type ProductionService interface {
	GetProductionByID(productionID uint) (dto.ProductionResponse, error)
	// CreateQuotation(input dto.CreateQuotation) (*dto.QuotationResponse, error)
}

type productionService struct {
	productionRepo repository.ProductionRepository
	userRepo       repository.UserRepository
	schoolRepo     repository.SchoolRepository
}

func NewProductionService(productionRepo repository.ProductionRepository, userRepo repository.UserRepository, schoolRepo repository.SchoolRepository) ProductionService {
	return &productionService{productionRepo: productionRepo, userRepo: userRepo, schoolRepo: schoolRepo}
}

func (s *productionService) GetProductionByID(productionID uint) (dto.ProductionResponse, error) {
	production, err := s.productionRepo.FindProductionByID(productionID)
	if err != nil {
		return dto.ProductionResponse{}, err
	}

	productionItemMap := []dto.ProductionItem{}
	for _, item := range production.Items {
		productionItemMap = append(productionItemMap, dto.ProductionItem{
			ProductTitle: item.ProductTitle,
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
		UserName:        production.User.Username,
		StoreName:       production.User.StoreName,
		SchoolName:      production.School.Name,
		SchoolAddress:   production.School.Address,
		SchoolTelephone: production.School.Telephone,
		Remark:          production.Remark,
		Items:           productionItemMap,
	}, nil
}
