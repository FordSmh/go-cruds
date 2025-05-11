package middleware

import (
	"github.com/gin-gonic/gin"
	"mindlab/internal/repo"
	"mindlab/internal/service"
	"net/http"
)

func AuthRequired(tokenRepo repo.TokenRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("Authorization")
		if len(token) == 0 || token[:7] != "Bearer " {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is required"})
			ctx.Abort()
			return
		}

		token = token[7:]

		claims, err := service.ValidateToken(ctx, tokenRepo, token)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			ctx.Abort()
			return
		}

		ctx.Set("username", claims.Username)

		// todo миддлвари теперь неявно связаны, быстро, но неаккуратно
		ctx.Set("role", claims.Role)

		ctx.Next()
	}
}
