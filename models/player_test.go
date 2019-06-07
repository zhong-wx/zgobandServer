package models

import "testing"

func TestCreatePlayer(t *testing.T) {
	db.LogMode(true)
	Open()
	ret  := CreatePlayer("test1", "test", "test")
	if(ret != nil) {
		t.Error(ret.Error())
	}
	Close()
}

func TestQueryPlayer(t *testing.T) {
	Open()
	db.LogMode(true)
	ret := QueryPlayer("test")
	if ret == nil {
		t.Error("ret = nil")
	}
	ret.Display()
	Close()
}