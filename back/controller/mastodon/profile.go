package mastodon

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moonchan-xyz/moonchan/back/mastodon/profile"
)

// https://docs.joinmastodon.org/methods/profile/#delete-profile-avatar
func DeleteProfileAvatar(c *gin.Context) {
	authorizationID := c.MustGet("authorizationID").(string)
	if authorizationID == "-" {
		errorJSON(c, http.StatusUnauthorized, errUnauthorized)
		return
	}

	id := authorizationID

	resp, err := profile.DeleteAvatar(id)

	if err != nil {
		errorJSON(c, http.StatusInternalServerError, newErrorMsg(err.Error()))
		return
	}
	c.JSON(http.StatusOK, resp)
}

// https://docs.joinmastodon.org/methods/profile/#delete-profile-header
func DeleteProfileHeader(c *gin.Context) {
	authorizationID := c.MustGet("authorizationID").(string)
	if authorizationID == "-" {
		errorJSON(c, http.StatusUnauthorized, errUnauthorized)
		return
	}

	id := authorizationID

	resp, err := profile.DeleteHeader(id)

	if err != nil {
		errorJSON(c, http.StatusInternalServerError, newErrorMsg(err.Error()))
		return
	}
	c.JSON(http.StatusOK, resp)
}
