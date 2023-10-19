package mastodon

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthChecker() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization") // TODO
		authorizationID := AuthToID(authorization)
		c.Set("authorizationID", authorizationID)
		c.Next()
	}
}

func AuthToID(auth string) string {
	id := strings.TrimPrefix(auth, "Barer ")
	return id
}
