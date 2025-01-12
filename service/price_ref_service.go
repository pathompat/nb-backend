package service

import (
	"errors"
	"notebook-backend/handler/dto"
	"notebook-backend/repository"
	"notebook-backend/repository/model"
	"sort"

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
	if err != nil {
		return []dto.PriceRefResponse{}, err
	}

	user, err := s.userRepo.FindByID(parsedUUID)
	if err != nil {
		return []dto.PriceRefResponse{}, err
	}

	priceRefs, err := s.priceRefRepo.FindByTierID(user.TierID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []dto.PriceRefResponse{}, nil
		}
		return nil, err
	}

	// Group data by categoryID
	categoryMap := make(map[uint]*dto.PriceRefResponse)
	for _, priceRef := range priceRefs {
		if _, exists := categoryMap[priceRef.CategoryID]; !exists {
			categoryMap[priceRef.CategoryID] = &dto.PriceRefResponse{
				CategoryID:     priceRef.CategoryID,
				CategoryNameTH: priceRef.Category.NameTH,
				Options:        []dto.Option{},
			}
		}

		// Append options to the category
		categoryMap[priceRef.CategoryID].Options = append(
			categoryMap[priceRef.CategoryID].Options,
			dto.Option{
				Gram:    priceRef.Gram,
				Pattern: []string(priceRef.Pattern),
				Page:    priceRef.Page,
				Color:   priceRef.Color,
				Price:   priceRef.Price,
			},
		)
	}

	// Convert map to slice
	result := []dto.PriceRefResponse{}
	for _, category := range categoryMap {
		result = append(result, *category)
	}

	// Sort result by CategoryID
	sort.Slice(result, func(i, j int) bool {
		return result[i].CategoryID < result[j].CategoryID
	})

	return result, nil
}

func (s *priceRefService) CreatePriceRef(priceRefInput []dto.CreatePriceRef) ([]dto.PriceRefResponse, error) {
	var priceRefs []model.PriceReference
	for _, input := range priceRefInput {
		priceRefs = append(priceRefs, model.PriceReference{
			TierID:     input.TierID,
			CategoryID: input.CategoryID,
			Gram:       input.Gram,
			Color:      input.Color,
			Page:       input.Page,
			Pattern:    input.Pattern,
			Price:      input.Price,
		})
	}

	createdPriceRefs, err := s.priceRefRepo.CreatePriceRef(priceRefs)
	if err != nil {
		return nil, err
	}

	// Group data by categoryID
	categoryMap := make(map[uint]*dto.PriceRefResponse)
	for _, priceRef := range createdPriceRefs {
		if _, exists := categoryMap[priceRef.CategoryID]; !exists {
			categoryMap[priceRef.CategoryID] = &dto.PriceRefResponse{
				CategoryID:     priceRef.CategoryID,
				CategoryNameTH: priceRef.Category.NameTH,
				Options:        []dto.Option{},
			}
		}

		// Append options to the category
		categoryMap[priceRef.CategoryID].Options = append(
			categoryMap[priceRef.CategoryID].Options,
			dto.Option{
				Gram:    priceRef.Gram,
				Pattern: []string(priceRef.Pattern),
				Page:    priceRef.Page,
				Color:   priceRef.Color,
				Price:   priceRef.Price,
			},
		)
	}

	// Convert map to slice
	result := []dto.PriceRefResponse{}
	for _, category := range categoryMap {
		result = append(result, *category)
	}

	// Sort result by CategoryID
	sort.Slice(result, func(i, j int) bool {
		return result[i].CategoryID < result[j].CategoryID
	})

	return result, nil
}
