// TODO

package mastodon

import (
	"log"
	"mime/multipart"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/moonchan-xyz/moonchan/back/mastodon/entities"
)

type errorMsg struct {
	Error string `json:"error"`
}

var errUnauthorized = &errorMsg{"The access token is invalid"}                       // 401
var errNotFound = &errorMsg{"Record not found"}                                      // 404
var errUnprocessableEntity = &errorMsg{"This method requires an authenticated user"} // 422
var errServiceUnavailable = &errorMsg{"Remote data could not be fetched"}            // 503

// user defined
func newErrorMsg(err string) *errorMsg {
	return &errorMsg{err}
}

func errorJSON(c *gin.Context, statusType int, err *errorMsg) {
	c.JSON(statusType, err)
}

// for reading form
type formExactor struct {
	*multipart.Form
}

func (fe *formExactor) String(key string) *string {
	ps := extractStringFromFormData(key, fe.Form)
	return ps
}
func (fe *formExactor) DefaultString(key string, d string) string {
	ps := extractStringFromFormData(key, fe.Form)
	s := defaultString(ps, d)
	return s
}

func (fe *formExactor) Bool(key string) *bool {
	ps := extractStringFromFormData(key, fe.Form)
	pb := extractBoolPtrFromStringPtr(ps)
	return pb
}
func (fe *formExactor) DefaultBool(key string, d bool) bool {
	ps := extractStringFromFormData(key, fe.Form)
	pb := extractBoolPtrFromStringPtr(ps)
	b := defaultBool(pb, d)
	return b
}

func (fe *formExactor) DefaultInt(key string, d int) int {
	ps := extractStringFromFormData(key, fe.Form)
	if ps == nil {
		return d
	}
	i := mustAtoi(*ps, d)
	return i
}

func (fe *formExactor) Hash(key string) entities.Hash {
	h := make(entities.Hash)
	for i := 0; i < mustAtoi(os.Getenv("MAX_HASH"), -1); i++ {
		si := strconv.Itoa(i)
		nameKey := key + "[" + si + "][name]"
		name := fe.String(nameKey)
		if name == nil {
			break
		}
		valueKey := key + "[" + si + "][value]"
		value := fe.String(valueKey)
		if value == nil {
			break
		}
		keyStr := si
		h.SetNV(keyStr, *name, *value)
	}
	return h
}

func (fe *formExactor) StringArray(key string) []string {
	arr := extractStringArrayFromFormData(key, fe.Form)
	return arr
}

func (fe *formExactor) URLString(key string) *entities.URLstring {
	// todo
	// 要做上传的
	files := fe.Form.File[key]
	for _, file := range files {
		src, err := file.Open()
		if err != nil {
			log.Println(err.Error())
			continue
		}
		_ = src // todo
	}
	return nil
}

func mustAtoi(s string, fallback int) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		return fallback
	}
	return n
}

func extractStringFromFormData(name string, form *multipart.Form) *string {
	if values, ok := form.Value[name]; !ok {
		return nil
	} else {
		for _, v := range values {
			var s string = v
			return &s
		}
	}
	return nil
}
func extractStringArrayFromFormData(name string, form *multipart.Form) []string {
	if values, ok := form.Value[name]; !ok {
		return nil
	} else {
		return values
	}
}
func extractBoolPtrFromStringPtr(ps *string) *bool {
	if ps == nil {
		return nil
	}
	r := isTrue(*ps)
	return &r
}

func isTrue(s string) bool {
	r := s == "true" || s == "1"
	return r
}

func defaultString(ps *string, d string) string {
	if ps == nil {
		return d
	}
	return *ps
}
func defaultBool(pb *bool, d bool) bool {
	if pb == nil {
		return d
	}
	return *pb
}
