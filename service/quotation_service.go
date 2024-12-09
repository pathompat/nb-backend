package service

import (
	"fmt"
	"notebook-backend/handler/dto"
	"notebook-backend/repository"
	"notebook-backend/repository/model"

	"log/slog"

	"github.com/google/uuid"
)

const (
	Q_DOC_TYPE       string = "QUOTATION"
	Q_STAT_REVIEWING string = "REVIEWING"
	Q_STAT_APPROVED  string = "APPROVED"
	Q_STAT_CANCELED  string = "CANCELED"
	P_DOC_TYPE       string = "PRODUCTION"
	P_STAT_DESIGNING string = "DESIGNING"
	U_ROLE_ADMIN     string = "ADMIN"
	U_ROLE_CUSTOMER  string = "CUSTOMER"
)

type QuotationService interface {
	GetAllQuotation(filter dto.QuotationFilter) ([]dto.QuotationResponse, error)
	GetQuotationByID(quotationID uint) (dto.QuotationResponse, error)
	CountQuotationByStatus(userId uuid.UUID) ([]dto.CountByStatus, error)
	CreateQuotation(input dto.CreateQuotation) (*dto.QuotationResponse, error)
	UpdateQuotation(id uint, input dto.UpdateQuotation) (*dto.QuotationResponse, error)
}

type quotationService struct {
	quotationRepo  repository.QuotationRepository
	userRepo       repository.UserRepository
	schoolRepo     repository.SchoolRepository
	productionRepo repository.ProductionRepository
}

func NewQuotationService(quotationRepo repository.QuotationRepository, userRepo repository.UserRepository, schoolRepo repository.SchoolRepository, productionRepo repository.ProductionRepository) QuotationService {
	return &quotationService{quotationRepo: quotationRepo, userRepo: userRepo, schoolRepo: schoolRepo, productionRepo: productionRepo}
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
		if quotation.Production != nil {
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
		}

		var production *dto.Production
		var productionId *uint
		if filter.IncludeProduction && quotation.Production != nil {
			productionId = &quotation.Production.ID
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
			ProductionID:    productionId,
			Production:      production,
		})

	}
	return quotationMap, nil
}

func (s *quotationService) GetQuotationByID(quotationID uint) (dto.QuotationResponse, error) {
	quotation, err := s.quotationRepo.FindById(quotationID)
	if err != nil {
		return dto.QuotationResponse{}, err
	}

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

	return dto.QuotationResponse{
		ID:              quotation.ID,
		UserID:          quotation.User.UserID,
		Username:        quotation.User.Username,
		StoreName:       quotation.User.StoreName,
		SchoolName:      quotation.SchoolName,
		SchoolAddress:   quotation.SchoolAddress,
		SchoolTelephone: quotation.SchoolTelephone,
		DueDateAt:       quotation.DueDateAt,
		Status:          quotation.Status,
		Items:           quotationItemMap,
		CreatedAt:       quotation.CreatedAt,
		UpdatedAt:       quotation.UpdatedAt,
		ProductionID:    nil,
		Remark:          quotation.Remark,
	}, nil
}

func (s *quotationService) CountQuotationByStatus(userId uuid.UUID) ([]dto.CountByStatus, error) {
	user, err := s.userRepo.FindByID(userId)
	if err != nil {
		return nil, err
	}

	var filterId *uint
	if user.Role == U_ROLE_CUSTOMER {
		filterId = &user.ID
	}

	slog.Info(fmt.Sprintf("Count status: %s, userID: %d", user.Role, filterId))

	statusCount := []dto.CountByStatus{}
	quotationStat, err := s.quotationRepo.CountByStatus(filterId)
	if err != nil {
		return nil, err
	}

	for _, item := range quotationStat {
		statusCount = append(statusCount, dto.CountByStatus{
			Status: item.Status,
			Count:  item.Count,
			Type:   Q_DOC_TYPE,
		})
	}

	productionStat, err := s.productionRepo.CountItemByStatus(filterId)
	if err != nil {
		return nil, err
	}

	for _, item := range productionStat {
		statusCount = append(statusCount, dto.CountByStatus{
			Status: item.Status,
			Count:  item.Count,
			Type:   P_DOC_TYPE,
		})
	}

	return statusCount, nil
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

func (s *quotationService) UpdateQuotation(id uint, input dto.UpdateQuotation) (*dto.QuotationResponse, error) {
	quotation, err := s.quotationRepo.FindById(id)
	if err != nil {
		return nil, err
	}

	quotation.Status = input.Status

	// Update item price, plate on QuotationItems
	productionItem := []model.ProductionItem{}
	for i, request := range input.Items {
		for j, item := range quotation.Items {
			if request.ID == item.ID {
				input.Items[i].Plate, quotation.Items[j].Plate = request.Plate, request.Plate
				input.Items[i].Price, quotation.Items[j].Price = request.Price, request.Price
				productionItem = append(productionItem, model.ProductionItem{
					ProductTitle: item.ProductTitle,
					Plate:        request.Plate,
					Gram:         item.Gram,
					Color:        item.Color,
					Page:         item.Page,
					Pattern:      item.Pattern,
					HasReference: item.HasReference,
					Quantity:     item.Quantity,
					Status:       P_STAT_DESIGNING,
				})
			}
		}
	}

	updatedQuotation, err := s.quotationRepo.Update(*quotation)
	if err != nil {
		return nil, err
	}

	// create production if status = APPROVED
	var createdProduction *model.Production
	if updatedQuotation.Status == Q_STAT_APPROVED {
		production := model.Production{
			UserID:      quotation.User.ID,
			SchoolID:    quotation.SchoolID,
			QuotationID: quotation.ID,
			Remark:      updatedQuotation.Remark,
			Items:       productionItem,
		}
		createdProduction, err = s.productionRepo.Create(production)
		if err != nil {
			return nil, err
		}
	}

	return &dto.QuotationResponse{
		ID:              updatedQuotation.ID,
		ProductionID:    &createdProduction.ID,
		UserID:          quotation.User.UserID,
		StoreName:       updatedQuotation.StoreName,
		SchoolName:      updatedQuotation.SchoolName,
		SchoolAddress:   updatedQuotation.SchoolAddress,
		SchoolTelephone: updatedQuotation.SchoolTelephone,
		AppointmentAt:   updatedQuotation.AppointmentAt,
		DueDateAt:       updatedQuotation.DueDateAt,
		Status:          updatedQuotation.Status,
		Remark:          updatedQuotation.Remark,
		Items:           input.Items,
		CreatedAt:       updatedQuotation.CreatedAt,
		UpdatedAt:       updatedQuotation.UpdatedAt,
	}, nil
}
