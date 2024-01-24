package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/golang-auth-jwtToken/internal/domain/auth/dto"
	"github.com/hudayberdipolat/golang-auth-jwtToken/internal/domain/auth/service"
	"github.com/hudayberdipolat/golang-auth-jwtToken/internal/utils/response"
	"github.com/hudayberdipolat/golang-auth-jwtToken/internal/utils/validate"
	"net/http"
)

type authHandlerImp struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) AuthHandler {
	return authHandlerImp{
		authService: authService,
	}
}

func (a authHandlerImp) Register(ctx *fiber.Ctx) error {
	var registerRequest dto.RegisterRequest

	if err := ctx.BodyParser(&registerRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "body parser error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	// validate
	if errValidate := validate.ValidateStruct(registerRequest); errValidate != nil {
		errResponse := response.Error(http.StatusBadRequest, "validate error", errValidate.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	// register user
	registerResponse, err := a.authService.Register(registerRequest)
	if err != nil {
		errResponse := response.Error(http.StatusBadRequest, "user can't registered", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusCreated, "user registered successfully", registerResponse)
	return ctx.Status(http.StatusCreated).JSON(successResponse)
}

func (a authHandlerImp) Login(ctx *fiber.Ctx) error {
	var loginRequest dto.LoginRequest
	if err := ctx.BodyParser(&loginRequest); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "body parser error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	// validate
	if errValidate := validate.ValidateStruct(loginRequest); errValidate != nil {
		errResponse := response.Error(http.StatusBadRequest, "validate error", errValidate.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	// login user
	loginResponse, err := a.authService.Login(loginRequest)
	if err != nil {
		errResponse := response.Error(http.StatusBadRequest, "user can't registered", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "user login successfully", loginResponse)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}
