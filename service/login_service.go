package service

import (
	"errors"
	"log/slog"
	"notebook-backend/handler/dto"
	"notebook-backend/helper"
	"notebook-backend/repository"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

type LoginService interface {
	Login(credential dto.Login) (*dto.ResponseWithToken, error)
}

type loginService struct {
	userRepo repository.UserRepository
}

func NewLoginService(userRepo repository.UserRepository) LoginService {
	return &loginService{userRepo: userRepo}
}

func (s *loginService) Login(credential dto.Login) (*dto.ResponseWithToken, error) {
	user, err := s.userRepo.FindByUsername(credential.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, helper.ErrInvalidUserOrPwd
		}
		return nil, helper.ErrSelectRecord
	}

	if !helper.CheckPasswordHash(credential.Password, user.Password) {
		slog.Error("check password hash error")
		return nil, helper.ErrUnauthorized
	}

	expiredIn, err := strconv.Atoi(os.Getenv("JWT_EXPIRED_IN"))
	if err != nil {
		slog.Error(err.Error())
		return nil, helper.ErrUnauthorized
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":   user.UserId,
		"username": user.Username,
		"role":     user.Role,
		"exp":      time.Now().Add(time.Second * time.Duration(expiredIn)).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		slog.Error(err.Error())
		return nil, helper.ErrUnauthorized
	}

	return &dto.ResponseWithToken{Token: tokenString, ExpiredIn: expiredIn}, nil
}
