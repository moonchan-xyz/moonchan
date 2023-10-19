package activitypub

import (
	"bytes"
	"crypto"
	"encoding/json"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"

	"github.com/iancoleman/orderedmap"
)

var jar, _ = cookiejar.New(nil)
var proxyURL, _ = url.Parse("http://localhost:10809")

// use this in formal version
// var client = &http.Client{
// // add this to disable redirect, do not use it
// 	CheckRedirect: func(req *http.Request, via []*http.Request) error {
// 		return http.ErrUseLastResponse
// 	},
// 	Jar: jar,
// }

var client = &http.Client{
	Transport: &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	},
	Jar: jar,
}

func Do(req *http.Request) (*http.Response, error) {
	return client.Do(req)
}

func NewSingedRequest(
	privateKey crypto.PrivateKey, pubKeyID string,
	method, url string, body []byte,
) (r *http.Request, err error) {
	r, err = http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		return
	}

	r.Header.Set("host", r.URL.Host)
	r.Header.Set("date", time.Now().UTC().Format(http.TimeFormat))
	r.Header.Set("content-type", "application/activity+json")

	err = Sign(privateKey, pubKeyID, r, body)
	if err != nil {
		return
	}

	return
}

func FetchOrderedMap(method, url string, body io.Reader, headers ...map[string][]string) (o *orderedmap.OrderedMap, err error) {
	r, err := http.NewRequest(method, url, body)
	if err != nil {
		return
	}

	r.Header.Set("Accept", "application/ld+json")
	r.Header.Set("Accept", "application/activity+json")

	for _, h := range headers {
		for k, va := range h {
			for _, v := range va {
				r.Header.Set(k, v)
			}
		}
	}

	resp, err := Do(r)
	if err != nil {
		return
	}

	// need review
	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	o = orderedmap.New()
	err = json.Unmarshal(respData, &o)
	if err != nil {
		return
	}

	return
}
