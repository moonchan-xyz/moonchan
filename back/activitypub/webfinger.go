package activitypub

import (
	"net/http"

	"github.com/iancoleman/orderedmap"
	"github.com/moonchan-xyz/moonchan/back/utils"
)

func GetWebfingerObj(acct string) (webfingerObj *orderedmap.OrderedMap, err error) {
	username, host, err := utils.ParseUserAndHost(acct)
	if err != nil {
		return nil, err
	}
	url := utils.ParseWebfingerUrl(username, host)
	webfingerObj, err = FetchOrderedMap(http.MethodGet, url, nil)
	return
}
