package controller

import (
	"github.com/gin-gonic/gin"
	"mindlab/internal/api/validator"
	"mindlab/internal/model"
	"mindlab/internal/repo"
	"mindlab/internal/service"
	"net/http"
	"strings"
)

type Credentials struct {
	Username string `json:"username" validate:"required,min=3,max=32"`
	Password string `json:"password" validate:"required,min=8,max=24"`
}

type LoginController struct {
	TokenRepo repo.TokenRepository
}

func NewLoginController(repo repo.TokenRepository) *LoginController {
	return &LoginController{TokenRepo: repo}
}

func (l *LoginController) Login(ctx *gin.Context) {
	cred := Credentials{}
	if err := ctx.ShouldBindJSON(&cred); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validator.Validate.Struct(cred); err != nil {
		errors := validator.PrepareErrorMessage(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errors})
		return
	}

	user := model.ValidateUser(cred.Username, cred.Password)
	if user == nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := service.GenerateJWT(ctx, l.TokenRepo, user)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (l *LoginController) Logout(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	if len(token) == 0 || !strings.HasPrefix(token, "Bearer ") {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Authorization token is required"})
		return
	}

	token = token[7:]

	err := service.RevokeToken(ctx, l.TokenRepo, token)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to log out"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}
