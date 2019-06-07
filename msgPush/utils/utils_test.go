package utils

import (
	"fmt"
	"testing"
)

func TestCreatMessage(t *testing.T) {
	m := map[string]interface{}{}
	m["1"] = 1
	m["2"] = "2"
	m["3"] = map[string]interface{}{"asdf":234}
	bs, err := CreatMessage(m)
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Println(bs)
	fmt.Println(string(bs))
}
