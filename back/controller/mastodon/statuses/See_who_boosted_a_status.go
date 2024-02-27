package statuses
import (
  "net/http"
  "github.com/gin-gonic/gin"
)

// See_who_boosted_a_status godoc
//	@Summary		See who boosted a status
//	@Description	View who boosted a given status.
//	@Tags			mastodon,statuses
//	@Produce		json
//	@Param			id				path	string	true	"REQUIRED String. The ID of the Status in the database."
//	@Param			Authorization	header	string	false	"Provide this header with Bearer <user token> to gain authorized access to this API method."
//	@Param			max_id			query	string	false	"Internal parameter. Use HTTP Link header for pagination."
//	@Param			since_id		query	string	false	"Internal parameter. Use HTTP Link header for pagination."
//	@Param			limit			query	int		false	"Integer. Maximum number of results to return. Defaults to 40 accounts. Max 80 accounts."
//	@Success		200				array	entities.Account
//	@Router			/api/v1/statuses/:id/reblogged_by [get]
func See_who_boosted_a_status(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}