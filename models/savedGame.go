package models

import "time"

type SavedGame struct {
	Id int			`gorm:"column:id;AUTO_INCREMENT"`
	GameName string	`gorm:"column:gameName"`
	SaveTime time.Time 		`gorm:"column:saveTime"`
	Record	string	`gorm:"column:record"`
	Account string 	`gorm:"column:account"`
	OtherSide string `gorm:"column:otherSide"`
	SeatID int		`gorm:"column:seatID"`
}

func (SavedGame) TableName() string {
	return "savedgame"
}

func SaveGame(gameName string, record string, account string, otherSide string, seatID int) (error, int) {
	savedGame := &SavedGame{GameName:gameName, SaveTime:time.Now(), Record:record, Account:account, OtherSide:otherSide, SeatID:seatID}
	err := db.Create(savedGame).Error
	if err != nil {
		return err, -1
	}
	return nil, savedGame.Id
}

func GetSaveGame(id int) *SavedGame {
	savedGame := &SavedGame{Id:id}
	err := db.First(savedGame).Error
	if err != nil {
		return nil
	}
	return savedGame
}

func GetAllSavedGame(account string) []SavedGame {
	var savedGameList []SavedGame
	db.Where("account = ?", account).Find(&savedGameList)
	return savedGameList
}

func DelSavedGame(id int) {
	savedGame := SavedGame{Id:id}
	db.Delete(&savedGame)
}