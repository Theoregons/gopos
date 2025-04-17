package service

import (
	"errors"
	"pos/config"
	"pos/entity"
	"pos/middleware"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(name, email, password, role string) (entity.User, error) {
	var existing entity.User

	if err := config.DB.Where("email = ?", email).First(&existing).Error; err == nil {
		return entity.User{}, errors.New("email already in use")
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if role == "" {
		role = "user"
	}

	user := entity.User{
		Name:     name,
		Email:    email,
		Password: string(hashed),
		Role:     role,
	}

	if err := config.DB.Create(&user).Error; err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func Login(email, password string) (string, error) {
	var user entity.User
	if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return "", errors.New("email atau password salah")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("email atau password salah")
	}
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(middleware.JWT_SECRET)

	if err != nil {
		return "", err
	}

	return tokenStr, nil
}
