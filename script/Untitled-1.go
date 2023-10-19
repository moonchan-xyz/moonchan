
// Mute godoc
//
//	@Summary		Lookup account ID from Webfinger address
//	@Description	Quickly lookup a username to see if it is available, skipping WebFinger resolution.
//	@Description	https://docs.joinmastodon.org/methods/accounts/#unblock
//	@Tags			mastodon/accounts
//	@Produce		json
//	@Param			Authorization	header		string	true	"Provide this header with `Bearer <user token>` to gain authorized access to this API method."
//	@Param			id				path 		string	true	"The ID of the Account in the database."
// @Param notifications formData boolean false "Mute notifications in addition to statuses? Defaults to true."
// @Param duration formData number false "How long the mute should last, in seconds. Defaults to 0 (indefinite)."
//	@Success		200		{object}	entities.Relationship
//	@Router			/api/v1/accounts/:id/mute [post]
func Mute(c *gin.Context) {
	authorizationID := c.MustGet("authorizationID").(string)
	if authorizationID == "-" {
		errorJSON(c, http.StatusUnauthorized, errUnauthorized)
		return
	}
	id := c.Param("id")

	resp, err := accounts.Unblock(authorizationID, id)
	if err != nil {
		errorJSON(c, http.StatusNotFound, errNotFound)
		return
	}
	c.JSON(http.StatusOK, resp)
}
