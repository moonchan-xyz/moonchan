package entities

import (
	"log"
	"testing"
)

func TestHash(t *testing.T) {
	key := "key"
	name := "name"
	value := "value"
	h := make(Hash)
	log.Println(h)
	// h[key] = NameValuePair{name, value}
	h.SetNV(key, name, value) // no  problem
	log.Println(h)
}
