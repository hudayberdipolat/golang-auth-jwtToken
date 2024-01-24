package constructor

import (
	"github.com/hudayberdipolat/golang-auth-jwtToken/internal/domain/user/handler"
	"github.com/hudayberdipolat/golang-auth-jwtToken/internal/domain/user/repository"
	"github.com/hudayberdipolat/golang-auth-jwtToken/internal/domain/user/service"
	"gorm.io/gorm"
)

var userRepository repository.UserRepository
var userService service.UserService
var UserHandler handler.UserHandler

func UserRequirementsCreator(db *gorm.DB) {
	userRepository = repository.NewUserRepository(db)
	userService = service.NewUserService(userRepository)
	UserHandler = handler.NewUserHandler(userService)
}
