package controller

import (
	"fmt"
	"io"
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
)

// Ping godoc
//
//	@Summary		For test
//	@Description	return http request
//	@Tags
//	@Produce	plain
//	@Response	200	{body}	string	"return the http request received"
//	@Router		/echo [get]
func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

// Echo godoc
//
//	@Summary		For test
//	@Description	return http request
//	@Tags
//	@Produce	plain
//	@Response	200	{body}	string	"return the http request received"
//	@Router		/echo [get]
//	@Router		/echo [post]
//	@Router		/echo [delete]
//	@Router		/echo [put]
//	@Router		/echo [patch]
func Echo(c *gin.Context) {
	fullpath := c.FullPath()

	header := c.Request.Header
	keys := make([]string, 0, len(header))
	for k := range header {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	headerString := ""
	for _, k := range keys {
		for _, value := range header[k] {
			headerString = headerString + fmt.Sprintln(k, value) + "\n"
		}
	}

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	resp := fullpath + "\n\n" + headerString + "\n\n" + string(body)
	c.String(http.StatusOK, resp)

}

// Test godoc

//	@Summary		For test
//	@Description	test form params
//	@Tags
//	@Accept	multipart/form-data

//	@Produce	plain
//	@param		something[somekey]	formData	string	false	"test"
//	@param		display_name		formData	string	false	"The display name to use for the profile."
//	@param		note				formData	string	false	"The account bio."
//	@param		avatar				formData	string	false	"Avatar image encoded using multipart/form-data"
//	@param		header				formData	string	false	"Header image encoded using multipart/form-data"
//	@param		locked				formData	boolean	false	"Whether manual approval of follow requests is required."
//	@param		bot					formData	boolean	false	"Whether the account has a bot flag."
//	@param		discoverable		formData	boolean	false	"Whether the account should be shown in the profile directory."
//	@param		pic					formData	any		false	"Whether the account should be shown in the profile directory."
//
//	@Response	200					{body}		string	"return the http request received"
//	@Router		/test [post]
func Test(c *gin.Context) {
	fullpath := c.FullPath()

	header := c.Request.Header
	keys := make([]string, 0, len(header))
	for k := range header {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	headerString := ""
	for _, k := range keys {
		for _, value := range header[k] {
			headerString = headerString + fmt.Sprintln(k, value) + "\n"
		}
	}

	formData := make(map[string]string)
	c.ShouldBind(&formData)
	fmt.Println(formData)

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	resp := fullpath + "\n------------------------------\n" + headerString + "\n------------------------------\n" + string(body)
	c.String(http.StatusOK, resp)

}
