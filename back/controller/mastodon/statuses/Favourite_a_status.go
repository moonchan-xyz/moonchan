package statuses
import (
  "net/http"
  "github.com/gin-gonic/gin"
)

// Favourite_a_status godoc
//	@Summary		Favourite a status
//	@Description	Add a status to your favourites list.
//	@Tags			mastodon,statuses
//	@Produce		json
//	@Param			id				path	string	true	"REQUIRED String. The ID of the Status in the database."
//	@Param			Authorization	header	string	true	"REQUIRED Provide this header with Bearer <user token> to gain authorized access to this API method."
//	@Success		200				object	entities.Status
//	@Router			/api/v1/statuses/:id/favourite [post]
func Favourite_a_status(c *gin.Context){
c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})
}