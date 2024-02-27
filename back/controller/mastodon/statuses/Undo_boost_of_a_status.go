package statuses
import (
  "net/http"
  "github.com/gin-gonic/gin"
)

// Undo_boost_of_a_status godoc
//	@Summary		Undo boost of a status
//	@Description	Undo a reshare of a status.
//	@Tags			mastodon,statuses
//	@Produce		json
//	@Param			id				path	string	true	"REQUIRED String. The ID of the Status in the database."
//	@Param			Authorization	header	string	true	"REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
//	@Success		200				object	entities.Status
//	@Router			/api/v1/statuses/:id/unreblog [post]
func Undo_boost_of_a_status(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}