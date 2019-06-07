package models

import (
	"fmt"
	"testing"
)

func TestQueryAllDesk(t *testing.T) {
	Open()
	db.LogMode(true)
	desks, err:= QueryAllDesk()
	if(err != nil) {
		t.Error(err.Error())
	}
	for i:=1; i< len(desks); i++ {
		fmt.Println(desks[i].Deskid, desks[i].Player1, desks[i].Player2, desks[i].Ready1[0], desks[i].Ready2[0])
	}
	Close()
}

func TestAddSeat(t *testing.T) {
	Open()
	db.LogMode(true)
	succ, err := AddSeat(5, 2, "tet")
	if err != nil {
		t.Error(succ, err.Error())
	}
	Close()
}

func TestDelSeat(t *testing.T) {
	Open()
	db.LogMode(true)
	err := DelSeat(1, 1)
	if err != nil {
		t.Error(err.Error())
	}
	Close()
}

func TestSetReady(t *testing.T) {
	Open()
	db.LogMode(true)
	err := SetReady(2, 1, false)
	if err != nil {
		t.Error(err.Error())
	}

	err = SetReady(3, 1, true)
	if err != nil {
		t.Error(err.Error())
	}
	Close()
}

func TestGetSeatAccountInfo(t *testing.T) {
	Open()
	db.LogMode(true)
	_, str, err := GetSeatAccountInfo(2, 1)
	if err != nil {
		t.Error(err.Error())
	} else {
		fmt.Println(str)
	}
	Close()
}