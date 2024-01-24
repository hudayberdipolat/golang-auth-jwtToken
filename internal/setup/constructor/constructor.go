package constructor

import (
	"github.com/hudayberdipolat/golang-auth-jwtToken/internal/app"
	authConstructor "github.com/hudayberdipolat/golang-auth-jwtToken/internal/domain/auth/constructor"
)

func Build(dependencies *app.Dependencies) {
	authConstructor.AuthRequirementsCreator(dependencies.DB)
}
