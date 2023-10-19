package webfinger

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/moonchan-xyz/moonchan/back/db"
	"github.com/moonchan-xyz/moonchan/back/utils"
	"github.com/moonchan-xyz/moonchan/back/webfinger"
	"gorm.io/gorm"
)

// Webfinger godoc
//	@Summary	webfinger
//	@Description
//	@Tags		webfinger
//	@Produce	json
//	@Param		resource	query		string	true	"search webfinger object by account"
//	@Success	200			{object}	object{subject=string,links=[]any}
//	@Failure	404			{object}	object{error=string}
//	@Failure	422			{object}	object{error=string}
//	@Failure	500			{object}	object{error=string}
//	@Router		/.well-known/webfinger [get]
func Webfinger(c *gin.Context) {
	// 被访问
	// 取得query
	resource := c.Query("resource")
	if !strings.HasPrefix(resource, "acct:") {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "format error"})
		return
	}
	// 取出成为 uesrname 和 host
	acct := strings.TrimPrefix(resource, "acct:")
	username, host, err := utils.ParseUserAndHost(acct)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "format error"})
		return
	}
	// 判断是否存在 not used
	_, err = db.ReadAccountByAcct(acct)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// 形成返回数据
	webfingerObj, err := webfinger.CreateWebfingerObj(username, host)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, webfingerObj)
}
