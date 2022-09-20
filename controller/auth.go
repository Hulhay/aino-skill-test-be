package controller

import (
	"fmt"
	"net/http"

	"github.com/Hulhay/aino-skill-test-be/model"
	"github.com/Hulhay/aino-skill-test-be/shared"
	"github.com/Hulhay/aino-skill-test-be/usecase"
	"github.com/Hulhay/aino-skill-test-be/usecase/auth"
	"github.com/Hulhay/aino-skill-test-be/usecase/token"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type authController struct {
	authUC  usecase.Auth
	tokenUC usecase.Token
}

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
	Logout(ctx *gin.Context)
}

func NewAuthController(authUC usecase.Auth, tokenUC usecase.Token) AuthController {
	return &authController{
		authUC:  authUC,
		tokenUC: tokenUC,
	}
}

func (c *authController) Register(ctx *gin.Context) {

	var (
		register auth.RegisterRequest
		err      error
	)

	err = ctx.ShouldBind(&register)
	if err != nil {
		res := shared.BuildErrorResponse("Failed to process request", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	err = c.authUC.Register(ctx, register)
	if err != nil {
		res := shared.BuildErrorResponse("Register Failed!", err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := shared.BuildResponse("Register Success!", nil)
	ctx.JSON(http.StatusCreated, res)
}

func (c *authController) Login(ctx *gin.Context) {

	var (
		response *token.ResultResponse
		login    auth.LoginRequest
		user     *model.User
		err      error
	)

	err = ctx.ShouldBind(&login)
	if err != nil {
		res := shared.BuildErrorResponse("Failed to process request", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	user, err = c.authUC.Login(ctx, login)
	if err != nil {
		res := shared.BuildErrorResponse("Login Failed!", err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	response, _ = c.tokenUC.GenerateToken(ctx, user)

	res := shared.BuildResponse("Login Success!", response)
	ctx.JSON(http.StatusOK, res)
}

func (c *authController) Logout(ctx *gin.Context) {

	authHeader := ctx.GetHeader("Authorization")
	token, err := c.tokenUC.ValidateToken(authHeader)
	if err != nil {
		return
	}
	claims := token.Claims.(jwt.MapClaims)
	email := fmt.Sprintf("%v", claims["email"])
	err = c.authUC.Logout(ctx, email)
	if err != nil {
		res := shared.BuildErrorResponse("Login Failed!", err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := shared.BuildResponse("Logout Success!", nil)
	ctx.JSON(http.StatusOK, res)
}
