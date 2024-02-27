package follow_requests
import (
  "net/http"
  "github.com/gin-gonic/gin"
)

// Accept_follow_request godoc
//	@Summary	Accept follow request
//	@Description
//	@Tags		mastodon,follow_requests
//	@Produce	json
//	@Param		account_id		path	string	true	"REQUIRED String. The ID of the Account in the database."
//	@Param		Authorization	header	string	true	"REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
//	@Success	200				object	entities.Relationship
//	@Router		/api/v1/follow_requests/:account_id/authorize [post]
func Accept_follow_request(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}