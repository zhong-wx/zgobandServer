package models

import "fmt"

type PlayerStat struct {
	Account string	`gorm:"primary_key"`
	Stat int
	DeskID int	`gorm:"column:deskID"`
	SeatID int	`gorm:"column:seatID"`
}

func (PlayerStat) TableName() string {
	return "playerstat"
}

func CreateOrUpdatePlayerStat(account string, stat int, deskID int, seatID int) error {
	playerStat := PlayerStat{Account: account, Stat: stat, DeskID:deskID, SeatID:seatID}
	db.NewRecord(playerStat)
	err := db.Create(&playerStat).Error
	if err == nil {
		return nil
	}

	fmt.Println(err.Error())
	err = db.Save(&playerStat).Error
	if err != nil {
		fmt.Println(err.Error())
	}

	return err
}

func DeletePlayerStat(account string) error {
	p := PlayerStat{Account:account}
	return db.Delete(&p).Error
}

func QueryPlayerStat(account string) (*PlayerStat, error) {
	p := PlayerStat{Account:account}
	err := db.First(&p).Error
	if err != nil {
		return nil, err
	}
	return &p, nil
}