package entities

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMarshal(t *testing.T) {
	a := &Source{
		"123123131",
		"12",
		false,
		"",
		"",
		[]Field{{"name", "what", nil}},
		999,
	}
	b, e := json.Marshal(a)
	if e != nil {
		t.Error(e)
	}
	fmt.Printf("%s", b)
}
