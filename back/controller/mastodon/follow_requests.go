package mastodon

import (
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/moonchan-xyz/moonchan/back/mastodon/methods"
)

func GetFollowRequests(c *gin.Context) {
	authorizationID := c.MustGet("authorizationID").(string)
	if authorizationID == "-" {
		errorJSON(c, http.StatusUnauthorized, errUnauthorized)
		return
	}
	id := authorizationID
	maxID := c.Query("max_id")
	sinceID := c.Query("since_id")
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		limit, err = strconv.Atoi(os.Getenv("DEFAULT_ACCOUNT_QUERY_LIMIT"))
		if err != nil {
			errorJSON(c, http.StatusInternalServerError, newErrorMsg(err.Error()))
			return
		}
	}

	resp, err := methods.GetFollowRequests(id, maxID, sinceID, limit)

	if err != nil {
		errorJSON(c, http.StatusInternalServerError, newErrorMsg(err.Error()))
		return
	}
	c.JSON(http.StatusOK, resp)
}

func AcceptFollowRequest(c *gin.Context) {
	authorizationID := c.MustGet("authorizationID").(string)
	if authorizationID == "-" {
		errorJSON(c, http.StatusUnauthorized, errUnauthorized)
		return
	}
	id := authorizationID
	accountID := c.Param("account_id")

	resp, err := methods.AcceptFollowRequest(id, accountID)

	if err != nil {
		errorJSON(c, http.StatusInternalServerError, newErrorMsg(err.Error()))
		return
	}
	c.JSON(http.StatusOK, resp)
}

func RejectFollowRequest(c *gin.Context) {
	authorizationID := c.MustGet("authorizationID").(string)
	if authorizationID == "-" {
		errorJSON(c, http.StatusUnauthorized, errUnauthorized)
		return
	}
	id := authorizationID
	accountID := c.Param("account_id")

	resp, err := methods.RejectFollowRequest(id, accountID)

	if err != nil {
		errorJSON(c, http.StatusInternalServerError, newErrorMsg(err.Error()))
		return
	}
	c.JSON(http.StatusOK, resp)
}
