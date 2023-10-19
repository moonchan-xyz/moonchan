package mastodon

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moonchan-xyz/moonchan/back/mastodon/statuses"
)

func CreateStatus(c *gin.Context) {
	authorizationID := c.MustGet("authorizationID").(string)
	if authorizationID == "-" {
		errorJSON(c, http.StatusUnauthorized, errUnauthorized)
		return
	}
	id := authorizationID
	key := c.GetHeader("Idempotency-Key")

	form, err := c.MultipartForm() // todo should check type
	if err != nil {
		errorJSON(c, http.StatusBadRequest, newErrorMsg(err.Error()))
	}
	fe := formExactor{Form: form}
	status := defaultString(fe.String("status"), "")
	mediaIDs := fe.StringArray("media_ids[]")
	pollOptions := fe.StringArray("poll[options][]")
	pollExpiresIn := mustAtoi(defaultString(fe.String("poll[expires_in]"), "0"), 0)
	pollMutiple := fe.Bool("poll[multiple]")
	pollHideTotals := fe.Bool("poll[hide_totals]]")
	inReplyToID := fe.String("in_reply_to_id")
	sensitive := fe.Bool("sensitive")
	spoilerText := fe.String("spoiler_text")
	visibility := fe.String("visibility")
	language := fe.String("language")
	scheduledAt := fe.String("scheduled_at")
	resp, err := statuses.Create(
		id,
		key,
		status,
		mediaIDs,
		pollOptions,
		pollExpiresIn,
		pollMutiple,
		pollHideTotals,
		inReplyToID,
		sensitive,
		spoilerText,
		visibility,
		language,
		scheduledAt,
	)

	if err != nil {
		errorJSON(c, http.StatusInternalServerError, newErrorMsg(err.Error()))
		return
	}
	c.JSON(http.StatusOK, resp)
}
