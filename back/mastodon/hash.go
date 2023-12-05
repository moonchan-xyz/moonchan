package mastodon

import "mime/multipart"

type Hash interface {
	File(key string) []*multipart.FileHeader
	Get(key string) (v any, e bool)
	String(key string) (string, bool)
	StringArray(key string) ([]string, bool)
	AuthorizationID(authorizationID string) string
}
