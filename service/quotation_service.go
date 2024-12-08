package service

import (
	"notebook-backend/handler/dto"
	"notebook-backend/repository"
	"notebook-backend/repository/model"
)

const (
	Q_STAT_REVIEWING string = "REVIEWING"
	Q_STAT_APPROVED  string = "APPROVED"
	Q_STAT_CANCELED  string = "CANCELED"
)

type QuotationService interface {
	GetAllQuotation(filter dto.QuotationFilter) ([]dto.QuotationResponse, error)
	CreateQuotation(input dto.CreateQuotation) (*dto.QuotationResponse, error)
}

type quotationService struct {
	quotationRepo repository.QuotationRepository
	userRepo      repository.UserRepository
	schoolRepo    repository.SchoolRepository
}

func NewQuotationService(quotationRepo repository.QuotationRepository, userRepo repository.UserRepository, schoolRepo repository.SchoolRepository) QuotationService {
	return &quotationService{quotationRepo: quotationRepo, userRepo: userRepo, schoolRepo: schoolRepo}
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
				HasReference: &item.HasReference,
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

		var production *dto.Production
		if filter.IncludeProduction {
			production = &dto.Production{
				ID:     quotation.Production.ID,
				Remark: quotation.Production.Remark,
				Items:  productionItemMap,
			}
		}
		quotationMap = append(quotationMap, dto.QuotationResponse{
			ID:              quotation.ID,
			UserID:          quotation.User.UserID,
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

func (s *quotationService) CreateQuotation(input dto.CreateQuotation) (*dto.QuotationResponse, error) {
	user, err := s.userRepo.FindByID(input.UserID)
	if err != nil {
		return nil, err
	}

	school, err := s.schoolRepo.FindById(input.SchoolID)
	if err != nil {
		return nil, err
	}

	items := []model.QuotationItem{}
	for _, item := range input.Items {
		items = append(items, model.QuotationItem{
			ProductTitle: item.ProductTitle,
			Plate:        item.Plate,
			Gram:         item.Gram,
			Color:        item.Color,
			Page:         item.Page,
			Pattern:      item.Pattern,
			HasReference: *item.HasReference,
			Quantity:     item.Quantity,
			Price:        item.Price,
		})
	}

	quotationMap := model.Quotation{
		UserID:          user.ID,
		StoreName:       user.StoreName,
		SchoolID:        school.ID,
		SchoolName:      school.Name,
		SchoolAddress:   school.Address,
		SchoolTelephone: school.Telephone,
		AppointmentAt:   input.AppointmentAt,
		DueDateAt:       input.DueDateAt,
		Status:          Q_STAT_REVIEWING,
		Remark:          input.Remark,
		Items:           items,
	}

	createdQuotation, err := s.quotationRepo.Create(quotationMap)
	if err != nil {
		return nil, err
	}

	return &dto.QuotationResponse{
		UserID:          user.UserID,
		StoreName:       user.StoreName,
		SchoolName:      school.Name,
		SchoolAddress:   school.Address,
		SchoolTelephone: school.Telephone,
		AppointmentAt:   createdQuotation.AppointmentAt,
		DueDateAt:       createdQuotation.DueDateAt,
		Status:          createdQuotation.Status,
		Remark:          createdQuotation.Remark,
		Items:           input.Items,
		CreatedAt:       createdQuotation.CreatedAt,
		UpdatedAt:       createdQuotation.UpdatedAt,
	}, nil
}
