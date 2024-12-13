package service

import (
	"errors"
	"notebook-backend/handler/dto"
	"notebook-backend/repository"
	"notebook-backend/repository/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PriceRefService interface {
	GetPriceRefByUserID(userID string) ([]dto.PriceRefResponse, error)
	CreatePriceRef(priceRefInput []dto.CreatePriceRef) ([]dto.PriceRefResponse, error)
}

type priceRefService struct {
	priceRefRepo repository.PriceRefRepository
	userRepo     repository.UserRepository
}

func NewPriceRefService(priceRefRepo repository.PriceRefRepository, userRepo repository.UserRepository) PriceRefService {
	return &priceRefService{priceRefRepo: priceRefRepo, userRepo: userRepo}
}

func (s *priceRefService) GetPriceRefByUserID(userID string) ([]dto.PriceRefResponse, error) {
	parsedUUID, err := uuid.Parse(userID)

	user, err := s.userRepo.FindByID(parsedUUID)
	if err != nil {
		return []dto.PriceRefResponse{}, errors.New("User not found")
	}

	priceRefs, err := s.priceRefRepo.FindByTierID(user.TierID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []dto.PriceRefResponse{}, nil
		}
		return nil, errors.New("database error")
	}

	priceRefMap := []dto.PriceRefResponse{}
	for _, priceRef := range priceRefs {
		priceRefMap = append(priceRefMap, dto.PriceRefResponse{
			Category:     priceRef.Category,
			Plate:        priceRef.Plate,
			Gram:         priceRef.Gram,
			Color:        priceRef.Color,
			Page:         priceRef.Page,
			Pattern:      priceRef.Pattern,
			HasReference: priceRef.HasReference,
			Price:        priceRef.Price,
		})
	}
	return priceRefMap, nil
}

func (s *priceRefService) CreatePriceRef(priceRefInput []dto.CreatePriceRef) ([]dto.PriceRefResponse, error) {
	var priceRefs []model.PriceReference
	for _, input := range priceRefInput {
		priceRefs = append(priceRefs, model.PriceReference{
			TierID:       input.TierID,
			Category:     input.Category,
			Plate:        input.Plate,
			Gram:         input.Gram,
			Color:        input.Color,
			Page:         input.Page,
			Pattern:      input.Pattern,
			HasReference: input.HasReference,
			Price:        input.Price,
		})
	}

	createdPriceRefs, err := s.priceRefRepo.CreatePriceRef(priceRefs)
	if err != nil {
		return nil, err
	}

	var responsePriceRefs []dto.PriceRefResponse
	for _, ref := range createdPriceRefs {
		responsePriceRefs = append(responsePriceRefs, dto.PriceRefResponse{
			Category:     ref.Category,
			Plate:        ref.Plate,
			Gram:         ref.Gram,
			Color:        ref.Color,
			Page:         ref.Page,
			Pattern:      ref.Pattern,
			HasReference: ref.HasReference,
			Price:        ref.Price,
		})
	}

	return responsePriceRefs, nil
}
