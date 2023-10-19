package utils

import "github.com/iancoleman/orderedmap"

type KV struct {
	K string
	V any
}

func (kv *KV) KV() (string, any) {
	return kv.K, kv.V
}

func OrderedMapByKVList(kvlist []KV) (o *orderedmap.OrderedMap) {
	o = orderedmap.New()
	for _, kv := range kvlist {
		o.Set(kv.KV())
	}
	return
}

// utils
func OrderedMap(keys []string, values []interface{}) *orderedmap.OrderedMap {
	o := orderedmap.New()
	for i, key := range keys {
		o.Set(key, values[i])
	}
	return o
}

func IDType(id, typestr string) *orderedmap.OrderedMap {
	return OrderedMap([]string{"@id", "@type"}, []any{id, typestr})
}
