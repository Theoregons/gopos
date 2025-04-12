package controller

import (
	"errors"
	"fmt"
	"net/http"
	"pos/config"
	"pos/entity"
	"pos/middleware"
	"pos/utils"
	"strings"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var input struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
		Role     string `json:"role"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {

		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			e := ve[0]
			field := e.Field()
			tag := e.Tag()

			var msg string

			switch tag {
			case "required":
				msg = fmt.Sprintf("field %s tidak boleh kosong", strings.ToLower(field))
			case "email":
				msg = "format email tidak valid"
			}
			utils.ResponseError(c, http.StatusBadRequest, msg)
			return
		}
		utils.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	role := input.Role
	if role == "" {
		role = "user"
	}

	user := entity.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashed),
		Role:     role,
	}

	if err := config.DB.Create(&user).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, "Email sudah digunakan")
		return
	}
	utils.ResponseSuccess(c, http.StatusCreated, "Register Sukses!")
}

func Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var user entity.User
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		utils.ResponseError(c, http.StatusUnauthorized, "email atau password salah")
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		utils.ResponseError(c, http.StatusUnauthorized, "email atau password salah")
		return
	}
	claims := jwt.MapClaims{
		"user_id": user.Id,
		"role":    user.Role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, _ := token.SignedString(middleware.JWT_SECRET)

	utils.ResponseSuccess(c, http.StatusOK, tokenStr)
}
