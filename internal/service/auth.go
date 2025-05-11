package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"mindlab/internal/model"
	"mindlab/internal/repo"
	"time"
)

var jwtKey = []byte("ochen-bezopasniy-i-ochen-sekretniy-klyuch-pon")

type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

func GenerateJWT(ctx *gin.Context, repo repo.TokenRepository, u *model.User) (string, error) {

	expirationTime := time.Now().Add(15 * time.Minute)
	claims := &Claims{
		Username: u.Username,
		Role:     u.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	err = repo.Save(ctx, tokenString)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(ctx *gin.Context, repo repo.TokenRepository, tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	_, err = repo.Find(ctx, tokenString)
	if err != nil {
		return nil, fmt.Errorf("token revoked or not found")
	}

	return claims, nil
}

func RevokeToken(ctx *gin.Context, repo repo.TokenRepository, tokenString string) error {
	err := repo.Delete(ctx, tokenString)
	if err != nil {
		return err
	}
	return nil
}
