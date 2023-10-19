package webfinger

import (
	"errors"
	"fmt"

	"github.com/iancoleman/orderedmap"
	"github.com/moonchan-xyz/moonchan/back/utils"
)

func CreateWebfingerObj(username, host string) (o *orderedmap.OrderedMap, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%s", e)
		}
	}()
	o = orderedmap.New()
	o.Set("subject", "acct:"+username+"@"+host)
	o.Set("links", []*orderedmap.OrderedMap{
		utils.OrderedMapByKVList([]utils.KV{
			{K: "rel", V: "self"},
			{K: "type", V: "application/activity+json"},
			{K: "href", V: "https://" + host + "/users/" + username},
		}),
		utils.OrderedMapByKVList([]utils.KV{
			{K: "rel", V: "http://webfinger.net/rel/profile-page"},
			{K: "type", V: "text/html"},
			{K: "href", V: "https://" + host + "/@" + username},
		}),
		utils.OrderedMapByKVList([]utils.KV{ // dunno what it is
			{K: "rel", V: "http://ostatus.org/schema/1.0/subscribe"},
			{K: "template", V: "https://" + host + "/authorize-follow?acct={uri}"},
		}),
	})
	return
}

// Parse UserID from webfinger
func ParseUserID(webfingerObj *orderedmap.OrderedMap) (userID string, err error) {
	if links, ok := webfingerObj.Get("links"); !ok {
		err = (errors.New("notfound links"))
		return
	} else {
		if arr, ok := links.([]any); !ok {
			err = (errors.New("links not array"))
			return
		} else {
			for _, li := range arr {
				if lo, ok := li.(orderedmap.OrderedMap); ok {
					key, ok := lo.Get("rel")
					if ok && key == "self" {
						id, ok := lo.Get("href")
						if !ok {
							err = (errors.New("self link doesn't have href"))
							return
						} else {
							userID = id.(string)
							return
						}
					}
				}
			}
		}
	}
	err = (errors.New("id not found"))
	return // never
}
