package service

import (
	"notebook-backend/handler/dto"
	"notebook-backend/repository"
)

type QuotationService interface {
	GetAllQuotation(filter dto.QuotationFilter) ([]dto.QuotationResponse, error)
}

type quotationService struct {
	quotationRepo repository.QuotationRepository
}

func NewQuotationService(quotationRepo repository.QuotationRepository) QuotationService {
	return &quotationService{quotationRepo: quotationRepo}
}

func (s *quotationService) GetAllQuotation(filter dto.QuotationFilter) ([]dto.QuotationResponse, error) {
	quotations, err := s.quotationRepo.FindAll(filter)
	if err != nil {
		return nil, err
	}

	quotationMap := []dto.QuotationResponse{}
	for _, quotation := range quotations {
		quotationItemMap := []dto.QuotationItem{}
		for _, item := range quotation.Items {
			quotationItemMap = append(quotationItemMap, dto.QuotationItem{
				ProductTitle: item.ProductTitle,
				Plate:        item.Plate,
				Gram:         item.Gram,
				Color:        item.Color,
				Page:         item.Page,
				Pattern:      item.Pattern,
				HasReference: item.HasReference,
				Quantity:     item.Quantity,
				Price:        item.Price,
			})
		}
		productionItemMap := []dto.ProductionItem{}
		for _, item := range quotation.Production.Items {
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

		var production *dto.ProductionResponse
		if filter.IncludeProduction {
			production = &dto.ProductionResponse{
				ID:     quotation.Production.ID,
				Remark: quotation.Production.Remark,
				Items:  productionItemMap,
			}
		}
		quotationMap = append(quotationMap, dto.QuotationResponse{
			ID:              quotation.ID,
			StoreName:       quotation.StoreName,
			SchoolName:      quotation.SchoolName,
			SchoolAddress:   quotation.SchoolAddress,
			SchoolTelephone: quotation.SchoolTelephone,
			AppointmentAt:   quotation.AppointmentAt,
			DueDateAt:       quotation.DueDateAt,
			Status:          quotation.Status,
			Items:           quotationItemMap,
			CreatedAt:       quotation.CreatedAt,
			UpdatedAt:       quotation.UpdatedAt,
			Production:      production,
		})

	}
	return quotationMap, nil
}
