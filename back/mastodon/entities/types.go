package entities

import (
	"time"
)

type URLstring string
type HTMLstring string
type DATEstring string

func (date *DATEstring) FromTimestamp(ms int64) {
	t := time.Unix(ms/1e3, ms%1e3*1e6).UTC()
	s := t.Format(time.RFC3339)
	*date = DATEstring(s)
}

func (date *DATEstring) ToTimestamp() (int64, error) {
	timestamp, err := time.Parse(time.RFC3339, string(*date))
	if err != nil {
		return 0, err
	}
	return timestamp.UnixMilli(), err
}

type NameValuePair struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func (nv *NameValuePair) NV() (string, string) {
	return nv.Name, nv.Value
}

type Hash map[string]NameValuePair

func (h Hash) SetNV(key, name, value string) {
	h[key] = NameValuePair{name, value}
}
