package service

import (
	"errors"
	"notebook-backend/domain/dto"
	"notebook-backend/repository"
)

type UserService interface {
	GetAllUsers() ([]dto.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) GetAllUsers() ([]dto.User, error) {
	resp, err := s.repo.FindAll()
	if err != nil {
		return nil, errors.New("database error")
	}

	users := []dto.User{}
	for _, user := range resp {
		users = append(users, dto.User{
			ID:        int(user.ID),
			Username:  user.Username,
			StoreName: user.StoreName,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}
	return users, nil
}
