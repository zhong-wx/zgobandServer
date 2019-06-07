package models

import "fmt"

type Desk struct {
	Deskid int32  `gorm:"column:deskID;primary_key"`
	Player1 string
	Player2 string
	Ready1 []uint8
	Ready2 []uint8
}

func QueryAllDesk() ([]Desk, error) {
	var desks []Desk
	db = db.Find(&desks)
	err := db.Error
	return desks, err
}

func AddSeat(deskID int32, seatID int32, account string) (bool, error) {
	tx := db.Begin()
	desk := &Desk{}
	tdb := tx.Set("gorm:query_option", "FOR UPDATE").First(desk, deskID)
	err := tdb.Error
	if err != nil && err.Error() != "record not found" {
		tx.Commit()
		return false, err
	}

	if(seatID == 1 && desk.Player1 != "") {
		tx.Commit()
		return false, nil
	}
	if(seatID == 2 && desk.Player2 != "") {
		tx.Commit()
		return false, nil
	}

	// set seatID
	if(seatID == 1) {
		desk.Player1 = account
	} else if(seatID == 2) {
		desk.Player2 = account
	} else {
		fmt.Println("updateSeatError: unkown seatID")
		tx.Commit()
		return false, err
	}

	// save or create
	if err == nil {
		err = tx.Save(desk).Error
	} else {
		desk.Ready2 = append(desk.Ready2, 0)
		desk.Ready1 = append(desk.Ready1, 0)
		desk.Deskid = deskID
		err = tx.Create(desk).Error
	}

	tx.Commit()
	if err != nil {
		return false, err
	}
	return true, nil
}

func DelSeat(deskID int32, seatID int32) error {
	var err error
	var b []uint8
	b = append(b, 0)
	if seatID == 1 {
		err = db.Table("desks").Where("deskID = ?", deskID).Updates(map[string]interface{}{"player1":"", "ready1":b}).Error
	} else if seatID == 2 {
		err = db.Table("desks").Where("deskID = ?", deskID).Updates(map[string]interface{}{"player2":"", "ready2":b}).Error
	} else {
		fmt.Println("DelSeat unknow seatID")
		return  fmt.Errorf("DelSeat unknow seatID")
	}

	if(err == nil) {
		return  nil
	}
	return  err
}

func SetReady(deskID int32, seatID int32, isReady bool) error {
	var err error = nil
	var ready []uint8
	if isReady {
		ready = append(ready, 1)
	} else {
		ready = append(ready, 0)
	}
	if(seatID == 1) {
		err = db.Table("desks").Where("deskID = ?", deskID).Update("ready1", ready).Error
	} else if(seatID == 2) {
		err = db.Table("desks").Where("deskID = ?", deskID).Update("ready2", ready).Error
	} else {
		fmt.Println("DelSeat unknow seatID")
		return  fmt.Errorf("DelSeat unknow seatID")
	}
	return err
}

func GetSeatAccountInfo(deskID int32, seatID int32) (string, bool, error) {
	desk, err := GetDesk(deskID)
	if err != nil {
		return "", false, err
	}
	ready1 := desk.Ready1[0] == 1
	ready2 := desk.Ready2[0] == 1
	switch seatID {
	case 1: return desk.Player1, ready1, nil
	case 2: return desk.Player2, ready2, nil
	}
	fmt.Println("getSeatAccountInfo unknow seatID")
	return "", false, fmt.Errorf("getSeatAccountInfo unknow seatID")
}

func GetDesk(deskID int32) (*Desk, error) {
	desk := &Desk{}
	err := db.First(desk, deskID).Error
	if err != nil {
		return desk, err
	}
	return desk, nil
}