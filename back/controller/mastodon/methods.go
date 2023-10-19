package mastodon

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moonchan-xyz/moonchan/back/mastodon/methods"
)

// https://docs.joinmastodon.org/methods/preferences/
func GetPreferences(c *gin.Context) {
	authorizationID := c.MustGet("authorizationID").(string)
	if authorizationID == "-" {
		errorJSON(c, http.StatusUnauthorized, errUnauthorized)
		return
	}

	id := authorizationID

	resp, err := methods.GetPreferences(id)

	if err != nil {
		errorJSON(c, http.StatusInternalServerError, newErrorMsg(err.Error()))
		return
	}
	c.JSON(http.StatusOK, resp)
}
