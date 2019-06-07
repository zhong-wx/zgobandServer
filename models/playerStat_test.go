package models

import (
	"fmt"
	"testing"
)

func TestCreateOrUpdatePlayerStat(t *testing.T) {
	Open()
	db.LogMode(true)
	CreateOrUpdatePlayerStat("test", 1, 1, 2)
	CreateOrUpdatePlayerStat("test", 2, 2, 1)
	Close()
}

func TestQueryPlayerStat(t *testing.T) {
	Open()
	db.LogMode(true)
	p, err := QueryPlayerStat("test")
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Println(p)
	Close()
}

func TestDeletePlayerStat(t *testing.T) {
	Open()
	db.LogMode(true)
	err := DeletePlayerStat("test")
	t.Error(err.Error())
	Close()
}