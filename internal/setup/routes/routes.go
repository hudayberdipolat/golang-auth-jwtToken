package routes

import (
	"github.com/gofiber/fiber/v2"
	authConstructor "github.com/hudayberdipolat/golang-auth-jwtToken/internal/domain/auth/constructor"
	userConstructor "github.com/hudayberdipolat/golang-auth-jwtToken/internal/domain/user/constructor"
	"github.com/hudayberdipolat/golang-auth-jwtToken/internal/middleware"
)

func Routes(route *fiber.App) {
	apiRoute := route.Group("/api")
	// auth route
	authRoute := apiRoute.Group("auth")
	authRoute.Post("/register", authConstructor.AuthHandler.Register)
	authRoute.Post("/login", authConstructor.AuthHandler.Login)

	// user profile data
	userRoute := apiRoute.Group("user")
	userRoute.Use(middleware.AuthMiddleware)
	userRoute.Get("/user-profile", userConstructor.UserHandler.GetUser)
	userRoute.Put("/update-profile", userConstructor.UserHandler.UpdateUserData)
	userRoute.Put("/update-password", userConstructor.UserHandler.UpdateUserPassword)
}
