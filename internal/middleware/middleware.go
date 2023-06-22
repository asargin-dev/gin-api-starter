package middleware

import (
	"github.com/asargin-dev/soundproof-backend-demo/pkg/tokenizer"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, "/user") {
			err := tokenizer.TokenValid(c)
			if err != nil {
				c.String(http.StatusUnauthorized, "Unauthorized")
				c.Abort()
				return
			}
			c.Next()
		}

	}
}
