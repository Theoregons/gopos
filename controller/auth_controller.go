package controller

import (
	"errors"
	"fmt"
	"net/http"
	"pos/service"
	"pos/utils"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Register(c *gin.Context) {
	var input struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
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
				msg = fmt.Sprintf("field %s cannot be empty", strings.ToLower(field))
			case "email":
				msg = "format must be email"
			case "min":
				msg = "minimum password 6 characters"
			}

			utils.ResponseError(c, http.StatusBadRequest, msg)
			return
		}
		utils.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	_, err := service.RegisterUser(input.Name, input.Email, input.Password, input.Role)

	if err != nil {
		utils.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.ResponseSuccess(c, http.StatusCreated, nil, "Register Sukses!")
}

func Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := service.Login(input.Email, input.Password)
	if err != nil {
		utils.ResponseError(c, http.StatusUnauthorized, err.Error())
		return
	}

	utils.ResponseSuccess(c, http.StatusOK, token, "login success")
}
