package controller

import (
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

type Hash struct {
	*gin.Context
	m               map[string]any
	file            map[string][]*multipart.FileHeader
	authorizationID string
}

func HashBuilder(c *gin.Context) *Hash {
	hash := &Hash{
		Context:         c,
		m:               make(map[string]any),
		file:            make(map[string][]*multipart.FileHeader),
		authorizationID: "-",
	}
	return hash
}

func (h *Hash) Parse(typ string, key string, isArray bool) {
	switch typ {
	case "query":
		if isArray {
			v := h.QueryArray(key)
			h.m[key] = v
		} else {
			if v, e := h.GetQuery(key); e {
				h.m[key] = v
			}
		}
	case "x-www-form-urlencoded", "body":
		if isArray {
			v := h.PostFormArray(key)
			h.m[key] = v
		} else {
			if v, e := h.GetPostForm(key); e {
				h.m[key] = v
			}
		}
	case "header":
		h.m[key] = h.GetHeader(key)
	case "multiple/form-data", "mpfd":
		form, _ := h.MultipartForm()
		for k, v := range form.Value {
			h.m[k] = v
		}
		h.file = form.File
	}
}

// get params
func (h *Hash) File(key string) []*multipart.FileHeader {
	return h.file[key]
}

func (h *Hash) Get(key string) (v any, e bool) {
	v, e = h.string(key)
	if e {
		return v, e
	}
	v, e = h.stringArray(key)
	if e {
		return v, e
	}
	return nil, false
}

func (h *Hash) String(key string) (string, bool) {
	s, e := h.string(key)
	if e {
		return s, e
	}
	sa, e := h.stringArray(key)
	if e && len(sa) > 0 {
		return sa[0], e
	}
	return "", false
}

func (h *Hash) string(key string) (string, bool) {
	v, e := h.m[key]
	if e {
		s, e := v.(string)
		if e {
			return s, e
		}
	}
	return "", false
}

func (h *Hash) StringArray(key string) ([]string, bool) {
	v, e := h.m[key]
	if e {
		sa, e := v.([]string)
		if e {
			return sa, e
		}
	}
	return nil, false
}

func (h *Hash) stringArray(key string) ([]string, bool) {
	v, e := h.m[key]
	if e {
		sa, e := v.([]string)
		if e {
			return sa, e
		}
	}
	return nil, false
}

func (h *Hash) SetAuthorizationID(authorizationID string) {
	h.authorizationID = authorizationID
}

func (h *Hash) AuthorizationID(authorizationID string) string {
	return h.authorizationID
}
