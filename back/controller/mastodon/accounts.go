// https://docs.joinmastodon.org/methods/accounts

package mastodon

import (
	"errors"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/moonchan-xyz/moonchan/back/mastodon/accounts"
	"gorm.io/gorm"
)

// Create godoc
//
//	@Summary		Register an account
//	@Description	Creates a user and account records. Returns an account access token for the app that initiated the request. The app should save this token for later, and should wait for the user to confirm their account by clicking a link in their email inbox.
//	@Description	https://docs.joinmastodon.org/methods/accounts/#create
//	@Tags			mastodon/accounts
//	@Produce		json
//	@Param			Authorization	header		string	true	"Provide this header with `Bearer <user token>` to gain authorized access to this API method."
//	@Param			username		formData	string	true	"The desired username for the account"
//	@Param			email			formData	string	true	"The email address to be used for login"
//	@Param			password		formData	string	true	"The password to be used for login"
//	@Param			aggrement		formData	boolean	true	" Whether the user agrees to the local rules, terms, and policies. These should be presented to the user in order to allow them to consent before setting this parameter to TRUE."
//	@Param			locale			formData	string	true	"The language of the confirmation email that will be sent."
//	@Param			reason			formData	string	false	"If registrations require manual approval, this text will be reviewed by moderators."
//	@Param			id				path		string	true	"The ID of the Account in the database."
//	@Success		200				{object}	entities.Account
//	@Failure		404				{object}	errorMsg
//	@Router			/api/v1/accounts [post]
func Create(c *gin.Context) {
	authorizationID := c.MustGet("authorizationID").(string)
	if authorizationID != "-" {
		errorJSON(c, http.StatusUnauthorized, errUnauthorized)
		return
	}
	username := c.PostForm("username")
	email := c.PostForm("email")
	password := c.PostForm("password")
	agreement := isTrue(c.PostForm("agreement"))
	locale := c.PostForm("locale")
	reason := c.PostForm("reason")
	err := accounts.Create(
		username,
		email,
		password,
		agreement,
		locale,
		reason,
	)
	if err != nil {
		errorJSON(c, http.StatusUnprocessableEntity, errUnprocessableEntity)
		return
	}
	c.JSON(http.StatusOK, nil)
}

// VerifyCredentials godoc
//
//	@Summary		Verify account credentials
//	@Description	Test to make sure that the user token works.
//	@Description	https://docs.joinmastodon.org/methods/accounts/#verify_credentials
//	@Tags			mastodon/accounts
//	@Produce		json
//	@Param			Authorization	header		string	true	"Provide this header with `Bearer <user token>` to gain authorized access to this API method."
//	@Success		200				{object}	entities.CredentialAccount
//	@Failure		401				{object}	errorMsg
//	@Failure		403				{object}	errorMsg
//	@Failure		404				{object}	errorMsg
//	@Failure		422				{object}	errorMsg
//	@Router			/api/v1/accounts/verify_credentials [get]
func VerifyCredentials(c *gin.Context) {
	authorizationID := c.MustGet("authorizationID").(string)
	if authorizationID == "-" {
		errorJSON(c, http.StatusUnauthorized, errUnauthorized)
		return
	}
	account, err := accounts.GetCredentials(authorizationID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			errorJSON(c, http.StatusNotFound, errNotFound)
		} else {
			errorJSON(c, http.StatusInternalServerError, newErrorMsg(err.Error()))
		}
		return
	}
	c.JSON(http.StatusOK, account)
}

// UpdateCredentials godoc
//
//	@Summary		Update account credentials
//	@Description	Update the user’s display and preferences.
//	@Description	https://docs.joinmastodon.org/methods/accounts/#update_credentials
//	@Tags			mastodon/accounts
//	@Accept			mpfd
//	@Produce		json
//	@Param			Authorization							header		string	true	"Provide this header with `Bearer <user token>` to gain authorized access to this API method."
//	@param			account[display_name]					formData	string	false	"The display name to use for the profile."
//	@Param			account[note]							formData	string	false	"The account bio."
//	@Param			account[avatar]							formData	file	false	"Avatar image encoded using multipart/form-data"
//	@Param			account[header]							formData	file	false	"Header image encoded using multipart/form-data"
//	@Param			account[locked]							formData	boolean	false	"Whether manual approval of follow requests is required."
//	@Param			account[bot]							formData	boolean	false	"Whether the account has a bot flag."
//	@Param			account[discoverable]					formData	boolean	false	"Whether the account should be shown in the profile directory."
//	@Param			account[fields_attributes][0][name]		formData	string	false	"hash"
//	@Param			account[fields_attributes][0][value]	formData	string	false	"hash"
//	@Param			account[fields_attributes][1][name]		formData	string	false	"hash"
//	@Param			account[fields_attributes][1][value]	formData	string	false	"hash"
//	@Param			account[fields_attributes][2][name]		formData	string	false	"hash"
//	@Param			account[fields_attributes][2][value]	formData	string	false	"hash"
//	@Param			account[fields_attributes][3][name]		formData	string	false	"hash"
//	@Param			account[fields_attributes][3][value]	formData	string	false	"hash"
//	@Param			account[discoverable]					formData	boolean	false	"Whether the account should be shown in the profile directory."
//	@Param			account[discoverable]					formData	boolean	false	"Whether the account should be shown in the profile directory."
//	@Param			account[discoverable]					formData	boolean	false	"Whether the account should be shown in the profile directory."
//	@Success		200										{object}	entities.CredentialAccount
//	@Failure		401										{object}	errorMsg
//	@Failure		403										{object}	errorMsg
//	@Failure		404										{object}	errorMsg
//	@Failure		422										{object}	errorMsg
//	@Failure		500										{object}	errorMsg
//	@Router			/api/v1/accounts/update_credentials [patch]
func UpdateCredentials(c *gin.Context) {
	authorizationID := c.MustGet("authorizationID").(string)
	if authorizationID == "-" {
		errorJSON(c, http.StatusUnauthorized, errUnauthorized)
		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		errorJSON(c, http.StatusBadRequest, newErrorMsg(err.Error()))
	}

	fe := formExactor{Form: form}

	account, err := accounts.UpdateCredentials(
		authorizationID,
		fe.String("account[display_name]"),    //displayName,
		fe.String("account[note]"),            // 	note,
		fe.URLString("account[avatar]"),       // 	avatar, // todo
		fe.URLString("account[header]"),       // 	header, // todo this is file
		fe.Bool("account[locked]"),            // 	locked,
		fe.Bool("account[bot]"),               // 	bot,
		fe.String("account[discoverable]"),    // 	discoverable,
		fe.Hash("account[fields_attributes]"), // 	fieldsAttributes,
		fe.String("account[display_name]"),    // 	sourcePrivacy, // todo, check key name
		fe.Bool("source[sensitive]"),          // 	sourceSensitive, // todo, check key name
		fe.String("source[language]"),         // 	sourceLanguage) // todo, check key name
	)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			errorJSON(c, http.StatusNotFound, errNotFound)
		} else {
			errorJSON(c, http.StatusInternalServerError, newErrorMsg(err.Error()))
		}
		return
	}
	c.JSON(http.StatusOK, account)
}

// GetAccount godoc
//
//	@Summary		Get account
//	@Description	View information about a profile.
//	@Description	https://docs.joinmastodon.org/methods/accounts/#get
//	@Tags			mastodon/accounts
//	@Produce		json
//	@Param			Authorization	header		string	false	"Provide this header with `Bearer <user token>` to gain authorized access to this API method."
//	@Param			id				path		string	true	"The ID of the Account in the database."
//	@Success		200				{object}	entities.Account
//	@Failure		404				{object}	errorMsg
//	@Router			/api/v1/accounts/{id} [get]
func GetAccount(c *gin.Context) {
	authorizationID := c.MustGet("authorizationID").(string)
	if isTrue(os.Getenv("ALLOW_VISITORS_WITHOUT_AUTHORIZATION")) || authorizationID == "-" {
		errorJSON(c, http.StatusUnauthorized, errUnauthorized)
		return
	}
	id := c.Param("id")
	account, err := accounts.Get(id)
	if err != nil {
		errorJSON(c, http.StatusNotFound, errNotFound)
		return
	}
	c.JSON(http.StatusOK, account)
}

// Statuses godoc
//
//	@Summary		Get account’s statuses
//	@Description	Statuses posted to the given account.
//	@Description	https://docs.joinmastodon.org/methods/accounts/#statuses
//	@Tags			mastodon/accounts
//	@Produce		json
//	@Param			Authorization	header		string	false	"Provide this header with `Bearer <user token>` to gain authorized access to this API method."
//	@Param			id				path		string	true	"The ID of the Account in the database."
//	@Param			max_id			query		string	false	"All results returned will be lesser than this ID. In effect, sets an upper bound on results."
//	@Param			since_id		query		string	false	"All results returned will be greater than this ID. In effect, sets a lower bound on results."
//	@Param			min_id			query		string	false	"Returns results immediately newer than this ID. In effect, sets a cursor at this ID and paginates forward."
//	@Param			limit			query		integer	false	"Maximum number of results to return. Defaults to 20 statuses. Max 40 statuses."
//	@Param			only_media		query		boolean	false	"Filter out statuses without attachments."
//	@Param			exclude_replies	query		boolean	false	"Filter out statuses in reply to a different account."
//	@Param			exclude_reblogs	query		boolean	false	"Filter out boosts from the response."
//	@Param			pinned			query		boolean	false	"Filter for pinned statuses only. Defaults to false, which includes all statuses. Pinned statuses do not receive special priority in the order of the returned results."
//	@Param			tagged			query		string	false	"Filter for statuses using a specific hashtag."
//	@Success		200				array		entities.Status
//	@Failure		404				{object}	errorMsg
//	@Router			/api/v1/accounts/{id}/statuses [get]
func Status(c *gin.Context) {
	authorizationID := c.MustGet("authorizationID").(string)
	if isTrue(os.Getenv("ALLOW_VISITORS_WITHOUT_AUTHORIZATION")) || authorizationID == "-" {
		errorJSON(c, http.StatusUnauthorized, errUnauthorized)
		return
	}
	id := c.Param("id")
	maxID := c.Query("max_id")
	sinceID := c.Query("since_id")
	minID := c.Query("min_id")
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		limit, err = strconv.Atoi(os.Getenv("DEFAULT_STATUS_QUERY_LIMIT"))
		if err != nil {
			errorJSON(c, http.StatusInternalServerError, newErrorMsg(err.Error()))
			return
		}
	}
	onlyMedia := isTrue(c.Query("only_media"))
	excludeReplies := isTrue(c.Query("exclude_replies"))
	excludeReblogs := isTrue(c.Query("exclude_reblogs"))
	pinned := isTrue(c.Query("pinned"))
	tagged := c.Query("tagged")
	statuses, err := accounts.Status(
		id,
		maxID,
		sinceID,
		minID,
		limit,
		onlyMedia,
		excludeReplies,
		excludeReblogs,
		pinned,
		tagged,
	)

	if err != nil {
		errorJSON(c, http.StatusInternalServerError, newErrorMsg(err.Error()))
		return
	}
	c.JSON(http.StatusOK, statuses)
}

// Followers godoc
//
//	@Summary		Lookup account ID from Webfinger address
//	@Description	Quickly lookup a username to see if it is available, skipping WebFinger resolution.
//	@Description	https://docs.joinmastodon.org/methods/accounts/#lookup
//	@Tags			mastodon/accounts
//	@Produce		json
//	@Param			Authorization	header		string	true	"Provide this header with `Bearer <user token>` to gain authorized access to this API method."
//	@Param			id				path		string	true	"The ID of the Account in the database."
//	@Param			max_id			query		string	false	"Internal parameter. Use HTTP Link header for pagination."
//	@Param			since_id		query		string	false	"Internal parameter. Use HTTP Link header for pagination."
//	@Param			min_id			query		string	false	"Internal parameter. Use HTTP Link header for pagination."
//	@Param			limit			query		integer	false	"Maximum number of results to return. Defaults to 40 accounts. Max 80 accounts."
//	@Success		200				array		[]entities.Account
//	@Failure		404				{object}	errorMsg
//	@Router			/api/v1/accounts/{id}/followers  [get]
func Followers(c *gin.Context) {
	authorizationID := c.MustGet("authorizationID").(string)
	if isTrue(os.Getenv("ALLOW_VISITORS_WITHOUT_AUTHORIZATION")) || authorizationID == "-" {
		errorJSON(c, http.StatusUnauthorized, errUnauthorized)
		return
	}
	id := c.Param("id")
	maxID := c.Query("max_id")
	sinceID := c.Query("since_id")
	minID := c.Query("min_id")
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		limit, err = strconv.Atoi(os.Getenv("DEFAULT_ACCOUNT_QUERY_LIMIT"))
		if err != nil {
			errorJSON(c, http.StatusInternalServerError, newErrorMsg(err.Error()))
			return
		}
	}

	followers, err := accounts.Followers(
		id,
		maxID,
		sinceID,
		minID,
		limit,
	)

	if err != nil {
		errorJSON(c, http.StatusInternalServerError, newErrorMsg(err.Error()))
		return
	}
	c.JSON(http.StatusOK, followers)
}

// Following godoc
//
//	@Summary		Lookup account ID from Webfinger address
//	@Description	Quickly lookup a username to see if it is available, skipping WebFinger resolution.
//	@Description	https://docs.joinmastodon.org/methods/accounts/#lookup
//	@Tags			mastodon/accounts
//	@Produce		json
//	@Param			Authorization	header		string	true	"Provide this header with `Bearer <user token>` to gain authorized access to this API method."
//	@Param			id				path		string	true	"The ID of the Account in the database."
//	@Param			max_id			query		string	false	"Internal parameter. Use HTTP Link header for pagination."
//	@Param			since_id		query		string	false	"Internal parameter. Use HTTP Link header for pagination."
//	@Param			min_id			query		string	false	"Internal parameter. Use HTTP Link header for pagination."
//	@Param			limit			query		integer	false	"Maximum number of results to return. Defaults to 40 accounts. Max 80 accounts."
//	@Success		200				array		[]entities.Account
//	@Failure		404				{object}	errorMsg
//	@Router			/api/v1/accounts/{id}/following  [get]
func Following(c *gin.Context) {
	authorizationID := c.MustGet("authorizationID").(string)
	if isTrue(os.Getenv("ALLOW_VISITORS_WITHOUT_AUTHORIZATION")) || authorizationID == "-" {
		errorJSON(c, http.StatusUnauthorized, errUnauthorized)
		return
	}
	id := c.Param("id")
	maxID := c.Query("max_id")
	sinceID := c.Query("since_id")
	minID := c.Query("min_id")
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		limit, err = strconv.Atoi(os.Getenv("DEFAULT_ACCOUNT_QUERY_LIMIT"))
		if err != nil {
			errorJSON(c, http.StatusInternalServerError, newErrorMsg(err.Error()))
			return
		}
	}

	following, err := accounts.Following(
		id,
		maxID,
		sinceID,
		minID,
		limit,
	)

	if err != nil {
		errorJSON(c, http.StatusInternalServerError, newErrorMsg(err.Error()))
		return
	}
	c.JSON(http.StatusOK, following)
}

// todo

// Follow godoc
//
//	@Summary		Lookup account ID from Webfinger address
//	@Description	Quickly lookup a username to see if it is available, skipping WebFinger resolution.
//	@Description	https://docs.joinmastodon.org/methods/accounts/#lookup
//	@Tags			mastodon/accounts
//	@Produce		json
//	@Param			Authorization	header		string			true	"Provide this header with `Bearer <user token>` to gain authorized access to this API method."
//	@Param			id				path		string			true	"The ID of the Account in the database."
//
//	@Param			reblogs			formData	boolean			false	"Receive this account’s reblogs in home timeline? Defaults to true."
//	@Param			notify			formData	boolean			false	"Receive notifications when this account posts a status? Defaults to false."
//	@Param			languages		formData	array{string}	false	"Array of String (ISO 639-1 language two-letter code). Filter received statuses for these languages. If not provided, you will receive this account’s posts in all languages."
//
//	@Success		200				{object}	entities.Relationship
//	@Router			/api/v1/accounts/:id/follow [post]
func Follow(c *gin.Context) {
	authorizationID := c.MustGet("authorizationID").(string)
	if authorizationID == "-" {
		errorJSON(c, http.StatusUnauthorized, errUnauthorized)
		return
	}
	id := c.Param("id")

	form, err := c.MultipartForm() // todo should check type
	if err != nil {
		errorJSON(c, http.StatusBadRequest, newErrorMsg(err.Error()))
	}
	fe := formExactor{Form: form}

	reblogs := isTrue(defaultString(fe.String("reblogs"), "true"))
	notify := isTrue(defaultString(fe.String("notify"), "false"))
	languages := fe.String("languages") // todo check what passed. 不想实现了

	resp, err := accounts.Follow(authorizationID, id, reblogs, notify, languages)
	if err != nil {
		errorJSON(c, http.StatusNotFound, errNotFound)
		return
	}
	c.JSON(http.StatusOK, resp)
}

// Unfollow godoc
//
//	@Summary		Lookup account ID from Webfinger address
//	@Description	Quickly lookup a username to see if it is available, skipping WebFinger resolution.
//	@Description	https://docs.joinmastodon.org/methods/accounts/#unfollow
//	@Tags			mastodon/accounts
//	@Produce		json
//	@Param			Authorization	header		string	true	"Provide this header with `Bearer <user token>` to gain authorized access to this API method."
//	@Param			id				path		string	true	"The ID of the Account in the database."
//	@Success		200				{object}	entities.Relationship
//	@Router			/api/v1/accounts/:id/unfollow [post]
func Unfollow(c *gin.Context) {
	authorizationID := c.MustGet("authorizationID").(string)
	if authorizationID == "-" {
		errorJSON(c, http.StatusUnauthorized, errUnauthorized)
		return
	}
	id := c.Param("id")

	resp, err := accounts.Unfollow(authorizationID, id)
	if err != nil {
		errorJSON(c, http.StatusNotFound, errNotFound)
		return
	}
	c.JSON(http.StatusOK, resp)
}

// RemoveFromFollowers godoc
//
//	@Summary		Lookup account ID from Webfinger address
//	@Description	Quickly lookup a username to see if it is available, skipping WebFinger resolution.
//	@Description	https://docs.joinmastodon.org/methods/accounts/#remove_from_followers
//	@Tags			mastodon/accounts
//	@Produce		json
//	@Param			Authorization	header		string	true	"Provide this header with `Bearer <user token>` to gain authorized access to this API method."
//	@Param			id				path		string	true	"The ID of the Account in the database."
//	@Success		200				{object}	entities.Relationship
//	@Router			/api/v1/accounts/:id/remove_from_followers [post]
func RemoveFromFollowers(c *gin.Context) {
	authorizationID := c.MustGet("authorizationID").(string)
	if authorizationID == "-" {
		errorJSON(c, http.StatusUnauthorized, errUnauthorized)
		return
	}
	id := c.Param("id")

	resp, err := accounts.RemoveFromFollowers(authorizationID, id)
	if err != nil {
		errorJSON(c, http.StatusNotFound, errNotFound)
		return
	}
	c.JSON(http.StatusOK, resp)
}

// Block godoc
//
//	@Summary		Lookup account ID from Webfinger address
//	@Description	Quickly lookup a username to see if it is available, skipping WebFinger resolution.
//	@Description	https://docs.joinmastodon.org/methods/accounts/#block
//	@Tags			mastodon/accounts
//	@Produce		json
//	@Param			Authorization	header		string	true	"Provide this header with `Bearer <user token>` to gain authorized access to this API method."
//	@Param			id				path		string	true	"The ID of the Account in the database."
//	@Success		200				{object}	entities.Relationship
//	@Router			/api/v1/accounts/:id/block [post]
func Block(c *gin.Context) {
	authorizationID := c.MustGet("authorizationID").(string)
	if authorizationID == "-" {
		errorJSON(c, http.StatusUnauthorized, errUnauthorized)
		return
	}
	id := c.Param("id")

	resp, err := accounts.Block(authorizationID, id)
	if err != nil {
		errorJSON(c, http.StatusNotFound, errNotFound)
		return
	}
	c.JSON(http.StatusOK, resp)
}

// Unblock godoc
//
//	@Summary		Lookup account ID from Webfinger address
//	@Description	Quickly lookup a username to see if it is available, skipping WebFinger resolution.
//	@Description	https://docs.joinmastodon.org/methods/accounts/#unblock
//	@Tags			mastodon/accounts
//	@Produce		json
//	@Param			Authorization	header		string	true	"Provide this header with `Bearer <user token>` to gain authorized access to this API method."
//	@Param			id				path		string	true	"The ID of the Account in the database."
//	@Success		200				{object}	entities.Relationship
//	@Router			/api/v1/accounts/:id/unblokc [post]
func Unblock(c *gin.Context) {
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

// Mute godoc
//
//	@Summary		Lookup account ID from Webfinger address
//	@Description	Quickly lookup a username to see if it is available, skipping WebFinger resolution.
//	@Description	https://docs.joinmastodon.org/methods/accounts/#unblock
//	@Tags			mastodon/accounts
//	@Produce		json
//	@Param			Authorization	header		string	true	"Provide this header with `Bearer <user token>` to gain authorized access to this API method."
//	@Param			id				path		string	true	"The ID of the Account in the database."
//
//	@Param			notifications	formData	boolean	false	"Mute notifications in addition to statuses? Defaults to true."
//	@Param			duration		formData	number	false	"How long the mute should last, in seconds. Defaults to 0 (indefinite)."
//
//	@Success		200				{object}	entities.Relationship
//	@Router			/api/v1/accounts/:id/mute [post]
func Mute(c *gin.Context) {
	authorizationID := c.MustGet("authorizationID").(string)
	if authorizationID == "-" {
		errorJSON(c, http.StatusUnauthorized, errUnauthorized)
		return
	}
	id := c.Param("id")
	form, err := c.MultipartForm() // todo should check type
	if err != nil {
		errorJSON(c, http.StatusBadRequest, newErrorMsg(err.Error()))
	}
	fe := formExactor{Form: form}
	notifications := defaultBool(fe.Bool("notifications"), true)
	duration := mustAtoi(defaultString(fe.String("duration"), "0"), 0)

	resp, err := accounts.Mute(authorizationID, id, notifications, duration)
	if err != nil {
		errorJSON(c, http.StatusNotFound, errNotFound)
		return
	}
	c.JSON(http.StatusOK, resp)
}

// Unmute godoc
//
//	@Summary		Lookup account ID from Webfinger address
//	@Description	Quickly lookup a username to see if it is available, skipping WebFinger resolution.
//	@Description	https://docs.joinmastodon.org/methods/accounts/#unblock
//	@Tags			mastodon/accounts
//	@Produce		json
//	@Param			Authorization	header		string	true	"Provide this header with `Bearer <user token>` to gain authorized access to this API method."
//	@Param			id				path		string	true	"The ID of the Account in the database."
//	@Success		200				{object}	entities.Relationship
//	@Router			/api/v1/accounts/:id/unblokc [post]
func Unmute(c *gin.Context) {
	authorizationID := c.MustGet("authorizationID").(string)
	if authorizationID == "-" {
		errorJSON(c, http.StatusUnauthorized, errUnauthorized)
		return
	}
	id := c.Param("id")

	resp, err := accounts.Unmute(authorizationID, id)
	if err != nil {
		errorJSON(c, http.StatusNotFound, errNotFound)
		return
	}
	c.JSON(http.StatusOK, resp)
}

// todo

// Note godoc
//
//	@Summary		Set private note on profile
//	@Description	Sets a private note on a user.
//	@Description	https://docs.joinmastodon.org/methods/accounts/#note
//	@Tags			mastodon/accounts
//	@Produce		json
//	@Param			Authorization	header		string	true	"Provide this header with `Bearer <user token>` to gain authorized access to this API method."
//	@Param			id				path		string	true	"The ID of the Account in the database."
//	@Param			comment			formData	string	false	"The comment to be set on that user. Provide an empty string or leave out this parameter to clear the currently set note."
//	@Success		200				array		entities.Relationship
//	@Failure		404				{object}	errorMsg
//	@Router			/api/v1/accounts/{id}/note [post]
func Note(c *gin.Context) {
	authorizationID := c.MustGet("authorizationID").(string)
	if authorizationID == "-" {
		errorJSON(c, http.StatusUnauthorized, errUnauthorized)
		return
	}
	id := c.Param("id")

	form, err := c.MultipartForm() // todo should check type
	if err != nil {
		errorJSON(c, http.StatusBadRequest, newErrorMsg(err.Error()))
	}
	fe := formExactor{Form: form}
	comment := fe.String("comment")

	relationship, err := accounts.Note(authorizationID, id, comment)
	if err != nil {
		errorJSON(c, http.StatusInternalServerError, newErrorMsg(err.Error()))
		return
	}
	c.JSON(http.StatusOK, relationship)

}

// Relationships godoc
//
//	@Summary		Check relationships to other accounts
//	@Description	Find out whether a given account is followed, blocked, muted, etc.
//	@Description	https://docs.joinmastodon.org/methods/accounts/#relationships
//	@Tags			mastodon/accounts
//	@Produce		json
//	@Param			id[]	query		array{string}	true	"The username or Webfinger address to lookup."
//	@Success		200		array		entities.Relationship
//	@Failure		404		{object}	errorMsg
//	@Router			/api/v1/accounts/relationships [get]
func Relationships(c *gin.Context) {
	authorizationID := c.MustGet("authorizationID").(string)
	if authorizationID == "-" {
		errorJSON(c, http.StatusUnauthorized, errUnauthorized)
		return
	}
	id := authorizationID
	idArray := c.QueryArray("id[]")

	relationships, err := accounts.Relationship(id, idArray)
	if err != nil {
		errorJSON(c, http.StatusInternalServerError, newErrorMsg(err.Error()))
		return
	}
	c.JSON(http.StatusOK, relationships)
}

//todo

// 这个很后面，放在最后，
// LookUp godoc
//
//	@Summary		Lookup account ID from Webfinger address
//	@Description	Quickly lookup a username to see if it is available, skipping WebFinger resolution.
//	@Description	https://docs.joinmastodon.org/methods/accounts/#lookup
//	@Tags			mastodon/accounts
//	@Produce		json
//	@Param			acct	query		string	true	"The username or Webfinger address to lookup."
//	@Success		200		{object}	entities.Account
//	@Failure		404		{object}	errorMsg
//	@Router			/api/v1/accounts/lookup [get]
func LookUp(c *gin.Context) {
	acct := c.Query("acct")
	resp, err := accounts.LookUp(acct)
	if err != nil {
		errorJSON(c, http.StatusNotFound, errNotFound)
		return
	}
	c.JSON(http.StatusOK, resp)
}
