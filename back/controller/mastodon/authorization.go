package mastodon

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func authChecker() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization") // TODO
		authorizationID := authToID(authorization)
		c.Set("authorizationID", authorizationID)
		c.Next()
	}
}

func authToID(auth string) string {
	id := strings.TrimPrefix(auth, "Barer ")
	return id
}
