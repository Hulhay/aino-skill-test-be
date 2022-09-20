package main

import (
	"github.com/Hulhay/aino-skill-test-be/config"
	"github.com/Hulhay/aino-skill-test-be/controller"
	"github.com/Hulhay/aino-skill-test-be/middleware"
	"github.com/Hulhay/aino-skill-test-be/repository"
	"github.com/Hulhay/aino-skill-test-be/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db      *gorm.DB                  = config.SetupDatabaseConnection()
	repo    repository.UserRepository = repository.NewUserRepository(db)
	tokenUC usecase.Token             = usecase.NewTokenUc()
	authUC  usecase.Auth              = usecase.NewAuthUC(repo)
	ac      controller.AuthController = controller.NewAuthController(authUC, tokenUC)
)

func main() {

	r := gin.Default()

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/register", ac.Register)
		authRoutes.POST("/login", ac.Login)
		authRoutes.POST("/logout", ac.Logout, middleware.AuthorizeJWT(tokenUC))
	}

	r.Run()
}
