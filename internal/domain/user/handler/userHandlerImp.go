package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/golang-auth-jwtToken/internal/domain/user/dto"
	"github.com/hudayberdipolat/golang-auth-jwtToken/internal/domain/user/service"
	"github.com/hudayberdipolat/golang-auth-jwtToken/internal/utils/response"
	"github.com/hudayberdipolat/golang-auth-jwtToken/internal/utils/validate"
	"net/http"
)

type userHandlerImp struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) UserHandler {
	return userHandlerImp{
		userService: userService,
	}
}

func (u userHandlerImp) GetUser(ctx *fiber.Ctx) error {
	username := ctx.Locals("username").(string)
	getUser, err := u.userService.GetUser(username)
	if err != nil {
		errResponse := response.Error(http.StatusBadRequest, "user not found", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "user data", getUser)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (u userHandlerImp) UpdateUserData(ctx *fiber.Ctx) error {
	var updateUser dto.ChangeUserData
	userID := ctx.Locals("user_id").(int)
	if err := ctx.BodyParser(&updateUser); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "body parser error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	// validate
	if errValidate := validate.ValidateStruct(updateUser); errValidate != nil {
		errResponse := response.Error(http.StatusBadRequest, "validate error", errValidate.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	// update user data
	if err := u.userService.UpdateUserData(userID, updateUser); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "user data can't updated", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, "userData updated successfully", nil)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}

func (u userHandlerImp) UpdateUserPassword(ctx *fiber.Ctx) error {
	var updateUserPassword dto.ChangeUserPassword
	userID := ctx.Locals("user_id").(int)
	if err := ctx.BodyParser(&updateUserPassword); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "body parser error", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	// validate
	if errValidate := validate.ValidateStruct(updateUserPassword); errValidate != nil {
		errResponse := response.Error(http.StatusBadRequest, "validate error", errValidate.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}

	if err := u.userService.UpdateUserPassword(userID, updateUserPassword); err != nil {
		errResponse := response.Error(http.StatusBadRequest, "password can't updated!!!", err.Error(), nil)
		return ctx.Status(http.StatusBadRequest).JSON(errResponse)
	}
	successResponse := response.Success(http.StatusOK, " password changed successfully", nil)
	return ctx.Status(http.StatusOK).JSON(successResponse)
}
