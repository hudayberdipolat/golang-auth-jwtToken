package constructor

import (
	"github.com/hudayberdipolat/golang-auth-jwtToken/internal/domain/auth/handler"
	"github.com/hudayberdipolat/golang-auth-jwtToken/internal/domain/auth/repository"
	"github.com/hudayberdipolat/golang-auth-jwtToken/internal/domain/auth/service"
	"gorm.io/gorm"
)

var authRepo repository.AuthRepository
var authService service.AuthService
var AuthHandler handler.AuthHandler

func AuthRequirementsCreator(db *gorm.DB) {
	authRepo = repository.NewAuthRepository(db)
	authService = service.NewAuthService(authRepo)
	AuthHandler = handler.NewAuthHandler(authService)
}
