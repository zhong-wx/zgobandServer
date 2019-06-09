package models

import (
	"encoding/json"
	"fmt"
	"testing"
	"../vo"
)

func TestSaveGame(t *testing.T) {
	Open()
	db.LogMode(true)
	p := vo.Pos{Row:3, Column:2}
	js, err := json.Marshal(p)
	if err != nil {
		t.Error(err.Error())
	}

	fmt.Println(js)
	err,_ = SaveGame("default", string(js), "test", "test1", 1)
	if err != nil {
		t.Error(err.Error())
	}
	Close()
}
