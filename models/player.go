package models

import "fmt"

type Player struct {
	Account string	`gorm:"primary_key"`
	Password string	`gorm:"column:password"`
	Nickname string	`gorm:"column:nickname"`
	Score int		`gorm:"column:score"`
	Winround int	`gorm:"column:winRound"`
	Loseround int	`gorm:"column:loseRound"`
	Drawround int	`gorm:"column:drawRound"`
	Escaperound int	`gorm:"column:escapeRound"`
}

func (p *Player) Display() {
	fmt.Println("account:", p.Account, " nickname:", p.Nickname, " winround:", p.Winround)
}

func CreatePlayer(account string, password string, nickname string) error {
	player := Player{Account: account, Password: password, Nickname:nickname}
	db.NewRecord(player)
	err := db.Create(&player).Error
	return err
}

func QueryPlayer(account string) *Player {
	player := Player{}
	err := db.Where("account = ?", account).First(&player).Error
	if(err != nil) {
		fmt.Println("db.First error:", err.Error())
		return nil
	}

	return &player
}

func WinUpdate(account string) error {
	player := QueryPlayer(account)
	if player == nil {
		return fmt.Errorf("cannot find the account:%s", account)
	}
	player.Winround += 1
	player.Score += 1
	err := db.Save(&player).Error
	if err != nil {
		fmt.Println("db.Save error:", err.Error())
		return err
	}
	return err
}

func LoseUpdate(account string) error {
	player := QueryPlayer(account)
	if player == nil {
		return fmt.Errorf("cannot find the account:%s", account)
	}
	player.Loseround += 1
	player.Score -= 1
	err := db.Save(&player).Error
	if err != nil {
		fmt.Println("db.Save error:", err.Error())
		return err
	}
	return err
}

func EscapeUpdate(account string) error {
	player := QueryPlayer(account)
	if player == nil {
		return fmt.Errorf("cannot find the account:%s", account)
	}
	player.Escaperound += 1
	player.Loseround += 1
	player.Score -= 5
	err := db.Save(&player).Error
	if err != nil {
		fmt.Println("db.Save error:", err.Error())
		return err
	}
	return err
}

func DrawUpdate(account string)  error {
	player := QueryPlayer(account)
	if player == nil {
		return fmt.Errorf("cannot find the account:%s", account)
	}
	player.Drawround += 1
	err := db.Save(&player).Error
	if err != nil {
		fmt.Println("db.Save error:", err.Error())
		return err
	}
	return err
}