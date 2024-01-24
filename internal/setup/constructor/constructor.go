package constructor

import (
	"github.com/hudayberdipolat/golang-auth-jwtToken/internal/app"
	authConstructor "github.com/hudayberdipolat/golang-auth-jwtToken/internal/domain/auth/constructor"
	userConstructor "github.com/hudayberdipolat/golang-auth-jwtToken/internal/domain/user/constructor"
)

func Build(dependencies *app.Dependencies) {
	authConstructor.AuthRequirementsCreator(dependencies.DB)
	userConstructor.UserRequirementsCreator(dependencies.DB)
}
