package routes

import (
	"github.com/gofiber/fiber/v2"
	authConstructor "github.com/hudayberdipolat/golang-auth-jwtToken/internal/domain/auth/constructor"
)

func Routes(route *fiber.App) {
	apiRoute := route.Group("/api")
	// auth route
	authRoute := apiRoute.Group("auth")
	authRoute.Post("/register", authConstructor.AuthHandler.Register)
	authRoute.Post("/login", authConstructor.AuthHandler.Login)
}
